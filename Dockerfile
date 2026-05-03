# Build stage
FROM golang:alpine AS builder

RUN apk add --no-cache ca-certificates git

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0

RUN go build -ldflags="-s -w" -o /bin/ghpgps .

# Runtime stage
FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /bin/ghpgps /bin/ghpgps
COPY templates/ /templates/
COPY static/ /static/

EXPOSE 5001

USER nobody

CMD ["/bin/ghpgps"]
