# How to run terraform

```
terraform plan -var tenable_linking_key=foo -out terraform.plan
```

check the variables.tf on what else you can set or override with tfvars

then you can do terraform apply on the plan file

```
terraform apply terraform.plan
```
