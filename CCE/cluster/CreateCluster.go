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

	request := &model.CreateClusterRequest{}
	containerNetworkSpec := &model.ContainerNetwork{
		Mode: model.GetContainerNetworkModeEnum().VPC_ROUTER,
	}
	hostNetworkSpec := &model.HostNetwork{
		Vpc:    COMMON.VpcId,
		Subnet: COMMON.SubnetId,
	}
	categorySpec := model.GetClusterSpecCategoryEnum().CCE
	typeSpec := model.GetClusterSpecTypeEnum().VIRTUAL_MACHINE
	ipv6enableSpec := false
	specbody := &model.ClusterSpec{
		Category:         &categorySpec,
		Type:             &typeSpec,
		Flavor:           COMMON.CceFlavor,
		Ipv6enable:       &ipv6enableSpec,
		HostNetwork:      hostNetworkSpec,
		ContainerNetwork: containerNetworkSpec,
	}
	metadatabody := &model.ClusterMetadata{
		Name: COMMON.CceClusterName,
	}
	request.Body = &model.Cluster{
		Spec:       specbody,
		Metadata:   metadatabody,
		ApiVersion: "v3",
		Kind:       "cluster",
	}

	/* 创建集群 */
	response, err := client.CreateCluster(request)
	if err == nil {
		fmt.Printf("Start to create cluster...\n")
	} else {
		fmt.Printf("Create cluster failed...\n")
		fmt.Printf("Error: %+v\n", err)
		return
	}

	if response != nil && response.Status != nil && response.Status.JobID != nil {
		waitClusterReady(client, *response.Status.JobID)
	} else {
		fmt.Printf("Get cluster state failed...\n")
	}
}

func waitClusterReady(c *cce.CceClient, jobID string) {
	req := &model.ShowJobRequest{
		JobId: jobID,
	}
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	timeout := time.After(30 * time.Minute)
	for {
		select {
		case <-timeout:
			fmt.Printf("Wait cluster state timed out")
			return
		case <-ticker.C:
			resp, err := c.ShowJob(req)
			if err != nil {
				fmt.Printf("Show create cluster job failed...\n")
				fmt.Printf("Error: %+v\n", err)
				return
			}
			if resp != nil && resp.Status != nil && resp.Status.Phase != nil {
				fmt.Printf("Wait cluster ready, job is %s...\n", *resp.Status.Phase)
				if *resp.Status.Phase == "Success" {
					fmt.Printf("Create cluster successfully!")
					return
				} else if *resp.Status.Phase == "Failed" {
					fmt.Printf("Create cluster failed!")
					return
				}
			}
		}
	}
}
