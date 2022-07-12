package entities

import (
	"database/sql"
)

type NullString struct {
	sql.NullString
}

type NullInt64 struct {
	sql.NullInt64
}

type NullFloat64 struct {
	sql.NullFloat64
}

type NullBool struct {
	sql.NullBool
}
