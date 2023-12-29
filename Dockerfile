FROM golang AS builder
ENV GO111MODULE=on
WORKDIR /go/src/invoice-app/
COPY go.mod .
COPY go.sum .
RUN go mod tidy

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /go/src/invoice-app/app .


FROM scratch

COPY --from=builder /go/src/invoice-app/app /app
COPY --from=builder /go/src/invoice-app/.env /.env
COPY --from=builder /go/src/invoice-app/html /html
COPY --from=builder /go/src/invoice-app/static /static
COPY --from=builder /etc/ssl /etc/ssl

ENTRYPOINT ["/app"]