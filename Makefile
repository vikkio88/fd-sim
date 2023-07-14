build:
	go build -o bin/
run:
	go run .
tests:
	go clean -testcache && go test ./...
tests-bench:
	go test -v --bench . --benchmem ./...
clean:
	rm -rf bin/ fdsim.db test.db db/test.db db_test/test.db
gen-db:
	go run runner/main.go && mv test.db fdsim.db
cp-testdb:
	cp db_test/test.db fdsim.db