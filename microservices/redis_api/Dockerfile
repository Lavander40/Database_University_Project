FROM golang:1.21.2

WORKDIR /redis

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o ./bin/redis_api.exe -v ./main.go

EXPOSE 4041

CMD [ "./bin/redis_api.exe" ]