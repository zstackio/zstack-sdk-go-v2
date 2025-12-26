// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeAccessControlListServerGroup changes AccessControlListServerGroup
func (cli *ZSClient) ChangeAccessControlListServerGroup(uuid string, params param.ChangeAccessControlListServerGroupParam) (*view.ChangeAccessControlListServerGroupEventView, error) {
	resp := view.ChangeAccessControlListServerGroupEventView{}
	if err := cli.Put("v1/load-balancers/listener/acl/{aclUuid}/servergroup/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
