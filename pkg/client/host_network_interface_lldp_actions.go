// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// GetHostNetworkInterfaceLldp gets HostNetworkInterfaceLldp by uuid
func (cli *ZSClient) GetHostNetworkInterfaceLldp(interfaceUuid string) (*view.HostNetworkInterfaceLldpInventoryView, error) {
	var resp view.HostNetworkInterfaceLldpInventoryView
	err := cli.GetWithSpec("v1/hostNetworkInterface/lldp", fmt.Sprintf(\"%s/info\", interfaceUuid), "", "", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
