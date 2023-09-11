package services

import (
	"lido-core/v1/app/tasks"
	"lido-core/v1/pkg/workers"
)

func HealthCheck() error {
	return workers.Delay("core_api_queue", "Worker.HealthCheck", tasks.HealthCheck, int64(1))
}
