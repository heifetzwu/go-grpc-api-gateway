# FROM golang:1.15.12-alpine3.13 as builder
FROM golang:1.18.5-alpine3.16 as builder

RUN apk update && apk upgrade && apk --update add git make

WORKDIR /app

COPY . .

RUN make engine


FROM alpine:latest
RUN apk update --no-cache && apk upgrade && apk --update --no-cache add tzdata && mkdir /app
WORKDIR /app
EXPOSE 3000

COPY --from=builder /app/go-grpc-api-gateway /app 
COPY --from=builder /app/pkg/config /app/pkg/config
COPY --from=builder /app/pkg/front /app/pkg/front

CMD /app/go-grpc-api-gateway 
# CMD echo "run" && sleep 60