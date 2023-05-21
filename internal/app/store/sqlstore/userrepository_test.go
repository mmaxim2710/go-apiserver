package sqlstore_test

import (
	"github.com/mmaxim2710/firstWebApp/internal/app/model"
	"github.com/mmaxim2710/firstWebApp/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t)
	defer teardown("users")
	s := sqlstore.New(db)
	u := model.TestUser(t)
	err := s.User().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t)
	defer teardown("users")

	s := sqlstore.New(db)
	email := "test_user@example.org"

	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email
	err = s.User().Create(u)
	if err != nil {
		t.Fatal(err)
	}
	found, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, email, found.Email)
}
