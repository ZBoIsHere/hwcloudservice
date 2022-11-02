package main

import (
	"HuaweiCloudService/COMMON"
	dns "HuaweiCloudService/DNS"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/region"
)

func main() {
	dns.ListNameserver()

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

	request := &model.CreateSubnetRequest{}
	var listDnsListSubnet = []string{}
	var primaryDnsSubnet, secondaryDnsSubnet string
	for _, v := range COMMON.SubnetDnsList {
		listDnsListSubnet = append(listDnsListSubnet, *v.Address)
		if *v.Priority == 1 {
			primaryDnsSubnet = *v.Address
		} else if *v.Priority == 2 {
			secondaryDnsSubnet = *v.Address
		}
	}

	descriptionSubnet := COMMON.SubnetName + `description`
	ipv6EnableSubnet := false
	dhcpEnableSubnet := true

	subnetbody := &model.CreateSubnetOption{
		Name:         COMMON.SubnetName,
		Description:  &descriptionSubnet,
		Cidr:         COMMON.SubnetCIDR,
		VpcId:        COMMON.VpcId,
		GatewayIp:    COMMON.SubnetGWIP,
		Ipv6Enable:   &ipv6EnableSubnet,
		DhcpEnable:   &dhcpEnableSubnet,
		PrimaryDns:   &primaryDnsSubnet,
		SecondaryDns: &secondaryDnsSubnet,
		DnsList:      &listDnsListSubnet,
	}
	request.Body = &model.CreateSubnetRequestBody{
		Subnet: subnetbody,
	}
	response, err := client.CreateSubnet(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
