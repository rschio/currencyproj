package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/rschio/currencyproj/currency/cmd/util"
	"github.com/rschio/currencyproj/currency/currency/proto"
	"github.com/rschio/currencyproj/currency/server"
	"github.com/rschio/currencyproj/currency/storage"
	"github.com/rschio/currencyproj/currency/storage/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	mysqlUser = util.MustEnv("MYSQL_USER")
	mysqlPass = util.MustEnv("MYSQL_PASS")
	mysqlPort = util.MustEnv("MYSQL_PORT")
	mysqlDB   = util.MustEnv("MYSQL_DB")
	port      = util.MustEnv("PORT")
)

func connectDB() storage.Storage {
	info := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?parseTime=true", mysqlUser, mysqlPass, mysqlPort, mysqlDB)
	conn, err := sql.Open("mysql", info)
	if err != nil {
		log.Fatalf("failed to connect DB: %v", err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatalf("failed to ping DB: %v", err)
	}
	return mysql.NewMySQL(conn)
}

func main() {
	l, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}

	db := connectDB()
	defer db.Close()

	srv := server.New(db)
	grpcSrv := grpc.NewServer()
	proto.RegisterExchangerServer(grpcSrv, srv)

	reflection.Register(grpcSrv)

	if err := grpcSrv.Serve(l); err != nil {
		log.Fatal(err)
	}
}
