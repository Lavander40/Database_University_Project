FROM golang:1.21.2

WORKDIR /postgre

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./

RUN go build -o ./bin/postgre_api.exe -v ./main.go

EXPOSE 4042

CMD [ "./bin/postgre_api.exe" ]