FROM golang:1.16 as builder

WORKDIR /shortURL
COPY . /shortURL

ENV CGO_ENABLED=0
ENV GO111MODULE=on
RUN go build -mod=readonly -v -o server main.go

FROM alpine:3

RUN apk add --no-cache ca-certificates
RUN mkdir -p /shortURL
WORKDIR /shortURL

COPY --from=builder /shortURL/server /shortURL/server
COPY --from=builder /shortURL/deployment /shortURL/deployment

# Create appuser.
ENV USER=appuser
ENV UID=51197

RUN adduser \
--disabled-password \
--gecos "application user" \
--no-create-home \
--uid "${UID}" \
"${USER}"

RUN chown appuser:appuser /shortURL
RUN chown appuser:appuser /shortURL/*
USER appuser:appuser

ENTRYPOINT ["/shortURL/server"]