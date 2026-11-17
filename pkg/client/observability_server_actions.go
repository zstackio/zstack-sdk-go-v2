// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// CreateObservabilityServer creates ObservabilityServer
func (cli *ZSClient) CreateObservabilityServer(params param.CreateObservabilityServerParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post("v1/observability-servers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachServiceToObservabilityServer operates on ServiceToObservabilityServer
func (cli *ZSClient) AttachServiceToObservabilityServer(observabilityServerUuid string, params param.AttachServiceToObservabilityServerParam) (*view.ObservabilityServerVmInventoryView, error) {
	resp := view.ObservabilityServerVmInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/observability-server/%s/service", observabilityServerUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachServiceFromObservabilityServer operates on ServiceFromObservabilityServer
func (cli *ZSClient) DetachServiceFromObservabilityServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/observability-server", uuid, string(deleteMode))
}

// GetObservabilityServerServiceData gets ObservabilityServerServiceData by uuid
func (cli *ZSClient) GetObservabilityServerServiceData(observabilityServerUuid string, params param.GetObservabilityServerServiceDataParam) (*view.ObservabilityServerServiceDataInventoryView, error) {
	resp := view.ObservabilityServerServiceDataInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/observability-server/%s/service-data", observabilityServerUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
