package datasources

import (
	"errors"
	"fmt"
	"os"

	"github.com/aaalik/api-alik/internal/api/constants"
	"github.com/aaalik/api-alik/pkg/alog"
	"github.com/jmoiron/sqlx"
)

// Prepare prepare sql statements or exit api if fails or error
func Prepare(db *sqlx.DB, query string) *sqlx.Stmt {
	s, err := db.Preparex(query)
	if err != nil {
		alog.Logger.Error(errors.New(fmt.Sprintf("error while preparing statement: %s", err)))
		alog.Logger.Error(errors.New(fmt.Sprintf("query: %s", query)))

		os.Exit(constants.ExitPrepareStmtFail)
	}
	return s
}

// PrepareNamed prepare sql statements with named bindvars or exit api if fails or error
func PrepareNamed(db *sqlx.DB, query string) *sqlx.NamedStmt {
	s, err := db.PrepareNamed(query)
	if err != nil {
		alog.Logger.Error(errors.New(fmt.Sprintf("error while preparing statement: %s", err)))
		alog.Logger.Error(errors.New(fmt.Sprintf("query: %s", query)))

		os.Exit(constants.ExitPrepareStmtFail)
	}
	return s
}
