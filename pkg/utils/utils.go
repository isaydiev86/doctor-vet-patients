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

// NilIfEmpty проверяет значение на "пустоту" и возвращает nil, если оно пустое.
func NilIfEmpty(value interface{}) interface{} {
	switch v := value.(type) {
	case string:
		if v == "" {
			return nil
		}
	case int, int8, int16, int32, int64:
		if v == 0 {
			return nil
		}
	case uint, uint8, uint16, uint32, uint64:
		if v == 0 {
			return nil
		}
	case float32, float64:
		if v == 0.0 {
			return nil
		}
	case nil:
		return nil
	}
	return value
}
