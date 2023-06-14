FROM golang

COPY . /app

WORKDIR /app

RUN go get .

ENTRYPOINT go run main.go

EXPOSE 3000