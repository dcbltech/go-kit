package email

import "context"

type Emailer interface {
	SendTemplateEmail(ctx context.Context, template int, name, email string, data map[string]any) error
}
