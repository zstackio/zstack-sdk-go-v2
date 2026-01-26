// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateObservabilityServerOffering creates ObservabilityServerOffering
func (cli *ZSClient) CreateObservabilityServerOffering(params param.CreateObservabilityServerOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	resp := view.InstanceOfferingInventoryView{}
	if err := cli.Post("v1/instance-offerings/observability-servers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
