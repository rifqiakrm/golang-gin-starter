package service

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"

	"golang.org/x/crypto/bcrypt"

	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/auth/v1/repository"
	"gin-starter/utils"
)

type AuthService struct {
	cfg      config.Config
	authRepo repository.AuthRepositoryUseCase
}

type AuthUseCase interface {
	AuthValidate(ctx context.Context, email, password string) (*entity.User, error)
	AuthValidateCMS(ctx context.Context, email, password string) (*entity.User, error)
	GenerateAccessToken(ctx context.Context, user *entity.User) (*entity.Token, error)
	GenerateAccessTokenCMS(ctx context.Context, user *entity.User) (*entity.Token, error)
}

func NewAuthService(
	cfg config.Config,
	authRepo repository.AuthRepositoryUseCase,
) *AuthService {
	return &AuthService{
		cfg:      cfg,
		authRepo: authRepo,
	}
}

func (as *AuthService) AuthValidate(ctx context.Context, email, password string) (*entity.User, error) {
	user, err := as.authRepo.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, entity.ErrPasswordMismatch.Error
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

		if err != nil {
			return nil, entity.ErrPasswordMismatch.Error
		}
	}

	otp := utils.GenerateOTP(4)

	if user.Email == "user-test@gmail.com" {
		otp = "1234"
	}

	if err := as.authRepo.UpdateOTP(ctx, user, otp); err != nil {
		return nil, err
	}

	t, err := template.ParseFiles("./template/email/send_otp.html")
	if err != nil {
		log.Println(fmt.Errorf("failed to load email template: %w", err))
		return nil, err
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
		return nil, err
	}

	payload := entity.EmailPayload{
		To:       user.Email,
		Subject:  "Login Verification",
		Content:  body.String(),
		Category: entity.EmailCategorySendOTP,
	}

	if err := utils.SendTopic(context.Background(), as.cfg, "send-email", payload); err != nil {
		log.Println(err)
	}

	return user, nil
}

func (as *AuthService) AuthValidateCMS(ctx context.Context, email, password string) (*entity.User, error) {
	user, err := as.authRepo.GetAdminByEmail(ctx, email)

	if err != nil {
		return user, err
	}

	if user == nil {
		return nil, entity.ErrPasswordMismatch.Error
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

		if err != nil {
			return nil, entity.ErrPasswordMismatch.Error
		}
	}

	return user, nil
}

func (as *AuthService) GenerateAccessToken(ctx context.Context, user *entity.User) (*entity.Token, error) {
	token, err := utils.JWTEncode(as.cfg, user.ID, as.cfg.JWTConfig.Issuer)

	if err != nil {
		return nil, err
	}

	return &entity.Token{
		Token: token,
	}, nil
}

func (as *AuthService) GenerateAccessTokenCMS(ctx context.Context, user *entity.User) (*entity.Token, error) {
	token, err := utils.JWTEncode(as.cfg, user.ID, as.cfg.JWTConfig.IssuerCMS)

	if err != nil {
		return nil, err
	}

	return &entity.Token{
		Token: token,
	}, nil
}
