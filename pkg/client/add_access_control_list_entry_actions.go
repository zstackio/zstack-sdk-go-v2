// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddAccessControlListEntry adds AccessControlListEntry
func (cli *ZSClient) AddAccessControlListEntry(params param.AddAccessControlListEntryParam) (*view.AddAccessControlListEntryEventView, error) {
	resp := view.AddAccessControlListEntryEventView{}
	if err := cli.Post("v1/access-control-lists/{aclUuid}/ipentries", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
