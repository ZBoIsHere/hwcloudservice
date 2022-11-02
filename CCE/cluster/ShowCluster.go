package main

import (
	"HuaweiCloudService/COMMON"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	cce "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cce/v3"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cce/v3/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cce/v3/region"
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

	request := &model.ShowClusterRequest{}
	request.ClusterId = COMMON.CceClusterId
	response, err := client.ShowCluster(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
