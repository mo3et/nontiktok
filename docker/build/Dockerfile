FROM golang:1.20.7 AS builder

# download modules
WORKDIR /nontiktok
COPY ./go.mod .
RUN go mod download

# build
FROM builder AS runner
ARG SERVICE_NAME

WORKDIR /nontiktok
COPY . .

RUN go build -o bin/${SERVICE_NAME} /nontiktok/app/${SERVICE_NAME}/service/cmd

# run
FROM debian:stable-slim
ARG SERVICE_NAME

ENV SERVICE_NAME ${SERVICE_NAME}

WORKDIR /app
COPY --from=runner /nontiktok/bin/${SERVICE_NAME} bin/
COPY --from=runner /nontiktok/configs/service/${SERVICE_NAME} configs/

CMD bin/${SERVICE_NAME} -conf configs