// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateHostIpmi updates HostIpmi
func (cli *ZSClient) UpdateHostIpmi(ctx context.Context, uuid string, params param.UpdateHostIpmiParam) (*view.HostIpmiInventoryView, error) {
	resp := view.HostIpmiInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hosts/ipmi", uuid, "", map[string]interface{}{
		"updateHostIpmi": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
