package domain

import "encoding/json"

type UpdateOneReq[T any] struct {
	CallSdkReq
	Data T `json:"data"`
}

func (req *UpdateOneReq[T]) ToJson() []byte {
	bdata, _ := json.Marshal(req)
	return bdata
}

type UpdateWrapper[T any] struct {
	Schema     string         `json:"schema"`
	Collection string         `json:"collection"`
	Query      map[string]any `json:"query"`
	UpdateSet  map[string]any `json:"update_set"`
}

func (query *UpdateWrapper[T]) ToJson() []byte {
	bdata, _ := json.Marshal(query)
	return bdata
}

func NewUpdateWrapper[T any](schema string, collection string) UpdateWrapper[T] {
	return UpdateWrapper[T]{
		Schema:     schema,
		Collection: collection,
		Query:      map[string]any{},
		UpdateSet: map[string]any{
			"$set": map[string]any{},
		},
	}
}

func (query *UpdateWrapper[T]) Eq(fieldName string, value any) *UpdateWrapper[T] {
	query.Query[fieldName] = map[string]any{
		"$eq": value,
	}
	return query
}

func (query *UpdateWrapper[T]) Gt(fieldName string, value any) *UpdateWrapper[T] {
	query.Query[fieldName] = map[string]any{
		"$gt": value,
	}
	return query
}

func (query *UpdateWrapper[T]) Lt(fieldName string, value any) *UpdateWrapper[T] {
	query.Query[fieldName] = map[string]any{
		"$lt": value,
	}
	return query
}

func (query *UpdateWrapper[T]) In(fieldName string, value []any) *UpdateWrapper[T] {
	query.Query[fieldName] = map[string]any{
		"$in": value,
	}
	return query
}

func (query *UpdateWrapper[T]) Search(fieldName string, value any) *UpdateWrapper[T] {
	query.Query[fieldName] = map[string]any{
		"$text": map[string]any{
			"$search": value,
		},
	}
	return query
}

func (query *UpdateWrapper[T]) Expr(fieldName string, value string) *UpdateWrapper[T] {
	query.Query[fieldName] = map[string]any{
		"$expr": value,
	}
	return query
}

func (query *UpdateWrapper[T]) Or(value map[string]any) *UpdateWrapper[T] {
	if _, ok := query.Query["$or"]; !ok {
		query.Query["$or"] = []map[string]any{}
	}
	query.Query["$or"] = append(query.Query["$or"].([]map[string]any), value)
	return query
}

func (query *UpdateWrapper[T]) Set(fieldName string, value any) *UpdateWrapper[T] {
	query.UpdateSet["$set"].(map[string]any)[fieldName] = value
	return query
}
