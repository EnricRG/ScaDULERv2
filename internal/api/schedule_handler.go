package api

import (
	"net/http"

	"github.com/EnricRG/openscheduler-backend/internal/api/mapper"
	"github.com/EnricRG/openscheduler-backend/internal/service"
	"github.com/labstack/echo/v4"
)

type ScheduleHandlers struct {
	server  *server
	service service.SchedulerService
}

func (handler *ScheduleHandlers) getUserSchedules(ctx echo.Context) (err error) {
	handler.server.logUrl(ctx)

	var userId uint
	if err = echo.PathParamsBinder(ctx).Uint("uid", &userId).BindError(); err != nil {
		s.logger.Error().Err(err).Msg("Could not bind user identifier from url")
		return ctx.String(http.StatusBadRequest, "Invalid user identifier")
	}

	userSchedules, err := s.services.Shcedule.GetUserSchedules(userId)
	if err != nil {
		s.logger.Error().Err(err).Msgf("Unexpected error while retrieving schedules for user '%d'", userId)
		return s.internalServerError(ctx)
	}

	viewSchedules := mapper.NewScheduleMapper().ToViews(userSchedules)
	return ctx.JSON(http.StatusOK, viewSchedules)
}
