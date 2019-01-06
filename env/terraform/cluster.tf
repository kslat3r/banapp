resource "google_container_cluster" "banapp-prod" {
  name = "banapp-prod"
  zone = "${local.zone}"
  initial_node_count = "${local.initial_node_count}"
  
  node_config {
    machine_type = "${local.machine_type}"
    disk_size_gb = "${local.disk_size_gb}"
    oauth_scopes = "${local.oauth_scopes}"
  }
}

resource "google_container_cluster" "banapp-staging" {
  name = "banapp-staging"
  zone = "${local.zone}"
  initial_node_count = "${local.initial_node_count}"
  
  node_config {
    machine_type = "${local.machine_type}"
    disk_size_gb = "${local.disk_size_gb}"
    oauth_scopes = "${local.oauth_scopes}"
  }
}

output "get-credentials-prod-cmd" {
  value = "${local.get_credentials} ${google_container_cluster.banapp-prod.id} --zone ${google_container_cluster.banapp-prod.zone}"
}

output "get-credentials-staging-cmd" {
  value = "${local.get_credentials} ${google_container_cluster.banapp-staging.id} --zone ${google_container_cluster.banapp-staging.zone}"
}