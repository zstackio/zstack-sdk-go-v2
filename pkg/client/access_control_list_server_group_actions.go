// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeAccessControlListServerGroup 操作AccessControlListServerGroup
func (cli *ZSClient) ChangeAccessControlListServerGroup(uuid string, params param.ChangeAccessControlListServerGroupParam) (*view.ChangeAccessControlListServerGroupEventView, error) {
	resp := view.ChangeAccessControlListServerGroupEventView{}
	if err := cli.Put("v1/load-balancers/listener/acl/{aclUuid}/servergroup/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

