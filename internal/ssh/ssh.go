package ssh

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/astrago/precheck/internal/config"

	"golang.org/x/crypto/ssh"
)

// Client SSH 클라이언트 래퍼
type Client struct {
	hostname string
	username string
	port     int
	keyFile  string
	password string
	client   *ssh.Client
}

// NewClient 새로운 SSH 클라이언트를 생성합니다
func NewClient(hostname, username string, port int, keyFile, password string) *Client {
	return &Client{
		hostname: hostname,
		username: username,
		port:     port,
		keyFile:  keyFile,
		password: password,
	}
}

// Connect SSH 연결을 설정합니다
func (c *Client) Connect() error {
	sshConfig := &ssh.ClientConfig{
		User:            c.username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         config.DefaultSSHTimeout * time.Second,
	}

	// 키 파일 또는 비밀번호로 인증
	if c.keyFile != "" {
		key, err := readPrivateKey(c.keyFile)
		if err != nil {
			return fmt.Errorf("키 파일 읽기 실패: %w", err)
		}
		sshConfig.Auth = []ssh.AuthMethod{ssh.PublicKeys(key)}
	} else if c.password != "" {
		sshConfig.Auth = []ssh.AuthMethod{ssh.Password(c.password)}
	} else {
		return fmt.Errorf("SSH 키 또는 비밀번호가 필요합니다")
	}

	addr := fmt.Sprintf("%s:%d", c.hostname, c.port)
	client, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return fmt.Errorf("SSH 연결 실패: %w", err)
	}

	c.client = client
	return nil
}

// Run 원격 명령을 실행합니다
func (c *Client) Run(cmd string, timeout time.Duration) (int, string, string) {
	if c.client == nil {
		if err := c.Connect(); err != nil {
			return 1, "", err.Error()
		}
	}

	session, err := c.client.NewSession()
	if err != nil {
		return 1, "", err.Error()
	}
	defer session.Close()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan error, 1)
	var stdoutBuilder, stderrBuilder strings.Builder

	session.Stdout = &stdoutBuilder
	session.Stderr = &stderrBuilder

	go func() {
		done <- session.Run(cmd)
	}()

	select {
	case err := <-done:
		if err != nil {
			if exitErr, ok := err.(*ssh.ExitError); ok {
				return exitErr.ExitStatus(), strings.TrimSpace(stdoutBuilder.String()), strings.TrimSpace(stderrBuilder.String())
			}
			return 1, strings.TrimSpace(stdoutBuilder.String()), err.Error()
		}
		return 0, strings.TrimSpace(stdoutBuilder.String()), strings.TrimSpace(stderrBuilder.String())
	case <-ctx.Done():
		session.Close()
		return 1, "", "Timeout"
	}
}

// Close SSH 연결을 종료합니다
func (c *Client) Close() error {
	if c.client != nil {
		return c.client.Close()
	}
	return nil
}

func readPrivateKey(keyFile string) (ssh.Signer, error) {
	key, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	return signer, nil
}
