// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddSNSFeiShuAtPerson adds SNSFeiShuAtPerson
func (cli *ZSClient) AddSNSFeiShuAtPerson(params param.AddSNSFeiShuAtPersonParam) (*view.AddSNSFeiShuAtPersonEventView, error) {
	resp := view.AddSNSFeiShuAtPersonEventView{}
	if err := cli.Post("v1/sns/application-endpoints/feishu/at-persons", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
