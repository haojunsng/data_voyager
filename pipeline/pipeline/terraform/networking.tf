variable "vpc_id" {
  type        = string
  default     = "vpc-063fb7fa08e361e2a"
  description = "Non-confidential AWS Default VPC ID"
}

variable "first_subnet_id" {
  type        = string
  default     = "subnet-017da5cd32e14f976"
  description = "Non-confidential AWS Default first subnet ID"
}

variable "second_subnet_id" {
  type        = string
  default     = "subnet-0a3caf18c6fea93a0"
  description = "Non-confidential AWS Default second subnet ID"
}

variable "security_group_id" {
  type        = string
  default     = "sg-020eeb87bcea509af"
  description = "Non-confidential AWS Default security group ID"
}
