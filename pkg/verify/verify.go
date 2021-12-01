package verify

import (
	"fmt"

	"github.com/Berger7/yunxiao-go-client/config"
)

func RunVerify() {
	fmt.Println("⚠️  Configurations just support with file.")
	CheckStringParamIsNotEmpty("accesskey", config.GetAK())
	CheckStringParamIsNotEmpty("secretkey", config.GetSK())
	CheckStringParamIsNotEmpty("endpoint", config.GetEndpoint())
}

func CheckStringParamIsNotEmpty(key, value string) {
	if value == "" {
		panic(fmt.Sprintf("💔 [%s] is empty", key))
	}
	fmt.Printf("🍺 [%s] check ...... \033[32mready\033[0m\n", key)
}
