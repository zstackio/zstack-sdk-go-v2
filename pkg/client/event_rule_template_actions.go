// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddEventRuleTemplate adds EventRuleTemplate
func (cli *ZSClient) AddEventRuleTemplate(params param.AddEventRuleTemplateParam) (*view.EventRuleTemplateInventoryView, error) {
	var resp view.AddEventRuleTemplateEventView
	if err := cli.Post("v1/zwatch/monitortemplates/evenrules", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteEventRuleTemplate deletes EventRuleTemplate
func (cli *ZSClient) DeleteEventRuleTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/monitortemplates/evenrules/{uuid}", uuid, string(deleteMode))
}
// UpdateEventRuleTemplate updates EventRuleTemplate
func (cli *ZSClient) UpdateEventRuleTemplate(uuid string, params param.UpdateEventRuleTemplateParam) (*view.EventRuleTemplateInventoryView, error) {
	var resp view.UpdateEventRuleTemplateEventView
	if err := cli.Put("v1/zwatch/monitortemplates/evenrules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryEventRuleTemplate queries EventRuleTemplate list
func (cli *ZSClient) QueryEventRuleTemplate(params *param.QueryParam) ([]view.EventRuleTemplateInventoryView, error) {
	var resp []view.EventRuleTemplateInventoryView
	return resp, cli.List("v1/zwatch/monitortemplates/evenrules", params, &resp)
}
