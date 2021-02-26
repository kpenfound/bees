data_dir  = "/opt/nomad/data"

bind_addr = "0.0.0.0" # the default

server {
  enabled          = true
  bootstrap_expect = 1
}

client {
  enabled       = true
  network_interface = "lo"
}

consul {
  address = "localhost:8500"
}
