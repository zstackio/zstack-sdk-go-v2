// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteResourceStack deletes ResourceStack
func (cli *ZSClient) DeleteResourceStack(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/cloudformation/stack", uuid, string(deleteMode))
}
// QueryResourceStack queries ResourceStack list
func (cli *ZSClient) QueryResourceStack(params *param.QueryParam) ([]view.ResourceStackInventoryView, error) {
	var resp []view.ResourceStackInventoryView
	return resp, cli.List("v1/cloudformation/stack", params, &resp)
}

func (cli *ZSClient) GetResourceStack(uuid string) (*view.ResourceStackInventoryView, error) {
	var resp view.ResourceStackInventoryView
	if err := cli.Get("v1/cloudformation/stack", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageResourceStack Pagination
func (cli *ZSClient) PageResourceStack(params *param.QueryParam) ([]view.ResourceStackInventoryView, int, error) {
	var resourceStacks []view.ResourceStackInventoryView
	total, err := cli.Page("v1/cloudformation/stack", params, &resourceStacks)
	return resourceStacks, total, err
}
// CreateResourceStack creates ResourceStack
func (cli *ZSClient) CreateResourceStack(params param.CreateResourceStackParam) (*view.ResourceStackInventoryView, error) {
	resp := view.ResourceStackInventoryView{}
	if err := cli.Post("v1/cloudformation/stack", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateResourceStack updates ResourceStack
func (cli *ZSClient) UpdateResourceStack(uuid string, params param.UpdateResourceStackParam) (*view.ResourceStackInventoryView, error) {
	resp := view.ResourceStackInventoryView{}
	if err := cli.PutWithRespKey("v1/cloudformation/stack", uuid, "", map[string]interface{}{
		"updateResourceStack": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceStackVmStatus gets ResourceStackVmStatus by uuid
func (cli *ZSClient) GetResourceStackVmStatus() (*view.GetResourceStackVmStatusView, error) {
	var resp view.GetResourceStackVmStatusView
	if err := cli.GetWithRespKey("v1/cloudformation/stack/monitor/vmstatus", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEventFromResourceStack queries EventFromResourceStack list
func (cli *ZSClient) QueryEventFromResourceStack(params *param.QueryParam) ([]view.CloudFormationStackEventInventoryView, error) {
	var resp []view.CloudFormationStackEventInventoryView
	return resp, cli.List("v1/cloudformation/event", params, &resp)
}

func (cli *ZSClient) GetEventFromResourceStack(uuid string) (*view.CloudFormationStackEventInventoryView, error) {
	var resp view.CloudFormationStackEventInventoryView
	if err := cli.Get("v1/cloudformation/event", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEventFromResourceStack Pagination
func (cli *ZSClient) PageEventFromResourceStack(params *param.QueryParam) ([]view.CloudFormationStackEventInventoryView, int, error) {
	var eventFromResourceStacks []view.CloudFormationStackEventInventoryView
	total, err := cli.Page("v1/cloudformation/event", params, &eventFromResourceStacks)
	return eventFromResourceStacks, total, err
}

// RestartResourceStack operates on RestartResourceStack
func (cli *ZSClient) RestartResourceStack(uuid string, params param.RestartResourceStackParam) (*view.ResourceStackInventoryView, error) {
	resp := view.ResourceStackInventoryView{}
	if err := cli.PutWithRespKey("v1/cloudformation/stack", uuid, "", map[string]interface{}{
		"restartResourceStack": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteResourceStackVmPortMonitor deletes ResourceStackVmPortMonitor
func (cli *ZSClient) DeleteResourceStackVmPortMonitor(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/cloudformation/stack/monitor/delvm", uuid, string(deleteMode))
}

// AddResourceStackVmPortMonitor adds ResourceStackVmPortMonitor
func (cli *ZSClient) AddResourceStackVmPortMonitor(params param.AddResourceStackVmPortMonitorParam) (*view.AddResourceStackVmPortMonitorEventView, error) {
	resp := view.AddResourceStackVmPortMonitorEventView{}
	if err := cli.Post("v1/cloudformation/stack/monitor/addvm", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateResourceStackFromApp creates ResourceStackFromApp
func (cli *ZSClient) CreateResourceStackFromApp(params param.CreateResourceStackFromAppParam) (*view.ResourceStackInventoryView, error) {
	resp := view.ResourceStackInventoryView{}
	if err := cli.Post("v1/appcenter/app/stack", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PreviewResourceStack operates on PreviewResourceStack
func (cli *ZSClient) PreviewResourceStack(params param.PreviewResourceStackParam) (*view.PreviewResourceStackView, error) {
	resp := view.PreviewResourceStackView{}
	if err := cli.Post("v1/cloudformation/stack/preview", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceFromResourceStack gets ResourceFromResourceStack by uuid
func (cli *ZSClient) GetResourceFromResourceStack() (*view.GetResourceFromResourceStackView, error) {
	var resp view.GetResourceFromResourceStackView
	if err := cli.GetWithRespKey("v1/cloudformation/stack/resources", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSupportedCloudFormationResources gets SupportedCloudFormationResources by uuid
func (cli *ZSClient) GetSupportedCloudFormationResources() (*view.GetSupportedCloudFormationResourcesView, error) {
	var resp view.GetSupportedCloudFormationResourcesView
	if err := cli.GetWithRespKey("v1/cloudformation/resources", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceStackFromResource gets ResourceStackFromResource by uuid
func (cli *ZSClient) GetResourceStackFromResource() (*view.GetResourceStackFromResourceView, error) {
	var resp view.GetResourceStackFromResourceView
	if err := cli.GetWithRespKey("v1/cloudformation/resources/stack", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
