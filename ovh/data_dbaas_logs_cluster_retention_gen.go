// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package ovh

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	ovhtypes "github.com/ovh/terraform-provider-ovh/v2/ovh/types"
)

func DbaasLogsClusterRetentionDataSourceSchema(ctx context.Context) schema.Schema {
	attrs := map[string]schema.Attribute{
		"cluster_id": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Cluster ID",
			MarkdownDescription: "Cluster ID",
		},
		"duration": schema.StringAttribute{
			CustomType: ovhtypes.TfStringType{},
			Computed:   true,
			Optional:   true,
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("retention_id")),
			},
			Description:         "Indexed duration expressed in ISO-8601 format",
			MarkdownDescription: "Indexed duration expressed in ISO-8601 format",
		},
		"is_supported": schema.BoolAttribute{
			CustomType:          ovhtypes.TfBoolType{},
			Computed:            true,
			Description:         "Indicates if a new stream can use it",
			MarkdownDescription: "Indicates if a new stream can use it",
		},
		"retention_id": schema.StringAttribute{
			CustomType: ovhtypes.TfStringType{},
			Computed:   true,
			Optional:   true,
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("duration")),
				stringvalidator.ConflictsWith(path.MatchRoot("retention_type")),
			},
			Description:         "Retention ID",
			MarkdownDescription: "Retention ID",
		},
		"retention_type": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Optional:            true,
			Description:         "Retention type",
			MarkdownDescription: "Retention type",
			Validators: []validator.String{
				stringvalidator.OneOf(
					"LOGS_COLD_STORAGE",
					"LOGS_INDEXING",
					"METRICS_TENANT",
				),
				stringvalidator.ConflictsWith(path.MatchRoot("retention_id")),
			},
		},
		"service_name": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Service name",
			MarkdownDescription: "Service name",
		},
	}

	return schema.Schema{
		Attributes: attrs,
	}
}

type DbaasLogsClusterRetentionModel struct {
	ClusterId     ovhtypes.TfStringValue `tfsdk:"cluster_id" json:"clusterId"`
	Duration      ovhtypes.TfStringValue `tfsdk:"duration" json:"duration"`
	IsSupported   ovhtypes.TfBoolValue   `tfsdk:"is_supported" json:"isSupported"`
	RetentionId   ovhtypes.TfStringValue `tfsdk:"retention_id" json:"retentionId"`
	RetentionType ovhtypes.TfStringValue `tfsdk:"retention_type" json:"retentionType"`
	ServiceName   ovhtypes.TfStringValue `tfsdk:"service_name" json:"serviceName"`
}

func (v *DbaasLogsClusterRetentionModel) MergeWith(other *DbaasLogsClusterRetentionModel) {
	if (v.ClusterId.IsUnknown() || v.ClusterId.IsNull()) && !other.ClusterId.IsUnknown() {
		v.ClusterId = other.ClusterId
	}

	if (v.Duration.IsUnknown() || v.Duration.IsNull()) && !other.Duration.IsUnknown() {
		v.Duration = other.Duration
	}

	if (v.IsSupported.IsUnknown() || v.IsSupported.IsNull()) && !other.IsSupported.IsUnknown() {
		v.IsSupported = other.IsSupported
	}

	if (v.RetentionId.IsUnknown() || v.RetentionId.IsNull()) && !other.RetentionId.IsUnknown() {
		v.RetentionId = other.RetentionId
	}

	if (v.RetentionType.IsUnknown() || v.RetentionType.IsNull()) && !other.RetentionType.IsUnknown() {
		v.RetentionType = other.RetentionType
	}

	if (v.ServiceName.IsUnknown() || v.ServiceName.IsNull()) && !other.ServiceName.IsUnknown() {
		v.ServiceName = other.ServiceName
	}
}
