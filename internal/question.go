package internal

// Question Document
/*
 *------------------------------------------------------
 * @brief Question
 *
 * @param[in]
 *
 * @return[want] string
 * @return[want] string
 *
 * @rules
 *
 *------------------------------------------------------
 */
func Question() (string, string) {

	q := map[string]string{
		"ありがとうございます": "ありがとうございます",
		"どういたしまして":   "どういたしまして",
		"おはようございます":  "おはようございます",
		"こんにちは":      "こんにちは",
		"こんばんは":      "こんばんは",
		"お買い物":       "おかいもの",
		"お仕事中":       "おしごとちゅう",
		"趣味に没頭":      "しゅみにぼっとう",
		"映画鑑賞":       "えいがかんしょう",
		"飲食店":        "いんしょくてん",
		"最新機器":       "さいしんきき",
		"電話中":        "でんわちゅう",
		"弦楽器":        "げんがっき",
		"管楽器":        "かんがっき",
		"歌唱力":        "かしょうりょく",
		"世界大戦":       "せかいたいせん",
		"人道的":        "じんどうてき",
		"独裁者":        "どくさいしゃ",
		"偽善者":        "ぎぜんしゃ",
		"働き者":        "はたらきもの",
		"怠け者":        "なまけもの",
		"自堕落":        "じだらく",
		"早寝早起き":      "はやねはやおき",
		"大富豪":        "だいふごう",
		"麻雀":         "まーじゃん",
		"趣味趣向":       "しゅみしゅこう",
		"出向":         "しゅっこう",
		"出発進行":       "しゅっぱつしんこう",
		"一触即発":       "いっしょくそくはつ",
		"雰囲気":        "ふんいき",
		"最高傑作":       "さいこうけっさく",
		"会心の出来":      "かいしんのでき",
		"努力":         "どりょく",
		"目が覚める":      "めがさめる",
		"目が充血":       "めがじゅうけつ",
		"煮込み料理":      "にこみりょうり",
		"目分量":        "めぶんりょう",
		"先行投資":       "せんこうとうし",
		"余計な一言":      "よけいなひとこと",
		"眉目秀麗":       "びもくしゅうれい",
		"頭脳明晰":       "ずのうめいせき",
		"筋骨隆々":       "きんこつりゅうりゅう",
		"相対性理論":      "そうたいせいりろん",
		"超新星爆発":      "ちょうしんせいばくはつ",
		"お天気速報":      "おてんきそくほう",
		"疲労困憊":       "ひろうこんぱい",
		"コンセント":      "こんせんと",
		"一向一揆":       "いっこういっき",
		"廃藩置県":       "はいはんちけん",
		"背水の陣":       "はいすいのじん",
		"三度目の正直":     "さんどめのしょうじき",
		"遺伝子操作":      "いでんしそうさ",
		"科学の発展":      "かがくのはってん",
		"魅力的な人物":     "みりょくてきなじんぶつ",
		"新たな発見":      "あらたなはっけん",
		"猫じゃらし":      "ねこじゃらし",
		"健康第一":       "けんこうだいいち",
		"老若男女":       "ろうにゃくなんにょ",
		"石橋を叩いて渡る":   "いしばしをたたいてわたる",
		"初日の出":       "はつひので",
		"驚愕の事実":      "きょうがくのじじつ",
		"思い出を語る":     "おもいでをかたる",
		"傾聴":         "けいちょう",
		"開発者":        "かいはつしゃ",
		"桃源郷":        "とうげんきょう",
		"周回要素":       "しゅうかいようそ",
		"葉緑素":        "ようりょくそ",
		"暑中見舞い":      "しょちゅうみまい",
		"空前絶後":       "くうぜんぜつご",
		"狂喜乱舞":       "きょうきらんぶ",
		"焼肉定食":       "やきにくていしょく",
		"目玉焼き":       "めだまやき",
		"調理器具":       "ちょうりきぐ",
		"時限要素":       "じげんようそ",
		"同窓会":        "どうそうかい",
		"問題数":        "もんだいすう",
		"同衾":         "どうきん",
		"長距離":        "ちょうきょり",
		"警備隊":        "けいびたい",
		"装飾過多":       "そうしょくかた",
		"顕微鏡":        "けんびきょう",
		"搦手":         "からめて",
	}

	for k, v := range q {
		return k, v
	}
	return "", ""
}
