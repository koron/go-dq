package dq1

import "testing"

func testEncode(t *testing.T, d *SaveData, v int) (string, error) {
	p, err := Encode(d, v)
	if err != nil {
		return "", err
	}
	return Format(p), nil
}

func checkEncode(t *testing.T, d *SaveData, v int, expected string) {
	s, err := testEncode(t, d, v)
	if err != nil {
		t.Fatal(err)
	}
	if s != expected {
		t.Fatalf("failed %q #%d expected %q but actually %q",
			d.Name, v, expected, s)
	}
}

func TestEncode(t *testing.T) {
	d := &SaveData{
		Name: "とんぬら",
	}
	checkEncode(t, d, 0, "れぎざぶい かころぐじでぶ いかこせつ せねふ")
	checkEncode(t, d, 1, "こたとぼえ くしがごぜばぼ えくしたと たはほ")
	checkEncode(t, d, 2, "ゆれぎでぶ いにだどべかこ せつにはほ はやり")
	checkEncode(t, d, 4, "るぐじべう きさわげずうき さそてぬひ ぬむゆ")
	checkEncode(t, d, 7, "かさそどべ うぬぢばぼそて ぬひまもら もがご")
}

func TestEncodeMax(t *testing.T) {
	d := &SaveData{
		Name:    "とんぬら",
		Gold:    65535,
		Exp:     65535,
		KeyNum:  7,
		HerbNum: 7,
		Weapon:  "ロトのつるぎ",
		Armor:   "ロトのよろい",
		Shield:  "みかがみのたて",
	}
	checkEncode(t, d, 0, "こさほれぎ ぎあめよれむや らこなのふ ひゆる")
	checkEncode(t, d, 7, "よよづゆる るしがごぜずだ でへりわぐ ぎぢば")
}
