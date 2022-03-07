package provider

import (
	"context"

	"github.com/firehydrant/terraform-provider-firehydrant/firehydrant"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceServiceDependency() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataFireHydrantServiceDependency,
		Schema: map[string]*schema.Schema{
			"service_dependency_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"connected_service_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"notes": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataFireHydrantServiceDependency(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(firehydrant.Client)
	id := d.Get("service_dependency_id").(string)

	r, err := ac.GetServiceDependency(ctx, id)
	if err != nil {
		return diag.FromErr(err)
	}

	var ds diag.Diagnostics
	env := map[string]string{
		"service_id":           r.ServiceID,
		"connected_service_id": r.ConnectedServiceID,
		"notes":                r.Notes,
	}

	for key, val := range env {
		if err := d.Set(key, val); err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(r.ID)

	return ds
}