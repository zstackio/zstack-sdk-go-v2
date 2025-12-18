// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ExpungeImage 操作Image
func (cli *ZSClient) ExpungeImage(uuid string, params param.ExpungeImageParam) (*view.ExpungeImageEventView, error) {
	resp := view.ExpungeImageEventView{}
	if err := cli.Put("v1/images/{imageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

