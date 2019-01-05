data "google_container_registry_repository" "ban-app" {}

output "gcr_location" {
    value = "${data.google_container_registry_repository.ban-app.repository_url}"
}