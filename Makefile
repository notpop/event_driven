.PHONY: all build-api build-job run-api run-job start-redis stop-redis clean run stop redis-ping

# Default task
all: build-api build-job

# Build API server
build-api:
	go build -o bin/api ./api

# Build Job Processor
build-job:
	go build -o bin/job_processor ./job_processor

# Run API server
run-api:
	./bin/api

# Run Job Processor
run-job:
	./bin/job_processor

# Install Redis server
install-redis:
	brew install redis
	brew services start redis

# Start Redis server
start-redis:
	brew services start redis

# Stop Redis server
stop-redis:
	brew services stop redis

redis-ping:
	redis-cli ping

# Clean built binaries
clean:
	rm -f bin/api bin/job_processor

# Start all services
run: start-redis build-api build-job
	make -j 2 run-api run-job

# Stop all services
stop: stop-redis

analyze:
	go run tools/endpoint_analyzer.go
