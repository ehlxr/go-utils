package ip

import (
	"fmt"
	"net"
	"strings"
	"testing"
)

func TestIp(t *testing.T) {

	external_ip := get_external()

	external_ip = strings.Replace(external_ip, "\n", "", -1)
	fmt.Println("公网ip是: ", external_ip)

	fmt.Println("------Dividing Line------")

	ip := net.ParseIP(external_ip)
	if ip == nil {
		fmt.Println("您输入的不是有效的IP地址，请重新输入！")
	} else {
		result := TabaoAPI(string(external_ip))
		if result != nil {
			fmt.Println("国家：", result.Data.Country)
			fmt.Println("地区：", result.Data.Area)
			fmt.Println("城市：", result.Data.City)
			fmt.Println("运营商：", result.Data.Isp)
		}
	}

	fmt.Println("------Dividing Line------GetIntranetIp")

	GetIntranetIp()

	fmt.Println("------Dividing Line------")

	ip_int := inet_aton(net.ParseIP(external_ip))
	fmt.Println("Convert IPv4 address to decimal number(base 10) :", ip_int)

	ip_result := inet_ntoa(ip_int)
	fmt.Println("Convert decimal number(base 10) to IPv4 address:", ip_result)

	fmt.Println("------Dividing Line------")

	is_between := IpBetween(net.ParseIP("0.0.0.0"), net.ParseIP("255.255.255.255"), net.ParseIP(external_ip))
	fmt.Println("check result: ", is_between)

	fmt.Println("------Dividing Line------external_ip")
	is_public_ip := IsPublicIP(net.ParseIP(external_ip))
	fmt.Println("It is public ip: ", is_public_ip)

	is_public_ip = IsPublicIP(net.ParseIP("169.254.85.131"))
	fmt.Println("It is public ip: ", is_public_ip)

	fmt.Println("------Dividing Line------GetPulicIP")
	fmt.Println(GetPulicIP())
}
