ports {
  grpc = 8502
}

connect {
  enabled = true
}

server = true
bootstrap = true

client_addr = "0.0.0.0"
bind_addr = "127.0.0.1"

data_dir = "/opt/consul/"

ui_config {
  enabled = true
}
