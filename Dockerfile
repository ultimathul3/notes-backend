FROM golang:1.20

WORKDIR /usr/src/notes-backend

COPY . .

RUN go build -o notes-backend cmd/app/main.go

EXPOSE 8080

CMD ["./notes-backend"]
