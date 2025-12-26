// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddSNSFeiShuAtPerson adds SNSFeiShuAtPerson
func (cli *ZSClient) AddSNSFeiShuAtPerson(params param.AddSNSFeiShuAtPersonParam) (*view.AddSNSFeiShuAtPersonEventView, error) {
	resp := view.AddSNSFeiShuAtPersonEventView{}
	if err := cli.Post("v1/sns/application-endpoints/feishu/at-persons", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
