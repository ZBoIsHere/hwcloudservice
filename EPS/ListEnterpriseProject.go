package main

import (
	"HuaweiCloudService/COMMON"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
	eps "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eps/v1"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eps/v1/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eps/v1/region"
)

func main() {
	ak := COMMON.HW_AK
	sk := COMMON.HW_SK

	auth := global.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := eps.NewEpsClient(
		eps.EpsClientBuilder().
			WithRegion(region.ValueOf("cn-north-4")).
			WithCredential(auth).
			Build())

	request := &model.ListEnterpriseProjectRequest{}
	response, err := client.ListEnterpriseProject(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
