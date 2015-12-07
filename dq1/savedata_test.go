package dq1

import "testing"

func checkPassword(t *testing.T, d *SaveData, v int, expected string) {
	s, err := d.Password(v)
	if err != nil {
		t.Fatal(err)
	}
	if s != expected {
		t.Fatalf("failed %q #%d expected %q but actually %q",
			d.Name, v, expected, s)
	}
}

func TestPassword(t *testing.T) {
	d := &SaveData{
		Name: "とんぬら",
	}
	checkPassword(t, d, 0, "れぎざぶい かころぐじでぶ いかこせつ せねふ")
	checkPassword(t, d, 1, "こたとぼえ くしがごぜばぼ えくしたと たはほ")
	checkPassword(t, d, 2, "ゆれぎでぶ いにだどべかこ せつにはほ はやり")
	checkPassword(t, d, 4, "るぐじべう きさわげずうき さそてぬひ ぬむゆ")
	checkPassword(t, d, 7, "かさそどべ うぬぢばぼそて ぬひまもら もがご")
}

func TestPasswordMax(t *testing.T) {
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
	checkPassword(t, d, 0, "こさほれぎ ぎあめよれむや らこなのふ ひゆる")
	checkPassword(t, d, 7, "よよづゆる るしがごぜずだ でへりわぐ ぎぢば")
}
