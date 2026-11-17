// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

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

// PageSNSEmailAddress Pagination
func (cli *ZSClient) PageSNSEmailAddress(params *param.QueryParam) ([]view.SNSEmailAddressInventoryView, int, error) {
	var sNSEmailAddress []view.SNSEmailAddressInventoryView
	total, err := cli.Page("v1/sns/application-endpoints/emails/email-addresses", params, &sNSEmailAddress)
	return sNSEmailAddress, total, err
}

// AddEmailAddressToSNSEmailEndpoint adds EmailAddressToSNSEmailEndpoint
func (cli *ZSClient) AddEmailAddressToSNSEmailEndpoint(params param.AddEmailAddressToSNSEmailEndpointParam) (*view.SNSEmailAddressInventoryView, error) {
	resp := view.SNSEmailAddressInventoryView{}
	if err := cli.Post("v1/sns/application-endpoints/emails/email-addresses", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteEmailAddressOfSNSEmailEndpoint deletes EmailAddressOfSNSEmailEndpoint
func (cli *ZSClient) DeleteEmailAddressOfSNSEmailEndpoint(endpointUuid string, emailAddressUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/sns/application-endpoints/emails", endpointUuid, fmt.Sprintf("email-addresses/%s", emailAddressUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// UpdateEmailAddressOfSNSEmailEndpoint updates EmailAddressOfSNSEmailEndpoint
func (cli *ZSClient) UpdateEmailAddressOfSNSEmailEndpoint(params param.UpdateEmailAddressOfSNSEmailEndpointParam) (*view.SNSEmailAddressInventoryView, error) {
	resp := view.SNSEmailAddressInventoryView{}
	if err := cli.PutWithRespKey("v1/sns/application-endpoints/emails/email-addresses", "", "", map[string]interface{}{
		"updateEmailAddressOfSNSEmailEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
