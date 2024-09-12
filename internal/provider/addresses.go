package provider

// Note:  This file is automatically generated.  Manually made changes
// will be overwritten when the provider is generated.

import (
	"context"
	"errors"
	"fmt"

	"github.com/PaloAltoNetworks/pango"
	sdkerrors "github.com/PaloAltoNetworks/pango/errors"
	"github.com/PaloAltoNetworks/pango/objects/address"

	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	rsschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	sdkmanager "github.com/PaloAltoNetworks/terraform-provider-panos/internal/manager"
)

// Generate Terraform Data Source object.
var (
	_ datasource.DataSource              = &AddressesDataSource{}
	_ datasource.DataSourceWithConfigure = &AddressesDataSource{}
)

func NewAddressesDataSource() datasource.DataSource {
	return &AddressesDataSource{}
}

type AddressesDataSource struct {
	client  *pango.Client
	manager *sdkmanager.EntryObjectManager[*address.Entry, address.Location, *address.Service]
}

type AddressesDataSourceFilter struct {
	// TODO: Generate Data Source filter via function
}
type AddressesDataSourceTfid struct {
	Name     string           `json:"name"`
	Location address.Location `json:"location"`
}

func (o *AddressesDataSourceTfid) IsValid() error {
	if o.Name == "" {
		return fmt.Errorf("name is unspecified")
	}
	return o.Location.IsValid()
}

type AddressesDataSourceModel struct {
	Location  AddressesLocation `tfsdk:"location"`
	Addresses types.Map         `tfsdk:"addresses"`
}
type AddressesDataSourceAddressesObject struct {
	Description types.String `tfsdk:"description"`
	Tags        types.List   `tfsdk:"tags"`
	Fqdn        types.String `tfsdk:"fqdn"`
	IpWildcard  types.String `tfsdk:"ip_wildcard"`
	IpNetmask   types.String `tfsdk:"ip_netmask"`
	IpRange     types.String `tfsdk:"ip_range"`
}

func (o *AddressesDataSourceAddressesObject) CopyToPango(ctx context.Context, obj **address.Entry, encrypted *map[string]types.String) diag.Diagnostics {
	var diags diag.Diagnostics
	description_value := o.Description.ValueStringPointer()
	tags_pango_entries := make([]string, 0)
	diags.Append(o.Tags.ElementsAs(ctx, &tags_pango_entries, false)...)
	if diags.HasError() {
		return diags
	}
	ipNetmask_value := o.IpNetmask.ValueStringPointer()
	ipRange_value := o.IpRange.ValueStringPointer()
	fqdn_value := o.Fqdn.ValueStringPointer()
	ipWildcard_value := o.IpWildcard.ValueStringPointer()

	if (*obj) == nil {
		*obj = new(address.Entry)
	}
	(*obj).Description = description_value
	(*obj).Tags = tags_pango_entries
	(*obj).IpNetmask = ipNetmask_value
	(*obj).IpRange = ipRange_value
	(*obj).Fqdn = fqdn_value
	(*obj).IpWildcard = ipWildcard_value

	return diags
}

func (o *AddressesDataSourceAddressesObject) CopyFromPango(ctx context.Context, obj *address.Entry, encrypted *map[string]types.String) diag.Diagnostics {
	var diags diag.Diagnostics
	var tags_list types.List
	{
		var list_diags diag.Diagnostics
		tags_list, list_diags = types.ListValueFrom(ctx, types.StringType, obj.Tags)
		diags.Append(list_diags...)
	}
	var description_value types.String
	if obj.Description != nil {
		description_value = types.StringValue(*obj.Description)
	}
	var ipRange_value types.String
	if obj.IpRange != nil {
		ipRange_value = types.StringValue(*obj.IpRange)
	}
	var fqdn_value types.String
	if obj.Fqdn != nil {
		fqdn_value = types.StringValue(*obj.Fqdn)
	}
	var ipWildcard_value types.String
	if obj.IpWildcard != nil {
		ipWildcard_value = types.StringValue(*obj.IpWildcard)
	}
	var ipNetmask_value types.String
	if obj.IpNetmask != nil {
		ipNetmask_value = types.StringValue(*obj.IpNetmask)
	}
	o.Tags = tags_list
	o.Description = description_value
	o.IpRange = ipRange_value
	o.Fqdn = fqdn_value
	o.IpWildcard = ipWildcard_value
	o.IpNetmask = ipNetmask_value

	return diags
}

