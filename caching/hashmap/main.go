package main

import (
	"fmt"
	"sync"
)

// Define a estrutura para o nosso cache em memória com hashmap
type MemoryCache struct {
	items map[string]interface{}
	mutex sync.RWMutex // mutex simples para garantir a sincronização durante a leitura/escrita
}

// cacheInstance é uma instância do cache, será usado para implementar o padrão singleton
// Garantindo que a criação do cache seja realizada apenas uma vez, independente de quantas
// Vezes for recuperada pela aplicação
var cacheInstance *MemoryCache
var once sync.Once

// GetCacheInstance retorna a instância única do cache
func GetCacheInstance() *MemoryCache {
	once.Do(func() {
		cacheInstance = &MemoryCache{
			items: make(map[string]interface{}),
		}
	})
	return cacheInstance
}

// Adiciona ou atualiza um valor no cache com a chave fornecida
func (c *MemoryCache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.items[key] = value
}

// Get retorna um valor do cache se ele existir
func (c *MemoryCache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	value, found := c.items[key]
	return value, found
}

// Utilizando o padrão de cache criado
func main() {

	// Obtendo a instância do cache
	cache := GetCacheInstance()

	// Adicionando alguns usuários hipotéticos ao cache
	cache.Set("user:1", "Matheus Fidelis")
	cache.Set("user:2", "Tarsila Bianca")

	// Teste 1: Recuperando valores do cache
	if userName, found := cache.Get("user:1"); found {
		fmt.Println("Found user:1 ->", userName)
	} else {
		fmt.Println("user:1 not found in cache")
	}

	// Teste 2: Recuperando valores do cache
	if userName, found := cache.Get("user:2"); found {
		fmt.Println("Found user:2 ->", userName)
	} else {
		fmt.Println("user:2 not found in cache")
	}

	// Teste 3: Procurando um item que não existe em cache
	if userName, found := cache.Get("user:3"); found {
		fmt.Println("Found user:3 ->", userName)
	} else {
		fmt.Println("user:3 não encontrado em cache")
	}

}
