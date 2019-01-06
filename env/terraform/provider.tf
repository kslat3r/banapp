provider "google" {
  credentials = "${file("gcp-credentials.json")}"
  project = "banapp"
  region = "europe-west2"
  zone = "europe-west2-a"
}
