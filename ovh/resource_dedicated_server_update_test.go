package ovh

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDedicatedServerUpdate_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckCredentials(t)
			testAccPreCheckDedicatedServer(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDedicatedServerUpdateConfig(false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("test", "true"),
					resource.TestCheckResourceAttr(
						"ovh_dedicated_server_update.server", "monitoring", "true"),
					resource.TestCheckResourceAttr(
						"ovh_dedicated_server_update.server", "state", "ok"),
					resource.TestCheckResourceAttr(
						"ovh_dedicated_server_update.server", "boot_script", ""),
					resource.TestCheckResourceAttr(
						"ovh_dedicated_server_update.server", "efi_bootloader_path", "\\efi\\boot\\bootx64.efi"),
					resource.TestCheckResourceAttr(
						"ovh_dedicated_server_update.server", "display_name", "An awesome display name"),
				),
			},
			{
				Config: testAccDedicatedServerUpdateConfig(true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("test", "true"),
					resource.TestCheckResourceAttr(
						"ovh_dedicated_server_update.server", "monitoring", "false"),
					resource.TestCheckResourceAttr(
						"ovh_dedicated_server_update.server", "state", "ok"),
					resource.TestCheckResourceAttr(
						"ovh_dedicated_server_update.server", "boot_script", ""),
					resource.TestCheckResourceAttr(
						"ovh_dedicated_server_update.server", "efi_bootloader_path", "\\efi\\boot\\bootx64.efi"),
					resource.TestCheckResourceAttr(
						"ovh_dedicated_server_update.server", "display_name", "An awesome display name restored"),
				),
			},
		},
	})
}

func testAccDedicatedServerUpdateConfig(restore bool) string {
	dedicated_server := os.Getenv("OVH_DEDICATED_SERVER")
	if restore {
		return fmt.Sprintf(
			testAccDedicatedServerUpdateConfig_Restore,
			dedicated_server,
		)
	}
	return fmt.Sprintf(
		testAccDedicatedServerUpdateConfig_Basic,
		dedicated_server,
	)

}

const testAccDedicatedServerUpdateConfig_Basic = `
data ovh_dedicated_server_boots "rescue" {
  service_name = "%s"
  boot_type    = "rescue"
  kernel       = "rescue-customer"
}

resource ovh_dedicated_server_update "server" {
  service_name = data.ovh_dedicated_server_boots.rescue.service_name
  boot_id      = data.ovh_dedicated_server_boots.rescue.result[0]
  monitoring   = true
  state        = "ok"
  display_name = "An awesome display name"
  efi_bootloader_path = "\\efi\\boot\\bootx64.efi"
}

output test {
  value = tostring(data.ovh_dedicated_server_boots.rescue.result[0] == ovh_dedicated_server_update.server.boot_id )
}
`

const testAccDedicatedServerUpdateConfig_Restore = `
data ovh_dedicated_server_boots "harddisk" {
  service_name = "%s"
  boot_type    = "harddisk"
}

resource ovh_dedicated_server_update "server" {
  service_name = data.ovh_dedicated_server_boots.harddisk.service_name
  boot_id      = data.ovh_dedicated_server_boots.harddisk.result[0]
  monitoring   = false
  state        = "ok"
  display_name = "An awesome display name restored"
  efi_bootloader_path = "\\efi\\boot\\bootx64.efi"
}

output test {
  value = tostring(data.ovh_dedicated_server_boots.harddisk.result[0] == ovh_dedicated_server_update.server.boot_id )
}
`
