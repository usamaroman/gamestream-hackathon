package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"proc/pb"

	"github.com/google/uuid"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"google.golang.org/grpc"
)

const bucketName = "images"

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

	log.Println("starting server", cfg.Port)
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

func (srv *server) Produce(ctx context.Context, req *pb.ProduceRequest) (*pb.ProduceResponse, error) {
	p := req.Img.Value

	log.Println("size", len(p))

	fileName := uuid.New().String() + ".png"

	file, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = file.Write(convertToBytes(p))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer file.Close()

	info, err := srv.minio.FPutObject(ctx, bucketName, fileName, fileName, minio.PutObjectOptions{
		ContentType: "image/png",
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("minio size", info.Size)

	err = os.Remove(fileName)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.ProduceResponse{
		Status: pb.Status_Ok,
		Image:  fileName,
	}, nil
}

func (srv *server) Consume(ctx context.Context, req *pb.ConsumeRequest) (*pb.ConsumeResponse, error) {
	const img1 = "img1.png"
	const img2 = "img2.png"

	err := srv.minio.FGetObject(ctx, bucketName, req.Image1, img1, minio.GetObjectOptions{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = srv.minio.FGetObject(ctx, bucketName, req.Image2, img2, minio.GetObjectOptions{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp := &pb.ConsumeResponse{
		Image1: &pb.Image{},
		Image2: &pb.Image{},
	}

	p, err := os.ReadFile(img1)
	if err != nil {
		return nil, err
	}
	resp.Image1.Value = convertToUint32(p)

	p, err = os.ReadFile(img2)
	if err != nil {
		return nil, err
	}
	resp.Image2.Value = convertToUint32(p)

	return resp, nil
}

func convertToBytes(p []uint32) []byte {
	res := make([]byte, len(p))

	for i, b := range p {
		res[i] = byte(b)
	}

	return res
}

func convertToUint32(p []byte) []uint32 {
	res := make([]uint32, len(p))

	for i, b := range p {
		res[i] = uint32(b)
	}

	return res
}
