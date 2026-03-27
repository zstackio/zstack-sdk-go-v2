// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteReservedIpRange deletes ReservedIpRange
func (cli *ZSClient) DeleteReservedIpRange(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/l3-networks/reserved-ip-ranges", uuid, string(deleteMode))
}
// AddReservedIpRange adds ReservedIpRange
func (cli *ZSClient) AddReservedIpRange(ctx context.Context, params param.AddReservedIpRangeParam) (*view.ReservedIpRangeInventoryView, error) {
	resp := view.ReservedIpRangeInventoryView{}
	if err := cli.Post(ctx, "v1/l3-networks/{l3NetworkUuid}/reserved-ip-ranges", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
