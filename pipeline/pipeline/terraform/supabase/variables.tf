# Prefix your env vars with TF_VAR_
variable "supabase_pat" {
  type      = string
  sensitive = true
}

variable "supabase_organization_slug" {
  type      = string
  sensitive = true
}

variable "supabase_database_password" {
  type      = string
  sensitive = true
}
