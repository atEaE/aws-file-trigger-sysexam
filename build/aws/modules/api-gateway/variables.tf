variable "name" {
    description = "AWS内でプロジェクトリソースを特定するための名前を定義してください."
    type        = string
}

variable "resource_name" {
    description = "作成するAPI Gatewayのリソース名を設定してください.設定値はリソースのパスになります。"
    type        = string
}

variable "resource_description" {
    description = "作成するAPI Gatewayのリソース概要を記載してください."
    type        = string
}
