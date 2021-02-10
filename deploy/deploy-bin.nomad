job "onhub-web" {
  datacenters = ["dc1"]

  group "onhub-web" {
    task "service" {
      driver = "raw_exec"

      config {
        command = "/opt/software/onhub-web/onhub-web"
      }

      resources {
        network {
          port "http" {
            static = 4209
          }
        }
      }
    }
  }
}
