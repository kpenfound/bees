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
	addr := os.Getenv("BEE_NOMAD_ADDR")
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

	dc := os.Getenv("NOMAD_DC")
	if dc == "" {
		dc = "dc1"
	}

	job.Datacenters = append(job.Datacenters, dc)

	_, _, err = jobsApi.Register(job, &nomad.WriteOptions{})
	return err
}

var NomadBeeHCL = fmt.Sprintf(`
job "bee" {
	group "bee" {
		network {
		  mode = "bridge"
		}
		service {
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
				BEE_REDIS_ADDR = "${NOMAD_UPSTREAM_ADDR_redis}"
		  }
		  config {
				image = "ghcr.io/kpenfound/bees:latest"
				args = ["bee"]
		  }
		}
	}
}`)
