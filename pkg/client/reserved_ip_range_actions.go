// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteReservedIpRange deletes ReservedIpRange
func (cli *ZSClient) DeleteReservedIpRange(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/l3-networks/reserved-ip-ranges", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// AddReservedIpRange adds ReservedIpRange
func (cli *ZSClient) AddReservedIpRange(params param.AddReservedIpRangeParam) (*view.ReservedIpRangeInventoryView, error) {
	var resp view.AddReservedIpRangeEventView
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/reserved-ip-ranges", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
