variable "hostname" {default = "Nessus Scanner"}
variable "region" { default = "us-east-1"}
variable "iam_role" { default= "TenableIO" }
# this is the key you got from Tenable
variable "tenable_linking_key" {}
variable "instance_type" { default = "t2.medium" }
