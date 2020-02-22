FROM golang:1.13
WORKDIR /src
COPY . .
RUN go build -o http-echo main.go

FROM ubuntu
COPY --from=0 /src/http-echo .
CMD ["./http-echo"]
