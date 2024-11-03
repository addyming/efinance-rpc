package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type EmStockWatch struct {
	ent.Schema
}

func (EmStockWatch) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").StorageKey("stock_watch_id"),
		field.String("stock_watch_code").Optional(),
		field.String("stock_watch_name").Optional(),
		field.String("stock_watch_up_price").Optional(),
		field.String("stock_watch_fall_price").Optional(),
		field.String("stock_watch_up_change").Optional(),
		field.String("stock_watch_fall_change").Optional()}
}
func (EmStockWatch) Edges() []ent.Edge {
	return nil
}

func (EmStockWatch) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "em_stock_watch"},
	}
}
