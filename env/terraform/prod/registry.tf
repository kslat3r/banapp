
data "google_container_registry_repository" "banapp" {}

output "gcr_location" {
    value = "${data.google_container_registry_repository.banapp.repository_url}"
}