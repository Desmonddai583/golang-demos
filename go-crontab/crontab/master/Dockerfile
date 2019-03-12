FROM base:latest AS builder
RUN go build -o /crontab_master_release ./master/main/master.go
COPY ./main/master.json /master.json
COPY ./main/webroot /webroot

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /crontab_master_release /
COPY --from=builder /master.json /
COPY --from=builder /webroot /webroot
CMD [ "/crontab_master_release", "-config", "/master.json" ]
LABEL Name=docker_go_crontab_master Version=0.0.1