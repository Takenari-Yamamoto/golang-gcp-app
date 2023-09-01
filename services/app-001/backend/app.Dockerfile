FROM golang:1.20 as builder

# ビルド用イメージ
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o sample-app-001-server cmd/server/main.go

# リリース用イメージ
FROM alpine
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
COPY --from=builder /app/sample-app-001-server /sample-app-001-server
CMD /sample-app-001-server