package main

import (
	"fmt"
	"os"

	nomad "github.com/hashicorp/nomad/api"
)

type NomadAPI struct {
	client *nomad.Client
}

func NewNomad() (*NomadAPI, error) {
	addr := os.Getenv("NOMAD_ADDR")
	if addr == "" {
		addr = "http://localhost:4646"
	}
	c, err := nomad.NewClient(&nomad.Config{Address: addr})

	return &NomadAPI{client: c}, err
}

func (n *NomadAPI) CreateBeeJob(b *Bee) error {
	jobsApi := n.client.Jobs()
	job, err := jobsApi.ParseHCL(NomadBeeHCL, true)
	if err != nil {
		return err
	}

	job.ID = &b.Id
	job.Name = &b.Id

	_, _, err = jobsApi.Register(job, &nomad.WriteOptions{})
	return err
}

var NomadBeeHCL = fmt.Sprintf(`
job "bee" {
	group "bee" {
		network {
		  mode = "bridge"
		}
		restart {
		  attempts = 2
		  delay    = "5s"
		}
		service "bee" {
		  name = "bee"
		  connect {
			sidecar_service {
			  proxy {
				upstreams {
				  destination_name = "redis"
				  local_bind_port  = 6379
				}
			  }
			}
		  }
		}
		task "bee" {
		  driver = "docker"
		  env {
			"REDIS_ADDR": "${NOMAD_UPSTREAM_ADDR_redis}"
		  }
		  config {
			image = "bees:local"
		  }
		} 
	  }
}`)
