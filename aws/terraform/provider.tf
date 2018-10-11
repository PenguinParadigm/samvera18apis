# AWS Provider
# set region 
provider "aws" {
  region  = "${var.aws_region}"
  profile = "${var.profile}"
  version = "1.40.0"
}
