package ncloud

import (
	"os"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NCLOUD_ACCESS_KEY", os.Getenv("NCLOUD_ACCESS_KEY")),
				Description: descriptions["access_key"],
			},
			"secret_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NCLOUD_SECRET_KEY", os.Getenv("NCLOUD_SECRET_KEY")),
				Description: descriptions["secret_key"],
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NCLOUD_REGION", os.Getenv("NCLOUD_REGION")),
				Description: descriptions["region"],
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"ncloud_regions":               dataSourceNcloudRegions(),
			"ncloud_zones":                 dataSourceNcloudZones(),
			"ncloud_server_image":          dataSourceNcloudServerImage(),
			"ncloud_server_images":         dataSourceNcloudServerImages(),
			"ncloud_member_server_image":   dataSourceNcloudMemberServerImage(),
			"ncloud_member_server_images":  dataSourceNcloudMemberServerImages(),
			"ncloud_server_product":        dataSourceNcloudServerProduct(),
			"ncloud_server_products":       dataSourceNcloudServerProducts(),
			"ncloud_port_forwarding_rule":  dataSourceNcloudPortForwardingRule(),
			"ncloud_port_forwarding_rules": dataSourceNcloudPortForwardingRules(),
			"ncloud_nas_volume":            dataSourceNcloudNasVolume(),
			"ncloud_nas_volumes":           dataSourceNcloudNasVolumes(),
			"ncloud_access_control_group":  dataSourceNcloudAccessControlGroup(),
			"ncloud_access_control_groups": dataSourceNcloudAccessControlGroups(),
			"ncloud_access_control_rule":   dataSourceNcloudAccessControlRule(),
			"ncloud_access_control_rules":  dataSourceNcloudAccessControlRules(),
			"ncloud_root_password":         dataSourceNcloudRootPassword(),
			"ncloud_public_ip":             dataSourceNcloudPublicIP(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"ncloud_instance":                      resourceNcloudInstance(),
			"ncloud_block_storage":                 resourceNcloudBlockStorage(),
			"ncloud_public_ip":                     resourceNcloudPublicIPInstance(),
			"ncloud_login_key":                     resourceNcloudLoginKey(),
			"ncloud_nas_volume":                    resourceNcloudNasVolume(),
			"ncloud_port_forwarding_rule":          resourceNcloudPortForwadingRule(),
			"ncloud_load_balancer":                 resourceNcloudLoadBalancer(),
			"ncloud_load_balancer_ssl_certificate": resourceNcloudLoadBalancerSSLCertificate(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AccessKey: d.Get("access_key").(string),
		SecretKey: d.Get("secret_key").(string),
	}

	if region, ok := d.GetOk("region"); ok && os.Getenv("NCLOUD_REGION") == "" {
		os.Setenv("NCLOUD_REGION", region.(string))
	}

	sdk, err := config.Client()
	if err != nil {
		return nil, err
	}
	return sdk, nil
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"access_key": "Access key of ncloud",
		"secret_key": "Secret key of ncloud",
		"region":     "Region of ncloud",
	}
}
