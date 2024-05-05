package handlers

import (
	"io"
	"strings"

	"github.com/gin-gonic/gin"
)

// request
// =======================================================
type Request interface {
	Bind(c *gin.Context) error
}

type RawRequest struct {
	Headers map[string][]string `json:"-"`
	Body    string              `json:"-"`
}

func (r *RawRequest) Bind(c *gin.Context) error {
	// logger := log.GetLogger("Bind")
	// copy headers
	r.Headers = make(map[string][]string)
	for h, v := range c.Request.Header {
		r.Headers[strings.ToLower(h)] = v
	}

	// copy body
	if body, exist := c.Get("VerifiedBody"); exist {
		// logger.Debug("found verified body")
		r.Body = strings.Clone(body.(string))
	} else if c.Request.Body != nil {
		// logger.Debug("found body, but not verified body")
		if bodyBytes, err := io.ReadAll(c.Request.Body); err == nil {
			r.Body = strings.Clone(string(bodyBytes))
		} else {
			return err
		}
	} else {
		// logger.Debug("body not found")
		r.Body = ""
	}

	return nil
}

// response
// =======================================================
type Response interface {
}

// handler
// =======================================================
type Handler[TReq Request, TRes Response] interface {
	Handle(rq TReq) (TRes, int, error)
	CreateRequestObject() TReq
}
