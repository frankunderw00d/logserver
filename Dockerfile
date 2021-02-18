FROM golang:1.14-alpine as builder
WORKDIR /usr/src/logserver
COPY ./logserver ./
RUN apk add --no-cache tzdata upx
RUN upx --best logserver -o _upx_logserver && \
mv -f _upx_logserver logserver

FROM scratch
WORKDIR /opt/logserver
COPY --from=builder /usr/src/logserver/logserver ./
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/
ENV TZ=Asia/Shanghai
CMD ["./logserver"]