func AddressesDataSourceSchema() dsschema.Schema {
	return dsschema.Schema{
		Attributes: map[string]dsschema.Attribute{

			"location": AddressesDataSourceLocationSchema(),

			"addresses": dsschema.MapNestedAttribute{
				Description:  "",
				Required:     true,
				Optional:     false,
				Computed:     false,
				Sensitive:    false,
				NestedObject: AddressesDataSourceAddressesSchema(),
			},
		},
	}
}

func (o *AddressesDataSourceModel) getTypeFor(name string) attr.Type {
	schema := AddressesDataSourceSchema()
	if attr, ok := schema.Attributes[name]; !ok {
		panic(fmt.Sprintf("could not resolve schema for attribute %s", name))
	} else {
		switch attr := attr.(type) {
		case dsschema.ListNestedAttribute:
			return attr.NestedObject.Type()
		case dsschema.MapNestedAttribute:
			return attr.NestedObject.Type()
		default:
			return attr.GetType()
		}
	}

	panic("unreachable")
}

func AddressesDataSourceAddressesSchema() dsschema.NestedAttributeObject {
	return dsschema.NestedAttributeObject{
		Attributes: map[string]dsschema.Attribute{

			"description": dsschema.StringAttribute{
				Description: "The description.",
				Computed:    true,
				Required:    false,
				Optional:    true,
				Sensitive:   false,
			},

			"tags": dsschema.ListAttribute{
				Description: "The administrative tags.",
				Required:    false,
				Optional:    true,
				Computed:    true,
				Sensitive:   false,
				ElementType: types.StringType,
			},

			"ip_netmask": dsschema.StringAttribute{
				Description: "The IP netmask value.",
				Computed:    true,
				Required:    false,
				Optional:    true,
				Sensitive:   false,

				Validators: []validator.String{
					stringvalidator.ExactlyOneOf(path.Expressions{
						path.MatchRelative().AtParent().AtName("ip_netmask"),
						path.MatchRelative().AtParent().AtName("ip_range"),
						path.MatchRelative().AtParent().AtName("fqdn"),
						path.MatchRelative().AtParent().AtName("ip_wildcard"),
					}...),
				},
			},

			"ip_range": dsschema.StringAttribute{
				Description: "The IP range value.",
				Computed:    true,
				Required:    false,
				Optional:    true,
				Sensitive:   false,
			},

			"fqdn": dsschema.StringAttribute{
				Description: "The FQDN value.",
				Computed:    true,
				Required:    false,
				Optional:    true,
				Sensitive:   false,
			},

			"ip_wildcard": dsschema.StringAttribute{
				Description: "The IP wildcard value.",
				Computed:    true,
				Required:    false,
				Optional:    true,
				Sensitive:   false,
			},
		},
	}
}

func (o *AddressesDataSourceAddressesObject) getTypeFor(name string) attr.Type {
	schema := AddressesDataSourceAddressesSchema()
	if attr, ok := schema.Attributes[name]; !ok {
		panic(fmt.Sprintf("could not resolve schema for attribute %s", name))
	} else {
		switch attr := attr.(type) {
		case dsschema.ListNestedAttribute:
			return attr.NestedObject.Type()
		case dsschema.MapNestedAttribute:
			return attr.NestedObject.Type()
		default:
			return attr.GetType()
		}
	}

	panic("unreachable")
}

func AddressesDataSourceLocationSchema() rsschema.Attribute {
	return AddressesLocationSchema()
}

// Metadata returns the data source type name.
func (d *AddressesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_addresses"
}

// Schema defines the schema for this data source.
func (d *AddressesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AddressesDataSourceSchema()
}

