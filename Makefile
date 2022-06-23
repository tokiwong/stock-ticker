REGISTRY ?= docker.io/tokiwong
IMAGE = $(REGISTRY)/stock-ticker

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

.PHONY: run
run: fmt vet build ## Run the app locally
	export API_KEY=C227WD9W3LUVKVV9
	bin/stonk

.PHONY: container
container:
	docker build --pull -t $(IMAGE):latest .

.PHONY: container-push
container-push:
	docker push $(IMAGE):latest
