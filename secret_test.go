package secretapi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_testing(t *testing.T) {
	type CreateSecretRequestBody struct {
		Data  map[string]string `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
		Name  *string           `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
		Scope *string           `form:"scope,omitempty" json:"scope,omitempty" xml:"scope,omitempty"`
	}
	marshal, err := json.Marshal(CreateSecretRequestBody{
		Data:  map[string]string{"key1": "value1"},
		Name:  stringp("my-secret"),
		Scope: stringp("my-scope"),
	})
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}

func stringp(s string) *string {
	return &s
}
