// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddAccessControlListEntry 操作AddAccessControlListEntry
func (cli *ZSClient) AddAccessControlListEntry(params param.AddAccessControlListEntryParam) (*view.AddAccessControlListEntryEventView, error) {
	resp := view.AddAccessControlListEntryEventView{}
	if err := cli.Post("v1/access-control-lists/{aclUuid}/ipentries", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

