package rdns_test

import (
	"fmt"
	"github.com/terraform-providers/terraform-provider-hcloud/internal/floatingip"
	"github.com/terraform-providers/terraform-provider-hcloud/internal/rdns"
	"github.com/terraform-providers/terraform-provider-hcloud/internal/server"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/terraform-providers/terraform-provider-hcloud/internal/testsupport"
	"github.com/terraform-providers/terraform-provider-hcloud/internal/testtemplate"
)

func TestRDNSResource_Server(t *testing.T) {
	var s hcloud.Server

	tmplMan := testtemplate.Manager{}
	resServer := &server.RData{
		Name:  "server-rdns",
		Type:  "cx11",
		Image: "ubuntu-20.04",
		Labels: map[string]string{
			"tf-test": fmt.Sprintf("tf-test-rdns-%d", tmplMan.RandInt),
		},
	}
	resServer.SetRName("server_rdns")
	resource.Test(t, resource.TestCase{
		PreCheck:     testsupport.AccTestPreCheck(t),
		Providers:    testsupport.AccTestProviders(),
		CheckDestroy: testsupport.CheckResourcesDestroyed(server.ResourceType, server.ByID(t, &s)),
		Steps: []resource.TestStep{
			{
				// Create a new RDNS using the required values
				// only.
				Config: tmplMan.Render(t,
					"testdata/r/hcloud_server", resServer,
					"testdata/r/hcloud_rdns", rdns.NewRData(t, "rdnstest", resServer.TFID()+".id", "", resServer.TFID()+".ipv4_address", "example.hetzner.cloud"),
				),
				Check: resource.ComposeTestCheckFunc(
					testsupport.CheckResourceExists(resServer.TFID(), server.ByID(t, &s)),
					resource.TestCheckResourceAttr("hcloud_rdns.rdnstest", "dns_ptr", "example.hetzner.cloud"),
				),
			},
		},
	})
}

func TestRDNSResource_FloatingIP_IPv4(t *testing.T) {
	var fl hcloud.FloatingIP

	tmplMan := testtemplate.Manager{}
	restFloatingIP := &floatingip.RData{
		Name:             "floating-ipv4-rdns",
		Type:             "ipv4",
		HomeLocationName: "fsn1",
	}
	restFloatingIP.SetRName("floating_ips_rdns_v4")
	resource.Test(t, resource.TestCase{
		PreCheck:     testsupport.AccTestPreCheck(t),
		Providers:    testsupport.AccTestProviders(),
		CheckDestroy: testsupport.CheckResourcesDestroyed(floatingip.ResourceType, floatingip.ByID(t, &fl)),
		Steps: []resource.TestStep{
			{
				// Create a new SSH Key using the required values
				// only.
				Config: tmplMan.Render(t,
					"testdata/r/hcloud_floating_ip", restFloatingIP,
					"testdata/r/hcloud_rdns", rdns.NewRData(t, "floating_ips_rdns_v4", "", restFloatingIP.TFID()+".id", restFloatingIP.TFID()+".ip_address", "example.hetzner.cloud"),
				),
				Check: resource.ComposeTestCheckFunc(
					testsupport.CheckResourceExists(restFloatingIP.TFID(), floatingip.ByID(t, &fl)),
					resource.TestCheckResourceAttr("hcloud_rdns.floating_ips_rdns_v4", "dns_ptr", "example.hetzner.cloud"),
				),
			},
		},
	})
}

func TestRDNSResource_FloatingIP_IPv6(t *testing.T) {
	var fl hcloud.FloatingIP

	tmplMan := testtemplate.Manager{}
	restFloatingIP := &floatingip.RData{
		Name:             "floating-ipv6-rdns",
		Type:             "ipv6",
		HomeLocationName: "fsn1",
	}
	restFloatingIP.SetRName("floating_ips_rdns_v6")
	resource.Test(t, resource.TestCase{
		PreCheck:     testsupport.AccTestPreCheck(t),
		Providers:    testsupport.AccTestProviders(),
		CheckDestroy: testsupport.CheckResourcesDestroyed(floatingip.ResourceType, floatingip.ByID(t, &fl)),
		Steps: []resource.TestStep{
			{
				// Create a new SSH Key using the required values
				// only.
				Config: tmplMan.Render(t,
					"testdata/r/hcloud_floating_ip", restFloatingIP,
					"testdata/r/hcloud_rdns", rdns.NewRData(t, "floating_ips_rdns_v6", "", restFloatingIP.TFID()+".id", restFloatingIP.TFID()+".ip_address", "example.hetzner.cloud"),
				),
				Check: resource.ComposeTestCheckFunc(
					testsupport.CheckResourceExists(restFloatingIP.TFID(), floatingip.ByID(t, &fl)),
					resource.TestCheckResourceAttr("hcloud_rdns.floating_ips_rdns_v6", "dns_ptr", "example.hetzner.cloud"),
				),
			},
		},
	})
}
