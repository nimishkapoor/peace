package polling

import (
	"context"
	"errors"
	"fmt"
	"go-kit/gql/ent"
	"go-kit/gql/ent/usermodel"

	"github.com/go-kit/kit/log"
	_ "github.com/lib/pq"
)

var RepoErr = errors.New("unable to handle repo request")

type repo struct {
	client *ent.Client
	logger log.Logger
}

func NewRepo(client *ent.Client, logger log.Logger) Repository {
	return &repo{
		client: client,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateNewUser(ctx context.Context, user User) (string, error) {
	if user.Username == "" || user.Password == "" {
		return "", errors.New("username or password cannot be empty")
	}
	pswd, _ := GeneratehashPassword(user.Password)
	_, err := repo.client.UserModel.
		Create().SetUserName(user.Username).
		SetPswd(pswd).
		Save(ctx)

	if err != nil {
		return "", errors.New("user already exists")
	}
	return "", nil
}

func (repo *repo) AuthLogin(ctx context.Context, user User) (string, error) {
	ret, err := repo.client.UserModel.Query().
		Where(usermodel.UserNameEQ(user.Username)).
		First(ctx)
	if err != nil {
		return "", errors.New("user does not exists")
	} else if !CheckPasswordHash(user.Password, ret.Pswd) {
		return "", errors.New("incorrect password")
	}
	validToken, _ := GenerateJWT(user.Username)
	return validToken, nil
}

func (repo *repo) GetTickets(ctx context.Context, token string) ([]*Ticket, error) {

	claims, err := ValidateJWT(token)
	if err != nil {
		return nil, err
	}
	tickets := make([]*Ticket, 0)

	rows, err := repo.client.TicketModel.Query().Order(ent.Asc(TicketModel.FieldID)).
		All(ctx)

	for _, row := range tickets {
		ticket := new(Ticket)
		ticket.Id = row.ID
		ticket.Subject = row.Subject
		ticket.UserID = row.UserID
		ticket.Body = row.Body
		ticket.Category = row.Category
		ticket.Severity = row.Severity
		ticket.Assignee = row.AssigneeID
		ticket.Label = row.label
		ticket.Status = row.Status
		ticket.ClientPriority = row.ClientPriority
		ticket.Source = row.Source
	}

	if err != nil {
		fmt.Println("Failed to execute query: ", err)
		return nil, err
	} else {
		fmt.Printf("Feeding Tickets!\n")
		return tickets, nil
	}
}

func (repo *repo) CreateTicket(ctx context.Context, token string) (string, error) {
	claims, err := ValidateJWT(token)
	if err != nil {
		return "", err
	}

	kapTctx := ctx.Value(contextKey)
	req := kapTctx.request.(CreateTicketRequest)

	ret, err := repo.client.TicketModel.Create().SetSubject(Subject).
		SetUserName(claims["username"].(string)).
		SetBody(req.Body).
		SetCategory(req.Category).
		SetSeverity(req.Severity).
		Save(ctx)

	if err != nil {
		fmt.Println("Failed to execute query: ", err)
		return "", errors.New("internal server error")
	}

	fmt.Printf("Ticket Created!\n")
	return "Worked!", nil
}

func (repo *repo) UpdateTicket(ctx context.Context, token string) (string, error) {
	claims, err := ValidateJWT(token)
	if err != nil {
		return "", err
	}

	return "Ticket Updated!", nil
}

func (repo *repo) DeleteTicket(ctx context.Context, token string) (string, error) {
	_, err := ValidateJWT(token)
	if err != nil {
		return nil, err
	}
	return "Ticket Deleted!", nil
}

func (repo *repo) CreateCategory(ctx context.Context, token string) (string, error) {
	claims, err := ValidateJWT(token)
	if err != nil {
		return "", err
	}

	fmt.Printf("Category Created!\n")
	return "Worked!", nil
}

func (repo *repo) UpdateCategory(ctx context.Context, token string) (string, error) {
	claims, err := ValidateJWT(token)
	if err != nil {
		return "", err
	}

	return "Category Updated!", nil
}

func (repo *repo) DeleteCategory(ctx context.Context, token string) (string, error) {
	_, err := ValidateJWT(token)
	if err != nil {
		return nil, err
	}
	return "Category Deleted!", nil
}

func (repo *repo) GetCategories(ctx context.Context, token string) ([]*Category, error) {

	claims, err := ValidateJWT(token)
	if err != nil {
		return nil, err
	}
	categories := make([]*Category, 0)

	if err != nil {
		fmt.Println("Failed to execute query: ", err)
		return nil, err
	} else {
		fmt.Printf("Feeding Categories!\n")
		return categories, nil
	}
}

func (repo *repo) CreateComment(ctx context.Context, token string) (string, error) {
	claims, err := ValidateJWT(token)
	if err != nil {
		return "", err
	}

	fmt.Printf("Comment Created!\n")
	return "Worked!", nil
}

func (repo *repo) UpdateComment(ctx context.Context, token string) (string, error) {
	claims, err := ValidateJWT(token)
	if err != nil {
		return "", err
	}

	return "Comment Updated!", nil
}

func (repo *repo) DeleteComment(ctx context.Context, token string) (string, error) {
	_, err := ValidateJWT(token)
	if err != nil {
		return nil, err
	}
	return "Comment Deleted!", nil
}

func (repo *repo) GetComments(ctx context.Context, token string) ([]*Comment, error) {

	claims, err := ValidateJWT(token)
	if err != nil {
		return nil, err
	}
	categories := make([]*Category, 0)

	if err != nil {
		fmt.Println("Failed to execute query: ", err)
		return nil, err
	} else {
		fmt.Printf("Feeding Comments!\n")
		return categories, nil
	}
}
