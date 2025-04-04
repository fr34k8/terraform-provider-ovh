package ovh

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ovh/terraform-provider-ovh/v2/ovh/helpers"
)

func resourceDbaasLogsInput() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDbaasLogsInputCreate,
		UpdateContext: resourceDbaasLogsInputUpdate,
		ReadContext:   resourceDbaasLogsInputRead,
		DeleteContext: resourceDbaasLogsInputDelete,
		Importer: &schema.ResourceImporter{
			State: resourceDbaasLogsInputImportState,
		},

		Schema: resourceDbaasLogsInputSchema(),
	}
}

func resourceDbaasLogsInputSchema() map[string]*schema.Schema {
	schema := map[string]*schema.Schema{
		"service_name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"description": {
			Type:        schema.TypeString,
			Description: "Input description",
			Required:    true,
		},
		"engine_id": {
			Type:        schema.TypeString,
			Description: "Input engine ID",
			Required:    true,
		},
		"stream_id": {
			Type:        schema.TypeString,
			Description: "Associated Graylog stream",
			Required:    true,
		},
		"title": {
			Type:        schema.TypeString,
			Description: "Input title",
			Required:    true,
		},
		"configuration": {
			Type:        schema.TypeList,
			Required:    true,
			ForceNew:    true,
			Description: "Input configuration",
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"flowgger": {
						Type:          schema.TypeList,
						Optional:      true,
						Description:   "Flowgger configuration",
						ConflictsWith: []string{"configuration.0.logstash"},
						MaxItems:      1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"log_format": {
									Type:        schema.TypeString,
									Description: "Type of format to decode",
									Required:    true,
									ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
										err := helpers.ValidateStringEnum(strings.ToUpper(v.(string)), []string{
											"RFC5424",
											"LTSV",
											"GELF",
											"CAPNP",
										})
										if err != nil {
											errors = append(errors, err)
										}
										return
									},
								},
								"log_framing": {
									Type:        schema.TypeString,
									Description: "Indicates how messages are delimited",
									Required:    true,
									ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
										err := helpers.ValidateStringEnum(strings.ToUpper(v.(string)), []string{
											"LINE",
											"NUL",
											"SYSLEN",
											"CAPNP",
										})
										if err != nil {
											errors = append(errors, err)
										}
										return
									},
								},
							},
						},
					},
					"logstash": {
						Type:          schema.TypeList,
						ConflictsWith: []string{"configuration.0.flowgger"},
						Optional:      true,
						MaxItems:      1,
						Description:   "Logstash configuration",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"filter_section": {
									Type:        schema.TypeString,
									Description: "The filter section of logstash.conf",
									Optional:    true,
									Default:     "",
								},
								"input_section": {
									Type:        schema.TypeString,
									Description: "The filter section of logstash.conf",
									Required:    true,
								},
								"pattern_section": {
									Type:        schema.TypeString,
									Description: "The list of customs Grok patterns",
									Optional:    true,
									Default:     "",
								},
							},
						},
					},
				},
			},
		},

		// Optional
		"allowed_networks": {
			Type:        schema.TypeSet,
			Description: "IP blocks",
			Elem:        &schema.Schema{Type: schema.TypeString},
			Optional:    true,
			Computed:    true,
		},
		"exposed_port": {
			Type:        schema.TypeString,
			Description: "Port",
			Optional:    true,
			Computed:    true,
		},
		"nb_instance": {
			Type:          schema.TypeInt,
			Description:   "Number of instance running",
			Optional:      true,
			ConflictsWith: []string{"autoscale"},
		},
		"autoscale": {
			Type:          schema.TypeBool,
			Description:   "Whether the workload is auto-scaled",
			Optional:      true,
			ConflictsWith: []string{"nb_instance"},
		},
		"min_scale_instance": {
			Type:        schema.TypeInt,
			Description: "Minimum number of instances in auto-scaled mode",
			Optional:    true,
		},
		"max_scale_instance": {
			Type:        schema.TypeInt,
			Description: "Maximum number of instances in auto-scaled mode",
			Optional:    true,
		},

		// computed
		"created_at": {
			Type:        schema.TypeString,
			Description: "Input creation",
			Computed:    true,
		},
		"hostname": {
			Type:        schema.TypeString,
			Description: "Hostname",
			Computed:    true,
		},
		"input_id": {
			Type:        schema.TypeString,
			Description: "Input ID",
			Computed:    true,
		},
		"is_restart_required": {
			Type:        schema.TypeBool,
			Description: "Indicate if input need to be restarted",
			Computed:    true,
		},
		"public_address": {
			Type:        schema.TypeString,
			Description: "Input IP address",
			Computed:    true,
		},
		"ssl_certificate": {
			Type:        schema.TypeString,
			Description: "Input SSL certificate",
			Computed:    true,
			Sensitive:   true,
		},
		"status": {
			Type:        schema.TypeString,
			Description: "init: configuration required, pending: ready to start, running: available",
			Computed:    true,
		},
		"updated_at": {
			Type:        schema.TypeString,
			Description: "Input last update",
			Computed:    true,
		},
		"current_nb_instance": {
			Type:        schema.TypeInt,
			Description: "Number of instance running (returned by the API)",
			Computed:    true,
		},
	}

	return schema
}

