package models

// CheckResult 체크 결과를 나타내는 구조체
type CheckResult struct {
	OK     bool   `json:"ok"`
	Detail string `json:"detail"`
}

// NodeReport 노드 체크 리포트를 나타내는 구조체
type NodeReport struct {
	IP          string                `json:"ip"`
	Role        string                `json:"role"`
	IsLocal     bool                  `json:"is_local"`
	Ping        CheckResult           `json:"ping"`
	SSH         CheckResult           `json:"ssh"`
	Firewall    CheckResult           `json:"firewall"`
	Ports       map[string]CheckResult `json:"ports"`
	Disk        CheckResult           `json:"disk"`
	ResolvConf  CheckResult           `json:"resolv_conf"`
	Sudo        CheckResult           `json:"sudo"`
	NTP         CheckResult           `json:"ntp"`
	CPUMemGPU   CheckResult           `json:"cpu_mem_gpu"`
	Internet    CheckResult           `json:"internet"`
	Swap        CheckResult           `json:"swap"`
	KubeVIPNIC  *CheckResult          `json:"kube_vip_nic,omitempty"`
}

