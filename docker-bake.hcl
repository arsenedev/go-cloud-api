
variable "IMAGE_NAME" {
  default = "ghcr.io/replace-me/go-cloud-api"
}

variable "TAG" {
  default = "dev"
}

group "default" {
  targets = ["api"]
}

target "api" {
  context    = "."
  dockerfile = "Dockerfile"

  tags = [
    "${IMAGE_NAME}:${TAG}"
  ]

  platforms = ["linux/amd64"]
}
