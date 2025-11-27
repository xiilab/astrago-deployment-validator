# AstraGo Deployment Validator

Kubernetes 클러스터 배포 전 Linux 노드 사전 점검 도구입니다. SSH를 통한 원격 체크와 로컬 체크를 지원하며, JSON, Markdown, HTML 형식의 리포트를 자동 생성합니다.

## 주요 기능
- ✅ **다중 노드 지원**: 여러 노드를 한 번에 점검
- ✅ **SSH/로컬 실행**: 원격 SSH 접속 또는 로컬 실행 모드
- ✅ **다양한 체크 항목**: 네트워크, 시스템 설정, 리소스 등 종합 점검
- ✅ **다중 리포트 형식**: JSON, Markdown, HTML 자동 생성
- ✅ **Kubernetes 특화**: kube-vip NIC 검증 등 Kubernetes 배포에 필요한 체크 포함
- ✅ **단일 바이너리**: Go로 작성되어 단일 실행 파일로 배포 가능
- ✅ **크로스 컴파일**: 한 플랫폼에서 여러 아키텍처 빌드 지원

## 체크 항목
각 노드에 대해 다음 항목들을 점검합니다:
1. **Ping**: 네트워크 연결 확인
2. **SSH**: SSH 접속 가능 여부
3. **Firewall**: 방화벽 상태 (firewalld/ufw/iptables)
4. **Ports**: 지정된 포트의 접속/리스닝 상태 (기본: 22, 6443, 2379, 10250)
5. **Disk**: 디스크 사용량 및 파일시스템 정보
6. **DNS**: `/etc/resolv.conf` 설정 확인
7. **Sudo**: sudo 권한 확인
8. **NTP**: 시간 동기화 상태 (timedatectl/chronyc/ntpq)
9. **CPU/Memory/GPU**: 시스템 리소스 정보
10. **Internet**: 인터넷 연결 확인 (8.8.8.8, google.com)
11. **Swap**: Swap 활성화 여부
12. **Kube VIP NIC**: Kubernetes VIP와 노드 IP의 NIC 일치 여부 (master 노드 전용)

## 프로젝트 구조
```text
.
├── main.go                    # 메인 진입점
├── go.mod                     # Go 모듈 정의
├── internal/                  # 내부 패키지
│   ├── config/                # 설정 및 CLI 옵션
│   │   └── config.go
│   ├── models/                # 데이터 모델
│   │   └── models.go
│   ├── utils/                 # 공통 유틸리티 함수
│   │   └── utils.go
│   ├── ssh/                   # SSH 클라이언트
│   │   └── ssh.go
│   ├── checker/               # 노드 체크 로직
│   │   └── checker.go
│   └── report/                # 리포트 생성
│       └── report.go
├── build-linux.sh             # Docker 기반 Linux 빌드 스크립트
└── build-linux-native.sh      # 네이티브 Linux 빌드 스크립트
```

## 설치

### 필수 요구사항
- **빌드 시**: Go 1.24 이상 (실행 시에는 Go 불필요)
- Linux 환경 (또는 macOS/Windows에서 원격 체크)

### Go 설치

Go가 설치되어 있지 않다면 [Go-INSTALL.md](Go-INSTALL.md) 파일을 참조하여 설치하세요.

