FROM golang:1.19.2 as builder
WORKDIR /app
RUN go mod init anime
COPY *.go ./
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /anime

FROM gcr.io/distroless/base-debian11
WORKDIR /
COPY --from=builder /anime /anime
ENV PORT 8080
USER nonroot:nonroot

EXPOSE 8080

CMD ["/anime"]