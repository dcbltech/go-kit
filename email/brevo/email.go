package brevo

import (
	"context"

	"github.com/dcbltech/go-kit/email"
	brevo "github.com/getbrevo/brevo-go/lib"
)

type Client struct {
	client *brevo.APIClient
}

func Must(apiKey string) email.Emailer {
	c := &Client{}

	cfg := brevo.NewConfiguration()
	cfg.AddDefaultHeader("api-key", apiKey)

	c.client = brevo.NewAPIClient(cfg)

	return c
}

func (c *Client) SendTemplateEmail(ctx context.Context, template int64, name string, email string, data map[string]any) error {
	_, _, err := c.client.TransactionalEmailsApi.SendTransacEmail(
		ctx,
		brevo.SendSmtpEmail{
			To:         []brevo.SendSmtpEmailTo{{Name: name, Email: email}},
			TemplateId: template,
			Params:     data,
		},
	)

	return err
}
