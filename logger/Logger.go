package logger

import (
	"net"
	"net/http"
	"strings"
	"time"
)

// LogRecord includes access log Content
type LogRecord struct {
	Time     time.Time
	Host     string
	URI      string
	Method   string
	Status   int
	Protocol string
	Size     int64
	From     string
}

// LoggingWriter includes http.ResponseWriter and LogRecord
type LoggingWriter struct {
	http.ResponseWriter
	logRecord LogRecord
}

// Logger incldes Log function
type Logger interface {
	Log(record LogRecord)
}

// LoggingHandler includes http.Handler and Logger
type LoggingHandler struct {
	handler http.Handler
	logger  Logger
}

// Write is used instead of http.ResponseWriter.Write
func (lw *LoggingWriter) Write(p []byte) (int, error) {
	written, err := lw.ResponseWriter.Write(p)
	lw.logRecord.Size += int64(written)
	return written, err
}

// WriteHeader is used instead of http.ResponseWriter.WriteHeader
func (lw *LoggingWriter) WriteHeader(status int) {
	lw.logRecord.Status = status
	lw.ResponseWriter.WriteHeader(status)
}

// GetIP gets X-Real-Ip, X-Forwarded-For or RemoteAddr from request
func (lh *LoggingHandler) GetIP(r *http.Request) string {
	xRealIP := r.Header.Get("X-Real-Ip")
	if xRealIP != "" {
		return xRealIP
	}

	xForwardedFor := r.Header.Get("X-Forwarded-For")
	for _, address := range strings.Split(xForwardedFor, ",") {
		address = strings.TrimSpace(address)
		if address != "" {
			return address
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

// ServeHTTP is used instead of http.ServeHTTP
func (lh *LoggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ip := lh.GetIP(r)
	startTime := time.Now()
	writer := &LoggingWriter{
		ResponseWriter: w,
		logRecord: LogRecord{
			Time:     startTime,
			Host:     r.Host,
			URI:      r.RequestURI,
			Method:   r.Method,
			Status:   0,
			Protocol: r.Proto,
			Size:     0,
			From:     ip,
		},
	}
	lh.handler.ServeHTTP(writer, r)
	lh.logger.Log(writer.logRecord)
}

// NewLoggingHandler returns
func NewLoggingHandler(handler http.Handler, logger Logger) http.Handler {
	return &LoggingHandler{
		handler: handler,
		logger:  logger,
	}
}
