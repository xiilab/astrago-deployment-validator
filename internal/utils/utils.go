package utils

import (
	"context"
	"fmt"
	"net"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"
)

// RunLocalCmd 로컬에서 쉘 명령을 실행합니다
func RunLocalCmd(cmd string, timeout time.Duration) (int, string, string) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var shell string
	var args []string

	if runtime.GOOS == "windows" {
		shell = "cmd"
		args = []string{"/C", cmd}
	} else {
		shell = "/bin/sh"
		args = []string{"-c", cmd}
	}

	execCmd := exec.CommandContext(ctx, shell, args...)
	stdout, err := execCmd.Output()
	stderr := ""
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			stderr = string(exitErr.Stderr)
			return exitErr.ExitCode(), strings.TrimSpace(string(stdout)), strings.TrimSpace(stderr)
		}
		return 1, strings.TrimSpace(string(stdout)), err.Error()
	}

	return 0, strings.TrimSpace(string(stdout)), strings.TrimSpace(stderr)
}

// validateIPOrHostname IP 주소 또는 호스트명 형식을 검증합니다
func validateIPOrHostname(addr string) bool {
	// IPv4 주소 검증
	if net.ParseIP(addr) != nil {
		return true
	}

	// IPv6 주소 검증 (대괄호 포함)
	if strings.HasPrefix(addr, "[") && strings.HasSuffix(addr, "]") {
		ipv6 := strings.Trim(addr, "[]")
		if net.ParseIP(ipv6) != nil {
			return true
		}
	}

	// 호스트명 검증 (도메인명, 호스트명)
	// 허용: 영문자, 숫자, 하이픈, 점, 최대 253자
	hostnameRegex := regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?(\.[a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?)*$`)
	if hostnameRegex.MatchString(addr) && len(addr) <= 253 {
		return true
	}

	return false
}

// PingHost 호스트에 ping 테스트를 수행합니다
// IP 주소 형식 검증을 통해 명령어 인젝션 공격을 방지합니다
func PingHost(ip string, count int, timeout time.Duration) bool {
	// IP 주소 또는 호스트명 형식 검증
	if !validateIPOrHostname(ip) {
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout+2*time.Second)
	defer cancel()

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		// Windows: ping -n [count] -w [timeout_ms] [ip]
		cmd = exec.CommandContext(ctx, "ping", "-n", fmt.Sprintf("%d", count), "-w", fmt.Sprintf("%d", int(timeout.Seconds()*1000)), ip)
	} else {
		// Linux/Unix: ping -c [count] -W [timeout_sec] [ip]
		cmd = exec.CommandContext(ctx, "ping", "-c", fmt.Sprintf("%d", count), "-W", fmt.Sprintf("%d", int(timeout.Seconds())), ip)
	}

	err := cmd.Run()
	return err == nil
}

// CheckTCPPort TCP 포트 접속 가능 여부를 확인합니다
func CheckTCPPort(ip string, port int, timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	dialer := &net.Dialer{
		Timeout: timeout,
	}
	conn, err := dialer.DialContext(ctx, "tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// SafeSplitLines 텍스트를 줄 단위로 분리하고 빈 줄을 제거합니다
func SafeSplitLines(text string) []string {
	lines := strings.Split(text, "\n")
	result := make([]string, 0, len(lines))
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			result = append(result, line)
		}
	}
	return result
}
