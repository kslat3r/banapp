provider "google" {
  credentials = "${file("../gcp-credentials.json")}"
  project = "bananapp"
  region = "europe-west2"
}
