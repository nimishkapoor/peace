// Code generated by ent, DO NOT EDIT.

package ticketmodel

import (
	"kapt/kapt/gql/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldUserID, v))
}

// TenantID applies equality check predicate on the "tenant_id" field. It's identical to TenantIDEQ.
func TenantID(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldTenantID, v))
}

// Subject applies equality check predicate on the "subject" field. It's identical to SubjectEQ.
func Subject(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldSubject, v))
}

// Body applies equality check predicate on the "body" field. It's identical to BodyEQ.
func Body(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldBody, v))
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldCategory, v))
}

// AssigneeID applies equality check predicate on the "assignee_id" field. It's identical to AssigneeIDEQ.
func AssigneeID(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldAssigneeID, v))
}

// Severity applies equality check predicate on the "severity" field. It's identical to SeverityEQ.
func Severity(v int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldSeverity, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldStatus, v))
}

// Time applies equality check predicate on the "time" field. It's identical to TimeEQ.
func Time(v time.Time) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldTime, v))
}

// ClientPriority applies equality check predicate on the "client_priority" field. It's identical to ClientPriorityEQ.
func ClientPriority(v int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldClientPriority, v))
}

// Source applies equality check predicate on the "source" field. It's identical to SourceEQ.
func Source(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldSource, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLTE(FieldUserID, v))
}

// TenantIDEQ applies the EQ predicate on the "tenant_id" field.
func TenantIDEQ(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldTenantID, v))
}

// TenantIDNEQ applies the NEQ predicate on the "tenant_id" field.
func TenantIDNEQ(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNEQ(FieldTenantID, v))
}

// TenantIDIn applies the In predicate on the "tenant_id" field.
func TenantIDIn(vs ...uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldIn(FieldTenantID, vs...))
}

// TenantIDNotIn applies the NotIn predicate on the "tenant_id" field.
func TenantIDNotIn(vs ...uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNotIn(FieldTenantID, vs...))
}

// TenantIDGT applies the GT predicate on the "tenant_id" field.
func TenantIDGT(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGT(FieldTenantID, v))
}

// TenantIDGTE applies the GTE predicate on the "tenant_id" field.
func TenantIDGTE(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGTE(FieldTenantID, v))
}

// TenantIDLT applies the LT predicate on the "tenant_id" field.
func TenantIDLT(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLT(FieldTenantID, v))
}

// TenantIDLTE applies the LTE predicate on the "tenant_id" field.
func TenantIDLTE(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLTE(FieldTenantID, v))
}

// SubjectEQ applies the EQ predicate on the "subject" field.
func SubjectEQ(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldSubject, v))
}

// SubjectNEQ applies the NEQ predicate on the "subject" field.
func SubjectNEQ(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNEQ(FieldSubject, v))
}

// SubjectIn applies the In predicate on the "subject" field.
func SubjectIn(vs ...string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldIn(FieldSubject, vs...))
}

// SubjectNotIn applies the NotIn predicate on the "subject" field.
func SubjectNotIn(vs ...string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNotIn(FieldSubject, vs...))
}

// SubjectGT applies the GT predicate on the "subject" field.
func SubjectGT(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGT(FieldSubject, v))
}

// SubjectGTE applies the GTE predicate on the "subject" field.
func SubjectGTE(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGTE(FieldSubject, v))
}

// SubjectLT applies the LT predicate on the "subject" field.
func SubjectLT(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLT(FieldSubject, v))
}

// SubjectLTE applies the LTE predicate on the "subject" field.
func SubjectLTE(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLTE(FieldSubject, v))
}

// SubjectContains applies the Contains predicate on the "subject" field.
func SubjectContains(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldContains(FieldSubject, v))
}

// SubjectHasPrefix applies the HasPrefix predicate on the "subject" field.
func SubjectHasPrefix(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldHasPrefix(FieldSubject, v))
}

// SubjectHasSuffix applies the HasSuffix predicate on the "subject" field.
func SubjectHasSuffix(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldHasSuffix(FieldSubject, v))
}

// SubjectEqualFold applies the EqualFold predicate on the "subject" field.
func SubjectEqualFold(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEqualFold(FieldSubject, v))
}

// SubjectContainsFold applies the ContainsFold predicate on the "subject" field.
func SubjectContainsFold(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldContainsFold(FieldSubject, v))
}

// BodyEQ applies the EQ predicate on the "body" field.
func BodyEQ(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldBody, v))
}

// BodyNEQ applies the NEQ predicate on the "body" field.
func BodyNEQ(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNEQ(FieldBody, v))
}

// BodyIn applies the In predicate on the "body" field.
func BodyIn(vs ...string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldIn(FieldBody, vs...))
}

// BodyNotIn applies the NotIn predicate on the "body" field.
func BodyNotIn(vs ...string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNotIn(FieldBody, vs...))
}

