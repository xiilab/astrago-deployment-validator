package report

import (
	"encoding/json"
	"fmt"
	"html"
	"os"
	"path/filepath"
	"strings"

	"github.com/astrago/precheck/internal/models"
)

// Meta 리포트 메타 정보
type Meta struct {
	GeneratedAt  string `json:"generated_at"`
	KubeVIP      string `json:"kube_vip"`
	PortsChecked []int  `json:"ports_checked"`
}

// Report 전체 리포트 구조체
type Report struct {
	Meta  Meta                          `json:"meta"`
	Nodes map[string]*models.NodeReport `json:"nodes"`
}

// SaveAll JSON, Markdown, HTML 리포트를 모두 저장합니다
func SaveAll(r *Report, outputPath string) error {
	basePath := strings.TrimSuffix(outputPath, filepath.Ext(outputPath))

	// JSON 저장
	jsonPath := basePath + ".json"
	if err := SaveJSON(r, jsonPath); err != nil {
		return fmt.Errorf("JSON 저장 실패: %w", err)
	}

	// Markdown 저장
	mdPath := basePath + ".md"
	if err := SaveMarkdown(r, mdPath); err != nil {
		return fmt.Errorf("Markdown 저장 실패: %w", err)
	}

	// HTML 저장
	htmlPath := basePath + ".html"
	if err := SaveHTML(r, htmlPath); err != nil {
		return fmt.Errorf("HTML 저장 실패: %w", err)
	}

	return nil
}

// SaveJSON JSON 리포트를 저장합니다
func SaveJSON(r *Report, path string) error {
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// SaveMarkdown Markdown 리포트를 저장합니다
func SaveMarkdown(r *Report, path string) error {
	var lines []string

	// 헤더
	lines = append(lines, "# AstraGo Precheck Report\n")
	lines = append(lines, fmt.Sprintf("- 생성 시각: **%s**", r.Meta.GeneratedAt))
	lines = append(lines, fmt.Sprintf("- kube-vip: **%s**", r.Meta.KubeVIP))
	portsStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(r.Meta.PortsChecked)), ", "), "[]")
	lines = append(lines, fmt.Sprintf("- 체크 포트: `%s`\n", portsStr))
	lines = append(lines, "---\n")

	// 노드별 리포트
	for ip, nodeReport := range r.Nodes {
		lines = append(lines, generateNodeSection(ip, nodeReport)...)
	}

	return os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
}

// SaveHTML HTML 리포트를 저장합니다
func SaveHTML(r *Report, path string) error {
	mdPath := strings.TrimSuffix(path, ".html") + ".md"
	mdData, err := os.ReadFile(mdPath)
	if err != nil {
		return err
	}

	htmlBody := convertMarkdownToHTML(string(mdData))
	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="ko">
<head>
<meta charset="UTF-8">
<title>AstraGo Precheck Report</title>
<style>
body { font-family: Arial, sans-serif; padding: 20px; }
h1, h2 { color: #333; }
pre { background: #f4f4f4; padding: 12px; border-radius: 6px; }
</style>
</head>
<body>
%s
</body>
</html>`, htmlBody)

	return os.WriteFile(path, []byte(html), 0644)
}

func generateNodeSection(ip string, report *models.NodeReport) []string {
	var lines []string
	lines = append(lines, fmt.Sprintf("## Node: %s (role: %s)\n", ip, report.Role))

	addMarkdownSection(&lines, "Ping", report.Ping)
	addMarkdownSection(&lines, "SSH", report.SSH)
	addMarkdownSection(&lines, "Firewall", report.Firewall)

	// 포트 체크 결과 통합
	portTotal := true
	var portDetails []string
	for port, result := range report.Ports {
		if !result.OK {
			portTotal = false
		}
		portDetails = append(portDetails, fmt.Sprintf("[%s] %s: %s", port, formatStatus(result.OK), result.Detail))
	}
	portResult := models.CheckResult{
		OK:     portTotal,
		Detail: strings.Join(portDetails, "\n"),
	}
	addMarkdownSection(&lines, "Ports", portResult)

	addMarkdownSection(&lines, "Disk", report.Disk)
	addMarkdownSection(&lines, "resolv.conf", report.ResolvConf)
	addMarkdownSection(&lines, "sudo 권한", report.Sudo)
	addMarkdownSection(&lines, "NTP", report.NTP)
	addMarkdownSection(&lines, "CPU/Memory/GPU", report.CPUMemGPU)
	addMarkdownSection(&lines, "Internet", report.Internet)
	addMarkdownSection(&lines, "Swap", report.Swap)

	if report.KubeVIPNIC != nil {
		addMarkdownSection(&lines, "Kube VIP NIC", *report.KubeVIPNIC)
	}

	lines = append(lines, "---\n")
	return lines
}

func addMarkdownSection(lines *[]string, title string, result models.CheckResult) {
	*lines = append(*lines, fmt.Sprintf("### %s", title))
	*lines = append(*lines, fmt.Sprintf("- 결과: **%s**", formatStatus(result.OK)))
	*lines = append(*lines, "```")
	*lines = append(*lines, result.Detail)
	*lines = append(*lines, "```\n")
}

func formatStatus(ok bool) string {
	if ok {
		return "OK"
	}
	return "FAIL"
}

func convertMarkdownToHTML(mdText string) string {
	var htmlLines []string
	lines := strings.Split(mdText, "\n")
	inCode := false

	for _, line := range lines {
		if strings.HasPrefix(line, "```") {
			if !inCode {
				htmlLines = append(htmlLines, "<pre><code>")
				inCode = true
			} else {
				htmlLines = append(htmlLines, "</code></pre>")
				inCode = false
			}
			continue
		}

		if inCode {
			// 코드 블록 내부도 HTML 이스케이핑 필요
			htmlLines = append(htmlLines, html.EscapeString(line))
			continue
		}

		if strings.HasPrefix(line, "# ") {
			content := strings.TrimSpace(line[2:])
			htmlLines = append(htmlLines, fmt.Sprintf("<h1>%s</h1>", html.EscapeString(content)))
		} else if strings.HasPrefix(line, "## ") {
			content := strings.TrimSpace(line[3:])
			htmlLines = append(htmlLines, fmt.Sprintf("<h2>%s</h2>", html.EscapeString(content)))
		} else if strings.HasPrefix(line, "### ") {
			content := strings.TrimSpace(line[4:])
			htmlLines = append(htmlLines, fmt.Sprintf("<h3>%s</h3>", html.EscapeString(content)))
		} else if strings.HasPrefix(line, "- ") {
			content := strings.TrimSpace(line[2:])
			htmlLines = append(htmlLines, fmt.Sprintf("<li>%s</li>", html.EscapeString(content)))
		} else if strings.TrimSpace(line) == "---" {
			htmlLines = append(htmlLines, "<hr>")
		} else if strings.TrimSpace(line) != "" {
			htmlLines = append(htmlLines, fmt.Sprintf("<p>%s</p>", html.EscapeString(line)))
		}
	}

	return strings.Join(htmlLines, "\n")
}
