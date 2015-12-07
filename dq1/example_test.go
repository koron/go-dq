package dq1_test

import (
	"fmt"

	"github.com/koron/go-dq/dq1"
)

func ExampleSaveData_Password() {
	d := &dq1.SaveData{
		Name: "とんぬら",
	}
	p, err := d.Password(0)
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
	// Output: れぎざぶい かころぐじでぶ いかこせつ せねふ
}
