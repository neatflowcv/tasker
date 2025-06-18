# Build stage
FROM golang:1.24.3-alpine AS builder

WORKDIR /app

# 의존성 복사 및 다운로드
COPY go.mod go.sum ./
RUN go mod download

# 소스코드 복사
COPY . .

# 바이너리 빌드
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tasker ./cmd/tasker

# Production stage
FROM alpine:latest

# 필요한 패키지 설치
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# 빌드된 바이너리 복사
COPY --from=builder /app/tasker .

# 포트 노출
EXPOSE 8080

# 애플리케이션 실행
CMD ["./tasker"] 