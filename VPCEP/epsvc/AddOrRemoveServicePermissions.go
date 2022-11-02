package main

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	vpcep "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpcep/v1"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpcep/v1/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpcep/v1/region"
)

func main() {
	ak := "<YOUR AK>"
	sk := "<YOUR SK>"

	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := vpcep.NewVpcepClient(
		vpcep.VpcepClientBuilder().
			WithRegion(region.ValueOf("cn-north-4")).
			WithCredential(auth).
			Build())

	request := &model.AddOrRemoveServicePermissionsRequest{}
	request.VpcEndpointServiceId = "5fdb39b0-e58c-4168-8275-81a2b317104f"
	var listPermissionsbody = []string{
		"iam:domain::d39b170abcfd4983979dedc340ba6a25",
	}
	request.Body = &model.AddOrRemoveServicePermissionsRequestBody{
		Action:      model.GetAddOrRemoveServicePermissionsRequestBodyActionEnum().ADD,
		Permissions: listPermissionsbody,
	}
	response, err := client.AddOrRemoveServicePermissions(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
