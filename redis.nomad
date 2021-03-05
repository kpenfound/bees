job "redis" {
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
}
