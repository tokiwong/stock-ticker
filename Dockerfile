# Build the stock ticker binary
FROM golang:1.18 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.sum main.go ./

RUN go mod download

COPY pkg/ pkg/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o stonk main.go

# Use distroless as minimal base image to package the ticker binary
FROM gcr.io/distroless/static-debian11
WORKDIR /
COPY --from=builder /workspace/stonk .

ENTRYPOINT ["/stonk"]
