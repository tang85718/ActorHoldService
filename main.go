package main

import (
	consul "github.com/hashicorp/consul/api"
	"fmt"
	"strings"
	"github.com/micro/go-micro"
	"proto/asylum"
	"proto/crm"
	"proto/gm"
	"golang.org/x/net/context"
	"time"
	"log"
)

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	config := consul.DefaultConfig()
	config.Address = "localhost:8500"
	cli, err := consul.NewClient(config)
	failOnError(err)

	service := micro.NewService(micro.Name("ConsulMonitorService"))
	service.Init()

	asylum := asylum_api.NewAsylumServiceClient("AsylumService", service.Client())
	crm := crm_api.NewCRMServiceClient("crmService", service.Client())
	game := gm_api.NewGameServiceClient("GameService", service.Client())

	for {
		fmt.Println(" ")
		fmt.Println("Show Service List:")
		fmt.Println(" ")
		list, err := cli.Agent().Services()
		failOnError(err)

		for k, v := range list {
			if 0 == strings.Compare(k, "consul") {
				continue
			}

			fmt.Printf("[ %s ] %s\n", v.Service, v.ID)

			if v.Service == "AsylumService" {
				_, err := asylum.AsylumPing(context.TODO(), &asylum_api.AsylumPingReq{})
				if err != nil {
					log.Println(err)
					cli.Agent().ServiceDeregister(k)
				}
			}

			if v.Service == "crmService" {
				_, err := crm.CRMPing(context.TODO(), &crm_api.CRMPingReq{})
				if err != nil {
					log.Println(err)
					cli.Agent().ServiceDeregister(k)
				}
			}

			if v.Service == "GameService" {
				_, err := game.PingGame(context.TODO(), &gm_api.PingGameReq{})
				if err != nil {
					log.Println(err)
					cli.Agent().ServiceDeregister(v.ID)
				}
			}
		}
		time.Sleep(time.Second * 5)
	}
}
