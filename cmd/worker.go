package cmd

import (
	"lido-core/v1/app/tasks"
	"lido-core/v1/pkg/configs"
	"lido-core/v1/pkg/workers"
)

func Setting(broker, resultBackend string) *configs.WorkerConfig {
	return &configs.WorkerConfig{
		Broker:        broker,
		ResultBackend: resultBackend,
		Workers: map[string]configs.Task{
			"core_api_queue": map[string]interface{}{
				"Worker.HealthCheck": tasks.HealthCheck,
				"Worker.QuizAnswer":  tasks.QuizAnswer,
			},
		},
	}
}

func WorkerExecute(queueName, consume string, concurrency int) error {
	wcf := Setting(configs.BrokerUrl, configs.ResultBackend)
	cnf := configs.NewWorker(queueName, wcf)
	return workers.Execute(cnf, consume, concurrency)
}
