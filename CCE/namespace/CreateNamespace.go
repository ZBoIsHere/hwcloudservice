package main

import (
	"HuaweiCloudService/CCE/client"
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

func main() {
	cli := client.GetK8sClient()
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "roboartisan",
		},
	}
	ret, err := cli.CoreV1().Namespaces().Create(context.TODO(), ns, metav1.CreateOptions{})
	if err != nil {
		klog.Errorln(err)
	} else {
		klog.Infof("Create namespace %s successful", ret.Name)
	}
}
