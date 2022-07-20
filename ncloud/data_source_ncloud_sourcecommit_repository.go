package ncloud

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	RegisterDataSource("ncloud_sourcecommit_repository", dataSourceNcloudSourceCommitRepository())
}

func dataSourceNcloudSourceCommitRepository() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSousrceNcloudSourceCommitRepositoryRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creator": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"git_https": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"git_ssh": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"filesafer": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSousrceNcloudSourceCommitRepositoryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	config := meta.(*ProviderConfig)

	name := d.Get("name").(string)

	logCommonRequest("GetSourceCommitRepository", "")
	repository, err := getRepository(ctx, config, name)

	var diags diag.Diagnostics

	if err != nil {
		logErrorResponse("GetSourceCommitRepository", err, "")
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to search repository",
			Detail:   "Unable to search repository - detail",
		})
		return diags
	}

	if repository == nil {
		logErrorResponse("GetSourceCommitRepository", err, "")
		d.SetId("")
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "there is no such repository",
			Detail:   "there is no such repository - detail",
		})
		return diags
	}

	logResponse("GetSourceCommitRepository", repository)
	d.SetId(strconv.Itoa(*repository.Id))
	d.Set("name", repository.Name)
	d.Set("description", repository.Description)
	d.Set("creator", repository.Created.User)
	d.Set("git_https", repository.Git.Https)
	d.Set("git_ssh", repository.Git.Ssh)
	d.Set("filesafer", repository.Linked.FileSafer)

	return nil
}
