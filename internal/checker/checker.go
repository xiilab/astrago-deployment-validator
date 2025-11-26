package checker

import (
	"fmt"
	"strings"
	"time"

	"github.com/astrago/precheck/internal/config"
	"github.com/astrago/precheck/internal/models"
	"github.com/astrago/precheck/internal/ssh"
	"github.com/astrago/precheck/internal/utils"
)

// NodeChecker 노드 상태를 체크하는 구조체
type NodeChecker struct {
	node   *config.Node
	config *config.Config
	client *ssh.Client
}

// NewNodeChecker 새로운 NodeChecker를 생성합니다
func NewNodeChecker(node *config.Node, cfg *config.Config) *NodeChecker {
	return &NodeChecker{
		node:   node,
		config: cfg,
	}
}

// runCmd 로컬 또는 원격에서 명령을 실행합니다
func (nc *NodeChecker) runCmd(cmd string, timeout time.Duration) (int, string, string) {
	if nc.node.IsLocal {
		return utils.RunLocalCmd(cmd, timeout)
	}

	if nc.client == nil {
		nc.client = ssh.NewClient(nc.node.IP, nc.config.User, nc.config.SSHPort, nc.config.KeyFile, nc.config.Password)
		if err := nc.client.Connect(); err != nil {
			return 1, "", err.Error()
		}
	}

	return nc.client.Run(cmd, timeout)
}

// close SSH 연결을 종료합니다
func (nc *NodeChecker) close() {
	if nc.client != nil {
		nc.client.Close()
	}
}

// CheckPing Ping 테스트를 수행합니다
func (nc *NodeChecker) CheckPing() models.CheckResult {
	ok := utils.PingHost(nc.node.IP, 1, config.DefaultPingTimeout*time.Second)
	detail := "Ping 성공"
	if !ok {
		detail = "Ping 실패"
	}
	return models.CheckResult{OK: ok, Detail: detail}
}

// CheckSSH SSH 접속 가능 여부를 확인합니다
func (nc *NodeChecker) CheckSSH() models.CheckResult {
	if nc.node.IsLocal {
		return models.CheckResult{OK: true, Detail: "로컬 실행"}
	}

	client := ssh.NewClient(nc.node.IP, nc.config.User, nc.config.SSHPort, nc.config.KeyFile, nc.config.Password)
	if err := client.Connect(); err != nil {
		return models.CheckResult{OK: false, Detail: fmt.Sprintf("SSH 실패: %v", err)}
	}
	client.Close()
	return models.CheckResult{OK: true, Detail: "SSH 접속 성공"}
}

// CheckFirewall 방화벽 상태를 확인합니다
func (nc *NodeChecker) CheckFirewall() models.CheckResult {
	for _, cmd := range config.FirewallCommands {
		code, out, _ := nc.runCmd(cmd.Cmd, config.DefaultTimeout*time.Second)
		if code == 0 && out != "" {
			return models.CheckResult{OK: true, Detail: fmt.Sprintf("[%s]\n%s", cmd.Name, out)}
		}
	}
	return models.CheckResult{OK: false, Detail: "firewall 정보 없음"}
}

// CheckPorts 포트 접속 가능 여부를 확인합니다
func (nc *NodeChecker) CheckPorts() map[string]models.CheckResult {
	results := make(map[string]models.CheckResult)
	for _, port := range nc.config.Ports {
		tcpOK := utils.CheckTCPPort(nc.node.IP, port, config.DefaultTCPTimeout*time.Second)
		_, out, _ := nc.runCmd("ss -lnt || netstat -lnt", config.DefaultTimeout*time.Second)
		listening := strings.Contains(out, fmt.Sprintf(":%d ", port))
		portStr := fmt.Sprintf("%d", port)
		results[portStr] = models.CheckResult{
			OK:     tcpOK || listening,
			Detail: fmt.Sprintf("TCP접속:%v, listening:%v", tcpOK, listening),
		}
	}
	return results
}

// CheckDisk 디스크 사용량을 확인합니다
func (nc *NodeChecker) CheckDisk() models.CheckResult {
	code, out, err := nc.runCmd("df -hT || df -h", config.DefaultTimeout*time.Second)
	detail := out
	if detail == "" {
		detail = err
	}
	return models.CheckResult{OK: code == 0, Detail: detail}
}

// CheckResolv DNS 설정 파일을 확인합니다
func (nc *NodeChecker) CheckResolv() models.CheckResult {
	code, out, _ := nc.runCmd("cat /etc/resolv.conf", config.DefaultTimeout*time.Second)
	return models.CheckResult{OK: code == 0, Detail: out}
}

// CheckSudo sudo 권한을 확인합니다
func (nc *NodeChecker) CheckSudo() models.CheckResult {
	code, _, err := nc.runCmd("sudo -n true", config.DefaultTimeout*time.Second)
	if code == 0 {
		return models.CheckResult{OK: true, Detail: "sudo OK"}
	}
	return models.CheckResult{OK: false, Detail: err}
}

