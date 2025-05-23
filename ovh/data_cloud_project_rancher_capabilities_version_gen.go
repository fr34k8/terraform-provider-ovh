// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package ovh

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	ovhtypes "github.com/ovh/terraform-provider-ovh/v2/ovh/types"
)

func CloudProjectRancherCapabilitiesVersionDataSourceSchema(ctx context.Context) schema.Schema {
	attrs := map[string]schema.Attribute{
		"versions": schema.SetNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"cause": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Cause for an unavailability",
						MarkdownDescription: "Cause for an unavailability",
					},
					"changelog_url": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Changelog URL of the version",
						MarkdownDescription: "Changelog URL of the version",
					},
					"message": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Human-readable description of the unavailability cause",
						MarkdownDescription: "Human-readable description of the unavailability cause",
					},
					"name": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Name of the version",
						MarkdownDescription: "Name of the version",
					},
					"status": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Status of the version",
						MarkdownDescription: "Status of the version",
					},
				},
				CustomType: CloudProjectRancherCapabilitiesVersionType{
					ObjectType: types.ObjectType{
						AttrTypes: CloudProjectRancherCapabilitiesVersionValue{}.AttributeTypes(ctx),
					},
				},
			},
			CustomType: ovhtypes.NewTfListNestedType[CloudProjectRancherCapabilitiesVersionValue](ctx),
			Computed:   true,
		},
		"project_id": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Project ID",
			MarkdownDescription: "Project ID",
		},
		"rancher_id": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Rancher ID",
			MarkdownDescription: "Rancher ID",
		},
	}

	return schema.Schema{
		Description: "List available and current versions for the given managed Rancher service",
		Attributes:  attrs,
	}
}

type CloudProjectRancherCapabilitiesVersionModel struct {
	Versions  ovhtypes.TfListNestedValue[CloudProjectRancherCapabilitiesVersionValue] `tfsdk:"versions" json:"versions"`
	ProjectId ovhtypes.TfStringValue                                                  `tfsdk:"project_id" json:"projectId"`
	RancherId ovhtypes.TfStringValue                                                  `tfsdk:"rancher_id" json:"rancherId"`
}

func (v *CloudProjectRancherCapabilitiesVersionModel) MergeWith(other *CloudProjectRancherCapabilitiesVersionModel) {

	if (v.Versions.IsUnknown() || v.Versions.IsNull()) && !other.Versions.IsUnknown() {
		v.Versions = other.Versions
	}

	if (v.ProjectId.IsUnknown() || v.ProjectId.IsNull()) && !other.ProjectId.IsUnknown() {
		v.ProjectId = other.ProjectId
	}

	if (v.RancherId.IsUnknown() || v.RancherId.IsNull()) && !other.RancherId.IsUnknown() {
		v.RancherId = other.RancherId
	}

}

var _ basetypes.ObjectTypable = CloudProjectRancherCapabilitiesVersionType{}

type CloudProjectRancherCapabilitiesVersionType struct {
	basetypes.ObjectType
}

