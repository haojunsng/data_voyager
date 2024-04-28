resource "aws_s3_bucket" "gomu_landing_bucket" {
  bucket        = "gomu-landing-bucket"
  force_destroy = true
}

resource "aws_s3_bucket_public_access_block" "privatise_bucket" {
  bucket = aws_s3_bucket.gomu_landing_bucket.id

  block_public_acls   = true
  block_public_policy = true
  ignore_public_acls  = true
}
