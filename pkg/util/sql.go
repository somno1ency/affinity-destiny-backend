package util

import (
	"database/sql"
	"time"
)

func ConvertString(str string) sql.NullString {
	if str == "" {
		return sql.NullString{
			String: "",
			Valid:  false,
		}
	}

	return sql.NullString{
		String: str,
		Valid:  true,
	}
}

func ConvertInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: i,
		Valid: true,
	}
}

func ConvertTime(time time.Time) sql.NullTime {
	if time.IsZero() {
		return sql.NullTime{
			Time:  time,
			Valid: false,
		}
	}

	return sql.NullTime{
		Time:  time,
		Valid: true,
	}
}
