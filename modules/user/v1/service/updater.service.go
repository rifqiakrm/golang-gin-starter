package service

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/user/v1/repository"
	"gin-starter/utils"
)

type UserUpdater struct {
	cfg      config.Config
	userRepo repository.UserRepositoryUseCase
}

type UserUpdaterUseCase interface {
	VerifyOTP(ctx context.Context, userID uuid.UUID, otp string) (bool, error)
	ResendOTP(ctx context.Context, userID uuid.UUID) error
	ChangePassword(ctx context.Context, userID uuid.UUID, oldPassword, newPassword string) error
	ForgotPasswordRequest(ctx context.Context, email string) error
	ForgotPassword(ctx context.Context, userID uuid.UUID, newPassword string) error
	Update(ctx context.Context, user *entity.User) error
}

func NewUserUpdater(
	cfg config.Config,
	userRepo repository.UserRepositoryUseCase,
) *UserUpdater {
	return &UserUpdater{
		cfg:      cfg,
		userRepo: userRepo,
	}
}

func (uu *UserUpdater) VerifyOTP(ctx context.Context, userID uuid.UUID, otp string) (bool, error) {
	user, err := uu.userRepo.GetUserByID(ctx, userID)

	if err != nil {
		return false, err
	}

	if user == nil {
		return false, entity.ErrUserNotFound.Error
	}

	if user.OTP.Valid && (user.OTP.String != otp) {
		return false, nil
	}

	if err := uu.userRepo.UpdateOTP(ctx, user, ""); err != nil {
		return false, err
	}

	return true, nil
}

func (uu *UserUpdater) ResendOTP(ctx context.Context, userID uuid.UUID) error {
	user, err := uu.userRepo.GetUserByID(ctx, userID)

	if err != nil {
		return err
	}

	otp := utils.GenerateOTP(4)

	if err := uu.userRepo.UpdateOTP(ctx, user, otp); err != nil {
		return err
	}

	t, err := template.ParseFiles("./template/email/send_otp.html")
	if err != nil {
		log.Println(fmt.Errorf("failed to load email template: %w", err))
		return err
	}

	var body bytes.Buffer

	err = t.Execute(&body, struct {
		Name string
		OTP  string
	}{
		Name: user.Name,
		OTP:  otp,
	})
	if err != nil {
		log.Println(fmt.Errorf("failed to execute email data: %w", err))
	}

	if err := utils.SendMail(uu.cfg, user.Email, "", "Login Verification", body.String()); err != nil {
		return err
	}

	return nil
}

func (uu *UserUpdater) ChangePassword(ctx context.Context, userID uuid.UUID, oldPassword, newPassword string) error {
	user, err := uu.userRepo.GetUserByID(ctx, userID)

	if err != nil {
		return err
	}

	if user == nil {
		return entity.ErrUserNotFound.Error
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return entity.ErrOldPasswordMismatch.Error
	}

	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.MinCost)
	if err != nil {
		return err
	}

	if err := uu.userRepo.ChangePassword(ctx, user, string(newPasswordHash)); err != nil {
		return err
	}

	return nil
}

func (uu *UserUpdater) ForgotPasswordRequest(ctx context.Context, email string) error {
	user, err := uu.userRepo.GetUserByEmail(ctx, email)

	if err != nil {
		return err
	}

	user.ForgotPasswordToken = utils.StringToNullString(utils.RandStringBytes(30))

	if err := uu.userRepo.Update(ctx, user); err != nil {
		return err
	}

	t, err := template.ParseFiles("./template/email/forgot_password.html")
	if err != nil {
		log.Println(fmt.Errorf("failed to load email template: %w", err))
		return err
	}

	var body bytes.Buffer

	err = t.Execute(&body, struct {
		Name string
		URL  string
	}{
		Name: user.Name,
		URL:  fmt.Sprintf("%s/%s", uu.cfg.URL.ForgotPasswordURL, user.ForgotPasswordToken.String),
	})

	if err != nil {
		log.Println(fmt.Errorf("failed to execute email data: %w", err))
	}

	if err := utils.SendMail(uu.cfg, user.Email, "", "Password Reset", body.String()); err != nil {
		return err
	}

	return nil
}

func (uu *UserUpdater) ForgotPassword(ctx context.Context, userID uuid.UUID, newPassword string) error {
	user, err := uu.userRepo.GetUserByID(ctx, userID)

	if err != nil {
		return err
	}

	if user == nil {
		return entity.ErrUserNotFound.Error
	}

	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.MinCost)
	if err != nil {
		return err
	}

	if err := uu.userRepo.ChangePassword(ctx, user, string(newPasswordHash)); err != nil {
		return err
	}

	return nil
}

func (uu *UserUpdater) Update(ctx context.Context, user *entity.User) error {
	if err := uu.userRepo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}
