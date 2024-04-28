# Configure api settings for the linked project
resource "supabase_settings" "production" {
  project_ref = supabase_project.strava_pipeline

  api = jsonencode({
    db_schema            = "public"
    db_extra_search_path = "public,extensions"
    max_rows             = 1000
  })
}
