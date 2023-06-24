package service

import "github.com/EnricRG/resteduler-backend/internal/domain"

type SchedulerService interface {
	GetUserSchedules(userId uint) ([]domain.Schedule, error)
}
