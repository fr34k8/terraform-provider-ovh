## 2.6.0 (July 25, 2025)

⚙️ General:

* Bump dependencies ([#1052](https://github.com/ovh/terraform-provider-ovh/pull/1052))

🎉 Features:

* New resource: `ovh_cloud_project_region` ([#1044](https://github.com/ovh/terraform-provider-ovh/pull/1044))

* New datasource: `ovh_cloud_project_gateway` ([#1061](https://github.com/ovh/terraform-provider-ovh/pull/1061))
* New datasource: `ovh_location` ([#1063](https://github.com/ovh/terraform-provider-ovh/pull/1063))

💪 Improvements:

* `r/ovh_dbaas_logs_token`: Resource can now be imported ([#1045](https://github.com/ovh/terraform-provider-ovh/pull/1045))
* Ensure all resources have an ID ([#1047](https://github.com/ovh/terraform-provider-ovh/pull/1047))

⚠️ Deprecations:

* `r/ovh_dedicated_server_update`: Resource is now deprecated ([#1036](https://github.com/ovh/terraform-provider-ovh/pull/1036))

🐜 Bug fixes:

* `r/ovh_dedicated_server`: Changing `service_name` forces recreation of the resource ([#1033](https://github.com/ovh/terraform-provider-ovh/pull/1033))
* `r/ovh_dedicated_server_update`: Changing `service_name` forces recreation of the resource ([#1033](https://github.com/ovh/terraform-provider-ovh/pull/1033))
* `r/ovh_savings_plan`: Perform an exact match when searching for a product code ([#1040](https://github.com/ovh/terraform-provider-ovh/pull/1040))
* `r/ovh_dbaas_logs_input`: Fix fields in example ([#1053](https://github.com/ovh/terraform-provider-ovh/pull/1053))
* `r/ovh_cloud_project_failover_ip_attach`: Correctly check errors and improve error messages ([#1056](https://github.com/ovh/terraform-provider-ovh/pull/1056))
* `r/ovh_cloud_project_containerregistry_oidc`: Attribute `oidc_verify_cert` is computed when not input by the user ([#1057](https://github.com/ovh/terraform-provider-ovh/pull/1057))
* `r/ovh_cloud_project_containerregistry_oidc`: Recreate OIDC config when removed outside Terraform ([#1058](https://github.com/ovh/terraform-provider-ovh/pull/1058))
* `r/ovh_cloud_project_storage`: Set the right type for replication tags ([#1060](https://github.com/ovh/terraform-provider-ovh/pull/1060))

📚 Documentation:

* Update subcategories of several public cloud resources ([#1035](https://github.com/ovh/terraform-provider-ovh/pull/1035))
* `r/ovh_cloud_project_containerregistry_user`: Fix resource documentation ([#1037](https://github.com/ovh/terraform-provider-ovh/pull/1037))
* Ensure naming convention in examples ([#1046](https://github.com/ovh/terraform-provider-ovh/pull/1046))
* `r/ovh_vps`: Add clarifications on image selection ([#1055](https://github.com/ovh/terraform-provider-ovh/pull/1055))

❤️ Thanks for your contributions ❤️

## 2.5.0 (July 1, 2025)

⚙️ General:

* Bump dependencies ([#1010](https://github.com/ovh/terraform-provider-ovh/pull/1010))

🎉 Features:

* New resource: `ovh_domain_zone_dynhost_login` ([#1001](https://github.com/ovh/terraform-provider-ovh/pull/1001))
* New resource: `ovh_domain_zone_dynhost_record` ([#1001](https://github.com/ovh/terraform-provider-ovh/pull/1001))
* New resource: `ovh_storage_efs_share` ([#1003](https://github.com/ovh/terraform-provider-ovh/pull/1003))
* New resource: `ovh_storage_efs_share_acl` ([#1003](https://github.com/ovh/terraform-provider-ovh/pull/1003))
* New resource: `ovh_storage_efs_share_snapshot` ([#1003](https://github.com/ovh/terraform-provider-ovh/pull/1003))

* New datasource: `ovh_storage_efs` ([#1003](https://github.com/ovh/terraform-provider-ovh/pull/1003))
* New datasource: `ovh_storage_efs_share_access_path` ([#1003](https://github.com/ovh/terraform-provider-ovh/pull/1003))
* New datasource: `ovh_storage_efs_share_access_paths` ([#1003](https://github.com/ovh/terraform-provider-ovh/pull/1003))

💪 Improvements:

* `r/ovh_ip_reverse`: Add field `readiness_timeout_duration` ([#1021](https://github.com/ovh/terraform-provider-ovh/pull/1021))
* `r/ovh_cloud_project_network_private`: Add field `regions_openstack_ids` ([#1011](https://github.com/ovh/terraform-provider-ovh/pull/1011))
* `r/ovh_vps`: Add fields `public_ssh_key` and `image_id` to allow installation of the server ([#1015](https://github.com/ovh/terraform-provider-ovh/pull/1015))

* `d/ovh_cloud_project_network_private`: Add field `regions_openstack_ids` ([#1011](https://github.com/ovh/terraform-provider-ovh/pull/1011))

🐜 Bug fixes:

* Retry failing calls when fetching order plans ([#1018](https://github.com/ovh/terraform-provider-ovh/pull/1018))
* `r/ovh_dedicated_server`: Correctly send string parameters when updating the resource ([#1002](https://github.com/ovh/terraform-provider-ovh/pull/1002))

📚 Documentation:

* Licenses update ([#1008](https://github.com/ovh/terraform-provider-ovh/pull/1008))
* Fix contributing doc ([#1012](https://github.com/ovh/terraform-provider-ovh/pull/1012))
* `r/ovh_me_identity_group`: Add import documentation ([#1024](https://github.com/ovh/terraform-provider-ovh/pull/1024))

❤️ Thanks for your contributions ❤️

## 2.4.0 (June 6, 2025)

⚙️ General:

* Add global option `api_rate_limit` to limit the number of concurrent calls to OVHcloud API ([#979](https://github.com/ovh/terraform-provider-ovh/pull/979))

🎉 Features:

* New resource: `ovh_cloud_project_containerregistry_iam` ([#907](https://github.com/ovh/terraform-provider-ovh/pull/907))
* New resource: `ovh_cloud_project_database_valkey_user` ([#994](https://github.com/ovh/terraform-provider-ovh/pull/994))

* New datasource: `ovh_cloud_project_database_valkey_user` ([#994](https://github.com/ovh/terraform-provider-ovh/pull/994))

💪 Improvements:

* `r/ovh_cloud_project_containerregistry`: Add `iam_enabled` field ([#907](https://github.com/ovh/terraform-provider-ovh/pull/907))
* `r/ovh_cloud_project_database`: Add `valkey` engine ([#994](https://github.com/ovh/terraform-provider-ovh/pull/994))
* Add ability to configure the creation timeout for resources that are created via an order ([#1004](https://github.com/ovh/terraform-provider-ovh/pull/1004))

* `d/ovh_cloud_project_containerregistry`: Add `iam_enabled` field ([#907](https://github.com/ovh/terraform-provider-ovh/pull/907))
* `d/ovh_cloud_project_containerregistries`: Add `iam_enabled` field ([#907](https://github.com/ovh/terraform-provider-ovh/pull/907))


📚 Documentation:

* `r/ovh_cloud_project_kube_nodepool`: Add missing field `availability_zones` ([#989](https://github.com/ovh/terraform-provider-ovh/pull/989))

❤️ Thanks for your contributions ❤️

## 2.3.0 (May 22, 2025)

⚙️ General:

* Add `user_agent_extra` field in provider configuration to append extra information in the user-agent header in API calls ([#973](https://github.com/ovh/terraform-provider-ovh/pull/973))
* Use `go-1.23` worker model in CI ([#974](https://github.com/ovh/terraform-provider-ovh/pull/974))

🎉 Features:

* New resource: `ovh_ovhcloud_connect_pop_config` ([#970](https://github.com/ovh/terraform-provider-ovh/pull/970))
* New resource: `ovh_ovhcloud_connect_pop_datacenter_config` ([#970](https://github.com/ovh/terraform-provider-ovh/pull/970))
* New resource: `ovh_ovhcloud_connect_pop_datacenter_extra_config` ([#970](https://github.com/ovh/terraform-provider-ovh/pull/970))

* New datasource: `ovh_cloud_project_ssh_key` ([#972](https://github.com/ovh/terraform-provider-ovh/pull/972))
* New datasource: `ovh_cloud_project_ssh_keys` ([#972](https://github.com/ovh/terraform-provider-ovh/pull/972))
* New datasource: `ovh_cloud_project_rancher_capabilities_plan` ([#971](https://github.com/ovh/terraform-provider-ovh/pull/971))
* New datasource: `ovh_cloud_project_rancher_capabilities_version` ([#971](https://github.com/ovh/terraform-provider-ovh/pull/971))
* New datasource: `ovh_cloud_project_flavor` ([#969](https://github.com/ovh/terraform-provider-ovh/pull/969))
* New datasource: `ovh_ovhcloud_connect_datacenters` ([#942](https://github.com/ovh/terraform-provider-ovh/pull/942))
* New datasource: `ovh_ovhcloud_connect_config_pops` ([#946](https://github.com/ovh/terraform-provider-ovh/pull/946))
* New datasource: `ovh_ovhcloud_connect_config_pop_datacenters` ([#946](https://github.com/ovh/terraform-provider-ovh/pull/946))
* New datasource: `ovh_ovhcloud_connect_config_pop_datacenter_extras` ([#946](https://github.com/ovh/terraform-provider-ovh/pull/946))

💪 Improvements:

* `r/ovh_cloud_project_user`: Add support for password rotation ([#965](https://github.com/ovh/terraform-provider-ovh/pull/965))
* `r/ovh_cloud_project_instance`: Add customizable creation timeout ([#964](https://github.com/ovh/terraform-provider-ovh/pull/964))
* `r/ovh_cloud_project_region`: Add missing fields in resource ([#967](https://github.com/ovh/terraform-provider-ovh/pull/967))
* `r/ovh_cloud_project_flavors`: Add `name_filter` field to filter results by name ([#969](https://github.com/ovh/terraform-provider-ovh/pull/969))
* `r/ovh_vrack_ip`: Add `region` optional field ([#982](https://github.com/ovh/terraform-provider-ovh/pull/982))
* `r/ovh_vrack_ipv6`: Add `region` computed field ([#982](https://github.com/ovh/terraform-provider-ovh/pull/982))
* `r/ovh_savings_plan`: Use environment variable `OVH_CLOUD_PROJECT_SERVICE` as defaut value for `service_name` field ([#984](https://github.com/ovh/terraform-provider-ovh/pull/984))
* `r/ovh_dedicated_server`: Add `reboot_rescue` option in pre-destroy available actions ([#988](https://github.com/ovh/terraform-provider-ovh/pull/988))

🐜 Bug fixes:

* `r/ovh_cloud_project_instance`: Use the right API parameter for `floating_ip` field ([#987](https://github.com/ovh/terraform-provider-ovh/pull/987))
* `r/ovh_cloud_project_database_log_subscription`: Add missing required field `kind` ([#976](https://github.com/ovh/terraform-provider-ovh/pull/976))
* `r/ovh_ip_firewall`: Avoid sending `enabled` argument at creation ([#991](https://github.com/ovh/terraform-provider-ovh/pull/991))

📚 Documentation:

* `r/ovh_cloud_project_user`: Add missing roles in the documentation ([#966](https://github.com/ovh/terraform-provider-ovh/pull/966))
* `r/ovh_cloud_project_gateway`: Update docomentation of import ([#993](https://github.com/ovh/terraform-provider-ovh/pull/993))

❤️ Thanks for your contributions ❤️

## 2.2.0 (April 30, 2025)

⚙️ General:

* Bump dependencies ([#950](https://github.com/ovh/terraform-provider-ovh/pull/950))

🎉 Features:

* New resource: `ovh_cloud_project_ssh_key` ([#926](https://github.com/ovh/terraform-provider-ovh/pull/926))
* New resource: `ovh_vrack_dedicated_cloud_datacenter` ([#908](https://github.com/ovh/terraform-provider-ovh/pull/908))
* New resource: `ovh_vrack_ipv6_routed_subrange` ([#947](https://github.com/ovh/terraform-provider-ovh/pull/947))

💪 Improvements:

* `r/ovh_cloud_project_instance`: Allow instance creation on a private network ([#931](https://github.com/ovh/terraform-provider-ovh/pull/931))
* `r/ovh_vrack_ipv6`: Add `bridged_subrange` field ([#947](https://github.com/ovh/terraform-provider-ovh/pull/947))
* `r/ovh_dedicated_server`: Add new fields to control a dedicated server reinstallation: `prevent_install_on_create` and `prevent_install_on_import`. Also add a new parameter to avoid terminating the OVHcloud service when removing the resource in Terraform: `keep_service_after_destroy` and a parameter to run specific actions before destroying the resource: `run_actions_before_destroy` ([#937](https://github.com/ovh/terraform-provider-ovh/pull/937), [#960](https://github.com/ovh/terraform-provider-ovh/pull/960))
* `r/ovh_cloud_project_kube_nodepool`: Add new field `availability_zones` ([#957](https://github.com/ovh/terraform-provider-ovh/pull/957))

* `d/ovh_cloud_project_kube_nodepool`: Add new field `availability_zones` ([#957](https://github.com/ovh/terraform-provider-ovh/pull/957))

🐜 Bug fixes:

* `r/ovh_cloud_project_database`: Correctly handle `backup_time` field ([#938](https://github.com/ovh/terraform-provider-ovh/pull/938))
* `r/ovh_cloud_project_rancher`: Fix version in tests ([#945](https://github.com/ovh/terraform-provider-ovh/pull/945))
* `r/ovh_iploadbalancing_tcp_farm`: Allow setting a null port ([#949](https://github.com/ovh/terraform-provider-ovh/pull/949))
* `r/ovh_iploadbalancing_tcp_farm_server`: Allow setting a null port ([#949](https://github.com/ovh/terraform-provider-ovh/pull/949))
* `r/ovh_iploadbalancing_http_farm`: Allow setting a null port ([#949](https://github.com/ovh/terraform-provider-ovh/pull/949))
* `r/ovh_iploadbalancing_http_farm_server`: Allow setting a null port ([#949](https://github.com/ovh/terraform-provider-ovh/pull/949))
* `r/ovh_domain_name`: Avoid unnecessary diff when `target_spec` is not set ([#954](https://github.com/ovh/terraform-provider-ovh/pull/954))

📚 Documentation:

* Add missing docs and fix some wrong locations ([#925](https://github.com/ovh/terraform-provider-ovh/pull/925))
* Add badges in README ([#941](https://github.com/ovh/terraform-provider-ovh/pull/941))
* Set an up-to-date version for `dbaas_logs` resources ([#955](https://github.com/ovh/terraform-provider-ovh/pull/955))

❤️ Thanks for your contributions ❤️

## 2.1.0 (March 24, 2025)

⚙️ General:

* Bump dependencies ([#903](https://github.com/ovh/terraform-provider-ovh/pull/903))

🎉 Features:

* New resource: `ovh_cloud_project_instance_snapshot` ([#887](https://github.com/ovh/terraform-provider-ovh/pull/887))
* New resource: `ovh_vrack_dedicated_cloud` ([#904](https://github.com/ovh/terraform-provider-ovh/pull/904))

* New datasource: `ovh_cloud_project_storage_object` ([#882](https://github.com/ovh/terraform-provider-ovh/pull/882))
* New datasource: `ovh_cloud_project_storage_objects` ([#882](https://github.com/ovh/terraform-provider-ovh/pull/882))
* New datasource: `ovh_cloud_project_image` ([#909](https://github.com/ovh/terraform-provider-ovh/pull/909))
* New datasource: `ovh_cloud_project_images` ([#909](https://github.com/ovh/terraform-provider-ovh/pull/909))
* New datasource: `ovh_dedicated_cloud` ([#889](https://github.com/ovh/terraform-provider-ovh/pull/889))
* New datasource: `ovh_vmware_cloud_director_organization` ([#900](https://github.com/ovh/terraform-provider-ovh/pull/900))
* New datasource: `ovh_vmware_cloud_director_backup` ([#902](https://github.com/ovh/terraform-provider-ovh/pull/902))

💪 Improvements:

* `r/ovh_cloud_project_region_storage_presign`: Update resource with new fields ([#880](https://github.com/ovh/terraform-provider-ovh/pull/880))

🐜 Bug fixes:

* `r/ovh_cloud_project_rancher`: Make sure `bootstrap_password` is saved on creation ([#888](https://github.com/ovh/terraform-provider-ovh/pull/888))
* `r/ovh_cloud_project_storage`: Remove objects on bucket deletion ([#886](https://github.com/ovh/terraform-provider-ovh/pull/886))
* `r/ovh_dedicated_server_reinstall_task`: Raid type was ignored when 0 was given as value ([#911](https://github.com/ovh/terraform-provider-ovh/pull/911))
* `r/ovh_cloud_project_loadbalancer`: Fix errors in schema and invalid values sent to the API ([#912](https://github.com/ovh/terraform-provider-ovh/pull/912))

📚 Documentation:

* Migrate documentation from legacy to new directory ([#883](https://github.com/ovh/terraform-provider-ovh/pull/883))
* Improve Rancher documentation ([#879](https://github.com/ovh/terraform-provider-ovh/pull/879))
* `r/ovh_dedicated_server_reinstall_task`: Fix invalid examples in documentation ([#910](https://github.com/ovh/terraform-provider-ovh/pull/910))

❤️ Thanks for your contributions ❤️

## 2.0.0 (March 4, 2025)

🧨 Breaking changes:

__In the context of IAM integration across all OVHcloud products, we are deprecating the routes in the API section `/me` that are used to store/retrieve information about dedicated server installation templates and partition schemes.
As a result, the resources `ovh_me_installation_template`, `ovh_me_installation_template_partition_scheme`, `ovh_me_installation_template_partition_scheme_hardware_raid` and `ovh_me_installation_template_partition_scheme_partition` are removed in favor of the new resource `ovh_dedicated_server_reinstall_task` and new parameters in resource `ovh_dedicated_server`.
This will simplify the reinstallation and OS change on dedicated servers, especially when partitioning needs to be customized.__

* `r/ovh_dedicated_server`: The fields used to reinstall a dedicated server have been modified do to API changes (the call [POST  /dedicated/server/{serviceName}/install/start](https://eu.api.ovh.com/console/?section=%2Fdedicated%2Fserver&branch=v1#post-/dedicated/server/-serviceName-/install/start) will be replaced by [POST /dedicated/server/{serviceName}/reinstall](https://eu.api.ovh.com/console/?section=%2Fdedicated%2Fserver&branch=v1#post-/dedicated/server/-serviceName-/reinstall)). In this context, the following fields have been removed: `user_metadata`, `template_name`, `partition_scheme_name`, `details`, and the following fields are replacing them: `os`, `storage`, `customizations`, and `properties` ([#835](https://github.com/ovh/terraform-provider-ovh/pull/835))
* Resource `ovh_dedicated_server_install_task` has been replaced by the new resource `ovh_dedicated_server_reinstall_task` ([#835](https://github.com/ovh/terraform-provider-ovh/pull/835))
* Resources removed: `ovh_me_installation_template`, `ovh_me_installation_template_partition_scheme`, `ovh_me_installation_template_partition_scheme_hardware_raid`, `ovh_me_installation_template_partition_scheme_partition` ([#835](https://github.com/ovh/terraform-provider-ovh/pull/835))

🎉 Features:

* New resource: `ovh_domain_name` ([#837](https://github.com/ovh/terraform-provider-ovh/pull/837))
* New resource: `ovh_domain_name_servers` ([#870](https://github.com/ovh/terraform-provider-ovh/pull/870))
* New resource: `ovh_domain_ds_records` ([#870](https://github.com/ovh/terraform-provider-ovh/pull/870))
* New resource: `ovh_cloud_project_storage` ([#834](https://github.com/ovh/terraform-provider-ovh/pull/834), [#858](https://github.com/ovh/terraform-provider-ovh/pull/858))
* New resource: `ovh_cloud_project_loadbalancer` ([#855](https://github.com/ovh/terraform-provider-ovh/pull/855))
* New resource: `ovh_cloud_project_volume_backup` ([#878](https://github.com/ovh/terraform-provider-ovh/pull/878))
* New resource: `ovh_cloud_project_rancher` ([#872](https://github.com/ovh/terraform-provider-ovh/pull/872))
* New resource: `ovh_dbaas_logs_role` ([#859](https://github.com/ovh/terraform-provider-ovh/pull/859))
* New resource: `ovh_dbaas_logs_role_permission_stream` ([#859](https://github.com/ovh/terraform-provider-ovh/pull/859))
* New resource: `ovh_vrack_ipv6` ([#838](https://github.com/ovh/terraform-provider-ovh/pull/838))
* New resource: `ovh_vrack_vrackservices` ([#839](https://github.com/ovh/terraform-provider-ovh/pull/839))
* New resource: `ovh_vrack_ovhcloudconnect` ([#875](https://github.com/ovh/terraform-provider-ovh/pull/875))
* New resource: `ovh_dedicated_server_reinstall_task` ([#835](https://github.com/ovh/terraform-provider-ovh/pull/835))

* New datasource: `ovh_cloud_project_flavors` ([#865](https://github.com/ovh/terraform-provider-ovh/pull/865))
* New datasource: `ovh_cloud_project_loadbalancer_flavors` ([#866](https://github.com/ovh/terraform-provider-ovh/pull/866))
* New datasource: `ovh_cloud_project_storage` ([#867](https://github.com/ovh/terraform-provider-ovh/pull/867))
* New datasource: `ovh_cloud_project_storages` ([#867](https://github.com/ovh/terraform-provider-ovh/pull/867))
* New datasource: `ovh_cloud_project_rancher` ([#876](https://github.com/ovh/terraform-provider-ovh/pull/876))
* New datasource: `ovh_cloud_project_rancher_version` ([#874](https://github.com/ovh/terraform-provider-ovh/pull/874))
* New datasource: `ovh_cloud_project_rancher_plan` ([#874](https://github.com/ovh/terraform-provider-ovh/pull/874))
* New datasource: `ovh_ovhcloud_connect` ([#871](https://github.com/ovh/terraform-provider-ovh/pull/871))
* New datasource: `ovh_ovhcloud_connects` ([#871](https://github.com/ovh/terraform-provider-ovh/pull/871))

💪 Improvements:

* `r/ovh_cloud_project_instance`: Add `availability_zone` field ([#863](https://github.com/ovh/terraform-provider-ovh/pull/863))
* `d/ovh_cloud_project_instance`: Add `availability_zone` field ([#863](https://github.com/ovh/terraform-provider-ovh/pull/863))
* `d/ovh_cloud_project_instances`: Add `availability_zone` field ([#863](https://github.com/ovh/terraform-provider-ovh/pull/863))

🐜 Bug fixes:

* `r/ovh_cloud_project_instance`: Correctly handle SSH keys ([#861](https://github.com/ovh/terraform-provider-ovh/pull/861))
* `r/ovh_cloud_project_kube`: Increase update timeout ([#868](https://github.com/ovh/terraform-provider-ovh/pull/868))
* `r/ovh_ip_move`: Fix import and add documentation ([#864](https://github.com/ovh/terraform-provider-ovh/pull/864))

📚 Documentation:

* Reorganization of the documentation ([#840](https://github.com/ovh/terraform-provider-ovh/pull/840), [#873](https://github.com/ovh/terraform-provider-ovh/pull/873))
* `r/ovh_vps`: Fix typo in documentation ([#836](https://github.com/ovh/terraform-provider-ovh/pull/836))

❤️ Thanks for your contributions ❤️

## 1.6.0 (February 6, 2025)

⚙️ General:

* Update issue template ([#812](https://github.com/ovh/terraform-provider-ovh/pull/812))
* Bump `go-ovh` dependency ([#830](https://github.com/ovh/terraform-provider-ovh/pull/830))

🎉 Features:

* New datasource: `ovh_cloud_project_floatingips` ([#811](https://github.com/ovh/terraform-provider-ovh/pull/811))
* New datasource: `ovh_dedicated_nasha_partition` ([#817](https://github.com/ovh/terraform-provider-ovh/pull/817))

💪 Improvements:

* `r/ovh_savings_plan`: Add a list of accepted flavors ([#813](https://github.com/ovh/terraform-provider-ovh/pull/813))
* `r/ovh_dedicated_nasha`: Add new field `partitions_list` ([#817](https://github.com/ovh/terraform-provider-ovh/pull/817))
* `r/ovh_dedicated_server`: Add new field `efi_bootloader_path` ([#755](https://github.com/ovh/terraform-provider-ovh/pull/755))
* `r/ovh_dedicated_server_update`: Add new field `efi_bootloader_path` ([#755](https://github.com/ovh/terraform-provider-ovh/pull/755))

* `d/ovh_dedicated_server`: Add new field `efi_bootloader_path` ([#755](https://github.com/ovh/terraform-provider-ovh/pull/755))

🐜 Bug fixes:

* `r/ovh_cloud_project_network_private_subnet_v2`: Set the right property name for allocation pools ([#816](https://github.com/ovh/terraform-provider-ovh/pull/816))
* `r/ovh_me_installation_template`: Fix panic when template does not exist ([#829](https://github.com/ovh/terraform-provider-ovh/pull/829))

📚 Documentation:

* `r/ovh_cloud_project_region_network`: Fix wrong indentations in documentation ([#824](https://github.com/ovh/terraform-provider-ovh/pull/824))
* `r/ovh_cloud_project_network_private`: Update documentation ([#826](https://github.com/ovh/terraform-provider-ovh/pull/826))
* `r/ovh_cloud_project_network_private_subnet`: Update documentation ([#827](https://github.com/ovh/terraform-provider-ovh/pull/827))
* `r/ovh_ip_service`: Add missing configurations in documentation ([#828](https://github.com/ovh/terraform-provider-ovh/pull/828))

❤️ Thanks for your contributions ❤️

## 1.5.0 (January 24, 2025)

⚙️ General:

* Bump golang dependencies ([#803](https://github.com/ovh/terraform-provider-ovh/pull/803))

🎉 Features:

* New resource: `ovh_cloud_project_instance` ([#809](https://github.com/ovh/terraform-provider-ovh/pull/809))
* New resource: `ovh_cloud_project_region_network` ([#808](https://github.com/ovh/terraform-provider-ovh/pull/808))
* New resource: `ovh_cloud_project_database_prometheus` ([#787](https://github.com/ovh/terraform-provider-ovh/pull/787))
* New resource: `ovh_cloud_project_database_mongodb_prometheus` ([#787](https://github.com/ovh/terraform-provider-ovh/pull/787))

* New datasource: `ovh_cloud_project_instance` ([#809](https://github.com/ovh/terraform-provider-ovh/pull/809))
* New datasource: `ovh_cloud_project_instances` ([#809](https://github.com/ovh/terraform-provider-ovh/pull/809))
* New datasource: `ovh_cloud_project_database_prometheus` ([#787](https://github.com/ovh/terraform-provider-ovh/pull/787))
* New datasource: `ovh_cloud_project_database_mongodb_prometheus` ([#787](https://github.com/ovh/terraform-provider-ovh/pull/787))

💪 Improvements:

* `r/ovh_ip_firewall_rule`: Add import capability to resource ([#797](https://github.com/ovh/terraform-provider-ovh/pull/797))
* `r/ovh_savings_plan`: For uppercase for flavors ([#799](https://github.com/ovh/terraform-provider-ovh/pull/799))
* `r/ovh_cloud_project_volume`: Allow size edition ([#804](https://github.com/ovh/terraform-provider-ovh/pull/804))

* `d/ovh_dbaas_logs_cluster_retention`: Add `retention_type` field ([#806](https://github.com/ovh/terraform-provider-ovh/pull/806))

📚 Documentation:

* Licenses update ([#796](https://github.com/ovh/terraform-provider-ovh/pull/796))

❤️ Thanks for your contributions ❤️

## 1.4.0 (December 26, 2024)

🎉 Features:

* New resource: `ovh_savings_plan` ([#793](https://github.com/ovh/terraform-provider-ovh/pull/793))

📚 Documentation:

* `r/ovh_cloud_project_volume`: Update documentation ([#792](https://github.com/ovh/terraform-provider-ovh/pull/792))

❤️ Thanks for your contributions ❤️

## 1.3.0 (December 20, 2024)

🎉 Features:

* New resource: `ovh_cloud_project_volume` ([#778](https://github.com/ovh/terraform-provider-ovh/pull/778))

🐜 Bug fixes:

* `r/ovh_cloud_project_network_private_subnet_v2`: Correctly handle DNS nameservers definition by adding `use_default_public_dns_resolver` parameter ([#789](https://github.com/ovh/terraform-provider-ovh/pull/789))
* `r/ovh_okms_credential`: Modifying field `identity_urns` now triggers a resource re-creation ([#790](https://github.com/ovh/terraform-provider-ovh/pull/790))

❤️ Thanks for your contributions ❤️

## 1.2.0 (December 17, 2024)

🎉 Features:

* New resource: `ovh_okms_service_key_jwk` ([#777](https://github.com/ovh/terraform-provider-ovh/pull/777))

* New datasource: `ovh_okms_service_key_pem` ([#777](https://github.com/ovh/terraform-provider-ovh/pull/777))

🐜 Bug fixes:

* `r/ovh_dbaas_logs_input`: Use Set instead of List for `allowed_networks` ([#783](https://github.com/ovh/terraform-provider-ovh/pull/783))
* `r/ovh_cloud_project_failover_ip_attach`: Retry attach if IP is not allocated yet ([#785](https://github.com/ovh/terraform-provider-ovh/pull/785))

📚 Documentation:

* `r/ovh_cloud_project_kube`: Clean example on resource with private network ([#775](https://github.com/ovh/terraform-provider-ovh/pull/775))

❤️ Thanks for your contributions ❤️

## 1.1.0 (November 19, 2024)

💪 Improvements:

* `r/ovh_cloud_project_database_m3db_namespace`: Allow managing the default namespace ([#761](https://github.com/ovh/terraform-provider-ovh/pull/761))
* `r/ovh_cloud_project_region_storage_presign`: Force region to be uppercased ([#772](https://github.com/ovh/terraform-provider-ovh/pull/772))
* `r/ovh_okms`: Support the new region name format ([#773](https://github.com/ovh/terraform-provider-ovh/pull/773))

* `d/ovh_dedicated_server`: Add missing fields ([#766](https://github.com/ovh/terraform-provider-ovh/pull/766))

📚 Documentation:

* Licenses update ([#764](https://github.com/ovh/terraform-provider-ovh/pull/764))
* `r/ovh_cloud_project_region_storage_presign`: Fix documentation ([#770](https://github.com/ovh/terraform-provider-ovh/pull/770))
* `r/ovh_cloud_project`: Add an example of cloud project creation with a pre-existing vRack ([#771](https://github.com/ovh/terraform-provider-ovh/pull/771))

❤️ Thanks for your contributions ❤️

## 1.0.0 (October 21, 2024)

🎉 Features:

* New resource: `ovh_dbaas_logs_output_opensearch_alias` ([#742](https://github.com/ovh/terraform-provider-ovh/pull/742))
* New resource: `ovh_dbaas_logs_output_opensearch_index` ([#742](https://github.com/ovh/terraform-provider-ovh/pull/742))
* New datasource: `ovh_cloud_project_volume` ([#736](https://github.com/ovh/terraform-provider-ovh/pull/736))
* New datasource: `ovh_cloud_project_volumes` ([#736](https://github.com/ovh/terraform-provider-ovh/pull/736))
* New datasource: `ovh_dbaas_logs_output_opensearch_index` ([#742](https://github.com/ovh/terraform-provider-ovh/pull/742))

🧨 Breaking changes:

* `r/ovh_ip_service`: **BREAKING CHANGE** Fix resource import and use `service_name` as ID. This is a breaking change because the ID of the resource has been modified, but already existing resources will continue working. ([#754](https://github.com/ovh/terraform-provider-ovh/pull/754))
* `r/ovh_cloud_project`: **BREAKING CHANGE** Fix resource import and use `project_id` as ID. This is a breaking change because the ID of the resource has been modified, but already existing resources will continue working. ([#760](https://github.com/ovh/terraform-provider-ovh/pull/760))
* `r/ovh_domain_zone`: **BREAKING CHANGE** Fix resource import and use `name` as ID. This is a breaking change because the ID of the resource has been modified, but already existing resources will continue working. ([#760](https://github.com/ovh/terraform-provider-ovh/pull/760))
* `r/ovh_hosting_privatedatabase`: **BREAKING CHANGE** Fix resource import and use `service_name` as ID. This is a breaking change because the ID of the resource has been modified, but already existing resources will continue working. ([#760](https://github.com/ovh/terraform-provider-ovh/pull/760))
* `r/ovh_iploadbalancing`: **BREAKING CHANGE** Fix resource import and use `service_name` as ID. This is a breaking change because the ID of the resource has been modified, but already existing resources will continue working. ([#760](https://github.com/ovh/terraform-provider-ovh/pull/760))

💪 Improvements:

* Don't read order-related information (`ovh_subsidiary`, `plan` and `plan_option` fields) in Read functions of resources that are created using an order ([#754](https://github.com/ovh/terraform-provider-ovh/pull/754))

🐜 Bug fixes:

* `r/ovh_cloud_project_database`: Prevent setting `backup_time` on several engines ([#756](https://github.com/ovh/terraform-provider-ovh/pull/756))
* Fix marshal and unmarshal of null and unknown values ([#750](https://github.com/ovh/terraform-provider-ovh/pull/750))
* Fix products termination in the US region ([#754](https://github.com/ovh/terraform-provider-ovh/pull/754))

❤️ Thanks for your contributions ❤️

## 0.51.0 (October 11, 2024)

🎉 Features:

* New resource: `ovh_cloud_project_network_private_subnet_v2` ([#732](https://github.com/ovh/terraform-provider-ovh/pull/732))

🐜 Bug fixes:

* `r/ovh_iploadbalancing_udp_farm`: Fix schema errors ([#735](https://github.com/ovh/terraform-provider-ovh/pull/735))
* `r/ovh_iploadbalancing_udp_farm_server`: Fix schema errors ([#735](https://github.com/ovh/terraform-provider-ovh/pull/735))
* `r/ovh_iploadbalancing_udp_frontend`: Fix schema errors ([#735](https://github.com/ovh/terraform-provider-ovh/pull/735))
* `r/ovh_iploadbalancing_ssl`: Force replacement if fingerprint changes ([#733](https://github.com/ovh/terraform-provider-ovh/pull/733))
* `r/ovh_ip_service`: Correctly retrieve `service_name` in US region ([#740](https://github.com/ovh/terraform-provider-ovh/pull/740))
* `r/ovh_dedicated_server`: Only save `user_metadata` coming from the plan in the state ([#744](https://github.com/ovh/terraform-provider-ovh/pull/744), [#745](https://github.com/ovh/terraform-provider-ovh/pull/745))
* `r/ovh_cloud_project_database`: Allow setting `maintenance_time` ([#748](https://github.com/ovh/terraform-provider-ovh/pull/748))

📚 Documentation:

* Licenses update ([#731](https://github.com/ovh/terraform-provider-ovh/pull/731))
* `r/ovh_cloud_project_kube`: Add details on `nodes_subnet_id` and `load_balancers_subnet_id` ([#738](https://github.com/ovh/terraform-provider-ovh/pull/738))
* Reorganize `Public cloud network` and `Log subscriptions` documentations ([#746](https://github.com/ovh/terraform-provider-ovh/pull/746))

❤️ Thanks for your contributions ❤️

## 0.50.0 (September 20, 2024)

🎉 Features:

* New resource: `ovh_okms` ([#709](https://github.com/ovh/terraform-provider-ovh/pull/709))
* New resource: `ovh_okms_credential` ([#709](https://github.com/ovh/terraform-provider-ovh/pull/709))
* New resource: `ovh_okms_service_key` ([#709](https://github.com/ovh/terraform-provider-ovh/pull/709))
* New resource: `ovh_iploadbalancing_ssl` ([#718](https://github.com/ovh/terraform-provider-ovh/pull/718))
* New resource: `ovh_iploadbalancing_udp_farm` ([#719](https://github.com/ovh/terraform-provider-ovh/pull/719))
* New resource: `ovh_iploadbalancing_udp_farm_server` ([#719](https://github.com/ovh/terraform-provider-ovh/pull/719))

* New datasource: `ovh_okms_credential` ([#709](https://github.com/ovh/terraform-provider-ovh/pull/709))
* New datasource: `ovh_okms_resource` ([#709](https://github.com/ovh/terraform-provider-ovh/pull/709))
* New datasource: `ovh_okms_service_key` ([#709](https://github.com/ovh/terraform-provider-ovh/pull/709))
* New datasource: `ovh_okms_service_key_jwk` ([#709](https://github.com/ovh/terraform-provider-ovh/pull/709))
* New datasource: `ovh_cloud_project_network_private` ([#723](https://github.com/ovh/terraform-provider-ovh/pull/723))
* New datasource: `ovh_cloud_project_network_privates` ([#723](https://github.com/ovh/terraform-provider-ovh/pull/723))
* New datasource: `ovh_cloud_project_network_private_subnets` ([#723](https://github.com/ovh/terraform-provider-ovh/pull/723))

💪 Improvements:

* `r/ovh_iploadbalancing_udp_frontend`: Add ability to import resource ([#719](https://github.com/ovh/terraform-provider-ovh/pull/719))

📚 Documentation:

* Reorganize cloud project docs ([#725](https://github.com/ovh/terraform-provider-ovh/pull/725))
* `r/ovh_vrack`: Update example ([#728](https://github.com/ovh/terraform-provider-ovh/pull/728))

❤️ Thanks for your contributions ❤️

## 0.49.0 (September 11, 2024)

🎉 Features:

* New resource: `ovh_dedicated_server` ([#711](https://github.com/ovh/terraform-provider-ovh/pull/711))
* New resource: `ovh_domain_zone_import` ([#716](https://github.com/ovh/terraform-provider-ovh/pull/716))

💪 Improvements:

* `r/ovh_vrack`: Add resource termination ([#713](https://github.com/ovh/terraform-provider-ovh/pull/713))
* `r/ovh_iploadbalancing_http_farm`: Added `uri` as a possible balance value ([#715](https://github.com/ovh/terraform-provider-ovh/pull/715))
* `r/ovh_dbaas_logs_input`: Add missing parameters for autoscaling ([#720](https://github.com/ovh/terraform-provider-ovh/pull/720))

🐜 Bug fixes:

* `r/ovh_hosting_privatedatabase_whitelist`: Add default netmask when user does not define it ([#703](https://github.com/ovh/terraform-provider-ovh/pull/703))
* `r/cloud_project_user_s3_policy`: Read policy field from API ([#707](https://github.com/ovh/terraform-provider-ovh/pull/707))
* `r/ovh_cloud_project`: Fetch plan configurations when reading orders ([#708](https://github.com/ovh/terraform-provider-ovh/pull/708))

* `d/ovh_hosting_privatedatabase_whitelist`: Add default netmask when user does not define it ([#703](https://github.com/ovh/terraform-provider-ovh/pull/703))

⛔️ Deletions:

* `r/ovh_dedicated_server_install_task`: **BREAKING CHANGE** Removal of deprecated properties `post_installation_script_link` and `post_installation_script_return` ([#710](https://github.com/ovh/terraform-provider-ovh/pull/710))
* `r/ovh_me_installation_template`: **BREAKING CHANGE** Removal of deprecated properties `post_installation_script_link` and `post_installation_script_return` ([#710](https://github.com/ovh/terraform-provider-ovh/pull/710))

* `d/ovh_me_installation_template`: **BREAKING CHANGE** Removal of deprecated properties `post_installation_script_link` and `post_installation_script_return` ([#710](https://github.com/ovh/terraform-provider-ovh/pull/710))

❤️ Thanks for your contributions ❤️

## 0.48.0 (July 31, 2024)

🎉 Features:

* New resource: `ovh_cloud_project_gateway_interface` ([#697](https://github.com/ovh/terraform-provider-ovh/pull/697))
* New resource: `ovh_cloud_project_region_loadbalancer_log_subscription` ([#657](https://github.com/ovh/terraform-provider-ovh/pull/657))

* New datasource: `ovh_cloud_project_gateway_interface` ([#697](https://github.com/ovh/terraform-provider-ovh/pull/697))
* New datasource: `ovh_cloud_project_region_loadbalancer_log_subscription` ([#657](https://github.com/ovh/terraform-provider-ovh/pull/657))
* New datasource: `ovh_cloud_project_region_loadbalancer_log_subscriptions` ([#657](https://github.com/ovh/terraform-provider-ovh/pull/657))

🐜 Bug fixes:

* `r/ovh_cloud_project_kube`: Fixed wrong dependencies between fields ([#695](https://github.com/ovh/terraform-provider-ovh/pull/695))

* `d/ovh_dbaas_logs_cluster`: Correctly set `cluster_id` field ([#694](https://github.com/ovh/terraform-provider-ovh/pull/694))

❤️ Thanks for your contributions ❤️

## 0.47.0 (July 19, 2024)

🎉 Features:

* New datasource: `ovh_dbaas_logs_cluster_retention` ([#691](https://github.com/ovh/terraform-provider-ovh/pull/691))

💪 Improvements:

* `r/ovh_dbaas_logs_output_graylog_stream`: Added computed property `write_token` ([#689](https://github.com/ovh/terraform-provider-ovh/pull/689))
* `r/ovh_dedicated_server_update`: Added property `display_name` ([#690](https://github.com/ovh/terraform-provider-ovh/pull/690))

* `d/ovh_dbaas_logs_output_graylog_stream`: Added computed property `write_token` ([#689](https://github.com/ovh/terraform-provider-ovh/pull/689))

📚 Documentation:

* Added documentation to explain how to define the retention period of a DBaaS Logs Graylog stream ([#692](https://github.com/ovh/terraform-provider-ovh/pull/692))

❤️ Thanks for your contributions ❤️

## 0.46.1 (July 8, 2024)

🎉 Features:

* Added `access_token` authentication method ([#668](https://github.com/ovh/terraform-provider-ovh/pull/668))

💪 Improvements:

* `r/ovh_cloud_project_kube`: Added properties `load_balancers_subnet_id` and `nodes_subnet_id` ([#661](https://github.com/ovh/terraform-provider-ovh/pull/661))
* `r/ovh_dedicated_nasha_partition_access`: Added field `acl_description` ([#678](https://github.com/ovh/terraform-provider-ovh/pull/678))

* `d/ovh_cloud_project_kube`: Added properties `load_balancers_subnet_id` and `nodes_subnet_id` ([#662](https://github.com/ovh/terraform-provider-ovh/pull/662))

🐜 Bug fixes:

* `r/ovh_cloud_project_kube_nodepool`: Increased default timeouts ([#656](https://github.com/ovh/terraform-provider-ovh/pull/656))
* `r/ovh_cloud_project_containerregistry_ip_restrictions_management`: Use `TypeSet` for IP restrictions ([#645](https://github.com/ovh/terraform-provider-ovh/pull/645))
* `r/ovh_cloud_project_containerregistry_ip_restrictions_registry`: Use `TypeSet` for IP restrictions ([#645](https://github.com/ovh/terraform-provider-ovh/pull/645))
* `r/ovh_domain_zone_record`: Use correct default value for TTL ([#672](https://github.com/ovh/terraform-provider-ovh/pull/672), [#679](https://github.com/ovh/terraform-provider-ovh/pull/679))
* `r/ovh_dedicated_server_install_task`: Don't retry task if creation failed ([#676](https://github.com/ovh/terraform-provider-ovh/pull/676))
* `r/ovh_cloud_project`: Retrieve order information when reading the resource to avoid re-creation ([#680](https://github.com/ovh/terraform-provider-ovh/pull/680))

* `d/ovh_cloud_project_containerregistry_ip_restrictions_management`: Use `TypeSet` for IP restrictions ([#645](https://github.com/ovh/terraform-provider-ovh/pull/645))
* `d/ovh_cloud_project_containerregistry_ip_restrictions_registry`: Use `TypeSet` for IP restrictions ([#645](https://github.com/ovh/terraform-provider-ovh/pull/645))

📚 Documentation:

* Added link to the documentation page that explains how to manage API keys ([#665](https://github.com/ovh/terraform-provider-ovh/pull/665))

❤️ Thanks for your contributions ❤️

## 0.45.0 (May 21, 2024)

🎉 Features:

* New resource: `ovh_dbaas_logs_token` ([#649](https://github.com/ovh/terraform-provider-ovh/pull/649))

* New datasource: `ovh_cloud_project_loadbalancer` ([#648](https://github.com/ovh/terraform-provider-ovh/pull/648))
* New datasource: `ovh_cloud_project_loadbalancers` ([#648](https://github.com/ovh/terraform-provider-ovh/pull/648))

💪 Improvements:

* `r/ovh_iploadbalancing_tcp_frontend`: Add field `denied_source` ([#652](https://github.com/ovh/terraform-provider-ovh/pull/652))

📚 Documentation:

* `r/ovh_cloud_project`: Fix example of project ordering with HDS certification ([#651](https://github.com/ovh/terraform-provider-ovh/pull/651))

❤️ Thanks for your contributions ❤️

## 0.44.0 (May 6, 2024)

🎉 Features:

* New datasource: `ovh_dedicated_installation_template` ([#639](https://github.com/ovh/terraform-provider-ovh/pull/639))

💪 Improvements:

* `r/ovh_vrack`: A vRack is now imported using its `service_name` instead of its `order_id` ([#630](https://github.com/ovh/terraform-provider-ovh/pull/630))
* `r/ovh_cloud_project_gateway`: Add fields `external_information` and `interfaces` ([#638](https://github.com/ovh/terraform-provider-ovh/pull/638))
* `r/ovh_me_installation_template`: New properties: `end_of_install`, `inputs`, `no_partitioning`, `soft_raid_only_mirroring` and `subfamily` ([#639](https://github.com/ovh/terraform-provider-ovh/pull/639))

* `d/ovh_me_installation_template`: New properties: `end_of_install`, `inputs`, `lvm_ready`, `no_partitioning`, `soft_raid_only_mirroring` and `subfamily` ([#639](https://github.com/ovh/terraform-provider-ovh/pull/639))

🐜 Bug fixes:

* `r/ovh_cloud_project_network_private`: Add possibility to update regions of a private network ([#624](https://github.com/ovh/terraform-provider-ovh/pull/624))
* `r/ovh_cloud_project_gateway`: Fix resource import ([#634](https://github.com/ovh/terraform-provider-ovh/pull/634))

* `d/ovh_me_identity_group`: Add missing URN in datasource ([#643](https://github.com/ovh/terraform-provider-ovh/pull/643))

📚 Documentation:

* `r/ovh_vrack_cloudproject`: Fix import documentation ([#635](https://github.com/ovh/terraform-provider-ovh/pull/635))
* `r/ovh_cloud_project_kube`: Fix wrong fields in documentation ([#640](https://github.com/ovh/terraform-provider-ovh/pull/640))

⛔️ Deletions:

* `r/ovh_me_ssh_key`: **BREAKING CHANGE** Resource removed ([#639](https://github.com/ovh/terraform-provider-ovh/pull/639))
* `r/ovh_dedicated_server_install_task`: **BREAKING CHANGE** Removal of deprecated properties `language` and `use_spla` ([#639](https://github.com/ovh/terraform-provider-ovh/pull/639))
* `r/ovh_me_installation_template`: **BREAKING CHANGE** Removal of deprecated properties `default_language`, `ssh_key_name`, and `available_languages` ([#639](https://github.com/ovh/terraform-provider-ovh/pull/639))

* `d/ovh_me_installation_template`: **BREAKING CHANGE** Removal of deprecated properties `default_language`, `ssh_key_name` and `available_languages` ([#639](https://github.com/ovh/terraform-provider-ovh/pull/639))
* `d/ovh_me_ssh_key`: **BREAKING CHANGE** Datasource removed ([#639](https://github.com/ovh/terraform-provider-ovh/pull/639))
* `d/ovh_me_ssh_keys`: **BREAKING CHANGE** Datasource removed ([#639](https://github.com/ovh/terraform-provider-ovh/pull/639))

❤️ Thanks for your contributions ❤️

## 0.43.1 (April 23, 2024)

🐜 Bug fixes:

* `r/ovh_cloud_project_user`: Fix user roles validation ([#627](https://github.com/ovh/terraform-provider-ovh/pull/627))

❤️ Thanks for your contributions ❤️

## 0.43.0 (April 22, 2024)

⚙️ General:

* Bump golang dependencies ([#622](https://github.com/ovh/terraform-provider-ovh/pull/622))

🎉 Features:

* New datasource: `ovh_dedicated_server_specifications_network` ([#617](https://github.com/ovh/terraform-provider-ovh/pull/617))

💪 Improvements:

* Support OAuth2.0 authentication ([#620](https://github.com/ovh/terraform-provider-ovh/pull/620))
* `r/ovh_dedicated_server_install_task`: Add `userMetadata` field to customize dedicated server installation ([#619](https://github.com/ovh/terraform-provider-ovh/pull/619))
* `r/ovh_cloud_project_database`: Add `kafka_schema_registry` field ([#623](https://github.com/ovh/terraform-provider-ovh/pull/623))

* `d/ovh_cloud_project_database`: Add `kafka_schema_registry` field ([#623](https://github.com/ovh/terraform-provider-ovh/pull/623))

🐜 Bug fixes:

* `r/ovh_cloud_project_database_mongodb_user`: Correctly manage `admin` user ([#609](https://github.com/ovh/terraform-provider-ovh/pull/609))
* `r/ovh_cloud_project_user`: Validate user roles against API ([#621](https://github.com/ovh/terraform-provider-ovh/pull/621))
* `r/ovh_cloud_project_database`: Kafka engine does not have a `backups` field ([#623](https://github.com/ovh/terraform-provider-ovh/pull/623))

⛔️ Deletions:

* `r/ovh_dedicated_server_install_task`: **Breaking change** Removed property `install_sql_server` ([#619](https://github.com/ovh/terraform-provider-ovh/pull/619))
* `r/ovh_me_installation_template`: **Breaking change** Removed properties `beta`, `deprecated`, `last_modification` and `supports_sql_server` ([#619](https://github.com/ovh/terraform-provider-ovh/pull/619))

* `d/ovh_me_installation_template`: **Breaking change** Removed properties `rating`, `beta`, `deprecated`, `last_modification` and `supports_sql_server` ([#619](https://github.com/ovh/terraform-provider-ovh/pull/619))

📚 Documentation:

* `r/ovh_cloud_project_kube_oidc`: Fix fields names ([#613](https://github.com/ovh/terraform-provider-ovh/pull/613))
* `d/ovh_cloud_project_kube_oidc`: Fix fields names ([#613](https://github.com/ovh/terraform-provider-ovh/pull/613))

❤️ Thanks for your contributions ❤️

## 0.42.0 (April 9, 2024)

🎉 Features:

* New datasource: `ovh_cloud_project_database_postgresql_connection_pools` ([#607](https://github.com/ovh/terraform-provider-ovh/pull/607))

💪 Improvements:

* `r/ovh_cloud_project_database`: Allow IP restrictions declaration directly in the resource instead of using `ovh_cloud_project_database_ip_restriction` (which is now deprecated) ([#600](https://github.com/ovh/terraform-provider-ovh/pull/600) and [#606](https://github.com/ovh/terraform-provider-ovh/pull/606))
* `d/ovh_cloud_project_database`: Add IP restrictions in the datasource ([#600](https://github.com/ovh/terraform-provider-ovh/pull/600))

🐜 Bug fixes:

* `r/ovh_ip_firewall_rule`: Fix type mismatch between request and response bodies ([#610](https://github.com/ovh/terraform-provider-ovh/pull/610))
* `r/ovh_cloud_project_database_postgresql_connection_pool`: Added missing retry on connection pool creation ([#607](https://github.com/ovh/terraform-provider-ovh/pull/607))

❤️ Thanks for your contributions ❤️

## 0.41.0 (April 4, 2024)

🎉 Features:

* New resource: `ovh_iploadbalancing_udp_frontend` ([#587](https://github.com/ovh/terraform-provider-ovh/pull/587))
* New resource: `ovh_domain_zone_dnssec` ([#589](https://github.com/ovh/terraform-provider-ovh/pull/589))
* New resource: `ovh_vps` ([#594](https://github.com/ovh/terraform-provider-ovh/pull/594))
* New resource: `ovh_ip_firewall` ([#596](https://github.com/ovh/terraform-provider-ovh/pull/596))
* New resource: `ovh_ip_firewall_rule` ([#601](https://github.com/ovh/terraform-provider-ovh/pull/601))
* New resource: `ovh_ip_mitigation` ([#604](https://github.com/ovh/terraform-provider-ovh/pull/604))

* New datasource: `ovh_domain_zone_dnssec` ([#589](https://github.com/ovh/terraform-provider-ovh/pull/589))
* New datasource: `ovh_ip_firewall` ([#596](https://github.com/ovh/terraform-provider-ovh/pull/596))
* New datasource: `ovh_ip_firewall_rule` ([#601](https://github.com/ovh/terraform-provider-ovh/pull/601))
* New datasource: `ovh_ip_mitigation` ([#604](https://github.com/ovh/terraform-provider-ovh/pull/604))
* New datasource: `ovh_cloud_project` ([#603](https://github.com/ovh/terraform-provider-ovh/pull/603))
* New datasource: `ovh_cloud_projects` ([#603](https://github.com/ovh/terraform-provider-ovh/pull/603))

💪 Improvements:

* Products ordering is now usable by resources developed with SDK `terraform-plugin-framework` ([#594](https://github.com/ovh/terraform-provider-ovh/pull/594))

🐜 Bug fixes:

* `r/ovh_iploadbalancing_tcp_farm`: Allow update of field `stickiness` ([#586](https://github.com/ovh/terraform-provider-ovh/pull/586))
* `r/ovh_iploadbalancing_tcp_frontend`: Use type Set for fields `allowed_source` and `dedicated_ipfo` ([#590](https://github.com/ovh/terraform-provider-ovh/pull/590))
* `r/ovh_iploadbalancing_http_frontend`: Use type Set for fields `allowed_source` and `dedicated_ipfo` ([#590](https://github.com/ovh/terraform-provider-ovh/pull/590))
* `r/ovh_cloud_project_database`: Increase default timeout for database creation ([#591](https://github.com/ovh/terraform-provider-ovh/pull/591))

⛔️ Deletions:

* `r/ovh_dedicated_server_install_task`: **Breaking change** Removed properties `change_log`, `install_rtm`, `reset_hw_raid` and `use_distrib_kernel` ([#593](https://github.com/ovh/terraform-provider-ovh/pull/593))
* `r/ovh_me_installation_template`: **Breaking change** Removed properties `change_log`, `use_distribution_kernel`, `supports_distribution_kernel` and `supports_rtm` ([#593](https://github.com/ovh/terraform-provider-ovh/pull/593))

* `d/ovh_me_installation_template`: **Breaking change** Removed properties `change_log`, `supports_distribution_kernel`, `supports_rtm` and `use_distribution_kernel` ([#593](https://github.com/ovh/terraform-provider-ovh/pull/593))

📚 Documentation:

* Licenses update ([#597](https://github.com/ovh/terraform-provider-ovh/pull/597))

❤️ Thanks for your contributions ❤️

## 0.40.0 (March 19, 2024)

⚙️ General:

* Bump golang dependencies ([#578](https://github.com/ovh/terraform-provider-ovh/pull/578))

🎉 Features:

* New resource: `ovh_cloud_project_gateway` ([#571](https://github.com/ovh/terraform-provider-ovh/pull/571))
* New resource: `ovh_ip_move` ([#510](https://github.com/ovh/terraform-provider-ovh/pull/510))

* New datasource: `ovh_dedicated_server_specifications_hardware` ([#580](https://github.com/ovh/terraform-provider-ovh/pull/580))

💪 Improvements:

* Configuration is loaded like it is done in go-ovh library ([#575](https://github.com/ovh/terraform-provider-ovh/pull/575))
* `d/ovh_dedicated_server`: Add property `display_name` ([#581](https://github.com/ovh/terraform-provider-ovh/pull/581))

📚 Documentation:

* `r/ovh_cloud_project_kube`: Update documentation page ([#579](https://github.com/ovh/terraform-provider-ovh/pull/579))

❤️ Thanks for your contributions ❤️

## 0.39.0 (March 11, 2024)

📚 Documentation:

* Licenses update ([#572](https://github.com/ovh/terraform-provider-ovh/pull/572))

❤️ Thanks for your contributions ❤️

## 0.38.0 (March 8, 2024)

🎉 Features:

* New resource: `ovh_cloud_project_alerting` ([#564](https://github.com/ovh/terraform-provider-ovh/pull/564))

🐜 Bug fixes:

* `r/ovh_cloud_project_database_m3db_user`: Set `password` field as computed when `password_reset` is used ([#561](https://github.com/ovh/terraform-provider-ovh/pull/561))
* `r/ovh_cloud_project_database_mongodb_user`: Set `password` field as computed when `password_reset` is used ([#561](https://github.com/ovh/terraform-provider-ovh/pull/561))
* `r/ovh_cloud_project_database_opensearch_user`: Set `password` field as computed when `password_reset` is used ([#561](https://github.com/ovh/terraform-provider-ovh/pull/561))
* `r/ovh_cloud_project_database_postgresql_user`: Set `password` field as computed when `password_reset` is used ([#561](https://github.com/ovh/terraform-provider-ovh/pull/561))
* `r/ovh_cloud_project_database_redis_user`: Set `password` field as computed when `password_reset` is used ([#561](https://github.com/ovh/terraform-provider-ovh/pull/561))
* `r/ovh_cloud_project_database_user`: Set `password` field as computed when `password_reset` is used ([#561](https://github.com/ovh/terraform-provider-ovh/pull/561))
* `r/ovh_domain_zone_record`: Resource is now recreated when zone is modified ([#563](https://github.com/ovh/terraform-provider-ovh/pull/563))

⛔️ Deletions:

* `r/ovh_me_ipxe_script`: Removed ([#562](https://github.com/ovh/terraform-provider-ovh/pull/562))

* `d/ovh_me_ipxe_script`: Removed ([#562](https://github.com/ovh/terraform-provider-ovh/pull/562))
* `d/ovh_me_ipxe_scripts`: Removed ([#562](https://github.com/ovh/terraform-provider-ovh/pull/562))

📚 Documentation:

* `r/cloud_project_database`: Update documentation page ([#556](https://github.com/ovh/terraform-provider-ovh/pull/556))

❤️ Thanks for your contributions ❤️

## 0.37.0 (February 14, 2024)

🎉 Features:

* New resource: `ovh_cloud_project_database_postgresql_connection_pool` ([#514](https://github.com/ovh/terraform-provider-ovh/pull/514))
* New resource: `ovh_iam_permissions_group` ([#521](https://github.com/ovh/terraform-provider-ovh/pull/521))
* New resource: `ovh_cloud_project_containerregistry_ip_restrictions_management` ([#552](https://github.com/ovh/terraform-provider-ovh/pull/552))
* New resource: `ovh_cloud_project_containerregistry_ip_restrictions_registry` ([#552](https://github.com/ovh/terraform-provider-ovh/pull/552))

* New datasource: `ovh_cloud_project_database_postgresql_connection_pool` ([#514](https://github.com/ovh/terraform-provider-ovh/pull/514))
* New datasource: `ovh_iam_permissions_group` ([#521](https://github.com/ovh/terraform-provider-ovh/pull/521))
* New datasource: `ovh_iam_permissions_groups` ([#521](https://github.com/ovh/terraform-provider-ovh/pull/521))
* New datasource: `ovh_cloud_project_containerregistry_ip_restrictions_management` ([#552](https://github.com/ovh/terraform-provider-ovh/pull/552))
* New datasource: `ovh_cloud_project_containerregistry_ip_restrictions_registry` ([#552](https://github.com/ovh/terraform-provider-ovh/pull/552))

💪 Improvements:

* IAM resources URNs are not computed anymore but fetched from the API ([#537](https://github.com/ovh/terraform-provider-ovh/pull/537))
* Ability to use fidelity account to pay orders ([#540](https://github.com/ovh/terraform-provider-ovh/pull/540))

* `r/ovh_cloud_project_kube_nodepool`: Add autoscaling settings ([#543](https://github.com/ovh/terraform-provider-ovh/pull/543))
* `r/ovh_cloud_project_user`: Add ability to update resource ([#548](https://github.com/ovh/terraform-provider-ovh/pull/548))
* `r/ovh_cloud_project_database`: Add ability to define custom backup ([#553](https://github.com/ovh/terraform-provider-ovh/pull/553))
* `r/ovh_dedicated_server_update`: Add ability to update boot script ([#545](https://github.com/ovh/terraform-provider-ovh/pull/545))

* `d/ovh_cloud_project_kube_nodepool`: Add autoscaling settings ([#543](https://github.com/ovh/terraform-provider-ovh/pull/543))
* `d/ovh_cloud_project_database`: Add custom backup ([#553](https://github.com/ovh/terraform-provider-ovh/pull/553))
* `d/ovh_dedicated_server_update`: Add boot script ([#545](https://github.com/ovh/terraform-provider-ovh/pull/545))

🐜 Bug fixes:

* `r/ovh_cloud_project_kube_nodepool`: Fix validation of the given taints ([#535](https://github.com/ovh/terraform-provider-ovh/pull/535))
* `r/ovh_cloud_project_kube_nodepool`: Nodepool are now created with the correct desired_nodes ([#538](https://github.com/ovh/terraform-provider-ovh/pull/538))
* `r/ovh_cloud_project_database_mongodb_user`: A breaking change has been applied on the 02/14/2024 on this ressource. **Since version 0.37.0, the authentication database must be indicated for all roles** ([#536](https://github.com/ovh/terraform-provider-ovh/pull/536))
* `r/ovh_cloud_project`: Fix bug on project creation ([#478](https://github.com/ovh/terraform-provider-ovh/pull/478))

* `d/ovh_cloud_project_database_mongodb_user`: A breaking change has been applied on the 02/14/2024 on this ressource. **Since version 0.37.0, the authentication database must be indicated for all roles** ([#536](https://github.com/ovh/terraform-provider-ovh/pull/536))

📚 Documentation:

* `index`: Update documentation page ([#503](https://github.com/ovh/terraform-provider-ovh/pull/503), [#542](https://github.com/ovh/terraform-provider-ovh/pull/542))
* `r/ovh_cloud_project`: Update documentation page ([#541](https://github.com/ovh/terraform-provider-ovh/pull/541))

❤️ Thanks for your contributions ❤️

# 0.36.1 (January 9, 2024)

🐜 Bug fixes:

* `core`: Regression while migrating to `github.com/hashicorp/terraform-plugin-framework` ([#528](https://github.com/ovh/terraform-provider-ovh/issues/528))

❤️ Thanks for your contributions ❤️

## 0.36.0 (January 9, 2024)

🎉 Features:

* New resource: `ovh_me_api_oauth2_client` ([#488](https://github.com/ovh/terraform-provider-ovh/pull/488))
* New resource: `ovh_iam_resource_group` ([#457](https://github.com/ovh/terraform-provider-ovh/pull/457))

* New datasource: `ovh_me_api_oauth2_client` ([#488](https://github.com/ovh/terraform-provider-ovh/pull/488))
* New datasource: `ovh_iam_resource_group` ([#457](https://github.com/ovh/terraform-provider-ovh/pull/457))
* New datasource: `ovh_iam_resource_groups` ([#457](https://github.com/ovh/terraform-provider-ovh/pull/457))
* New datasource: `ovh_dbaas_logs_cluster` ([#446](https://github.com/ovh/terraform-provider-ovh/pull/446))
* New datasource: `ovh_cloud_project_vrack` ([#504](https://github.com/ovh/terraform-provider-ovh/pull/504))

⚙️ General:

* Bump golang dependencies ([#482](https://github.com/ovh/terraform-provider-ovh/pull/482), [#516](https://github.com/ovh/terraform-provider-ovh/pull/516), [#525](https://github.com/ovh/terraform-provider-ovh/pull/525))

💪 Improvements:

* `r/ovh_iam_policy`: Add support for deny in IAM policies ([#483](https://github.com/ovh/terraform-provider-ovh/pull/483))

🐜 Bug fixes:

* `r/ovh_cloud_project_database_ip_restriction`: Prevent terraform apply useless retries when an IP restriction is already set on a database instance ([#489](https://github.com/ovh/terraform-provider-ovh/pull/489))
* `r/ovh_dbaas_logs_cluster`: Add optional parameter `cluster_id` ([#446](https://github.com/ovh/terraform-provider-ovh/pull/446))

* `d/ovh_me_installation_template`: Fix a bug when referencing a template that doesn't exist ([#499](https://github.com/ovh/terraform-provider-ovh/pull/499))
* `d/ovh_dbaas_logs_cluster`: Add optional parameter `cluster_id` ([#446](https://github.com/ovh/terraform-provider-ovh/pull/446))

📚 Documentation:

* `r/ovh_domain_zone`: Update documentation page ([#486](https://github.com/ovh/terraform-provider-ovh/pull/486))
* `r/ovh_vrack`: Add documentation for import ([#513](https://github.com/ovh/terraform-provider-ovh/pull/513))
* `r/ovh_domain_zone`: Improve parameters documentation ([#511](https://github.com/ovh/terraform-provider-ovh/pull/511))

❤️ Thanks for your contributions ❤️

## 0.35.0 (November 7, 2023)

⚙️ General:

* Bump golang dependencies ([#480](https://github.com/ovh/terraform-provider-ovh/pull/480))

🐜 Bug fixes:

* `r/cloud_project_user_s3_credential`: A breaking change has been applied on the 11/06/2023 on the routes behind this ressource. In order to use it or if you have it in your state, **an upgrade to the v0.35.0 version is mandatory.** ([#492](https://github.com/ovh/terraform-provider-ovh/pull/492))
* `d/cloud_project_user_s3_credential`: A breaking change has been applied on the 11/06/2023 on the routes behind this data source. In order to use it or if you have it in your state, **an upgrade to the v0.35.0 version is mandatory.** ([#492](https://github.com/ovh/terraform-provider-ovh/pull/492))
* `r/cloud_project_user`: Add AI training read role ([#475](https://github.com/ovh/terraform-provider-ovh/pull/475))

📚 Documentation:

* `examples/kube-nodepool-deployment`: Add an example with Kube, NodePool and an app ([#468](https://github.com/ovh/terraform-provider-ovh/pull/468))
* `d/dedicated_server_boots`: Update documentation page ([#473](https://github.com/ovh/terraform-provider-ovh/pull/473))
* `r/me_ssh_key`: Update documentation page ([#467](https://github.com/ovh/terraform-provider-ovh/pull/467))

❤️ Thanks for your contributions ❤️

## 0.34.0 (September 11, 2023)

⚙️ General:

* Bump github.com/ovh/go-ovh from 1.4.1 to 1.4.2 ([#460](https://github.com/ovh/terraform-provider-ovh/pull/460))

💪 Improvements:

* `r/cloud_project_containerregistry_oidc`: Add OIDC configuration to containerRegistry ([#459](https://github.com/ovh/terraform-provider-ovh/pull/459))

* `d/cloud_project_containerregistry_oidc`: Add OIDC configuration to containerRegistry ([#459](https://github.com/ovh/terraform-provider-ovh/pull/459))

🐜 Bug fixes:

* `r/iploadbalancing_http_frontend`: Fix behavior when OVH API returns unordered response - Fix [#439](https://github.com/ovh/terraform-provider-ovh/issues/439) ([#458](https://github.com/ovh/terraform-provider-ovh/pull/458))
* `r/cloud_project_kube_test`: Use an environment variable for the Kubernetes previous version ([#466](https://github.com/ovh/terraform-provider-ovh/pull/466))
* `r/order`: Creation, Update and import are now solved for US accounts ([#455](https://github.com/ovh/terraform-provider-ovh/pull/455))
* `r/cloud_project`: Creation, Update and import are now solved for US accounts ([#455](https://github.com/ovh/terraform-provider-ovh/pull/455))

📚 Documentation:

* `index`: Update documentation page ([#466](https://github.com/ovh/terraform-provider-ovh/pull/466))

* `r/cloud_project`: Update documentation page ([#456](https://github.com/ovh/terraform-provider-ovh/pull/456))
* `r/hosting_privatedatabase`: Update documentation page ([#456](https://github.com/ovh/terraform-provider-ovh/pull/456))
* `r/ip_service`: Update documentation page ([#456](https://github.com/ovh/terraform-provider-ovh/pull/456))
* `r/iploadbalancing`: Update documentation page ([#456](https://github.com/ovh/terraform-provider-ovh/pull/456))
* `r/ovh_domain_zone`: Update documentation page ([#456](https://github.com/ovh/terraform-provider-ovh/pull/456))
* `r/vrack`: Update documentation page ([#456](https://github.com/ovh/terraform-provider-ovh/pull/456))
* `r/cloud_project_kube`: Update documentation page ([#461](https://github.com/ovh/terraform-provider-ovh/pull/461))
* `r/cloud_project_kube`: Update documentation page ([#461](https://github.com/ovh/terraform-provider-ovh/pull/461))
* `r/cloud_project_containerregistry`: Update documentation page ([#462](https://github.com/ovh/terraform-provider-ovh/pull/462))
* `r/cloud_project_containerregistry_oidc`: Add documentation page ([#459](https://github.com/ovh/terraform-provider-ovh/pull/459))
* `r/cloud_project_containerregistry_users`: Update documentation page ([#459](https://github.com/ovh/terraform-provider-ovh/pull/459))
* `r/vrack_ip`: Update documentation page ([#465](https://github.com/ovh/terraform-provider-ovh/pull/465))

* `d/order_cart`: Update documentation page ([#456](https://github.com/ovh/terraform-provider-ovh/pull/456))
* `d/order_cart_product`: Update documentation page ([#465](https://github.com/ovh/terraform-provider-ovh/pull/465))
* `d/order_cart_product_options`: Update documentation page ([#465](https://github.com/ovh/terraform-provider-ovh/pull/465))
* `d/order_cart_product_options_plan`: Update documentation page ([#465](https://github.com/ovh/terraform-provider-ovh/pull/465))
* `d/order_cart_product_plan`: Update documentation page ([#465](https://github.com/ovh/terraform-provider-ovh/pull/465))
* `d/cloud_project_containerregistry`: Update documentation page ([#459](https://github.com/ovh/terraform-provider-ovh/pull/459))
* `d/cloud_project_containerregistry_oidc`: Add documentation page ([#459](https://github.com/ovh/terraform-provider-ovh/pull/459))
* `d/cloud_project_containerregistry_users`: Update documentation page ([#459](https://github.com/ovh/terraform-provider-ovh/pull/459))

❤️ Thanks for your contributions ❤️

## 0.33.0 (August 24, 2023)

🎉 Features:

* New resource: `ovh_cloud_project_database_kafka_schemaregistryacl` ([#449](https://github.com/ovh/terraform-provider-ovh/pull/449))

* New datasource: `ovh_cloud_project_database_kafka_schemaregistryacl` ([#449](https://github.com/ovh/terraform-provider-ovh/pull/449))
* New datasource: `ovh_cloud_project_database_kafka_schemaregistryacls` ([#449](https://github.com/ovh/terraform-provider-ovh/pull/449))

⚙️ General:

* Bump google.golang.org/grpc from 1.48.0 to 1.53.0 ([#440](https://github.com/ovh/terraform-provider-ovh/pull/440))

💪 Improvements:

* `r/ovh_dedicated_server_networking`: Specify that this route is for internal use only ([#451](https://github.com/ovh/terraform-provider-ovh/pull/451))

🐜 Bug fixes:

* `r/me_installation_template`: Remove `useDistribKernel` deprecated field and fix the test case ([#452](https://github.com/ovh/terraform-provider-ovh/pull/452))
* `r/cloud_project_database_kafka_schemaregistryacl`: Fix test case ([#454](https://github.com/ovh/terraform-provider-ovh/pull/454))
* `r/me_identity_user`: Fix user update ([#443](https://github.com/ovh/terraform-provider-ovh/pull/443))
* `r/cloud_project_kube_nodepool`: Fix terraform schema where nodepool template is optional but all its attributes are required & fix the issue [427](https://github.com/ovh/terraform-provider-ovh/issues/427) ([#433](https://github.com/ovh/terraform-provider-ovh/pull/433))

* `d/cloud_project_database_kafka_schemaregistryacl`: Fix test case ([#454](https://github.com/ovh/terraform-provider-ovh/pull/454))
* `d/cloud_project_database_kafka_schemaregistryacls`: Fix test case ([#454](https://github.com/ovh/terraform-provider-ovh/pull/454))

📚 Documentation:

* `r/ovh_dedicated_server_networking`: Update documentation page ([#451](https://github.com/ovh/terraform-provider-ovh/pull/451))
* `r/ovh_cloud_project_database_kafka_schemaregistryacl`: Add documentation page ([#451](https://github.com/ovh/terraform-provider-ovh/pull/451))
* `r/iam_policy`: Update documentation page ([#445](https://github.com/ovh/terraform-provider-ovh/pull/445))
* `r/ovh_domain_zone_record`: Update documentation page ([#442](https://github.com/ovh/terraform-provider-ovh/pull/442))
* `r/iploadbalancing_http_farm`: Update documentation page ([#432](https://github.com/ovh/terraform-provider-ovh/pull/432))
* `r/iploadbalancing_http_farm_server`: Update documentation page ([#432](https://github.com/ovh/terraform-provider-ovh/pull/432))
* `r/iploadbalancing_http_frontend`: Update documentation page ([#432](https://github.com/ovh/terraform-provider-ovh/pull/432))
* `r/iploadbalancing_http_route`: Update documentation page ([#432](https://github.com/ovh/terraform-provider-ovh/pull/432))
* `r/iploadbalancing_http_route_rule`: Update documentation page ([#432](https://github.com/ovh/terraform-provider-ovh/pull/432))
* `r/iploadbalancing_tcp_farm`: Update documentation page ([#432](https://github.com/ovh/terraform-provider-ovh/pull/432))
* `r/iploadbalancing_tcp_farm_server`: Update documentation page ([#432](https://github.com/ovh/terraform-provider-ovh/pull/432))
* `r/iploadbalancing_tcp_frontend`: Update documentation page ([#432](https://github.com/ovh/terraform-provider-ovh/pull/432))
* `r/iploadbalancing_tcp_route`: Update documentation page ([#432](https://github.com/ovh/terraform-provider-ovh/pull/432))
* `r/iploadbalancing_tcp_route_rule`: Update documentation page ([#432](https://github.com/ovh/terraform-provider-ovh/pull/432))

* `d/ovh_cloud_project_database_kafka_schemaregistryacl`: Add documentation page ([#451](https://github.com/ovh/terraform-provider-ovh/pull/451))
* `d/ovh_cloud_project_database_kafka_schemaregistryacls`: Add documentation page ([#451](https://github.com/ovh/terraform-provider-ovh/pull/451))
* `d/iam_policies`: Update documentation page ([#445](https://github.com/ovh/terraform-provider-ovh/pull/445))
* `d/iam_policy`: Update documentation page ([#445](https://github.com/ovh/terraform-provider-ovh/pull/445))
* `d/iam_reference_actions`: Update documentation page ([#445](https://github.com/ovh/terraform-provider-ovh/pull/445))
* `d/iam_reference_resource_type`: Update documentation page ([#445](https://github.com/ovh/terraform-provider-ovh/pull/445))

❤️ Thanks for your contributions ❤️

## 0.32.0 (July 18, 2023)

🎉 Features:

* New resource: `ovh_iam_policy` ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))

* New datasource: `ovh_iam_policy` ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* New datasource: `ovh_iam_policies` ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* New datasource: `ovh_iam_reference_action` ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* New datasource: `ovh_iam_reference_resource` ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))

⚙️ General:

* Internal provider authentication: Check authenticaton using /auth/currentCredential instead of /auth/details ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))

💪 Improvements:

* `d/ovh_dbaas_logs_cluster`: Add computed identity URN ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_dedicated_ceph`: Add computed identity URN ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_dedicated_nasha`: Add computed identity URN ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_dedicated_server`: Add computed identity URN ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_domain_zone`: Add computed identity URN ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_hosting_privatedatabase`: Add computed identity URN ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_iploadbalancing`: Add computed identity URN ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_me_identity_user`: Add computed identity URN ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_me`: Add computed identity URN ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_vps`: Add computed identity URN ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))

* `r/ovh_cloud_project`: Add computed identity URN ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `r/ovh_dedicated_server_install_task`: Allow retry on install task creation and add import ([#444](https://github.com/ovh/terraform-provider-ovh/pull/444))
* `r/ovh_domain_zone`: Add computed identity URN ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `r/ovh_hosting_privatedatabase`: Add computed identity URN ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `r/ovh_iploadbalancing`: Add computed identity URN ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `r/ovh_me_identity_group`: Add computed identity URN ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `r/ovh_me_identity_user`: Add computed identity URN ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `r/ovh_vrack`: Add computed identity URN ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))

🐜 Bug fixes:

* `r/ovh_me_identity_user`: Fix IAM policy sweeps after tests ([#438](https://github.com/ovh/terraform-provider-ovh/pull/438))

📚 Documentation:

* `d/ovh_dbaas_logs_cluster`: Update documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_dedicated_ceph`: Update documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_dedicated_nasha`: Update documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_dedicated_server`: Update documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_domain_zone`: Update documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_hosting_privatedatabase`: Update documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_iam_policies`: Add documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_iam_policy`: Add documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_iam_reference_action`: Add documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_iam_reference_resource`: Add documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_iploadbalancing`: Update documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_me_identity_user`: Update documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_me`: Update documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `d/ovh_vps`: Update documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))

* `r/ovh_cloud_project`: Update documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `r/ovh_dedicated_server_install_task`: Update documentation page ([#444](https://github.com/ovh/terraform-provider-ovh/pull/444))
* `r/ovh_domain_zone`: Update documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `r/ovh_hosting_privatedatabase`: Update documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `r/ovh_iam_policy`: Add documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `r/ovh_iploadbalancing`: Update documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `r/ovh_me_identity_group`: Update documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `r/ovh_me_identity_user`: Update documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))
* `r/ovh_vrack`: Update documentation page ([#424](https://github.com/ovh/terraform-provider-ovh/pull/424))

❤️ Thanks for your contributions ❤️

## 0.31.0 (June 19, 2023)

Welcome to the the first step of the IAM integration (beta) 🎉.

🎉 Features:

* New resource: `me_identity_group` ([#420](https://github.com/ovh/terraform-provider-ovh/pull/420))

* New datasource: `me_identity_group` ([#420](https://github.com/ovh/terraform-provider-ovh/pull/420))
* New datasource: `me_identity_groups` ([#420](https://github.com/ovh/terraform-provider-ovh/pull/420))

⚙️ General:

* bump go-ovh version to `1.4.1` to fix OVH-Query-Id log ([#425](https://github.com/ovh/terraform-provider-ovh/pull/425))

💪 Improvements:

* `r/ovh_iploadbalancing_tcp_farm_server`: Skip validate func for proxy_protocol_version if nil ([#428](https://github.com/ovh/terraform-provider-ovh/pull/428))
* `r/ovh_iploadbalancing_http_farm_server`: Skip validate func for proxy_protocol_version if nil ([#428](https://github.com/ovh/terraform-provider-ovh/pull/428))
* `r/ovh_cloud_project_database`: Add retry logic to avoid Conflict error on Kafka Topic + Clean code ([#426](https://github.com/ovh/terraform-provider-ovh/pull/426))
* `r/ovh_cloud_project_database_integration`: Add retry logic to avoid Conflict error on Kafka Topic + Clean code ([#426](https://github.com/ovh/terraform-provider-ovh/pull/426))
* `r/ovh_cloud_project_database_kafka_topic`: Add retry logic to avoid Conflict error on Kafka Topic + Clean code ([#426](https://github.com/ovh/terraform-provider-ovh/pull/426))
* `r/ovh_cloud_project_database_m3db_namespace`: Add retry logic to avoid Conflict error on Kafka Topic + Clean code ([#426](https://github.com/ovh/terraform-provider-ovh/pull/426))
* `r/ovh_cloud_project_database_redis_user`: Add retry logic to avoid Conflict error on Kafka Topic + Clean code ([#426](https://github.com/ovh/terraform-provider-ovh/pull/426))
* `r/ovh_cloud_project_database`: Remove unnecessary warning in database resources ([#421](https://github.com/ovh/terraform-provider-ovh/pull/421))
* `r/ovh_cloud_project_database_integration`: Remove unnecessary warning in database resources ([#421](https://github.com/ovh/terraform-provider-ovh/pull/421))
* `r/ovh_cloud_project_database_kafka_topic`: Remove unnecessary warning in database resources ([#421](https://github.com/ovh/terraform-provider-ovh/pull/421))
* `r/ovh_cloud_project_database_m3db_namespace`: Remove unnecessary warning in database resources ([#421](https://github.com/ovh/terraform-provider-ovh/pull/421))
* `r/ovh_cloud_project_database_redis_user`: Remove unnecessary warning in database resources ([#421](https://github.com/ovh/terraform-provider-ovh/pull/421))
* `r/ovh_cloud_project`: Fix typos ([#422](https://github.com/ovh/terraform-provider-ovh/pull/422))
* `r/ovh_domain_zone`: Fix typos ([#422](https://github.com/ovh/terraform-provider-ovh/pull/422))
* `r/ovh_hosting_privatedatabase`: Fix typos ([#422](https://github.com/ovh/terraform-provider-ovh/pull/422))
* `r/ovh_iam_policy`: Fix typos ([#422](https://github.com/ovh/terraform-provider-ovh/pull/422))
* `r/ovh_iploadbalancing`: Fix typos ([#422](https://github.com/ovh/terraform-provider-ovh/pull/422))
* `r/ovh_me_identity_group`: Fix typos ([#422](https://github.com/ovh/terraform-provider-ovh/pull/422))
* `r/ovh_me_identity_user`: Fix typos ([#422](https://github.com/ovh/terraform-provider-ovh/pull/422))
* `r/ovh_vrack`: Fix typos ([#422](https://github.com/ovh/terraform-provider-ovh/pull/422))

* `d/ovh_order_cart`: Fix typos ([#422](https://github.com/ovh/terraform-provider-ovh/pull/422))

📚 Documentation:

* Home Page: Add documentation on multiple providers usage ([#417](https://github.com/ovh/terraform-provider-ovh/pull/417))

* `r/ovh_dedicated_server`: Update documentation page ([#423](https://github.com/ovh/terraform-provider-ovh/pull/423))
* `r/me_identity_group`: Add documentation page ([#420](https://github.com/ovh/terraform-provider-ovh/pull/420))
* `r/ovh_cloud_project_kube`: Update documentation page ([#418](https://github.com/ovh/terraform-provider-ovh/pull/418))

* `d/ovh_order_cart`: Update documentation page ([#422](https://github.com/ovh/terraform-provider-ovh/pull/422))
* `d/me_identity_group`: Add documentation page ([#420](https://github.com/ovh/terraform-provider-ovh/pull/420))
* `d/me_identity_groups`: Add documentation page ([#420](https://github.com/ovh/terraform-provider-ovh/pull/420))

❤️ Thanks for your contributions ❤️

## 0.30.0 (May 3, 2023)

⚙️ General:

* `go` version upgrade to `1.20.0` ([#403](https://github.com/ovh/terraform-provider-ovh/pull/403))
* `terraform-plugin-sdk` upgrade to `2.24.0` ([#406](https://github.com/ovh/terraform-provider-ovh/pull/406))

💪 Improvements:

* `r/ovh_cloud_project_database`: Manage grafana in the generic database user resources ([#412](https://github.com/ovh/terraform-provider-ovh/pull/412))
* `r/ovh_cloud_project_database_*`: Use new function with context and diagnostic - Stop to use deprecated functions ([#403](https://github.com/ovh/terraform-provider-ovh/pull/403))

* `d/ovh_cloud_project_database_*`: Use new function with context and diagnostic - Stop to use deprecated functions ([#403](https://github.com/ovh/terraform-provider-ovh/pull/403))

🐜 Bug fixes:

* `r/ovh_dedicated_ceph_acl`: Fix the acceptance test ([#416](https://github.com/ovh/terraform-provider-ovh/pull/416))
* `r/ovh_cloud_project_database`: Fix the bug that resetted avnadmin password on creation of ovh_cloud_project_database / grafana ([#412](https://github.com/ovh/terraform-provider-ovh/pull/412))
* Fix a lot of acceptance tests ([#406](https://github.com/ovh/terraform-provider-ovh/pull/406))

📚 Documentation:

* `r/ovh_cloud_project_database`: Update documentation page ([#413](https://github.com/ovh/terraform-provider-ovh/pull/413))
* `r/ovh_cloud_project_database_user`: Update documentation page ([#412](https://github.com/ovh/terraform-provider-ovh/pull/412))
* `r/ovh_domain_zone_record`: Update documentation page ([#408](https://github.com/ovh/terraform-provider-ovh/pull/408))

* `d/ovh_cloud_project_database_user`: Update documentation page ([#412](https://github.com/ovh/terraform-provider-ovh/pull/412))

❤️ Thanks for your contributions ❤️

## 0.29.0 (March 24, 2023)

💪 Improvements:

* `r/ovh_domain_zone_record`: Validate zone_record TTL is >=60 ([#397](https://github.com/ovh/terraform-provider-ovh/pull/397))
* `r/ovh_cloud_project_kube_iprestrictions`: Add acceptance tests for all Managed Kubernetes Service resources ([#387](https://github.com/ovh/terraform-provider-ovh/pull/387))
* `r/ovh_cloud_project_kube_nodepool`: Add acceptance tests for all Managed Kubernetes Service resources ([#387](https://github.com/ovh/terraform-provider-ovh/pull/387))
* `r/ovh_cloud_project_kube_oidc`: Add acceptance tests for all Managed Kubernetes Service resources ([#387](https://github.com/ovh/terraform-provider-ovh/pull/387))
* `r/ovh_cloud_project_kube`: Add acceptance tests for all Managed Kubernetes Service resources ([#387](https://github.com/ovh/terraform-provider-ovh/pull/387))

🐜 Bug fixes:

* `r/ovh_ip_reverse`: Update separator for ip reverse to fix the acceptance test ([#394](https://github.com/ovh/terraform-provider-ovh/pull/394))
* `r/ovh_cloud_project_kube_nodepool`: desired_nodes to 0 was not taken into account ([#389](https://github.com/ovh/terraform-provider-ovh/pull/389))

* `d/ovh_me_paymentmean_bankaccount`: Fix issue to be able to order with BankAccount or CreditCard ([#390](https://github.com/ovh/terraform-provider-ovh/pull/390))
* `d/ovh_me_paymentmean_creditcard`: Fix issue to be able to order with BankAccount or CreditCard ([#390](https://github.com/ovh/terraform-provider-ovh/pull/390))

📚 Documentation:

* Add subcategory in documentation page ([#400](https://github.com/ovh/terraform-provider-ovh/pull/400))

* `r/ovh_dedicated_nasha_partition`: Update documentation page ([#401](https://github.com/ovh/terraform-provider-ovh/pull/401))
* `r/ovh_dedicated_nasha_partition_access`: Update documentation page ([#401](https://github.com/ovh/terraform-provider-ovh/pull/401))
* `r/ovh_dedicated_nasha_partition_snapshot`: Update documentation page ([#401](https://github.com/ovh/terraform-provider-ovh/pull/401))
* `r/ovh_iploadbalancing_http_farm`: Update documentation page ([#401](https://github.com/ovh/terraform-provider-ovh/pull/401))
* `r/ovh_iploadbalancing_http_farm_server`: Update documentation page ([#401](https://github.com/ovh/terraform-provider-ovh/pull/401))
* `r/ovh_iploadbalancing_http_frontend`: Update documentation page ([#401](https://github.com/ovh/terraform-provider-ovh/pull/401))
* `r/ovh_iploadbalancing_refresh`: Update documentation page ([#399](https://github.com/ovh/terraform-provider-ovh/pull/399))
* `r/ovh_iploadbalancing_tcp_farm`: Update documentation page ([#399](https://github.com/ovh/terraform-provider-ovh/pull/399))
* `r/ovh_iploadbalancing_tcp_farm_server`: Update documentation page ([#399](https://github.com/ovh/terraform-provider-ovh/pull/399))
* `r/ovh_domain_zone_record`: Update documentation page ([#397](https://github.com/ovh/terraform-provider-ovh/pull/397))
* `r/ovh_cloud_project_kube`: Update documentation page ([#395](https://github.com/ovh/terraform-provider-ovh/pull/395))
* `r/ovh_me_installation_template_partition_scheme_partition`: Update documentation page ([#393](https://github.com/ovh/terraform-provider-ovh/pull/393))
* `r/ovh_cloud_project_workflow_backup`: Update documentation page ([#392](https://github.com/ovh/terraform-provider-ovh/pull/392))
* `r/ovh_cloud_project`: Update documentation page ([#390](https://github.com/ovh/terraform-provider-ovh/pull/390))
* `r/ovh_hosting_privatedatabase`: Update documentation page ([#390](https://github.com/ovh/terraform-provider-ovh/pull/390))
* `r/ovh_ip_service`: Update documentation page ([#390](https://github.com/ovh/terraform-provider-ovh/pull/390))
* `r/ovh_iploadbalancing`: Update documentation page ([#390](https://github.com/ovh/terraform-provider-ovh/pull/390))
* `r/ovh_domain_zone`: Update documentation page ([#390](https://github.com/ovh/terraform-provider-ovh/pull/390))
* `r/ovh_vrack`: Update documentation page ([#390](https://github.com/ovh/terraform-provider-ovh/pull/390))
* `r/ovh_vrack_ip`: Update documentation page ([#390](https://github.com/ovh/terraform-provider-ovh/pull/390))

* `d/cloud_project_failover_ip_attach`: Update documentation page ([#401](https://github.com/ovh/terraform-provider-ovh/pull/401))

❤️ Thanks for your contributions ❤️

## 0.28.0 (February 22, 2023)

⚠️ Deprecation:

* `r/cloud_project_kube`: TypeSet `customization.apiserver` is now deprecated in favor of `customization_apiserver` ([#381](https://github.com/ovh/terraform-provider-ovh/pull/381))
* `d/cloud_project_kube`: TypeSet `customization.apiserver` is now deprecated in favor of `customization_apiserver` ([#381](https://github.com/ovh/terraform-provider-ovh/pull/381))

🎉 Features:

* New resource: `ovh_cloud_project_workflow_backup` ([#368](https://github.com/ovh/terraform-provider-ovh/pull/368))
* New resource: `ovh_dbaas_logs_cluster` ([#364](https://github.com/ovh/terraform-provider-ovh/pull/364))
* New datasource: `ovh_dbaas_logs_cluster` ([#364](https://github.com/ovh/terraform-provider-ovh/pull/364))

💪 Improvements:

* `r/cloud_project_kube`: Add kube proxy configuration ([#381](https://github.com/ovh/terraform-provider-ovh/pull/381))
* `r/cloud_project_kube`: Make kubeconfig attributes available as resource output ([#378](https://github.com/ovh/terraform-provider-ovh/pull/378))
* `r/dbaas_logs_input`: Improve acceptance test ([#366](https://github.com/ovh/terraform-provider-ovh/pull/366))
* `r/cloud_project_database`: Add Advanced Configuration & Manage avnadmin user ([#360](https://github.com/ovh/terraform-provider-ovh/pull/360))
* `r/cloud_project_database_m3db_user`: Add Advanced Configuration & Manage avnadmin user ([#360](https://github.com/ovh/terraform-provider-ovh/pull/360))
* `r/cloud_project_database_mongodb_user`: Add Advanced Configuration & Manage avnadmin user ([#360](https://github.com/ovh/terraform-provider-ovh/pull/360))
* `r/cloud_project_database_opensearch_user`: Add Advanced Configuration & Manage avnadmin user ([#360](https://github.com/ovh/terraform-provider-ovh/pull/360))
* `r/cloud_project_database_postgresql_user`: Add Advanced Configuration & Manage avnadmin user ([#360](https://github.com/ovh/terraform-provider-ovh/pull/360))
* `r/cloud_project_database_redis_user`: Add Advanced Configuration & Manage avnadmin user ([#360](https://github.com/ovh/terraform-provider-ovh/pull/360))
* `r/cloud_project_database_user`: Add Advanced Configuration & Manage avnadmin user ([#360](https://github.com/ovh/terraform-provider-ovh/pull/360))

* `d/cloud_project_kube`: Add kube proxy configuration ([#381](https://github.com/ovh/terraform-provider-ovh/pull/381))
* `d/cloud_project_kube`: Make kubeconfig attributes available as resource output ([#378](https://github.com/ovh/terraform-provider-ovh/pull/378))
* `d/dbaas_logs_input_engine`: Improve acceptance test ([#366](https://github.com/ovh/terraform-provider-ovh/pull/366))
* `d/cloud_project_database`: Add Advanced Configuration ([#360](https://github.com/ovh/terraform-provider-ovh/pull/360))

📚 Documentation:

* `r/cloud_project_kube`: Update documentation page ([#381](https://github.com/ovh/terraform-provider-ovh/pull/381))
* `r/cloud_project_kube_iprestrictions`: Update documentation page ([#386](https://github.com/ovh/terraform-provider-ovh/pull/386))
* `r/cloud_project_kube_nodepool`: Update documentation page ([#386](https://github.com/ovh/terraform-provider-ovh/pull/386))
* `r/cloud_project_kube_oidc`: Update documentation page ([#386](https://github.com/ovh/terraform-provider-ovh/pull/386))
* `r/cloud_project_workflow_backup`: Add documentation page ([#368](https://github.com/ovh/terraform-provider-ovh/pull/368))
* `r/ovh_dbaas_logs_cluster`: Add documentation page ([#364](https://github.com/ovh/terraform-provider-ovh/pull/364))
* `r/cloud_project_database`: Update documentation page ([#360](https://github.com/ovh/terraform-provider-ovh/pull/360))
* `r/cloud_project_database_m3db_user`: Update documentation page ([#360](https://github.com/ovh/terraform-provider-ovh/pull/360))
* `r/cloud_project_database_opensearch_user`: Update documentation page ([#360](https://github.com/ovh/terraform-provider-ovh/pull/360))
* `r/cloud_project_database_postgresql_user`: Update documentation page ([#360](https://github.com/ovh/terraform-provider-ovh/pull/360))
* `r/cloud_project_database_redis_user`: Update documentation page ([#360](https://github.com/ovh/terraform-provider-ovh/pull/360))
* `r/cloud_project_database_user`: Update documentation page ([#360](https://github.com/ovh/terraform-provider-ovh/pull/360))

* `d/cloud_project_kube`: Update documentation page ([#381](https://github.com/ovh/terraform-provider-ovh/pull/381))
* `d/ovh_dbaas_logs_cluster`: Add documentation page ([#364](https://github.com/ovh/terraform-provider-ovh/pull/364))
* `d/cloud_project_database`: Update documentation page ([#360](https://github.com/ovh/terraform-provider-ovh/pull/360))
* `d/cloud_project_database_users`: Update documentation page ([#360](https://github.com/ovh/terraform-provider-ovh/pull/360))

❤️ Thanks for your contributions ❤️

## 0.27.0 (February 9, 2023)

💪 Improvements:

* `r/cloud_project_kube`: Add customized timeouts ([#374](https://github.com/ovh/terraform-provider-ovh/pull/374))
* `r/cloud_project_kube_nodepool`: Add customized timeouts ([#374](https://github.com/ovh/terraform-provider-ovh/pull/374))
* `r/cloud_project_kube_iprestrictions`: Add customized timeouts ([#374](https://github.com/ovh/terraform-provider-ovh/pull/374))
* `r/cloud_project_kube_oidc`: Add customized timeouts ([#374](https://github.com/ovh/terraform-provider-ovh/pull/374))
* `r/iploadbalancing_http_farm_server`: Add `on_marked_down` attribute ([#369](https://github.com/ovh/terraform-provider-ovh/pull/369))
* `r/iploadbalancing_tcp_farm_server`: Add `on_marked_down` attribute ([#369](https://github.com/ovh/terraform-provider-ovh/pull/369))
* `r/iploadbalancing_http_frontend`: Add `hsts` attribute ([#365](https://github.com/ovh/terraform-provider-ovh/pull/365))

🐜 Bug fixes:

* `d/order_cart`: Update expire time format to RFC3339 constant to avoid local time issues ([#363](https://github.com/ovh/terraform-provider-ovh/pull/363))

📚 Documentation:

* `r/iploadbalancing_http_farm_server`: Update documentation page ([#375](https://github.com/ovh/terraform-provider-ovh/pull/375))
* `r/iploadbalancing_tcp_farm_server`: Update documentation page ([#375](https://github.com/ovh/terraform-provider-ovh/pull/375))
* `r/iploadbalancing_refresh`: Update documentation page ([#375](https://github.com/ovh/terraform-provider-ovh/pull/375))
* `r/cloud_project_kube`: Update documentation page ([#374](https://github.com/ovh/terraform-provider-ovh/pull/374))
* `r/cloud_project_kube_iprestrictions`: Update documentation page ([#374](https://github.com/ovh/terraform-provider-ovh/pull/374))
* `r/cloud_project_kube_nodepool`: Update documentation page ([#374](https://github.com/ovh/terraform-provider-ovh/pull/374))
* `r/cloud_project_kube_oidc`: Update documentation page ([#374](https://github.com/ovh/terraform-provider-ovh/pull/374))
* `r/iploadbalancing_http_frontend`: Update documentation page ([#365](https://github.com/ovh/terraform-provider-ovh/pull/365))
* `r/cloud_project`: Update documentation page ([#361](https://github.com/ovh/terraform-provider-ovh/pull/361))

* `d/vpss`: Update documentation page ([#372](https://github.com/ovh/terraform-provider-ovh/pull/372))

❤️ Thanks for your contributions ❤️

## 0.26.0 (January 9, 2023)

First of all, we wish you a happy new year 🎉.

🎉 Features:

* New resource: `dedicated_server_networking` ([#351](https://github.com/ovh/terraform-provider-ovh/pull/351))
* New resource: `dedicated_nasha_partition` ([#349](https://github.com/ovh/terraform-provider-ovh/pull/349))
* New resource: `dedicated_nasha_partition_access` ([#349](https://github.com/ovh/terraform-provider-ovh/pull/349))
* New resource: `dedicated_nasha_partition_snapshot` ([#349](https://github.com/ovh/terraform-provider-ovh/pull/349))

* New datasource: `dedicated_nasha` ([#349](https://github.com/ovh/terraform-provider-ovh/pull/349))

💪 Improvements:

* `r/iploadbalancing_http_frontend`: Add support for HTTP header ([#356](https://github.com/ovh/terraform-provider-ovh/pull/356))
* `r/installation_template`: Remove UEFI & GPT support for installation template ([#352](https://github.com/ovh/terraform-provider-ovh/pull/352))
* `d/installation_template`: Remove UEFI & GPT support for installation template ([#352](https://github.com/ovh/terraform-provider-ovh/pull/352))

🐜 Bug fixes:

* `d/cloud_project_kube_nodepool`: Fix Node Pool read template ([#354](https://github.com/ovh/terraform-provider-ovh/pull/354))
* `r/ip_reverse`: Fix import with IPv6 ([#346](https://github.com/ovh/terraform-provider-ovh/pull/346))

📚 Documentation:

* `r/iploadbalancing_http_frontend`: Update documentation page ([#356](https://github.com/ovh/terraform-provider-ovh/pull/356))
* `r/cloud_project_kube`: Update documentation page ([#355](https://github.com/ovh/terraform-provider-ovh/pull/355))
* `r/cloud_project_kube_nodepool`: Update documentation page ([#353](https://github.com/ovh/terraform-provider-ovh/pull/353))
* `r/me_installation_template`: Update documentation page ([#352](https://github.com/ovh/terraform-provider-ovh/pull/352))
* `r/ovh_dedicated_server_networking`: Add documentation page ([#351](https://github.com/ovh/terraform-provider-ovh/pull/351))
* `r/dedicated_nasha_partition`: Add documentation page ([#349](https://github.com/ovh/terraform-provider-ovh/pull/349))
* `r/dedicated_nasha_partition_access`: Add documentation page ([#349](https://github.com/ovh/terraform-provider-ovh/pull/349))
* `r/dedicated_nasha_partition_snapshot`: Add documentation page ([#349](https://github.com/ovh/terraform-provider-ovh/pull/349))
* `r/ovh_domain_zone_record`: Update documentation page ([#348](https://github.com/ovh/terraform-provider-ovh/pull/348))
* `r/ip_reverse`: Update documentation page ([#346](https://github.com/ovh/terraform-provider-ovh/pull/346))

* `d/me_installation_template`: Update documentation page ([#352](https://github.com/ovh/terraform-provider-ovh/pull/352))
* `d/dedicated_nasha`: Add documentation page ([#349](https://github.com/ovh/terraform-provider-ovh/pull/349))

❤️ Thanks for your contributions ❤️

## 0.25.0 (December 16, 2022)

🎉 Features:

* New datasource: `d/vpss` ([#345](https://github.com/ovh/terraform-provider-ovh/pull/345))

💪 Improvements:

* `d/ovh_dbaas_logs_input_engine`: Now take in account a required service_name ([#347](https://github.com/ovh/terraform-provider-ovh/pull/347))


🐜 Bug fixes:

* `d/ovh_dbaas_logs_output_graylog_stream`: Fix acceptance test ([#347](https://github.com/ovh/terraform-provider-ovh/pull/347))

📚 Documentation:

* `d/ovh_dbaas_logs_input_engine`: Update documentation page ([#347](https://github.com/ovh/terraform-provider-ovh/pull/347))
* `d/ovh_dbaas_logs_output_graylog_stream`: Update documentation page ([#347](https://github.com/ovh/terraform-provider-ovh/pull/347))
ovh_dbaas_logs_output_graylog_stream
* `d/vpss`: Add documentation page ([#345](https://github.com/ovh/terraform-provider-ovh/pull/345))
* `d/vps`: Update documentation page ([#345](https://github.com/ovh/terraform-provider-ovh/pull/345))

❤️ Thanks for your contributions ❤️

## 0.24.0 (December 6, 2022)

🎉 Features:

* New resource: `r/ovh_hosting_privatedatabase` ([#247](https://github.com/ovh/terraform-provider-ovh/pull/247))
* New resource: `r/ovh_hosting_privatedatabase_database` ([#247](https://github.com/ovh/terraform-provider-ovh/pull/247))
* New resource: `r/ovh_hosting_privatedatabase_user` ([#247](https://github.com/ovh/terraform-provider-ovh/pull/247))
* New resource: `r/ovh_hosting_privatedatabase_user_grant` ([#247](https://github.com/ovh/terraform-provider-ovh/pull/247))
* New resource: `r/ovh_hosting_privatedatabase_whitelist` ([#247](https://github.com/ovh/terraform-provider-ovh/pull/247))

* New datasource: `d/ovh_hosting_privatedatabase` ([#247](https://github.com/ovh/terraform-provider-ovh/pull/247))
* New datasource: `d/ovh_hosting_privatedatabase_database` ([#247](https://github.com/ovh/terraform-provider-ovh/pull/247))
* New datasource: `d/ovh_hosting_privatedatabase_user` ([#247](https://github.com/ovh/terraform-provider-ovh/pull/247))
* New datasource: `d/ovh_hosting_privatedatabase_user_grant` ([#247](https://github.com/ovh/terraform-provider-ovh/pull/247))
* New datasource: `d/ovh_hosting_privatedatabase_whitelist` ([#247](https://github.com/ovh/terraform-provider-ovh/pull/247))
* New datasource: `d/ovh_cloud_project_kube_nodepool_nodes` ([#340](https://github.com/ovh/terraform-provider-ovh/pull/340))
* New datasource: `d/ovh_cloud_project_kube_nodes` ([#340](https://github.com/ovh/terraform-provider-ovh/pull/340))

💪 Improvements:

* `d/ovh_order_cart`: Add feature to set catalogName ([#250](https://github.com/ovh/terraform-provider-ovh/pull/250))

🐜 Bug fixes:

* `r/ovh_cloud_project_database`: Fix disk type and size ([#341](https://github.com/ovh/terraform-provider-ovh/pull/341))

📚 Documentation:

* Homepage: new environment variables ([#247](https://github.com/ovh/terraform-provider-ovh/pull/247))

* `r/ovh_hosting_privatedatabase`: Add documentation page ([#247](https://github.com/ovh/terraform-provider-ovh/pull/247))
* `r/ovh_hosting_privatedatabase_database`: Add documentation page ([#247](https://github.com/ovh/terraform-provider-ovh/pull/247))
* `r/ovh_hosting_privatedatabase_user`: Add documentation page ([#247](https://github.com/ovh/terraform-provider-ovh/pull/247))
* `r/ovh_hosting_privatedatabase_user_grant`: Add documentation page ([#247](https://github.com/ovh/terraform-provider-ovh/pull/247))
* `r/ovh_hosting_privatedatabase_whitelist`: Add documentation page ([#247](https://github.com/ovh/terraform-provider-ovh/pull/247))

* `d/ovh_hosting_privatedatabase`: Add documentation page ([#247](https://github.com/ovh/terraform-provider-ovh/pull/247))
* `d/ovh_hosting_privatedatabase_database`: Add documentation page ([#247](https://github.com/ovh/terraform-provider-ovh/pull/247))
* `d/ovh_hosting_privatedatabase_user`: Add documentation page ([#247](https://github.com/ovh/terraform-provider-ovh/pull/247))
* `d/ovh_hosting_privatedatabase_user_grant`: Add documentation page ([#247](https://github.com/ovh/terraform-provider-ovh/pull/247))
* `d/ovh_hosting_privatedatabase_whitelist`: Add documentation page ([#247](https://github.com/ovh/terraform-provider-ovh/pull/247))
* `d/ovh_cloud_project_kube_nodepool_nodes`: Add documentation page ([#340](https://github.com/ovh/terraform-provider-ovh/pull/340))
* `d/ovh_cloud_project_kube_nodes`: Add documentation page ([#340](https://github.com/ovh/terraform-provider-ovh/pull/340))
* `d/order_cart`: Update documentation page ([#250](https://github.com/ovh/terraform-provider-ovh/pull/250))

❤️ Thanks for your contributions ❤️


## 0.23.0 (November 22, 2022)

⚠️ Deprecation:

* `r/ovh_vrack_dedicated_server`: this resource is now deprecated, please use `ovh_vrack_dedicated_server_interface` instead ([#337](https://github.com/ovh/terraform-provider-ovh/pull/337))

🎉 Features:

* New resource: `r/ovh_cloud_project_region_storage_presign` ([#326](https://github.com/ovh/terraform-provider-ovh/pull/326))
* New datasource: `d/ovh_cloud_project_kube_oidc` ([#339](https://github.com/ovh/terraform-provider-ovh/pull/339))

💪 Improvements:

* `r/ovh_cloud_project_kube_oidc`: Add more OIDC parameters to configure ([#339](https://github.com/ovh/terraform-provider-ovh/pull/339))
* `r/ovh_cloud_project_database`: Add `disk_size` and `disk_type` information ([#333](https://github.com/ovh/terraform-provider-ovh/pull/333))
* `r/ovh_cloud_project_kube_nodepool`: Replace resizing to upscaling/downscaling ([#328](https://github.com/ovh/terraform-provider-ovh/pull/328))
* `d/ovh_cloud_project_database`: Add `disk_size` and `disk_type` information ([#333](https://github.com/ovh/terraform-provider-ovh/pull/333))

* Add User-Agent on go-ovh client to identify calls that comes from Terraform ([#338](https://github.com/ovh/terraform-provider-ovh/pull/338))

🐜 Bug fixes:

* `r/ovh_cloud_project_kube_nodepool`: Fix the import ([#334](https://github.com/ovh/terraform-provider-ovh/pull/334))

📚 Documentation:

* `r/ovh_cloud_project_kube_oidc`: Update documentation page ([#339](https://github.com/ovh/terraform-provider-ovh/pull/339))
* `r/ovh_vrack_dedicated_server`: Update documentation page ([#337](https://github.com/ovh/terraform-provider-ovh/pull/337))
* `r/ovh_vrack_dedicated_server_interface`: Update documentation page ([#337](https://github.com/ovh/terraform-provider-ovh/pull/337))
* `r/ovh_dedicated_server_update`: Update documentation page ([#337](https://github.com/ovh/terraform-provider-ovh/pull/337))
* `r/ovh_dedicated_server_reboot_task`: Update documentation page ([#337](https://github.com/ovh/terraform-provider-ovh/pull/337))
* `r/ovh_dedicated_server_install_task`: Update documentation page ([#337](https://github.com/ovh/terraform-provider-ovh/pull/337))
* `r/ovh_cloud_project_kube_nodepool`: Fix the import in the documentation page ([#334](https://github.com/ovh/terraform-provider-ovh/pull/334))
* `r/ovh_vrack_cloudproject`: Update documentation page ([#327](https://github.com/ovh/terraform-provider-ovh/pull/327))
* `r/ovh_cloud_project_network_private_subnet`: Update documentation page ([#327](https://github.com/ovh/terraform-provider-ovh/pull/327))
* `r/cloud_project_network_private`: Update documentation page ([#327](https://github.com/ovh/terraform-provider-ovh/pull/327))
* `r/ovh_cloud_project_database`: Update documentation page ([#333](https://github.com/ovh/terraform-provider-ovh/pull/333))
* `r/ovh_cloud_project_region_storage_presign`: Add documentation page ([#326](https://github.com/ovh/terraform-provider-ovh/pull/326))

* `d/ovh_cloud_project_kube_oidc`: Add documentation page ([#339](https://github.com/ovh/terraform-provider-ovh/pull/339))
* `d/ovh_cloud_project_database`: Update documentation page ([#333](https://github.com/ovh/terraform-provider-ovh/pull/333))

* Fix indent in many documentation pages ([#324](https://github.com/ovh/terraform-provider-ovh/pull/324))

❤️ Thanks for your contributions ❤️

## 0.22.0 (October 6, 2022)

🎉 Features:

* New resource: `r/ovh_cloud_project_database_integration` ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* New resource: `r/cloud_project_database_database` ([#313](https://github.com/ovh/terraform-provider-ovh/pull/313))
* New resource: `r/cloud_project_database_m3db_namespace` ([#313](https://github.com/ovh/terraform-provider-ovh/pull/313))
* New resource: `r/cloud_project_database_m3db_user` ([#313](https://github.com/ovh/terraform-provider-ovh/pull/313))

* New datasource: `d/ovh_cloud_project_database_certificates` ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* New datasource: `d/ovh_cloud_project_database_integration` ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* New datasource: `d/ovh_cloud_project_database_integrations` ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* New datasource: `d/cloud_project_database_capabilities` ([#313](https://github.com/ovh/terraform-provider-ovh/pull/313))
* New datasource: `d/cloud_project_database_database` ([#313](https://github.com/ovh/terraform-provider-ovh/pull/313))
* New datasource: `d/cloud_project_database_databases` ([#313](https://github.com/ovh/terraform-provider-ovh/pull/313))
* New datasource: `d/cloud_project_database_m3db_namespace` ([#313](https://github.com/ovh/terraform-provider-ovh/pull/313))
* New datasource: `d/cloud_project_database_m3db_namespaces` ([#313](https://github.com/ovh/terraform-provider-ovh/pull/313))
* New datasource: `d/cloud_project_database_m3db_user` ([#313](https://github.com/ovh/terraform-provider-ovh/pull/313))

💪 Improvements:

* `r/cloud_project_database_m3db_user`: Add Password Reset system ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `r/cloud_project_database_mongodb_user`: Add Password Reset system ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `r/cloud_project_database_opensearch_user`: Add Password Reset system ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `r/cloud_project_database_postgresql_user`: Add Password Reset system ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `r/cloud_project_database_redis_user`: Add Password Reset system ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `r/cloud_project_database_user`: Add Password Reset system ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `r/cloud_project_database`: Add tips on Network_id ([#317](https://github.com/ovh/terraform-provider-ovh/pull/317))

* `d/cloud_project_database_kafka_certificates`: Generalize Certificate feature ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))

🐜 Bug fixes:

* `r/dedicated_server_install_task`: Fix the resource deletion ([#322](https://github.com/ovh/terraform-provider-ovh/pull/322))
* `r/cloud_project_kube`: Fix updating from a version to the next one ([#319](https://github.com/ovh/terraform-provider-ovh/pull/319))
* `r/cloud_project_kube`: Now we have kubeconfig after importing a Kubernetes cluster ([#315](https://github.com/ovh/terraform-provider-ovh/pull/315))

📚 Documentation:

* `r/cloud_project_database_integration`: Add documentation page ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `r/cloud_project_database_database`: Add documentation page ([#313](https://github.com/ovh/terraform-provider-ovh/pull/313))
* `r/cloud_project_database_m3db_namespace`: Add documentation page ([#313](https://github.com/ovh/terraform-provider-ovh/pull/313))
* `r/cloud_project_database_m3db_user`: Add documentation page ([#313](https://github.com/ovh/terraform-provider-ovh/pull/313))
* `r/cloud_project_database_database`: Update documentation page ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `r/cloud_project_database_ip_restriction`: Update documentation page ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `r/cloud_project_database_kafka_acl`: Update documentation page ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `r/cloud_project_database_m3db_user`: Update documentation page ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `r/cloud_project_database_mongodb_user`: Update documentation page ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `r/cloud_project_database_opensearch_user`: Update documentation page ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `r/cloud_project_database_postgresql_user`: Update documentation page ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `r/cloud_project_database_redis_user`: Update documentation page ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `r/cloud_project_database_user`: Update documentation page ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `r/cloud_project_database`: Update documentation page ([#317](https://github.com/ovh/terraform-provider-ovh/pull/317))
* `r/cloud_project_kube`: Update documentation page ([#316](https://github.com/ovh/terraform-provider-ovh/pull/316))
* `r/cloud_project_kube_iprestrictions`: Update documentation page ([#311](https://github.com/ovh/terraform-provider-ovh/pull/311))
* `r/cloud_project_kube_nodepool`: Update documentation page ([#311](https://github.com/ovh/terraform-provider-ovh/pull/311))

* `d/cloud_project_database_certificates`: Add documentation page ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `d/cloud_project_database_integration`: Add documentation page ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `d/cloud_project_database_integrations`: Add documentation page ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `d/cloud_project_database_capabilities`: Add documentation page ([#313](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `d/cloud_project_database_database`: Add documentation page ([#313](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `d/cloud_project_database_databases`: Add documentation page ([#313](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `d/cloud_project_database_m3db_namespace`: Add documentation page ([#313](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `d/cloud_project_database_m3db_namespaces`: Add documentation page ([#313](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `d/cloud_project_database_m3db_user`: Add documentation page ([#313](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `d/cloud_project_database`: Update documentation page ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `d/cloud_project_database_database`: Update documentation page ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `d/cloud_project_database_databases`: Update documentation page ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `d/cloud_project_database_ip_restrictions`: Update documentation page ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `d/cloud_project_database_user`: Update documentation page ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `d/cloud_project_database_users`: Update documentation page ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `d/cloud_project_databases`: Update documentation page ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))
* `d/cloud_project_database_kafka_certificates`: Delete documentation page ([#321](https://github.com/ovh/terraform-provider-ovh/pull/321))

* Better indentation on a lot of documentation pages 😉
* Full review and fixes in the whole documentation 🙂

❤️ Thanks for your contributions ❤️

## 0.21.0 (September 14, 2022)

💪 Improvements:

* `r/ovh_cloud_project_database`: Add customizable timeouts ([#303](https://github.com/ovh/terraform-provider-ovh/pull/303))
* `r/ovh_cloud_project_database_ip_restriction`: Add customizable timeouts ([#303](https://github.com/ovh/terraform-provider-ovh/pull/303))
* `r/ovh_cloud_project_database_kafka_acl`: Add customizable timeouts ([#303](https://github.com/ovh/terraform-provider-ovh/pull/303))
* `r/ovh_cloud_project_database_kafka_topic`: Add customizable timeouts ([#303](https://github.com/ovh/terraform-provider-ovh/pull/303))
* `r/ovh_cloud_project_database_mongodb_user`: Add customizable timeouts ([#303](https://github.com/ovh/terraform-provider-ovh/pull/303))
* `r/ovh_cloud_project_database_opensearch_pattern`: Add customizable timeouts ([#303](https://github.com/ovh/terraform-provider-ovh/pull/303))
* `r/ovh_cloud_project_database_opensearch_user`: Add customizable timeouts ([#303](https://github.com/ovh/terraform-provider-ovh/pull/303))
* `r/ovh_cloud_project_database_postgresql_user`: Add customizable timeouts ([#303](https://github.com/ovh/terraform-provider-ovh/pull/303))
* `r/ovh_cloud_project_database_redis_user`: Add customizable timeouts ([#303](https://github.com/ovh/terraform-provider-ovh/pull/303))
* `r/ovh_cloud_project_database_user`: Add customizable timeouts ([#303](https://github.com/ovh/terraform-provider-ovh/pull/303))
* `r/ovh_cloud_project_kube`: Add customization feature especially for API Server admission plugins at cluster creation and update ([#304](https://github.com/ovh/terraform-provider-ovh/pull/304))

* `d/ovh_cloud_project_kube`: Add customization feature ([#304](https://github.com/ovh/terraform-provider-ovh/pull/304))

🐜 Bug fixes:

* `r/ovh_iploadbalancing_tcp_farm_server`: Fix the resource creation ([#302](https://github.com/ovh/terraform-provider-ovh/pull/302))
* `r/ovh_cloud_project_kube`: Fix helper function  to avoid to bord effect and do not transform 0 value of an int pointer to a nil ([#304](https://github.com/ovh/terraform-provider-ovh/pull/304))
* `r/ovh_cloud_project_kube`: Fix acceptance test ([#305](https://github.com/ovh/terraform-provider-ovh/pull/305))

* `d/cloud_project_database_opensearch_user`: Fix acceptance test ([#300](https://github.com/ovh/terraform-provider-ovh/pull/300))

📚 Documentation:

* `r/cloud_project_containerregistry_user`: Update documentation page ([#299](https://github.com/ovh/terraform-provider-ovh/pull/299))
* `r/cloud_project_database`: Update documentation page ([#299](https://github.com/ovh/terraform-provider-ovh/pull/299))
* `r/cloud_project_database_ip_restriction`: Update documentation page ([#299](https://github.com/ovh/terraform-provider-ovh/pull/299))
* `r/cloud_project_database_kafka_acl`: Update documentation page ([#299](https://github.com/ovh/terraform-provider-ovh/pull/299))
* `r/cloud_project_database_kafka_topic`: Update documentation page ([#299](https://github.com/ovh/terraform-provider-ovh/pull/299))
* `r/cloud_project_database_mongodb_user`: Update documentation page ([#299](https://github.com/ovh/terraform-provider-ovh/pull/299))
* `r/cloud_project_database_opensearch_pattern`: Update documentation page ([#299](https://github.com/ovh/terraform-provider-ovh/pull/299))
* `r/cloud_project_database_opensearch_user`: Update documentation page ([#299](https://github.com/ovh/terraform-provider-ovh/pull/299))
* `r/cloud_project_database_postgresql_user`: Update documentation page ([#299](https://github.com/ovh/terraform-provider-ovh/pull/299))
* `r/cloud_project_database_redis_user`: Update documentation page ([#299](https://github.com/ovh/terraform-provider-ovh/pull/299))
* `r/cloud_project_database_user`: Update documentation page ([#299](https://github.com/ovh/terraform-provider-ovh/pull/299))
* `r/cloud_project_user_s3_credential`: Update documentation page ([#299](https://github.com/ovh/terraform-provider-ovh/pull/299))
* `r/cloud_project_user_s3_policy`: Update documentation page ([#299](https://github.com/ovh/terraform-provider-ovh/pull/299))
* `r/ovh_cloud_project_kube`: Update documentation page ([#304](https://github.com/ovh/terraform-provider-ovh/pull/304))

* `d/cloud_project_database`: Update documentation page ([#299](https://github.com/ovh/terraform-provider-ovh/pull/299))

## 0.20.0 (September 8, 2022)

🎉 Features:

* New resource: `r/ovh_cloud_project_database_kafka_acl` ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* New resource: `r/ovh_cloud_project_database_kafka_topic` ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* New resource: `r/ovh_cloud_project_database_opensearch_pattern` ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* New resource: `r/ovh_cloud_project_database_opensearch_user` ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* New resource: `r/ovh_cloud_project_user_s3_policy` ([#294](https://github.com/ovh/terraform-provider-ovh/pull/294))
* New resource: `r/ovh_cloud_project_user_s3_credential` ([#291](https://github.com/ovh/terraform-provider-ovh/pull/291))
* New resource: `r/ovh_cloud_project_database_mongodb_user` ([#283](https://github.com/ovh/terraform-provider-ovh/pull/283))
* New resource: `r/ovh_cloud_project_database_redis_user` ([#283](https://github.com/ovh/terraform-provider-ovh/pull/283))

* New datasource: `d/ovh_cloud_project_database_kafka_acl` ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* New datasource: `d/ovh_cloud_project_database_kafka_acls` ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* New datasource: `d/ovh_cloud_project_database_kafka_certificates` ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* New datasource: `d/ovh_cloud_project_database_kafka_topic` ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* New datasource: `d/ovh_cloud_project_database_kafka_topics` ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* New datasource: `d/ovh_cloud_project_database_kafka_user_access` ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* New datasource: `d/ovh_cloud_project_database_opensearch_pattern` ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* New datasource: `d/ovh_cloud_project_database_opensearch_patterns` ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* New datasource: `d/ovh_cloud_project_database_opensearch_user` ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* New datasource: `d/ovh_cloud_project_user` ([#294](https://github.com/ovh/terraform-provider-ovh/pull/294))
* New datasource: `d/ovh_cloud_project_users` ([#294](https://github.com/ovh/terraform-provider-ovh/pull/294))
* New datasource: `d/ovh_cloud_project_user_s3_policy` ([#294](https://github.com/ovh/terraform-provider-ovh/pull/294))
* New datasource: `d/ovh_cloud_project_user_s3_credential` ([#291](https://github.com/ovh/terraform-provider-ovh/pull/291))
* New datasource: `d/ovh_cloud_project_user_s3_credentials` ([#291](https://github.com/ovh/terraform-provider-ovh/pull/291))
* New datasource: `d/ovh_cloud_project_database_mongodb_user` ([#283](https://github.com/ovh/terraform-provider-ovh/pull/283))
* New datasource: `d/ovh_cloud_project_database_redis_user` ([#283](https://github.com/ovh/terraform-provider-ovh/pull/283))

💪 Improvements:

* Extend terminate product timeout ([#295](https://github.com/ovh/terraform-provider-ovh/pull/295))
* Improved resource `r/ovh_cloud_project_kube`: update_policy attribute can now be used at cluster creation ([#293](https://github.com/ovh/terraform-provider-ovh/pull/293))

🐜 Bug fixes:

* `r/ovh_cloud_project_network_private`: region_attributes now contains all the regions not only one ([#286](https://github.com/ovh/terraform-provider-ovh/pull/286))
* `r/ovh_cloud_project_containerregistry`: updating a container registry now will no longer trigger error on terraform side ([#282](https://github.com/ovh/terraform-provider-ovh/pull/282))


📚 Documentation:

* `r/ovh_cloud_project_databases`: Update documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `r/ovh_cloud_project_database_ip_restriction`: Update documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `r/ovh_cloud_project_database_kafka_acl` : Add documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `r/ovh_cloud_project_database_kafka_topic` : Add documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `r/ovh_cloud_project_database_opensearch_pattern` : Add documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `r/ovh_cloud_project_database_opensearch_user` : Add documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `r/ovh_cloud_project_database_postgresql_user`: Update documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `r/ovh_cloud_project_database_user`: Update documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `r/ovh_cloud_project_user_s3_policy` : Add documentation page ([#294](https://github.com/ovh/terraform-provider-ovh/pull/294))
* `r/ovh_cloud_project_kube`: Update documentation page ([#293](https://github.com/ovh/terraform-provider-ovh/pull/293))
* `r/ovh_cloud_project_user_s3_credential` : Add documentation page ([#291](https://github.com/ovh/terraform-provider-ovh/pull/291))
* `r/ovh_cloud_project_database`: Update documentation page ([#290](https://github.com/ovh/terraform-provider-ovh/pull/290))
* `r/ovh_cloud_project_database_user`: Update documentation page ([#289](https://github.com/ovh/terraform-provider-ovh/pull/289))
* `r/ovh_dedicated_server_reboot_task`: Update documentation page ([#289](https://github.com/ovh/terraform-provider-ovh/pull/289))
* `r/ovh_dedicated_server_update`: Update documentation page ([#289](https://github.com/ovh/terraform-provider-ovh/pull/289))
* `r/ovh_vrack_ip`: Update documentation page ([#289](https://github.com/ovh/terraform-provider-ovh/pull/289))
* `r/ovh_vrack_ip_loadbalancing`: Update documentation page ([#289](https://github.com/ovh/terraform-provider-ovh/pull/289))
* `r/ovh_vracks`: Update documentation page ([#284](https://github.com/ovh/terraform-provider-ovh/pull/284))
* `r/ovh_cloud_project_network_private`: Update documentation page ([#284](https://github.com/ovh/terraform-provider-ovh/pull/284))
* `r/ovh_iploadbalancing_http_farm`: Update documentation page ([#284](https://github.com/ovh/terraform-provider-ovh/pull/284))
* `r/ovh_iploadbalancing_http_farm_server`: Update documentation page ([#284](https://github.com/ovh/terraform-provider-ovh/pull/284))
* `r/ovh_iploadbalancing_refresh`: Update documentation page ([#284](https://github.com/ovh/terraform-provider-ovh/pull/284))
* `r/ovh_iploadbalancing_tcp_farm`: Update documentation page ([#284](https://github.com/ovh/terraform-provider-ovh/pull/284))
* `r/ovh_iploadbalancing_tcp_farm_server`: Update documentation page ([#284](https://github.com/ovh/terraform-provider-ovh/pull/284))
* `r/ovh_vrack_cloudproject`: Update documentation page ([#284](https://github.com/ovh/terraform-provider-ovh/pull/284))
* `r/ovh_vrack_dedicated_server`: Update documentation page ([#284](https://github.com/ovh/terraform-provider-ovh/pull/284))
* `r/ovh_cloud_project_database_mongodb_user`: Add documentation page ([#283](https://github.com/ovh/terraform-provider-ovh/pull/283))
* `r/ovh_cloud_project_database_redis_user`: Add documentation page ([#283](https://github.com/ovh/terraform-provider-ovh/pull/283))


* `d/ovh_cloud_project_database`: Update documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `d/ovh_cloud_project_database_ip_restrictions`: Update documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `d/ovh_cloud_project_database_kafka_acl` : Add documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `d/ovh_cloud_project_database_kafka_acls` : Add documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `d/ovh_cloud_project_database_kafka_certificates` : Add documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `d/ovh_cloud_project_database_kafka_topic` : Add documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `d/ovh_cloud_project_database_kafka_topics` : Add documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `d/ovh_cloud_project_database_kafka_user_access` : Add documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `d/ovh_cloud_project_database_postgresql_user`: Update documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `d/ovh_cloud_project_database_user`: Update documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `d/ovh_cloud_project_database_users`: Update documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `d/ovh_cloud_project_databases`: Update documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `d/ovh_cloud_project_database_opensearch_pattern` : Add documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `d/ovh_cloud_project_database_opensearch_patterns` : Add documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `d/ovh_cloud_project_database_opensearch_user` : Add documentation page ([#296](https://github.com/ovh/terraform-provider-ovh/pull/296))
* `d/ovh_cloud_project_user` : Add documentation page ([#294](https://github.com/ovh/terraform-provider-ovh/pull/294))
* `d/ovh_cloud_project_users` : Add documentation page ([#294](https://github.com/ovh/terraform-provider-ovh/pull/294))
* `d/ovh_cloud_project_user_s3_policy` : Add documentation page ([#294](https://github.com/ovh/terraform-provider-ovh/pull/294))
* `d/ovh_cloud_project_user_s3_credential` : Add documentation page ([#291](https://github.com/ovh/terraform-provider-ovh/pull/291))
* `d/ovh_cloud_project_user_s3_credentials` : Add documentation page ([#291](https://github.com/ovh/terraform-provider-ovh/pull/291))
* `d/ovh_dedicated_server_boots`: Update documentation page ([#289](https://github.com/ovh/terraform-provider-ovh/pull/289))
* `d/ovh_cloud_project_capabilities_containerregistry_filter`: Update documentation page ([#287](https://github.com/ovh/terraform-provider-ovh/pull/287))
* `d/ovh_cloud_project_capabilities_containerregistry`: Update documentation page ([#287](https://github.com/ovh/terraform-provider-ovh/pull/287))
* `d/ovh_cloud_project_database_mongodb_user`: Add documentation page ([#283](https://github.com/ovh/terraform-provider-ovh/pull/283))
* `d/ovh_cloud_project_database_redis_user`: Update documentation page ([#283](https://github.com/ovh/terraform-provider-ovh/pull/283))

❤️ Thanks for your contributions ❤️


## 0.19.0 (July 28, 2022)

🎉 Features:

* New resource: `r/ovh_cloud_project_database` ([#269](https://github.com/ovh/terraform-provider-ovh/pull/269))
* New resource: `r/ovh_cloud_project_kube_oidc` ([#273](https://github.com/ovh/terraform-provider-ovh/pull/273))
* New resource: `r/ovh_cloud_project_kube_iprestrictions` ([#274](https://github.com/ovh/terraform-provider-ovh/pull/274))
* New resource: `r/ovh_cloud_project_database_ip_restriction` ([#276](https://github.com/ovh/terraform-provider-ovh/pull/276))
* New resource: `r/ovh_cloud_project_database_user` ([#278](https://github.com/ovh/terraform-provider-ovh/pull/278))
* New resource: `r/ovh_cloud_project_database_postgresql_user` ([#278](https://github.com/ovh/terraform-provider-ovh/pull/278))
* New datasource: `d/ovh_cloud_project_kube_nodepool` ([#274](https://github.com/ovh/terraform-provider-ovh/pull/274))
* New datasource: `d/ovh_cloud_project_kube_iprestrictions` ([#274](https://github.com/ovh/terraform-provider-ovh/pull/274))
* New datasource: `d/ovh_cloud_project_databases` ([#276](https://github.com/ovh/terraform-provider-ovh/pull/276))
* New datasource: `d/ovh_cloud_project_database` ([#276](https://github.com/ovh/terraform-provider-ovh/pull/276))
* New datasource: `d/ovh_cloud_project_database_ip_restrictions` ([#276](https://github.com/ovh/terraform-provider-ovh/pull/276))
* New datasource: `d/ovh_cloud_project_database_user` ([#278](https://github.com/ovh/terraform-provider-ovh/pull/278))
* New datasource: `d/ovh_cloud_project_database_postgresql_user` ([#278](https://github.com/ovh/terraform-provider-ovh/pull/278))
* New datasource: `d/ovh_cloud_project_database_users` ([#278](https://github.com/ovh/terraform-provider-ovh/pull/278))

Improvements:

* Improved resource `r/ovh_cloud_project_kube_nodepool`: add nodepool template parameters and update autoscale parameter versus destroy ([#272](https://github.com/ovh/terraform-provider-ovh/pull/272))


📚 Documentation:

* `r/ovh_cloud_project_database`: Add documentation page ([#269](https://github.com/ovh/terraform-provider-ovh/pull/269))
* `r/ovh_cloud_project_kube_nodepool`: Add template information ([#272](https://github.com/ovh/terraform-provider-ovh/pull/272))
* `r/ovh_cloud_project_kube_oidc`: Add documentation page ([#273](https://github.com/ovh/terraform-provider-ovh/pull/273))
* `r/ovh_cloud_project_kube_iprestrictions`: Add documentation page ([#274](https://github.com/ovh/terraform-provider-ovh/pull/274))
* `r/cloud_project_kube`: Add private network configuration & fix version explanation ([#274](https://github.com/ovh/terraform-provider-ovh/pull/274))
* `r/ovh_cloud_project_database_ip_restriction`: Add documentation page ([#276](https://github.com/ovh/terraform-provider-ovh/pull/276))
* `r/ovh_cloud_project_database`: Fix example ([#276](https://github.com/ovh/terraform-provider-ovh/pull/276))
* `r/ovh_cloud_project_database_user`: Add documentation page ([#278](https://github.com/ovh/terraform-provider-ovh/pull/278))
* `r/ovh_cloud_project_database_postgresql_user`: Add documentation page ([#278](https://github.com/ovh/terraform-provider-ovh/pull/278))

* `d/ovh_cloud_project_database`: Add documentation page ([#276](https://github.com/ovh/terraform-provider-ovh/pull/276))
* `d/ovh_cloud_project_database_ip_restrictions`: Add documentation page ([#276](https://github.com/ovh/terraform-provider-ovh/pull/276))
* `d/ovh_cloud_project_databases`: Add documentation page ([#276](https://github.com/ovh/terraform-provider-ovh/pull/276))
* `d/ovh_cloud_project_kube_iprestrictions`: Add documentation page ([#274](https://github.com/ovh/terraform-provider-ovh/pull/274))
* `d/ovh_cloud_project_kube_nodepool`: Add documentation page ([#274](https://github.com/ovh/terraform-provider-ovh/pull/274))
* `d/ovh_cloud_project_database_user`: Add documentation page ([#278](https://github.com/ovh/terraform-provider-ovh/pull/278))
* `d/ovh_cloud_project_database_postgresql_user`: Add documentation page Add documentation page ([#278](https://github.com/ovh/terraform-provider-ovh/pull/278))
* `d/ovh_cloud_project_database_users`: Add documentation page Add documentation page ([#278](https://github.com/ovh/terraform-provider-ovh/pull/278))
* Home Page: fixed local installation section ([#279](https://github.com/ovh/terraform-provider-ovh/pull/273))

❤️ Thanks for your contributions ❤️

## 0.18.1 (June 14, 2022)

BUG FIXES:

* `r/dbaas_logs_output_graylog_stream`: Fix typo in `cold_storage_content` validation list. ([#266](https://github.com/ovh/terraform-provider-ovh/pull/266))


## 0.18.0 (June 9, 2022)

FEATURES:

* __New Datasource:__ `d/ovh_me` ([#262](https://github.com/ovh/terraform-provider-ovh/pull/262))


BUG FIXES:

* `r/cloud_project_kube`: fix kubernetes cluster import. ([#257](https://github.com/ovh/terraform-provider-ovh/pull/257), [#258](https://github.com/ovh/terraform-provider-ovh/pull/258))

Documentation:

* Home Page: add instructions on how to use a locally built provider ([#264](https://github.com/ovh/terraform-provider-ovh/pull/264))


## 0.17.1 (April 6, 2022)

NOTE:

This release fixes the v0.17.0 release.

Documentation:

* Home Page: add missing end of blockquote ([#249](https://github.com/ovh/terraform-provider-ovh/pull/249))


## 0.17.0 (April 6, 2022)


Improvements:

* Adding failover IP address data and resource blocks for lookup and attachment to existing instances ([#234](https://github.com/ovh/terraform-provider-ovh/pull/234))

BUG FIXES:

* `r/domain_zone_record`: fix panic when del. outside tf ([#231](https://github.com/ovh/terraform-provider-ovh/pull/231))
* `r/domain_zone_record`: dont raise error when deleting a missing resource ([#232](https://github.com/ovh/terraform-provider-ovh/pull/232))
* `r/cloud_project_network_private`: remove deprecated regions with mising regions_attributes ([#238](https://github.com/ovh/terraform-provider-ovh/pull/238))
* `r/cloud_project_kube`: correct the import method ([#239](https://github.com/ovh/terraform-provider-ovh/pull/239))

Documentation:

* `r/cloud_project_kube`: Add missing `id` attribute ([#236](https://github.com/ovh/terraform-provider-ovh/pull/236))
* `r/cloud_project_network_private`: fix & improve the list of exported attributes ([#238](https://github.com/ovh/terraform-provider-ovh/pull/238))
* `r/iploadbalancing_http_farm`: fix `oko` to `oco` typo ([#243](https://github.com/ovh/terraform-provider-ovh/pull/243))
* `r/iploadbalancing_tcp_farm`: fix `oko` to `oco` typo ([#243](https://github.com/ovh/terraform-provider-ovh/pull/243))
* Home Page: add OVH_ENDPOINT in alternative configuration of the provider ([#244](https://github.com/ovh/terraform-provider-ovh/pull/244))
* Home Page: define the provider configuration for terraform CLI version 0.12- and 0.13+ ([#248](https://github.com/ovh/terraform-provider-ovh/pull/248))

## 0.16.0 (October 25, 2021)


FEATURES:

* __New Resources:__ `r/ovh_iploadbalancing_tcp_route`, `r/ovh_iploadbalancing_tcp_route_rule` ([#222](https://github.com/ovh/terraform-provider-ovh/pull/222))

Improvements:

* Add regions regions_status deprecation. ([#198](https://github.com/ovh/terraform-provider-ovh/pull/198), [#227](https://github.com/ovh/terraform-provider-ovh/pull/198))
* fix & improve: data sources and kubernetes resources ([#226](https://github.com/ovh/terraform-provider-ovh/pull/226))
* `r/cloud_project_user`: add importer ([#220](https://github.com/ovh/terraform-provider-ovh/pull/220))
* Add missing Ovh subsidiaries ([#224](https://github.com/ovh/terraform-provider-ovh/pull/224))

BUG FIXES:

* fix: use the right json annotation ([#218](https://github.com/ovh/terraform-provider-ovh/pull/218))
* data/cloud/project/kube: fix acctest match on version ([#223](https://github.com/ovh/terraform-provider-ovh/pull/223))



## 0.15.0 (July 7, 2021)

BREAKING CHANGES:

* `r/ip_reverse`: `ipreverse` is renamed `ip_reverse` and is now mandatory ([#209](https://github.com/ovh/terraform-provider-ovh/pull/209))


BUG FIXES:

* `r/dbaas_logs_input`: fix import function ([#205](https://github.com/ovh/terraform-provider-ovh/pull/205))

Improvements:

* Provider is now built against go v1.16. ([#206](https://github.com/ovh/terraform-provider-ovh/pull/206))


## 0.13.1 (June 28, 2021)

NOTE:

This release fixes the v0.13.0 release, which should have included #194 patchset, but
due to an issue during the release process, the resulting binaries published
on the terraform registry didn't include it.

BUG FIXES:

* `r/cloud_project_kube`: fix issue with empty version([#194](https://github.com/ovh/terraform-provider-ovh/pull/194))

## 0.14.0 (June 23, 2021)

IMPORTANT: This release introduces a new kind of resources which are able to order and terminate OVH products.
OVH products are generally not on demand products, and thus may generate heavy costs. To use these
resources, you have to register a default payment mean on your account. These resources are still in
beta, and should be used with care.


FEATURES:

* __New Datasource:__ `d/dbaas_logs_input_engine` ([#202](https://github.com/ovh/terraform-provider-ovh/pull/202))
* __New Datasource:__ `d/dbaas_logs_output_graylog_stream` ([#202](https://github.com/ovh/terraform-provider-ovh/pull/202))
* __New Datasource:__ `d/ip_service` ([#202](https://github.com/ovh/terraform-provider-ovh/pull/202))
* __New Datasource:__ `d/order_cart` ([#202](https://github.com/ovh/terraform-provider-ovh/pull/202))
* __New Datasource:__ `d/order_cart_product` ([#202](https://github.com/ovh/terraform-provider-ovh/pull/202))
* __New Datasource:__ `d/order_cart_product_options` ([#202](https://github.com/ovh/terraform-provider-ovh/pull/202))
* __New Datasource:__ `d/order_cart_product_options_plan` ([#202](https://github.com/ovh/terraform-provider-ovh/pull/202))
* __New Datasource:__ `d/order_cart_product_plan` ([#202](https://github.com/ovh/terraform-provider-ovh/pull/202))
* __New Resource:__ `r/cloud_project` ([#202](https://github.com/ovh/terraform-provider-ovh/pull/202))
* __New Resource:__ `r/dbaas_logs_input` ([#202](https://github.com/ovh/terraform-provider-ovh/pull/202))
* __New Resource:__ `r/dbaas_logs_output_graylog_stream` ([#202](https://github.com/ovh/terraform-provider-ovh/pull/202))
* __New Resource:__ `r/domain_zone` ([#202](https://github.com/ovh/terraform-provider-ovh/pull/202))
* __New Resource:__ `r/ip_service` ([#202](https://github.com/ovh/terraform-provider-ovh/pull/202))
* __New Resource:__ `r/iploadbalancing` ([#202](https://github.com/ovh/terraform-provider-ovh/pull/202))
* __New Resource:__ `r/vrack` ([#202](https://github.com/ovh/terraform-provider-ovh/pull/202))
* __New Resource:__ `r/vrack_ip` ([#202](https://github.com/ovh/terraform-provider-ovh/pull/202))


Improvements:

* Forgotten documentation change about private networking in Kubernetes. ([#195](https://github.com/ovh/terraform-provider-ovh/pull/195))
* Do not search for iploadbalancing service when exact service name was passed ([#200](https://github.com/ovh/terraform-provider-ovh/pull/200))

## 0.13.0 (May 11, 2021)

BREAKING CHANGES:

`OVH_VRACK_ID` and `OVH_PROJECT_ID` environment variables support are removed in favor of `OVH_VRACK_SERVICE` and `OVH_CLOUD_PROJECT_SERVICE`. Accordingly, `project_id` and `vrack_id` arguments are removed from resources in favor of `service_name`.

* `d/cloud_region`: removed ([#193](https://github.com/ovh/terraform-provider-ovh/pull/193))
* `d/cloud_regions`: removed ([#193](https://github.com/ovh/terraform-provider-ovh/pull/193))
* `r/cloud_private_network`: removed ([#193](https://github.com/ovh/terraform-provider-ovh/pull/193))
* `r/cloud_private_network_subnet`: removed ([#193](https://github.com/ovh/terraform-provider-ovh/pull/193))
* `r/cloud_user`: removed ([#193](https://github.com/ovh/terraform-provider-ovh/pull/193))

Improvements:

* increase vrack task timeouts fomr 20m to 60m

BUG FIXES:

* `r/cloud_project_containerregistry`: fix region and plan_id arguments issue ([#193](https://github.com/ovh/terraform-provider-ovh/pull/193))
* `r/dedicated_server_install_task`: fix issue when OVH cleans its tasks database ([#193](https://github.com/ovh/terraform-provider-ovh/pull/193))
* `r/dedicated_server_reboot_task`: fix issue when OVH cleans its tasks database ([#193](https://github.com/ovh/terraform-provider-ovh/pull/193))
* `r/cloud_project_kube`: fix issue with empty version([#194](https://github.com/ovh/terraform-provider-ovh/pull/194))
* `r/cloud_project_kube_nodepool`: fix issue with computed arguments 
* Documentation: fix typo ([#191](https://github.com/ovh/terraform-provider-ovh/pull/191))


## 0.12.0 (May 7, 2021)

FEATURES:

* __New Datasource:__ `ovh_cloud_project_capabilities_containerregistry` ([#184](https://github.com/ovh/terraform-provider-ovh/pull/184))
* __New Datasource:__ `ovh_cloud_project_capabilities_containerregistry_filter` ([#184](https://github.com/ovh/terraform-provider-ovh/pull/184))
* __New Datasource:__ `ovh_cloud_project_containerregistries` ([#184](https://github.com/ovh/terraform-provider-ovh/pull/184))
* __New Datasource:__ `ovh_cloud_project_containerregistry` ([#184](https://github.com/ovh/terraform-provider-ovh/pull/184))
* __New Datasource:__ `ovh_cloud_project_containerregistry_users` ([#184](https://github.com/ovh/terraform-provider-ovh/pull/184))
* __New Resource:__ `ovh_cloud_project_containerregistry` ([#184](https://github.com/ovh/terraform-provider-ovh/pull/184))
* __New Resource:__ `ovh_cloud_project_containerregistry_user` ([#184](https://github.com/ovh/terraform-provider-ovh/pull/184))

IMPROVEMENTS:

* Various documentation fixes, improvements. ([#182](https://github.com/ovh/terraform-provider-ovh/pull/182), [#185](https://github.com/ovh/terraform-provider-ovh/pull/185))
* Better handle error type casting for ovh api errors ([#186](https://github.com/ovh/terraform-provider-ovh/pull/186))
* Add `private_network_id` and `autoscale` arguments to ovh_cloud_project_kube resources.
([#189](https://github.com/ovh/terraform-provider-ovh/pull/189))
* Add missing attributes to ovh_cloud_project_kube_nodepool resource.
([#189](https://github.com/ovh/terraform-provider-ovh/pull/189))

BUG FIXES:

* Updated the GPG key used to verify Terraform installs in response to the Terraform GPG key rotation. ([#750](https://github-redirect.dependabot.com/hashicorp/terraform-plugin-sdk/issues/739))
cf [Terraform Updates for HCSEC-2021-12](https://discuss.hashicorp.com/t/terraform-updates-for-hcsec-2021-12/23570)

## 0.11.0 (March 3, 2021)

FEATURES:

* __New Datasource:__ `ovh_cloud_project_kube` ([#180](https://github.com/ovh/terraform-provider-ovh/pull/180))
* __New Resource:__ `ovh_cloud_project_kube` ([#180](https://github.com/ovh/terraform-provider-ovh/pull/180))
* __New Resource:__ `ovh_cloud_project_kube_nodepool` ([#180](https://github.com/ovh/terraform-provider-ovh/pull/180))

IMPROVEMENTS:

* Documentation improvements on provider setup ([#181](https://github.com/ovh/terraform-provider-ovh/pull/181))


## 0.10.0 (December 7, 2020)

BREAKING CHANGES:

* d/publiccloud*: all datasources ovh_publiccloud_* are removed in favor of ovh_cloud_project_* ([#175](https://github.com/ovh/terraform-provider-ovh/pull/175))
* r/publiccloud*: all resources ovh_publiccloud_* are removed in favor of ovh_cloud_project_* ([#175](https://github.com/ovh/terraform-provider-ovh/pull/175))

NOTES/DEPRECATIONS:

* d/cloud*: all datasources ovh_cloud_* are deprecated in favor of ovh_cloud_project_* ([#175](https://github.com/ovh/terraform-provider-ovh/pull/175))
* r/cloud*: all resources ovh_cloud_* are deprecated in favor of ovh_cloud_project_* ([#175](https://github.com/ovh/terraform-provider-ovh/pull/175))
* d/cloud*: use service_name for identifier ([#173](https://github.com/ovh/terraform-provider-ovh/pull/173))
* r/cloud*: use service_name for identifier ([#173](https://github.com/ovh/terraform-provider-ovh/pull/173))

FEATURES:

* __New Datasource:__ `ovh_me_identity_user` ([#166](https://github.com/ovh/terraform-provider-ovh/pull/166))
* __New Datasource:__ `ovh_me_identity_users` ([#166](https://github.com/ovh/terraform-provider-ovh/pull/166))
* __New Resource:__ `ovh_me_identity_user` ([#166](https://github.com/ovh/terraform-provider-ovh/pull/166))

IMPROVEMENTS:

* enforce CheckDeleted on all resources read operations ([#176](https://github.com/ovh/terraform-provider-ovh/pull/176))
* cicd acceptance tests now run on OVH CDS build system, travis-ci is removed ([#174](https://github.com/ovh/terraform-provider-ovh/pull/174))
* migrate to new lib ovh/terraform-ovh-provider ([#172](https://github.com/ovh/terraform-provider-ovh/pull/172))
* r/iploadbalancing*: add missing sweepers ([#171](https://github.com/ovh/terraform-provider-ovh/pull/171))
* go-ovh lib: bump to v1.1.0 ([#170](https://github.com/ovh/terraform-provider-ovh/pull/170))
* add freebsd support ([#164](https://github.com/ovh/terraform-provider-ovh/pull/164))
* increase vrack task timeout to 20 minutes ([#38b610e](https://github.com/ovh/terraform-provider-ovh/commit/38b610e310b7478d5cbe53bdc2e3dd09581b1340))

BUG FIXES:

* r/iploadbalancing_http_farm: fix probe handling ([#178](https://github.com/ovh/terraform-provider-ovh/pull/178))
* r/iploadbalancing_tcp_farm: fix probe handling ([#178](https://github.com/ovh/terraform-provider-ovh/pull/178))
* r/dedicated_server_update: fix monitoring update ([#178](https://github.com/ovh/terraform-provider-ovh/pull/178))
* d/vps: Fix erroneous types([#164](https://github.com/ovh/terraform-provider-ovh/pull/164))
* r/me_ssh_key: fix setting key default property and handle key not found error ([#158](https://github.com/ovh/terraform-provider-ovh/pull/158))


## 0.9.0 (August 26, 2020)

BREAKING CHANGES:

* provider: This release includes a Terraform SDK upgrade with compatibility for Terraform >= v0.12. The provider is not compatible with Terraform < v0.12 anymore. This update should have no significant changes in behavior for the provider. ([#154](https://github.com/ovh/terraform-provider-ovh/pull/154))

FEATURES:

* __New Datasource:__ `ovh_dedicated_ceph` ([#150](https://github.com/ovh/terraform-provider-ovh/pull/150))
* __New Resource:__ `ovh_dedicated_ceph_acl` ([#150](https://github.com/ovh/terraform-provider-ovh/pull/150))

IMPROVEMENTS:

* Fetch all IPs for dedicated servers. ([#149](https://github.com/ovh/terraform-provider-ovh/pull/149))
* r/ovh_cloud_user: Add roles/service_name attributes. Deprecate project_id attribute. ([#151](https://github.com/ovh/terraform-provider-ovh/pull/151))

BUG FIXES:

* r/iploadbalancing_tcp_frontend, r/iploadbalancing_http_frontend: Fix allowed_source,dedicated_ipfo updates ([#155](https://github.com/ovh/terraform-provider-ovh/pull/155))

## 0.8.0 (May 28, 2020)

NOTES/DEPRECATIONS:

* `*/ovh_iploadbalancing_vrack_network_*`: Deprecate `farm_id` attribute as it conflicts with other resources. ([#144](https://github.com/ovh/terraform-provider-ovh/issues/144))

FEATURES:

* __New Datasource:__ `ovh_me_ipxe_script` ([#141](https://github.com/ovh/terraform-provider-ovh/pull/141))
* __New Datasource:__ `ovh_me_ipxe_scripts` ([#141](https://github.com/ovh/terraform-provider-ovh/pull/141))
* __New Resource:__ `ovh_me_ipxe_script` ([#141](https://github.com/ovh/terraform-provider-ovh/pull/141))

IMPROVEMENTS:

* Stop failing sweepers if mandatory env vars are missing. ([#142](https://github.com/ovh/terraform-provider-ovh/pull/142))
* r/iploadbalancing_*: Add importers and tests([#140](https://github.com/ovh/terraform-provider-ovh/pull/140))
* r/iploadbalancing_tcp_farm, r/iploadbalancing_http_farm: Extend read function to get all values ([#140](https://github.com/ovh/terraform-provider-ovh/pull/140))
* r/iploadbalancing_http_route: Extend read function to get action value ([#140](https://github.com/ovh/terraform-provider-ovh/pull/140))
* r/iploadbalancing_http_route_rule: Read all values ([#140](https://github.com/ovh/terraform-provider-ovh/pull/140))
* r/iploadbalancing_tcp_frontend, r/iploadbalancing_http_frontend: Some code refactoring according to linter ([#140](https://github.com/ovh/terraform-provider-ovh/pull/140))

BUG FIXES:

* r/iploadbalancing_vrack_network: fix sweepers ([#142](https://github.com/ovh/terraform-provider-ovh/pull/142))
* r/iploadbalancing_tcp_farm, r/iploadbalancing_http_farm: Fix typo in 'oco' probe type ([#140](https://github.com/ovh/terraform-provider-ovh/pull/140))
* r/iploadbalancing_tcp_farm_server, r/iploadbalancing_http_farm_server: Allow port to have a nil value  ([#140](https://github.com/ovh/terraform-provider-ovh/pull/140))

## 0.7.0 (March 02, 2020)

FEATURES:

* __New Datasource:__ `ovh_vps` ([#126](https://github.com/ovh/terraform-provider-ovh/pull/126))

IMPROVEMENTS:

* r/iploadbalancing_http_farm: add cookie stickiness ([#133](https://github.com/ovh/terraform-provider-ovh/pull/133))
* r/dedicated_server_reboot_task, r/dedicated_server_install_task: retry task on 500/404 errors due to API unstability ([#134](https://github.com/ovh/terraform-provider-ovh/pull/134))

BUG FIXES:

* r/dedicated_server_reboot_task, r/dedicated_server_install_task: fix missing ForcesNew attributes ([#135](https://github.com/ovh/terraform-provider-ovh/pull/135))
* r/me_ssh_key: fix missing ForcesNew attribute ([#136](https://github.com/ovh/terraform-provider-ovh/pull/136))
* r/domain_zone_record, domain_zone_redirection: don't fail sweepers on missing OVH_ZONE env var. ([#138](https://github.com/ovh/terraform-provider-ovh/pull/138))
* r/cloud_network_private: fix sweeper. ([#138](https://github.com/ovh/terraform-provider-ovh/pull/138))

## 0.6.0 (January 15, 2020)

FEATURES:

* __New Datasource:__ `ovh_dedicated_server` ([#100](https://github.com/ovh/terraform-provider-ovh/pull/100))
* __New Datasource:__ `ovh_dedicated_servers` ([#100](https://github.com/ovh/terraform-provider-ovh/pull/100))
* __New Datasource:__ `ovh_dedicated_server_boots` ([#105](https://github.com/ovh/terraform-provider-ovh/pull/105))
* __New Datasource:__ `ovh_dedicated_server_boots` ([#105](https://github.com/ovh/terraform-provider-ovh/pull/105))
* __New Datasource:__ `ovh_dedicated_server_installation_templates` ([#101](https://github.com/ovh/terraform-provider-ovh/pull/101))
* __New Datasource:__ `ovh_iploadbalancing_vrack_network` ([#127](https://github.com/ovh/terraform-provider-ovh/pull/127))
* __New Datasource:__ `ovh_iploadbalancing_vrack_networks` ([#127](https://github.com/ovh/terraform-provider-ovh/pull/127))
* __New Datasource:__ `ovh_me_installation_template` ([#103](https://github.com/ovh/terraform-provider-ovh/pull/103))
* __New Datasource:__ `ovh_me_installation_templates` ([#103](https://github.com/ovh/terraform-provider-ovh/pull/103))
* __New Datasource:__ `ovh_me_ssh_key` ([#93](https://github.com/ovh/terraform-provider-ovh/pull/93))
* __New Datasource:__ `ovh_me_ssh_keys` ([#93](https://github.com/ovh/terraform-provider-ovh/pull/93))
* __New Datasource:__ `ovh_vracks` ([#114](https://github.com/ovh/terraform-provider-ovh/pull/114))
* __New Resource:__ `ovh_dedicated_server_install_task` ([#117](https://github.com/ovh/terraform-provider-ovh/pull/117))
* __New Resource:__ `ovh_dedicated_server_reboot_task` ([#116](https://github.com/ovh/terraform-provider-ovh/pull/116))
* __New Resource:__ `ovh_dedicated_server_update` ([#116](https://github.com/ovh/terraform-provider-ovh/pull/116))
* __New Resource:__ `ovh_iploadbalancing_vrack_network` ([#127](https://github.com/ovh/terraform-provider-ovh/pull/127),[#129](https://github.com/ovh/terraform-provider-ovh/pull/129))
* __New Resource:__ `ovh_me_installation_template` ([#103](https://github.com/ovh/terraform-provider-ovh/pull/103))
* __New Resource:__ `ovh_me_installation_template_partition_scheme` ([#103](https://github.com/ovh/terraform-provider-ovh/pull/103))
* __New Resource:__ `ovh_me_installation_template_partition_scheme_partition` ([#103](https://github.com/ovh/terraform-provider-ovh/pull/103))
* __New Resource:__ `ovh_me_installation_template_partition_scheme_hardware_raid` ([#103](https://github.com/ovh/terraform-provider-ovh/pull/103))
* __New Resource:__ `ovh_me_ssh_key` ([#93](https://github.com/ovh/terraform-provider-ovh/pull/93))
* __New Resource:__ `ovh_vrack_dedicated_server` ([#115](https://github.com/ovh/terraform-provider-ovh/pull/115))
* __New Resource:__ `ovh_vrack_dedicated_server_interface` ([#115](https://github.com/ovh/terraform-provider-ovh/pull/115))
* __New Resource:__ `ovh_vrack_iploadbalancing` ([#127](https://github.com/ovh/terraform-provider-ovh/pull/127))

IMPROVEMENTS:

* provider: bump to go 1.13 ([#104](https://github.com/ovh/terraform-provider-ovh/pull/104),[#118](https://github.com/ovh/terraform-provider-ovh/pull/118))
* provider: migrate to terraform-plugin-sdk ([#98](https://github.com/ovh/terraform-provider-ovh/pull/98))
* provider: skip testacc if required env vars are missing ([#106](https://github.com/ovh/terraform-provider-ovh/pull/106))
* d/cloud_regions: add "has_services_up" filter ([#112](https://github.com/ovh/terraform-provider-ovh/pull/112))
* r/ip_reverse: add sweeper (([#99](https://github.com/ovh/terraform-provider-ovh/pull/99), [#102](https://github.com/ovh/terraform-provider-ovh/pull/102))
* acceptance tests: Add PreCheck for HTTP Loadbalancing ([#94](https://github.com/ovh/terraform-provider-ovh/pull/94))
* r/cloud_network_private_subnet: add importer ([#124](https://github.com/ovh/terraform-provider-ovh/pull/124))

BUG FIXES:

* helpers: Fix nil pointer funcs which return wrong golang values in case of HCL nil values ([#120](https://github.com/ovh/terraform-provider-ovh/pull/120))
* r/cloud_network_private, r/cloud_network_private_subnet: fix acctest & rework ([#113](https://github.com/ovh/terraform-provider-ovh/pull/113))
* handle record id bigger than 32bits ([#109](https://github.com/ovh/terraform-provider-ovh/pull/109))
* docs: Correct variable escaping in ovh_iploadbalancing_http_route example ([#97](https://github.com/ovh/terraform-provider-ovh/pull/97))
* docs: Add "ovh_iploadbalancing_refresh" to the website sidebar ([#96](https://github.com/ovh/terraform-provider-ovh/pull/96))

## 0.5.0 (May 22, 2019)

NOTES:

* provider: This release includes only a Terraform SDK upgrade with compatibility for Terraform v0.12. The provider remains backwards compatible with Terraform v0.11 and this update should have no significant changes in behavior for the provider. ([#86](https://github.com/ovh/terraform-provider-ovh/issues/86))

## 0.4.0 (May 22, 2019)

NOTES/DEPRECATIONS:

* `*/ovh_publiccloud_*`: Deprecate `publiccloud` data sources & resources (to issue warnings when used) ([#81](https://github.com/ovh/terraform-provider-ovh/issues/81))

FEATURES:

* __New Resource:__ `ovh_iploadbalancing_tcp_frontend` ([#58](https://github.com/ovh/terraform-provider-ovh/issues/58))

IMPROVEMENTS:

* provider: Enable request/response logging in `>=DEBUG` mode ([#77](https://github.com/ovh/terraform-provider-ovh/issues/77))
* provider: Make homedir detection more robust ([#82](https://github.com/ovh/terraform-provider-ovh/issues/82))

BUG FIXES:

* resource/ovh_domain_zone_record: Attempt retries to avoid errors caused by eventual consistency after creation/update/deletion ([#77](https://github.com/ovh/terraform-provider-ovh/issues/77))
* resource/ovh_domain_zone_record: Make fieldtype non-updatable ([#84](https://github.com/ovh/terraform-provider-ovh/issues/84))
* resource/ovh_domain_zone_redirection: Return errors from refreshing after creation/update/deletion ([#77](https://github.com/ovh/terraform-provider-ovh/issues/77))

## 0.3.0 (July 11, 2018)

DEPRECATIONS / NOTES:

Resources and datasources names now reflects the OVH API endpoints. As such,
resources & datasources that doesn't comply are now deprecated and will be removed
in next release.

* data-source/ovh_publiccloud_region: Deprecated in favor of data-source/ovh_cloud_region ([#31](https://github.com/ovh/terraform-provider-ovh/pull/31))
* data-source/ovh_publiccloud_regions: Deprecated in favor of data-source/ovh_cloud_regions ([#31](https://github.com/ovh/terraform-provider-ovh/pull/31))
* resource/ovh_publiccloud_private_network: Deprecated in favor of resource/ovh_cloud_network_private ([#31](https://github.com/ovh/terraform-provider-ovh/pull/31))
* resource/ovh_publiccloud_private_network_subnet: Deprecated in favor of resource/ovh_cloud_network_private_subnet ([#31](https://github.com/ovh/terraform-provider-ovh/pull/31))
* resource/ovh_publiccloud_user: Deprecated in favor of resource/ovh_cloud_user ([#31](https://github.com/ovh/terraform-provider-ovh/pull/31))
* resource/ovh_vrack_publiccloud_attachment: Deprecated in favor of resource/ovh_vrack_cloudproject ([#31](https://github.com/ovh/terraform-provider-ovh/pull/31))

FEATURES

* __New Datasource:__ `ovh_cloud_region` ([#31](https://github.com/ovh/terraform-provider-ovh/pull/31))
* __New Datasource:__ `ovh_cloud_regions` ([#31](https://github.com/ovh/terraform-provider-ovh/pull/31))
* __New Datasource:__ `ovh_domain_zone` ([#39](https://github.com/ovh/terraform-provider-ovh/pull/39))
* __New Datasource:__ `ovh_iploadbalancing` ([#40](https://github.com/ovh/terraform-provider-ovh/pull/40))
* __New Datasource:__ `ovh_me_paymentmean_creditcard` ([#34](https://github.com/ovh/terraform-provider-ovh/pull/34),[#52](https://github.com/ovh/terraform-provider-ovh/pull/52))
* __New Datasource:__ `ovh_me_paymentmean_bankaccount` ([#34](https://github.com/ovh/terraform-provider-ovh/pull/34),[#52](https://github.com/ovh/terraform-provider-ovh/pull/52))
* __New Resource:__ `ovh_iploadbalancing_tcp_farm` ([#32](https://github.com/ovh/terraform-provider-ovh/pull/32))
* __New Resource:__ `ovh_iploadbalancing_tcp_farm_server` ([#33](https://github.com/ovh/terraform-provider-ovh/pull/33))
* __New Resource:__ `ovh_iploadbalancing_http_route` ([#35](https://github.com/ovh/terraform-provider-ovh/pull/35))
* __New Resource:__ `ovh_iploadbalancing_http_route_rule` ([#35](https://github.com/ovh/terraform-provider-ovh/pull/35))
* __New Resource:__ `ovh_domain_zone_redirection` ([#36](https://github.com/ovh/terraform-provider-ovh/pull/36))
* __New Resource:__ `ovh_cloud_network_private` ([#31](https://github.com/ovh/terraform-provider-ovh/pull/31))
* __New Resource:__ `ovh_cloud_network_private_subnet` ([#31](https://github.com/ovh/terraform-provider-ovh/pull/31))
* __New Resource:__ `ovh_cloud_user` ([#31](https://github.com/ovh/terraform-provider-ovh/pull/31))
* __New Resource:__ `ovh_vrack_cloudproject` ([#31](https://github.com/ovh/terraform-provider-ovh/pull/31))

IMPROVEMENTS

* Various doc improvements ([#14](https://github.com/ovh/terraform-provider-ovh/pull/14),[#16](https://github.com/ovh/terraform-provider-ovh/pull/16),[#17](https://github.com/ovh/terraform-provider-ovh/pull/17),[#26](https://github.com/ovh/terraform-provider-ovh/pull/26),[#53](https://github.com/ovh/terraform-provider-ovh/pull/51),[#51](https://github.com/ovh/terraform-provider-ovh/pull/53))
* provider: Fallback to get current home directory ([#19](https://github.com/ovh/terraform-provider-ovh/pull/19))
* provider: bump to terraform v0.10.8 ([#49](https://github.com/ovh/terraform-provider-ovh/pull/49))
* r/ovh_domain_zone_record: add sweeper ([#50](https://github.com/ovh/terraform-provider-ovh/pull/50))
* r/ovh_domain_zone_redirection: add sweeper ([#50](https://github.com/ovh/terraform-provider-ovh/pull/50))


BUG FIXES:

* resource/ovh_domain_zone_record: Fixes [[#25](https://github.com/ovh/terraform-provider-ovh/issues/25)] by id removal, cleans up naming and struct repetition for domain zone record resource
* provider: Fixes [[#27](https://github.com/ovh/terraform-provider-ovh/issues/27)] by switching to `/auth/currentCredential` for client validation

## 0.2.0 (January 10, 2018)

BACKWARDS INCOMPATIBILITIES / NOTES:

* d/ovh_publiccloud_region: Deprecated fields which don't comply
  with lowercase & underscore convention (`continentCode`, `datacenterLocation`).
  Use `continent_code` and `datacenter_location` instead. ([#4](https://github.com/ovh/terraform-provider-ovh/issues/4))

FEATURES

* __New Resource:__ `ovh_domain_zone_record` ([#3](https://github.com/ovh/terraform-provider-ovh/issues/3))

IMPROVEMENTS

* The provider config can now source its credentials from `~/.ovh.conf` ([#10](https://github.com/ovh/terraform-provider-ovh/issues/10))

## 0.1.0 (June 21, 2017)

NOTES:

* Same functionality as that of Terraform 0.9.8. Repacked as part of [Provider Splitout](https://www.hashicorp.com/blog/upcoming-provider-changes-in-terraform-0-10/)
