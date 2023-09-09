package middleware

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"math"
	"net/http"
	"net/http/httptest"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/relvox/iridescence_go/logging"
	"github.com/relvox/iridescence_go/utils"
)

type LoggingOptions uint16

const (
	Method      = 1 << 0 // 0. LogMethod
	RemoteAddr  = 1 << 1 // 1. LogRemoteAddr
	RequestID   = 1 << 2 // 2. LogRequestID
	UserAgent   = 1 << 3 // 3. LogUserAgent
	Query       = 1 << 4 // 4. Query
	ContentType = 1 << 5 // 5. LogContentType
	Request     = 1 << 6 // 6. Request

	Response            = 1 << 10          // 10. LogResponse
	ResponseContentType = 1<<11 + Response // 11. LogResponseContentType (+10)
	ResponseLength      = 1<<12 + Response // 12. LogResponseLength (+10)
	ResponseStatus      = 1<<13 + Response // 13. LogResponseStatus (+10)

	Duration = 1 << 15 // 15. LogDuration
)

var AllOptions LoggingOptions = math.MaxUint16

func (o LoggingOptions) LogMethod() bool              { return (o & Method) != 0 }
func (o LoggingOptions) LogRemoteAddr() bool          { return (o & RemoteAddr) != 0 }
func (o LoggingOptions) LogRequestID() bool           { return (o & RequestID) != 0 }
func (o LoggingOptions) LogUserAgent() bool           { return (o & UserAgent) != 0 }
func (o LoggingOptions) LogQuery() bool               { return (o & Query) != 0 }
func (o LoggingOptions) LogContentType() bool         { return (o & ContentType) != 0 }
func (o LoggingOptions) LogRequest() bool             { return (o & Request) != 0 }
func (o LoggingOptions) LogResponse() bool            { return (o & Response) != 0 }
func (o LoggingOptions) LogResponseContentType() bool { return (o & ResponseContentType) != 0 }
func (o LoggingOptions) LogResponseLength() bool      { return (o & ResponseLength) != 0 }
func (o LoggingOptions) LogResponseStatus() bool      { return (o & ResponseStatus) != 0 }
func (o LoggingOptions) LogDuration() bool            { return (o & Duration) != 0 }

func LoggingMiddleware(log *slog.Logger, opts ...LoggingOptions) mux.MiddlewareFunc {
	config := utils.Sum(opts...)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			logFields := []any{slog.String("url", r.URL.String())}
			if config.LogMethod() {
				logFields = append(logFields, slog.String("method", r.Method))
			}
			if config.LogRemoteAddr() {
				logFields = append(logFields, slog.String("remote_addr", r.RemoteAddr))
			}
			if config.LogRequestID() {
				logFields = append(logFields, slog.String("request_id", r.Header.Get("X-Request-ID")))
			}
			if config.LogUserAgent() {
				logFields = append(logFields, slog.String("user_agent", r.UserAgent()))
			}
			if config.LogQuery() {
				queries := r.URL.Query()
				logFields = append(logFields, slog.String("url_queries", queries.Encode()))
			}
			if r.Method != "GET" && r.Method != "DELETE" {
				if config.LogContentType() {
					logFields = append(logFields, slog.String("content_type", r.Header.Get("Content-Type")))
				}
				if config.LogRequest() {
					bodyBytes, err := io.ReadAll(r.Body)
					if err != nil {
						log.Error("reading body", logging.Error(err))
					}
					logFields = append(logFields, slog.String("request_body", string(bodyBytes)))
					r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
				}
			}
			log.Info("Incoming Request", logFields...)

			ctx := context.WithValue(r.Context(), Key("log"), log)
			r = r.WithContext(ctx)

			rec := httptest.NewRecorder()
			if config.LogResponse() {
				next.ServeHTTP(rec, r)
				for k, v := range rec.Header() {
					w.Header()[k] = v
				}
				w.WriteHeader(rec.Code)
				_, err := w.Write(rec.Body.Bytes())
				if err != nil {
					log.Error("write body to actual response", logging.Error(err))
				}
			} else {
				next.ServeHTTP(w, r)
			}

			logFields = []any{slog.String("url", r.URL.String())}
			if config.LogMethod() {
				logFields = append(logFields, slog.String("method", r.Method))
			}
			if config.LogRemoteAddr() {
				logFields = append(logFields, slog.String("remote_addr", r.RemoteAddr))
			}
			if config.LogRequestID() {
				logFields = append(logFields, slog.String("request_id", r.Header.Get("X-Request-ID")))
			}

			if config.LogResponse() {
				if config.LogResponseContentType() {
					logFields = append(logFields, slog.String("response_content_type", rec.Header().Get("Content-Type")))
				}
				logFields = append(logFields, slog.String("response", rec.Body.String()))
				if config.LogResponseLength() {
					logFields = append(logFields, slog.String("response_length", strconv.Itoa(rec.Body.Len())))
				}
				if config.LogResponseStatus() {
					logFields = append(logFields, slog.String("response_status", strconv.Itoa(rec.Code)))
				}
			}
			if config.LogDuration() {
				logFields = append(logFields, slog.String("duration", time.Since(start).String()))
			}
			log.Info("Request Handled", logFields...)
		})
	}
}
