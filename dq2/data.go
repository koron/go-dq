package dq2

// PassKana is 復活の呪文用の「かな」
var PassKana = []rune{
	'あ', 'い', 'う', 'え', 'お', 'か', 'き', 'く', 'け', 'こ',
	'さ', 'し', 'す', 'せ', 'そ', 'た', 'ち', 'つ', 'て', 'と',
	'な', 'に', 'ぬ', 'ね', 'の', 'は', 'ひ', 'ふ', 'へ', 'ほ',
	'ま', 'み', 'む', 'め', 'も', 'や', 'ゆ', 'よ', 'ら', 'り',
	'る', 'れ', 'ろ', 'わ', 'が', 'ぎ', 'ぐ', 'げ', 'ご', 'ざ',
	'じ', 'ず', 'ぜ', 'ぞ', 'ば', 'び', 'ぶ', 'べ', 'ぼ', 'ぱ',
	'ぴ', 'ぷ', 'ぺ', 'ぽ',
}

// NameKana is 名前用の「かな」
var NameKana = []rune{
	'あ', 'い', 'う', 'え', 'お', 'か', 'き', 'く', 'け', 'こ',
	'さ', 'し', 'す', 'せ', 'そ', 'た', 'ち', 'つ', 'て', 'と',
	'な', 'に', 'ぬ', 'ね', 'の', 'は', 'ひ', 'ふ', 'へ', 'ほ',
	'ま', 'み', 'む', 'め', 'も', 'や', 'ゆ', 'よ', 'ら', 'り',
	'る', 'れ', 'ろ', 'わ', 'を', 'ん', 'っ', 'ゃ', 'ゅ', 'ょ',
	'゛', '゜', '　',
}

// Items is アイテム一覧
var Items = []string{
	"", "ひのきのぼう", "せいなるナイフ", "まどうしのつえ",
	"いかずちのつえ", "こんぼう", "どうのつるぎ", "くさりがま",
	"てつのやり", "はやぶさのけん", "はがねのつるぎ", "おおかなずち",
	"はかいのつるぎ", "ドラゴンキラー", "ひかりのつるぎ", "ロトのつるぎ",
	"いなずまのけん", "ぬののふく", "みかわしのふく", "みずのはごろも",
	"ミンクのコート", "かわのよろい", "くさりかたびら", "あくまのよろい",
	"まほうのよろい", "はがねのよろい", "ガイアのよろい", "ロトのよろい",
	"かわのたて", "ちからのたて", "はがねのたて", "しにがみのたて",
	"ロトのたて", "ふしぎなぼうし", "てつかぶと", "ロトのかぶと",
	"ロトのしるし", "ふねのざいほう", "つきのかけら", "ルビスのまもり",
	"じゃしんのぞう", "せかいじゅのは", "やまびこのふえ", "ラーのかがみ",
	"あまつゆのいと", "せいなりおりき", "かぜのマント", "あくまのしっぽ",
	"まよけのすず", "ふっかつのたま", "ゴールドカード", "ふくびきけん",
	"せいすい", "キメラのつばさ", "みみせん", "きんのかぎ",
	"ぎんのかぎ", "ろうやのかぎ", "すいもんのかぎ", "どくけしそう",
	"やくそう", "いのりのゆびわ", "しのオルゴール", "あぶないみずぎ",
}

// CheckPoints is 復活する場所一覧
var CheckPoints = []string{
	"ローレシア", "サマルトリア", "ラダトーム", "デルコンダル",
	"ベラヌール", "ロンダルキア", "ムーンペタ",
}

// Crests is 紋章一覧
var Crests = []string{"太陽", "星", "月", "水", "命"}

// SamantoriaNames is サマルトリアの王子の名前一覧
var SamantoriaNames = []string{
	"パウロ", "ランド", "カイン", "アーサー",
	"コナン", "クッキー", "トンヌラ", "すけさん",
}

// MoonbrookeNames is ムーンブルクの王女の名前一覧
var MoonbrookeNames = []string{
	"まいこ", "リンダ", "サマンサ", "アイリン",
	"マリア", "ナナ", "あきな", "プリン",
}
