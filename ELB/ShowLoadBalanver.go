package main

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	elb "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v2/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v2/region"
)

func main() {
	ak := "<YOUR AK>"
	sk := "<YOUR SK>"

	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := elb.NewElbClient(
		elb.ElbClientBuilder().
			WithRegion(region.ValueOf("cn-north-4")).
			WithCredential(auth).
			Build())

	request := &model.ShowLoadbalancerRequest{}
	request.LoadbalancerId = "2c7c748d-05f6-4507-a4bc-2eb879e70cd5"
	response, err := client.ShowLoadbalancer(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
