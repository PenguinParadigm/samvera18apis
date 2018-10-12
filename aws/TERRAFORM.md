# Deploying TAQUITO using Terraform

[Terraform](https://www.terraform.io/) is a way to plan and deploy infrastructure as code.

## Dependencies

1. Terraform
2. You need to manually create a key pair named "taquito"

## Prepare to run terraform

### Set variables

[Variables](https://www.terraform.io/docs/configuration/variables.html) serve as input parameters in terraform.

For this simple example, we'll only set 2 variables

* VPC ID
* AWS Profile (if different from "default")

Once you've logged into your AWS account console, visit the [VPCs list](https://console.aws.amazon.com/vpc/home?region=us-east-1#vpcs:) you should only have 1. 

Copy the VPC ID listed and past it into [terraform/variables.tf](terraform/variables.tf):

```
variable "vpc_id" {
    default = "PASTE VPC ID HERE"
}
```

If you named your AWS profile (locally) something different than "default", set that name in the profile variable:

```
variable profile {
    default = "default"
}
```

## Edit variables in task-definition.json (optional)

Currently, the task-definition for our container is setup to run in the **us-east-1** region. If you are running in a different region, change those references. 

# We are now ready to run Terraform

Open a terminal window and cd into the aws/terraform directory

```
cd $GOPATH/src/github.com/PenguinParadigm/samvera18apis/aws/terraform
```

## Initialize Terraform

For any terraform project, when getting started or changing any provider or module, initilizing is required:
```
terraform init
```

You should see output like:

```
Initializing provider plugins...
- Checking for available provider plugins on https://releases.hashicorp.com...
- Downloading plugin for provider "aws" (1.40.0)...

Terraform has been successfully initialized!

You may now begin working with Terraform. Try running "terraform plan" to see
any changes that are required for your infrastructure. All Terraform commands
should now work.

If you ever set or change modules or backend configuration for Terraform,
rerun this command to reinitialize your working directory. If you forget, other
commands will detect it and remind you to do so if necessary.
```

## Create an execution plan

```
terraform plan
```

Will produce a lot of output that looks like:


```
An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  + aws_ecs_cluster.taquito
      id:                                    <computed>
      arn:                                   <computed>
      name:                                  "taquito"
 ...
 ...
 ...
 
 Plan: 9 to add, 0 to change, 0 to destroy.

------------------------------------------------------------------------

Note: You didn't specify an "-out" parameter to save this plan, so Terraform
can't guarantee that exactly these actions will be performed if
"terraform apply" is subsequently run.     
```


## Apply the execution plan

```
terraform apply
```

A lot of the same input, with a required response:

```
Plan: 9 to add, 0 to change, 0 to destroy.

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value:
```

Enter "yes"

... More output:

```
aws_ecs_cluster.taquito: Creating...
  arn:  "" => "<computed>"
  name: "" => "taquito"

  aws_instance.ec2-instance: Creation complete after 44s (ID: i-0514bf0a4bc72c824)

Apply complete! Resources: 9 added, 0 changed, 0 destroyed.

Outputs:

public_endpoint = ec2-107-23-232-107.compute-1.amazonaws.com
```

Notice we specifically asked to be told about the public DNS for our instance in [terraform/outputs.tf](terraform/outputs.tf)

## Tear down

Just as before, we don't want to keep our infrastructure running, and terrform makes teardown quick and easy.

```
terraform destroy
```

Output:

```
An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  - destroy

Terraform will perform the following actions:

  - aws_ecs_cluster.taquito

  - aws_ecs_service.taquito-service

  - aws_ecs_task_definition.taquito-task

  - aws_iam_instance_profile.ec2-profile

  - aws_iam_role.ec2-role

  - aws_iam_role_policy_attachment.role-attach-1

  - aws_iam_role_policy_attachment.role-attach-2

  - aws_instance.ec2-instance

  - aws_security_group.allow_all


Plan: 0 to add, 0 to change, 9 to destroy.

Do you really want to destroy all resources?
  Terraform will destroy all your managed infrastructure, as shown above.
  There is no undo. Only 'yes' will be accepted to confirm.

   Enter a value:
```

Answer "yes"

Final Output:

```
  aws_ecs_service.taquito-service: Destroying... (ID: arn:aws:ecs:us-east-1:[YOUR USER ID]:service/taquito)
aws_instance.ec2-instance: Destroying... (ID: i-0514bf0a4bc72c824)
aws_iam_role_policy_attachment.role-attach-1: Destroying... (ID: ec2-role-20181011173248948800000003)
aws_iam_role_policy_attachment.role-attach-2: Destroying... (ID: ec2-role-20181011173248916800000002)
aws_iam_role_policy_attachment.role-attach-2: Destruction complete after 1s
aws_iam_role_policy_attachment.role-attach-1: Destruction complete after 1s
aws_instance.ec2-instance: Still destroying... (ID: i-0514bf0a4bc72c824, 10s elapsed)
aws_ecs_service.taquito-service: Still destroying... (ID: arn:aws:ecs:us-east-1:[YOUR USER ID]:service/taquito, 10s elapsed)
aws_instance.ec2-instance: Still destroying... (ID: i-0514bf0a4bc72c824, 20s elapsed)
aws_ecs_service.taquito-service: Still destroying... (ID: arn:aws:ecs:us-east-1:[YOUR USER ID]:service/taquito, 20s elapsed)
aws_ecs_service.taquito-service: Destruction complete after 28s
aws_ecs_task_definition.taquito-task: Destroying... (ID: taquito)
aws_ecs_cluster.taquito: Destroying... (ID: arn:aws:ecs:us-east-1:[YOUR USER ID]:cluster/taquito)
aws_ecs_task_definition.taquito-task: Destruction complete after 0s
aws_ecs_cluster.taquito: Destruction complete after 0s
aws_instance.ec2-instance: Still destroying... (ID: i-0514bf0a4bc72c824, 30s elapsed)
aws_instance.ec2-instance: Still destroying... (ID: i-0514bf0a4bc72c824, 40s elapsed)
aws_instance.ec2-instance: Still destroying... (ID: i-0514bf0a4bc72c824, 50s elapsed)
aws_instance.ec2-instance: Still destroying... (ID: i-0514bf0a4bc72c824, 1m0s elapsed)
aws_instance.ec2-instance: Still destroying... (ID: i-0514bf0a4bc72c824, 1m10s elapsed)
aws_instance.ec2-instance: Destruction complete after 1m12s
aws_iam_instance_profile.ec2-profile: Destroying... (ID: ec2-profile)
aws_security_group.allow_all: Destroying... (ID: sg-076f644ee15c1c395)
aws_iam_instance_profile.ec2-profile: Destruction complete after 1s
aws_iam_role.ec2-role: Destroying... (ID: ec2-role)
aws_security_group.allow_all: Destruction complete after 1s
aws_iam_role.ec2-role: Destruction complete after 1s

Destroy complete! Resources: 9 destroyed.
```