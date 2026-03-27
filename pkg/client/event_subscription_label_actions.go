// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateEventSubscriptionLabel updates EventSubscriptionLabel
func (cli *ZSClient) UpdateEventSubscriptionLabel(ctx context.Context, uuid string, params param.UpdateEventSubscriptionLabelParam) (*view.EventSubscriptionLabelInventoryView, error) {
	resp := view.EventSubscriptionLabelInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/zwatch/events/subscriptions/labels", uuid, "", map[string]interface{}{
		"updateEventSubscriptionLabel": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
