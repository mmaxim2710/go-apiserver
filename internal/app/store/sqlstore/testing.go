package sqlstore

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/mmaxim2710/firstWebApp/internal/app/config"
	"os"
	"strings"
	"testing"
)

func TestDB(t *testing.T) (*sql.DB, func(...string)) {
	t.Helper()
	newConfig, err := config.NewConfig("../../../../configs/apiserver_test.yaml")
	if err != nil {
		t.Fatal(err)
	}

	pwd, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		t.Fatal("Not specified DB_PASSWORD env variable")
	}

	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			newConfig.DB.Host, newConfig.DB.Port, newConfig.DB.User, pwd, newConfig.DB.DBName,
		),
	)

	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			_, _ = db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
		}

		_ = db.Close()
	}
}
