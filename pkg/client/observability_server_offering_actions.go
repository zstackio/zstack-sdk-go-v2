// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateObservabilityServerOffering creates ObservabilityServerOffering
func (cli *ZSClient) CreateObservabilityServerOffering(ctx context.Context, params param.CreateObservabilityServerOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	resp := view.InstanceOfferingInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/instance-offerings/observability-servers", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
