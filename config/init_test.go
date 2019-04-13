package config

import (
	"blog-api/app/utils"
	"fmt"
	"testing"
)

func TestInitConfig(t *testing.T) {
	c := GetConfig("../app.yaml")
	fmt.Println(utils.JsonMarshalIndent(c, "	", "	"))
}

func BenchmarkGetConfig(b *testing.B) {
	b.N = 2000000
	for i := 0; i < b.N; i++ {
		GetConfig("../app.yaml")
	}
}
func BenchmarkGetConfig2(b *testing.B) {
	b.N = 200
	for i := 0; i < b.N; i++ {
		GetConfigNoSingleInstance("../app.yaml")
	}
}
