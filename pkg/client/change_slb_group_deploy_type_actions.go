// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeSlbGroupDeployType changes SlbGroupDeployType
func (cli *ZSClient) ChangeSlbGroupDeployType(uuid string, params param.ChangeSlbGroupDeployTypeParam) (*view.ChangeSlbGroupDeployTypeEventView, error) {
	resp := view.ChangeSlbGroupDeployTypeEventView{}
	if err := cli.Put("v1/load-balancers/slb/groups/{slbGroupUuid}/deployType", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