func (t CloudProjectRancherCapabilitiesVersionType) Equal(o attr.Type) bool {
	other, ok := o.(CloudProjectRancherCapabilitiesVersionType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t CloudProjectRancherCapabilitiesVersionType) String() string {
	return "CloudProjectRancherCapabilitiesVersionType"
}

func (t CloudProjectRancherCapabilitiesVersionType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	causeAttribute, ok := attributes["cause"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`cause is missing from object`)

		return nil, diags
	}

	causeVal, ok := causeAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`cause expected to be ovhtypes.TfStringValue, was: %T`, causeAttribute))
	}

	changelogUrlAttribute, ok := attributes["changelog_url"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`changelog_url is missing from object`)

		return nil, diags
	}

	changelogUrlVal, ok := changelogUrlAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`changelog_url expected to be ovhtypes.TfStringValue, was: %T`, changelogUrlAttribute))
	}

	messageAttribute, ok := attributes["message"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`message is missing from object`)

		return nil, diags
	}

	messageVal, ok := messageAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`message expected to be ovhtypes.TfStringValue, was: %T`, messageAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return nil, diags
	}

	nameVal, ok := nameAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be ovhtypes.TfStringValue, was: %T`, nameAttribute))
	}

	statusAttribute, ok := attributes["status"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`status is missing from object`)

		return nil, diags
	}

	statusVal, ok := statusAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`status expected to be ovhtypes.TfStringValue, was: %T`, statusAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return CloudProjectRancherCapabilitiesVersionValue{
		Cause:        causeVal,
		ChangelogUrl: changelogUrlVal,
		Message:      messageVal,
		Name:         nameVal,
		Status:       statusVal,
		state:        attr.ValueStateKnown,
	}, diags
}

func NewCloudProjectRancherCapabilitiesVersionValueNull() CloudProjectRancherCapabilitiesVersionValue {
	return CloudProjectRancherCapabilitiesVersionValue{
		state: attr.ValueStateNull,
	}
}

func NewCloudProjectRancherCapabilitiesVersionValueUnknown() CloudProjectRancherCapabilitiesVersionValue {
	return CloudProjectRancherCapabilitiesVersionValue{
		state: attr.ValueStateUnknown,
	}
}

func NewCloudProjectRancherCapabilitiesVersionValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (CloudProjectRancherCapabilitiesVersionValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing CloudProjectRancherCapabilitiesVersionValue Attribute Value",
				"While creating a CloudProjectRancherCapabilitiesVersionValue value, a missing attribute value was detected. "+
					"A CloudProjectRancherCapabilitiesVersionValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("CloudProjectRancherCapabilitiesVersionValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid CloudProjectRancherCapabilitiesVersionValue Attribute Type",
				"While creating a CloudProjectRancherCapabilitiesVersionValue value, an invalid attribute value was detected. "+
					"A CloudProjectRancherCapabilitiesVersionValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("CloudProjectRancherCapabilitiesVersionValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("CloudProjectRancherCapabilitiesVersionValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra CloudProjectRancherCapabilitiesVersionValue Attribute Value",
				"While creating a CloudProjectRancherCapabilitiesVersionValue value, an extra attribute value was detected. "+
					"A CloudProjectRancherCapabilitiesVersionValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra CloudProjectRancherCapabilitiesVersionValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewCloudProjectRancherCapabilitiesVersionValueUnknown(), diags
	}

	causeAttribute, ok := attributes["cause"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`cause is missing from object`)

		return NewCloudProjectRancherCapabilitiesVersionValueUnknown(), diags
	}

	causeVal, ok := causeAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`cause expected to be ovhtypes.TfStringValue, was: %T`, causeAttribute))
	}

	changelogUrlAttribute, ok := attributes["changelog_url"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`changelog_url is missing from object`)

		return NewCloudProjectRancherCapabilitiesVersionValueUnknown(), diags
	}

	changelogUrlVal, ok := changelogUrlAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`changelog_url expected to be ovhtypes.TfStringValue, was: %T`, changelogUrlAttribute))
	}

	messageAttribute, ok := attributes["message"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`message is missing from object`)

		return NewCloudProjectRancherCapabilitiesVersionValueUnknown(), diags
	}

	messageVal, ok := messageAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`message expected to be ovhtypes.TfStringValue, was: %T`, messageAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return NewCloudProjectRancherCapabilitiesVersionValueUnknown(), diags
	}

	nameVal, ok := nameAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be ovhtypes.TfStringValue, was: %T`, nameAttribute))
	}

	statusAttribute, ok := attributes["status"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`status is missing from object`)

		return NewCloudProjectRancherCapabilitiesVersionValueUnknown(), diags
	}

	statusVal, ok := statusAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`status expected to be ovhtypes.TfStringValue, was: %T`, statusAttribute))
	}

	if diags.HasError() {
		return NewCloudProjectRancherCapabilitiesVersionValueUnknown(), diags
	}

	return CloudProjectRancherCapabilitiesVersionValue{
		Cause:        causeVal,
		ChangelogUrl: changelogUrlVal,
		Message:      messageVal,
		Name:         nameVal,
		Status:       statusVal,
		state:        attr.ValueStateKnown,
	}, diags
}

func NewCloudProjectRancherCapabilitiesVersionValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) CloudProjectRancherCapabilitiesVersionValue {
	object, diags := NewCloudProjectRancherCapabilitiesVersionValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewCloudProjectRancherCapabilitiesVersionValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t CloudProjectRancherCapabilitiesVersionType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewCloudProjectRancherCapabilitiesVersionValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewCloudProjectRancherCapabilitiesVersionValueUnknown(), nil
	}

	if in.IsNull() {
		return NewCloudProjectRancherCapabilitiesVersionValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewCloudProjectRancherCapabilitiesVersionValueMust(CloudProjectRancherCapabilitiesVersionValue{}.AttributeTypes(ctx), attributes), nil
}

func (t CloudProjectRancherCapabilitiesVersionType) ValueType(ctx context.Context) attr.Value {
	return CloudProjectRancherCapabilitiesVersionValue{}
}

var _ basetypes.ObjectValuable = CloudProjectRancherCapabilitiesVersionValue{}

