#!/bin/bash
set -e

echo "=== Go Linux 네이티브 빌드 시작 ==="

# 현재 시스템 아키텍처 확인
ARCH=$(uname -m)
case "$ARCH" in
    x86_64)
        GOARCH="amd64"
        ;;
    aarch64|arm64)
        GOARCH="arm64"
        ;;
    *)
        echo "지원하지 않는 아키텍처: $ARCH"
        exit 1
        ;;
esac

GOOS="linux"
output_name="astrago-precheck-linux-${GOARCH}"

echo "빌드 중: $GOOS/$GOARCH"
echo "출력 파일: dist/$output_name"

# dist 디렉토리 생성
mkdir -p dist

# 빌드 실행
GOOS=$GOOS GOARCH=$GOARCH go build -o "dist/$output_name" -ldflags '-s -w' .

echo "=== 빌드 완료 ==="
ls -lh "dist/$output_name"
