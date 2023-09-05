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
		Tickets []*Ticket `json:"tickets"`
	}
	CreateTicketRequest struct {
		Token    string `json:"token"`
		Subject  string `json:"subject"`
		Category string `json:"category"`
		Body     string `json:"body"`
		Severity string `json:"severity"`
	}
	CreateTicketResponse struct {
		UUID string `json:"uuid"`
	}
	UpdateTicketRequest struct {
		Token    string `json:"token"`
		UUID     int    `json:"ticket_id"`
		Subject  string `json:"subject"`
		Category string `json:"category"`
		Body     string `json:"body"`
		Severity string `json:"severity"`
	}
	UpdateTicketResponse struct {
		Msg string `json:"message"`
	}
	DeleteTicketRequest struct {
		Token string `json:"token"`
		UUID  int    `json:"ticket_id"`
	}
	DeleteTicketResponse struct {
		Msg string `json:"message"`
	}
	DeleteCategoryRequest struct {
		Token string `json:"token"`
		UUID  int    `json:"category_id"`
	}
	DeleteCategoryResponse struct {
		Msg string `json:"message"`
	}
	CreateCategoryRequest struct {
		Token string `json:"token"`
		Name  string `json:"name"`
	}
	CreateCategoryResponse struct {
		UUID string `json:"uuid"`
	}
	GetCategoriesRequest struct {
		Token string `json:"token"`
	}
	GetCategoriesResponse struct {
		Categories []*Category `json:"categories"`
	}
	CreateCommentRequest struct {
		Token string `json:"token"`
		Body  string `json:"body"`
	}
	CreateCommentResponse struct {
		UUID string `json:"uuid"`
	}
	DeleteCommentRequest struct {
		Token string `json:"token"`
		UUID  string `json:"uuid"`
	}
	DeleteCommentResponse struct {
		Msg string `json:"message"`
	}
	UpdateCommentRequest struct {
		Token string `json:"token"`
		UUID  string `json:"uuid"`
		Body  string `json:"body"`
	}
	UpdateCommentResponse struct {
		Msg string `json:"message"`
	}
	GetCommentsRequest struct {
		Token    string `json:"token"`
		TicketID string `json:"ticket_id"`
	}
	GetCommentsResponse struct {
		Comments []*Comment `json:"comments"`
	}
	GraphQLResponse struct {
		r []*QueryResponse `json:"query"`
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

func decodeGetTicketReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetTicketRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeCreateTicketReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateTicketRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeUpdateTicketReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UpdateTicketRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteTicketReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req DeleteTicketRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteCategoryReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req DeleteCategoryRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeCreateCategoryReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateCategoryRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetCategoriesReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetCategoriesRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeCreateCommentReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateCommentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeUpdateCommentReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UpdateCommentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteCommentReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req DeleteCommentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetCommentsReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetCommentsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGraphQLReq(ctx context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}
