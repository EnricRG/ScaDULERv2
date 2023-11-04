package api

import (
	"net/http"

	"github.com/EnricRG/openscheduler-backend/internal/logging"
	"github.com/EnricRG/openscheduler-backend/internal/service"
	"github.com/labstack/echo/v4"
	log "github.com/rs/zerolog"
)

// Api definition of the service
func (s *server) Api() []ApiEntry {
	return []ApiEntry{
		{"GET", "/api/v1/users/:uid/schedules", s.getUserSchedules},
	}
}

type server struct {
	port     uint16
	services service.Services
	logger   log.Logger
}

func NewServer(port uint16, services service.Services) *server {
	return &server{
		port,
		services,
		logging.GetLogger(),
	}
}

// logUrl logs request URI.
func (s *server) logUrl(ctx echo.Context) {
	s.logger.Info().Str("url", ctx.Request().URL.String())
}

// internalServerError returns generic internal server error into context.
func (s *server) internalServerError(ctx echo.Context) error {
	return ctx.String(http.StatusInternalServerError, "Internal Server Error")
}
