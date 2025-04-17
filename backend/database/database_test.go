package database

import (
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/arly/arlyApi/config"
	"github.com/arly/arlyApi/models"
	"github.com/arly/arlyApi/utilities"
	"github.com/stretchr/testify/assert"
)

func TestConnectDb(t *testing.T) {
	utilities.InitializeLogger()

	os.Setenv("ENV", "test")

	patchConfig := gomonkey.ApplyFunc(config.DbConfig, func() string {
		return "postgres://root:examplepassword@localhost:5432/arlydb?sslmode=disable"
	})
	defer patchConfig.Reset()

	os.Setenv("ADMIN_FIRSTNAME", "Test")
	os.Setenv("ADMIN_LASTNAME", "Admin")
	os.Setenv("ADMIN_EMAIL", "test_admin@example.com")
	os.Setenv("ADMIN_PASSWORD", "password")

	ConnectDb()

	assert.NotNil(t, Database.Db, "Database connection should be established")

	hasUserTable := Database.Db.Migrator().HasTable(&models.User{})
	assert.True(t, hasUserTable, "Users table should exist after migrations")

}

func TestSeedAdminUser_CreatesAdminWhenNotExists(t *testing.T) {
	utilities.InitializeLogger()

	os.Setenv("ENV", "test")
	os.Setenv("ADMIN_FIRSTNAME", "SeedTest")
	os.Setenv("ADMIN_LASTNAME", "User")
	os.Setenv("ADMIN_EMAIL", "seed_test_admin@example.com")
	os.Setenv("ADMIN_PASSWORD", "seedpassword")

	patchConfig := gomonkey.ApplyFunc(config.DbConfig, func() string {
		return "postgres://root:examplepassword@localhost:5432/arlydb?sslmode=disable"
	})
	defer patchConfig.Reset()

	ConnectDb()
	db := Database.Db

	db.Exec("DELETE FROM users")

	seedAdminUser(db)

	var admin models.User
	result := db.First(&admin, 1)
	assert.Nil(t, result.Error, "Seeded admin user should be found")
	assert.Equal(t, "SeedTest", admin.FirstName, "Admin first name should match")
	assert.Equal(t, "User", admin.LastName, "Admin last name should match")
	assert.Equal(t, "seed_test_admin@example.com", admin.Email, "Admin email should match")

}

func TestSeedAdminUser_DoesNothingIfAdminExists(t *testing.T) {
	utilities.InitializeLogger()

	os.Setenv("ENV", "test")
	os.Setenv("ADMIN_FIRSTNAME", "SeedTest")
	os.Setenv("ADMIN_LASTNAME", "User")
	os.Setenv("ADMIN_EMAIL", "seed_test_admin@example.com")
	os.Setenv("ADMIN_PASSWORD", "seedpassword")

	patchConfig := gomonkey.ApplyFunc(config.DbConfig, func() string {
		return "postgres://root:examplepassword@localhost:5432/arlydb?sslmode=disable"
	})
	defer patchConfig.Reset()

	ConnectDb()
	db := Database.Db

	seedAdminUser(db)
	var admin models.User
	_ = db.First(&admin, 1)
	originalFirstName := admin.FirstName
	seedAdminUser(db)
	var adminAfter models.User
	_ = db.First(&adminAfter, 1)
	assert.Equal(t, originalFirstName, adminAfter.FirstName, "Admin user should remain unchanged if already exists")
}
