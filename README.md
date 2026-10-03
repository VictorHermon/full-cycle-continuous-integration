# full-cycle-continuous-integration

Exercício do curso **Full Cycle** sobre integração contínua com GitHub Actions, usando um programa simples em Go.

## O que o pipeline faz

A cada pull request para `develop`, o workflow [`ci.yaml`](.github/workflows/ci.yaml):

1. Configura o Go com uma matrix de versões.
2. Executa o programa (`go run math.go`).
3. Roda os testes (`go test -v`).
4. Monta a imagem Docker com Buildx, sem publicar.

## Como rodar localmente

```bash
go run math.go     # imprime o resultado da soma
go test -v         # roda os testes
docker build -t full-cycle-ci .
docker run --rm full-cycle-ci
```

Há também uma configuração de **Dev Container** em `.devcontainer/`.

## Tecnologias

Go · GitHub Actions · Docker · Dev Containers
