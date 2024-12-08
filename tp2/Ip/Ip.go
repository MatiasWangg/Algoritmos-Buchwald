package Ip

import (
	"strconv"
	"strings"
)

type IP struct {
	octetos [4]int
}

func CrearIp(ip string) *IP {
	nuevaip := new(IP)
	numeros := strings.Split(ip, ".")
	nuevaip.octetos = [4]int{}
	for i, e := range numeros {
		nuevaip.octetos[i], _ = strconv.Atoi(e)
	}
	return nuevaip
}

func CompararIps(ip1, ip2 *IP) int {
	for i, e := range ip1.octetos {
		if e > ip2.octetos[i] {
			return 1
		} else if e < ip2.octetos[i] {
			return -1
		}
	}
	return 0
}

func ObtenerOcteto(ip string, posicion int) int {
	octetos := strings.Split(ip, ".")
	valor, _ := strconv.Atoi(octetos[posicion])
	return valor
}
