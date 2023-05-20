package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURl string
)

func TestMain(m *testing.M) {
	databaseURl = os.Getenv("DATABASE_URL")
	if databaseURl == "" {
		databaseURl = "host=localhost dbname=apiserver_test sslmode=disable user=postgres password=Qw2165322710ss!!"
	}

	os.Exit(m.Run())
}