// Configure prepares the struct.
func (d *AddressesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*pango.Client)
	specifier, _, err := address.Versioning(d.client.Versioning())
	if err != nil {
		resp.Diagnostics.AddError("Failed to configure SDK client", err.Error())
		return
	}
	d.manager = sdkmanager.NewEntryObjectManager(d.client, address.NewService(d.client), specifier, address.SpecMatches)
}

func (o *AddressesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state AddressesDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing resource create", map[string]any{
		"resource_name": "panos_addresses",
		"function":      "Create",
	})

	var location address.Location

	if !state.Location.Shared.IsNull() && state.Location.Shared.ValueBool() {
		location.Shared = true
	}
	if state.Location.Vsys != nil {
		location.Vsys = &address.VsysLocation{

			NgfwDevice: state.Location.Vsys.NgfwDevice.ValueString(),
			Vsys:       state.Location.Vsys.Name.ValueString(),
		}
	}
	if !state.Location.FromPanoramaShared.IsNull() && state.Location.FromPanoramaShared.ValueBool() {
		location.FromPanoramaShared = true
	}
	if state.Location.FromPanoramaVsys != nil {
		location.FromPanoramaVsys = &address.FromPanoramaVsysLocation{

			Vsys: state.Location.FromPanoramaVsys.Vsys.ValueString(),
		}
	}
	if state.Location.DeviceGroup != nil {
		location.DeviceGroup = &address.DeviceGroupLocation{

			PanoramaDevice: state.Location.DeviceGroup.PanoramaDevice.ValueString(),
			DeviceGroup:    state.Location.DeviceGroup.Name.ValueString(),
		}
	}

	elements := make(map[string]AddressesDataSourceAddressesObject)
	resp.Diagnostics.Append(state.Addresses.ElementsAs(ctx, &elements, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	entries := make([]*address.Entry, 0, len(elements))
	for name, elt := range elements {
		var entry *address.Entry
		resp.Diagnostics.Append(elt.CopyToPango(ctx, &entry, nil)...)
		if resp.Diagnostics.HasError() {
			return
		}
		entry.Name = name
		entries = append(entries, entry)
	}

	readEntries, err := o.manager.ReadMany(ctx, location, entries)
	if err != nil {
		if errors.Is(err, sdkmanager.ErrObjectNotFound) {
			resp.State.RemoveResource(ctx)
		} else {
			resp.Diagnostics.AddError("Failed to read entries from the server", err.Error())
		}
		return
	}

	objects := make(map[string]AddressesDataSourceAddressesObject)
	for _, elt := range readEntries {
		var object AddressesDataSourceAddressesObject
		resp.Diagnostics.Append(object.CopyFromPango(ctx, elt, nil)...)
		if resp.Diagnostics.HasError() {
			return
		}
		objects[elt.Name] = object
	}

	var map_diags diag.Diagnostics
	state.Addresses, map_diags = types.MapValueFrom(ctx, state.getTypeFor("addresses"), objects)
	resp.Diagnostics.Append(map_diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)

}

// Generate Terraform Resource object
var (
	_ resource.Resource                = &AddressesResource{}
	_ resource.ResourceWithConfigure   = &AddressesResource{}
	_ resource.ResourceWithImportState = &AddressesResource{}
)

func NewAddressesResource() resource.Resource {
	return &AddressesResource{}
}

type AddressesResource struct {
	client  *pango.Client
	manager *sdkmanager.EntryObjectManager[*address.Entry, address.Location, *address.Service]
}
type AddressesResourceTfid struct {
	Name     string           `json:"name"`
	Location address.Location `json:"location"`
}

func (o *AddressesResourceTfid) IsValid() error {
	if o.Name == "" {
		return fmt.Errorf("name is unspecified")
	}
	return o.Location.IsValid()
}

func AddressesResourceLocationSchema() rsschema.Attribute {
	return AddressesLocationSchema()
}

type AddressesResourceModel struct {
	Location  AddressesLocation `tfsdk:"location"`
	Addresses types.Map         `tfsdk:"addresses"`
}
type AddressesResourceAddressesObject struct {
	Description types.String `tfsdk:"description"`
	Tags        types.List   `tfsdk:"tags"`
	IpNetmask   types.String `tfsdk:"ip_netmask"`
	IpRange     types.String `tfsdk:"ip_range"`
	Fqdn        types.String `tfsdk:"fqdn"`
	IpWildcard  types.String `tfsdk:"ip_wildcard"`
}

func (r *AddressesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_addresses"
}

func (r *AddressesResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
}

// <ResourceSchema>

func AddressesResourceSchema() rsschema.Schema {
	return rsschema.Schema{
		Attributes: map[string]rsschema.Attribute{

			"location": AddressesResourceLocationSchema(),

			"addresses": rsschema.MapNestedAttribute{
				Description:  "",
				Required:     true,
				Optional:     false,
				Computed:     false,
				Sensitive:    false,
				NestedObject: AddressesResourceAddressesSchema(),
			},
		},
	}
}

func (o *AddressesResourceModel) getTypeFor(name string) attr.Type {
	schema := AddressesResourceSchema()
	if attr, ok := schema.Attributes[name]; !ok {
		panic(fmt.Sprintf("could not resolve schema for attribute %s", name))
	} else {
		switch attr := attr.(type) {
		case rsschema.ListNestedAttribute:
			return attr.NestedObject.Type()
		case rsschema.MapNestedAttribute:
			return attr.NestedObject.Type()
		default:
			return attr.GetType()
		}
	}

	panic("unreachable")
}

func AddressesResourceAddressesSchema() rsschema.NestedAttributeObject {
	return rsschema.NestedAttributeObject{
		Attributes: map[string]rsschema.Attribute{

			"description": rsschema.StringAttribute{
				Description: "The description.",
				Computed:    false,
				Required:    false,
				Optional:    true,
				Sensitive:   false,
			},

			"tags": rsschema.ListAttribute{
				Description: "The administrative tags.",
				Required:    false,
				Optional:    true,
				Computed:    false,
				Sensitive:   false,
				ElementType: types.StringType,
			},

			"ip_wildcard": rsschema.StringAttribute{
				Description: "The IP wildcard value.",
				Computed:    false,
				Required:    false,
				Optional:    true,
				Sensitive:   false,

				Validators: []validator.String{
					stringvalidator.ExactlyOneOf(path.Expressions{
						path.MatchRelative().AtParent().AtName("fqdn"),
						path.MatchRelative().AtParent().AtName("ip_wildcard"),
						path.MatchRelative().AtParent().AtName("ip_netmask"),
						path.MatchRelative().AtParent().AtName("ip_range"),
					}...),
				},
			},

			"ip_netmask": rsschema.StringAttribute{
				Description: "The IP netmask value.",
				Computed:    false,
				Required:    false,
				Optional:    true,
				Sensitive:   false,
			},

			"ip_range": rsschema.StringAttribute{
				Description: "The IP range value.",
				Computed:    false,
				Required:    false,
				Optional:    true,
				Sensitive:   false,
			},

			"fqdn": rsschema.StringAttribute{
				Description: "The FQDN value.",
				Computed:    false,
				Required:    false,
				Optional:    true,
				Sensitive:   false,
			},
		},
	}
}

func (o *AddressesResourceAddressesObject) getTypeFor(name string) attr.Type {
	schema := AddressesResourceAddressesSchema()
	if attr, ok := schema.Attributes[name]; !ok {
		panic(fmt.Sprintf("could not resolve schema for attribute %s", name))
	} else {
		switch attr := attr.(type) {
		case rsschema.ListNestedAttribute:
			return attr.NestedObject.Type()
		case rsschema.MapNestedAttribute:
			return attr.NestedObject.Type()
		default:
			return attr.GetType()
		}
	}

	panic("unreachable")
}

func (r *AddressesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = AddressesResourceSchema()
}

// </ResourceSchema>

func (r *AddressesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*pango.Client)
	specifier, _, err := address.Versioning(r.client.Versioning())
	if err != nil {
		resp.Diagnostics.AddError("Failed to configure SDK client", err.Error())
		return
	}
	r.manager = sdkmanager.NewEntryObjectManager(r.client, address.NewService(r.client), specifier, address.SpecMatches)
}

