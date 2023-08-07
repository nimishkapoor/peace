package kapt

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type (
	CreateNewUserRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	CreateNewUserResponse struct {
		Msg string `json:"message"`
	}
	AuthLoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	AuthLoginResponse struct {
		TokenString string `json:"token"`
	}
	ErrorResponse struct {
		Error string `json:"error"`
	}
	GetTicketRequest struct {
		Token string `json:"token"`
	}
	GetTicketsResponse struct {
		Tickets []*Ticket `json:"polls"`
	}
	CreateTicketRequest struct {
		Token    string `json:"token"`
		Subject  string `json:"subject"`
		Category string `json:"category"`
		Body     string `json:"body"`
		Severity string `json:"severity"`
	}
	CreateTicketResponse struct {
		UUID string `json:"message"`
	}
	UpdateTicketRequest struct {
		Token    string `json:"token"`
		UUID     int    `json:"optid"`
		Subject  string `json:"subject"`
		Category string `json:"category"`
		Body     string `json:"body"`
		Severity string `json:"severity"`
	}
	VoteResponse struct {
		Msg string `json:"message"`
	}
	VotedByRequest struct {
		Token string `json:"token"`
		OptId int    `json:"optid"`
	}
	VotedByResponse struct {
		Votes []string `json"votes"`
	}
	GraphQLResponse struct {
		r []*QueryResponse `json"query"`
	}
)

// GraphQLWriter struct to get http response data
type GraphQLWriter struct {
	http.ResponseWriter
	buff *bytes.Buffer
}

func (gwr *GraphQLWriter) Write(p []byte) (int, error) {
	return gwr.buff.Write(p)
}

func encodeGraphQLResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("Querying GraphQL Server")
	rsp := response.(QueryResponse)

	r := rsp.r
	svc := rsp.qsvc

	gWriter := &GraphQLWriter{
		ResponseWriter: w,
		buff:           &bytes.Buffer{},
	}
	svc.ServeHTTP(gWriter, r)

	resp := gWriter.buff.Bytes()
	//log.Println("GQL Response Data is: +%v", resp)

	w.Write([]byte(resp))
	return nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateNewUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeAuthReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req AuthLoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetPollsReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetPollsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeCreatePollReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreatePollRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeVoteReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req VoteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeVotedByReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req VotedByRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGraphQLReq(ctx context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}
