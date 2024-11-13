FROM golang:1.22 as builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /weather-api cmd/weather-api/main.go

FROM scratch
COPY --from=builder /weather-api /weather-api
EXPOSE 8080
ENTRYPOINT ["/weather-api"]
