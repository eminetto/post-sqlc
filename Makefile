build: mod
	go build -o bin/post-sqlc cmd/api/main.go

mod:
	go mod download

test: clean generate-mocks sqlc-generate
	go test -v -coverprofile=cp.out ./...

sqlc-generate:
	@sqlc generate

generate-mocks:
	@mockery --output person/mocks --dir person/ --all

clean:
	@rm -rf person/mocks/*
