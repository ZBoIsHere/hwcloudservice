package main

import (
	"fmt"
	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
)

func main() {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "https://120.46.160.30:8788",
		Username: "rwusr",
		Password: "ecs@245680!",
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	q := client.NewQuery("select * from cpu", "db0", "ns")
	if response, err := c.Query(q); err == nil && response.Error() == nil {
		fmt.Println("the result is: ", response.Results)
	}
}
