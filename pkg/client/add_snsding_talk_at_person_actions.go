// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddSNSDingTalkAtPerson adds SNSDingTalkAtPerson
func (cli *ZSClient) AddSNSDingTalkAtPerson(params param.AddSNSDingTalkAtPersonParam) (*view.AddSNSDingTalkAtPersonEventView, error) {
	resp := view.AddSNSDingTalkAtPersonEventView{}
	if err := cli.Post("v1/sns/application-endpoints/ding-talk/at-persons", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
