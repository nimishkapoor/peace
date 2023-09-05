package kapt

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type KapTContext struct {
	Ctx     context.Context
	Request interface{}
}

const contextKey string = "kapt_context"

func NewKapTContext(ctx context.Context, req interface{}) context.Context {
	kapTCtx := &KapTContext{
		Ctx:     ctx,
		Request: req,
	}
	return context.WithValue(ctx, contextKey, kapTCtx)
}

type Endpoints struct {
	CreateNewUser  endpoint.Endpoint
	AuthLogin      endpoint.Endpoint
	GetTickets     endpoint.Endpoint
	CreateTicket   endpoint.Endpoint
	UpdateTicket   endpoint.Endpoint
	DeleteTicket   endpoint.Endpoint
	GraphQL        endpoint.Endpoint
	CreateCategory endpoint.Endpoint
	DeleteCategory endpoint.Endpoint
	CreateComment  endpoint.Endpoint
	UpdateComment  endpoint.Endpoint
	DeleteComment  endpoint.Endpoint
	GetComments    endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateNewUser:  makeCreateNewUserEndpoint(s),
		AuthLogin:      makeAuthLoginEndpoint(s),
		GetTickets:     makeGetTicketsEndpoint(s),
		CreateTicket:   makeCreateTicketEndpoint(s),
		UpdateTicket:   makeUpdateTicketEndpoint(s),
		DeleteTicket:   makeDeleteTicketEndpoint(s),
		GraphQL:        makeGraphQLEndpoint(s),
		CreateCategory: makeCreateCategoryEndpoint(s),
		DeleteCategory: makeDeleteCategoryEndpoint(s),
		CreateComment:  makeCreateCommentEndpoint(s),
		UpdateComment:  makeUpdateCommentEndpoint(s),
		DeleteComment:  makeDeleteCommentEndpoint(s),
		GetComments:    makeGetCommentsEndpoint(s),
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
		req := request.(GetTicketRequest)
		tickets, err := s.GetTickets(ctx, req.Token)
		if err != nil {
			return ErrorResponse{Error: err.Error()}, nil
		}
		return GetTicketsResponse{Tickets: tickets}, err
	}
}

func makeCreateTicketEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTicketRequest)
		ctx = NewKapTContext(ctx, request)
		ticket, err := s.CreateTicket(ctx, req.Token)
		if err != nil {
			return ErrorResponse{Error: err.Error()}, nil
		}
		return CreateTicketResponse{UUID: ticket}, err
	}
}

func makeUpdateTicketEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateTicketRequest)
		ok, err := s.UpdateTicket(ctx, req.Token)
		if err != nil {
			return ErrorResponse{Error: err.Error()}, nil
		}
		return UpdateTicketResponse{Msg: ok}, err
	}
}

func makeDeleteTicketEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteTicketRequest)
		ok, err := s.DeleteTicket(ctx, req.Token)
		if err != nil {
			return ErrorResponse{Error: err.Error()}, nil
		}
		return DeleteTicketResponse{Msg: ok}, err
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

func makeCreateCategoryEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCategoryRequest)
		category, err := s.CreateCategory(ctx, req.Token)
		if err != nil {
			return ErrorResponse{Error: err.Error()}, nil
		}
		return CreateCategoryResponse{UUID: category}, err
	}
}

func makeDeleteCategoryEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteCategoryRequest)
		ok, err := s.DeleteCategory(ctx, req.Token)
		if err != nil {
			return ErrorResponse{Error: err.Error()}, nil
		}
		return DeleteCategoryResponse{Msg: ok}, err
	}
}

func makeGetCategoriesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCategoriesRequest)

		ok, err := s.GetCategories(ctx, req.Token)
		if err != nil {
			return ErrorResponse{Error: err.Error()}, nil
		}
		return GetCategoriesResponse{Categories: ok}, err
	}
}

func makeCreateCommentEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCommentRequest)
		ok, err := s.CreateComment(ctx, req.Token)
		if err != nil {
			return ErrorResponse{Error: err.Error()}, nil
		}
		return CreateCommentResponse{UUID: ok}, err
	}
}

func makeUpdateCommentEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateCommentRequest)
		ok, err := s.UpdateComment(ctx, req.Token)
		if err != nil {
			return ErrorResponse{Error: err.Error()}, nil
		}
		return UpdateCommentResponse{Msg: ok}, err
	}
}

func makeDeleteCommentEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteCommentRequest)
		ok, err := s.DeleteComment(ctx, req.Token)
		if err != nil {
			return ErrorResponse{Error: err.Error()}, nil
		}
		return DeleteCommentResponse{Msg: ok}, err
	}
}

func makeGetCommentsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCommentsRequest)
		ok, err := s.GetComments(ctx, req.Token)
		if err != nil {
			return ErrorResponse{Error: err.Error()}, nil
		}
		return GetCommentsResponse{Comments: ok}, err
	}
}
