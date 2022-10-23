FROM golang:1.19-alpine3.16

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /demo-helm-app

CMD ["/demo-helm-app", "-f", "/app/config.yaml"]
