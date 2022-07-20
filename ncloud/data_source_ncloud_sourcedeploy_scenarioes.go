package ncloud

import (
	"context"

	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/vsourcedeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	RegisterDataSource("ncloud_sourcedeploy_scenarioes", dataSourceNcloudSourceDeployscenarioesContext())
}

func dataSourceNcloudSourceDeployscenarioesContext() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNcloudSourceDeployScenarioesReadContext,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"stage_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"output_file": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"filter": dataSourceFiltersSchema(),
			"scenarioes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNcloudSourceDeployScenarioesReadContext(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*ProviderConfig)

	if !config.SupportVPC {
		return diag.FromErr(NotSupportClassic("dataSource `ncloud_sourcedeploy_scenarioes`"))
	}

	projectId := ncloud.IntString(d.Get("project_id").(int))
	stageId := ncloud.IntString(d.Get("stage_id").(int))
	resp, err := GetScenarioes(ctx, config, projectId, stageId)
	if err != nil {
		return diag.FromErr(err)
	}
	logResponse("GetScenarioesList", resp)

	resources := []map[string]interface{}{}
	for _, r := range resp.ScenarioList {
		project := map[string]interface{}{
			"id":   *r.Id,
			"name": *r.Name,
		}

		resources = append(resources, project)
	}

	if f, ok := d.GetOk("filter"); ok {
		resources = ApplyFilters(f.(*schema.Set), resources, dataSourceNcloudSourceDeployscenarioesContext().Schema)
	}
	d.SetId(config.RegionCode)
	d.Set("scenarioes", resources)

	if output, ok := d.GetOk("output_file"); ok && output.(string) != "" {
		return diag.FromErr(writeToFile(output.(string), resources))
	}

	return nil
}

func GetScenarioes(ctx context.Context, config *ProviderConfig, projectId *string, stageId *string)(*vsourcedeploy.GetScenarioListResponse, error) {
	
	reqParams := make(map[string]interface{})
	logCommonRequest("GetScenarioes", reqParams)
	resp, err := config.Client.vsourcedeploy.V1Api.GetScenarioes(ctx, projectId, stageId ,reqParams)

	if err != nil {
		logErrorResponse("GetScenarioes", err, "")
		return nil, err
	}
	logResponse("GetScenarioes", resp)

	return resp, nil
}
