package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/utils"
)

func main() {
	cfg, err := config.LoadConfig(".env")
	checkError(err)

	db, err := utils.NewPostgresGormDB(&cfg.Postgres)
	checkError(err)

	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	createSampleUser(db)
	createAdminUser(db)

}

func createSampleUser(db *gorm.DB) {
	userID := uuid.New()

	dob, _ := utils.DateStringToTime("1996-11-04")

	if err := db.WithContext(context.Background()).
		Model(&entity.User{}).
		Create(entity.NewUser(
			userID,
			"Bayu Novianto",
			"bayunoviantoo9@gmail.com",
			"test123",
			utils.TimeToNullTime(dob),
			"",
			"0895346419497",
			"system",
		)).
		Error; err != nil {
		panic(err)
	}

	userID2 := uuid.New()

	dob2, _ := utils.DateStringToTime("1996-11-04")

	if err := db.WithContext(context.Background()).
		Model(&entity.User{}).
		Create(entity.NewUser(
			userID2,
			"User Test",
			"user-test@gmail.com",
			"testingApp23!",
			utils.TimeToNullTime(dob2),
			"",
			"",
			"system",
		)).
		Error; err != nil {
		panic(err)
	}
}

func createAdminUser(db *gorm.DB) {
	roleID := uuid.New()

	if err := db.WithContext(context.Background()).
		Model(&entity.Role{}).
		Create(entity.NewRole(
			roleID,
			"Super Admin",
			"system",
		)).
		Error; err != nil {
		panic(err)
	}

	userID := uuid.New()

	dobUser, _ := utils.DateStringToTime("1999-10-30")

	if err := db.WithContext(context.Background()).
		Model(&entity.User{}).
		Create(entity.NewUser(
			userID,
			"Rifqi Akram",
			"rifqiakram57@gmail.com",
			"testing1234",
			utils.TimeToNullTime(dobUser),
			"",
			"0895346419497",
			"system",
		)).
		Error; err != nil {
		panic(err)
	}

	if err := db.WithContext(context.Background()).
		Model(&entity.UserRole{}).
		Create(entity.NewUserRole(
			userID,
			roleID,
			"system",
		)).
		Error; err != nil {
		panic(err)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
