FROM golang:1.22-alpine AS builder

ENV MINIO_HOST=""
ENV MINIO_PORT=""
ENV MINIO_USER=""
ENV MINIO_PASSWORD=""

WORKDIR /usr/local/src/proc

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/proc main.go

FROM alpine AS runner

COPY --from=builder /usr/local/bin/proc /usr/local/bin/proc

WORKDIR /usr/local/bin
EXPOSE 8000
CMD [ "proc" ]