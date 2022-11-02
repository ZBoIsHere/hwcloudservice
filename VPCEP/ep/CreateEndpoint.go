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

	request := &model.CreateEndpointRequest{}
	portIpCreateEndpointRequestBody := "192.168.1.144"
	enableDnsCreateEndpointRequestBody := false
	subnetIdCreateEndpointRequestBody := "e6cb1393-78e9-4431-b2a3-9c6af85bc6bf"
	request.Body = &model.CreateEndpointRequestBody{
		PortIp:            &portIpCreateEndpointRequestBody,
		EnableDns:         &enableDnsCreateEndpointRequestBody,
		VpcId:             "111",
		EndpointServiceId: "5fdb39b0-e58c-4168-8275-81a2b317104f",
		SubnetId:          &subnetIdCreateEndpointRequestBody,
	}
	response, err := client.CreateEndpoint(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
