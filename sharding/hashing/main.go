package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strings"
)

// Calcula o hash SHA-256 do valor do tenant e o converte para um inteiro
func hashTenant(tenant string) int {
	tenant = strings.ToLower(tenant) // Converte o valor do tenant para minúsculas
	hash := sha256.New()
	hash.Write([]byte(tenant))
	hashBytes := hash.Sum(nil)
	hashInt := binary.BigEndian.Uint64(hashBytes)

	// Converte o valor para um valor positivo caso o hashint venha a ser negativo
	if int(hashInt) < 0 {
		return -int(hashInt)
	}
	return int(hashInt)
}

// Recebe uma string do tenant e o número de shards, retornando o número do shard correspondente
func getShardByTenant(tenant string) int {
	numShards := 3

	hashValue := hashTenant(tenant)
	shard := hashValue % numShards
	return shard
}

func main() {
	// Lista de tenants
	tenants := []string{
		"Petshops-Souza",
		"Pizzarias-Carvalho",
		"Mecanica-Dois-Irmaos",
		"Padaria-Estrela-Filial-1",
		"Padaria-Estrela-Filial-2",
		"Padaria-Estrela-Filial-3",
		"Hortifruti-Oba",
		"Acougue-Zona-Leste",
		"Acougue-Zona-Oeste",
		"Acougue-Zona-Norte",
	}

	// Verifica a distribuição dos tenants entre os shards
	for _, tenant := range tenants {
		shard := getShardByTenant(tenant)
		fmt.Printf("Tenant: %s, Shard: %d\n", tenant, shard)
	}
}
