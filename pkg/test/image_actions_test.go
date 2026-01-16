// Copyright (c) ZStack.io, Inc.

package test

import (
	"fmt"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryImage1(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryImage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryImage error: %v", err)
		return
	}
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s\t%s\t%s\t%d\t%s", r.UUID, r.Name, r.Platform, r.GuestOsType, r.Format, r.Status, r.Size, r.Description)
	}
	golog.Infof("======================================")

	queryParam.AddQ(fmt.Sprintf("platform=%s", "Linux"))
	queryParam.Start(0).Limit(10).ReplyWithCount(true)
	result, err = accountLoginCli.QueryImage(&queryParam)
	if err != nil {
		golog.Errorf("ZSClient.QueryImage error:%v", err)
		return
	}
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s\t%s\t%s\t%d\t%s", r.UUID, r.Name, r.Platform, r.GuestOsType, r.Format, r.Status, r.Size, r.Description)
	}
}

func TestQueryImage2(t *testing.T) {
	params := param.NewQueryParam()
	//params.AddQ("state=Enabled")
	//params.AddQ("type=zstack")
	//params.AddQ("format!=vmtx")
	//params.AddQ("status=Ready")
	//params.AddQ("system=false")
	params.AddQ("name=centos-test")
	//params.AddQ("mediaType=DataVolumeTemplate")
	//params.AddQ("backupStorage.zone.uuid=6e8191bfd57745f282f78cb013b732b6")
	result, err := accountLoginCli.QueryImage(&params)
	if err != nil {
		golog.Errorf("ZSClient.QueryImage error:%v", err)
		return
	}
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s\t%s\t%s\t%d\t%s", r.UUID, r.Name, r.Platform, r.GuestOsType, r.Format, r.Status, r.Size, r.Description)
	}
	golog.Infof("======================================")
}

func TestQueryImage3(t *testing.T) {
	params := param.NewQueryParam()
	//params.AddQ("state=Enabled")
	//params.AddQ("type=zstack")
	//params.AddQ("format!=vmtx")
	//params.AddQ("status=Ready")
	//params.AddQ("system=false")
	params.AddQ("format!=vmtx")
	params.AddQ("system=true")
	//params.AddQ("mediaType=DataVolumeTemplate")
	//params.AddQ("backupStorage.zone.uuid=6e8191bfd57745f282f78cb013b732b6")
	result, err := accountLoginCli.QueryImage(&params)
	if err != nil {
		golog.Errorf("ZSClient.QueryImage error:%v", err)
		return
	}
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s\t%s\t%s\t%d\t%s", r.UUID, r.Name, r.Platform, r.GuestOsType, r.Format, r.Status, r.Size, r.Description)
	}
	golog.Infof("======================================")
}

