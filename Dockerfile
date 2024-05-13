FROM golang:1.22.1 AS final
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

WORKDIR /src/server

RUN CGO_ENABLED=1 GOOS=linux go build -o /app/server

EXPOSE 8080
ENTRYPOINT [ "/app/server" ]