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
	cm := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "haproxycfg",
			Namespace: "roboartisan",
		},
		Data: map[string]string{
			"haproxy.cfg": `    
	global
        maxconn 102400

    defaults
        mode            tcp
        log             global
        option          dontlognull
        option http-server-close
        option          redispatch
        retries         10
        timeout http-request 1m
        timeout queue   1m
        timeout connect 10s
        timeout client  1m
        timeout server  1m
        timeout http-keep-alive 1m
        timeout check   10s
        maxconn         102400

    frontend    echo_frontend
        bind        *:13307
        mode        tcp
        default_backend echo_server

    backend     echo_server
        mode tcp
        balance roundrobin
        server echo-01 192.168.1.108:8635
`,
		},
	}

	ret, err := cli.CoreV1().ConfigMaps(`roboartisan`).Create(context.TODO(), cm, metav1.CreateOptions{})
	if err != nil {
		klog.Errorln(err)
	} else {
		klog.Infof("Create configmap %s successful", ret.Name)
	}

}
