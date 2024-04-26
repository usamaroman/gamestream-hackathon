package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"proc/pb"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"google.golang.org/grpc"
)

type config struct {
	Port  string `yaml:"port" env:"PORT" env-default:"8080"`
	Minio struct {
		Host     string `yaml:"host" env:"MINIO_HOST" env-default:"localhost"`
		Port     string `yaml:"port" env:"MINIO_PORT" env-default:"9000"`
		User     string `yaml:"user" env:"MINIO_USER" env-default:"minio"`
		Password string `yaml:"password" env:"MINIO_PASSWORD" env-default:"minio123"`
	} `yaml:"minio"`
}

func newCfg() (*config, error) {
	var cfg config

	err := cleanenv.ReadConfig("config.yaml", &cfg)
	if err != nil {
		return nil, err
	}

	log.Println("proc configuration", cfg)

	return &cfg, nil
}

type server struct {
	cfg   *config
	minio *minio.Client
	pb.UnimplementedImageServiceServer
}

func newMinio(cfg *config) *minio.Client {
	endpoint := fmt.Sprintf("%s:%s", cfg.Minio.Host, cfg.Minio.Port)
	accessKeyID := cfg.Minio.User
	secretAccessKey := cfg.Minio.Password

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Fatal(err)
	}

	bucketName := "images"
	location := "BLR"
	ctx := context.Background()

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Println("We already own")
		} else {
			log.Fatal(err)
		}
	} else {
		log.Println("Successfully created")
	}

	return minioClient
}

func main() {
	cfg, err := newCfg()
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", cfg.Port))
	if err != nil {
		log.Fatal("failed to listen", err)
	}

	srv := &server{
		cfg:   cfg,
		minio: newMinio(cfg),
	}

	rpcSrv := grpc.NewServer(
		grpc.UnaryInterceptor(srv.logInterceptor),
	)

	pb.RegisterImageServiceServer(rpcSrv, srv)

	err = rpcSrv.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}

func (srv *server) logInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("[Storage Service Interceptor]", info.FullMethod)

	m, err := handler(ctx, req)

	log.Println("post proc message", m)

	return m, err
}
