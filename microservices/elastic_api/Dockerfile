FROM golang:1.21.2

WORKDIR /elastic

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o ./bin/elastic_api.exe -v ./main.go

EXPOSE 4043

CMD [ "./bin/elastic_api.exe" ]