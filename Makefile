test:
	go test -v --coverprofile tests/coverage.out ./... 
	go tool cover -html=tests/coverage.out