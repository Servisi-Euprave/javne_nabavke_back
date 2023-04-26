FROM golang:alpine as build_container
WORKDIR /app
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main .


FROM alpine
WORKDIR /root/
COPY --from=build_container /app/main .
COPY ../public.pem . 

EXPOSE 8080
ENTRYPOINT ["./main"]
