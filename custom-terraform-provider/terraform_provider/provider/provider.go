package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-mockupstream/mock_upstream"
)

type (
	MockUpstreamProvider struct {
		version string
	}

	MockUpstreamProviderModel struct {
		BaseUrl types.String `tfsdk:"base_url"`
	}
)

var (
	_ provider.Provider = &MockUpstreamProvider{}
)

// Metadata https://developer.hashicorp.com/terraform/plugin/framework/providers#metadata-method
func (m *MockUpstreamProvider) Metadata(ctx context.Context, request provider.MetadataRequest, response *provider.MetadataResponse) {
	response.TypeName = "mockupstream"
	response.Version = m.version
}

// Schema https://developer.hashicorp.com/terraform/plugin/framework/providers#schema-method
func (m *MockUpstreamProvider) Schema(ctx context.Context, request provider.SchemaRequest, response *provider.SchemaResponse) {
	fmt.Println("Schema...")
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Base URL of the mock upstream",
				Required:    true,
			},
		},
	}
}

// Configure https://developer.hashicorp.com/terraform/plugin/framework/providers#configure-method
func (m *MockUpstreamProvider) Configure(ctx context.Context, request provider.ConfigureRequest, response *provider.ConfigureResponse) {
	var mockUpstreamProviderModel MockUpstreamProviderModel

	response.Diagnostics.Append(request.Config.Get(ctx, &mockUpstreamProviderModel)...)
	if response.Diagnostics.HasError() {
		return
	}

	mockUpstreamClient := mock_upstream.BuildClient(mockUpstreamProviderModel.BaseUrl.ValueString())

	response.ResourceData = mockUpstreamClient
	response.DataSourceData = mockUpstreamClient
}

// DataSources https://developer.hashicorp.com/terraform/plugin/framework/providers#datasources
func (m *MockUpstreamProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return nil
}

// Resources https://developer.hashicorp.com/terraform/plugin/framework/providers#resources
func (m *MockUpstreamProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewUserResource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &MockUpstreamProvider{
			version: version,
		}
	}
}
