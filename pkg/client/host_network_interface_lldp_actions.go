// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// GetHostNetworkInterfaceLldp gets HostNetworkInterfaceLldp by uuid
func (cli *ZSClient) GetHostNetworkInterfaceLldp(uuid string) (*view.HostNetworkInterfaceLldpInventoryView, error) {
	var resp view.HostNetworkInterfaceLldpInventoryView
	if err := cli.GetWithRespKey("v1/hostNetworkInterface/lldp", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
