package main

import (
	"fmt"
	"os"
	"time"

	"github.com/astrago/precheck/internal/checker"
	"github.com/astrago/precheck/internal/config"
	"github.com/astrago/precheck/internal/models"
	"github.com/astrago/precheck/internal/report"
)

func main() {
	cfg := config.ParseFlags()

	if err := run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "오류: %v\n", err)
		os.Exit(1)
	}
}

func run(cfg *config.Config) error {
	fmt.Println("=== AstraGo Precheck 시작 ===")

	nodes, err := buildNodeList(cfg)
	if err != nil {
		return err
	}

	reports := make(map[string]*models.NodeReport)

	for _, node := range nodes {
		fmt.Printf("[%s] 점검 시작...\n", node.IP)

		checker := checker.NewNodeChecker(node, cfg)
		report := checker.RunAll()
		reports[node.IP] = report

		fmt.Printf("[%s] 완료\n", node.IP)
	}

	result := &report.Report{
		Meta: report.Meta{
			GeneratedAt:  time.Now().Format("2006-01-02 15:04:05"),
			KubeVIP:      cfg.KubeVIP,
			PortsChecked: cfg.Ports,
		},
		Nodes: reports,
	}

	if err := report.SaveAll(result, cfg.Output); err != nil {
		return fmt.Errorf("리포트 저장 실패: %w", err)
	}

	fmt.Println("=== AstraGo Precheck 완료 ===")
	return nil
}

func buildNodeList(cfg *config.Config) ([]*config.Node, error) {
	if cfg.LocalOnly {
		hostname, err := os.Hostname()
		if err != nil {
			return nil, fmt.Errorf("호스트명 조회 실패: %w", err)
		}

		return []*config.Node{
			{
				IP:      hostname,
				Role:    "master",
				IsLocal: true,
			},
		}, nil
	}

	if len(cfg.Nodes) == 0 {
		return nil, fmt.Errorf("--node 옵션이 필요합니다")
	}

	nodes := make([]*config.Node, 0, len(cfg.Nodes))
	for i, ip := range cfg.Nodes {
		role := "unknown"
		if i < len(cfg.Roles) {
			role = cfg.Roles[i]
		}

		nodes = append(nodes, &config.Node{
			IP:      ip,
			Role:    role,
			IsLocal: false,
		})
	}

	return nodes, nil
}