func (o *AddressesResourceAddressesObject) CopyToPango(ctx context.Context, obj **address.Entry, encrypted *map[string]types.String) diag.Diagnostics {
	var diags diag.Diagnostics
	description_value := o.Description.ValueStringPointer()
	tags_pango_entries := make([]string, 0)
	diags.Append(o.Tags.ElementsAs(ctx, &tags_pango_entries, false)...)
	if diags.HasError() {
		return diags
	}
	ipNetmask_value := o.IpNetmask.ValueStringPointer()
	ipRange_value := o.IpRange.ValueStringPointer()
	fqdn_value := o.Fqdn.ValueStringPointer()
	ipWildcard_value := o.IpWildcard.ValueStringPointer()

	if (*obj) == nil {
		*obj = new(address.Entry)
	}
	(*obj).Description = description_value
	(*obj).Tags = tags_pango_entries
	(*obj).IpNetmask = ipNetmask_value
	(*obj).IpRange = ipRange_value
	(*obj).Fqdn = fqdn_value
	(*obj).IpWildcard = ipWildcard_value

	return diags
}

func (o *AddressesResourceAddressesObject) CopyFromPango(ctx context.Context, obj *address.Entry, encrypted *map[string]types.String) diag.Diagnostics {
	var diags diag.Diagnostics
	var tags_list types.List
	{
		var list_diags diag.Diagnostics
		tags_list, list_diags = types.ListValueFrom(ctx, types.StringType, obj.Tags)
		diags.Append(list_diags...)
	}
	var description_value types.String
	if obj.Description != nil {
		description_value = types.StringValue(*obj.Description)
	}
	var ipNetmask_value types.String
	if obj.IpNetmask != nil {
		ipNetmask_value = types.StringValue(*obj.IpNetmask)
	}
	var ipRange_value types.String
	if obj.IpRange != nil {
		ipRange_value = types.StringValue(*obj.IpRange)
	}
	var fqdn_value types.String
	if obj.Fqdn != nil {
		fqdn_value = types.StringValue(*obj.Fqdn)
	}
	var ipWildcard_value types.String
	if obj.IpWildcard != nil {
		ipWildcard_value = types.StringValue(*obj.IpWildcard)
	}
	o.Description = description_value
	o.Tags = tags_list
	o.IpNetmask = ipNetmask_value
	o.IpRange = ipRange_value
	o.Fqdn = fqdn_value
	o.IpWildcard = ipWildcard_value

	return diags
}

