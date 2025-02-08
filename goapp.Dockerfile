FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest
RUN go install github.com/joho/godotenv/cmd/godotenv@latest

COPY go.mod go.sum ./

RUN go mod download && go mod verify

CMD [ "air", "-c", ".air.linux.conf" ]