package teststore_test

import (
	"github.com/mmaxim2710/firstWebApp/internal/app/model"
	"github.com/mmaxim2710/firstWebApp/internal/app/store"
	"github.com/mmaxim2710/firstWebApp/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	err := s.User().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "test_user@example.org"

	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

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
