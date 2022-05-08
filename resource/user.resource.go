package resource

import (
	"mime/multipart"
	"os"

	"gin-starter/entity"
	"gin-starter/utils"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

type UpdateUserRequest struct {
	ID          string                `form:"id" json:"id"`
	Name        string                `form:"name" json:"name"`
	Email       string                `form:"email" json:"email"`
	DOB         string                `form:"dob" json:"dob"`
	PhoneNumber string                `form:"phone_number" json:"phone_number"`
	Photo       *multipart.FileHeader `form:"photo" json:"photo"`
}

type UserProfile struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	OTPIsNull   bool   `json:"otp_is_null"`
	PhoneNumber string `json:"phone_number"`
	DOB         string `json:"dob"`
	Status      string `json:"status"`
	Photo       string `json:"photo"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func NewUserProfile(user *entity.User) *UserProfile {
	otpIsNull := false
	if user.OTP.String != "" {
		otpIsNull = true
	}

	dob := "1970-01-01"
	if user.DOB.Valid {
		dob = user.DOB.Time.Format(timeFormat)
	}

	return &UserProfile{
		ID:          user.ID.String(),
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		DOB:         dob,
		Photo:       utils.ImageFullPath(os.Getenv("IMAGE_HOST"), user.Photo),
		Status:      user.Status,
		OTPIsNull:   otpIsNull,
		CreatedAt:   user.CreatedAt.Format(timeFormat),
		UpdatedAt:   user.UpdatedAt.Format(timeFormat),
	}
}

type ChangePasswordRequest struct {
	OldPassword             string `form:"old_password" json:"old_password" binding:"required"`
	NewPassword             string `form:"new_password" json:"new_password" binding:"required"`
	NewPasswordConfirmation string `form:"new_password_confirmation" json:"new_password_confirmation" binding:"required"`
}

type ForgotPasswordRequest struct {
	Email string `form:"email" json:"email" binding:"required"`
}

type ForgotPasswordChangeRequest struct {
	Token                   string `form:"token" json:"token" binding:"required"`
	NewPassword             string `form:"new_password" json:"new_password" binding:"required"`
	NewPasswordConfirmation string `form:"new_password_confirmation" json:"new_password_confirmation" binding:"required"`
}

type GetUserByForgotPasswordTokenRequest struct {
	Token string `uri:"token" json:"token" binding:"required"`
}

type VerifyOTPRequest struct {
	Code string `form:"code" json:"code" binding:"required"`
}
