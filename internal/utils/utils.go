package utils

import (
	"context"
	"fmt"
	"net"
	"os/exec"
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

// PingHost 호스트에 ping 테스트를 수행합니다
func PingHost(ip string, count int, timeout time.Duration) bool {
	var cmd string
	if runtime.GOOS == "windows" {
		cmd = fmt.Sprintf("ping -n %d -w %d %s", count, int(timeout.Seconds()*1000), ip)
	} else {
		cmd = fmt.Sprintf("ping -c %d -W %d %s", count, int(timeout.Seconds()), ip)
	}

	code, _, _ := RunLocalCmd(cmd, timeout+2*time.Second)
	return code == 0
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
