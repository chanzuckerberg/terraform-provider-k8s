resource "k8s_core_v1_service" "this" {
  metadata {
    name        = "${var.name}"
    namespace   = "${var.namespace}"
    labels      = "${local.labels}"
    annotations = "${var.annotations}"
  }

  spec {
    ports = [
      {
        name = "pulsar"
        port = "${var.port}"
      },
      {
        name = "http-admin"
        port = 8080
      },
    ]

    selector = "${local.labels}"
  }
}
