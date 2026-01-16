// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVolume(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryVolume(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVolume error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryVolume result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s\t%d", r.UUID, r.Name, r.State, r.Type, r.Size)
	}
	golog.Infof("======================================")
}

func TestQueryVolume2(t *testing.T) {
	// Query with conditions - Root volumes
	params := param.NewQueryParam()
	params.AddQ("type=Root")
	params.AddQ("status=Ready")
	params.Start(0).Limit(10).ReplyWithCount(true)
	result, err := accessKeyAuthCli.QueryVolume(&params)
	if err != nil {
		t.Errorf("TestQueryVolume2 error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("Found %d Ready Root Volumes:", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%d MB", r.UUID, r.Name, r.Format, r.Size/1024/1024)
	}
	golog.Infof("======================================")
}

func TestQueryVolume3(t *testing.T) {
	// Query with conditions - Data volumes
	params := param.NewQueryParam()
	params.AddQ("type=Data")
	params.Limit(5)
	result, err := accessKeyAuthCli.QueryVolume(&params)
	if err != nil {
		t.Errorf("TestQueryVolume3 error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("Found %d Data Volumes:", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%d MB", r.UUID, r.Name, r.Status, r.Size/1024/1024)
	}
	golog.Infof("======================================")
}

func TestPageVolume(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageVolume(&queryParam)
	if err != nil {
		t.Errorf("TestPageVolume error: %v", err)
		return
	}
	golog.Infof("PageVolume result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s", r.UUID, r.Name, r.State, r.Type)
	}
}

func TestGetVolume(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVolume(&queryParam)
	if err != nil {
		t.Errorf("TestGetVolume Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Volume found to test Get")
		return
	}

	// Get by UUID
	result, err := accessKeyAuthCli.GetVolume(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVolume error: %v", err)
		return
	}
	golog.Infof("GetVolume result: %s, Name: %s, Type: %s", result.UUID, result.Name, result.Type)
}
