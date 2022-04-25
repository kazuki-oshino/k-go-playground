package secretstring

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSecretString(t *testing.T) {
	c := NewConfidentialInfo(123, CreditCardCode(1234567890123456))

	act := fmt.Sprint(c.CreditCardCode)
	expect := "xxxx-xxxx-xxxx-xxxx"
	fmt.Println(c)
	fmt.Printf("%+v \n", c)
	fmt.Printf("%#v \n", c)
	if expect != act {
		t.Errorf("expect: %v, actually: %v", expect, act)
	}

	bytes, _ := json.Marshal(c)
	act = fmt.Sprint(string(bytes))
	expect = "{\"Id\":123,\"CreditCardCode\":1234567890123456}"
	fmt.Println(act)
	if expect != act {
		t.Errorf("expect: %v, actually: %v", expect, act)
	}
}
