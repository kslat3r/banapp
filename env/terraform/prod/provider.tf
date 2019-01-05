provider "google" {
  credentials = "${file("../gcp-credentials.json")}"
  project = "ban-app"
  region = "europe-west2"
}
