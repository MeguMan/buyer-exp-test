package apiserver

import (
	"database/sql"
	"fmt"
	"github.com/MeguMan/buyer-exp-test/internal/app/emailsender"
	"github.com/MeguMan/buyer-exp-test/internal/app/store/postgres_store"
	"net/http"
)

func Start(databaseURL string) error {
	fmt.Print(databaseURL)
	db, err := newDB(databaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	s := postgres_store.New(db)
	es := emailsender.New()
	server := NewServer(s, *es)
	return http.ListenAndServe(":8080", server)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
