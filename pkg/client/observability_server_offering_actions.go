// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateObservabilityServerOffering creates ObservabilityServerOffering
func (cli *ZSClient) CreateObservabilityServerOffering(params param.CreateObservabilityServerOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	var resp view.CreateInstanceOfferingEventView
	if err := cli.Post("v1/instance-offerings/observability-servers", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
