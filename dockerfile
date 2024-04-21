FROM golang:1.22.2-bullseye as dev
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]