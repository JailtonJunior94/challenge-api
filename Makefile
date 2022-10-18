test:
	go test -v --coverprofile tests/coverage.out ./... 
	go tool cover -html=tests/coverage.out

start:
	docker-compose -f deployment/docker-compose.yml up -d --build

stop: 
	docker-compose -f deployment/docker-compose.yml down