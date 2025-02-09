package repositories_test

import (
	"placelists/internal/entities"
	"placelists/internal/storage/repositories"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryGetByUsernameValid(t *testing.T) {
	db := InitDB()
	defer DropDB(db)
	r := repositories.NewRepository(db)

	u := &entities.User{ID: uuid.New(), Username: "main", PublicID: "main"}
    db.Create(u)

    foundUser, err := r.User.GetByUsername(u.Username)

	assert.NoError(t, err)
	assert.Equal(t, u.ID, foundUser.ID)
}

func TestUserRepositoryGetByUsernameInvalidNotExistingUser(t *testing.T) {
	db := InitDB()
	defer DropDB(db)
	r := repositories.NewRepository(db)

    _, err := r.User.GetByUsername("main")

	assert.Error(t, err)
}

func TestUserRepositoryCreateValid(t *testing.T) {
	db := InitDB()
	defer DropDB(db)
	r := repositories.NewRepository(db)

	u := &entities.User{ID: uuid.New(), Username: "main", PublicID: "main"}

	err := r.User.Create(u)

	var createdUser *entities.User
	db.First(&createdUser, u.ID)

	assert.NoError(t, err)
	assert.Equal(t, u.ID, createdUser.ID)
}

func TestUserRepositoryCreateInvalidExistingUsername(t *testing.T) {
	db := InitDB()
	defer DropDB(db)
	r := repositories.NewRepository(db)

	u := &entities.User{ID: uuid.New(), Username: "main", PublicID: "main"}

	err1 := r.User.Create(u)
	err2 := r.User.Create(u)

	assert.NoError(t, err1)
	assert.Error(t, err2)
}

func TestUserRepositoryUpdateValid(t *testing.T) {
	db := InitDB()
	defer DropDB(db)
	r := repositories.NewRepository(db)

	u := &entities.User{ID: uuid.New(), Username: "main", PublicID: "main"}
    db.Create(u)

    userUpdate := &entities.User{ID: u.ID, Username: "second", PublicID: "second"}

    err := r.User.Update(userUpdate)

	var updatedUser *entities.User
	db.First(&updatedUser, u.ID)

	assert.NoError(t, err)
	assert.Equal(t, u.ID, updatedUser.ID)
	assert.Equal(t, updatedUser.Username, userUpdate.Username)
}

func TestUserRepositoryUpdateInvalidExistingUsername(t *testing.T) {
	db := InitDB()
	defer DropDB(db)
	r := repositories.NewRepository(db)

	user1 := &entities.User{ID: uuid.New(), Username: "main", PublicID: "main"}
	user2 := &entities.User{ID: uuid.New(), Username: "second", PublicID: "second"}
    db.Create(user1)
    db.Create(user2)

    userUpdate := &entities.User{ID: user1.ID, Username: user2.Username, PublicID: user2.PublicID}

    err := r.User.Update(userUpdate)

	assert.Error(t, err)
}
