package repository_test

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"gin-starter/modules/auth/v1/repository"
	"gin-starter/test/helpers"
	"gin-starter/utils"
)

var (
	email = "test@mail.com"
)

type AuthServiceTestSuite struct {
	suite.Suite
	subReporter *helpers.SubReporter
	db          *gorm.DB
	mock        sqlmock.Sqlmock
	repo        *repository.AuthRepository
}

func TestAuth(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuite))
}

func (s *AuthServiceTestSuite) BeforeTest(string, string) {
	s.subReporter = helpers.NewSubReporter(s.T())

	db, mock, err := sqlmock.New()
	if err != nil {
		s.FailNow("error opening a stub db connection: ", err)
	}

	s.db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		s.FailNow("error initializing gorm connection: ", err)
	}

	s.mock = mock
	s.repo = repository.NewAuthRepository(s.db)
}

func (s *AuthServiceTestSuite) AfterTest(string, string) {
	if err := s.mock.ExpectationsWereMet(); err != nil {
		s.FailNow("there were unfulfilled expectations: ", err)
	}
}

func (s *AuthServiceTestSuite) TestGetUserByEmail() {
	s.Run("fail to find record due to failure when executing select statement", func() {
		defer s.subReporter.Add(s.T())()

		s.mock.
			ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL`)).
			WithArgs(email).
			WillReturnError(errors.New("error"))

		res, err := s.repo.GetUserByEmail(context.Background(), email)

		s.Nil(res)
		s.NotNil(err)
	})

	s.Run("fail to find record due to non-existent record", func() {
		defer s.subReporter.Add(s.T())()

		s.mock.
			ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL`)).
			WithArgs(email).
			WillReturnError(gorm.ErrRecordNotFound)

		res, err := s.repo.GetUserByEmail(context.Background(), email)

		s.Nil(res)
		s.Nil(err)
	})

	s.Run("successfully find a record", func() {
		defer s.subReporter.Add(s.T())()
		dob, _ := utils.DateStringToTime("2006-01-02 00:00:00")
		s.mock.
			ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL`)).
			WithArgs(email).
			WillReturnRows(
				sqlmock.
					NewRows([]string{
						"id",
						"name",
						"email",
						"password",
						"phone_number",
						"photo",
						"dob",
						"otp",
						"status",
						"forgot_password_token",
					}).
					AddRow(
						uuid.New(),
						"Test",
						"test@gmail.com",
						"ThePassword",
						"0895346419497",
						"img_path",
						dob,
						"",
						true,
						"",
					),
			)

		res, err := s.repo.GetUserByEmail(context.Background(), email)

		s.NotNil(res)
		s.Nil(err)
	})
}
