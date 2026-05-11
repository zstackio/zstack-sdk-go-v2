// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// GetHostNetworkInterfaceLldp gets HostNetworkInterfaceLldp by uuid
func (cli *ZSClient) GetHostNetworkInterfaceLldp(ctx context.Context, uuid string) (*view.GetHostNetworkInterfaceLldpView, error) {
	var resp view.GetHostNetworkInterfaceLldpView
	if err := cli.GetWithRespKey(ctx, "v1/hostNetworkInterface/lldp", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
