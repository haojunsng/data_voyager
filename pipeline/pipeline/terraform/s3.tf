resource "aws_s3_bucket" "landing_point" {
  bucket = "landing-point"
  force_destroy = true
}

resource "aws_s3_bucket_public_access_block" "privatise_bucket" {
  bucket = aws_s3_bucket.landing_point.id

  block_public_acls   = true
  block_public_policy = true
  ignore_public_acls  = true
}
