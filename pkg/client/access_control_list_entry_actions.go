// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// RemoveAccessControlListEntry removes AccessControlListEntry
func (cli *ZSClient) RemoveAccessControlListEntry(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/access-control-lists/{aclUuid}/ipentries/{uuid}", uuid, string(deleteMode))
}
// AddAccessControlListEntry adds AccessControlListEntry
func (cli *ZSClient) AddAccessControlListEntry(params param.AddAccessControlListEntryParam) (*view.AccessControlListEntryInventoryView, error) {
	var resp view.AddAccessControlListEntryEventView
	if err := cli.Post("v1/access-control-lists/{aclUuid}/ipentries", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
