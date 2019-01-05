resource "google_container_cluster" "ban-app-staging" {
  name = "ban-app-staging"
  zone = "europe-west2-a"
  initial_node_count = 3
  
  node_config {
    machine_type = "n1-standard-1"
  }
}

output "get-credentials-cmd" {
  value = "gcloud container clusters get-credentials ${google_container_cluster.ban-app-staging.id} --zone ${google_container_cluster.ban-app-staging.zone}"
}