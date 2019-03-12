FROM base:latest AS builder
RUN go build -o /crontab_worker_release ./worker/main/worker.go
COPY ./main/worker.json /worker.json

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates bash
COPY --from=builder /crontab_worker_release /
COPY --from=builder /worker.json /
CMD [ "/crontab_worker_release", "-config", "/worker.json" ]
LABEL Name=docker_go_crontab_worker Version=0.0.1