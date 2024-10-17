.PHONY: bench
bench:
	go test -bench . -benchmem -benchtime 5s
