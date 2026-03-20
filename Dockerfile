FROM golang:1.26

WORKDIR /app

COPY . /app/

RUN go build -o /app/kb2 /app/cmd/kb2

ENTRYPOINT [ "/app/kb2" ]
