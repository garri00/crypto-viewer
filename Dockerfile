FROM golang:1.19

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o /bin ./...

EXPOSE 8080

CMD ["bin"]