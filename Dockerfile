FROM golang:latest

WORKDIR /app
COPY . .
RUN go build -o myapp
EXPOSE 1234
CMD ["./myapp"]