package Ip

import (
	"strconv"
	"strings"
)

type IP struct {
	valores [4]int
}

func CrearIp(ip string) *IP {
	nuevaip := new(IP)
	numeros := strings.Split(ip, ".")
	nuevaip.valores = [4]int{}
	for i, e := range numeros {
		nuevaip.valores[i], _ = strconv.Atoi(e)
	}
	return nuevaip
}

func CompararIps(ip1, ip2 *IP) int {
	for i, e := range ip1.valores {
		if e > ip2.valores[i] {
			return 1
		} else if e < ip2.valores[i] {
			return -1
		}
	}
	return 0
}
