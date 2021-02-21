package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func NewNomad() *NomadAPI {
	addr := os.Getenv("NOMAD_ADDR")
	if addr == "" {
		addr = "http://localhost:4646"
	}
	api := NewAPI(fmt.Sprintf("%s/v1", addr))
	return &NomadAPI{
		api: api,
	}
}

type NomadAPI struct {
	api *API
}

func (n *NomadAPI) CreateJob(b *Bee) {
	job := b.GetJobspec()
	spec, err := json.Marshal(job)
	if err != nil {
		panic(err)
	}
	_, _ = n.api.Post("/jobs", spec)
}

func (n *NomadAPI) DeleteJob(b *Bee) {
	path := fmt.Sprintf("/job/%s?purge=true", b.Id)
	n.api.Delete(path)
}

type NomadJob struct {
	Job struct {
		ID          string   `json:"ID"`
		Name        string   `json:"Name"`
		Type        string   `json:"Type"`
		Datacenters []string `json:"Datacenters"`
		TaskGroups  []struct {
			Name  string `json:"Name"`
			Count int    `json:"Count"`
			Tasks []struct {
				Name   string `json:"Name"`
				Driver string `json:"Driver"`
				Config struct {
					Image string `json:"image"`
				} `json:"Config"`
				Resources struct {
					CPU      int `json:"CPU"`
					MemoryMB int `json:"MemoryMB"`
				} `json:"Resources"`
			} `json:"Tasks"`
		} `json:"TaskGroups"`
	} `json:"Job"`
}

var DefaultJob = fmt.Sprintf(`{
	"Job": {
	  "ID": "bzzz",
	  "Name": "bzzz",
	  "Type": "batch",
	  "Datacenters": ["dc1"],
	  "TaskGroups": [{
		  "Name": "bee",
		  "Count": 1,
		  "Tasks": [{
			  "Name": "bee",
			  "Driver": "docker",
			  "Config": {
				"image": "bees:local"
			  },
			  "Resources": {
				"CPU": 60,
				"MemoryMB": 128
			  }
			}]
		}]
	}
  }`)
