package main

import (
	"github.com/DoktorGhost/golibrary/pkg/users/internal/config"
	"github.com/DoktorGhost/golibrary/pkg/users/internal/provider/grpcProvider"
	"github.com/DoktorGhost/golibrary/pkg/users/internal/repositories/postgres"
	pb "github.com/DoktorGhost/golibrary/pkg/users/pkg/proto"
	"github.com/DoktorGhost/golibrary/pkg/users/pkg/storage/psg"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"

	"github.com/DoktorGhost/golibrary/pkg/users/internal/services"
)

func main() {
	//считываем переменные окружения
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println(err)
	}

	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		log.Println(err)
	}

	//экземпляр бд
	var db *pgxpool.Pool
	for i := 0; i < 5; i++ {
		db, err = psg.InitStorage(cfg, schema)
		if err == nil {
			log.Println("Connected to database")
			break
		}
		log.Println("Ошибка подключения к бд:", err, "попытка ", i+1)
		time.Sleep(5 * time.Second)
	}

	defer db.Close()

	userRepo := postgres.NewPostgresRepository(db)
	userService := services.NewUserService(userRepo)

	grpcServer := grpc.NewServer()
	userGRPCServer := grpcProvider.NewUserGRPCServer(userService)

	lis, err := net.Listen("tcp", ":"+cfg.Provider_port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pb.RegisterUsersServiceServer(grpcServer, userGRPCServer)
	reflection.Register(grpcServer)

	log.Println("gRPC server listening on :", cfg.Provider_port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}

}
