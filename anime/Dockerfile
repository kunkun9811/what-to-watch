FROM golang:latest as builder

WORKDIR /app
RUN go mod init anime
COPY *.go ./
RUN go mod tidy
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o anime

FROM scratch
EXPOSE 8080

COPY --from=builder /app /app
WORKDIR /app

ENTRYPOINT ["./anime"]