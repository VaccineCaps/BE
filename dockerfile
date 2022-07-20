FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go build -ldflags "-w -s" -o vaccine .

EXPOSE 8080

ENTRYPOINT [ "./vaccine" ]