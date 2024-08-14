package main

import (
	"fmt"

	"github.com/spaolacci/murmur3"
)

// Função para calcular o hash utilizando MurmurHash3
func murmurHash(tenant string) uint32 {
	return murmur3.Sum32([]byte(tenant))
}

// Função para determinar o shard a partir do valor de hash
func getShardByMurmurHash(tenant string, numShards int) int {
	hashValue := murmurHash(tenant)
	shard := hashValue % uint32(numShards)
	return int(shard)
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
		"Padaria-Estrela-Filial-3",
		"Hortifruti-Oba",
		"Acougue-Zona-Leste",
		"Acougue-Zona-Oeste",
		"Acougue-Zona-Norte",
		"Livraria-Cultura",
		"Drogaria-Sao-Paulo",
		"Supermercado-Extra",
		"Restaurantes-da-Fazenda",
		"Barbearia-Dois-Irmaos",
		"Salão-Beleza-Filial-1",
		"Salão-Beleza-Filial-2",
		"Auto-Peças-Sul",
		"Academias-BoaForma",
		"Escola-Livre",
		"Clínica-Medical",
		"Oficina-Mestre",
		"Padarias-Delicia",
		"Supermercado-Popular",
		"Restaurante-Maré",
		"Barber-Shop-Filial-1",
		"Barber-Shop-Filial-2",
		"Biblioteca-Publica",
		"Pet-Shop-Central",
		"Papelaria-Estudante",
		"Loja-Eletronicos",
		"Construtora-Axis",
		"Lavanderia-Prata",
		"Mercadinho-Barato",
		"Cinema-Central",
		"Loja-Roupas-Filial-1",
		"Loja-Roupas-Filial-2",
		"Café-BonsMomentos",
	}

	// Número de shards
	numShards := 5

	// Criação de um mapa para contar a distribuição dos tenants pelos shards
	shardDistribution := make(map[int]int)

	// Distribuição dos tenants nos shards
	for _, tenant := range tenants {
		shard := getShardByMurmurHash(tenant, numShards)
		fmt.Printf("Tenant: %s -> Shard: %d\n", tenant, shard)
		shardDistribution[shard]++
	}

	// Exibe a distribuição final dos shards
	fmt.Println("\nDistribuição Final dos Shards:")
	for i := 0; i < numShards; i++ {
		fmt.Printf("Shard %d: %d tenants\n", i, shardDistribution[i])
	}
}
