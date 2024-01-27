package main

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
)

// Abstração de um Mecanismo IP Hashing
type IPHashBalancer struct {
	hosts []string
}

func NewIPHashBalancer(hosts []string) *IPHashBalancer {
	return &IPHashBalancer{hosts: hosts}
}

// Retorna o host com base no hash do endereço IP do cliente
func (ipb *IPHashBalancer) getHost(clientIP string) string {
	// Calcula o hash MD5 do endereço IP
	// Qualquer outro mecanismo de hashing pode ser utilizado
	hasher := md5.New()
	hasher.Write([]byte(clientIP))
	hashBytes := hasher.Sum(nil)

	// Calcula o index a partir dos 4 primeiros bytes do hash
	// Transformamos ele em um Integer para facilitar o exemplo
	// O resultado é um índice entre 0 e len(ipb.hosts) - 1
	// que são os índices válidos para a nossa slice de hosts.
	hashIndex := binary.BigEndian.Uint32(hashBytes[:4]) % uint32(len(ipb.hosts))
	return ipb.hosts[hashIndex]
}

func main() {
	hosts := []string{"http://host1.com", "http://host2.com", "http://host3.com"}
	ipHashBalancer := NewIPHashBalancer(hosts)

	// Define uma lista de IP's fakes
	clientIPs := []string{
		"192.168.1.1", "10.0.0.2", "10.0.0.3",
		"172.16.1.1", "172.16.1.2", "192.168.2.1", "192.168.2.2",
	}

	// Simula 30 Requisições
	for i := 0; i < 20; i++ {
		clientIP := clientIPs[i%len(clientIPs)]
		host := ipHashBalancer.getHost(clientIP)
		fmt.Printf("Requisição %d do IP %s direcionada para: %s\n", i+1, clientIP, host)
	}
}
