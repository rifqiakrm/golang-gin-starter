package service

import (
	"context"
	"errors"
	"log"
	"strings"

	"cloud.google.com/go/pubsub"
	"github.com/mailgun/mailgun-go/v4"

	"gin-starter/config"
	"gin-starter/entity"
)

// EmailSenderUsecase is use case for creating new ptk
type EmailSenderUsecase interface {
	// SendWithAPI send email using mailgun api
	SendWithAPI(ctx context.Context, mID, from, to, subject, message, category, creator string, pubsubMessage *pubsub.Message) error
}

// EmailSentRepository is use case for creating new ptk
type EmailSentRepository interface {
	// Insert insert email sent log to database
	Insert(ctx context.Context, ent *entity.EmailSent) error
}

// EmailSender is use case for creating new ptk
type EmailSender struct {
	emailSentRepo EmailSentRepository
	mailgunConfig config.MailGun
}

// NewEmailSender is constructor for EmailSender
func NewEmailSender(repository EmailSentRepository, mailgunConfig config.MailGun) *EmailSender {
	return &EmailSender{
		emailSentRepo: repository,
		mailgunConfig: mailgunConfig,
	}
}

// SendWithAPI send email using mailgun api and save to database
func (s *EmailSender) SendWithAPI(ctx context.Context, mID, from, to, subject, message, category, creator string, pubsubMessage *pubsub.Message) error {
	var status string
	if len(strings.TrimSpace(to)) == 0 {
		status = entity.EmailSentStatusNoRecipient
	} else {
		status = entity.EmailSentStatusOutgoing
	}

	// save sent message to repository
	emailSent := entity.NewEmailSent(mID, from, to, subject, message,
		status, category, creator)

	err := s.emailSentRepo.Insert(ctx, emailSent)

	if err != nil {
		pubsubMessage.Nack()
		log.Println(err)
	}

	if status == entity.EmailSentStatusNoRecipient {
		pubsubMessage.Ack()
		return errors.New("no recipient")
	}

	mg := mailgun.NewMailgun(s.mailgunConfig.Domain, s.mailgunConfig.APIKey)

	m := mg.NewMessage(from, subject, "", to)
	m.SetHtml(message)

	_, _, err = mg.Send(ctx, m)

	if err != nil {
		pubsubMessage.Nack()
		log.Println(err)
	}

	pubsubMessage.Ack()
	return nil
}
