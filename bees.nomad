job "bees" {
  datacenters = ["dc1"]

  group "redis" {
    network {
      mode = "bridge"

      port "redis" {
        static = "6379"
        to = "6379"
      }
    }

    restart {
      attempts = 5
      delay    = "5s"
    }

    service {
      name = "redis"
      port = "6379"

      connect {
        sidecar_service {}
      }
    }

    task "redis" {
      driver = "docker"
      config {
        image = "redis:6.2.0"
      }

      resources {
        memory = 512
      }
    } 
  }

  group "world" {
    network {
      mode = "bridge"
    }

    restart {
      attempts = 0
    }

    reschedule {
      max_delay = "120s"
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
          }
        }
      }
    }

    task "world" {
      driver = "docker"

      env {
        REDIS_ADDR = "${NOMAD_UPSTREAM_ADDR_redis}"
      }
      config {
        image = "ghcr.io/kpenfound/bees:latest"
        entrypoint = ["/usr/bin/bees", "start"]
      }
    } 
  }
}
