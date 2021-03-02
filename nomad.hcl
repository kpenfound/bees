data_dir  = "/opt/nomad/data"

bind_addr = "0.0.0.0" # the default

server {
  enabled          = true
  bootstrap_expect = 1
}

client {
  enabled       = true
}

consul {
  address = "localhost:8500"
}

plugin_dir = "/opt/nomad/plugins"

plugin "nomad-driver-podman" {
  config {
    volumes {
      enabled = true
    }
    recover_stopped = false
  }
}
