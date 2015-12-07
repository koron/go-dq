# go-dq

go-dq はFC版ドラゴンクエストの、復活の呪文を生成するライブラリです。

現在は dq1/ しか実装していません。

使い方の例:

```go
package dq1_test

import (
	"fmt"

	"github.com/koron/go-dq/dq1"
)

func ExamplePassword() {
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
```

## References

以下の資料を参考にしました。

*   [「復活の呪文」資料室](http://www.imasy.or.jp/~yotti/dq-passwd.html)
