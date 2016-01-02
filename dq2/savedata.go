package dq2

import "fmt"

// SaveData is DQ2のセーブデータ
type SaveData struct {

	// CheckPoint is 最後に復活の呪文を聞いた場所.  Choosen from CheckPoints.
	CheckPoint string

	// Name is ローレシアの王子の名前. Max length is 4. Composed from chars in
	// NameKana.
	Name string

	// Gold is 所持ゴールド. Max is 65535.
	Gold int

	// Flags is 各イベントの進行フラグ.
	Flags struct {

		// SeaCave shows 海の洞窟を開いたか
		SeaCave bool

		// TepaGate shows テパの水門を開いたか
		TepaGate bool

		// Weaving shows みずのはごろもを織ってもらっているか
		Weaving bool

		// Ship shows ルプガナでのイベント進行度 (0-2)
		//   0:なし
		//   1:女の子を助けた
		//   2:助けて船を貰った
		Ship int

		// Prince shows サマルトリアの王子の進行度 (0-3)
		//   0:なし
		//   1:サマルトリア王に会った
		//   2:勇者の泉に探しに行った
		//   3:見つけた
		Prince int
	}

	// Crest shows 紋章の取得状況
	Crest struct {
		Sun   bool
		Star  bool
		Moon  bool
		Water bool
		Life  bool
	}

	// Lorasia is ローレシアの王子の状態.
	Lorasia Char

	// Samantoria is サマルトリアの王子の状態. 仲間になってない場合nil
	Samantoria *Char

	// Moonbrooke is ムーンブルクの王女の状態. 仲間になってない場合nil
	Moonbrooke *Char
}

func (d *SaveData) checkPoint() (int, error) {
	n := findString(d.CheckPoint, CheckPoints)
	if n < 0 {
		return 0, fmt.Errorf("invalid CheckPoint %q", d.CheckPoint)
	}
	return n, nil
}

func (d *SaveData) gold() (int, error) {
	if d.Gold < 0 || d.Gold > 65535 {
		return 0, fmt.Errorf("invalid Gold %d", d.Gold)
	}
	return d.Gold, nil
}

func (d *SaveData) name() ([]byte, error) {
	b := []byte{62, 62, 62, 62}
	i := 0
	for _, r := range d.Name {
		if i >= len(b) {
			return nil, fmt.Errorf("too long Name %q", d.Name)
		}
		n := findRune(r, NameKana)
		if n < 0 {
			return nil, fmt.Errorf("invalid character in Name %q", d.Name)
		}
		b[i] = byte(n)
		i++
	}
	return b, nil
}

func (d *SaveData) encodePassword(cryptKey int) ([]byte, error) {
	b := &Buffer{}

	cp, err := d.checkPoint()
	if err != nil {
		return nil, err
	}
	name, err := d.name()
	if err != nil {
		return nil, err
	}
	gold, err := d.gold()
	if err != nil {
		return nil, err
	}

	// check point
	b.WriteByte(byte(cp))
	// name and gold
	b.WriteByte(byte((name[2] << 2) | ((name[1] & 0x30) >> 4)))
	b.WriteByte(byte(gold >> 8))
	b.WriteByte(byte(((name[1] & 0x06) << 5) | name[0]))
	b.WriteByte(byte(gold))
	b.WriteByte(byte(((name[1] & 0x01) << 7) | (name[3] << 1) |
		((name[1] & 0x08) >> 3)))
	// crypt key and flags
	b.WriteByte(byte(((cryptKey & 0x01) << 7) |
		btoi(d.Flags.SeaCave, 0x40) |
		btoi(d.Flags.TepaGate, 0x20) |
		btoi(d.Flags.Weaving, 0x10) |
		((d.Flags.Ship & 0x03) << 2) |
		(d.Flags.Prince & 0x03)))
	b.WriteByte(byte(((cryptKey & 0x0e) << 4) |
		btoi(d.Crest.Life, 0x10) |
		btoi(d.Crest.Water, 0x08) |
		btoi(d.Crest.Moon, 0x04) |
		btoi(d.Crest.Star, 0x02) |
		btoi(d.Crest.Sun, 0x01)))
	b.WriteByte(0)

	// Lorasia
	b.WriteBits(d.Lorasia.Exp, 16)
	b.WriteBits(d.Lorasia.Exp>>16, 4)
	d.Lorasia.WriteItems(b)

	// Samantoria
	if d.Samantoria != nil {
		b.WriteBits(1, 1)
		b.WriteBits(d.Samantoria.Exp, 16)
		b.WriteBits(d.Samantoria.Exp>>16, 4)
		d.Samantoria.WriteItems(b)
		// Moonbrooke
		if d.Moonbrooke != nil {
			b.WriteBits(1, 1)
			b.WriteBits(d.Moonbrooke.Exp, 16)
			b.WriteBits(d.Moonbrooke.Exp>>16, 4)
			d.Moonbrooke.WriteItems(b)
		} else {
			b.WriteBits(0, 1)
		}
	} else {
		b.WriteBits(0, 1)
	}
	err = b.Flush()
	if err != nil {
		return nil, err
	}

	// TODO:

	return nil, nil
}

// Char is キャラクターの状態
type Char struct {

	// Exp is 経験値. Max is 1000000.
	Exp int

	// Items is 所持アイテム. Max is 8, choosen from Items.
	Items []string
}

func (c *Char) items() []int {
	a := make([]int, 0, len(c.Items))
	for _, s := range c.Items {
		n := findString(s, Items)
		if n < 0 {
			continue
		}
		a = append(a, n)
	}
	return a
}

func (c *Char) WriteItems(b *Buffer) {
	a := c.items()
	b.WriteBits(len(a), 4)
	for _, n := range a {
		b.WriteBits(n, 7)
	}
}
