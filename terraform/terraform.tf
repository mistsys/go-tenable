provider "aws"{
  region = "${var.region}"
}


data "aws_ami" "nessus_scanner"{
  most_recent = true
  filter {
    name = "name"
    # We want the pre-authorized images
    values = ["Nessus*Managed*"]
  }
  owners = ["679593333241"]
}

resource "aws_security_group" "nessus_scanner"{
  name = "nessus_scanner"
  description = "Nessus Scanner to ping home and later to be allowed to scan internal servers"

  egress {
    from_port = 443
    to_port = 443
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_iam_role" "nessus_server_role" {
  name = "${var.iam_role}"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": "1"
    }
  ]
}
EOF
}

resource "aws_iam_policy_attachment" "nessus_server_role_attach" {
  name = "nessus_server_policy_attachment"
  roles = ["${aws_iam_role.nessus_server_role.name}"]
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ReadOnlyAccess"
}

resource "aws_iam_instance_profile" "nessus_server_instance_profile" {
  name = "nessus_server_instance_profile"
  role = "${aws_iam_role.nessus_server_role.name}"
}

resource "aws_instance" "nessus_server"{
  ami = "${data.aws_ami.nessus_scanner.id}"
  instance_type = "${var.instance_type}"
  iam_instance_profile = "${aws_iam_instance_profile.nessus_server_instance_profile.id}"
  # TODO: You want to do it in a VPC
  security_groups = ["${aws_security_group.nessus_scanner.name}"]
  user_data= <<EOF
{
  "name": "${var.hostname}",
  "key": "${var.tenable_linking_key}",
  "iam_role": "${var.iam_role}"
}
EOF
}

output "nessus_ami" {
  value = "${data.aws_ami.nessus_scanner.id}"
}
