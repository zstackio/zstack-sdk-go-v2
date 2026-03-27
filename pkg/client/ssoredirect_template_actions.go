// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSSORedirectTemplate creates SSORedirectTemplate
func (cli *ZSClient) CreateSSORedirectTemplate(ctx context.Context, params param.CreateSSORedirectTemplateParam) (*view.SSORedirectTemplateInventoryView, error) {
	resp := view.SSORedirectTemplateInventoryView{}
	if err := cli.Post(ctx, "v1/create/sso/redirect/template/", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateSSORedirectTemplate updates SSORedirectTemplate
func (cli *ZSClient) UpdateSSORedirectTemplate(ctx context.Context, params param.UpdateSSORedirectTemplateParam) (*view.SSORedirectTemplateInventoryView, error) {
	resp := view.SSORedirectTemplateInventoryView{}
	if err := cli.Post(ctx, "v1/update/sso/redirectTemplate", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteSSORedirectTemplate deletes SSORedirectTemplate
func (cli *ZSClient) DeleteSSORedirectTemplate(params param.DeleteSSORedirectTemplateParam) (*view.DeleteSSORedirectTemplateEventView, error) {
	resp := view.DeleteSSORedirectTemplateEventView{}
	if err := cli.PostWithRespKey("v1/delete/sso/redirect/template", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
