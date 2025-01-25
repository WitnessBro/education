FROM golang:1.19.0

WORKDIR /app

COPY . . 

RUN go get -d -v ./...

RUN go build -o github.com/WitnessBro/education .

EXPOSE 8000

CMD ["./github.com/WitnessBro/education"]