func resourceDbaasLogsInputImportState(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	givenID := d.Id()
	splitID := strings.SplitN(givenID, "/", 2)
	if len(splitID) != 2 {
		return nil, fmt.Errorf("import ID is not service_name/id formatted")
	}
	serviceName := splitID[0]
	id := splitID[1]
	d.SetId(id)
	d.Set("service_name", serviceName)

	results := make([]*schema.ResourceData, 1)
	results[0] = d
	return results, nil
}

func resourceDbaasLogsInputCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)

	serviceName := d.Get("service_name").(string)

	log.Printf("[INFO] Will create dbaas logs input for: %s", serviceName)

	opts := (&DbaasLogsInputOpts{}).FromResource(d)
	res := &DbaasLogsOperation{}
	endpoint := fmt.Sprintf(
		"/dbaas/logs/%s/input",
		url.PathEscape(serviceName),
	)
	if err := config.OVHClient.Post(endpoint, opts, res); err != nil {
		return diag.Errorf("Error calling post %s:\n\t %q", endpoint, err)
	}

	// Wait for operation status
	op, err := waitForDbaasLogsOperation(ctx, config.OVHClient, serviceName, res.OperationId)
	if err != nil {
		return diag.FromErr(err)
	}

	id := op.InputId
	if id == nil {
		return diag.Errorf("Input Id is nil. This should not happen: operation is %s/%s", serviceName, res.OperationId)
	}

	d.SetId(*id)

	if err := dbaasLogsInputConfigurationUpdate(ctx, d, meta); err != nil {
		return diag.FromErr(err)
	}

	if err := dbaasLogsInputStart(ctx, d, meta); err != nil {
		return diag.FromErr(err)
	}

	return resourceDbaasLogsInputUpdate(ctx, d, meta)

}

func resourceDbaasLogsInputUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)

	serviceName := d.Get("service_name").(string)
	id := d.Id()

	log.Printf("[INFO] Will update dbaas logs input for: %s", serviceName)

	opts := (&DbaasLogsInputOpts{}).FromResource(d)
	res := &DbaasLogsOperation{}
	endpoint := fmt.Sprintf(
		"/dbaas/logs/%s/input/%s",
		url.PathEscape(serviceName),
		url.PathEscape(id),
	)
	if err := config.OVHClient.Put(endpoint, opts, res); err != nil {
		return diag.Errorf("Error calling Put %s:\n\t %q", endpoint, err)
	}

	// Wait for operation status
	if _, err := waitForDbaasLogsOperation(ctx, config.OVHClient, serviceName, res.OperationId); err != nil {
		return diag.FromErr(err)
	}

	if err := dbaasLogsInputConfigurationUpdate(ctx, d, meta); err != nil {
		return diag.FromErr(err)
	}

	if err := dbaasLogsInputStart(ctx, d, meta); err != nil {
		return diag.FromErr(err)
	}

	return resourceDbaasLogsInputRead(ctx, d, meta)
}

func resourceDbaasLogsInputRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	res, err := dbaasLogsInputRead(d, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	for k, v := range res.ToMap() {
		if k != "id" {
			d.Set(k, v)
		}
	}

	if err := dbaasLogsInputConfigurationRead(d, meta); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func dbaasLogsInputRead(d *schema.ResourceData, meta interface{}) (*DbaasLogsInput, error) {
	config := meta.(*Config)

	serviceName := d.Get("service_name").(string)
	id := d.Id()

	log.Printf("[INFO] Will read dbaas logs input: %s/%s", serviceName, id)
	res := &DbaasLogsInput{}
	endpoint := fmt.Sprintf(
		"/dbaas/logs/%s/input/%s",
		url.PathEscape(serviceName),
		url.PathEscape(id),
	)
	if err := config.OVHClient.Get(endpoint, &res); err != nil {
		log.Printf("[ERROR] %s: %v", endpoint, err)
		return nil, helpers.CheckDeleted(d, err, endpoint)
	}

	return res, nil
}

func resourceDbaasLogsInputDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if err := dbaasLogsInputEnd(ctx, d, meta); err != nil {
		return diag.FromErr(err)
	}

	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	id := d.Id()

	log.Printf("[INFO] Will delete dbaas logs input: %s/%s", serviceName, id)
	res := &DbaasLogsOperation{}
	endpoint := fmt.Sprintf(
		"/dbaas/logs/%s/input/%s",
		url.PathEscape(serviceName),
		url.PathEscape(id),
	)

	if err := config.OVHClient.Delete(endpoint, res); err != nil {
		return diag.FromErr(helpers.CheckDeleted(d, err, endpoint))
	}

	// Wait for operation status
	if _, err := waitForDbaasLogsOperation(ctx, config.OVHClient, serviceName, res.OperationId); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func dbaasLogsInputConfigurationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	serviceName := d.Get("service_name").(string)
	id := d.Id()

	res := &DbaasLogsOperation{}

	flowggerConf := d.Get("configuration.0.flowgger.#").(int)
	if flowggerConf > 0 {
		log.Printf("[INFO] Will update dbaas logs input configuration flowgger for: %s/%s", serviceName, id)
		flowggerOpts := (&DbaasLogsInputConfigurationFlowgger{}).FromResourceWithPath(d, "configuration.0.flowgger.0")
		endpoint := fmt.Sprintf(
			"/dbaas/logs/%s/input/%s/configuration/flowgger",
			url.PathEscape(serviceName),
			url.PathEscape(id),
		)
		if err := config.OVHClient.Put(endpoint, flowggerOpts, res); err != nil {
			return fmt.Errorf("error calling Put %s:\n\t %q", endpoint, err)
		}
	}

	logstashConf := d.Get("configuration.0.logstash.#").(int)
	if logstashConf > 0 {
		log.Printf("[INFO] Will update dbaas logs input configuration logstash for: %s/%s", serviceName, id)
		logstashOpts := (&DbaasLogsInputConfigurationLogstash{}).FromResourceWithPath(d, "configuration.0.logstash.0")
		endpoint := fmt.Sprintf(
			"/dbaas/logs/%s/input/%s/configuration/logstash",
			url.PathEscape(serviceName),
			url.PathEscape(id),
		)
		if err := config.OVHClient.Put(endpoint, logstashOpts, res); err != nil {
			return fmt.Errorf("error calling Put %s:\n\t %q", endpoint, err)
		}
	}

	// Wait for operation status
	if _, err := waitForDbaasLogsOperation(ctx, config.OVHClient, serviceName, res.OperationId); err != nil {
		return err
	}

	return nil
}

func dbaasLogsInputConfigurationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	serviceName := d.Get("service_name").(string)
	id := d.Id()
	configuration := map[string][]interface{}{}

	flowggerConf := d.Get("configuration.0.flowgger.#").(int)
	if flowggerConf > 0 {
		log.Printf("[INFO] Will read dbaas logs input configuration flowgger: %s/%s", serviceName, id)
		flowgger := &DbaasLogsInputConfigurationFlowgger{}
		endpoint := fmt.Sprintf(
			"/dbaas/logs/%s/input/%s/configuration/flowgger",
			url.PathEscape(serviceName),
			url.PathEscape(id),
		)
		if err := config.OVHClient.Get(endpoint, flowgger); err != nil {
			log.Printf("[ERROR] %s: %v", endpoint, err)
			return helpers.CheckDeleted(d, err, endpoint)
		}

		configuration["flowgger"] = []interface{}{flowgger.ToMap()}
	}

	logstashConf := d.Get("configuration.0.logstash.#").(int)
	if logstashConf > 0 {
		log.Printf("[INFO] Will read dbaas logs input configuration logstash: %s/%s", serviceName, id)
		logstash := &DbaasLogsInputConfigurationLogstash{}
		endpoint := fmt.Sprintf(
			"/dbaas/logs/%s/input/%s/configuration/logstash",
			url.PathEscape(serviceName),
			url.PathEscape(id),
		)
		if err := config.OVHClient.Get(endpoint, logstash); err != nil {
			log.Printf("[ERROR] %s: %v", endpoint, err)
			return helpers.CheckDeleted(d, err, endpoint)
		}

		configuration["logstash"] = []interface{}{logstash.ToMap()}
	}

	d.Set("configuration", []interface{}{configuration})
	return nil
}

func dbaasLogsInputStart(ctx context.Context, d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	id := d.Id()

	input, err := dbaasLogsInputRead(d, meta)
	if err != nil {
		return err
	}

	if input.Status != "RUNNING" && input.Status != "PENDING" && input.Status != "INIT" {
		return fmt.Errorf("start/restart: input %s/%s is in wrong state : %s",
			serviceName,
			id,
			input.Status,
		)
	}

	log.Printf("[INFO] Will restart dbaas logs input for: %s/%s", serviceName, id)
	res := &DbaasLogsOperation{}

	if input.Status == "RUNNING" {
		endpoint := fmt.Sprintf(
			"/dbaas/logs/%s/input/%s/restart",
			url.PathEscape(serviceName),
			url.PathEscape(id),
		)
		if err := config.OVHClient.Post(endpoint, nil, res); err != nil {
			return fmt.Errorf("error calling Put %s:\n\t %q", endpoint, err)
		}
	}

	if input.Status == "PENDING" || input.Status == "INIT" {
		endpoint := fmt.Sprintf(
			"/dbaas/logs/%s/input/%s/start",
			url.PathEscape(serviceName),
			url.PathEscape(id),
		)
		if err := config.OVHClient.Post(endpoint, nil, res); err != nil {
			return fmt.Errorf("error calling Put %s:\n\t %q", endpoint, err)
		}
	}

	// Wait for operation status
	if _, err := waitForDbaasLogsOperation(ctx, config.OVHClient, serviceName, res.OperationId); err != nil {
		return err
	}

	return nil
}

func dbaasLogsInputEnd(ctx context.Context, d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	id := d.Id()

	input, err := dbaasLogsInputRead(d, meta)
	if err != nil {
		return err
	}

	if input.Status == "PROCESSING" {
		return fmt.Errorf("stop: input %s/%s already has an ongoing action",
			serviceName,
			id,
		)
	}
	if input.Status != "RUNNING" {
		log.Printf("[DEBUG] end: input is not running for: %s/%s", serviceName, id)
		return nil
	}

	log.Printf("[INFO] Will end dbaas logs input for: %s/%s", serviceName, id)
	res := &DbaasLogsOperation{}
	endpoint := fmt.Sprintf(
		"/dbaas/logs/%s/input/%s/end",
		url.PathEscape(serviceName),
		url.PathEscape(id),
	)
	if err := config.OVHClient.Post(endpoint, nil, res); err != nil {
		return fmt.Errorf("error calling Put %s:\n\t %q", endpoint, err)
	}

	// Wait for operation status
	if _, err := waitForDbaasLogsOperation(ctx, config.OVHClient, serviceName, res.OperationId); err != nil {
		return err
	}

	return nil
}
