resource "aws_api_gateway_rest_api" "this" {
    name        = ""
    description = ""
}

resource "aws_api_gateway_resource" "this" {
    rest_api_id = "${aws_api_gateway_resource.this.id}"
    
}
