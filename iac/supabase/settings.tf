# Configure api settings for the linked project
resource "supabase_settings" "strava_pipeline" {
  project_ref = var.supabase_project_ref

  api = jsonencode({
    db_schema            = "public"
    db_extra_search_path = "public,extensions"
    max_rows             = 1000
  })
}
