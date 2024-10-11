package util

import (
	"database/sql"
	"time"
)

func ConvertString(value string) sql.NullString {
	if value == "" {
		return sql.NullString{
			String: "",
			Valid:  false,
		}
	}

	return sql.NullString{
		String: value,
		Valid:  true,
	}
}

func ConvertInt64(value int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: value,
		Valid: true,
	}
}

func ConvertTime(value time.Time) sql.NullTime {
	if value.IsZero() {
		return sql.NullTime{
			Time:  value,
			Valid: false,
		}
	}

	return sql.NullTime{
		Time:  value,
		Valid: true,
	}
}
