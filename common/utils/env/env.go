package env

import (
	"fmt"
	"os"
	"strings"
)

const (
	ENVPATH      string = "CONFIGPATH"
	DEFAULTPATH  string = "./resources"
	ENVIP        string = "PODIP"
	ENVIPREPLACE string = "${PODIP}"
	DEFAULTIP    string = "0.0.0.0"
)

func CheckPath() (prePath string) {
	//环境变量读取config路径
	path := os.Getenv(ENVPATH) // 部署环境默认 etc/resources
	if path == "" {
		return DEFAULTPATH
	}
	fmt.Sprintf(ENVPATH+" = %v\n", path)
	return path
}

func CheckIp(olds ...*string) {
	ip := os.Getenv(ENVIP)
	fmt.Sprintf(ENVIP+" = %s", ip)
	for i := 0; i < len(olds); i++ {
		if ip == "" {
			*olds[i] = strings.ReplaceAll(*olds[i], ENVIPREPLACE, DEFAULTIP)
		} else {
			*olds[i] = strings.ReplaceAll(*olds[i], ENVIPREPLACE, ip)
		}
	}
}