func (r *AddressesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var state AddressesResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing resource create", map[string]any{
		"resource_name": "panos_addresses_resource",
		"function":      "Create",
	})

	var location address.Location

	if !state.Location.Shared.IsNull() && state.Location.Shared.ValueBool() {
		location.Shared = true
	}
	if state.Location.Vsys != nil {
		location.Vsys = &address.VsysLocation{

			NgfwDevice: state.Location.Vsys.NgfwDevice.ValueString(),
			Vsys:       state.Location.Vsys.Name.ValueString(),
		}
	}
	if !state.Location.FromPanoramaShared.IsNull() && state.Location.FromPanoramaShared.ValueBool() {
		location.FromPanoramaShared = true
	}
	if state.Location.FromPanoramaVsys != nil {
		location.FromPanoramaVsys = &address.FromPanoramaVsysLocation{

			Vsys: state.Location.FromPanoramaVsys.Vsys.ValueString(),
		}
	}
	if state.Location.DeviceGroup != nil {
		location.DeviceGroup = &address.DeviceGroupLocation{

			DeviceGroup:    state.Location.DeviceGroup.Name.ValueString(),
			PanoramaDevice: state.Location.DeviceGroup.PanoramaDevice.ValueString(),
		}
	}

	type entryWithState struct {
		Entry    *address.Entry
		StateIdx int
	}

	var elements map[string]AddressesResourceAddressesObject
	state.Addresses.ElementsAs(ctx, &elements, false)
	entries := make([]*address.Entry, len(elements))
	idx := 0
	for name, elt := range elements {
		var entry *address.Entry
		resp.Diagnostics.Append(elt.CopyToPango(ctx, &entry, nil)...)
		if resp.Diagnostics.HasError() {
			return
		}
		entry.Name = name
		entries[idx] = entry
		idx++
	}

	created, err := r.manager.CreateMany(ctx, location, entries)
	if err != nil {
		resp.Diagnostics.AddError("Failed to create new entries", err.Error())
		return
	}

	for _, elt := range created {
		if _, found := elements[elt.Name]; !found {
			continue
		}
		var object AddressesResourceAddressesObject
		resp.Diagnostics.Append(object.CopyFromPango(ctx, elt, nil)...)
		if resp.Diagnostics.HasError() {
			return
		}
		elements[elt.Name] = object
	}

	var map_diags diag.Diagnostics
	state.Addresses, map_diags = types.MapValueFrom(ctx, state.getTypeFor("addresses"), elements)
	resp.Diagnostics.Append(map_diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)

}

