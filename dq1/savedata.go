package dq1

import (
	"bytes"
	"fmt"
)

// PassKana is 復活の呪文用の「かな」
var PassKana = []rune{
	'あ', 'い', 'う', 'え', 'お', 'か', 'き', 'く', 'け', 'こ',
	'さ', 'し', 'す', 'せ', 'そ', 'た', 'ち', 'つ', 'て', 'と',
	'な', 'に', 'ぬ', 'ね', 'の', 'は', 'ひ', 'ふ', 'へ', 'ほ',
	'ま', 'み', 'む', 'め', 'も', 'や', 'ゆ', 'よ', 'ら', 'り',
	'る', 'れ', 'ろ', 'わ', 'が', 'ぎ', 'ぐ', 'げ', 'ご', 'ざ',
	'じ', 'ず', 'ぜ', 'ぞ', 'だ', 'ぢ', 'づ', 'で', 'ど', 'ば',
	'び', 'ぶ', 'べ', 'ぼ',
}

// NameKana is 名前用の「かな」
var NameKana = []rune{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'あ', 'い', 'う', 'え', 'お', 'か', 'き', 'く', 'け', 'こ',
	'さ', 'し', 'す', 'せ', 'そ', 'た', 'ち', 'つ', 'て', 'と',
	'な', 'に', 'ぬ', 'ね', 'の', 'は', 'ひ', 'ふ', 'へ', 'ほ',
	'ま', 'み', 'む', 'め', 'も', 'や', 'ゆ', 'よ', 'ら', 'り',
	'る', 'れ', 'ろ', 'わ', 'を', 'ん', 'っ', 'ゃ', 'ゅ', 'ょ',
	'゛', '゜', 'ー', ' ',
}

// Weapons is 武器一覧
var Weapons = []string{
	"", "たけざお", "こんぼう", "どうのつるぎ", "てつのおの",
	"はがねのつるぎ", "ほのおのつるぎ", "ロトのつるぎ",
}

// Armors is 鎧一覧
var Armors = []string{
	"", "ぬののふく", "かわのふく", "くさりかたびら", "てつのよろい",
	"はがねのよろい", "まほうのよろい", "ロトのよろい",
}

// Shields is 盾一覧
var Shields = []string{
	"", "かわのたて", "てつのたて", "みかがみのたて",
}

// Items is アイテム一覧
var Items = []string{
	"", "たいまつ", "せいすい", "キメラのつばさ", "りゅうのうろこ",
	"ようせいのふえ", "せんしのゆびわ", "ロトのしるし", "おうじょのあい",
	"のろいのベルト", "ぎんのたてごと", "しのくびかざり", "たいようのいし",
	"あまぐものつえ", "にじのしずく", "(ありえないアイテム)",
}

var crcTable = []int{
	0x7e, 0xbf, 0xcf, 0xf7, 0x6b, 0xa5, 0xc2, 0x61,
	0xa0, 0x00, 0x68, 0xb4, 0x5a, 0xad, 0xc6, 0x63,
	0xa1, 0x40, 0xa0, 0x50, 0xa8, 0xd4, 0xea, 0x75,
	0xaa, 0xd5, 0x7a, 0x3d, 0x8e, 0x47, 0xb3, 0x49,
	0x34, 0x1a, 0x0d, 0x96, 0x4b, 0x35, 0x8a, 0x45,
	0xb2, 0xd9, 0xfc, 0xfe, 0xff, 0xef, 0x67, 0x23,
	0x81, 0xd0, 0x68, 0xb4, 0xda, 0x6d, 0xa6, 0xd3,
	0xf9, 0xec, 0xf6, 0x7b, 0xad, 0xc6, 0xe3, 0x61,
	0x20, 0x10, 0x08, 0x84, 0x42, 0xa1, 0x40, 0xa0,
	0x00, 0x68, 0xb4, 0x5a, 0x00, 0x06, 0x83, 0x51,
	0x38, 0x9c, 0x4e, 0xa7, 0xc3, 0xf1, 0x68, 0xb4,
	0x5a, 0xad, 0xc6, 0x63, 0xa1, 0xc0, 0x60, 0x30,
	0x98, 0xcc, 0xe6, 0x73, 0xa9, 0xc4, 0x62, 0x00,
	0x88, 0xc4, 0x62, 0x31, 0x08, 0x84, 0x42, 0x21,
}

// SaveData is DQ1のセーブデータ
type SaveData struct {

	// Name is 名前. Max length 4.  Composed from chars in NameKana
	Name string

	// Gold is 所持金. Max value is 65535
	Gold int

	// Exp is 経験値. Max value is 65535
	Exp int

	// KeyNum is 鍵の所持数. Max value is 7
	KeyNum int

	// HerbNum is 薬草の所持数. Max value is 7
	HerbNum int

	// Weapon is 武器. Should be choosen from Weapons
	Weapon string

	// Armors is 鎧. Should be choosen from Armors
	Armor string

	// Shield is 盾. Shield be choosen from Shields
	Shield string

	// Items is 所持アイテム.  Shol be choosen from Items
	Items []string

	// Uroko indicates 「りゅうのうろこ」を身につけたことがある
	Uroko bool

	// Oujyo indicates 王女を捕らえたドラゴンを倒した
	Oujyo bool

	// Gorem indicates ゴーレムを倒した
	Gorem bool

	// Ring indicates 「せんしのゆびわ」を身に着けている
	Ring bool

	// Death indicates 「しのくびかざり」を発見済み
	Death bool
}

