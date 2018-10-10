// Links to terraform provider docs inline
// AWS Provider: https://www.terraform.io/docs/providers/aws/index.html

// Doc: https://www.terraform.io/docs/providers/aws/d/vpc.html
data "aws_vpc" "main" {
  id = "${var.vpc_id}"
}

// Doc: https://www.terraform.io/docs/providers/aws/r/iam_role.html
resource "aws_iam_role" "ec2-role" {
  name = "ec2-role"

  assume_role_policy = "${file("../instance-role-trust-policy.json")}"
}

// Doc: https://www.terraform.io/docs/providers/aws/r/iam_role_policy_attachment.html
resource "aws_iam_role_policy_attachment" "role-attach-1" {
    role       = "${aws_iam_role.ec2-role.name}"
    policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role"
}

// Doc: https://www.terraform.io/docs/providers/aws/r/iam_role_policy_attachment.html
resource "aws_iam_role_policy_attachment" "role-attach-2" {
    role       = "${aws_iam_role.ec2_role.name}"
    policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceRole"
}

// Doc: https://www.terraform.io/docs/providers/aws/r/iam_instance_profile.html
resource "aws_iam_instance_profile" "ec2-profile" {
  name = "ec2-profile"
  role = "${aws_iam_role.ec2-role.name}"
}

// Doc: https://www.terraform.io/docs/providers/aws/r/security_group.html
resource "aws_security_group" "allow_all" {
  name        = "allow_all"
  description = "Allow all inbound traffic"
  vpc_id      = "${aws_vpc.main.id}"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

}

// Doc: https://www.terraform.io/docs/providers/aws/r/ecs_cluster.html
resource "aws_ecs_cluster" "taquito" {
  name = "taquito"
}

// Doc: https://www.terraform.io/docs/providers/aws/r/instance.html
resource "aws_instance" "ec2-instance" {
  ami           = "${var.prefered_ami}"
  instance_type = "t2.micro"
  iam_instance_profile = "${aws_iam_instance_profile.ec2-profile.name}"
  security_groups = ["${aws_security_group.allow_all.id}"]
  user_data = "${file("../user_data.txt")}"
  associate_public_ip_address = true

  tags {
    Name = "Taquito"
  }
}

// Doc: https://www.terraform.io/docs/providers/aws/r/ecs_task_definition.html
resource "aws_ecs_task_definition" "taquito-task" {
  family                = "taquito"
  container_definitions = "${file("../task-definition.json")}"
}

// Doc: https://www.terraform.io/docs/providers/aws/r/ecs_service.html
resource "aws_ecs_service" "taquito-service" {
  name            = "taquito"
  cluster         = "${aws_ecs_cluster.taquito.id}"
  task_definition = "${aws_ecs_task_definition.taquito-task.arn}"
  desired_count   = 1
  iam_role        = "${aws_iam_role.ec2-role.arn}"
  depends_on      = ["aws_iam_role_policy.ec2-role"]
}
