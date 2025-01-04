FROM golang:1.23

ARG APP_NAME=gocrud-server
ENV APP_NAME=${APP_NAME}

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# COPY . .
# RUN go build -v -o /usr/local/bin/${APP_NAME} ./cmd/${APP_NAME}

EXPOSE 8080
# CMD /usr/local/bin/${APP_NAME}
CMD ["air"]
