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

	request := &model.CreateEndpointServiceRequest{}
	clientPortPorts := int32(55555)
	serverPortPorts := int32(44444)
	protocolPorts := model.GetPortListProtocolEnum().TCP
	var listPortsbody = []model.PortList{
		{
			ClientPort: &clientPortPorts,
			ServerPort: &serverPortPorts,
			Protocol:   &protocolPorts,
		},
	}
	approvalEnabledCreateEndpointServiceRequestBody := false
	request.Body = &model.CreateEndpointServiceRequestBody{
		Ports:           listPortsbody,
		ServerType:      model.GetCreateEndpointServiceRequestBodyServerTypeEnum().LB,
		ApprovalEnabled: &approvalEnabledCreateEndpointServiceRequestBody,
		VpcId:           "416c8a88-6064-4510-81cb-b8450b108936",
		// todo zhangbo 获取lb详情中的vip_port_id字段
		PortId: "54e7fdec-254f-45e8-b406-757c520fbd9a",
	}
	response, err := client.CreateEndpointService(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
