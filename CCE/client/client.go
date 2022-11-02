package client

import (
	"HuaweiCloudService/COMMON"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
)

func GetK8sClient() *kubernetes.Clientset {
	cs, err := kubernetes.NewForConfig(&rest.Config{
		Host:        fmt.Sprintf("https://%s.cce.%s.myhuaweicloud.com", COMMON.CceClusterId, COMMON.DEBUG_REGION),
		BearerToken: COMMON.Token,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	})
	if err != nil {
		klog.Errorln(err)
	}
	return cs
}