func (o *AddressesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state AddressesResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing resource create", map[string]any{
		"resource_name": "panos_addresses",
		"function":      "Create",
	})

	var location address.Location

	if state.Location.FromPanoramaVsys != nil {
		location.FromPanoramaVsys = &address.FromPanoramaVsysLocation{

			Vsys: state.Location.FromPanoramaVsys.Vsys.ValueString(),
		}
	}
	if state.Location.DeviceGroup != nil {
		location.DeviceGroup = &address.DeviceGroupLocation{

			PanoramaDevice: state.Location.DeviceGroup.PanoramaDevice.ValueString(),
			DeviceGroup:    state.Location.DeviceGroup.Name.ValueString(),
		}
	}
	if !state.Location.Shared.IsNull() && state.Location.Shared.ValueBool() {
		location.Shared = true
	}
	if state.Location.Vsys != nil {
		location.Vsys = &address.VsysLocation{

			NgfwDevice: state.Location.Vsys.NgfwDevice.ValueString(),
			Vsys:       state.Location.Vsys.Name.ValueString(),
		}
	}
	if !state.Location.FromPanoramaShared.IsNull() && state.Location.FromPanoramaShared.ValueBool() {
		location.FromPanoramaShared = true
	}

	elements := make(map[string]AddressesResourceAddressesObject)
	resp.Diagnostics.Append(state.Addresses.ElementsAs(ctx, &elements, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	entries := make([]*address.Entry, 0, len(elements))
	for name, elt := range elements {
		var entry *address.Entry
		resp.Diagnostics.Append(elt.CopyToPango(ctx, &entry, nil)...)
		if resp.Diagnostics.HasError() {
			return
		}
		entry.Name = name
		entries = append(entries, entry)
	}

	readEntries, err := o.manager.ReadMany(ctx, location, entries)
	if err != nil {
		if errors.Is(err, sdkmanager.ErrObjectNotFound) {
			resp.State.RemoveResource(ctx)
		} else {
			resp.Diagnostics.AddError("Failed to read entries from the server", err.Error())
		}
		return
	}

	objects := make(map[string]AddressesResourceAddressesObject)
	for _, elt := range readEntries {
		var object AddressesResourceAddressesObject
		resp.Diagnostics.Append(object.CopyFromPango(ctx, elt, nil)...)
		if resp.Diagnostics.HasError() {
			return
		}
		objects[elt.Name] = object
	}

	var map_diags diag.Diagnostics
	state.Addresses, map_diags = types.MapValueFrom(ctx, state.getTypeFor("addresses"), objects)
	resp.Diagnostics.Append(map_diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)

}

func (r *AddressesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var state, plan AddressesResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing resource create", map[string]any{
		"resource_name": "panos_addresses_resource",
		"function":      "Create",
	})

	var location address.Location

	if !plan.Location.FromPanoramaShared.IsNull() && plan.Location.FromPanoramaShared.ValueBool() {
		location.FromPanoramaShared = true
	}
	if plan.Location.FromPanoramaVsys != nil {
		location.FromPanoramaVsys = &address.FromPanoramaVsysLocation{

			Vsys: plan.Location.FromPanoramaVsys.Vsys.ValueString(),
		}
	}
	if plan.Location.DeviceGroup != nil {
		location.DeviceGroup = &address.DeviceGroupLocation{

			PanoramaDevice: plan.Location.DeviceGroup.PanoramaDevice.ValueString(),
			DeviceGroup:    plan.Location.DeviceGroup.Name.ValueString(),
		}
	}
	if !plan.Location.Shared.IsNull() && plan.Location.Shared.ValueBool() {
		location.Shared = true
	}
	if plan.Location.Vsys != nil {
		location.Vsys = &address.VsysLocation{

			NgfwDevice: plan.Location.Vsys.NgfwDevice.ValueString(),
			Vsys:       plan.Location.Vsys.Name.ValueString(),
		}
	}

	// Basic logging.
	tflog.Info(ctx, "performing resource update", map[string]any{
		"resource_name": "panos_addresses_resource",
		"function":      "Update",
	})

	var elements map[string]AddressesResourceAddressesObject
	state.Addresses.ElementsAs(ctx, &elements, false)
	stateEntries := make([]*address.Entry, len(elements))
	idx := 0
	for name, elt := range elements {
		var entry *address.Entry
		resp.Diagnostics.Append(elt.CopyToPango(ctx, &entry, nil)...)
		if resp.Diagnostics.HasError() {
			return
		}
		entry.Name = name
		stateEntries[idx] = entry
		idx++
	}

	existing, err := r.manager.ReadMany(ctx, location, stateEntries)
	if err != nil && !sdkerrors.IsObjectNotFound(err) {
		resp.Diagnostics.AddError("Error while reading entries from the server", err.Error())
		return
	}

	existingEntriesByName := make(map[string]*address.Entry, len(existing))
	for _, elt := range existing {
		existingEntriesByName[elt.Name] = elt
	}

	plan.Addresses.ElementsAs(ctx, &elements, false)
	planEntries := make([]*address.Entry, len(elements))
	idx = 0
	for name, elt := range elements {
		entry, _ := existingEntriesByName[name]
		resp.Diagnostics.Append(elt.CopyToPango(ctx, &entry, nil)...)
		if resp.Diagnostics.HasError() {
			return
		}

		entry.Name = name
		planEntries[idx] = entry
		idx++
	}

	processed, err := r.manager.UpdateMany(ctx, location, stateEntries, planEntries)
	if err != nil {
		resp.Diagnostics.AddError("Error while updating entries", err.Error())
		return
	}

	objects := make(map[string]*AddressesResourceAddressesObject, len(processed))
	for _, elt := range processed {
		var object AddressesResourceAddressesObject
		copy_diags := object.CopyFromPango(ctx, elt, nil)
		resp.Diagnostics.Append(copy_diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		objects[elt.Name] = &object
	}

	var list_diags diag.Diagnostics
	plan.Addresses, list_diags = types.MapValueFrom(ctx, state.getTypeFor("addresses"), objects)
	resp.Diagnostics.Append(list_diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)

}

