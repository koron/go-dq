# go-dq

go-dq はFC版ドラゴンクエストの、復活の呪文を生成するライブラリです。

現在は dq1/ しか実装していません。

使い方の例:

```go
import (
	"fmt"

	"github.com/koron/go-dq/dq1"
)

func main() {
	d := &dq1.SaveData{
		Name: "とんぬら",
	}
	p, err := dq1.Encode(d, 0)
	if err != nil {
		panic(err)
	}
	// "れぎざぶい かころぐじでぶ いかこせつ せねふ"
	fmt.Println(dq1.Format(p))
}
```

## References

以下の資料を参考にしました。

*   [「復活の呪文」資料室](http://www.imasy.or.jp/~yotti/dq-passwd.html)
