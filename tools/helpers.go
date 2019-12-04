package tools

import (
	"github.com/9299381/wego/tools/tests"
	"math/rand"
	"net"
	"time"
)

func LocalIp() (string, error) {

	netInterfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String(), nil
					}
				}
			}
		}
	}
	return "", nil

}

func Test() *tests.TestStruct {
	return tests.NewTest()
}

// RandString 生成随机字符串
func RandString(length int, opt ...string) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if opt != nil {
		if opt[0] == "0" {
			str = "0123456789"
		} else if opt[0] == "a" {
			str = "abcdefghijklmnopqrstuvwxyz"
		} else if opt[0] == "A" {
			str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		} else if opt[0] == "aA" {
			str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		}
	}
	bytes := []byte(str)
	bytesLen := len(bytes)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(bytesLen)])
	}
	return string(result)
}
