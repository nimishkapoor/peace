package kapt

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateNewUser endpoint.Endpoint
	AuthLogin     endpoint.Endpoint
	GetTickets    endpoint.Endpoint
	CreateTicket  endpoint.Endpoint
	UpdateTicket  endpoint.Endpoint
	DeleteTicket  endpoint.Endpoint
	GraphQL       endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateNewUser: makeCreateNewUserEndpoint(s),
		AuthLogin:     makeAuthLoginEndpoint(s),
		GetTickets:    makeGetTicketsEndpoint(s),
		CreateTicket:  makeCreateTicketEndpoint(s),
		UpdateTicket:  makeUpdateTicketEndpoint(s),
		DeleteTicket:  makeDeleteTicketEndpoint(s),
		GraphQL:       makeGraphQLEndpoint(s),
	}
}

func makeCreateNewUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateNewUserRequest)
		ok, err := s.CreateNewUser(ctx, req.Username, req.Password)
		if err != nil {
			return ErrorResponse{Error: err.Error()}, nil
		}
		return CreateNewUserResponse{Msg: ok}, err
	}
}

func makeAuthLoginEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AuthLoginRequest)
		token, err := s.AuthLogin(ctx, req.Username, req.Password)
		if err != nil {
			return ErrorResponse{Error: err.Error()}, nil
		}
		return AuthLoginResponse{TokenString: token}, err
	}
}

func makeGetTicketsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPollsRequest)
		polls, err := s.GetPolls(ctx, req.Token)
		if err != nil {
			return ErrorResponse{Error: err.Error()}, nil
		}
		return GetPollsResponse{Polls: polls}, err
	}
}

func makeCreateTicketEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreatePollRequest)
		ok, err := s.CreatePoll(ctx, req.Token, req.Subject, req.Opt1, req.Opt2, req.Opt3, req.Opt4)
		if err != nil {
			return ErrorResponse{Error: err.Error()}, nil
		}
		return CreatePollResponse{Msg: ok}, err
	}
}

func makeUpdateTicketEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(VoteRequest)
		ok, err := s.Vote(ctx, req.Token, req.OptId)
		if err != nil {
			return ErrorResponse{Error: err.Error()}, nil
		}
		return VoteResponse{Msg: ok}, err
	}
}

func makeDeleteTicketEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(VotedByRequest)
		votes, err := s.VotedBy(ctx, req.Token, req.OptId)
		if err != nil {
			return ErrorResponse{Error: err.Error()}, nil
		}
		return VotedByResponse{Votes: votes}, err
	}
}

func makeGraphQLEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, _ := request.(*http.Request)
		votes, err := s.GraphQL(ctx, req)
		if err != nil {
			return ErrorResponse{Error: err.Error()}, nil
		}
		return votes, err
	}
}
