FROM golang:1.13-alpine3.10 AS builder
RUN  apk --update --no-cache add ca-certificates 
ARG VERSION
ARG COMMIT
ENV VERSION ${VERSION}
ENV COMMIT ${COMMIT}
WORKDIR /app
COPY . .
RUN GOOS=linux go build -o carpark -ldflags "-X github.com/x-mod/build.version=${VERSION} -X github.com/x-mod/build.commit=${COMMIT}"

FROM alpine:3.10
RUN  apk --update --no-cache add tzdata ca-certificates \
    && cp /usr/share/zoneinfo/Asia/Singapore /etc/localtime
WORKDIR /app
COPY --from=builder /app/carpark /usr/local/bin
ENTRYPOINT [ "carpark" ]
