provider "google" {
  credentials = "${file("../gcp-credentials.json")}"
  project = "ban-app"
  region = "europe-west2"
  zone = "europe-west2-a"
}
