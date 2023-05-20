package store

import (
	"fmt"
	"github.com/mmaxim2710/firstWebApp/internal/app/config"
	"github.com/sirupsen/logrus"
	"strings"
	"testing"
)

func TestStore(t *testing.T) (*Store, func(...string)) {
	t.Helper()
	newConfig, err := config.NewConfig("../../../configs/apiserver_test.yaml")
	if err != nil {
		t.Fatal(err)
	}
	s := New(newConfig, logrus.New())
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		if err := s.Close(); err != nil {
			t.Fatal(err)
		}
	}
}
