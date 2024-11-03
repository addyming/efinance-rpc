package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type EmStockBasicInfo struct {
	ent.Schema
}

func (EmStockBasicInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").StorageKey("stock_basic_info_id"),
		field.String("f1").Optional(),
		field.String("f2").Optional(),
		field.String("f3").Optional(),
		field.String("f4").Optional(),
		field.String("f5").Optional(),
		field.String("f6").Optional(),
		field.String("f7").Optional(),
		field.String("f8").Optional(),
		field.String("f9").Optional(),
		field.String("f10").Optional(),
		field.String("f11").Optional(),
		field.String("f12").Optional(),
		field.String("f13").Optional(),
		field.String("f14").Optional(),
		field.String("f15").Optional(),
		field.String("f16").Optional(),
		field.String("f17").Optional(),
		field.String("f18").Optional(),
		field.String("f19").Optional(),
		field.String("f20").Optional(),
		field.String("f21").Optional(),
		field.String("f22").Optional(),
		field.String("f23").Optional(),
		field.String("f24").Optional(),
		field.String("f25").Optional(),
		field.String("f26").Optional(),
		field.String("f27").Optional(),
		field.String("f28").Optional(),
		field.String("f29").Optional(),
		field.String("f30").Optional(),
		field.Time("create_time").Optional(),
		field.String("remarks").Optional(),
		field.String("status").Optional(),
		field.Int32("st1").Optional(),
		field.Int32("st2").Optional(),
		field.Int32("st3").Optional(),
		field.Int32("st4").Optional(),
		field.Int32("st5").Optional()}
}
func (EmStockBasicInfo) Edges() []ent.Edge {
	return nil
}

func (EmStockBasicInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "em_stock_basic_info"},
	}
}