// BodyGT applies the GT predicate on the "body" field.
func BodyGT(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGT(FieldBody, v))
}

// BodyGTE applies the GTE predicate on the "body" field.
func BodyGTE(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGTE(FieldBody, v))
}

// BodyLT applies the LT predicate on the "body" field.
func BodyLT(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLT(FieldBody, v))
}

// BodyLTE applies the LTE predicate on the "body" field.
func BodyLTE(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLTE(FieldBody, v))
}

// BodyContains applies the Contains predicate on the "body" field.
func BodyContains(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldContains(FieldBody, v))
}

// BodyHasPrefix applies the HasPrefix predicate on the "body" field.
func BodyHasPrefix(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldHasPrefix(FieldBody, v))
}

// BodyHasSuffix applies the HasSuffix predicate on the "body" field.
func BodyHasSuffix(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldHasSuffix(FieldBody, v))
}

// BodyEqualFold applies the EqualFold predicate on the "body" field.
func BodyEqualFold(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEqualFold(FieldBody, v))
}

// BodyContainsFold applies the ContainsFold predicate on the "body" field.
func BodyContainsFold(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldContainsFold(FieldBody, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNotIn(FieldCategory, vs...))
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGT(FieldCategory, v))
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGTE(FieldCategory, v))
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLT(FieldCategory, v))
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLTE(FieldCategory, v))
}

// LabelEQ applies the EQ predicate on the "label" field.
func LabelEQ(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldLabel, v))
}

// LabelNEQ applies the NEQ predicate on the "label" field.
func LabelNEQ(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNEQ(FieldLabel, v))
}

// LabelIn applies the In predicate on the "label" field.
func LabelIn(vs ...string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldIn(FieldLabel, vs...))
}

// LabelNotIn applies the NotIn predicate on the "label" field.
func LabelNotIn(vs ...string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNotIn(FieldLabel, vs...))
}

// LabelGT applies the GT predicate on the "label" field.
func LabelGT(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGT(FieldLabel, v))
}

// LabelGTE applies the GTE predicate on the "label" field.
func LabelGTE(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGTE(FieldLabel, v))
}

// LabelLT applies the LT predicate on the "label" field.
func LabelLT(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLT(FieldLabel, v))
}

// LabelLTE applies the LTE predicate on the "label" field.
func LabelLTE(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLTE(FieldLabel, v))
}

// LabelContains applies the Contains predicate on the "label" field.
func LabelContains(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldContains(FieldLabel, v))
}

// LabelHasPrefix applies the HasPrefix predicate on the "label" field.
func LabelHasPrefix(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldHasPrefix(FieldLabel, v))
}

// LabelHasSuffix applies the HasSuffix predicate on the "label" field.
func LabelHasSuffix(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldHasSuffix(FieldLabel, v))
}

// LabelEqualFold applies the EqualFold predicate on the "label" field.
func LabelEqualFold(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEqualFold(FieldLabel, v))
}

// LabelContainsFold applies the ContainsFold predicate on the "label" field.
func LabelContainsFold(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldContainsFold(FieldLabel, v))
}

// AssigneeIDEQ applies the EQ predicate on the "assignee_id" field.
func AssigneeIDEQ(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldAssigneeID, v))
}

// AssigneeIDNEQ applies the NEQ predicate on the "assignee_id" field.
func AssigneeIDNEQ(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNEQ(FieldAssigneeID, v))
}

// AssigneeIDIn applies the In predicate on the "assignee_id" field.
func AssigneeIDIn(vs ...uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldIn(FieldAssigneeID, vs...))
}

// AssigneeIDNotIn applies the NotIn predicate on the "assignee_id" field.
func AssigneeIDNotIn(vs ...uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNotIn(FieldAssigneeID, vs...))
}

// AssigneeIDGT applies the GT predicate on the "assignee_id" field.
func AssigneeIDGT(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGT(FieldAssigneeID, v))
}

// AssigneeIDGTE applies the GTE predicate on the "assignee_id" field.
func AssigneeIDGTE(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGTE(FieldAssigneeID, v))
}

// AssigneeIDLT applies the LT predicate on the "assignee_id" field.
func AssigneeIDLT(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLT(FieldAssigneeID, v))
}

// AssigneeIDLTE applies the LTE predicate on the "assignee_id" field.
func AssigneeIDLTE(v uuid.UUID) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLTE(FieldAssigneeID, v))
}

// SeverityEQ applies the EQ predicate on the "severity" field.
func SeverityEQ(v int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldSeverity, v))
}

// SeverityNEQ applies the NEQ predicate on the "severity" field.
func SeverityNEQ(v int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNEQ(FieldSeverity, v))
}

// SeverityIn applies the In predicate on the "severity" field.
func SeverityIn(vs ...int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldIn(FieldSeverity, vs...))
}