type Image struct {
	Name        string `json:"name"`
	UUID        string `json:"uuid"`
	Format      string `json:"format"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Status      string `json:"status"`
	State       string `json:"state"`
}

type QueryResult struct {
	Results []struct {
		Inventories []Image `json:"inventories"`
	} `json:"results"`
}

func TestZSClient_QueryByZql(t *testing.T) {

	//var reservedIpRanges []view.ReservedIpRangeInventoryView
	var queryResult QueryResult
	//	_, err := accountLoginCli.Zql(fmt.Sprintf("query Image "), &virtualRouterImages, "inventories")

	_, err := accountLoginCli.Zql(
		"query Image where system='true'",
		&queryResult,
	)
	if err != nil {
		golog.Errorf("failed to execute ZQL query: %v", err)
	}

	//_, err := accountLoginCli.Zql(fmt.Sprintf("query Image"), &virtualRouterImages, "inventories")
	// 提取结果
	if len(queryResult.Results) > 0 {
		inventories := queryResult.Results[0].Inventories
		fmt.Printf("Query Response: %+v\n", inventories)
	} else {
		fmt.Println("No inventories found.")
	}

}

func TestPageImage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, total, err := accountLoginCli.PageImage(&queryParam)
	if err != nil {
		t.Errorf("TestPageImage error: %v", err)
		return
	}
	golog.Infof("PageImage result count: %d, total: %d", len(result), total)
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s\t%s\t%s\t%d\t%s %d", r.UUID, r.Name, r.Platform, r.GuestOsType, r.Format, r.Status, r.Size, r.Description, total)
	}
}

func TestGetImage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryImage(&queryParam)
	if err != nil {
		t.Errorf("TestGetImage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Image found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetImage(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetImage error: %v", err)
		return
	}
	golog.Infof("GetImage result: %s", result.UUID)
}

/*
func TestUpdateImage(t *testing.T) {
	// First query to get a valid UUID
	desc := "test from sdk-chi"

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryImage(&queryParam)
	if err != nil {
		t.Errorf("TestGetImage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Image found to test Get")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateImageParam{
		Params: param.UpdateImageParamDetail{
			Uuid:        list[0].UUID,
			Name:        "sdk-test-0115",
			Description: desc,
			// Keep original values - just testing the API works
		},
	}

	v, err := accountLoginCli.UpdateImage(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateImage error: %v", err)
		return
	}
	golog.Infof("UpdateImage succeeded for UUID: %s", v.UUID)
}

func TestDeleteImage(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	//t.Skip("TestDeleteImage is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryImage(&queryParam)
	if err != nil {
		t.Errorf("TestGetImage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Image found to test Get")
		return
	}

	err = accountLoginCli.DeleteImage(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteImage error: %v", err)
		return
	}
	golog.Infof("DeleteImage succeeded for UUID: %s", list[0].UUID)
}

func TestExpungeImageByUUID(t *testing.T) {

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	queryParam.AddQ("status=Deleted")
	list, err := accountLoginCli.QueryImage(&queryParam)
	if err != nil {
		t.Errorf("TestGetImage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Image found to test Get")
		return
	}

	err = accountLoginCli.ExpungeImage(list[0].UUID)
	if err != nil {
		t.Errorf("TestExpungeImageByUUID error: %v", err)
		return
	}
	golog.Infof("ExpungeImageByUUID succeeded for UUID: %s", list[0].UUID)

}

func TestAddImage(t *testing.T) {
	// Add operation - similar to Create
	storage, err := accountLoginCli.QueryBackupStorage(&param.QueryParam{})
	if err != nil {
		t.Errorf("QueryBackupStorage error: %v", err)
		return
	}
	if len(storage) == 0 {
		t.Skip("No BackupStorage found to add image")
		return
	}

	imageParam := param.AddImageParam{
		BaseParam: param.BaseParam{
			SystemTags: []string{"bootMode::Legacy"},
		},
		Params: param.AddImageParamDetail{
			Name:               "esxi-6.7.0-ks.iso",
			Description:        "接口image",
			Url:                "http://192.168.200.100/mirror/jiajian.chi/os/base/esxi-6.7.0-ks.iso",
			MediaType:          "RootVolumeTemplate",
			GuestOsType:        "Linux",
			System:             false,
			Format:             "iso",
			Platform:           "Linux",
			BackupStorageUuids: []string{storage[0].UUID},
			Type:               "",
			ResourceUuid:       "",
			Architecture:       "x86_64",
			Virtio:             false,
		},
	}

	result, err := accountLoginCli.AddImage(imageParam)
	if err != nil {
		t.Errorf("AddImage error: %v", err)
		return
	}

	golog.Infof("AddImage succeeded, UUID: %s, Name: %s", result.UUID, result.Name)
	//创建失败情况

		imageParam = param.AddImageParam{
			BaseParam: param.BaseParam{
				SystemTags: []string{"bootMode::Legacy"},
			},
			Params: param.AddImageDetailParam{
				Name:               "image-4",
				Description:        "接口image",
				Url:                "https://image.baidu.com/search/detail?tn=baiduimagedetail&word=%E6%B8%90%E5%8F%98%E9%A3%8E%E6%A0%BC%E6%8F%92%E7%94%BB&album_tab=%E8%AE%BE%E8%AE%A1%E7%B4%A0%E6%9D%90&album_id=409&ie=utf-8&fr=albumsdetail&cs=4036010509,3445021118&pi=144521&pn=1&ic=0&objurl=https%3A%2F%2Ft7.baidu.com%2Fit%2Fu%3D4036010509%2C3445021118%26fm%3D193%26f%3DGIF",
				MediaType:          param.RootVolumeTemplate,
				GuestOsType:        "Windows 10",
				System:             false,
				Format:             param.Qcow2,
				Platform:           "Windows",
				BackupStorageUuids: []string{"26684790e4734a0bbb506f40907f57da"},
				Type:               "",
				ResourceUuid:       "",
				Architecture:       "x86_64",
				Virtio:             false,
			},
		}

		_, err = accountLoginCli.AddImage(imageParam)
		if err != nil {
			golog.Errorf("ZSClient.CreateImage error:%v", err)
		}

		golog.Infof("======================================")
		//删除
		err = accountLoginCli.DeleteImage(r.UUID, param.DeleteModeEnforcing)
		if err != nil {
			golog.Errorf("ZSClient.DeleteImage error:%v", err)
		}

		//彻底删除
		err = accountLoginCli.ExpungeImage(r.UUID)
		if err != nil {
			golog.Errorf("ZSClient.ExpungeImage error:%v", err)
		}

}

func TestChangeImageState(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryImage(&queryParam)
	if err != nil {
		t.Errorf("TestGetImage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Image found to test Get")
		return
	}

	state, err := accountLoginCli.ChangeImageState(list[0].UUID, param.ChangeImageStateParam{
		BaseParam: param.BaseParam{},
		Params: param.ChangeImageStateParamDetail{
			Uuid:       list[0].UUID,
			StateEvent: "enable",
		},
	})
	if err != nil {
		golog.Errorf("TestZSClient_UpdateVirtio error:%v", err)
	}
	fmt.Println(state)
}

func TestAddImageAsync(t *testing.T) {
	// Query backup storage first
	storage, err := accountLoginCli.QueryBackupStorage(&param.QueryParam{})
	if err != nil {
		t.Errorf("QueryBackupStorage error: %v", err)
		return
	}
	if len(storage) == 0 {
		t.Skip("No BackupStorage found to add image")
		return
	}

	imageParam := param.AddImageParam{
		BaseParam: param.BaseParam{
			SystemTags: []string{"bootMode::Legacy"},
		},
		Params: param.AddImageParamDetail{
			Name:               "async-test-image",
			Description:        "Test async image creation",
			Url:                "http://192.168.200.100/mirror/jiajian.chi/os/base/esxi-6.7.0-ks.iso",
			MediaType:          "RootVolumeTemplate",
			GuestOsType:        "Linux",
			System:             false,
			Format:             "iso",
			Platform:           "Linux",
			BackupStorageUuids: []string{storage[0].UUID},
			Architecture:       "x86_64",
			Virtio:             false,
		},
	}

	apiId, err := accountLoginCli.AddImageAsync(imageParam)
	if err != nil {
		t.Errorf("AddImageAsync error: %v", err)
		return
	}
	golog.Infof("AddImageAsync succeeded, API Job ID: %s", apiId)
}

func TestSyncImage(t *testing.T) {
	// Note: SyncImage API may not be available in all ZStack versions
	t.Skip("SyncImage API not supported in current ZStack version")

	// Query image store backup storage
	storage, err := accountLoginCli.QueryBackupStorage(&param.QueryParam{})
	if err != nil {
		t.Errorf("QueryBackupStorage error: %v", err)
		return
	}
	if len(storage) == 0 {
		t.Skip("No BackupStorage found to sync image")
		return
	}

	// Find an ImageStoreBackupStorage
	var imageStoreUuid string
	for _, s := range storage {
		if s.Type == "ImageStoreBackupStorage" {
			imageStoreUuid = s.UUID
			break
		}
	}
	if imageStoreUuid == "" {
		t.Skip("No ImageStoreBackupStorage found to sync image")
		return
	}

	syncParam := param.SyncImageParam{
		BaseParam: param.BaseParam{},
		Params: param.SyncImageParamDetail{
			ImageStoreUuid: imageStoreUuid,
		},
	}

	result, err := accountLoginCli.SyncImage(imageStoreUuid, syncParam)
	if err != nil {
		t.Errorf("SyncImage error: %v", err)
		return
	}
	golog.Infof("SyncImage succeeded: %v", result)
}

func TestRecoverImage(t *testing.T) {
	// Query deleted images
	queryParam := param.NewQueryParam()
	queryParam.AddQ("status=Deleted")
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryImage(&queryParam)
	if err != nil {
		t.Errorf("QueryImage error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No deleted Image found to test Recover")
		return
	}

	recoverParam := param.RecoverImageParam{
		BaseParam: param.BaseParam{},
		Params: param.RecoverImageParamDetail{
			ImageUuid: list[0].UUID,
		},
	}

	result, err := accountLoginCli.RecoverImage(list[0].UUID, recoverParam)
	if err != nil {
		t.Errorf("RecoverImage error: %v", err)
		return
	}
	golog.Infof("RecoverImage succeeded for UUID: %s", result.UUID)
}

func TestCloneImage(t *testing.T) {
	// Note: CloneImage API may not be available in all ZStack versions
	t.Skip("CloneImage API endpoint /v1/image/clone not available in current ZStack version")

	// Query existing images to get a valid UUID to clone
	queryParam := param.NewQueryParam()
	queryParam.AddQ("status=Ready")
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryImage(&queryParam)
	if err != nil {
		t.Errorf("QueryImage error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ready Image found to test Clone")
		return
	}

	cloneParam := param.CloneImageParam{
		BaseParam: param.BaseParam{},
		Params: param.CloneImageParamDetail{
			ImageUuid: list[0].UUID,
		},
	}

	result, err := accountLoginCli.CloneImage(cloneParam)
	if err != nil {
		t.Errorf("CloneImage error: %v", err)
		return
	}
	golog.Infof("CloneImage succeeded, new image UUID: %s", result.UUID)
}
*/
