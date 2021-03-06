VERSION := v0.1.0
PACKAGE_NAME := adrene

.PHONY: build
build:
	go build \
		-o ./bin/$(PACKAGE_NAME) \
		-ldflags "-X main.version=$(VERSION) -X main.name=$(PACKAGE_NAME)" .

.PHONY: test
test:
	go test ./...

.PHONY: install
install: test build
	chmod 755 ./bin/$(PACKAGE_NAME) && mv ./bin/$(PACKAGE_NAME) /usr/local/bin/

.PHONY: tag
tag:
	git tag $(VERSION)
	git push origin $(VERSION)
