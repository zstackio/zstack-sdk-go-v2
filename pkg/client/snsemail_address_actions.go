// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySNSEmailAddress queries SNSEmailAddress list
func (cli *ZSClient) QuerySNSEmailAddress(params *param.QueryParam) ([]view.SNSEmailAddressInventoryView, error) {
	var resp []view.SNSEmailAddressInventoryView
	return resp, cli.List("v1/sns/application-endpoints/emails/email-addresses", params, &resp)
}

func (cli *ZSClient) GetSNSEmailAddress(uuid string) (*view.SNSEmailAddressInventoryView, error) {
	var resp view.SNSEmailAddressInventoryView
	if err := cli.Get("v1/sns/application-endpoints/emails/email-addresses", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