간단한 설치 방법:
- **macOS**: `brew install go`
- **Linux**: `sudo apt-get install golang-go` 또는 [Go-INSTALL.md](Go-INSTALL.md) 참조
- **Windows**: [Go 설치파일 다운로드](https://go.dev/dl/) 에서 설치 파일 다운로드

### 빌드 방법

#### 1. 로컬 빌드 (현재 플랫폼용)
```bash
go build -o astrago-precheck .
```

#### 2. Docker를 사용한 크로스 컴파일 (Linux amd64/arm64)
```bash
# 모든 아키텍처 빌드 (기본값)
./build-linux.sh

# 특정 플랫폼만 빌드
PLATFORMS="linux/amd64" ./build-linux.sh

# Go 버전 지정
GO_VERSION=1.24 ./build-linux.sh
```

#### 3. Linux 서버에서 네이티브 빌드
```bash
# 서버에서 직접 실행
./build-linux-native.sh
```

빌드된 파일은 `dist/` 디렉토리에 생성됩니다:
- `dist/astrago-precheck-linux-amd64` (x86_64용)
- `dist/astrago-precheck-linux-arm64` (ARM64용)

## 사용법

### 기본 사용법

#### 1. 로컬 노드만 체크
```bash
./astrago-precheck --local-only
```

#### 2. 원격 노드 체크 (SSH 키 사용)
```bash
./astrago-precheck \
  --node 192.168.1.10 --role master \
  --node 192.168.1.11 --role worker \
  --node 192.168.1.12 --role worker \
  --user root \
  --key ~/.ssh/id_rsa \
  --kube-vip 192.168.1.100 \
  --ports "22,6443,2379,10250" \
  --output precheck_report
```

#### 3. 원격 노드 체크 (비밀번호 사용)
```bash
./astrago-precheck \
  --node 192.168.1.10 --role master \
  --user root \
  --password your_password \
  --kube-vip 192.168.1.100
```

#### 4. 커스텀 SSH 포트 사용
```bash
./astrago-precheck \
  --node 192.168.1.10 --role master \
  --user admin \
  --port 2222 \
  --key ~/.ssh/id_rsa
```

### 명령줄 옵션

| 옵션 | 설명 | 기본값 |
|------|------|--------|
| `--local-only` | 로컬 노드만 체크 | `false` |
| `--node` | 체크할 노드 IP 주소 (여러 개 지정 가능) | - |
| `--role` | 노드 역할 (master, worker 등) | `unknown` |
| `--user` | SSH 사용자명 | `root` |
| `--port` | SSH 포트 번호 | `22` |
| `--key` | SSH 키 파일 경로 | - |
| `--password` | SSH 비밀번호 | - |
| `--ports` | 체크할 포트 목록 (쉼표 구분) | `22,6443,2379,10250` |
| `--kube-vip` | Kubernetes VIP 주소 | - |
| `--output` | 출력 파일 경로 (확장자 제외) | `precheck_report` |

## 리포트 형식

도구는 세 가지 형식의 리포트를 자동 생성합니다:

1. **JSON** (`precheck_report.json`): 구조화된 데이터, 자동화 스크립트에서 활용
2. **Markdown** (`precheck_report.md`): 사람이 읽기 쉬운 형식
3. **HTML** (`precheck_report.html`): 브라우저에서 보기 좋은 형식

### 리포트 예시

```json
{
  "meta": {
    "generated_at": "2024-01-15 10:30:00",
    "kube_vip": "192.168.1.100",
    "ports_checked": [22, 6443, 2379, 10250]
  },
  "nodes": {
    "192.168.1.10": {
      "ip": "192.168.1.10",
      "role": "master",
      "ping": {"ok": true, "detail": "Ping 성공"},
      "ssh": {"ok": true, "detail": "SSH 접속 성공"},
      ...
    }
  }
}
```

## 빌드 스크립트 상세

### build-linux.sh (Docker 기반 크로스 컴파일)

Docker를 사용하여 여러 Linux 아키텍처용 바이너리를 빌드합니다.

**환경 변수:**
- `GO_VERSION`: Go 버전 (기본값: `1.24`)
- `PLATFORMS`: 빌드할 플랫폼 목록 (기본값: `linux/amd64,linux/arm64`)

**특징:**
- `CGO_ENABLED=0`로 정적 링크하여 GLIBC 의존성 제거
- Ubuntu 20.04 (GLIBC 2.31) 이상에서 실행 가능
- 순수 Go 네트워크 스택 사용 (`-tags netgo`)

**사용 예시:**
```bash
# 모든 아키텍처 빌드
./build-linux.sh

# amd64만 빌드
PLATFORMS="linux/amd64" ./build-linux.sh

# Go 1.24 사용
GO_VERSION=1.24 ./build-linux.sh
```

### build-linux-native.sh (네이티브 빌드)

Linux 서버에서 직접 빌드합니다. 현재 서버의 아키텍처에 맞는 바이너리를 생성합니다.

**특징:**
- GLIBC 호환성 문제 없음 (같은 환경에서 빌드)
- 빠른 빌드 속도
- 추가 도구 불필요

**사용 예시:**
```bash
# Linux 서버에서 실행
./build-linux-native.sh
```

## 호환성 정보
### Go 버전 요구사항
- **빌드 시**: Go 1.24 이상 필요 (`go.mod`에 명시)
  - `build-linux.sh`의 기본값은 Go 1.24입니다
  - 다른 버전 사용 시: `GO_VERSION=1.22 ./build-linux.sh` 형태로 지정 가능
- **실행 시**: Go가 필요 없음 (정적 링크된 바이너리)

### Linux 버전 호환성

현재 빌드 스크립트는 `CGO_ENABLED=0`로 정적 링크하여 빌드하므로:

- **GLIBC 의존성 없음**: 바이너리에 GLIBC가 정적으로 포함되어 있어 외부 라이브러리 의존성이 없습니다
- **지원되는 Linux 배포판**:
  - Ubuntu 18.04 LTS (GLIBC 2.27) 이상
  - Ubuntu 20.04 LTS (GLIBC 2.31) 이상 ✅
  - Ubuntu 22.04 LTS (GLIBC 2.35) 이상 ✅
  - Debian 10 (Buster) 이상
  - CentOS 7 이상
  - RHEL 7 이상
  - 기타 Linux 커널 3.10 이상의 배포판

**참고**: 정적 링크된 바이너리는 대부분의 최신 Linux 배포판에서 실행 가능하지만, 매우 오래된 배포판(커널 2.6 이하)에서는 실행되지 않을 수 있습니다.

## 문제 해결
### GLIBC 버전 호환성
Docker로 빌드한 바이너리가 서버에서 실행되지 않는 경우 (예: `GLIBC_2.32 not found`), 다음 방법을 시도하세요:

1. **네이티브 빌드 사용** (권장): 서버에서 직접 빌드
   ```bash
   ./build-linux-native.sh
   ```
   이 방법은 GLIBC 호환성 문제를 완전히 해결합니다.

2. **CGO 비활성화 빌드**: 현재 `build-linux.sh`는 `CGO_ENABLED=0`로 빌드하여 정적 링크를 사용합니다.
   - 이렇게 빌드된 바이너리는 GLIBC 의존성이 없어 대부분의 Linux 배포판에서 실행 가능합니다.
   - 만약 여전히 문제가 발생한다면, 서버에서 네이티브 빌드를 사용하는 것을 권장합니다.
   