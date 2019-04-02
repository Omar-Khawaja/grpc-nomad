job "greeting-gateway" {
  datacenters = ["dc1"]
  type = "service"

  group "greeting-gateway" {
    count = 1
    task "greeting-gateway" {
      driver = "docker"
      config {
        image = "omarkhawaja/greeting-gateway-grpc"
        port_map {
          http = 9090
        }
      }

      env {
	address = "fabio.service.consul:8888"
      }

      resources {
        network {
          mbits = 10
          port "http" {
	    static = 9090
	  }
        }
      }

      service {
        name = "greeting-gateway"
        port = "http"
      }
    }
  }
}
