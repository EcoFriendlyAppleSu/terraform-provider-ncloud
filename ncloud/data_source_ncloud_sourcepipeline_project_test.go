package ncloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNcloudSourcePipelineProject_classic_basic(t *testing.T) {
	dataName := "data.ncloud_sourcepipeline_project.foo"
	resourceName := "ncloud_sourcepipeline_project.test-project"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccClassicProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNcloudSourcePipelineProjectConfig("test-project", "description test"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataSourceID(dataName),
					resource.TestCheckResourceAttrPair(dataName, "id", resourceName, "id"),
					resource.TestCheckResourceAttrPair(dataName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataName, "description", resourceName, "description"),
				),
			},
		},
	})
}

func TestAccDataSourceNcloudSourcePipelineProject_vpc_basic(t *testing.T) {
	dataName := "data.ncloud_sourcepipeline_project.foo"
	resourceName := "ncloud_sourcepipeline_project.test-project"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNcloudSourcePipelineProjectConfig("test-project", "description test"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataSourceID(dataName),
					resource.TestCheckResourceAttrPair(dataName, "id", resourceName, "id"),
					resource.TestCheckResourceAttrPair(dataName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataName, "description", resourceName, "description"),
				),
			},
		},
	})
}

func testAccDataSourceNcloudSourcePipelineProjectConfig(name, description string) string {
	return fmt.Sprintf(`
data "ncloud_sourcebuild_compute" "compute" {
}

data "ncloud_sourcebuild_os" "os" {
}

data "ncloud_sourcebuild_runtime" "runtime" {
	os_id = data.ncloud_sourcebuild_os.os.os[0].id
}

data "ncloud_sourcebuild_runtime_version" "runtime_version" {
	os_id      = data.ncloud_sourcebuild_os.os.os[0].id
	runtime_id = data.ncloud_sourcebuild_runtime.runtime.runtime[0].id
}

data "ncloud_sourcebuild_docker" "docker" {
}
		
resource "ncloud_sourcecommit_repository" "test-repo" {
	name = "sourceCommit"
}

resource "ncloud_sourcebuild_project" "test-project" {
	name        = "souceBuild"
	description = "my build project"
	source {
		type = "SourceCommit"
		config {
			repository = ncloud_sourcecommit_repository.test-repo.name
			branch     = "master"
		}
	}
	env {
		compute {
			id = data.ncloud_sourcebuild_compute.compute.compute[0].id
		}
		platform {
			type = "SourceBuild"
			config {
				os {
					id = data.ncloud_sourcebuild_os.os.os[0].id
				}
				runtime {
					id = data.ncloud_sourcebuild_runtime.runtime.runtime[0].id
					version {
						id = data.ncloud_sourcebuild_runtime_version.runtime_version.runtime_version[0].id
					}
				}
			}
		}
		docker {
			use = true
			id = data.ncloud_sourcebuild_docker.docker.docker[0].id
		}
		timeout = 500
		env_vars {
			key   = "k1"
			value = "v1"
		}
	}
	cmd {
		pre   = ["pwd", "ls"]
		build = ["pwd", "ls"]
		post  = ["pwd", "ls"]
	}
}

resource "ncloud_sourcepipeline_project" "test-project" {
	name               = "%[1]s"
	description        = "%[2]s"
	tasks {
		name 		   = "task_name"
		type 		   = "SourceBuild"
		config {
			project_id   = ncloud_sourcebuild_project.test-project.id
		}
		linked_tasks   = []
		}
	trigger {
		setting = true
		sourcecommit {
			repository = ncloud_sourcecommit_repository.test-repo.name
		}
	}
}

data "ncloud_sourcepipeline_project" "foo" {
	id = ncloud_sourcepipeline_project.test-project.id
}
`, name, description)
}
