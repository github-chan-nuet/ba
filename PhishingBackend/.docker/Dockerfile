# Golang images: https://hub.docker.com/_/golang
FROM golang:1.24.1 AS build
WORKDIR /app
# copy only dependency related files in a first step to enable incremental build / improve performance
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# CGO_ENABLED flag decides, if static or dynamic C libs should be used. 0 = static, 1 = dynamic. CGO_ENABLED=0
# As no dynamic C libs are available at runtime, static ones will be used in the resulting binary.
# -o = define file output path
RUN CGO_ENABLED=0 go build -o /bin/phishingBackend cmd/main.go

# Alpine images: https://hub.docker.com/_/alpine
FROM alpine:3.21.3
COPY --from=build /bin/phishingBackend /app/phishingBackend
# create non root user to run the app
RUN adduser -D nonrootuser && chown -R nonrootuser:nonrootuser /app
USER nonrootuser
CMD ["/app/phishingBackend"]