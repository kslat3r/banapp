provider "google" {
  credentials = "${file("gcp-credentials.json")}"
  project = "banapp"
  region = "europe-west2"
  zone = "${local.zone}"
}
