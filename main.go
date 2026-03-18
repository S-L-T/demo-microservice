package main

import (
	repositoryimpl "github.com/S-L-T/demo-microservice/data/repository"
	"github.com/S-L-T/demo-microservice/domain/use_case"
	"github.com/S-L-T/demo-microservice/helper"
	presentationgrpchealth "github.com/S-L-T/demo-microservice/presentation/grpc/healthcheck"
	presentationgrpcuser "github.com/S-L-T/demo-microservice/presentation/grpc/user"
	presentationhttp "github.com/S-L-T/demo-microservice/presentation/http"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	err := helper.InitializeLogger(helper.TraceLevel)
	if err != nil {
		log.Fatal(err)
	}

	userRepo, err := repositoryimpl.NewMySQLUserRepository()
	if err != nil {
		helper.Log(err, helper.FatalLevel)
	}
	useCase := use_case.NewUserUseCase(&userRepo)
	go startHTTPServer(useCase)
	startGRPCServer(useCase)
}

func startHTTPServer(u use_case.User) {
	s := presentationhttp.NewServer(u)

	err := http.ListenAndServe(":8080", s.Router)

	if err != nil {
		helper.Log(err, helper.FatalLevel)
	}
}

func startGRPCServer(u use_case.User) {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		helper.Log(err, helper.FatalLevel)
	}

	s := grpc.NewServer([]grpc.ServerOption{}...)
	userServer := presentationgrpcuser.NewGRPCUserServer(u)
	presentationgrpcuser.RegisterUserServiceServer(s, &userServer)
	presentationgrpchealth.RegisterHealthServer(s, &presentationgrpchealth.HealthServerImpl{})

	err = s.Serve(listener)
	if err != nil {
		helper.Log(err, helper.FatalLevel)
	}

}
