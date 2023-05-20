package apiserver

import (
	"database/sql"
	"fmt"
	"github.com/mmaxim2710/firstWebApp/internal/app/config"
	"github.com/mmaxim2710/firstWebApp/internal/app/store/sqlstore"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func Start(config *config.Config, logger *logrus.Logger) error {
	db, err := newDB(config)
	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}(db)

	store := sqlstore.New(db, logger)
	s := newServer(store, logger)
	return http.ListenAndServe(config.Server.BindAddr, s)
}

func newDB(dbconf *config.Config) (*sql.DB, error) {
	pwd, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return nil, ErrEnvVariableNotFound
	}

	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbconf.DB.Host, dbconf.DB.Port, dbconf.DB.User, pwd, dbconf.DB.DBName,
		),
	)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