func (r *AddressesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state AddressesResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Basic logging.
	tflog.Info(ctx, "performing resource delete", map[string]any{
		"resource_name": "panos_addresses_resource",
		"function":      "Delete",
	})
	elements := make(map[string]AddressesResourceAddressesObject, len(state.Addresses.Elements()))
	state.Addresses.ElementsAs(ctx, &elements, false)

	var location address.Location

	if !state.Location.Shared.IsNull() && state.Location.Shared.ValueBool() {
		location.Shared = true
	}
	if state.Location.Vsys != nil {
		location.Vsys = &address.VsysLocation{

			Vsys:       state.Location.Vsys.Name.ValueString(),
			NgfwDevice: state.Location.Vsys.NgfwDevice.ValueString(),
		}
	}
	if !state.Location.FromPanoramaShared.IsNull() && state.Location.FromPanoramaShared.ValueBool() {
		location.FromPanoramaShared = true
	}
	if state.Location.FromPanoramaVsys != nil {
		location.FromPanoramaVsys = &address.FromPanoramaVsysLocation{

			Vsys: state.Location.FromPanoramaVsys.Vsys.ValueString(),
		}
	}
	if state.Location.DeviceGroup != nil {
		location.DeviceGroup = &address.DeviceGroupLocation{

			PanoramaDevice: state.Location.DeviceGroup.PanoramaDevice.ValueString(),
			DeviceGroup:    state.Location.DeviceGroup.Name.ValueString(),
		}
	}

	var names []string
	for name, _ := range elements {
		names = append(names, name)
	}
	err := r.manager.Delete(ctx, location, names)
	if err != nil {
		resp.Diagnostics.AddError("error while deleting entries", err.Error())
		return
	}

}

