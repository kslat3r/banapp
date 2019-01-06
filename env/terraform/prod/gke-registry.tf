data "google_container_registry_repository" "bananapp" {}

output "gcr_location" {
    value = "${data.google_container_registry_repository.bananapp.repository_url}"
}