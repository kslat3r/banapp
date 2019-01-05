resource "google_container_cluster" "ban-app-prod" {
  name = "ban-app-prod"
  zone = "europe-west2-a"
  initial_node_count = 3
  
  node_config {
    machine_type = "n1-standard-1"
  }
}

output "get-credentials-cmd" {
  value = "gcloud container clusters get-credentials ${google_container_cluster.ban-app-prod.id} --zone ${google_container_cluster.ban-app-prod.zone}"
}
