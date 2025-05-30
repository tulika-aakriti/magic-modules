package networksecurity_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccNetworkSecurityMirroringDeploymentGroup_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityMirroringDeploymentGroup_basic(context),
			},
			{
				ResourceName:            "google_network_security_mirroring_deployment_group.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
			{
				Config: testAccNetworkSecurityMirroringDeploymentGroup_update(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_network_security_mirroring_deployment_group.default", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_network_security_mirroring_deployment_group.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"update_time", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkSecurityMirroringDeploymentGroup_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "network" {
  name                    = "tf-test-example-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_network_security_mirroring_deployment_group" "default" {
  mirroring_deployment_group_id = "tf-test-example-dg%{random_suffix}"
  location                      = "global"
  network                       = google_compute_network.network.id
  description                   = "initial description"
  labels = {
    foo = "bar"
  }
}
`, context)
}

func testAccNetworkSecurityMirroringDeploymentGroup_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "network" {
  name                    = "tf-test-example-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_network_security_mirroring_deployment_group" "default" {
  mirroring_deployment_group_id = "tf-test-example-dg%{random_suffix}"
  location                      = "global"
  network                       = google_compute_network.network.id
  description                   = "updated description"
  labels = {
    foo = "goo"
  }
}
`, context)
}
