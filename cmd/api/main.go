package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/salsapunk/api-rest-go/internal/handler"
	"github.com/salsapunk/api-rest-go/internal/repository"
	"github.com/salsapunk/api-rest-go/internal/service"
)

func main() {

	// godotenv.Load() é usada para carregar o arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error openning enviroment file")
	}

	// cria o contexto e o seu cancelamento
	// WithCancel() retorna context e uma função de cancelamento cancel()
	// context.Background() retorna um não-nulo contexto vazio/zerado

	// o contexto é importante pois ele carrega sinais de cancelamentos, deadlines e outros
	// valores request-scoped pela API e pelos processos
	ctx, cancel := context.WithCancel(context.Background())
	// adia o cancelamento do contexto para o fim da função main
	defer cancel()

	// pgxpool.New() cria uma nova pool de conexões com o banco de dados
	// utilizando a URL presente em .env
	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	log.Println("Connected to the pool")

	// inicialização do repository, service e handlers.
	// o repository se comunica com o banco de dados, recebendo a pool,
	// o service cuida da lógica e regras de negócio, recebendo um ponteiro do repository e
	// o handler recebe as requisições http, passa para o service (que passa pro repository) e
	// retorna o JSON com as respostas
	TaskRepo := repository.NewTaskRepository(pool)
	TaskServ := service.NewTaskService(TaskRepo)
	TaskHand := handler.NewTaskHandler(TaskServ)

	// inicializa o multiplexer (tem o do gorilla, mas preferi usar o do próprio net/http)
	mux := http.NewServeMux()

	// define as rotas e as funções do handler que serão chamadas por elas
	mux.HandleFunc("/ping", TaskHand.HealthHandler)

	// define a "rota principal do multiplexer"
	http.Handle("/", mux)

	// avisa e roda a api na porta :8080
	log.Println("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
