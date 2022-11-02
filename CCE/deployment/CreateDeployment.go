package main

import (
	"HuaweiCloudService/CCE/client"
	"context"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

// todo zhangbo 先创建configmap
func main() {
	cli := client.GetK8sClient()
	var replicas int32 = 1
	var defaultMode int32 = 420
	var privileged = true
	deployment := &appv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "haproxy",
			Namespace: `roboartisan`,
		},
		Spec: appv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":     "haproxy",
					"version": "v1",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":     "haproxy",
						"version": "v1",
					},
				},
				Spec: corev1.PodSpec{
					Volumes: []corev1.Volume{
						{
							Name: "haproxycfg",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: `haproxycfg`,
									},
									DefaultMode: &defaultMode,
								},
							},
						},
					},
					Containers: []corev1.Container{
						{
							Name:  `haproxy`,
							Image: "haproxy:latest",
							Resources: corev1.ResourceRequirements{
								Limits: map[corev1.ResourceName]resource.Quantity{
									corev1.ResourceCPU:    resource.MustParse("4"),
									corev1.ResourceMemory: resource.MustParse("4Gi"),
								},
								Requests: map[corev1.ResourceName]resource.Quantity{
									corev1.ResourceCPU:    resource.MustParse("250m"),
									corev1.ResourceMemory: resource.MustParse("512Mi"),
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "haproxycfg",
									ReadOnly:  true,
									MountPath: "/usr/local/etc/haproxy/",
								},
							},
							SecurityContext: &corev1.SecurityContext{
								Privileged: &privileged,
							},
						},
					},
					HostNetwork: true,
				},
			},
		},
	}

	ret, err := cli.AppsV1().Deployments(`roboartisan`).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		klog.Errorln(err)
	} else {
		klog.Infof("Create deployment %s successful", ret.Name)
	}
}
