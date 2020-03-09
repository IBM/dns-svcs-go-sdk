/**
 * (C) Copyright IBM Corp. 2020.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package dnssvcsinstancesv2 : Operations and models for the DnsSvcsInstancesV2 service
package dnssvcsinstancesv2

import (
	"fmt"

	common "github.com/IBM/dns-svcs-go-sdk/common"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
)

// DnsSvcsInstancesV2 : Manage lifecycle of PDNS resource instance using Resource Controller APIs.
//
// Version: 2.0
type DnsSvcsInstancesV2 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://resource-controller.cloud.ibm.com/v2"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "dns_svcs_instances"

// DnsSvcsInstancesV2Options : Service options
type DnsSvcsInstancesV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewDnsSvcsInstancesV2UsingExternalConfig : constructs an instance of DnsSvcsInstancesV2 with passed in options and external configuration.
func NewDnsSvcsInstancesV2UsingExternalConfig(options *DnsSvcsInstancesV2Options) (dnsSvcsInstances *DnsSvcsInstancesV2, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	dnsSvcsInstances, err = NewDnsSvcsInstancesV2(options)
	if err != nil {
		return
	}

	err = dnsSvcsInstances.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = dnsSvcsInstances.Service.SetServiceURL(options.URL)
	}
	return
}

// NewDnsSvcsInstancesV2 : constructs an instance of DnsSvcsInstancesV2 with passed in options.
func NewDnsSvcsInstancesV2(options *DnsSvcsInstancesV2Options) (service *DnsSvcsInstancesV2, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &DnsSvcsInstancesV2{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (dnsSvcsInstances *DnsSvcsInstancesV2) SetServiceURL(url string) error {
	return dnsSvcsInstances.Service.SetServiceURL(url)
}

// ListResourceInstances : Get a list of all resource instances
// Get a list of all resource instances.
func (dnsSvcsInstances *DnsSvcsInstancesV2) ListResourceInstances(listResourceInstancesOptions *ListResourceInstancesOptions) (result *ResourceInstancesList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listResourceInstancesOptions, "listResourceInstancesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listResourceInstancesOptions, "listResourceInstancesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"resource_instances"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(dnsSvcsInstances.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listResourceInstancesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs_instances", "V2", "ListResourceInstances")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Accept-Language", "en-US")

	builder.AddQuery("resource_id", fmt.Sprint(*listResourceInstancesOptions.ResourceID))
	builder.AddQuery("type", fmt.Sprint(*listResourceInstancesOptions.Type))
	if listResourceInstancesOptions.Guid != nil {
		builder.AddQuery("guid", fmt.Sprint(*listResourceInstancesOptions.Guid))
	}
	if listResourceInstancesOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listResourceInstancesOptions.Name))
	}
	if listResourceInstancesOptions.ResourceGroupID != nil {
		builder.AddQuery("resource_group_id", fmt.Sprint(*listResourceInstancesOptions.ResourceGroupID))
	}
	if listResourceInstancesOptions.ResourcePlanID != nil {
		builder.AddQuery("resource_plan_id", fmt.Sprint(*listResourceInstancesOptions.ResourcePlanID))
	}
	if listResourceInstancesOptions.SubType != nil {
		builder.AddQuery("sub_type", fmt.Sprint(*listResourceInstancesOptions.SubType))
	}
	if listResourceInstancesOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listResourceInstancesOptions.Limit))
	}
	if listResourceInstancesOptions.UpdatedFrom != nil {
		builder.AddQuery("updated_from", fmt.Sprint(*listResourceInstancesOptions.UpdatedFrom))
	}
	if listResourceInstancesOptions.UpdatedTo != nil {
		builder.AddQuery("updated_to", fmt.Sprint(*listResourceInstancesOptions.UpdatedTo))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcsInstances.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResourceInstancesList(m)
		response.Result = result
	}

	return
}

// CreateResourceInstance : Create (provision) a new resource instance
// Provision a new resource in the specified location for the selected plan.
func (dnsSvcsInstances *DnsSvcsInstancesV2) CreateResourceInstance(createResourceInstanceOptions *CreateResourceInstanceOptions) (result *ResourceInstance, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createResourceInstanceOptions, "createResourceInstanceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createResourceInstanceOptions, "createResourceInstanceOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"resource_instances"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(dnsSvcsInstances.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createResourceInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs_instances", "V2", "CreateResourceInstance")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createResourceInstanceOptions.Name != nil {
		body["name"] = createResourceInstanceOptions.Name
	}
	if createResourceInstanceOptions.Target != nil {
		body["target"] = createResourceInstanceOptions.Target
	}
	if createResourceInstanceOptions.ResourceGroup != nil {
		body["resource_group"] = createResourceInstanceOptions.ResourceGroup
	}
	if createResourceInstanceOptions.ResourcePlanID != nil {
		body["resource_plan_id"] = createResourceInstanceOptions.ResourcePlanID
	}
	if createResourceInstanceOptions.Tags != nil {
		body["tags"] = createResourceInstanceOptions.Tags
	}
	if createResourceInstanceOptions.AllowCleanup != nil {
		body["allow_cleanup"] = createResourceInstanceOptions.AllowCleanup
	}
	if createResourceInstanceOptions.Parameters != nil {
		body["parameters"] = createResourceInstanceOptions.Parameters
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcsInstances.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResourceInstance(m)
		response.Result = result
	}

	return
}

// GetResourceInstance : Get a resource instance
// Retrieve a resource instance by ID.
func (dnsSvcsInstances *DnsSvcsInstancesV2) GetResourceInstance(getResourceInstanceOptions *GetResourceInstanceOptions) (result *ResourceInstance, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getResourceInstanceOptions, "getResourceInstanceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getResourceInstanceOptions, "getResourceInstanceOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"resource_instances"}
	pathParameters := []string{*getResourceInstanceOptions.ID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(dnsSvcsInstances.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getResourceInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs_instances", "V2", "GetResourceInstance")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcsInstances.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResourceInstance(m)
		response.Result = result
	}

	return
}

// DeleteResourceInstance : Delete a resource instance
// Delete a resource instance by ID.
func (dnsSvcsInstances *DnsSvcsInstancesV2) DeleteResourceInstance(deleteResourceInstanceOptions *DeleteResourceInstanceOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteResourceInstanceOptions, "deleteResourceInstanceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteResourceInstanceOptions, "deleteResourceInstanceOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"resource_instances"}
	pathParameters := []string{*deleteResourceInstanceOptions.ID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(dnsSvcsInstances.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteResourceInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs_instances", "V2", "DeleteResourceInstance")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcsInstances.Service.Request(request, nil)

	return
}

// UpdateResourceInstance : Update a resource instance
// Update a resource instance by ID.
func (dnsSvcsInstances *DnsSvcsInstancesV2) UpdateResourceInstance(updateResourceInstanceOptions *UpdateResourceInstanceOptions) (result *ResourceInstance, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateResourceInstanceOptions, "updateResourceInstanceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateResourceInstanceOptions, "updateResourceInstanceOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"resource_instances"}
	pathParameters := []string{*updateResourceInstanceOptions.ID}

	builder := core.NewRequestBuilder(core.PATCH)
	_, err = builder.ConstructHTTPURL(dnsSvcsInstances.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateResourceInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs_instances", "V2", "UpdateResourceInstance")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddHeader("Accept-Language", "en-US")

	body := make(map[string]interface{})
	if updateResourceInstanceOptions.Name != nil {
		body["name"] = updateResourceInstanceOptions.Name
	}
	if updateResourceInstanceOptions.Parameters != nil {
		body["parameters"] = updateResourceInstanceOptions.Parameters
	}
	if updateResourceInstanceOptions.ResourcePlanID != nil {
		body["resource_plan_id"] = updateResourceInstanceOptions.ResourcePlanID
	}
	if updateResourceInstanceOptions.AllowCleanup != nil {
		body["allow_cleanup"] = updateResourceInstanceOptions.AllowCleanup
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcsInstances.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResourceInstance(m)
		response.Result = result
	}

	return
}

// CreateResourceInstanceOptions : The CreateResourceInstance options.
type CreateResourceInstanceOptions struct {
	// The name of the instance. Must be 180 characters or less and cannot include any special characters other than
	// `(space) - . _ :`.
	Name *string `json:"name" validate:"required"`

	// The deployment location where the instance should be hosted.
	Target *string `json:"target" validate:"required"`

	// Short or long ID of resource group.
	ResourceGroup *string `json:"resource_group" validate:"required"`

	// The unique ID of the plan associated with the offering. This value is provided by and stored in the global catalog.
	ResourcePlanID *string `json:"resource_plan_id" validate:"required"`

	// Tags that are attached to the instance after provisioning. These tags can be searched and managed through the
	// Tagging API in IBM Cloud.
	Tags []string `json:"tags,omitempty"`

	// A boolean that dictates if the resource instance should be deleted (cleaned up) during the processing of a region
	// instance delete call.
	AllowCleanup *bool `json:"allow_cleanup,omitempty"`

	// Configuration options represented as key-value pairs that are passed through to the target resource brokers.
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateResourceInstanceOptions : Instantiate CreateResourceInstanceOptions
func (*DnsSvcsInstancesV2) NewCreateResourceInstanceOptions(name string, target string, resourceGroup string, resourcePlanID string) *CreateResourceInstanceOptions {
	return &CreateResourceInstanceOptions{
		Name:           core.StringPtr(name),
		Target:         core.StringPtr(target),
		ResourceGroup:  core.StringPtr(resourceGroup),
		ResourcePlanID: core.StringPtr(resourcePlanID),
	}
}

// SetName : Allow user to set Name
func (options *CreateResourceInstanceOptions) SetName(name string) *CreateResourceInstanceOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetTarget : Allow user to set Target
func (options *CreateResourceInstanceOptions) SetTarget(target string) *CreateResourceInstanceOptions {
	options.Target = core.StringPtr(target)
	return options
}

// SetResourceGroup : Allow user to set ResourceGroup
func (options *CreateResourceInstanceOptions) SetResourceGroup(resourceGroup string) *CreateResourceInstanceOptions {
	options.ResourceGroup = core.StringPtr(resourceGroup)
	return options
}

// SetResourcePlanID : Allow user to set ResourcePlanID
func (options *CreateResourceInstanceOptions) SetResourcePlanID(resourcePlanID string) *CreateResourceInstanceOptions {
	options.ResourcePlanID = core.StringPtr(resourcePlanID)
	return options
}

// SetTags : Allow user to set Tags
func (options *CreateResourceInstanceOptions) SetTags(tags []string) *CreateResourceInstanceOptions {
	options.Tags = tags
	return options
}

// SetAllowCleanup : Allow user to set AllowCleanup
func (options *CreateResourceInstanceOptions) SetAllowCleanup(allowCleanup bool) *CreateResourceInstanceOptions {
	options.AllowCleanup = core.BoolPtr(allowCleanup)
	return options
}

// SetParameters : Allow user to set Parameters
func (options *CreateResourceInstanceOptions) SetParameters(parameters map[string]interface{}) *CreateResourceInstanceOptions {
	options.Parameters = parameters
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateResourceInstanceOptions) SetHeaders(param map[string]string) *CreateResourceInstanceOptions {
	options.Headers = param
	return options
}

// DeleteResourceInstanceOptions : The DeleteResourceInstance options.
type DeleteResourceInstanceOptions struct {
	// The short or long ID of the instance.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteResourceInstanceOptions : Instantiate DeleteResourceInstanceOptions
func (*DnsSvcsInstancesV2) NewDeleteResourceInstanceOptions(id string) *DeleteResourceInstanceOptions {
	return &DeleteResourceInstanceOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *DeleteResourceInstanceOptions) SetID(id string) *DeleteResourceInstanceOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteResourceInstanceOptions) SetHeaders(param map[string]string) *DeleteResourceInstanceOptions {
	options.Headers = param
	return options
}

// GetResourceInstanceOptions : The GetResourceInstance options.
type GetResourceInstanceOptions struct {
	// The short or long ID of the instance.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetResourceInstanceOptions : Instantiate GetResourceInstanceOptions
func (*DnsSvcsInstancesV2) NewGetResourceInstanceOptions(id string) *GetResourceInstanceOptions {
	return &GetResourceInstanceOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetResourceInstanceOptions) SetID(id string) *GetResourceInstanceOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetResourceInstanceOptions) SetHeaders(param map[string]string) *GetResourceInstanceOptions {
	options.Headers = param
	return options
}

// ListResourceInstancesOptions : The ListResourceInstances options.
type ListResourceInstancesOptions struct {
	// The unique ID of the offering. This value is provided by and stored in the global catalog.
	ResourceID *string `json:"resource_id" validate:"required"`

	// The type of the instance. For example, `service_instance`.
	Type *string `json:"type" validate:"required"`

	// When you provision a new resource in the specified location for the selected plan, a GUID (globally unique
	// identifier) is created. This is a unique internal GUID managed by Resource controller that corresponds to the
	// instance.
	Guid *string `json:"guid,omitempty"`

	// The human-readable name of the instance.
	Name *string `json:"name,omitempty"`

	// Short ID of a resource group.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// The unique ID of the plan associated with the offering. This value is provided by and stored in the global catalog.
	ResourcePlanID *string `json:"resource_plan_id,omitempty"`

	// The sub-type of instance, e.g. `cfaas`.
	SubType *string `json:"sub_type,omitempty"`

	// Limit on how many items should be returned.
	Limit *string `json:"limit,omitempty"`

	// Start date inclusive filter.
	UpdatedFrom *string `json:"updated_from,omitempty"`

	// End date inclusive filter.
	UpdatedTo *string `json:"updated_to,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListResourceInstancesOptions.Type property.
// The type of the instance. For example, `service_instance`.
const (
	ListResourceInstancesOptions_Type_ServiceInstance = "service_instance"
)

// NewListResourceInstancesOptions : Instantiate ListResourceInstancesOptions
func (*DnsSvcsInstancesV2) NewListResourceInstancesOptions(resourceID string, typeVar string) *ListResourceInstancesOptions {
	return &ListResourceInstancesOptions{
		ResourceID: core.StringPtr(resourceID),
		Type:       core.StringPtr(typeVar),
	}
}

// SetResourceID : Allow user to set ResourceID
func (options *ListResourceInstancesOptions) SetResourceID(resourceID string) *ListResourceInstancesOptions {
	options.ResourceID = core.StringPtr(resourceID)
	return options
}

// SetType : Allow user to set Type
func (options *ListResourceInstancesOptions) SetType(typeVar string) *ListResourceInstancesOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetGuid : Allow user to set Guid
func (options *ListResourceInstancesOptions) SetGuid(guid string) *ListResourceInstancesOptions {
	options.Guid = core.StringPtr(guid)
	return options
}

// SetName : Allow user to set Name
func (options *ListResourceInstancesOptions) SetName(name string) *ListResourceInstancesOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetResourceGroupID : Allow user to set ResourceGroupID
func (options *ListResourceInstancesOptions) SetResourceGroupID(resourceGroupID string) *ListResourceInstancesOptions {
	options.ResourceGroupID = core.StringPtr(resourceGroupID)
	return options
}

// SetResourcePlanID : Allow user to set ResourcePlanID
func (options *ListResourceInstancesOptions) SetResourcePlanID(resourcePlanID string) *ListResourceInstancesOptions {
	options.ResourcePlanID = core.StringPtr(resourcePlanID)
	return options
}

// SetSubType : Allow user to set SubType
func (options *ListResourceInstancesOptions) SetSubType(subType string) *ListResourceInstancesOptions {
	options.SubType = core.StringPtr(subType)
	return options
}

// SetLimit : Allow user to set Limit
func (options *ListResourceInstancesOptions) SetLimit(limit string) *ListResourceInstancesOptions {
	options.Limit = core.StringPtr(limit)
	return options
}

// SetUpdatedFrom : Allow user to set UpdatedFrom
func (options *ListResourceInstancesOptions) SetUpdatedFrom(updatedFrom string) *ListResourceInstancesOptions {
	options.UpdatedFrom = core.StringPtr(updatedFrom)
	return options
}

// SetUpdatedTo : Allow user to set UpdatedTo
func (options *ListResourceInstancesOptions) SetUpdatedTo(updatedTo string) *ListResourceInstancesOptions {
	options.UpdatedTo = core.StringPtr(updatedTo)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListResourceInstancesOptions) SetHeaders(param map[string]string) *ListResourceInstancesOptions {
	options.Headers = param
	return options
}

// PlanHistoryItem : An element of the plan history of the instance.
type PlanHistoryItem struct {
	// The unique ID of the plan associated with the offering. This value is provided by and stored in the global catalog.
	ResourcePlanID *string `json:"resource_plan_id" validate:"required"`

	// The date on which the plan was changed.
	StartDate *strfmt.DateTime `json:"start_date" validate:"required"`
}

// UnmarshalPlanHistoryItem constructs an instance of PlanHistoryItem from the specified map.
func UnmarshalPlanHistoryItem(m map[string]interface{}) (result *PlanHistoryItem, err error) {
	obj := new(PlanHistoryItem)
	obj.ResourcePlanID, err = core.UnmarshalString(m, "resource_plan_id")
	if err != nil {
		return
	}
	obj.StartDate, err = core.UnmarshalDateTime(m, "start_date")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalPlanHistoryItemSlice unmarshals a slice of PlanHistoryItem instances from the specified list of maps.
func UnmarshalPlanHistoryItemSlice(s []interface{}) (slice []PlanHistoryItem, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'PlanHistoryItem'")
			return
		}
		obj, e := UnmarshalPlanHistoryItem(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalPlanHistoryItemSliceAsProperty unmarshals a slice of PlanHistoryItem instances that are stored as a property
// within the specified map.
func UnmarshalPlanHistoryItemSliceAsProperty(m map[string]interface{}, propertyName string) (slice []PlanHistoryItem, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'PlanHistoryItem'", propertyName)
			return
		}
		slice, err = UnmarshalPlanHistoryItemSlice(vSlice)
	}
	return
}

// ResourceInstance : A resource instance.
type ResourceInstance struct {
	// The ID associated with the instance.
	ID *string `json:"id,omitempty"`

	// When you create a new resource, a globally unique identifier (GUID) is assigned. This GUID is a unique internal
	// identifier managed by the resource controller that corresponds to the instance.
	Guid *string `json:"guid,omitempty"`

	// The full Cloud Resource Name (CRN) associated with the instance. For more information about this format, see [Cloud
	// Resource Names](https://cloud.ibm.com/docs/overview?topic=overview-crn).
	Crn *string `json:"crn,omitempty"`

	// When you provision a new resource, a relative URL path is created identifying the location of the instance.
	URL *string `json:"url,omitempty"`

	// The human-readable name of the instance.
	Name *string `json:"name,omitempty"`

	// The location of cloud geography/region/zone/data center that the resource resides in
	RegionID *string `json:"region_id,omitempty"`

	// An alpha-numeric value identifying the account ID.
	AccountID *string `json:"account_id,omitempty"`

	// The short ID of the resource group.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// The long ID (full CRN) of the resource group.
	ResourceGroupCrn *string `json:"resource_group_crn,omitempty"`

	// The unique ID of the offering. This value is provided by and stored in the global catalog.
	ResourceID *string `json:"resource_id,omitempty"`

	// The unique ID of the plan associated with the offering. This value is provided by and stored in the global catalog.
	ResourcePlanID *string `json:"resource_plan_id,omitempty"`

	// The full deployment CRN as defined in the global catalog. The Cloud Resource Name (CRN) of the deployment location
	// where the instance is provisioned.
	TargetCrn *string `json:"target_crn,omitempty"`

	// The current state of the instance. For example, if the instance is deleted, it will return removed.
	State *string `json:"state,omitempty"`

	// The type of the instance, e.g. `service_instance`.
	Type *string `json:"type,omitempty"`

	// The sub-type of instance, e.g. `cfaas`.
	SubType *string `json:"sub_type,omitempty"`

	// A boolean that dictates if the resource instance should be deleted (cleaned up) during the processing of a region
	// instance delete call.
	AllowCleanup *bool `json:"allow_cleanup,omitempty"`

	// The status of the last operation requested on the instance.
	LastOperation map[string]interface{} `json:"last_operation,omitempty"`

	// The resource-broker-provided URL to access administrative features of the instance.
	DashboardURL *string `json:"dashboard_url,omitempty"`

	// The plan history of the instance.
	PlanHistory []PlanHistoryItem `json:"plan_history,omitempty"`

	// The relative path to the resource aliases for the instance.
	ResourceAliasesURL *string `json:"resource_aliases_url,omitempty"`

	// The relative path to the resource bindings for the instance.
	ResourceBindingsURL *string `json:"resource_bindings_url,omitempty"`

	// The relative path to the resource keys for the instance.
	ResourceKeysURL *string `json:"resource_keys_url,omitempty"`

	// The date when the instance was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// The date when the instance was last updated.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	// The date when the instance was deleted.
	DeletedAt *strfmt.DateTime `json:"deleted_at,omitempty"`

	// A boolean that dictates if the resource instance is going to be migrated
	Migrated *bool `json:"migrated,omitempty"`
}

// UnmarshalResourceInstance constructs an instance of ResourceInstance from the specified map.
func UnmarshalResourceInstance(m map[string]interface{}) (result *ResourceInstance, err error) {
	obj := new(ResourceInstance)
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.Guid, err = core.UnmarshalString(m, "guid")
	if err != nil {
		return
	}
	obj.Crn, err = core.UnmarshalString(m, "crn")
	if err != nil {
		return
	}
	obj.URL, err = core.UnmarshalString(m, "url")
	if err != nil {
		return
	}
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	obj.RegionID, err = core.UnmarshalString(m, "region_id")
	if err != nil {
		return
	}
	obj.AccountID, err = core.UnmarshalString(m, "account_id")
	if err != nil {
		return
	}
	obj.ResourceGroupID, err = core.UnmarshalString(m, "resource_group_id")
	if err != nil {
		return
	}
	obj.ResourceGroupCrn, err = core.UnmarshalString(m, "resource_group_crn")
	if err != nil {
		return
	}
	obj.ResourceID, err = core.UnmarshalString(m, "resource_id")
	if err != nil {
		return
	}
	obj.ResourcePlanID, err = core.UnmarshalString(m, "resource_plan_id")
	if err != nil {
		return
	}
	obj.TargetCrn, err = core.UnmarshalString(m, "target_crn")
	if err != nil {
		return
	}
	obj.State, err = core.UnmarshalString(m, "state")
	if err != nil {
		return
	}
	obj.Type, err = core.UnmarshalString(m, "type")
	if err != nil {
		return
	}
	obj.SubType, err = core.UnmarshalString(m, "sub_type")
	if err != nil {
		return
	}
	obj.AllowCleanup, err = core.UnmarshalBool(m, "allow_cleanup")
	if err != nil {
		return
	}
	obj.DashboardURL, err = core.UnmarshalString(m, "dashboard_url")
	if err != nil {
		return
	}
	obj.PlanHistory, err = UnmarshalPlanHistoryItemSliceAsProperty(m, "plan_history")
	if err != nil {
		return
	}
	obj.ResourceAliasesURL, err = core.UnmarshalString(m, "resource_aliases_url")
	if err != nil {
		return
	}
	obj.ResourceBindingsURL, err = core.UnmarshalString(m, "resource_bindings_url")
	if err != nil {
		return
	}
	obj.ResourceKeysURL, err = core.UnmarshalString(m, "resource_keys_url")
	if err != nil {
		return
	}
	obj.CreatedAt, err = core.UnmarshalDateTime(m, "created_at")
	if err != nil {
		return
	}
	obj.UpdatedAt, err = core.UnmarshalDateTime(m, "updated_at")
	if err != nil {
		return
	}
	obj.Migrated, err = core.UnmarshalBool(m, "migrated")
	if err != nil {
		return
	}
	if m["deleted_at"] != nil {
		obj.DeletedAt, err = core.UnmarshalDateTime(m, "deleted_at")
		if err != nil {
			return
		}
	}

	result = obj
	return
}

// UnmarshalResourceInstanceSlice unmarshals a slice of ResourceInstance instances from the specified list of maps.
func UnmarshalResourceInstanceSlice(s []interface{}) (slice []ResourceInstance, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ResourceInstance'")
			return
		}
		obj, e := UnmarshalResourceInstance(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResourceInstanceSliceAsProperty unmarshals a slice of ResourceInstance instances that are stored as a property
// within the specified map.
func UnmarshalResourceInstanceSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ResourceInstance, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ResourceInstance'", propertyName)
			return
		}
		slice, err = UnmarshalResourceInstanceSlice(vSlice)
	}
	return
}

// ResourceInstancesList : A list of resource instances.
type ResourceInstancesList struct {
	// The URL for requesting the next page of results.
	NextURL *string `json:"next_url,omitempty"`

	// A list of resource instances.
	Resources []ResourceInstance `json:"resources" validate:"required"`

	// The number of resource instances in `resources`.
	RowsCount *int64 `json:"rows_count" validate:"required"`
}

// UnmarshalResourceInstancesList constructs an instance of ResourceInstancesList from the specified map.
func UnmarshalResourceInstancesList(m map[string]interface{}) (result *ResourceInstancesList, err error) {
	obj := new(ResourceInstancesList)
	if m["next_url"] != nil {
		obj.NextURL, err = core.UnmarshalString(m, "next_url")
		if err != nil {
			return
		}
	}

	obj.Resources, err = UnmarshalResourceInstanceSliceAsProperty(m, "resources")
	if err != nil {
		return
	}
	obj.RowsCount, err = core.UnmarshalInt64(m, "rows_count")
	if err != nil {
		return
	}
	result = obj
	return
}

// UpdateResourceInstanceOptions : The UpdateResourceInstance options.
type UpdateResourceInstanceOptions struct {
	// The short or long ID of the instance.
	ID *string `json:"id" validate:"required"`

	// The new name of the instance. Must be 180 characters or less and cannot include any special characters other than
	// `(space) - . _ :`.
	Name *string `json:"name,omitempty"`

	// The new configuration options for the instance.
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// The unique ID of the plan associated with the offering. This value is provided by and stored in the global catalog.
	ResourcePlanID *string `json:"resource_plan_id,omitempty"`

	// A boolean that dictates if the resource instance should be deleted (cleaned up) during the processing of a region
	// instance delete call.
	AllowCleanup *bool `json:"allow_cleanup,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateResourceInstanceOptions : Instantiate UpdateResourceInstanceOptions
func (*DnsSvcsInstancesV2) NewUpdateResourceInstanceOptions(id string) *UpdateResourceInstanceOptions {
	return &UpdateResourceInstanceOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *UpdateResourceInstanceOptions) SetID(id string) *UpdateResourceInstanceOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateResourceInstanceOptions) SetName(name string) *UpdateResourceInstanceOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetParameters : Allow user to set Parameters
func (options *UpdateResourceInstanceOptions) SetParameters(parameters map[string]interface{}) *UpdateResourceInstanceOptions {
	options.Parameters = parameters
	return options
}

// SetResourcePlanID : Allow user to set ResourcePlanID
func (options *UpdateResourceInstanceOptions) SetResourcePlanID(resourcePlanID string) *UpdateResourceInstanceOptions {
	options.ResourcePlanID = core.StringPtr(resourcePlanID)
	return options
}

// SetAllowCleanup : Allow user to set AllowCleanup
func (options *UpdateResourceInstanceOptions) SetAllowCleanup(allowCleanup bool) *UpdateResourceInstanceOptions {
	options.AllowCleanup = core.BoolPtr(allowCleanup)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateResourceInstanceOptions) SetHeaders(param map[string]string) *UpdateResourceInstanceOptions {
	options.Headers = param
	return options
}
