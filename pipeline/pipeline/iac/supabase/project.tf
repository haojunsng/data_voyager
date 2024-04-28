resource "supabase_project" "strava_pipeline" {
  organization_id   = var.supabase_organization_slug
  name              = "strava_pipeline"
  database_password = var.supabase_database_password
  region            = "ap-southeast-1"

  lifecycle {
    ignore_changes = [database_password]
  }
}
