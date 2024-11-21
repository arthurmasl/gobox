run:
	@time go run ./cmd/app/main.go

bench:
	@go test ./cmd/app/... -bench=. -benchmem

test:
	@go test ./cmd/app/... -v
	@make bench