func (d *SaveData) name() ([4]int, error) {
	ret := [4]int{63, 63, 63, 63}
	i := 0
	for _, r := range d.Name {
		if i >= len(ret) {
			return ret, fmt.Errorf("long name")
		}
		n := findRune(r, NameKana)
		if n < 0 {
			return ret, fmt.Errorf("invalid kana:%c in Name at #%d", r, i)
		}
		ret[i] = n
		i++
	}
	return ret, nil
}

func (d *SaveData) weapon() (int, error) {
	n := findString(d.Weapon, Weapons)
	if n < 0 {
		return -1, fmt.Errorf("invalid Weapon:%q", d.Weapon)
	}
	return n, nil
}

func (d *SaveData) armor() (int, error) {
	n := findString(d.Armor, Armors)
	if n < 0 {
		return -1, fmt.Errorf("invalid Armor:%q", d.Armor)
	}
	return n, nil
}

func (d *SaveData) shield() (int, error) {
	n := findString(d.Shield, Shields)
	if n < 0 {
		return -1, fmt.Errorf("invalid Shield:%q", d.Shield)
	}
	return n, nil
}

func (d *SaveData) items() ([8]int, error) {
	var ret [8]int
	if len(d.Items) > 8 {
		return ret, fmt.Errorf("too many items")
	}
	for i, t := range d.Items {
		n := findString(t, Items)
		if n < 0 {
			return ret, fmt.Errorf("invalid item:%q in Items at #%d", t, i)
		}
		ret[i] = n
	}
	return ret, nil
}

// Password generates 復活の呪文 with variation v (0-7).
func (d *SaveData) Password(v int) (string, error) {
	c, err := d.encodePassword(v)
	if err != nil {
		return "", err
	}
	return formatPassword(c), nil
}

// encodePassword generates 復活の呪文(password) from SaveData
func (d *SaveData) encodePassword(v int) ([]rune, error) {
	name, err := d.name()
	if err != nil {
		return nil, err
	}
	items, err := d.items()
	if err != nil {
		return nil, err
	}
	weapon, err := d.weapon()
	if err != nil {
		return nil, err
	}
	armor, err := d.armor()
	if err != nil {
		return nil, err
	}
	shield, err := d.shield()
	if err != nil {
		return nil, err
	}
	if d.Gold < 0 || d.Gold > 65535 {
		return nil, fmt.Errorf("gold overflow: %d", d.Gold)
	}
	if d.Exp < 0 || d.Exp > 65535 {
		return nil, fmt.Errorf("exp overflow: %d", d.Exp)
	}
	if d.KeyNum < 0 || d.KeyNum > 6 {
		return nil, fmt.Errorf("key-num overflow: %d", d.KeyNum)
	}
	if d.HerbNum < 0 || d.HerbNum > 6 {
		return nil, fmt.Errorf("herb-num overflow: %d", d.HerbNum)
	}

	checks := []bool{
		(v & 0x4) != 0,
		(v & 0x2) != 0,
		(v & 0x1) != 0,
	}

	code := []int{
		0,
		d.Exp & 0xff,
		btoi(checks[2], 128) | btoi(d.Death, 64) | name[2],
		(items[3] << 4) | items[2],
		d.Gold & 0xff,
		(name[0] << 2) | btoi(d.Gorem, 2) | btoi(checks[1], 1),
		(items[7] << 4) | items[6],
		btoi(checks[0], 128) | btoi(d.Oujyo, 64) | name[3],
		(weapon << 5) | (armor << 2) | shield,
		(d.Gold >> 8) & 0xff,
		(d.KeyNum << 4) | d.HerbNum,
		(items[5] << 4) | items[4],
		(d.Exp >> 8) & 0xff,
		btoi(d.Uroko, 0x80) | (name[1] << 1) | btoi(d.Ring, 1),
		(items[1] << 4) | items[0],
	}

	// calc crc and store it as last code
	crc := 0
	for i, c := range code[1:] {
		for j := 0; j < 8; j++ {
			if c&0x80 != 0 {
				crc ^= crcTable[i*8+j]
			}
			c <<= 1
		}
	}
	code[0] = crc & 0xff

	// make 6bit twisted sequence.
	data := make([]int, 0, 20)
	for i := 0; i < len(code); i += 3 {
		n := (code[i+2] << 16) | (code[i+1] << 8) | code[i]
		data = append(data, n&0x3f, (n>>6)&0x3f, (n>>12)&0x3f, (n>>18)&0x3f)
	}

	// encrypt and encode
	password := make([]rune, 0, len(data))
	a := 0
	for i, n := range data {
		data[i] = (n + a + 4) & 0x3f
		a = data[i]
		password = append(password, PassKana[a])
	}
	return password, nil
}

func formatPassword(c []rune) string {
	b := bytes.Buffer{}
	for i, r := range c {
		switch i {
		case 5, 12, 17:
			b.WriteRune(' ')
		}
		b.WriteRune(r)
	}
	return b.String()
}
