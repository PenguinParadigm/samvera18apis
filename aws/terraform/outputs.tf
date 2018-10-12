output public_endpoint {
    value = "${aws_instance.ec2-instance.public_dns}"
}