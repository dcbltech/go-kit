package email

import "context"

type Emailer interface {
	SendTemplateEmail(ctx context.Context, template int64, name, email string, data map[string]any) error
}
