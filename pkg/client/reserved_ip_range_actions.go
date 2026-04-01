// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteReservedIpRange deletes ReservedIpRange
func (cli *ZSClient) DeleteReservedIpRange(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks/reserved-ip-ranges", uuid, string(deleteMode))
}
// AddReservedIpRange adds ReservedIpRange
func (cli *ZSClient) AddReservedIpRange(l3NetworkUuid string, params param.AddReservedIpRangeParam) (*view.ReservedIpRangeInventoryView, error) {
	resp := view.ReservedIpRangeInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/l3-networks/%s/reserved-ip-ranges", l3NetworkUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
