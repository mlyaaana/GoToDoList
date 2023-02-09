package database

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"testing"
	"todolist/database/models"
)

func TestDatabase(t *testing.T) {
	db, err := New()
	require.NoError(t, err)

	var usersModels []models.User
	err = db.Find(&usersModels).Error
	require.NoError(t, err)
	spew.Dump(usersModels)
	//user := &models.User{}
	//err = db.Where(&models.User{Name: "p sune"}).First(user).Error
	//require.NoError(t, err)
	//spew.Dump(user)

	//err = db.Create(&models.User{
	//	Id:   uuid.New().String(),
	//	Name: random.String(10),
	//}).Error
	//require.NoError(t, err)
}
