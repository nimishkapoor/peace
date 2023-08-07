package kapt

import (
	"context"

	uuid "github.com/google/uuid"
)

type Tenant struct {
	TenantName  string `json:"tenant_name"`
	TenantID    string `json:"tenant_id"`
	Status      int    `json:"status"`
	TenantEmail string `json:"tenant_email"`
}

type User struct {
	UserID    uuid.UUID `json:"user_id"`
	UserName  string    `json:"user_name"`
	Password  string    `json:"password"`
	TenantID  uuid.UUID `json:"tenant_id"`
	Role      string    `json:"role"`
	UserEmail string    `json:"email"`
}

type Ticket struct {
	TicketId    uuid.UUID     `json:"ticket_id"`
	UserID      uuid.UUID     `json:"user_id"`
	TenantID    uuid.UUID     `json:"tenant_id"`
	Subject     string        `json:"subject"`
	Body        string        `json:"body"`
	Category    uuid.UUID     `json:"category"`
	Severity    int           `json:"severity"`
	Assignee    uuid.UUID     `json:"asignee"`
	Label       string        `json:"label"`
	Status      int           `json:"status"`
	Attachments []*Attachment `json:"attachments"`
}

type Category struct {
	CategoryID uuid.UUID `json:"category_id"`
	Name       string    `json:"name"`
	TenantID   uuid.UUID `json:"tenant_id"`
}

type Attachment struct {
	AttachmentID uuid.UUID `json:"attachment_id"`
	Link         string    `json:"link"`
	TicketID     uuid.UUID `json:"ticket_id"`
}

type Repository interface {
	CreateNewUser(ctx context.Context, user User) (string, error)
	CreateNewTenant(ctx context.Context, tenant Tenant, user User) (string, error)
	UpdateTenant(ctx context.Context, tenant Tenant, token string) (string, error)
	UpdateUser(ctx context.Context, user User, token string) (string, error)
	AuthLogin(ctx context.Context, user User) (string, error)
	GetTickets(ctx context.Context, token string) ([]*Ticket, error)
	CreateTicket(ctx context.Context, token string, ticket Ticket) (string, error)
	UpdateTicket(ctx context.Context, token string, ticketid string, ticket Ticket) (string, error)
	CreateCategory(ctx context.Context, token string, category Category) (string, error)
	UpdateCategory(ctx context.Context, token string, category Category) (string, error)
	StoreAttachment(ctx context.Context, token string, attachemt Attachment) (string, error)
}
