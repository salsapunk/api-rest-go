# Receita para uma API-REST em Go

Primeiramente, criamos o `go.mod` utilizando o comando `go mod init` no terminal aberto no diretório root do projeto da API. Com isso feito, podemos prosseguir para a criação dos diretórios em que ficarão separadas as partes do código. Serão criados os seguintes diretórios:

<img src="/home/salsa/Workspace/Go/api1/readme-assets/diretorios.png" width="200">

## cmd

É o diretório onde ficará armazenado o `main.go`, o arquivo principal do nosso código.

## db

É o diretório onde ficará guardado o `db.go`, arquivo responsável pela comunicação com o banco de dados que, nesse projeto, é o PostgreSQL, com integração com Docker.

## handlers

É o diretório responsável por cuidar dos nossos _handlers_, utilizando uma struct que a conecta com o _repository_ (e este, por sua vez, conecta-se com o banco de dados). Os _handlers_ usam diretamente a biblioteca `net/http`, nativa de Go, conectando-se com o cliente e o server, fazendo suas requisições para o _repository_ e retornando para o cliente/usuário o resultado ou o status do erro.

## models

O diretório onde ficarão guardados os modelos que iremos usar como constantes, como a **Task** e os **comandos SQL**.

## repository

É onde ficará armazenado o arquivo `taskrepository.go`, responsável por interagir com o banco de dados, importando a biblioteca `database/sql`.

## routes

É onde terá o arquivo em que vamos definir as rotas do nosso mux (utilizando a biblioteca presente em github.com/gorilla/mux), que armazenará as rotas e as chamará de acordo com o caminho fornecido.
