package main

import (
	"HuaweiCloudService/CCE/client"
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/klog/v2"
)

func main() {
	cli := client.GetK8sClient()

	lbSvc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "haproxysvc",
			Namespace: "roboartisan",
			Labels: map[string]string{
				"app":     "haproxy",
				"version": "v1",
			},
			Annotations: map[string]string{
				"kubernetes.io/elb.autocreate": `{
            		"type": "inner",
            		"name": "inner-elb"
				}`,
				`kubernetes.io/elb.class`: `union`,
			},
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:     "cce-service-0",
					Protocol: corev1.ProtocolTCP,
					Port:     33333,
					TargetPort: intstr.IntOrString{
						IntVal: 13307,
					},
				},
			},
			Selector: map[string]string{
				"app":     "haproxy",
				"version": "v1",
			},
			Type: "LoadBalancer",
		},
	}

	ret, err := cli.CoreV1().Services(`roboartisan`).Create(context.TODO(), lbSvc, metav1.CreateOptions{})
	if err != nil {
		klog.Errorln(err)
	} else {
		klog.Infof("Create loadbalancer service %s successful", ret.Name)
	}

}
