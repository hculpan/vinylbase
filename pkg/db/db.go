package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var dbase *sql.DB = nil

type QueryFunc func(index int, rows *sql.Rows) error

func InitDb(dbName, tursoToken string) error {
	if dbase != nil {
		return errors.New("attempted to initialize database that is already open")
	}

	db, err := sql.Open("libsql", fmt.Sprintf("%s?authToken=%s", dbName, tursoToken))
	if err != nil {
		return err
	}

	dbase = db
	return nil
}

func Execute(sql string) (int, error) {
	if dbase == nil {
		return 0, errors.New("cannot execute statement, database not open")
	}

	result, err := dbase.Exec(sql)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	return int(rows), err
}

func Query(sql string, queryFunc QueryFunc) (int, error) {
	rows, err := dbase.Query(sql)
	if err != nil {
		return 0, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		count++
		err := queryFunc(count, rows)
		if err != nil {
			break
		}
	}

	if err := rows.Err(); err != nil {
		return count, fmt.Errorf("error during rows iteration: %w", err)
	}

	return count, nil
}

func CloseDb() {
	if dbase != nil {
		dbase.Close()
		dbase = nil
	}
}
