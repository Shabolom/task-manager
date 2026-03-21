run:  ## Запустить сервис
	bash -c 'set -a; . ./build/local/.env; set +a; go run cmd/job_tracker/main.go'
