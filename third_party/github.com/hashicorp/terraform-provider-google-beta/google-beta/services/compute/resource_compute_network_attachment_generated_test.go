// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package compute_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccComputeNetworkAttachment_networkAttachmentBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":          envvar.GetTestOrgFromEnv(t),
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeNetworkAttachmentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNetworkAttachment_networkAttachmentBasicExample(context),
			},
			{
				ResourceName:            "google_compute_network_attachment.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccComputeNetworkAttachment_networkAttachmentBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network_attachment" "default" {
    provider = google-beta
    name = "tf-test-basic-network-attachment%{random_suffix}"
    region = "us-central1"
    description = "basic network attachment description"
    connection_preference = "ACCEPT_MANUAL"

    subnetworks = [
        google_compute_subnetwork.default.self_link
    ]

    producer_accept_lists = [
        google_project.accepted_producer_project.project_id
    ]

    producer_reject_lists = [
        google_project.rejected_producer_project.project_id
    ]
}

resource "google_compute_network" "default" {
    provider = google-beta
    name = "tf-test-basic-network%{random_suffix}"
    auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
    provider = google-beta
    name = "tf-test-basic-subnetwork%{random_suffix}"
    region = "us-central1"

    network = google_compute_network.default.id
    ip_cidr_range = "10.0.0.0/16"
}

resource "google_project" "rejected_producer_project" {
    provider = google-beta
    project_id      = "prj-rejected%{random_suffix}"
    name            = "prj-rejected%{random_suffix}"
    org_id          = "%{org_id}"
    billing_account = "%{billing_account}"
}

resource "google_project" "accepted_producer_project" {
    provider = google-beta
    project_id      = "prj-accepted%{random_suffix}"
    name            = "prj-accepted%{random_suffix}"
    org_id          = "%{org_id}"
    billing_account = "%{billing_account}"
}
`, context)
}

func testAccCheckComputeNetworkAttachmentDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_network_attachment" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/networkAttachments/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("ComputeNetworkAttachment still exists at %s", url)
			}
		}

		return nil
	}
}