// SeverityNotIn applies the NotIn predicate on the "severity" field.
func SeverityNotIn(vs ...int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNotIn(FieldSeverity, vs...))
}

// SeverityGT applies the GT predicate on the "severity" field.
func SeverityGT(v int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGT(FieldSeverity, v))
}

// SeverityGTE applies the GTE predicate on the "severity" field.
func SeverityGTE(v int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGTE(FieldSeverity, v))
}

// SeverityLT applies the LT predicate on the "severity" field.
func SeverityLT(v int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLT(FieldSeverity, v))
}

// SeverityLTE applies the LTE predicate on the "severity" field.
func SeverityLTE(v int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLTE(FieldSeverity, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLTE(FieldStatus, v))
}

// TimeEQ applies the EQ predicate on the "time" field.
func TimeEQ(v time.Time) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldTime, v))
}

// TimeNEQ applies the NEQ predicate on the "time" field.
func TimeNEQ(v time.Time) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNEQ(FieldTime, v))
}

// TimeIn applies the In predicate on the "time" field.
func TimeIn(vs ...time.Time) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldIn(FieldTime, vs...))
}

// TimeNotIn applies the NotIn predicate on the "time" field.
func TimeNotIn(vs ...time.Time) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNotIn(FieldTime, vs...))
}

// TimeGT applies the GT predicate on the "time" field.
func TimeGT(v time.Time) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGT(FieldTime, v))
}

// TimeGTE applies the GTE predicate on the "time" field.
func TimeGTE(v time.Time) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGTE(FieldTime, v))
}

// TimeLT applies the LT predicate on the "time" field.
func TimeLT(v time.Time) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLT(FieldTime, v))
}

// TimeLTE applies the LTE predicate on the "time" field.
func TimeLTE(v time.Time) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLTE(FieldTime, v))
}

// ClientPriorityEQ applies the EQ predicate on the "client_priority" field.
func ClientPriorityEQ(v int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldClientPriority, v))
}

// ClientPriorityNEQ applies the NEQ predicate on the "client_priority" field.
func ClientPriorityNEQ(v int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNEQ(FieldClientPriority, v))
}

// ClientPriorityIn applies the In predicate on the "client_priority" field.
func ClientPriorityIn(vs ...int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldIn(FieldClientPriority, vs...))
}

// ClientPriorityNotIn applies the NotIn predicate on the "client_priority" field.
func ClientPriorityNotIn(vs ...int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNotIn(FieldClientPriority, vs...))
}

// ClientPriorityGT applies the GT predicate on the "client_priority" field.
func ClientPriorityGT(v int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGT(FieldClientPriority, v))
}

// ClientPriorityGTE applies the GTE predicate on the "client_priority" field.
func ClientPriorityGTE(v int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGTE(FieldClientPriority, v))
}

// ClientPriorityLT applies the LT predicate on the "client_priority" field.
func ClientPriorityLT(v int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLT(FieldClientPriority, v))
}

// ClientPriorityLTE applies the LTE predicate on the "client_priority" field.
func ClientPriorityLTE(v int) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLTE(FieldClientPriority, v))
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEQ(FieldSource, v))
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNEQ(FieldSource, v))
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldIn(FieldSource, vs...))
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldNotIn(FieldSource, vs...))
}

// SourceGT applies the GT predicate on the "source" field.
func SourceGT(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGT(FieldSource, v))
}

// SourceGTE applies the GTE predicate on the "source" field.
func SourceGTE(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldGTE(FieldSource, v))
}

// SourceLT applies the LT predicate on the "source" field.
func SourceLT(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLT(FieldSource, v))
}

// SourceLTE applies the LTE predicate on the "source" field.
func SourceLTE(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldLTE(FieldSource, v))
}

// SourceContains applies the Contains predicate on the "source" field.
func SourceContains(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldContains(FieldSource, v))
}

// SourceHasPrefix applies the HasPrefix predicate on the "source" field.
func SourceHasPrefix(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldHasPrefix(FieldSource, v))
}

// SourceHasSuffix applies the HasSuffix predicate on the "source" field.
func SourceHasSuffix(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldHasSuffix(FieldSource, v))
}

// SourceEqualFold applies the EqualFold predicate on the "source" field.
func SourceEqualFold(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldEqualFold(FieldSource, v))
}

// SourceContainsFold applies the ContainsFold predicate on the "source" field.
func SourceContainsFold(v string) predicate.TicketModel {
	return predicate.TicketModel(sql.FieldContainsFold(FieldSource, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TicketModel) predicate.TicketModel {
	return predicate.TicketModel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TicketModel) predicate.TicketModel {
	return predicate.TicketModel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TicketModel) predicate.TicketModel {
	return predicate.TicketModel(func(s *sql.Selector) {
		p(s.Not())
	})
}
