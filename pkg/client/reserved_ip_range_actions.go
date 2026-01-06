// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteReservedIpRange deletes ReservedIpRange
func (cli *ZSClient) DeleteReservedIpRange(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks/reserved-ip-ranges/{uuid}", uuid, string(deleteMode))
}
// AddReservedIpRange adds ReservedIpRange
func (cli *ZSClient) AddReservedIpRange(params param.AddReservedIpRangeParam) (*view.ReservedIpRangeInventoryView, error) {
	var resp view.AddReservedIpRangeEventView
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/reserved-ip-ranges", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
