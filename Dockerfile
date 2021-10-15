FROM golang:alpine AS build
WORKDIR /app
ADD go.mod go.sum ./
RUN go mod download
ADD . .
ENV CGO_ENABLED=0 
ENV GOOS=linux
RUN go test -v ./internal/app
RUN go build -a -installsuffix cgo -o app ./cmd/app

FROM scratch
COPY --from=build /app /app

ENTRYPOINT ["/app"]