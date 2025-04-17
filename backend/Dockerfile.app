FROM golang:1.23.4

WORKDIR /app

COPY . .

RUN go build -o arlyApi .

EXPOSE 5050

CMD ["./arlyApi"]
