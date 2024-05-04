```bash
docker exec -it kafka kafka-topics --create --topic ecommerce_nova_venda --partitions 3 --replication-factor 1 --bootstrap-server localhost:9092
```