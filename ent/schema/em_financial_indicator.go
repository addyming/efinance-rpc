package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type EmFinancialIndicator struct {
	ent.Schema
}

func (EmFinancialIndicator) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").StorageKey("financial_indicators_id"),
		field.String("financial_indicators_key").Optional(),
		field.String("financial_indicators_value").Optional(),
		field.String("financial_indicators_name").Optional(),
		field.Time("create_time").Optional(),
		field.String("remarks").Optional(),
		field.String("status").Optional()}
}
func (EmFinancialIndicator) Edges() []ent.Edge {
	return nil
}

func (EmFinancialIndicator) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "em_financial_indicator"},
	}
}
