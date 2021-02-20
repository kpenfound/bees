job "docs" {
  datacenters = ["dc1"]

  group "bees" {
    task "server" {
      driver = "docker"
      config {
        image = "bees:latest"
      }

      resources {
        memory = 128
      }
    }
  }
}