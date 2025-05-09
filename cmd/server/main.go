package main

import (
	"github.com/jinfanchris/emailsender/pkg/config"
	"github.com/jinfanchris/emailsender/pkg/log"
	"github.com/jinfanchris/emailsender/pkg/service"
	"net"

	pb "github.com/jinfanchris/emailsender/pkg/grpc/mailer"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var Port = "52333"

func main() {
	a := ParseArgs()

	if err := log.Setup(true); err != nil {
		panic(err)
	}

	logrus.Infof("args: %#v", a)

	c, err := config.LoadConfig(a.ConfigF)
	if err != nil {
		logrus.Panic(err)
	}

	logrus.Infof("config: %#v", c)

	creds, err := credentials.NewServerTLSFromFile(a.CertF, a.KeyF)
	if err != nil {
		logrus.Panic(err)
	}

	apiKey := c.ApiKey

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	mailServer := service.NewServer(
		apiKey,
		c.Sender.Account,
		c.Sender.Password,
		c.Sender.SMTPServer,
		c.Sender.Port,
	)

	pb.RegisterMailServiceServer(grpcServer, mailServer)

	lis, err := net.Listen("tcp", ":"+Port)
	if err != nil {
		logrus.Panic(err)
	}

	logrus.Infof("Listening on port %s", Port)
	if err := grpcServer.Serve(lis); err != nil {
		logrus.Panic(err)
	}
}
