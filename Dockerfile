FROM golang:alpine AS builder

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container
COPY . .

# Set necessary environmet variables needed for our image and build the API server
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o apiserver ./cmd/apiserver/...

FROM scratch AS runner

# Copy binary from build to main folder
COPY --from=builder ["/build/apiserver", "/build/configs/apiserver.yml", "/dist/"]

# Export necessary port
EXPOSE 5000

# Command to run when starting the container
CMD ["/dist/apiserver"]