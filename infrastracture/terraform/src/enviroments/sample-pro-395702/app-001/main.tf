provider "google" {
  project = var.project_id
  region  = var.region
}

resource "google_cloud_run_service" "sample-app-001" {
  name     = "sample-app-001"
  location = var.region

  template {
    spec {
      containers {
        image = ""

        env {
          name  = "HOGEHOGE"
          value = "おりゃあaaa"
        }
      }
    }
  }

  lifecycle {
    ignore_changes = [
      template[0].spec[0].containers[0].image,
    ]
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

resource "google_cloud_run_service_iam_member" "public_access" {
  service  = "sample-app-001"
  location = var.region
  role     = "roles/run.invoker"
  member   = "allUsers"
}
