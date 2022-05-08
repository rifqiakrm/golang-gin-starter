package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"gin-starter/utils"
)

const (
	userTableName = "main.users"
)

type User struct {
	ID                  uuid.UUID      `json:"id"`
	Name                string         `json:"name"`
	Email               string         `json:"email"`
	Password            string         `json:"password"`
	PhoneNumber         string         `json:"phone_number"`
	Photo               string         `json:"photo"`
	DOB                 sql.NullTime   `json:"dob"`
	OTP                 sql.NullString `json:"otp"`
	Status              string         `json:"status"`
	ForgotPasswordToken sql.NullString `json:"forgot_password_token"`
	UserRole            *UserRole      `foreignKey:"ID" associationForeignKey:"UserID"`
	Auditable
}

// TableName specifies table name
func (model *User) TableName() string {
	return userTableName
}

func NewUser(
	id uuid.UUID,
	name string,
	email string,
	password string,
	dob sql.NullTime,
	photo string,
	phoneNumber string,
	createdBy string,
) *User {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return &User{
		ID:          id,
		Name:        name,
		Email:       email,
		Password:    string(passwordHash),
		PhoneNumber: phoneNumber,
		Photo:       photo,
		DOB:         dob,
		OTP:         sql.NullString{},
		Status:      "ACTIVATED",
		Auditable:   NewAuditable(createdBy),
	}
}

// MapUpdateFrom mapping from model
func (model *User) MapUpdateFrom(from *User) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"name":         model.Name,
			"email":        model.Email,
			"phone_number": model.PhoneNumber,
			"photo":        model.Photo,
			"otp":          model.OTP,
			"status":       model.Status,
			"updated_at":   model.UpdatedAt,
		}
	}

	mapped := make(map[string]interface{})

	if model.Name != from.Name {
		mapped["name"] = from.Name
	}

	if model.Email != from.Email {
		mapped["email"] = from.Email
	}

	if model.PhoneNumber != from.PhoneNumber {
		mapped["phone_number"] = from.PhoneNumber
	}

	if model.DOB != from.DOB {
		mapped["dob"] = from.DOB
	}

	if (model.Photo != from.Photo) && from.Photo != "" {
		mapped["photo"] = from.Photo
	}

	if model.OTP != from.OTP {
		mapped["otp"] = utils.StringToNullString(from.OTP.String)
	}

	if model.Status != from.Status {
		mapped["status"] = from.Status
	}

	if model.ForgotPasswordToken != from.ForgotPasswordToken {
		mapped["forgot_password_token"] = from.ForgotPasswordToken
	}

	mapped["updated_at"] = time.Now()
	return &mapped
}
