// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetIAM2ProjectRetirePolicy operates on SetIAM2ProjectRetirePolicy
func (cli *ZSClient) SetIAM2ProjectRetirePolicy(uuid string, params param.SetIAM2ProjectRetirePolicyParam) (*view.SetIAM2ProjectRetirePolicyEventView, error) {
	resp := view.SetIAM2ProjectRetirePolicyEventView{}
	if err := cli.Put("v1/iam2/projects/retire-policies/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
