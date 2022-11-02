package main

import (
	"HuaweiCloudService/COMMON"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	cce "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cce/v3"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cce/v3/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cce/v3/region"
	"time"
)

func main() {
	ak := COMMON.HW_AK
	sk := COMMON.HW_SK

	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := cce.NewCceClient(
		cce.CceClientBuilder().
			WithRegion(region.ValueOf(COMMON.DEBUG_REGION)).
			WithCredential(auth).
			Build())

	request := &model.CreateNodeRequest{}
	request.ClusterId = COMMON.CceClusterId
	nameRuntime := model.GetRuntimeNameEnum().DOCKER
	runtimeSpec := &model.Runtime{
		Name: &nameRuntime,
	}
	subnetIdPrimaryNic := COMMON.SubnetId
	primaryNicNodeNicSpec := &model.NicSpec{
		SubnetId: &subnetIdPrimaryNic,
	}
	nodeNicSpecSpec := &model.NodeNicSpec{
		PrimaryNic: primaryNicNodeNicSpec,
	}
	var listVirtualSpacesStorageGroups = []model.VirtualSpace{
		{
			Name: "runtime",
			Size: "90%",
		},
		{
			Name: "kubernetes",
			Size: "10%",
		},
	}
	var listSelectorNamesStorageGroups = []string{
		"cceUse",
	}
	cceManagedStorageGroups := true
	var listStorageGroupsStorage = []model.StorageGroups{
		{
			Name:          "vgpaas",
			CceManaged:    &cceManagedStorageGroups,
			SelectorNames: listSelectorNamesStorageGroups,
			VirtualSpaces: listVirtualSpacesStorageGroups,
		},
	}
	sizeMatchLabels := "100"
	volumeTypeMatchLabels := "ESSD"
	countMatchLabels := "1"
	matchLabelsStorageSelectors := &model.StorageSelectorsMatchLabels{
		Size:       &sizeMatchLabels,
		VolumeType: &volumeTypeMatchLabels,
		Count:      &countMatchLabels,
	}
	var listStorageSelectorsStorage = []model.StorageSelectors{
		{
			Name:        "cceUse",
			StorageType: "evs",
			MatchLabels: matchLabelsStorageSelectors,
		},
	}
	storageSpec := &model.Storage{
		StorageSelectors: listStorageSelectorsStorage,
		StorageGroups:    listStorageGroupsStorage,
	}
	hwPassthroughDataVolumes := true
	var listDataVolumesSpec = []model.Volume{
		{
			Size:          int32(100),
			Volumetype:    "ESSD",
			Hwpassthrough: &hwPassthroughDataVolumes,
		},
	}
	hwPassthroughRootVolume := true
	rootVolumeSpec := &model.Volume{
		Size:          int32(50),
		Volumetype:    "ESSD",
		Hwpassthrough: &hwPassthroughRootVolume,
	}
	usernameUserPassword := "root"
	userPasswordLogin := &model.UserPassword{
		Username: &usernameUserPassword,
		Password: COMMON.Password,
	}
	loginSpec := &model.Login{
		UserPassword: userPasswordLogin,
	}
	osSpec := "EulerOS 2.5"
	countSpec := int32(1)
	billingModeSpec := int32(0)
	specbody := &model.NodeSpec{
		Flavor:      "c7.xlarge.4",
		Az:          "cn-north-4a",
		Os:          &osSpec,
		Login:       loginSpec,
		RootVolume:  rootVolumeSpec,
		DataVolumes: listDataVolumesSpec,
		Storage:     storageSpec,
		NodeNicSpec: nodeNicSpecSpec,
		Count:       &countSpec,
		BillingMode: &billingModeSpec,
		Runtime:     runtimeSpec,
	}
	nameMetadata := "cce-node-zhangbotest"
	metadatabody := &model.NodeMetadata{
		Name: &nameMetadata,
	}
	request.Body = &model.NodeCreateRequest{
		Spec:       specbody,
		Metadata:   metadatabody,
		ApiVersion: "v3",
		Kind:       "Node",
	}
	response, err := client.CreateNode(request)
	if err == nil {
		fmt.Printf("Start to create node...\n")
	} else {
		fmt.Printf("Create node failed...\n")
		fmt.Printf("Error: %+v\n", err)
		return
	}

	if response != nil && response.Status != nil && response.Status.JobID != nil {
		waitNodeReady(client, *response.Status.JobID)
	} else {
		fmt.Printf("Get node state failed...\n")
	}
}

func waitNodeReady(c *cce.CceClient, jobID string) {
	req := &model.ShowJobRequest{
		JobId: jobID,
	}
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	timeout := time.After(30 * time.Minute)
	for {
		select {
		case <-timeout:
			fmt.Printf("Wait node state timed out.")
			return
		case <-ticker.C:
			resp, err := c.ShowJob(req)
			if err != nil {
				fmt.Printf("Show create node job failed...\n")
				fmt.Printf("Error: %+v\n", err)
				return
			}
			if resp != nil && resp.Status != nil && resp.Status.Phase != nil {
				fmt.Printf("Wait node ready, job is %s\n", *resp.Status.Phase)
				if *resp.Status.Phase == "Success" {
					fmt.Printf("Create node successfully!")
					return
				} else if *resp.Status.Phase == "Failed" {
					fmt.Printf("Create node failed!")
					return
				}
			}
		}
	}
}
