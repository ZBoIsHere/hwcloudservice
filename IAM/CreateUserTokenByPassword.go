package main

import (
	"HuaweiCloudService/COMMON"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
	iam "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/region"
)

func main() {
	ak := COMMON.HW_AK
	sk := COMMON.HW_SK

	auth := global.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := iam.NewIamClient(
		iam.IamClientBuilder().
			WithRegion(region.ValueOf(COMMON.DEBUG_REGION)).
			WithCredential(auth).
			Build())

	request := &model.KeystoneCreateUserTokenByPasswordRequest{}
	idProject := COMMON.ProjectID
	projectScope := &model.AuthScopeProject{
		Id: &idProject,
	}
	scopeAuth := &model.AuthScope{
		Project: projectScope,
	}
	domainUser := &model.PwdPasswordUserDomain{
		Name: "hwstaff_pub_CTOTI",
	}
	userPassword := &model.PwdPasswordUser{
		Domain:   domainUser,
		Name:     "CTO_TI_FBSYJG",
		Password: `ecil@2022Edge`,
	}
	passwordIdentity := &model.PwdPassword{
		User: userPassword,
	}
	var listMethodsIdentity = []model.PwdIdentityMethods{
		model.GetPwdIdentityMethodsEnum().PASSWORD,
	}
	identityAuth := &model.PwdIdentity{
		Methods:  listMethodsIdentity,
		Password: passwordIdentity,
	}
	authbody := &model.PwdAuth{
		Identity: identityAuth,
		Scope:    scopeAuth,
	}
	request.Body = &model.KeystoneCreateUserTokenByPasswordRequestBody{
		Auth: authbody,
	}
	response, err := client.KeystoneCreateUserTokenByPassword(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
