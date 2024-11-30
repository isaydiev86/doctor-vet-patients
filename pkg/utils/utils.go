package utils

import (
	"database/sql"
	"strconv"
	"time"
)

// nolint (gocritic)
func ToPtr[T comparable](v T) *T {
	return &v
}

// nolint (gocritic)
func FromPtr[T any](v *T) T {
	if v == nil {
		var dv T
		return dv
	}
	return *v
}

func Ter[T any](x bool, a, b T) T {
	if x {
		return a
	}
	return b
}

func ValidNullBool(b bool) sql.NullBool          { return sql.NullBool{Bool: b, Valid: true} }
func ValidNullString(s string) sql.NullString    { return sql.NullString{String: s, Valid: true} }
func ValidNullTime(t time.Time) sql.NullTime     { return sql.NullTime{Time: t, Valid: true} }
func ValidNullInt64(i int64) sql.NullInt64       { return sql.NullInt64{Int64: i, Valid: true} }
func ValidNullFloat64(i float64) sql.NullFloat64 { return sql.NullFloat64{Float64: i, Valid: true} }

func ValidInt64ToString(value int64) string {
	return strconv.FormatInt(value, 10)
}
