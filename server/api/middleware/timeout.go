package middleware

/*
	this code copied from:
	https://github.com/JacobSNGoodwin/memrizr/blob/master/account/handler/middleware/timeout.go
*/

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func Timeout(timout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		tw := &timeoutWriter{ResponseWriter: c.Writer, h: make(http.Header)}
		c.Writer = tw

		ctx, cancel := context.WithTimeout(c.Request.Context(), timout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		finished := make(chan struct{})
		panicChan := make(chan interface{}, 1)

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()

			c.Next()
			finished <- struct{}{}
		}()

		select {
		case <-panicChan:
			tw.ResponseWriter.WriteHeader(http.StatusInternalServerError)
			res, _ := json.Marshal(gin.H{
				"error": "cannot recover from timout's panic",
			})
			tw.ResponseWriter.Write(res)

		case <-finished:
			tw.mu.Lock()
			defer tw.mu.Unlock()

			dst := tw.ResponseWriter.Header()
			for k, vv := range tw.Header() {
				dst[k] = vv
			}

			tw.ResponseWriter.WriteHeader(tw.code)
			tw.ResponseWriter.Write(tw.wbuf.Bytes())

		case <-ctx.Done():
			tw.mu.Lock()
			defer tw.mu.Unlock()

			tw.ResponseWriter.Header().Set("Content-Type", "application/json")
			tw.ResponseWriter.WriteHeader(http.StatusRequestTimeout)
			res, _ := json.Marshal(gin.H{
				"err": "request got timedout",
			})
			tw.ResponseWriter.Write(res)
			c.Abort()
			tw.SetTimedOut()
		}
	}
}

type timeoutWriter struct {
	gin.ResponseWriter
	h    http.Header
	wbuf bytes.Buffer

	mu          sync.Mutex
	timedOut    bool
	wroteHeader bool
	code        int
}

func (tw *timeoutWriter) Write(b []byte) (int, error) {
	tw.mu.Lock()
	defer tw.mu.Unlock()

	if tw.timedOut {
		return 0, nil
	}

	return tw.wbuf.Write(b)
}

func (tw *timeoutWriter) WriteHeader(code int) {
	checkWriteHeaderCode(code)

	tw.mu.Lock()
	defer tw.mu.Unlock()

	if tw.timedOut || tw.wroteHeader {
		return
	}
	tw.writeHeader(code)
}

func (tw *timeoutWriter) writeHeader(code int) {
	tw.wroteHeader = true
	tw.code = code
}

func (tw *timeoutWriter) Header() http.Header {
	return tw.h
}

func (tw *timeoutWriter) SetTimedOut() {
	tw.timedOut = true
}

func checkWriteHeaderCode(code int) {
	if code < 100 || code > 999 {
		panic(fmt.Sprintf("invalide header code: %v", code))
	}
}
