.PHONY: run
run:
	go run main.go

.PHONY: reset
reset:
	rm -rf ./dst/
