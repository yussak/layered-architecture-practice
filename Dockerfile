FROM golang:1.23-alpine

WORKDIR /app

# COPY . .
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# COPY *.go ./
COPY . .

RUN go build -o main .
# RUN go build -o ./main

EXPOSE 8080

CMD [ "./main" ]

