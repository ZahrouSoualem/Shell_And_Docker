FROM golang:1.19-alpine3.16

WORKDIR /app

COPY . /app

RUN chmod +x ./script.sh

RUN go build -o main main.go

ARG D_P=Zahrou

ENV NAME $D_P

EXPOSE 8080

CMD [ "./script.sh" ]