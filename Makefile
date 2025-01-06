run:
	@go run ./cmd/app/main.go

bench:
	@go test ./cmd/profiler/... -bench=. -benchmem

test:
	@go test ./cmd/profiler/... -v

profile:
	@go tool pprof -http 127.0.0.1:8080 cpu_profile.prof

momo:
	@rm ./resources/momo.gif
	@ffmpeg -i resources/momo.mov -vf "fps=30,scale=175:-1" -c:v gif resources/momo.gif
	@go run ./cmd/app/main.go

# @ffmpeg -i resources/full.mov -vf "fps=30,scale=175:-1" -c:v gif resources/full.gif
dandadan:
	@go run ./cmd/app/main.go
