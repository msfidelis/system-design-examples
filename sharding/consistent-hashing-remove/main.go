package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Representa um nó no anel de hash.
type Node struct {
	ID   string
	Hash uint64
}

// Representa o hash ring que contém vários nós.
type ConsistentHashRing struct {
	Nodes       []Node
	NumReplicas int
}

// Cria um novo anel de hash ring.
func NewConsistentHashRing(numReplicas int) *ConsistentHashRing {
	return &ConsistentHashRing{
		Nodes:       []Node{},
		NumReplicas: numReplicas,
	}
}

// Adiciona um novo node ao Hash Ring
func (ring *ConsistentHashRing) AddNode(nodeID string) {
	for i := 0; i < ring.NumReplicas; i++ {
		replicaID := nodeID + strconv.Itoa(i)
		hash := hashTenant(replicaID)
		ring.Nodes = append(ring.Nodes, Node{ID: nodeID, Hash: hash})
	}
	sort.Slice(ring.Nodes, func(i, j int) bool {
		return ring.Nodes[i].Hash < ring.Nodes[j].Hash
	})
}

// Remove um node existente do Hash Ring
func (ring *ConsistentHashRing) RemoveNode(nodeID string) {
	var newNodes []Node
	for _, node := range ring.Nodes {
		if node.ID != nodeID {
			newNodes = append(newNodes, node)
		}
	}
	ring.Nodes = newNodes
	sort.Slice(ring.Nodes, func(i, j int) bool {
		return ring.Nodes[i].Hash < ring.Nodes[j].Hash
	})
}

// Retorna o node onde o Tenant deverá estar alocado
func (ring *ConsistentHashRing) GetTenantNode(key string) string {
	hash := hashTenant(key)
	idx := sort.Search(len(ring.Nodes), func(i int) bool {
		return ring.Nodes[i].Hash >= hash
	})

	// Se o índice estiver fora dos limites, retorna ao primeiro nó
	if idx == len(ring.Nodes) {
		idx = 0
	}

	return ring.Nodes[idx].ID
}

// Calcula o hash do tenant e a converte para uint64.
func hashTenant(s string) uint64 {
	s = strings.ToLower(s)
	hash := sha256.New()
	hash.Write([]byte(s))
	hashBytes := hash.Sum(nil)
	return binary.BigEndian.Uint64(hashBytes[:8])
}

func main() {
	// Cria um novo anel de hash consistente com 3 réplicas por nó.
	ring := NewConsistentHashRing(3)

	// Adiciona pseudo-nodes ao hash ring
	ring.AddNode("Shard-00")
	ring.AddNode("Shard-01")
	ring.AddNode("Shard-02")
	ring.AddNode("Shard-03")

	// Lista de Tenants
	keys := []string{
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

	// Distribuição Inicial dos Tenants pelos Nodes
	for _, key := range keys {
		node := ring.GetTenantNode(key)
		fmt.Printf("Tenant: %s, Node: %s\n", key, node)
	}

	// Remove um nó e exibe a nova distribuição de chaves.
	ring.RemoveNode("Shard-02")
	fmt.Println("\nRemovendo Shard-02:\n")
	for _, key := range keys {
		node := ring.GetTenantNode(key)
		fmt.Printf("Tenant: %s, Shard: %s\n", key, node)
	}

	ring.AddNode("Shard-04")
	fmt.Println("\nAdicionando Shard-04:\n")
	for _, key := range keys {
		node := ring.GetTenantNode(key)
		fmt.Printf("Tenant: %s, Shard: %s\n", key, node)
	}

}
