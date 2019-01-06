locals {
  zone = "europe-west2-a"
  initial_node_count = 3
  machine_type = "n1-standard-1"
  disk_size_gb = 10
  oauth_scopes = [
    "https://www.googleapis.com/auth/devstorage.read_only",
    "https://www.googleapis.com/auth/logging.write",
    "https://www.googleapis.com/auth/monitoring",
    "https://www.googleapis.com/auth/servicecontrol",
    "https://www.googleapis.com/auth/service.management.readonly",
    "https://www.googleapis.com/auth/trace.append"
  ]
  get_credentials = "gcloud container clusters get-credentials"
}

