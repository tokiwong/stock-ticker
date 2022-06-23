REGISTRY ?= docker.io/tokiwong
IMAGE = $(REGISTRY)/stock-ticker

ALPHA_KEY = API_KEY=C227WD9W3LUVKVV9

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

##@ Build

.PHONY: build
build: fmt vet ## Build binary.
	go build -o bin/stonk main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: run
run: fmt vet build ## Run the app locally
	export $(ALPHA_KEY)
	bin/stonk

.PHONY: container
container:
	docker build --pull -t $(IMAGE):latest .

.PHONY: container-run
container-run:
	docker run --env $(ALPHA_KEY) --expose 8080 $(IMAGE):latest .

.PHONY: container-push
container-push:
	docker push $(IMAGE):latest
