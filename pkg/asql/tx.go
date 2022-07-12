package asql

import (
	"github.com/jmoiron/sqlx"
	"log"
)

// NamedStmtTx determines whether to use transaction or not
func NamedStmtTx(s *sqlx.NamedStmt, tx *sqlx.Tx) *sqlx.NamedStmt {
	if tx != nil {
		return tx.NamedStmt(s)
	}
	return s
}

// StmtTx determines whether to use transaction or not
func StmtTx(s *sqlx.Stmt, tx *sqlx.Tx) *sqlx.Stmt {
	if tx != nil {
		return tx.Stmtx(s)
	}
	return s
}

// ReleaseTx clean db transaction by commit if no error, or rollback if an error occurred
func ReleaseTx(tx *sqlx.Tx, err *error) {
	if *err != nil {
		// if an error occurred, rollback transaction
		errRollback := tx.Rollback()
		if errRollback != nil {
			log.Printf("unable to rollback transaction: %s", errRollback)
		} else {
			log.Print("transaction rolled back")
		}
		return
	}
	// else, commit transaction
	errCommit := tx.Commit()
	if errCommit != nil {
		log.Printf("unable to commit transaction: %s", errCommit)
	}
}
