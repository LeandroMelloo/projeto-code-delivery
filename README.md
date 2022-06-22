# projeto_code_delivery
Projeto Code Delivery

# Docker
docker-compose up -d
docker-compose ps
docker exec -it simulator bash
docker exec -it kafka-kafka-1 bash

# Go
gomod = gerenciador de controle de versão de pacotes externos do Go
go mod init github.com/projetocodedelivery/simulator
go run main.go

# Kafka
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position --group=terminal
kafka-console-producer --bootstrap-server=localhost:9092 --topic=route.new-direction
