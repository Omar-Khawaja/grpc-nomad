job "greeting-server" {
  datacenters = ["dc1"]
  type = "service"

  group "greeting-server" {
    count = 1
    task "greeting-server" {
      driver = "docker"
      config {
        image = "omarkhawaja/greeting-server-grpc"
        port_map {
          http = 8080
        }
      }

      resources {
        network {
          mbits = 10
          port "http" {}
        }
      }

      service {
        name = "greeting-server"
	tags = [
		"urlprefix-/" 
	]
        port = "http"
	check {
          type     = "tcp"
          interval = "2s"
          timeout  = "2s"
        }
      }
    }
  }
}
