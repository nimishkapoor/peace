package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func GenerateJWT(username string) (string, error) {
	var mySigningKey = []byte("akakakakakakakak")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 300).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

const (
	DB_DSN = "postgres://postgres:123@localhost/postgres?sslmode=disable"
)

var db *sql.DB

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Pswd     string `json:"password"`
}

type Error struct {
	Msg string `json:"message"`
}

type Token struct {
	TokenString string `json:"token"`
}

type RequestBody struct {
	TokenString string `json:"token"`
	Subject     string `json:"subject"`
	Opt1        string `json:"opt1"`
	Opt2        string `json:"opt2"`
	Opt3        string `json:"opt3"`
	Opt4        string `json:"opt4"`
}

type VoteBody struct {
	TokenString string `json:"token"`
	OptId       int    `json:"optid"`
}

type Opt struct {
	Value  string `json:"value"`
	OptId  int    `json:"optid"`
	PostId int    `json:"postid"`
	Count  int    `json:"count"`
}

type Poll struct {
	Id        int    `json:"id"`
	UserName  string `json:"userName"`
	Subject   string `json:"subject"`
	Options   []*Opt `json:"options"`
	Voted     int    `json:"voted"`
	VotedByMe int    `json:"votedbyme"`
}

type Vote struct {
	Value string `json:"value"`
	Count int    `json:"count"`
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)
	var myUser User
	userSql := "Insert INTO users (user_name, pswd) values ($1, $2) RETURNING user_id"

	var err = db.QueryRow(userSql, user.UserName, user.Pswd).Scan(&myUser.Id)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		error_ := Error{
			Msg: "User Name already exists!",
		}
		json.NewEncoder(w).Encode(error_)
	} else {
		fmt.Printf("UserID %d, created!\n", myUser.Id)
		json.NewEncoder(w).Encode(user)
	}
}
func authLogin(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)

	var myUser User
	userSql := "SELECT pswd FROM users WHERE user_name = $1"

	var err = db.QueryRow(userSql, user.UserName).Scan(&myUser.Pswd)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		error_ := Error{
			Msg: "User does not exists!",
		}
		json.NewEncoder(w).Encode(error_)
	} else if user.Pswd != myUser.Pswd {
		error_ := Error{
			Msg: "Incorrect Password!",
		}
		json.NewEncoder(w).Encode(error_)
	} else {
		validToken, _ := GenerateJWT(user.UserName)
		var token Token
		token.TokenString = validToken
		json.NewEncoder(w).Encode(token)
	}
}

func getPolls(w http.ResponseWriter, r *http.Request) {

	var token_c Token
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &token_c)

	var mySigningKey = []byte("akakakakakakakak")

	token, err := jwt.Parse(token_c.TokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return mySigningKey, nil
	})

	if err != nil {
		error_ := Error{
			Msg: "Token Expired!",
		}
		json.NewEncoder(w).Encode(error_)
		return
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	polls := make([]*Poll, 0)
	pollsSql := "SELECT * FROM post"

	rows, err := db.Query(pollsSql)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		poll := new(Poll)
		if err := rows.Scan(&poll.Id, &poll.Subject, &poll.UserName); err != nil {
			panic(err)
		}
		options := make([]*Opt, 0)
		optSql := "SELECT value, post_id, opt_id, count FROM opt WHERE post_id = $1 ORDER BY opt_id"
		optrows, opterr := db.Query(optSql, poll.Id)
		if opterr != nil {
			log.Println(opterr)
		}
		for optrows.Next() {
			opt := new(Opt)
			if opterr := optrows.Scan(&opt.Value, &opt.PostId, &opt.OptId, &opt.Count); opterr != nil {
				panic(opterr)
			}
			options = append(options, opt)
		}
		poll.Options = options
		existSql := "SELECT vote_id, opt_id FROM vote WHERE post_id = $1 AND user_name = $2"

		var voteid int
		var votedopt int
		_ = db.QueryRow(existSql, poll.Id, claims["username"]).Scan(&voteid, &votedopt)
		if voteid != 0 {
			poll.Voted = 1
			poll.VotedByMe = votedopt
		} else {
			poll.Voted = 0
			poll.VotedByMe = 0
		}
		polls = append(polls, poll)
	}

	if err := rows.Err(); err != nil {
		log.Println("Failed to execute query: ", err)
		error_ := Error{
			Msg: "Internal Server!",
		}
		json.NewEncoder(w).Encode(error_)
	} else {
		fmt.Printf("Feeding Polls!\n")
		json.NewEncoder(w).Encode(polls)
	}
}

