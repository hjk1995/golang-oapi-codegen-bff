provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" "gin_server" {
  ami           = "ami-0c55b159cbfafe1f0" # Update with a valid AMI ID for your region
  instance_type = "t2.micro"
  key_name      = "your-key-pair"         # Update with your EC2 key pair

  # Enable inbound access to port 8080
  vpc_security_group_ids = [aws_security_group.gin_sg.id]

  user_data = <<-EOF
              #!/bin/bash
              sudo yum update -y
              sudo yum install -y golang git
              git clone https://github.com/yourusername/gin-oapi-terraform.git
              cd gin-oapi-terraform
              go run main.go
              EOF

  tags = {
    Name = "Gin-OAPI-Server"
  }
}

resource "aws_security_group" "gin_sg" {
  name_prefix = "gin_sg_"

  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # Open to all (for testing purposes; restrict for production)
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
