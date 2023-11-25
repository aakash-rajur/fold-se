FROM bitnami/golang:1.21 as builder

RUN mkdir -p /home/nobody && chown -R nobody:nogroup /home/nobody

WORKDIR /home/nobody

USER nobody

ENV GOPROXY=direct HOME=/home/nobody

COPY --chown=nobody:nogroup go.mod go.sum ./

RUN go mod tidy

COPY --chown=nobody:nogroup . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/app/app.go

FROM alpine:3

RUN mkdir -p /home/nobody && chown -R nobody:nogroup /home/nobody

WORKDIR /home/nobody

USER nobody

ENV HOME=/home/nobody

USER nobody

COPY --from=builder /home/nobody/app .

ENTRYPOINT ["/home/nobody/app"]
