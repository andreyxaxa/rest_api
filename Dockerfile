FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
COPY configs/apiserver.toml ./configs/apiserver.toml
RUN go build -o ./bin/app cmd/apiserver/main.go

FROM alpine AS runner

COPY --from=builder /app/bin/app /
COPY --from=builder /app/configs/apiserver.toml /configs/apiserver.toml
CMD [ "/app" ]