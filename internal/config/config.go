package config

import (
	"flag"
	"strconv"
	"strings"
)

const (
	DefaultSSHPort    = 22
	DefaultSSHUser    = "root"
	DefaultTimeout    = 10
	DefaultPingTimeout = 3
	DefaultTCPTimeout = 2
	DefaultSSHTimeout = 5

	DefaultPorts      = "22,6443,2379,10250"
	DefaultOutputFile = "precheck_report.json"

	RoleMaster  = "master"
	RoleUnknown = "unknown"
)

var (
	FirewallCommands = []Command{
		{Name: "firewalld", Cmd: "sudo firewall-cmd --state"},
		{Name: "ufw", Cmd: "sudo ufw status"},
		{Name: "iptables", Cmd: "sudo iptables -L"},
	}

	NTPCommands = []string{
		"timedatectl status",
		"chronyc tracking",
		"ntpq -p",
	}

	SystemInfoCommands = []string{
		"lscpu",
		"free -h",
		"nvidia-smi -L",
	}

	InternetTestCommands = []string{
		"ping -c 1 8.8.8.8",
		"ping -c 1 google.com",
	}
)

type Command struct {
	Name string
	Cmd  string
}

type Node struct {
	IP      string
	Role    string
	IsLocal bool
}

type Config struct {
	Nodes     []string
	Roles     []string
	LocalOnly bool

	User     string
	SSHPort  int
	KeyFile  string
	Password string

	Ports   []int
	KubeVIP string

	Output string
}

func ParseFlags() *Config {
	cfg := &Config{}

	flag.Var((*stringSlice)(&cfg.Nodes), "node", "체크할 노드 IP 주소 (여러 개 지정 가능)")
	flag.Var((*stringSlice)(&cfg.Roles), "role", "노드 역할 (master, worker 등)")
	flag.BoolVar(&cfg.LocalOnly, "local-only", false, "로컬 노드만 체크")

	flag.StringVar(&cfg.User, "user", DefaultSSHUser, "SSH 사용자명")
	flag.IntVar(&cfg.SSHPort, "port", DefaultSSHPort, "SSH 포트 번호")

	flag.StringVar(&cfg.KeyFile, "key", "", "SSH 키 파일 경로")
	flag.StringVar(&cfg.Password, "password", "", "SSH 비밀번호")

	portsStr := flag.String("ports", DefaultPorts, "체크할 포트 목록 (쉼표 구분)")
	flag.StringVar(&cfg.KubeVIP, "kube-vip", "", "Kubernetes VIP 주소")

	flag.StringVar(&cfg.Output, "output", DefaultOutputFile, "출력 파일 경로")

	flag.Parse()

	// 포트 파싱
	ports := strings.Split(*portsStr, ",")
	cfg.Ports = make([]int, 0, len(ports))
	for _, p := range ports {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		port, err := strconv.Atoi(p)
		if err != nil {
			continue
		}
		cfg.Ports = append(cfg.Ports, port)
	}

	return cfg
}

type stringSlice []string

func (s *stringSlice) String() string {
	return strings.Join(*s, ",")
}

func (s *stringSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}

