#!/bin/bash
set -e

# 환경 변수 설정
GO_VERSION=${GO_VERSION:-1.24}
PLATFORMS=${PLATFORMS:-"linux/amd64,linux/arm64"}

echo "=== Go Linux 빌드 시작 ==="

# 디렉토리 생성
mkdir -p dist

# 빌드할 플랫폼 목록
IFS=',' read -ra PLATFORM_ARRAY <<< "$PLATFORMS"

for platform in "${PLATFORM_ARRAY[@]}"; do
    platform=$(echo "$platform" | xargs)  # trim whitespace
    
    # 플랫폼에서 OS와 아키텍처 분리
    IFS='/' read -ra PLATFORM_PARTS <<< "$platform"
    os="${PLATFORM_PARTS[0]}"
    arch="${PLATFORM_PARTS[1]}"
    
    echo "빌드 중: $os/$arch"
    
    # 아키텍처별 출력 파일명
    output_name="astrago-precheck-${os}-${arch}"
    
    # Docker를 사용하여 크로스 컴파일
    # CGO 비활성화로 정적 링크하여 GLIBC 의존성 제거 (Ubuntu 20.04 호환)
    docker run --rm \
        --platform "$platform" \
        -v "$(pwd):/workspace" \
        -w /workspace \
        golang:${GO_VERSION} \
        sh -c "
            # Go 모듈 다운로드
            go mod download
            
            # CGO 비활성화 및 정적 링크로 크로스 컴파일
            # CGO_ENABLED=0: C 라이브러리 의존성 제거 (GLIBC 버전 문제 해결)
            # -ldflags '-s -w': 바이너리 크기 최적화
            # -tags netgo: 순수 Go 네트워크 스택 사용
            CGO_ENABLED=0 GOOS=$os GOARCH=$arch go build -tags netgo -o dist/$output_name -ldflags '-s -w' .
        "
    
    echo "완료: dist/$output_name"
done

echo "=== 빌드 완료 ==="
ls -lh dist/
