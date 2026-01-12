// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateHostIpmi updates HostIpmi
func (cli *ZSClient) UpdateHostIpmi(uuid string, params param.UpdateHostIpmiParam) (*view.HostIpmiInventoryView, error) {
	resp := view.HostIpmiInventoryView{}
	err := cli.PutWithSpec("v1/hosts/ipmi", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
