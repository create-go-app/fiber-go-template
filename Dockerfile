FROM golang:alpine AS builder

LABEL maintainer="Vic Shóstak <truewebartisans@gmail.com>"

WORKDIR /backend
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o apiserver cmd/apiserver/*.go

FROM scratch

COPY --from=builder ["/backend/apiserver", "/backend/configs/apiserver.yml", "/backend/"]
ENTRYPOINT ["/backend/apiserver", "-config-path", "/backend/apiserver.yml"]
