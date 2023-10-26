package graph

import (
	"fio_finder/internal/service"
	"fio_finder/pkg/logger"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Services service.Services
	Logger   logger.Logger
}
