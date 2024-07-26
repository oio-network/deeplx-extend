// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/oio-network/deeplx-extend/app/deeplx/internal/data/ent/accesslog"
	"github.com/oio-network/deeplx-extend/app/deeplx/internal/data/ent/predicate"
	"github.com/oio-network/deeplx-extend/app/deeplx/internal/data/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 2)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   accesslog.Table,
			Columns: accesslog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: accesslog.FieldID,
			},
		},
		Type: "AccessLog",
		Fields: map[string]*sqlgraph.FieldSpec{
			accesslog.FieldCreatedAt:   {Type: field.TypeTime, Column: accesslog.FieldCreatedAt},
			accesslog.FieldUserID:      {Type: field.TypeInt64, Column: accesslog.FieldUserID},
			accesslog.FieldIP:          {Type: field.TypeString, Column: accesslog.FieldIP},
			accesslog.FieldCountryName: {Type: field.TypeString, Column: accesslog.FieldCountryName},
			accesslog.FieldCountryCode: {Type: field.TypeString, Column: accesslog.FieldCountryCode},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: user.FieldID,
			},
		},
		Type: "User",
		Fields: map[string]*sqlgraph.FieldSpec{
			user.FieldCreatedAt: {Type: field.TypeTime, Column: user.FieldCreatedAt},
			user.FieldUpdatedAt: {Type: field.TypeTime, Column: user.FieldUpdatedAt},
			user.FieldToken:     {Type: field.TypeUUID, Column: user.FieldToken},
		},
	}
	graph.MustAddE(
		"owner_user",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesslog.OwnerUserTable,
			Columns: []string{accesslog.OwnerUserColumn},
			Bidi:    false,
		},
		"AccessLog",
		"User",
	)
	graph.MustAddE(
		"access_logs",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AccessLogsTable,
			Columns: []string{user.AccessLogsColumn},
			Bidi:    false,
		},
		"User",
		"AccessLog",
	)
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (alq *AccessLogQuery) addPredicate(pred func(s *sql.Selector)) {
	alq.predicates = append(alq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the AccessLogQuery builder.
func (alq *AccessLogQuery) Filter() *AccessLogFilter {
	return &AccessLogFilter{config: alq.config, predicateAdder: alq}
}

// addPredicate implements the predicateAdder interface.
func (m *AccessLogMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the AccessLogMutation builder.
func (m *AccessLogMutation) Filter() *AccessLogFilter {
	return &AccessLogFilter{config: m.config, predicateAdder: m}
}

// AccessLogFilter provides a generic filtering capability at runtime for AccessLogQuery.
type AccessLogFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *AccessLogFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int64 predicate on the id field.
func (f *AccessLogFilter) WhereID(p entql.Int64P) {
	f.Where(p.Field(accesslog.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *AccessLogFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(accesslog.FieldCreatedAt))
}

// WhereUserID applies the entql int64 predicate on the user_id field.
func (f *AccessLogFilter) WhereUserID(p entql.Int64P) {
	f.Where(p.Field(accesslog.FieldUserID))
}

// WhereIP applies the entql string predicate on the ip field.
func (f *AccessLogFilter) WhereIP(p entql.StringP) {
	f.Where(p.Field(accesslog.FieldIP))
}

// WhereCountryName applies the entql string predicate on the country_name field.
func (f *AccessLogFilter) WhereCountryName(p entql.StringP) {
	f.Where(p.Field(accesslog.FieldCountryName))
}

// WhereCountryCode applies the entql string predicate on the country_code field.
func (f *AccessLogFilter) WhereCountryCode(p entql.StringP) {
	f.Where(p.Field(accesslog.FieldCountryCode))
}

// WhereHasOwnerUser applies a predicate to check if query has an edge owner_user.
func (f *AccessLogFilter) WhereHasOwnerUser() {
	f.Where(entql.HasEdge("owner_user"))
}

// WhereHasOwnerUserWith applies a predicate to check if query has an edge owner_user with a given conditions (other predicates).
func (f *AccessLogFilter) WhereHasOwnerUserWith(preds ...predicate.User) {
	f.Where(entql.HasEdgeWith("owner_user", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (uq *UserQuery) addPredicate(pred func(s *sql.Selector)) {
	uq.predicates = append(uq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the UserQuery builder.
func (uq *UserQuery) Filter() *UserFilter {
	return &UserFilter{config: uq.config, predicateAdder: uq}
}

// addPredicate implements the predicateAdder interface.
func (m *UserMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the UserMutation builder.
func (m *UserMutation) Filter() *UserFilter {
	return &UserFilter{config: m.config, predicateAdder: m}
}

// UserFilter provides a generic filtering capability at runtime for UserQuery.
type UserFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *UserFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int64 predicate on the id field.
func (f *UserFilter) WhereID(p entql.Int64P) {
	f.Where(p.Field(user.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *UserFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(user.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *UserFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(user.FieldUpdatedAt))
}

// WhereToken applies the entql [16]byte predicate on the token field.
func (f *UserFilter) WhereToken(p entql.ValueP) {
	f.Where(p.Field(user.FieldToken))
}

// WhereHasAccessLogs applies a predicate to check if query has an edge access_logs.
func (f *UserFilter) WhereHasAccessLogs() {
	f.Where(entql.HasEdge("access_logs"))
}

// WhereHasAccessLogsWith applies a predicate to check if query has an edge access_logs with a given conditions (other predicates).
func (f *UserFilter) WhereHasAccessLogsWith(preds ...predicate.AccessLog) {
	f.Where(entql.HasEdgeWith("access_logs", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}
