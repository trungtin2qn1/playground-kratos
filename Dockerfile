FROM golang:1.20 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

WORKDIR /app

COPY --from=builder /src/bin .
RUN mkdir -p ./configs
COPY --from=builder /src/configs ./configs

EXPOSE 8000
EXPOSE 9000
# VOLUME /data/conf

CMD ["./playground-kratos", "-conf", "./configs"]