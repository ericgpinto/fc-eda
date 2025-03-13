# FC Wallet

FC Wallet é um sistema baseado em microsserviços para gerenciamento de transações financeiras entre contas. Ele é composto por dois microsserviços que se comunicam via Kafka.

## Arquitetura

O sistema é dividido em dois microsserviços principais:

- **fc-wallet-core** (Go): Desenvolvido durante o módulo de arquitetura baseada em eventos do curso FullCycle. Responsável por registrar novas transações e produzir mensagens para o Kafka.
- **fc-wallet-balance** (Java com Spring): Desenvolvido a parte para integrar com serviço wallet-core. Consome mensagens do Kafka para processar e atualizar os saldos das contas.

## Tecnologias Utilizadas

- **Linguagens**: Go, Java (Spring Boot)
- **Mensageria**: Apache Kafka
- **Banco de Dados**: MySQL
- **Containerização**: Docker, Docker Compose
- **Orquestração**: Docker Compose

## Estrutura do Projeto

```
fc-wallet/
│-- fc-wallet-core/        # Microsserviço em Go
│-- fc-wallet-balance/     # Microsserviço em Java com Spring Boot
│-- docker-compose.yml     # Configuração para levantar os containers
│-- README.md              # Documentação do projeto
```

## Como Executar

1. Clone o repositório:

   ```sh
   git clone https://github.com/seu-usuario/fc-wallet.git
   cd fc-wallet
   ```

2. Suba os containers com Docker Compose:

   ```sh
   docker-compose up -d
   ```

3. Verifique se os containers estão rodando:

   ```sh
   docker ps
   ```

## Configuração dos Microsserviços

### fc-wallet-core (Go)

- Porta padrão: `8080`
- Produz mensagens para o Kafka no tópico `transactions`
- Banco de dados: MySQL

### fc-wallet-balance (Java/Spring Boot)

- Porta padrão: `3003`
- Consome mensagens do tópico `transactions` no Kafka
- Atualiza os saldos das contas no banco de dados MySQL
- Exibe saldo das contas atualizado.

## Testes

Para testar a API do `fc-wallet-core`, utilize ferramentas como Postman ou cURL:

```sh
curl -X POST http://localhost:8080/transactions -d '{
    "account_id_from": "1b94b998-1f92-4897-a5e2-24bde6685b5d",
    "account_id_to": "bb835285-769c-439f-b1cb-a8788bdf8e72",
    "amount": 25
}' -H "Content-Type: application/json"
```

Utilize também o arquivo client.http dentro do projeto fc-wallet-core para realizar testes.

Para verificar os logs do consumidor `fc-wallet-balance`:

```sh
docker logs -f fc-wallet-balance
```

Para testar a API do `fc-wallet-balance` e ver o saldo das contas atualizado, utilize ferramentas como Postman ou cURL:

```sh
curl -X GET http://localhost:3003/balances/1b94b998-1f92-4897-a5e2-24bde6685b5d -H "Content-Type: application/json"
```

## Contribuição

Sinta-se à vontade para abrir issues e pull requests para melhorias no projeto.
