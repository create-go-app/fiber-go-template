FROM golang:1.14-alpine AS builder

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container
COPY . .

# Set necessary environmet variables needed for our image and build the API server
ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-w -s" -o apiserver ./cmd/apiserver/...

FROM scratch

# Copy binary from build to main folder
COPY --from=builder ["/build/apiserver", "/build/configs/apiserver.yml", "/dist/"]

# Export necessary port
EXPOSE 5000

# Command to run when starting the container
ENTRYPOINT ["/dist/apiserver"]