// CheckNTP NTP 동기화 상태를 확인합니다
func (nc *NodeChecker) CheckNTP() models.CheckResult {
	for _, cmd := range config.NTPCommands {
		code, out, _ := nc.runCmd(cmd, config.DefaultTimeout*time.Second)
		if code == 0 {
			return models.CheckResult{OK: true, Detail: out}
		}
	}
	return models.CheckResult{OK: false, Detail: "NTP 정보 조회 실패"}
}

// CheckCPUMemGPU CPU, Memory, GPU 정보를 확인합니다
func (nc *NodeChecker) CheckCPUMemGPU() models.CheckResult {
	var parts []string
	for _, cmd := range config.SystemInfoCommands {
		code, out, _ := nc.runCmd(cmd, config.DefaultTimeout*time.Second)
		if code == 0 {
			parts = append(parts, out)
		} else {
			parts = append(parts, fmt.Sprintf("%s 실패", cmd))
		}
	}
	return models.CheckResult{OK: true, Detail: strings.Join(parts, "\n\n")}
}

// CheckInternet 인터넷 연결 상태를 확인합니다
func (nc *NodeChecker) CheckInternet() models.CheckResult {
	var results []string
	ok := false
	for _, testCmd := range config.InternetTestCommands {
		code, _, _ := nc.runCmd(testCmd, config.DefaultTimeout*time.Second)
		status := "FAIL"
		if code == 0 {
			status = "OK"
			ok = true
		}
		results = append(results, fmt.Sprintf("%s: %s", testCmd, status))
	}
	return models.CheckResult{OK: ok, Detail: strings.Join(results, "\n")}
}

// CheckSwap Swap 상태를 확인합니다
func (nc *NodeChecker) CheckSwap() models.CheckResult {
	code, out, _ := nc.runCmd("swapon --show || cat /proc/swaps", config.DefaultTimeout*time.Second)
	if code == 0 && strings.TrimSpace(out) != "" {
		return models.CheckResult{OK: true, Detail: out}
	}
	return models.CheckResult{OK: false, Detail: "swap 비활성화"}
}

// CheckKubeVIPNIC VIP와 노드 IP의 NIC를 비교합니다
func (nc *NodeChecker) CheckKubeVIPNIC() *models.CheckResult {
	if nc.config.KubeVIP == "" || nc.node.Role != config.RoleMaster {
		return &models.CheckResult{OK: false, Detail: "master node 또는 VIP 미지정"}
	}

	code, out, _ := nc.runCmd("ip -4 addr show", config.DefaultTimeout*time.Second)
	if code != 0 {
		return &models.CheckResult{OK: false, Detail: "ip addr 실패"}
	}

	var nodeNIC, vipNIC, currentNIC string
	lines := utils.SafeSplitLines(out)

	for _, line := range lines {
		if strings.Contains(line, ":") && len(line) > 0 && line[0] >= '0' && line[0] <= '9' {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				nicPart := strings.Split(parts[1], "@")[0]
				currentNIC = strings.TrimSpace(nicPart)
			}
		}

		if strings.Contains(line, "inet ") {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				ip := strings.Split(fields[1], "/")[0]
				if ip == nc.node.IP {
					nodeNIC = currentNIC
				}
				if ip == nc.config.KubeVIP {
					vipNIC = currentNIC
				}
			}
		}
	}

	if nodeNIC == "" || vipNIC == "" {
		return &models.CheckResult{
			OK:     false,
			Detail: fmt.Sprintf("nic 정보 부족 node:%s, vip:%s", nodeNIC, vipNIC),
		}
	}

	same := nodeNIC == vipNIC
	return &models.CheckResult{
		OK:     same,
		Detail: fmt.Sprintf("node_nic=%s, vip_nic=%s", nodeNIC, vipNIC),
	}
}

// RunAll 모든 체크를 실행하고 결과를 반환합니다
func (nc *NodeChecker) RunAll() *models.NodeReport {
	defer nc.close()

	return &models.NodeReport{
		IP:         nc.node.IP,
		Role:       nc.node.Role,
		IsLocal:    nc.node.IsLocal,
		Ping:       nc.CheckPing(),
		SSH:        nc.CheckSSH(),
		Firewall:   nc.CheckFirewall(),
		Ports:      nc.CheckPorts(),
		Disk:       nc.CheckDisk(),
		ResolvConf: nc.CheckResolv(),
		Sudo:       nc.CheckSudo(),
		NTP:        nc.CheckNTP(),
		CPUMemGPU:  nc.CheckCPUMemGPU(),
		Internet:   nc.CheckInternet(),
		Swap:       nc.CheckSwap(),
		KubeVIPNIC: nc.CheckKubeVIPNIC(),
	}
}
