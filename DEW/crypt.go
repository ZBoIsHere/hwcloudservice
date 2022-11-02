package main

import (
	"encoding/base64"
	"fmt"

	"github.com/GehirnInc/crypt"
	_ "github.com/GehirnInc/crypt/sha256_crypt"
)

func main() {
	crypt := crypt.SHA256.New()
	ret, _ := crypt.Generate([]byte("ecs@245680!"), []byte("$5$salt"))
	fmt.Println(ret)

	err := crypt.Verify(ret, []byte("secret"))
	fmt.Println(err)

	password := base64.StdEncoding.EncodeToString([]byte(ret))
	fmt.Println(password)

	// Output:
	// $5$salt$kpa26zwgX83BPSR8d7w93OIXbFt/d3UOTZaAu5vsTM6
	// <nil>
}
