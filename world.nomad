variable "nomad_http_addr" {}

job "world" {
  datacenters = ["dc1"]

  group "world" {
    network {
      mode = "bridge"
    }

    restart {
      attempts = 0
    }

    reschedule {
      max_delay = "30s"
    }

    service {
      name = "world"

      connect {
        sidecar_service {
          proxy {
            upstreams {
              destination_name = "redis"
              local_bind_port  = 6379
            }
            upstreams {
              destination_name = "nomadserverproxy"
              local_bind_port  = 4646
            }
          }
        }
      }
    }

    task "world" {
      driver = "docker"

      env {
        BEE_REDIS_ADDR = "${NOMAD_UPSTREAM_ADDR_redis}"
        BEE_NOMAD_ADDR = var.nomad_http_addr
      }
      config {
        image = "ghcr.io/kpenfound/bees:latest"
        args = ["start"]
      }
    }
  }
}
