// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSSORedirectTemplate creates SSORedirectTemplate
func (cli *ZSClient) CreateSSORedirectTemplate(params param.CreateSSORedirectTemplateParam) (*view.SSORedirectTemplateInventoryView, error) {
	var resp view.CreateSSORedirectTemplateEventView
	if err := cli.Post("v1/create/sso/redirect/template/", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateSSORedirectTemplate updates SSORedirectTemplate
func (cli *ZSClient) UpdateSSORedirectTemplate(uuid string, params param.UpdateSSORedirectTemplateParam) (*view.SSORedirectTemplateInventoryView, error) {
	var resp view.UpdateSSORedirectTemplateEventView
	if err := cli.Put("v1/update/sso/redirectTemplate", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteSSORedirectTemplate deletes SSORedirectTemplate
func (cli *ZSClient) DeleteSSORedirectTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/delete/sso/redirect/template", uuid, string(deleteMode))
}
