package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/pkg/errors"

	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/notification/v1/service"
)

// SendEmailPubsubHandler struct
type SendEmailPubsubHandler struct {
	mailGun        config.MailGun
	emailSenderSvc service.EmailSenderUsecase
}

const (
	// SendEmailSubName is a subscriber name for SendEmailSub
	SendEmailSubName = "send-email-sub"
)

// NewSendEmailPubsubHandler create ptk pubsub handler
func NewSendEmailPubsubHandler(emailSenderSvc service.EmailSenderUsecase, mailGun config.MailGun) *SendEmailPubsubHandler {
	return &SendEmailPubsubHandler{
		mailGun:        mailGun,
		emailSenderSvc: emailSenderSvc,
	}
}

// SubscriptionName is a function for getting subscription name
func (pubsub *SendEmailPubsubHandler) SubscriptionName() string {
	return SendEmailSubName
}

// ProcessMessage is a function for processing message from pubsub
func (pubsub *SendEmailPubsubHandler) ProcessMessage(ctx context.Context, m *pubsub.Message) {
	// log message from pubsub
	log.Println(fmt.Sprintf("Received message: %s", m.Data))

	var payload entity.EmailPayload

	// parsing json payload
	if err := json.Unmarshal(m.Data, &payload); err != nil {
		log.Println(errors.Wrap(err, fmt.Sprintf("[SendEmailPubsubHandler-ProcessMessage] error unmarshal: %s", m.Attributes)))
		m.Ack()
		return
	}

	// send email
	err := pubsub.emailSenderSvc.SendWithAPI(ctx, m.ID, pubsub.mailGun.From, payload.To, payload.Subject, payload.Content, payload.Category, pubsub.SubscriptionName(), m)
	if err != nil {
		log.Println(errors.Wrap(err, fmt.Sprintf("[SendEmailPubsubHandler-ProcessMessage] error send email svc: %s", m.Attributes)))
		// m.Ack()
		return
	}
	m.Ack()
}
