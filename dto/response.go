package dto

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type responseType int

const (
	responseTypeRespond responseType = iota
	responseTypeRedirect
)

type Response struct {
	Type   responseType
	Status int
	Data   any
}

func (r *Response) Respond(c *gin.Context) {
	switch r.Type {
	case responseTypeRespond:
		r.respondFn(c)
	case responseTypeRedirect:
		r.redirectFn(c)
	}
}

func (r *Response) respondFn(c *gin.Context) {
	if r.Data == nil {
		c.Status(r.Status)
	} else {
		c.JSON(r.Status, r.Data)
	}
}

func (r *Response) redirectFn(c *gin.Context) {
	if url, ok := r.Data.(string); ok {
		log.Printf("Redirecting to: %s with status: %d", url, r.Status)
		c.Redirect(r.Status, url)
	} else {
		log.Printf("Invalid redirect data type: %T", r.Data)
		c.Status(http.StatusInternalServerError)
	}
}
