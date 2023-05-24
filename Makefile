build:
	go build -o bin/
run:
	go run .
tests:
	go test ./...
tests-bench:
	go test -v --bench . --benchmem ./...
clean:
	rm -rf bin/