package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-mockupstream/mock_upstream"
)

type (
	UserResource struct {
		client *mock_upstream.MockUpstreamClient
	}

	UserResourceModel struct {
		Name types.String `tfsdk:"name"`
		Age  types.Int32  `tfsdk:"age"`
		ID   types.Int64  `tfsdk:"id"`
	}
)

var (
	_ resource.Resource              = &UserResource{}
	_ resource.ResourceWithConfigure = &UserResource{}
)

// Configure https://developer.hashicorp.com/terraform/plugin/framework/resources/configure
func (u *UserResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	// Always perform a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if request.ProviderData == nil {
		return
	}

	client, ok := request.ProviderData.(*mock_upstream.MockUpstreamClient)
	if !ok {
		response.Diagnostics.AddError(
			"Unexpected type for client",
			fmt.Sprintf("Expected *mock_upstream.MockUpstreamClient; received %T", request.ProviderData),
		)

		return
	}

	u.client = client
}

// Metadata https://developer.hashicorp.com/terraform/plugin/framework/resources#metadata-method
func (u *UserResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_user"
}

// Schema https://developer.hashicorp.com/terraform/plugin/framework/resources#schema-method
func (u *UserResource) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Description: "Name of the user",
				Required:    true,
			},
			"age": schema.Int32Attribute{
				Description: "age of the user",
				Required:    true,
			},
			"id": schema.Int64Attribute{
				Description: "Id of the created user",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					// Computed means this id will be unknown during the update. This modifier is used
					// to copy the id from the previous state whenever there is an update
					int64planmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

// Create https://developer.hashicorp.com/terraform/plugin/framework/resources/create
func (u *UserResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var plan UserResourceModel

	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	user := mock_upstream.User{
		Name: plan.Name.ValueString(),
		Age:  plan.Age.ValueInt32(),
	}
	err := u.client.CreateUser(ctx, &user)
	if err != nil {
		response.Diagnostics.AddError("failed to create user", err.Error())
	}
	plan.ID = types.Int64Value(user.Id)

	response.Diagnostics.Append(response.State.Set(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}
}

// Read https://developer.hashicorp.com/terraform/plugin/framework/resources/read
func (u *UserResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var state UserResourceModel

	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	user := mock_upstream.User{
		Name: state.Name.ValueString(),
		Age:  state.Age.ValueInt32(),
		Id:   state.ID.ValueInt64(),
	}
	err := u.client.GetUser(ctx, &user)
	if err != nil {
		response.Diagnostics.AddError("failed to get user", err.Error())
	}

	state.Name = types.StringValue(user.Name)
	state.Age = types.Int32Value(user.Age)

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

// Update https://developer.hashicorp.com/terraform/plugin/framework/resources/update
func (u *UserResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var plan UserResourceModel

	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	user := mock_upstream.User{
		Name: plan.Name.ValueString(),
		Age:  plan.Age.ValueInt32(),
		Id:   plan.ID.ValueInt64(),
	}

	err := u.client.UpdateUser(ctx, &user)
	if err != nil {
		response.Diagnostics.AddError("failed to update user", err.Error())
	}

	response.Diagnostics.Append(response.State.Set(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}
}

// Delete https://developer.hashicorp.com/terraform/plugin/framework/resources/delete
func (u *UserResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var state UserResourceModel

	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	user := mock_upstream.User{
		Name: state.Name.ValueString(),
		Age:  state.Age.ValueInt32(),
		Id:   state.ID.ValueInt64(),
	}
	err := u.client.DeleteUser(ctx, &user)
	if err != nil {
		response.Diagnostics.AddError("failed to delete user", err.Error())
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func NewUserResource() resource.Resource {
	return &UserResource{}
}
