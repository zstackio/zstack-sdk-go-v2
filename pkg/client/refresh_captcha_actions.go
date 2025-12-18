// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RefreshCaptcha operates on RefreshCaptcha
func (cli *ZSClient) RefreshCaptcha(params param.RefreshCaptchaParam) (*view.RefreshCaptchaView, error) {
	var resp view.RefreshCaptchaView
	if err := cli.Get("v1/captcha/refresh", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
