package service_test

import (
	"context"
	"database/sql"
	"gin-starter/modules/auth/v1/service"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"

	"gin-starter/config"
	"gin-starter/entity"
	mockRepo "gin-starter/test/mock/modules/auth/repository"
	"gin-starter/utils"
)

type AuthServiceTestSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller

	cfg            config.Config
	authRepository *mockRepo.MockAuthRepositoryUseCase
	authService    *service.AuthService
}

func TestAuthServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuite))
}

func (suite *AuthServiceTestSuite) BeforeTest(string, string) {
	suite.mockCtrl = gomock.NewController(suite.T())

	suite.cfg = config.Config{}
	suite.authRepository = mockRepo.NewMockAuthRepositoryUseCase(suite.mockCtrl)

	suite.authService = service.NewAuthService(
		suite.cfg,
		suite.authRepository,
	)
}

func (suite *AuthServiceTestSuite) AfterTest(string, string) {
	defer suite.mockCtrl.Finish()
}

func (suite *AuthServiceTestSuite) TestNewAuthService() {
	suite.Run("successfully create a new instance of AuthService", func() {
		suite.NotNil(suite.authService)
	})
}

func (suite *AuthServiceTestSuite) TestAuthService_AuthValidate() {

	suite.Run("successfully auth validate", func() {
		if err := os.MkdirAll("template/email", os.ModePerm); err != nil {
			panic(err)
		}

		tempDir, err := os.Create("template/email/send_otp.html")

		if err != nil {
			panic(err)
		}

		defer tempDir.Close()

		defer func() {
			if err := os.RemoveAll("template"); err != nil {
				panic(err)
			}
		}()

		email := "test@mail.com"
		password := "test123"
		otp := "1234"

		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		returnedUser := &entity.User{
			ID:                  uuid.New(),
			Name:                "test",
			Email:               email,
			Password:            string(passwordHash),
			PhoneNumber:         "",
			Photo:               "",
			DOB:                 sql.NullTime{},
			OTP:                 utils.StringToNullString(otp),
			Status:              "",
			ForgotPasswordToken: sql.NullString{},
			UserRole:            nil,
		}

		suite.authRepository.EXPECT().GetUserByEmail(context.Background(), email).Return(returnedUser, nil)

		suite.authRepository.EXPECT().UpdateOTP(context.Background(), returnedUser, gomock.Any()).Return(nil)

		validate, err := suite.authService.AuthValidate(context.Background(), email, password)

		if err != nil {
			panic(err)
		}

		suite.NotNil(validate)
		suite.Nil(err)
	})
}
