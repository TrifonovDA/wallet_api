package main

import (
	"context"
	"fmt"
	"github.com/deliryo-io/wallet_api/config"
	Server "github.com/deliryo-io/wallet_api/internal/handlers"
	api "github.com/deliryo-io/wallet_api/internal/protoc"
	"github.com/deliryo-io/wallet_api/pkg/database_tools"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	server := grpc.NewServer()
	reflection.Register(server) //Создание образа для видимости необходимых для передачи атрибутов. В проде - к удалению
	ctx := context.Background()
	dbConnection := database_tools.NewConnection(ctx) //Создание пула коннектов к бд
	GRPCserver := Server.GRPCserver{dbConnection}

	srv := api.WalletApiServer(GRPCserver)
	api.RegisterWalletApiServer(server, srv)

	log.Println("Start to listen:")
	listener, err := net.Listen("tcp", config.Wallet_api_crd.Port) //Запуск листенера
	if err != nil {
		database_tools.CloseConnectionPool(dbConnection) //Закрытие пула коннектов к бд
		fmt.Println("Work is stopped!1")
		log.Fatal(err)
	}

	if err := server.Serve(listener); err != nil { //Serve возвращает ошибку, когда листенеры возвращают с фатальными ошибками.
		database_tools.CloseConnectionPool(dbConnection)
		fmt.Println("Work is stopped!2")
		log.Fatal(err)
	}
}
