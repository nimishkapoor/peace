package kapt

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

const dbsource = "postgres://postgres:123@localhost/postgres?sslmode=disable"

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("Post").Path("/new-user").Handler(httptransport.NewServer(
		endpoints.CreateNewUser,
		decodeUserReq,
		encodeResponse,
	))

	r.Methods("Post").Path("/login").Handler(httptransport.NewServer(
		endpoints.AuthLogin,
		decodeAuthReq,
		encodeResponse,
	))

	r.Methods("Post").Path("/tickets").Handler(httptransport.NewServer(
		endpoints.GetTickets,
		decodeGetTicketReq,
		encodeResponse,
	))

	r.Methods("Post").Path("/create-ticket").Handler(httptransport.NewServer(
		endpoints.CreateTicket,
		decodeCreateTicketReq,
		encodeResponse,
	))

	r.Methods("Post").Path("/update-ticket").Handler(httptransport.NewServer(
		endpoints.UpdateTicket,
		decodeUpdateTicketReq,
		encodeResponse,
	))

	r.Methods("Post").Path("/delete-ticket").Handler(httptransport.NewServer(
		endpoints.DeleteTicket,
		decodeDeleteTicketReq,
		encodeResponse,
	))

	r.Methods("Post").Path("/graph-ql").Handler(httptransport.NewServer(
		endpoints.GraphQL,
		decodeGraphQLReq,
		encodeGraphQLResponse,
	))

	r.Methods("Post").Path("/create-categoty").Handler(httptransport.NewServer(
		endpoints.CreateCategory,
		decodeCreateCategoryReq,
		encodeResponse,
	))

	r.Methods("Post").Path("/delete-categoty").Handler(httptransport.NewServer(
		endpoints.DeleteCategory,
		decodeDeleteCategoryReq,
		encodeResponse,
	))

	r.Methods("Post").Path("/create-comment").Handler(httptransport.NewServer(
		endpoints.CreateComment,
		decodeCreateCommentReq,
		encodeResponse,
	))

	r.Methods("Post").Path("/update-comment").Handler(httptransport.NewServer(
		endpoints.UpdateComment,
		decodeUpdateCommentReq,
		encodeResponse,
	))

	r.Methods("Post").Path("/delete-comment").Handler(httptransport.NewServer(
		endpoints.DeleteComment,
		decodeDeleteCommentReq,
		encodeResponse,
	))

	r.Methods("Post").Path("/get-comments").Handler(httptransport.NewServer(
		endpoints.GetComments,
		decodeGetCommentsReq,
		encodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
