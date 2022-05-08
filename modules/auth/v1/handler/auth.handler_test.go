package handler_test

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"

	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/auth/v1/handler"
	"gin-starter/modules/auth/v1/service"
	"gin-starter/test/helpers"
	mockRepo "gin-starter/test/mock/modules/auth/repository"
	"gin-starter/utils"
)

const (
	publicKeySample = `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAugp0k/VBdm8TvSbbPl4a
nVDWDUJYjM4U9v9QFE49HhlMoo7srzfRq2ZgmwX5C55ej5UY0kqnId9KRjBN7NlD
iysM181yR/oJFlarjYXla0ebc7RX6pQNOBB5337sxoebpZRSSPe8HxZl21tcC1cN
2C0ZBXcCsGVkHQJkvDB5TR7ncjGLUCMEatrBowiPnEKv4zEfqEchne3TF19guWrZ
zk81y+zxeXa7mJdL0x/ENCNrIuSSZto4fVif7ANw3iLg6C8gqCXum/12R+L6SooX
eJXcDlIa1EjU6RIMR3UxxsidpfTQZDTWDkBz2RPjXbncWx4T45vp4wAncxhhqVuy
HpcZR/u1iEFIWg5IsIHssMFBN6NIMaok29WpMR5Q1rva7c2gUIzA3TX4ZlG7wBag
Jgk3MNnMe55KL/JLCP17SfqP0bMFLkIGDIBHfLTa743YZyWmo+opAUNyxfWviDrf
BPmgSmnAwUM+9nQvdXGZiD2eS2zXHQLHqyIYZ9l2Rtmi2Mw91OttRPGNOwEtQ1bk
0UeaZImL/LTEvEKUkNnYm6eCBBfyqKUWOy+/4EHVLr3Mg54nJn4tBIzCXHULjDNi
YT4LJFdjm0gCgNTKTwbG+6m0+0yO89pDEflkpD4iUjPeapmQ3DlwUZipol+lKXaF
N1OyKBTsst3chHYM2p5TOXMCAwEAAQ==
-----END PUBLIC KEY-----`
	privateKeySample = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`
)

type AuthHandlerTestSuite struct {
	suite.Suite
	cfg            config.Config
	mockCtrl       *gomock.Controller
	subReporter    *helpers.SubReporter
	authRepository *mockRepo.MockAuthRepositoryUseCase
	service        service.AuthUseCase
	authHandler    *handler.AuthHandler
}

func (suite *AuthHandlerTestSuite) BeforeTest(string, string) {
	suite.subReporter = helpers.NewSubReporter(suite.T())
	suite.mockCtrl = gomock.NewController(suite.T())

	suite.cfg = config.Config{
		JWTConfig: config.JWTConfig{
			Public:    "./config/rsa-key/oauth-public.key",
			Private:   "./config/rsa-key/oauth-private.key",
			Issuer:    "test",
			IssuerCMS: "testCms",
		},
	}
	suite.authRepository = mockRepo.NewMockAuthRepositoryUseCase(suite.mockCtrl)
	suite.service = service.NewAuthService(suite.cfg, suite.authRepository)
	suite.authHandler = handler.NewAuthHandler(suite.service)

	if err := os.MkdirAll("template/email", os.ModePerm); err != nil {
		panic(err)
	}
	if err := os.MkdirAll("config/rsa-key/", os.ModePerm); err != nil {
		panic(err)
	}

	tempDir, _ := os.Create("template/email/send_otp.html")

	defer func() {
		if err := tempDir.Close(); err != nil {
			log.Println("can't close file while reading csv:", err)
		}
	}()

	tempDirPublic, _ := os.Create("config/rsa-key/oauth-public.key")

	if _, err := tempDirPublic.Write([]byte(publicKeySample)); err != nil {
		log.Println(err)
	}

	defer func() {
		if err := tempDirPublic.Close(); err != nil {
			log.Println(err)
		}
	}()

	tempDirPrivate, _ := os.Create("config/rsa-key/oauth-private.key")

	if _, err := tempDirPrivate.Write([]byte(privateKeySample)); err != nil {
		log.Println(err)
	}

	defer func() {
		if err := tempDirPrivate.Close(); err != nil {
			log.Println(err)
		}
	}()
}

func (suite *AuthHandlerTestSuite) AfterTest(string, string) {
	defer suite.mockCtrl.Finish()
}

func TestAuthHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthHandlerTestSuite))
}

func (suite *AuthHandlerTestSuite) TestAuthHandler(t *testing.T) {
	suite.Run("successfully create a new auth handler", func() {
		suite.NotNil(suite.authHandler)
	})
}

func (suite *AuthHandlerTestSuite) TestAuthHandler_Login() {
	suite.Run("successfully login", func() {
		defer func() {
			if err := os.RemoveAll("template"); err != nil {
				panic(err)
			}
		}()

		defer func() {
			if err := os.RemoveAll("config"); err != nil {
				panic(err)
			}
		}()

		email := "test@mail.com"
		password := "test123"

		form := url.Values{}
		form.Add("email", email)
		form.Add("password", password)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodPost, "/user/login", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		c.Request = req
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

		suite.authRepository.EXPECT().GetUserByEmail(c, email).Return(returnedUser, nil)

		suite.authRepository.EXPECT().UpdateOTP(c, returnedUser, gomock.Any()).Return(nil)

		suite.authHandler.Login(c)

		suite.Equal(http.StatusOK, w.Code)
	})
}
