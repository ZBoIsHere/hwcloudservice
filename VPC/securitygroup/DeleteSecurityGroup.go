package main

import (
	"HuaweiCloudService/COMMON"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/region"
)

func main() {
	ak := COMMON.HW_AK
	sk := COMMON.HW_SK

	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := vpc.NewVpcClient(
		vpc.VpcClientBuilder().
			WithRegion(region.ValueOf(COMMON.DEBUG_REGION)).
			WithCredential(auth).
			Build())

	request := &model.DeleteSecurityGroupRequest{}
	request.SecurityGroupId = COMMON.SecurityGroupId
	response, err := client.DeleteSecurityGroup(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
