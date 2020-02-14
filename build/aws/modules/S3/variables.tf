variable "name" {
    description = ""
    type        = string
}

variable "bucket_name" {
    description = "AWS S3に設定するBucketの名前を設定してください.(DNS準拠の名前にしてください)"
    type        = string
}

variable "version_enable" {
    description = "Bucket内ファイルのVersioningを許可しない場合はfalseを設定してください.(Default = true))"
    type        = bool
    default     = true
}

