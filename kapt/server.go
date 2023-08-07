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
		decodeGetPollsReq,
		encodeResponse,
	))

	r.Methods("Post").Path("/create-ticket").Handler(httptransport.NewServer(
		endpoints.CreateTicket,
		decodeCreatePollReq,
		encodeResponse,
	))

	r.Methods("Post").Path("/update-ticket").Handler(httptransport.NewServer(
		endpoints.UpdateTicket,
		decodeVoteReq,
		encodeResponse,
	))

	r.Methods("Post").Path("/delete-ticket").Handler(httptransport.NewServer(
		endpoints.DeleteTicket,
		decodeVotedByReq,
		encodeResponse,
	))

	r.Methods("Post").Path("/graph-ql").Handler(httptransport.NewServer(
		endpoints.GraphQL,
		decodeGraphQLReq,
		encodeGraphQLResponse,
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
