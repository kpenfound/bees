ports {
  grpc = 8502
}

connect {
  enabled = true
}

server = true
bootstrap = true

client_addr = "0.0.0.0"
data_dir = "~/.consul/"

ui_config {
  enabled = true
}