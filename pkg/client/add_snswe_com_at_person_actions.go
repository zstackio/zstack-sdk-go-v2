// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddSNSWeComAtPerson 操作AddSNSWeComAtPerson
func (cli *ZSClient) AddSNSWeComAtPerson(params param.AddSNSWeComAtPersonParam) (*view.AddSNSWeComAtPersonEventView, error) {
	resp := view.AddSNSWeComAtPersonEventView{}
	if err := cli.Post("v1/sns/application-endpoints/we-com/at-persons", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

