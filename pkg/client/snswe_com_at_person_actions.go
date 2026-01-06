// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddSNSWeComAtPerson adds SNSWeComAtPerson
func (cli *ZSClient) AddSNSWeComAtPerson(params param.AddSNSWeComAtPersonParam) (*view.SNSWeComAtPersonInventoryView, error) {
	var resp view.AddSNSWeComAtPersonEventView
	if err := cli.Post("v1/sns/application-endpoints/we-com/at-persons", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QuerySNSWeComAtPerson queries SNSWeComAtPerson list
func (cli *ZSClient) QuerySNSWeComAtPerson(params *param.QueryParam) ([]view.SNSWeComAtPersonInventoryView, error) {
	var resp []view.SNSWeComAtPersonInventoryView
	return resp, cli.List("v1/sns/application-endpoints/we-com/at-persons", params, &resp)
}
// RemoveSNSWeComAtPerson removes SNSWeComAtPerson
func (cli *ZSClient) RemoveSNSWeComAtPerson(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/application-endpoints/we-com/{endpointUuid}/at-persons/{userId}", uuid, string(deleteMode))
}