func createPoll(w http.ResponseWriter, r *http.Request) {

	var body RequestBody
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &body)

	var mySigningKey = []byte("akakakakakakakak")

	token, err := jwt.Parse(body.TokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return mySigningKey, nil
	})

	if err != nil {
		error_ := Error{
			Msg: "Token Expired!",
		}
		json.NewEncoder(w).Encode(error_)
		return
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	pollsSql := "INSERT INTO post (sub, user_name) VALUES ($1, $2) RETURNING post_id"

	var postid int
	var err_ = db.QueryRow(pollsSql, body.Subject, claims["username"]).Scan(&postid)

	if err_ != nil {
		log.Println("Failed to execute query: ", err_)
		error_ := Error{
			Msg: "Internal Server Error!",
		}
		json.NewEncoder(w).Encode(error_)
	}

	optSql := "INSERT INTO opt (value, post_id) VALUES ($1, $2)"

	_ = db.QueryRow(optSql, body.Opt1, postid)
	_ = db.QueryRow(optSql, body.Opt2, postid)
	_ = db.QueryRow(optSql, body.Opt3, postid)
	_ = db.QueryRow(optSql, body.Opt4, postid)

	fmt.Printf("Poll Created!\n")
	json.NewEncoder(w).Encode("Worked!")
}

func vote(w http.ResponseWriter, r *http.Request) {

	var body VoteBody
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &body)

	var mySigningKey = []byte("akakakakakakakak")

	token, err := jwt.Parse(body.TokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return mySigningKey, nil
	})

	if err != nil {
		error_ := Error{
			Msg: "Token Expired!",
		}
		json.NewEncoder(w).Encode(error_)
		return
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	var postid_ int
	_ = db.QueryRow("SELECT post_id FROM opt WHERE opt_id = $1", body.OptId).Scan(&postid_)

	existSql := "SELECT vote_id FROM vote WHERE post_id = $1 AND user_name = $2"

	var voteid int
	_ = db.QueryRow(existSql, postid_, claims["username"]).Scan(&voteid)

	if voteid != 0 {
		error_ := Error{
			Msg: "Cannot Vote Again!",
		}
		json.NewEncoder(w).Encode(error_)
		return
	}

	pollsSql := "INSERT INTO vote (opt_id, user_name, post_id) VALUES ($1, $2, $3)"

	_, err_ := db.Query(pollsSql, body.OptId, claims["username"], postid_)

	if err_ != nil {
		log.Println("Failed to execute query: 1", err_)
		error_ := Error{
			Msg: "Internal Server Error!",
		}
		json.NewEncoder(w).Encode(error_)
		return
	}

	var postid int
	optSql := "UPDATE opt SET count = count + 1 WHERE opt_id = $1 RETURNING post_id"

	_ = db.QueryRow(optSql, body.OptId).Scan(&postid)

	votes := make([]*Vote, 0)
	voteSql := "SELECT count, value FROM Opt WHERE post_id = $1 ORDER BY opt_id"
	voterows, voteerr := db.Query(voteSql, postid)
	if voteerr != nil {
		log.Println(voteerr)
	}
	defer voterows.Close()
	for voterows.Next() {
		vote := new(Vote)
		if voteerr := voterows.Scan(&vote.Count, &vote.Value); voteerr != nil {
			panic(voteerr)
		}
		votes = append(votes, vote)
	}
	fmt.Printf("Vote Created!\n")
	json.NewEncoder(w).Encode(votes)
}

func votedBy(w http.ResponseWriter, r *http.Request) {

	var body VoteBody
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &body)

	var mySigningKey = []byte("akakakakakakakak")

	_, err := jwt.Parse(body.TokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return mySigningKey, nil
	})

	if err != nil {
		error_ := Error{
			Msg: "Token Expired!",
		}
		json.NewEncoder(w).Encode(error_)
		return
	}

	voters := make([]string, 0)
	voteSql := "SELECT user_name FROM vote WHERE opt_id = $1"
	voterows, voteerr := db.Query(voteSql, body.OptId)
	if voteerr != nil {
		log.Println(voteerr)
	}
	for voterows.Next() {
		var voter string
		if voteerr := voterows.Scan(&voter); voteerr != nil {
			panic(voteerr)
		}
		voters = append(voters, voter)
	}
	fmt.Printf("Feeding Voters!\n")
	json.NewEncoder(w).Encode(voters)
}
func handleRequests() {
	log.Println("Starting development server at http://127.0.0.1:10000/")
	log.Println("Quit the server with CONTROL-C.")

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/new-user", createNewUser).Methods("POST")
	myRouter.HandleFunc("/login", authLogin).Methods("POST")
	myRouter.HandleFunc("/polls", getPolls).Methods("POST")
	myRouter.HandleFunc("/create-poll", createPoll).Methods("POST")
	myRouter.HandleFunc("/vote", vote).Methods("POST")
	myRouter.HandleFunc("/voted-by", votedBy).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(myRouter)))
}

func main() {
	var err error
	db, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}

	handleRequests()
}
