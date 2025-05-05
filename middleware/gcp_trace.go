package middleware

import (
	"fmt"
	"strings"

	"cloud.google.com/go/compute/metadata"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func WithCloudTraceContext(c *gin.Context) {
	pid, err := metadata.ProjectIDWithContext(c.Request.Context())
	if err != nil || pid == "" {
		c.Next()

		return
	}

	trace := ""
	traceHeader := c.Request.Header.Get("X-Cloud-Trace-Context")
	traceParts := strings.Split(traceHeader, "/")

	if len(traceParts) > 0 && len(traceParts[0]) > 0 {
		trace = fmt.Sprintf("projects/%s/traces/%s", pid, traceParts[0])
	}

	ctx := context.WithValue(c.Request.Context(), "trace", trace)
	c.Request = c.Request.WithContext(ctx)

	c.Next()
}
