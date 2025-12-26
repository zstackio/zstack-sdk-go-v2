// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeSlbGroupDeployType changes SlbGroupDeployType
func (cli *ZSClient) ChangeSlbGroupDeployType(uuid string, params param.ChangeSlbGroupDeployTypeParam) (*view.ChangeSlbGroupDeployTypeEventView, error) {
	resp := view.ChangeSlbGroupDeployTypeEventView{}
	if err := cli.Put("v1/load-balancers/slb/groups/{slbGroupUuid}/deployType", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
