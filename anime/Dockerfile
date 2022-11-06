FROM golang:latest

WORKDIR /app
RUN go mod init anime
COPY *.go ./
RUN go mod tidy

RUN go build -o /anime

EXPOSE 8080

CMD [ "/anime" ]