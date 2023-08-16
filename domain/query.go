package domain

import "encoding/json"

type QueryWrapper[T any] struct {
	Method     string         `json:"method"`
	Schema     string         `json:"schema"`
	Collection string         `json:"collection"`
	Query      map[string]any `json:"query"`
}

func (query *QueryWrapper[T]) ToJson() []byte {
	bdata, _ := json.Marshal(query)
	return bdata
}

func NewQueryWrapper[T any](schema string, collection string) QueryWrapper[T] {
	return QueryWrapper[T]{
		Schema:     schema,
		Collection: collection,
	}
}

func (query *QueryWrapper[T]) Eq(fieldName string, value any) *QueryWrapper[T] {
	query.Query[fieldName] = map[string]any{
		"$eq": value,
	}
	return query
}

func (query *QueryWrapper[T]) Gt(fieldName string, value any) *QueryWrapper[T] {
	query.Query[fieldName] = map[string]any{
		"$gt": value,
	}
	return query
}

func (query *QueryWrapper[T]) Lt(fieldName string, value any) *QueryWrapper[T] {
	query.Query[fieldName] = map[string]any{
		"$lt": value,
	}
	return query
}

func (query *QueryWrapper[T]) In(fieldName string, value []any) *QueryWrapper[T] {
	query.Query[fieldName] = map[string]any{
		"$in": value,
	}
	return query
}

func (query *QueryWrapper[T]) Search(fieldName string, value any) *QueryWrapper[T] {
	query.Query[fieldName] = map[string]any{
		"$text": map[string]any{
			"$search": value,
		},
	}
	return query
}

func (query *QueryWrapper[T]) Expr(fieldName string, value string) *QueryWrapper[T] {
	query.Query[fieldName] = map[string]any{
		"$expr": value,
	}
	return query
}
