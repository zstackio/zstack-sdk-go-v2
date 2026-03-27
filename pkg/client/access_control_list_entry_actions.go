// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// RemoveAccessControlListEntry removes AccessControlListEntry
func (cli *ZSClient) RemoveAccessControlListEntry(ctx context.Context, aclUuid string, uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/access-control-lists", aclUuid, fmt.Sprintf("ipentries/%s", uuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}
// AddAccessControlListEntry adds AccessControlListEntry
func (cli *ZSClient) AddAccessControlListEntry(ctx context.Context, params param.AddAccessControlListEntryParam) (*view.AccessControlListEntryInventoryView, error) {
	resp := view.AccessControlListEntryInventoryView{}
	if err := cli.Post(ctx, "v1/access-control-lists/{aclUuid}/ipentries", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
