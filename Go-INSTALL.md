# Go 설치 가이드

이 문서는 AstraGo Deployment Validator를 빌드하기 위해 필요한 Go 프로그래밍 언어 설치 방법을 설명합니다.

## 필수 요구사항

- Go 1.21 이상
- Linux 환경 (또는 macOS/Windows에서 원격 체크)

## 설치 방법

### macOS

#### Homebrew를 사용한 설치 (권장)
```bash
brew install go
```

#### 공식 설치 파일 다운로드
1. https://go.dev/dl/ 에서 macOS용 .pkg 파일 다운로드
2. 다운로드한 .pkg 파일을 실행하여 설치

### Linux

#### Ubuntu/Debian (패키지 관리자)
```bash
sudo apt-get update
sudo apt-get install -y golang-go
```

#### Snap을 사용한 최신 버전 설치
```bash
sudo snap install go --classic
```

#### 공식 바이너리 설치 (권장)
```bash
# 최신 버전 다운로드 (버전은 https://go.dev/dl/ 에서 확인)
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz

# 기존 설치 제거 (있는 경우)
sudo rm -rf /usr/local/go

# 압축 해제 및 설치
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz

# PATH 환경 변수 설정
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

### Windows

1. https://go.dev/dl/ 에서 Windows용 설치 파일(.msi) 다운로드
2. 설치 파일 실행 및 설치
3. 설치 후 새 터미널에서 `go version` 명령으로 확인

## 설치 확인

설치가 완료되었는지 확인하려면 다음 명령을 실행하세요:

```bash
go version
```

출력 예시:
```
go version go1.21.5 darwin/amd64
```

## 환경 변수 설정 (필요시)

Go 작업 디렉토리를 설정하려면 다음 환경 변수를 설정할 수 있습니다:

```bash
# Go 작업 디렉토리 설정 (선택사항)
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

### 영구 설정

환경 변수를 영구적으로 설정하려면 셸 설정 파일에 추가하세요:

**Bash 사용자:**
```bash
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc
source ~/.bashrc
```

**Zsh 사용자:**
```bash
echo 'export GOPATH=$HOME/go' >> ~/.zshrc
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.zshrc
source ~/.zshrc
```

## 추가 정보

- 공식 Go 웹사이트: https://go.dev/
- 공식 설치 가이드: https://go.dev/doc/install
- 최신 버전 다운로드: https://go.dev/dl/

