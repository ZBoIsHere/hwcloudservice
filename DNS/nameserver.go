package dns

import (
	"HuaweiCloudService/COMMON"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	dns "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dns/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dns/v2/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dns/v2/region"
)

func ListNameserver() {
	ak := COMMON.HW_AK
	sk := COMMON.HW_SK

	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := dns.NewDnsClient(
		dns.DnsClientBuilder().
			WithRegion(region.ValueOf(COMMON.DEBUG_REGION)).
			WithCredential(auth).
			Build())

	request := &model.ListNameServersRequest{}
	typeRequest := COMMON.DnsRequestType
	request.Type = &typeRequest
	regionRequest := COMMON.DEBUG_REGION
	request.Region = &regionRequest
	response, err := client.ListNameServers(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
		for _, v := range *response.Nameservers {
			if *v.Region == COMMON.DEBUG_REGION {
				COMMON.SubnetDnsList = *v.NsRecords
			}
		}
	} else {
		fmt.Println(err)
	}
}
