package dtos

import (
	"fmt"
	"strings"
)

type Queryable struct {
	Limit   int                    `json:"limit"`
	Offset  int                    `json:"offset"`
	Filters map[string]interface{} `json:"filters"`
}

func (q *Queryable) BuildQueryWithFilters(query string) (string, []interface{}) {
	var whereClauses []string
	var args []interface{}
	paramCount := 1
	for key, value := range q.Filters {
		whereClauses = append(whereClauses, fmt.Sprintf("%s = $%d", key, paramCount))
		args = append(args, value)
		paramCount++
	}

	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	if q.Limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", paramCount)
		args = append(args, q.Limit)
		paramCount++
	}
	if q.Offset > 0 {
		query += fmt.Sprintf(" OFFSET $%d", paramCount)
		args = append(args, q.Offset)
	}

	return query, args
}
