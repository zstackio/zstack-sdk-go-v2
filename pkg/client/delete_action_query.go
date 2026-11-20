// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

// deleteActionWithQuery validates required fields and sends collection entries
// as repeated keys, matching the REST server's DELETE query decoder.
func (cli *ZSClient) deleteActionWithQuery(resource, uuid, spec string, deleteMode param.DeleteMode, query url.Values, required ...string) error {
	if spec != "" || uuid != "" {
		var err error
		uuid, err = actionPathSegment("resource UUID", uuid)
		if err != nil {
			return err
		}
	}
	for _, key := range required {
		values := query[key]
		if len(values) == 0 {
			return fmt.Errorf("%s is required", key)
		}
		for _, value := range values {
			if strings.TrimSpace(value) == "" {
				return fmt.Errorf("%s must not contain empty values", key)
			}
		}
	}
	if deleteMode != "" {
		query.Set("deleteMode", string(deleteMode))
	}
	return cli.DeleteWithSpec(resource, uuid, spec, query.Encode(), nil)
}

func (cli *ZSClient) deleteResourceWithQuery(resource, uuid string, deleteMode param.DeleteMode, query url.Values, required ...string) error {
	if uuid == "" {
		return fmt.Errorf("resource UUID is required")
	}
	return cli.deleteActionWithQuery(resource, uuid, "", deleteMode, query, required...)
}

func actionPathSegment(field, value string) (string, error) {
	if strings.TrimSpace(value) == "" || value == "." || value == ".." || strings.Contains(value, "/") {
		return "", fmt.Errorf("%s must be a nonempty path segment", field)
	}
	return url.PathEscape(value), nil
}

func deleteModeOrDefault(mode param.DeleteMode, fallback *string) param.DeleteMode {
	if mode == "" && fallback != nil {
		return param.DeleteMode(*fallback)
	}
	return mode
}