type CloudProjectRancherCapabilitiesVersionValue struct {
	Cause        ovhtypes.TfStringValue `tfsdk:"cause" json:"cause"`
	ChangelogUrl ovhtypes.TfStringValue `tfsdk:"changelog_url" json:"changelogUrl"`
	Message      ovhtypes.TfStringValue `tfsdk:"message" json:"message"`
	Name         ovhtypes.TfStringValue `tfsdk:"name" json:"name"`
	Status       ovhtypes.TfStringValue `tfsdk:"status" json:"status"`
	state        attr.ValueState
}

func (v *CloudProjectRancherCapabilitiesVersionValue) UnmarshalJSON(data []byte) error {
	type JsonCloudProjectRancherCapabilitiesVersionValue CloudProjectRancherCapabilitiesVersionValue

	var tmp JsonCloudProjectRancherCapabilitiesVersionValue
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	v.Cause = tmp.Cause
	v.ChangelogUrl = tmp.ChangelogUrl
	v.Message = tmp.Message
	v.Name = tmp.Name
	v.Status = tmp.Status

	v.state = attr.ValueStateKnown

	return nil
}

func (v *CloudProjectRancherCapabilitiesVersionValue) MergeWith(other *CloudProjectRancherCapabilitiesVersionValue) {

	if (v.Cause.IsUnknown() || v.Cause.IsNull()) && !other.Cause.IsUnknown() {
		v.Cause = other.Cause
	}

	if (v.ChangelogUrl.IsUnknown() || v.ChangelogUrl.IsNull()) && !other.ChangelogUrl.IsUnknown() {
		v.ChangelogUrl = other.ChangelogUrl
	}

	if (v.Message.IsUnknown() || v.Message.IsNull()) && !other.Message.IsUnknown() {
		v.Message = other.Message
	}

	if (v.Name.IsUnknown() || v.Name.IsNull()) && !other.Name.IsUnknown() {
		v.Name = other.Name
	}

	if (v.Status.IsUnknown() || v.Status.IsNull()) && !other.Status.IsUnknown() {
		v.Status = other.Status
	}

	if (v.state == attr.ValueStateUnknown || v.state == attr.ValueStateNull) && other.state != attr.ValueStateUnknown {
		v.state = other.state
	}
}

func (v CloudProjectRancherCapabilitiesVersionValue) Attributes() map[string]attr.Value {
	return map[string]attr.Value{
		"cause":        v.Cause,
		"changelogUrl": v.ChangelogUrl,
		"message":      v.Message,
		"name":         v.Name,
		"status":       v.Status,
	}
}
func (v CloudProjectRancherCapabilitiesVersionValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 5)

	var val tftypes.Value
	var err error

	attrTypes["cause"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["changelog_url"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["message"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["name"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["status"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 5)

		val, err = v.Cause.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["cause"] = val

		val, err = v.ChangelogUrl.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["changelog_url"] = val

		val, err = v.Message.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["message"] = val

		val, err = v.Name.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["name"] = val

		val, err = v.Status.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["status"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v CloudProjectRancherCapabilitiesVersionValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v CloudProjectRancherCapabilitiesVersionValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v CloudProjectRancherCapabilitiesVersionValue) String() string {
	return "CloudProjectRancherCapabilitiesVersionValue"
}

func (v CloudProjectRancherCapabilitiesVersionValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	objVal, diags := types.ObjectValue(
		map[string]attr.Type{
			"cause":         ovhtypes.TfStringType{},
			"changelog_url": ovhtypes.TfStringType{},
			"message":       ovhtypes.TfStringType{},
			"name":          ovhtypes.TfStringType{},
			"status":        ovhtypes.TfStringType{},
		},
		map[string]attr.Value{
			"cause":         v.Cause,
			"changelog_url": v.ChangelogUrl,
			"message":       v.Message,
			"name":          v.Name,
			"status":        v.Status,
		})

	return objVal, diags
}

func (v CloudProjectRancherCapabilitiesVersionValue) Equal(o attr.Value) bool {
	other, ok := o.(CloudProjectRancherCapabilitiesVersionValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Cause.Equal(other.Cause) {
		return false
	}

	if !v.ChangelogUrl.Equal(other.ChangelogUrl) {
		return false
	}

	if !v.Message.Equal(other.Message) {
		return false
	}

	if !v.Name.Equal(other.Name) {
		return false
	}

	if !v.Status.Equal(other.Status) {
		return false
	}

	return true
}

func (v CloudProjectRancherCapabilitiesVersionValue) Type(ctx context.Context) attr.Type {
	return CloudProjectRancherCapabilitiesVersionType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v CloudProjectRancherCapabilitiesVersionValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"cause":         ovhtypes.TfStringType{},
		"changelog_url": ovhtypes.TfStringType{},
		"message":       ovhtypes.TfStringType{},
		"name":          ovhtypes.TfStringType{},
		"status":        ovhtypes.TfStringType{},
	}
}
