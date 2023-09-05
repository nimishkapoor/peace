package polling

import (
	"context"
	"fmt"
	"go-kit/gql"
	"go-kit/gql/ent"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s service) CreateNewUser(ctx context.Context, username string, password string) (string, error) {
	logger := log.With(s.logger, "method", "CreateNewUser")
	user := User{
		Username: username,
		Password: password,
	}

	_, err := s.repository.CreateNewUser(ctx, user)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("User Created!")
	return "Success", nil
}

func (s service) AuthLogin(ctx context.Context, username string, password string) (string, error) {
	logger := log.With(s.logger, "method", "AuthLogin")
	user := User{
		Username: username,
		Password: password,
	}
	token, err := s.repository.AuthLogin(ctx, user)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("Auth Successful!")
	return token, nil
}

func (s service) GetTickets(ctx context.Context, token string) ([]*Ticket, error) {
	logger := log.With(s.logger, "method", "GetTickets")
	fmt.Println(token)
	polls, err := s.repository.GetTickets(ctx, token)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}
	logger.Log("Feeding Tickets!")
	return polls, nil
}

func (s service) CreateTicket(ctx context.Context, token string, Subject string, Category string, Body string, Severity string) (string, error) {
	logger := log.With(s.logger, "method", "CreateTicket")
	_, err := s.repository.CreateTicket(ctx, token, Subject, Subject, Category, Body, Severity)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("Ticket Created!")
	return "Success", nil
}

func (s service) UpdateTicket(ctx context.Context, token string, optid int) (string, error) {
	logger := log.With(s.logger, "method", "UpdateTicket")
	_, err := s.repository.UpdateTicket(ctx, token, ticketID, Subject, Category, Body, Severity)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("Ticket Updated!")
	return "Success", nil
}

func (s service) DeleteTicket(ctx context.Context, token string, optid int) ([]string, error) {
	logger := log.With(s.logger, "method", "DeleteTicket")
	voters, err := s.repository.DeleteTicket(ctx, token, ticketID)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}
	logger.Log("Deleting Ticket!")
	return voters, nil
}

func (s service) GraphQL(ctx context.Context, body *http.Request) (QueryResponse, error) {
	logger := log.With(s.logger, "method", "GraphQL")
	client, err := ent.Open("postgres", dbsource)
	if err != nil {
		fmt.Println(err)
	}
	svr := handler.NewDefaultServer(gql.NewSchema(client))
	logger.Log("Executing GraphQL Query!")
	return QueryResponse{
		r:    body,
		qsvc: svr,
	}, nil
}
