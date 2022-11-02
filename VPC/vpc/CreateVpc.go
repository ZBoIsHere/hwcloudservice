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

	request := &model.CreateVpcRequest{}
	cidrVpc := COMMON.VpcCIDR
	nameVpc := COMMON.VpcName
	descriptionVpc := COMMON.VpcDescription
	enterpriseProjectIdVpc := COMMON.EnterpriseProjectIdVpc
	vpcbody := &model.CreateVpcOption{
		Cidr:                &cidrVpc,
		Name:                &nameVpc,
		Description:         &descriptionVpc,
		EnterpriseProjectId: &enterpriseProjectIdVpc,
	}
	request.Body = &model.CreateVpcRequestBody{
		Vpc: vpcbody,
	}
	response, err := client.CreateVpc(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
		COMMON.VpcId = response.Vpc.Id
	} else {
		fmt.Println(err)
	}
}
