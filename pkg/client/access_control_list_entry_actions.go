// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// RemoveAccessControlListEntry removes AccessControlListEntry
func (cli *ZSClient) RemoveAccessControlListEntry(aclUuid string, uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/access-control-lists", fmt.Sprintf("%s/ipentries/%s", aclUuid, uuid), string(deleteMode))
}
// AddAccessControlListEntry adds AccessControlListEntry
func (cli *ZSClient) AddAccessControlListEntry(params param.AddAccessControlListEntryParam) (*view.AccessControlListEntryInventoryView, error) {
	var resp view.AddAccessControlListEntryEventView
	if err := cli.Post("v1/access-control-lists/{aclUuid}/ipentries", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
