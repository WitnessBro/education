FROM golang:latest

WORKDIR /app

COPY . . 

RUN go get -d -v ./...

RUN go build -o github.com/WitnessBro/education ./cmd/task_manager

EXPOSE 8000

CMD ["./github.com/WitnessBro/education"]
