// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSSORedirectTemplate creates SSORedirectTemplate
func (cli *ZSClient) CreateSSORedirectTemplate(params param.CreateSSORedirectTemplateParam) (*view.SSORedirectTemplateInventoryView, error) {
	resp := view.SSORedirectTemplateInventoryView{}
	if err := cli.Post("v1/create/sso/redirect/template/", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateSSORedirectTemplate updates SSORedirectTemplate
func (cli *ZSClient) UpdateSSORedirectTemplate(uuid string, params param.UpdateSSORedirectTemplateParam) (*view.SSORedirectTemplateInventoryView, error) {
	resp := view.SSORedirectTemplateInventoryView{}
	if err := cli.Put("v1/update/sso/redirectTemplate", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteSSORedirectTemplate deletes SSORedirectTemplate
func (cli *ZSClient) DeleteSSORedirectTemplate(params param.DeleteSSORedirectTemplateParam) (*view.SSORedirectTemplateInventoryView, error) {
	resp := view.SSORedirectTemplateInventoryView{}
	if err := cli.Post("v1/delete/sso/redirect/template", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
