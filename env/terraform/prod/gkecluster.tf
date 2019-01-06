resource "google_container_cluster" "bananapp-prod" {
  name = "bananapp-prod"
  zone = "europe-west2-a"
  initial_node_count = 3
  
  node_config {
    machine_type = "n1-standard-1"
    oauth_scopes = [
      "compute-rw",
      "storage-ro",
      "logging-write",
      "monitoring"
    ]
  }
}

output "get-credentials-cmd" {
  value = "gcloud container clusters get-credentials ${google_container_cluster.bananapp-prod.id} --zone ${google_container_cluster.bananapp-prod.zone}"
}
