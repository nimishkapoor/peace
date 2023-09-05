package kapt

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
)

type QueryResponse struct {
	r    *http.Request
	qsvc *handler.Server
}

type Service interface {
	CreateNewUser(ctx context.Context, username string, password string) (string, error)
	AuthLogin(ctx context.Context, username string, password string) (string, error)
	GetTickets(ctx context.Context, token string) ([]*Ticket, error)
	CreateTicket(ctx context.Context, token string) (string, error)
	UpdateTicket(ctx context.Context, token string) (string, error)
	DeleteTicket(ctx context.Context, token string) (string, error)
	CreateCategory(ctx context.Context, token string) (string, error)
	UpdateCategory(ctx context.Context, token string) (string, error)
	DeleteCategory(ctx context.Context, token string) (string, error)
	GetCategories(ctx context.Context, token string) ([]*Category, error)
	CreateComment(ctx context.Context, token string) (string, error)
	UpdateComment(ctx context.Context, token string) (string, error)
	DeleteComment(ctx context.Context, token string) (string, error)
	GetComments(ctx context.Context, token string) ([]*Comment, error)
	GraphQL(ctx context.Context, body *http.Request) (QueryResponse, error)
}
