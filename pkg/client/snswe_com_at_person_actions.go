// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddSNSWeComAtPerson adds SNSWeComAtPerson
func (cli *ZSClient) AddSNSWeComAtPerson(params param.AddSNSWeComAtPersonParam) (*view.SNSWeComAtPersonInventoryView, error) {
	resp := view.SNSWeComAtPersonInventoryView{}
	if err := cli.Post("v1/sns/application-endpoints/we-com/at-persons", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySNSWeComAtPerson queries SNSWeComAtPerson list
func (cli *ZSClient) QuerySNSWeComAtPerson(params *param.QueryParam) ([]view.SNSWeComAtPersonInventoryView, error) {
	var resp []view.SNSWeComAtPersonInventoryView
	return resp, cli.List("v1/sns/application-endpoints/we-com/at-persons", params, &resp)
}

func (cli *ZSClient) GetSNSWeComAtPerson(uuid string) (*view.SNSWeComAtPersonInventoryView, error) {
	var resp view.SNSWeComAtPersonInventoryView
	if err := cli.Get("v1/sns/application-endpoints/we-com/at-persons", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSWeComAtPerson Pagination
func (cli *ZSClient) PageSNSWeComAtPerson(params *param.QueryParam) ([]view.SNSWeComAtPersonInventoryView, int, error) {
	var sNSWeComAtPersons []view.SNSWeComAtPersonInventoryView
	total, err := cli.Page("v1/sns/application-endpoints/we-com/at-persons", params, &sNSWeComAtPersons)
	return sNSWeComAtPersons, total, err
}
// RemoveSNSWeComAtPerson removes SNSWeComAtPerson
func (cli *ZSClient) RemoveSNSWeComAtPerson(endpointUuid string, userId string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/sns/application-endpoints/we-com", endpointUuid, fmt.Sprintf("at-persons/%s", userId), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}
