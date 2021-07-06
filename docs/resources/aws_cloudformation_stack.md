---
page_title: "aws_cloudformation_stack Resource - cloudconformity_terraform"
subcategory: "AWS"
description: |-
  Provides a CloudFormation Stack resource from AWS provider.
---

# Resource `aws_cloudformation_stack`
Provides a CloudFormation Stack resource from AWS provider.

## Example Usage
```terraform
resource "aws_cloudformation_stack" "cloud-conformity" {

  name         = "CloudConformity"

  template_url = "https://s3-us-west-2.amazonaws.com/cloudconformity/CloudConformity.template"

  parameters={
    AccountId  = "717210094962"
    ExternalId = "${data.conformity_external_id.all.external_id}"
  }

  capabilities = ["CAPABILITY_NAMED_IAM"]

}
```

## Argument reference

 - `name` - (Required) The name of your CloudFormation Stack (Do not change the value).
 - `template_url` - (Required) Default CloudFormation template (Do not change the value).
 - `AccountId` - (Required) Default Cloud Conformity AWS Account ID (Do not change the value).

## Attributes Reference

In addition to all the arguments above, the following attributes are imported from `cloudconformity_external_id` resource.

 - `ExternalId` - (Required) The external ID. Imported from `cloudconformity_external_id`.

## Import
Cloudformation Stacks can be imported using the `name`. e.g.

```hcl
terraform import aws_cloudformation_stack.stack CloudConformity
```

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 2.7.0 |
| <a name="requirement_conformity"></a> [conformity](#requirement\_conformity) | 0.1.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | >= 2.7.0 |
| <a name="provider_conformity"></a> [conformity](#provider\_conformity) | 0.1.0 |

## Resources

| Name | Type |
|------|------|
| [aws_cloudformation_stack.cloud-conformity](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudformation_stack) | resource |

| conformity_aws_account.aws | resource |

| conformity_external_id.external | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_access_key"></a> [access\_key](#input\_access\_key) | n/a | `string` | `""` | yes |
| <a name="input_apikey"></a> [apikey](#input\_apikey) | n/a | `string` | `""` | yes |
| <a name="input_environment"></a> [environment](#input\_environment) | n/a | `string` | `"Staging"` | yes |
| <a name="input_name"></a> [name](#input\_name) | n/a | `string` | `"Cloudconformity"` | yes |
| <a name="input_region"></a> [region](#input\_region) | n/a | `string` | `"us-west-2"` | yes |
| <a name="input_secret_key"></a> [secret\_key](#input\_secret\_key) | n/a | `string` | `""` | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_aws_account_name"></a> [aws\_account\_name](#output\_aws\_account\_name) | n/a |
| <a name="output_aws_environment"></a> [aws\_environment](#output\_aws\_environment) | n/a |
| <a name="output_aws_role_arn"></a> [aws\_role\_arn](#output\_aws\_role\_arn) | n/a |
| <a name="output_external_id"></a> [external\_id](#output\_external\_id) | n/a |
| <a name="output_role_arn"></a> [role\_arn](#output\_role\_arn) | n/a |