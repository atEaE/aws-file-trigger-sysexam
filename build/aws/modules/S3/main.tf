resource "aws_s3_bucket" "this" {
    bucket = ""
    acl    = ""

    versioning {
        enabled = false
    }

    tags = {
        Name = ""
    }
}