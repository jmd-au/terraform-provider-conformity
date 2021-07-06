package conformity

import (
	"context"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAwsAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAwsAccountCreate,
		ReadContext:   resourceAwsAccountRead,
		UpdateContext: resourceAwsAccountUpdate,
		DeleteContext: resourceAccountDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"environment": {
				Type:     schema.TypeString,
				Required: true,
			},
			"role_arn": {
				Type:     schema.TypeString,
				Required: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAwsAccountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	payload := cloudconformity.AccountPayload{}
	payload.Data.Attributes.Name = d.Get("name").(string)
	payload.Data.Attributes.Environment = d.Get("environment").(string)
	payload.Data.Attributes.Access.Keys.RoleArn = d.Get("role_arn").(string)
	payload.Data.Attributes.Access.Keys.ExternalId = d.Get("external_id").(string)
	accountId, err := client.CreateAwsAccount(payload)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(accountId)
	resourceAwsAccountRead(ctx, d, m)
	return diags
}

func resourceAwsAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	accountId := d.Id()

	// get both account details and access settings
	accountAccessAndDetails, err := client.GetAccount(accountId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", accountAccessAndDetails.AccountDetails.Data.Attributes.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("environment", accountAccessAndDetails.AccountDetails.Data.Attributes.Environment); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("role_arn", accountAccessAndDetails.AccessSettings.Attributes.Configuration.RoleArn); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("external_id", accountAccessAndDetails.AccessSettings.Attributes.Configuration.ExternalId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("tags", accountAccessAndDetails.AccountDetails.Data.Attributes.Tags); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceAwsAccountUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	if d.HasChange("name") || d.HasChange("environment") || d.HasChange("tags") {

		payload := cloudconformity.AccountPayload{}
		payload.Data.Attributes.Name = d.Get("name").(string)
		payload.Data.Attributes.Environment = d.Get("environment").(string)

		tags := d.Get("tags").(*schema.Set)
		for _, tag := range tags.List() {
			payload.Data.Attributes.Tags = append(payload.Data.Attributes.Tags, tag.(string))
		}

		accountId := d.Id()
		_, err := client.UpdateAccount(accountId, payload)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	if d.HasChange("role_arn") || d.HasChange("external_id") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to Update Conformity AWS account",
			Detail:   "'role_arn' and 'external_id' cannot be changed",
		})

		return diags
	}

	return resourceAwsAccountRead(ctx, d, m)
}

func resourceAccountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	accountId := d.Id()

	_, err := client.DeleteAccount(accountId)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.

	d.SetId("")
	return diags
}