resource "aws_iam_role" "ec2-role" {
  name = "ec2-role"

  assume_role_policy = "${file("../instance-role-trust-policy.json")}"
}

resource "aws_iam_role_policy_attachment" "role-attach-1" {
    role       = "${aws_iam_role.ec2-role.name}"
    policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role"
}

resource "aws_iam_role_policy_attachment" "role-attach-2" {
    role       = "${aws_iam_role.ec2_role.name}"
    policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceRole"
}

resource "aws_iam_instance_profile" "ec2-profile" {
  name = "ec2-profile"
  role = "${aws_iam_role.ec2-role.name}"
}