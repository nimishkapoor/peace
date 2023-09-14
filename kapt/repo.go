package polling

import (
	"context"
	"errors"
	"fmt"
	"go-kit/gql/ent"
	"go-kit/gql/ent/usermodel"
	"time"

	"github.com/go-kit/kit/log"
	uuid "github.com/google/uuid"
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

	for _, row := range rows {
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
		attachments := make([]*Attachment, 0)
		subRows, err := repo.client.AttachemtModel.Query().Order(ent.Asc(TicketModel.FieldID)).
			All(ctx)
		for _, subRow := range subrows {
			attachment := new(Attachment)
			attachment.AttachmentID = row.AttachmentID
			attachment.Link = row.Link
			attachment.TicketID = row.TicketID
			attachements = append(attachments, attachment)
		}
		ticket.Attachment = attachments
		tickets = append(tickets, ticket)
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

	ticketID := uuid.New()
	time := time.Now()

	ret, err := repo.client.TicketModel.Create().SetSubject(req.Subject).
		SetUserID(claims["username"].(string)).
		SetBody(req.Body).
		SetTenantID().
		SetCategory(req.Category).
		SetSeverity(req.Severity).
		SetLabel(req.Label).
		SetAssigneeID(req.AssigneeID).
		SetStatus(req.Status).
		SetTime().
		SetClientPriority().
		SetSource().
		Save(ctx)

	if err != nil {
		fmt.Println("Failed to execute query: ", err)
		return "", errors.New("internal server error")
	}

	if attachmets != nil {
		for _, row := range attachments {
			attachmentID := uuid.New()
			ret, err := repo.client.AttachmentModel.Create().SetID(attachmentID).
				SetLink(row.Link).
				SetTicketID(ticketID).
				Save(ctx)

			if err != nil {
				fmt.Println("Failed to execute query: ", err)
				return "", errors.New("internal server error")
			}
		}
	}

	fmt.Printf("Ticket Created!\n")
	return "Worked!", nil
}

func (repo *repo) UpdateTicket(ctx context.Context, token string) (string, error) {
	claims, err := ValidateJWT(token)
	if err != nil {
		return "", err
	}

	ticketID = ""

	ret, err := repo.client.TicketModel.UpdateOneID(ticketID).SetSubject(req.Subject).
		SetBody(req.Body).
		SetCategory(req.Category).
		SetSeverity(req.Severity).
		Save(ctx)

	if err != nil {
		fmt.Println("Failed to execute query: ", err)
		return "", errors.New("internal server error")
	}

	return "Ticket Updated!", nil
}

func (repo *repo) DeleteTicket(ctx context.Context, token string) (string, error) {
	_, err := ValidateJWT(token)
	if err != nil {
		return nil, err
	}

	ticketID = ""

	err := repo.client.TicketModel.DeleteOneID(ticketID).Exec(ctx)

	switch {
	case ent.IsNotFound(err):
		fmt.Println("ticket not found")
		return "", errors.New("ticket not found")
	case err != nil:
		fmt.Println("Failed to execute query: ", err)
		return "", errors.New("internal server error")
	}

	return "Ticket Deleted!", nil
}

func (repo *repo) CreateCategory(ctx context.Context, token string) (string, error) {
	claims, err := ValidateJWT(token)
	if err != nil {
		return "", err
	}

	tenantID := ""
	categoryID := uuid.New()

	ret, err := repo.client.CategoryModel.Create().SetID(categoryID).
		SetName(req.Name).
		SetTenantID(tenantID).
		Save(ctx)

	if err != nil {
		fmt.Println("Failed to execute query: ", err)
		return "", errors.New("internal server error")
	}

	fmt.Printf("Category Created!\n")
	return "Worked!", nil
}

func (repo *repo) UpdateCategory(ctx context.Context, token string) (string, error) {
	claims, err := ValidateJWT(token)
	if err != nil {
		return "", err
	}

	categoryID := ""

	ret, err := repo.client.CategoryModel.UpdateOneID(categoryID).SetID(categoryID).
		SetName(req.Name).
		SetTenantID(tenantID).
		Save(ctx)

	if err != nil {
		fmt.Println("Failed to execute query: ", err)
		return "", errors.New("internal server error")
	}

	return "Category Updated!", nil
}

func (repo *repo) DeleteCategory(ctx context.Context, token string) (string, error) {
	_, err := ValidateJWT(token)
	if err != nil {
		return nil, err
	}

	categoryID := ""

	err := repo.client.CateorgyModel.DeleteOneID(categoryID).Exec(ctx)

	switch {
	case ent.IsNotFound(err):
		fmt.Println("ticket not found")
		return "", errors.New("ticket not found")
	case err != nil:
		fmt.Println("Failed to execute query: ", err)
		return "", errors.New("internal server error")
	}

	return "Category Deleted!", nil
}

func (repo *repo) GetCategories(ctx context.Context, token string) ([]*Category, error) {

	claims, err := ValidateJWT(token)
	if err != nil {
		return nil, err
	}
	categories := make([]*Category, 0)

	rows, err := repo.client.CategoryModel.Query().Order(ent.Asc(CategoryModel.FieldID)).
		All(ctx)

	for _, row := range tickets {
		category := new(Category)
		category.Id = row.ID
		category.Name = row.Name
		category.TenantID = row.TenantID
		categories = append(categories, category)
	}

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

	ticketID := ""
	commentID := uuid.New()
	body := req.Body
	source := ""
	time := time.Now()

	ret, err := repo.client.ThreadModel.Create().SetID(commentID).
		SetBody(body).
		SetTicketUUID(ticketId).
		SetTime().
		SetSource(source).
		Save(ctx)

	if err != nil {
		fmt.Println("Failed to execute query: ", err)
		return "", errors.New("internal server error")
	}

	fmt.Printf("Comment Created!\n")
	return "Worked!", nil
}

func (repo *repo) UpdateComment(ctx context.Context, token string) (string, error) {
	claims, err := ValidateJWT(token)
	if err != nil {
		return "", err
	}

	ret, err := repo.client.ThreadModel.UpdateOneID(commentID).
		SetBody(body).
		SetTicketUUID(ticketId).
		SetTime().
		SetSource(source).
		Save(ctx)

	if err != nil {
		fmt.Println("Failed to execute query: ", err)
		return "", errors.New("internal server error")
	}

	return "Comment Updated!", nil
}

func (repo *repo) DeleteComment(ctx context.Context, token string) (string, error) {
	_, err := ValidateJWT(token)
	if err != nil {
		return nil, err
	}

	commentID := ""

	err := repo.client.ThreadModel.DeleteOneID(commentID).Exec(ctx)

	switch {
	case ent.IsNotFound(err):
		fmt.Println("comment not found")
		return "", errors.New("comment not found")
	case err != nil:
		fmt.Println("Failed to execute query: ", err)
		return "", errors.New("internal server error")
	}

	return "Comment Deleted!", nil
}

func (repo *repo) GetComments(ctx context.Context, token string) ([]*Comment, error) {

	claims, err := ValidateJWT(token)
	if err != nil {
		return nil, err
	}
	comments := make([]*Comment, 0)

	rows, err := repo.client.ThreadModel.Query().Order(ent.Asc(ThreadModel.FieldID)).
		All(ctx)

	for _, row := range tickets {
		comment := new(Comment)
		comment.Id = row.ID
		comment.Body = row.Body
		comment.TicketUUID = row.TicketUUID
		comment.Time = row.Time
		comment.Source = row.Source
		attachments := make([]*Attachment, 0)
		subRows, err := repo.client.AttachemtModel.Query().Order(ent.Asc(TicketModel.FieldID)).
			All(ctx)
		for _, subRow := range subrows {
			attachment := new(Attachment)
			attachment.AttachmentID = row.AttachmentID
			attachment.Link = row.Link
			attachemnt.TicketID = row.TicketID
			attachments = append(attachments, attachment)
		}
		comments.Attachments = attachments
		comments = append(commets, comment)
	}

	if err != nil {
		fmt.Println("Failed to execute query: ", err)
		return nil, err
	} else {
		fmt.Printf("Feeding Comments!\n")
		return comments, nil
	}
}
