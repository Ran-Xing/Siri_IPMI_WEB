package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"os"
	"os/exec"
)

var (
	USER      = "admin"
	PASSWORD  = "admin"
	IPADDRESS = "passwd"
	TOKEN     = "hhTp5eUSsc7iS5"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	var temp string
	if temp = os.Getenv("USER"); temp != "" {
		USER = temp
	}
	if temp = os.Getenv("PASSWORD"); temp != "" {
		PASSWORD = temp
	}
	if temp = os.Getenv("IPADDRESS"); temp != "" {
		IPADDRESS = temp
	}
}

func main() {
	log.Infof("Run models...")
	log.Infof("Server Start: http://%v:33659\n", getClientIp())
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusBadGateway, "Server error!")
	})

	cli := r.Use(Auth())
	cli.GET("/power", func(c *gin.Context) {
		typeTmp := c.Query("type")
		if typeTmp == "" {
			c.String(http.StatusBadGateway, "Server error!")
			return
		}
		cmd := exec.Command("ipmitool", "-I", "lan", "-U", USER, "-P", PASSWORD, "-H", IPADDRESS, "chassis", "power", typeTmp)
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Errorf("power %v failed with %s\n", typeTmp, err)
			c.String(http.StatusBadRequest, fmt.Sprintf("power %v failed with %v", typeTmp, err))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("power %v success with %v", typeTmp, string(out)))
	})
	r.GET("/fan", func(c *gin.Context) {
		//	ipmitool -I lan -U USER -P PASSWD -H IPADDRESS raw 0x30 0x70 0x66 0x01 0x01 0x20
		cmd := exec.Command("ipmitool", "-I", "lan", "-U", USER, "-P", PASSWORD, "-H", IPADDRESS, "raw", "0x30", "0x70", "0x66", "0x01", "0x01", "0x20")
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Errorf("set fan failed with %s\n", err)
			c.String(http.StatusBadRequest, fmt.Sprintf("set fan failed with %v", err))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("set fan success with %v", string(out)))
	})

	if err := r.Run(":33659"); err != nil {
		log.Errorf("Run Error: [%v]", err)
		return
	}

}

func getClientIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Errorf("获取本机 IP 失败: %v", err)
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tmp := context.Request.URL.Query().Get("token")
		if tmp == "" || tmp != TOKEN {
			context.String(http.StatusBadGateway, "Server error!")
			context.Abort()
			return
		}
		context.Next()
	}
}
