// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddEmailAddressToSNSEmailEndpoint 操作AddEmailAddressToSNSEmailEndpoint
func (cli *ZSClient) AddEmailAddressToSNSEmailEndpoint(params param.AddEmailAddressToSNSEmailEndpointParam) (*view.AddEmailAddressToSNSEmailEndpointEventView, error) {
	resp := view.AddEmailAddressToSNSEmailEndpointEventView{}
	if err := cli.Post("v1/sns/application-endpoints/emails/email-addresses", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

