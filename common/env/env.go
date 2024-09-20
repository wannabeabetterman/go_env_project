package env

import (
	"alger/config"
	"fmt"
	"os"
	"strings"
)

func CheckConfigPath() (prePath string) {
	//环境变量读取config路径
	configPath := os.Getenv("CONFIGPATH") // 部署环境默认 etc/resources
	if configPath == "" {
		return "./resources"
	}
	fmt.Printf("CONFIGPATH = %v\n", configPath)
	return configPath
}

func CheckIp(c *config.Config) {
	ip := os.Getenv("PODIP")
	fmt.Printf("PODIP = %v\n", ip)
	if ip == "" {
		c.Api.Host = strings.ReplaceAll(c.Api.Host, "${PODIP}", "0.0.0.0")
		//c.XxJob.ExecutorIp = strings.ReplaceAll(c.XxJob.ExecutorIp, "${PODIP}", "0.0.0.0")
		return
	}
	c.Api.Host = strings.ReplaceAll(c.Api.Host, "${PODIP}", ip)
	//c.XxJob.ExecutorIp = strings.ReplaceAll(c.XxJob.ExecutorIp, "${PODIP}", ip)
	return
}