func (r *AddressesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("tfid"), req, resp)
}

type AddressesFromPanoramaVsysLocation struct {
	Vsys types.String `tfsdk:"vsys"`
}
type AddressesDeviceGroupLocation struct {
	PanoramaDevice types.String `tfsdk:"panorama_device"`
	Name           types.String `tfsdk:"name"`
}
type AddressesVsysLocation struct {
	NgfwDevice types.String `tfsdk:"ngfw_device"`
	Name       types.String `tfsdk:"name"`
}
type AddressesLocation struct {
	FromPanoramaShared types.Bool                         `tfsdk:"from_panorama_shared"`
	FromPanoramaVsys   *AddressesFromPanoramaVsysLocation `tfsdk:"from_panorama_vsys"`
	DeviceGroup        *AddressesDeviceGroupLocation      `tfsdk:"device_group"`
	Shared             types.Bool                         `tfsdk:"shared"`
	Vsys               *AddressesVsysLocation             `tfsdk:"vsys"`
}

func AddressesLocationSchema() rsschema.Attribute {
	return rsschema.SingleNestedAttribute{
		Description: "The location of this object.",
		Required:    true,
		Attributes: map[string]rsschema.Attribute{
			"shared": rsschema.BoolAttribute{
				Description: "Located in shared.",
				Optional:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},

				Validators: []validator.Bool{
					boolvalidator.ExactlyOneOf(path.Expressions{
						path.MatchRelative().AtParent().AtName("from_panorama_shared"),
						path.MatchRelative().AtParent().AtName("from_panorama_vsys"),
						path.MatchRelative().AtParent().AtName("device_group"),
						path.MatchRelative().AtParent().AtName("shared"),
						path.MatchRelative().AtParent().AtName("vsys"),
					}...),
				},
			},
			"vsys": rsschema.SingleNestedAttribute{
				Description: "Located in a specific vsys.",
				Optional:    true,
				Attributes: map[string]rsschema.Attribute{
					"ngfw_device": rsschema.StringAttribute{
						Description: "The NGFW device.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString("localhost.localdomain"),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"name": rsschema.StringAttribute{
						Description: "The vsys.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString("vsys1"),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
				},
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplace(),
				},
			},
			"from_panorama_shared": rsschema.BoolAttribute{
				Description: "Located in shared in the config pushed from Panorama.",
				Optional:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
			"from_panorama_vsys": rsschema.SingleNestedAttribute{
				Description: "Located in a specific vsys in the config pushed from Panorama.",
				Optional:    true,
				Attributes: map[string]rsschema.Attribute{
					"vsys": rsschema.StringAttribute{
						Description: "The vsys.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString("vsys1"),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
				},
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplace(),
				},
			},
			"device_group": rsschema.SingleNestedAttribute{
				Description: "Located in a specific device group.",
				Optional:    true,
				Attributes: map[string]rsschema.Attribute{
					"panorama_device": rsschema.StringAttribute{
						Description: "The panorama device.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString("localhost.localdomain"),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"name": rsschema.StringAttribute{
						Description: "The device group.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(""),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
				},
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}
