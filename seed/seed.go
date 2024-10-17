package seed

import (
	"api/model"
	"fmt"

	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) error {
	var count int64
	db.Model(&model.Word{}).Count(&count)
	if count > 0 {
		fmt.Println("Seed data already exists, skipping seeding")
		return nil
	}

	words1 := []model.Word{
		{English: "achievement", Japanese: "達成", Mp3Path: "/audio/achievement.mp3"},
		{English: "allocate", Japanese: "割り当てる", Mp3Path: "/audio/allocate.mp3"},
		{English: "applicant", Japanese: "応募者", Mp3Path: "/audio/applicant.mp3"},
		{English: "appointment", Japanese: "予約", Mp3Path: "/audio/appointment.mp3"},
		{English: "assess", Japanese: "評価する", Mp3Path: "/audio/assess.mp3"},
		{English: "attend", Japanese: "出席する", Mp3Path: "/audio/attend.mp3"},
		{English: "budget", Japanese: "予算", Mp3Path: "/audio/budget.mp3"},
		{English: "candidate", Japanese: "候補者", Mp3Path: "/audio/candidate.mp3"},
		{English: "conference", Japanese: "会議", Mp3Path: "/audio/conference.mp3"},
		{English: "confirm", Japanese: "確認する", Mp3Path: "/audio/confirm.mp3"},
		{English: "consult", Japanese: "相談する", Mp3Path: "/audio/consult.mp3"},
		{English: "contribute", Japanese: "貢献する", Mp3Path: "/audio/contribute.mp3"},
		{English: "deliver", Japanese: "配達する", Mp3Path: "/audio/deliver.mp3"},
		{English: "department", Japanese: "部門", Mp3Path: "/audio/department.mp3"},
		{English: "efficient", Japanese: "効率的な", Mp3Path: "/audio/efficient.mp3"},
		{English: "estimate", Japanese: "見積もる", Mp3Path: "/audio/estimate.mp3"},
		{English: "expand", Japanese: "拡大する", Mp3Path: "/audio/expand.mp3"},
		{English: "facility", Japanese: "施設", Mp3Path: "/audio/facility.mp3"},
		{English: "frequent", Japanese: "頻繁な", Mp3Path: "/audio/frequent.mp3"},
		{English: "implement", Japanese: "実行する", Mp3Path: "/audio/implement.mp3"},
		{English: "income", Japanese: "収入", Mp3Path: "/audio/income.mp3"},
		{English: "invoice", Japanese: "請求書", Mp3Path: "/audio/invoice.mp3"},
		{English: "manage", Japanese: "管理する", Mp3Path: "/audio/manage.mp3"},
		{English: "negotiation", Japanese: "交渉", Mp3Path: "/audio/negotiation.mp3"},
		{English: "offer", Japanese: "提供する", Mp3Path: "/audio/offer.mp3"},
		{English: "participate", Japanese: "参加する", Mp3Path: "/audio/participate.mp3"},
		{English: "policy", Japanese: "方針", Mp3Path: "/audio/policy.mp3"},
		{English: "procedure", Japanese: "手続き", Mp3Path: "/audio/procedure.mp3"},
		{English: "schedule", Japanese: "予定", Mp3Path: "/audio/schedule.mp3"},
		{English: "submit", Japanese: "提出する", Mp3Path: "/audio/submit.mp3"},
	}

	if err := db.Create(&words1).Error; err != nil {
		return err
	}

	lesson1 := model.Lesson{
		Title:       "TOEICで出る重要単語30選",
		Description: "このレッスンでは、TOEICの試験で頻出する30個の重要単語を学びます。ビジネスシーンや日常生活で役立つ単語を、効率よく覚えることができます。各単語には音声も用意されており、発音を確認しながら単語を身につけましょう。",
	}

	if err := db.Create(&lesson1).Error; err != nil {
		return err
	}

	var lessonWords1 []model.LessonWord
	for _, word := range words1 {
		lw := model.LessonWord{
			LessonID: lesson1.ID,
			WordID:   word.ID,
		}
		lessonWords1 = append(lessonWords1, lw)
	}

	if err := db.Create(&lessonWords1).Error; err != nil {
		return err
	}

	fmt.Println("Seed data inserted successfully")
	return nil
}

func SeedLesson2(db *gorm.DB) error {
	var count int64
	db.Model(&model.Word{}).Where("english IN (?)", []string{"stand", "sit", "hold", "look at", "walk"}).Count(&count)
	if count > 0 {
		fmt.Println("Lesson 2 seed data already exists, skipping seeding")
		return nil
	}

	words2 := []model.Word{
		{English: "stand", Japanese: "立っている", Mp3Path: "/audio/stand.mp3"},
		{English: "sit", Japanese: "座っている", Mp3Path: "/audio/sit.mp3"},
		{English: "hold", Japanese: "持っている", Mp3Path: "/audio/hold.mp3"},
		{English: "look at", Japanese: "見ている", Mp3Path: "/audio/look_at.mp3"},
		{English: "walk", Japanese: "歩いている", Mp3Path: "/audio/walk.mp3"},
		{English: "carry", Japanese: "運んでいる", Mp3Path: "/audio/carry.mp3"},
		{English: "write", Japanese: "書いている", Mp3Path: "/audio/write.mp3"},
		{English: "read", Japanese: "読んでいる", Mp3Path: "/audio/read.mp3"},
		{English: "point", Japanese: "指を指している", Mp3Path: "/audio/point.mp3"},
		{English: "open", Japanese: "開ける", Mp3Path: "/audio/open.mp3"},
		{English: "close", Japanese: "閉じる", Mp3Path: "/audio/close.mp3"},
		{English: "wear", Japanese: "着ている", Mp3Path: "/audio/wear.mp3"},
		{English: "put on", Japanese: "着るところ", Mp3Path: "/audio/put_on.mp3"},
		{English: "take off", Japanese: "脱ぐ", Mp3Path: "/audio/take_off.mp3"},
		{English: "talk", Japanese: "話している", Mp3Path: "/audio/talk.mp3"},
		{English: "listen to", Japanese: "聞いている", Mp3Path: "/audio/listen_to.mp3"},
		{English: "use", Japanese: "使っている", Mp3Path: "/audio/use.mp3"},
		{English: "clean", Japanese: "掃除する", Mp3Path: "/audio/clean.mp3"},
		{English: "cook", Japanese: "料理する", Mp3Path: "/audio/cook.mp3"},
		{English: "serve", Japanese: "提供する", Mp3Path: "/audio/serve.mp3"},
		{English: "push", Japanese: "押す", Mp3Path: "/audio/push.mp3"},
		{English: "pull", Japanese: "引く", Mp3Path: "/audio/pull.mp3"},
		{English: "fill", Japanese: "満たす", Mp3Path: "/audio/fill.mp3"},
		{English: "empty", Japanese: "空にする", Mp3Path: "/audio/empty.mp3"},
		{English: "lift", Japanese: "持ち上げる", Mp3Path: "/audio/lift.mp3"},
		{English: "lower", Japanese: "下げる", Mp3Path: "/audio/lower.mp3"},
		{English: "turn on", Japanese: "つける", Mp3Path: "/audio/turn_on.mp3"},
		{English: "turn off", Japanese: "消す", Mp3Path: "/audio/turn_off.mp3"},
		{English: "repair", Japanese: "修理する", Mp3Path: "/audio/repair.mp3"},
		{English: "paint", Japanese: "塗る", Mp3Path: "/audio/paint.mp3"},
	}

	if err := db.Create(&words2).Error; err != nil {
		return err
	}

	lesson2 := model.Lesson{
		Title:       "TOEIC Part 1 で出る重要動詞30選",
		Description: "このレッスンでは、TOEIC Part 1（写真描写問題）で頻出する30個の重要な動詞を学びます。日常生活やビジネスの場面で使われる動作を表す単語を、写真をイメージしながら効率よく覚えることができます。各単語には発音を確認できる音声も付いており、リスニングと発音の練習にも最適です。",
	}

	if err := db.Create(&lesson2).Error; err != nil {
		return err
	}

	var lessonWords2 []model.LessonWord
	for _, word := range words2 {
		lw := model.LessonWord{
			LessonID: lesson2.ID,
			WordID:   word.ID,
		}
		lessonWords2 = append(lessonWords2, lw)
	}

	if err := db.Create(&lessonWords2).Error; err != nil {
		return err
	}

	fmt.Println("Lesson 2 seed data inserted successfully")
	return nil
}

func SeedLesson3(db *gorm.DB) error {
	var count int64
	db.Model(&model.Word{}).Where("english IN (?)", []string{
		"available", "traffic", "identification", "agenda", "deposit",
		"store", "training", "employee", "employer", "technician",
		"bill", "fee", "fare", "sales figures", "office equipment",
		"office supplies", "registration", "shipment", "ship", "reserve",
		"distribute", "reimburse", "supplier", "board", "accounting",
		"sales", "personnel", "human resources", "report to", "run out",
		"drop by", "draw up", "fill out", "prefer A to B", "hand out",
		"take place", "pass out", "take part in", "sign up",
		"responsible for", "in charge of", "out of stock",
	}).Count(&count)
	if count > 0 {
		fmt.Println("Lesson 3 seed data already exists, skipping seeding")
		return nil
	}

	words3 := []model.Word{
		{English: "available", Japanese: "利用可能", Mp3Path: "/audio/available.mp3"},
		{English: "traffic", Japanese: "交通", Mp3Path: "/audio/traffic.mp3"},
		{English: "identification", Japanese: "識別", Mp3Path: "/audio/identification.mp3"},
		{English: "agenda", Japanese: "議題", Mp3Path: "/audio/agenda.mp3"},
		{English: "deposit", Japanese: "預金", Mp3Path: "/audio/deposit.mp3"},
		{English: "store", Japanese: "店", Mp3Path: "/audio/store.mp3"},
		{English: "training", Japanese: "訓練", Mp3Path: "/audio/training.mp3"},
		{English: "employee", Japanese: "従業員", Mp3Path: "/audio/employee.mp3"},
		{English: "employer", Japanese: "雇用主", Mp3Path: "/audio/employer.mp3"},
		{English: "technician", Japanese: "技術者", Mp3Path: "/audio/technician.mp3"},
		{English: "bill", Japanese: "請求書", Mp3Path: "/audio/bill.mp3"},
		{English: "fee", Japanese: "料金", Mp3Path: "/audio/fee.mp3"},
		{English: "fare", Japanese: "運賃", Mp3Path: "/audio/fare.mp3"},
		{English: "sales figures", Japanese: "売上高", Mp3Path: "/audio/sales_figures.mp3"},
		{English: "office equipment", Japanese: "オフィス機器", Mp3Path: "/audio/office_equipment.mp3"},
		{English: "office supplies", Japanese: "事務用品", Mp3Path: "/audio/office_supplies.mp3"},
		{English: "registration", Japanese: "登録", Mp3Path: "/audio/registration.mp3"},
		{English: "shipment", Japanese: "出荷", Mp3Path: "/audio/shipment.mp3"},
		{English: "ship", Japanese: "船", Mp3Path: "/audio/ship.mp3"},
		{English: "reserve", Japanese: "予約する", Mp3Path: "/audio/reserve.mp3"},
		{English: "distribute", Japanese: "分配する", Mp3Path: "/audio/distribute.mp3"},
		{English: "reimburse", Japanese: "払い戻す", Mp3Path: "/audio/reimburse.mp3"},
		{English: "supplier", Japanese: "供給者", Mp3Path: "/audio/supplier.mp3"},
		{English: "board", Japanese: "取締役会", Mp3Path: "/audio/board.mp3"},
		{English: "accounting", Japanese: "会計", Mp3Path: "/audio/accounting.mp3"},
		{English: "sales", Japanese: "売上", Mp3Path: "/audio/sales.mp3"},
		{English: "personnel", Japanese: "人事", Mp3Path: "/audio/personnel.mp3"},
		{English: "human resources", Japanese: "人的資源", Mp3Path: "/audio/human_resources.mp3"},
		{English: "report to", Japanese: "～に報告する", Mp3Path: "/audio/report_to.mp3"},
		{English: "run out", Japanese: "切れる", Mp3Path: "/audio/run_out.mp3"},
		{English: "drop by", Japanese: "立ち寄る", Mp3Path: "/audio/drop_by.mp3"},
		{English: "draw up", Japanese: "作成する", Mp3Path: "/audio/draw_up.mp3"},
		{English: "fill out", Japanese: "記入する", Mp3Path: "/audio/fill_out.mp3"},
		{English: "prefer A to B", Japanese: "AをBより好む", Mp3Path: "/audio/prefer_a_to_b.mp3"},
		{English: "hand out", Japanese: "配る", Mp3Path: "/audio/hand_out.mp3"},
		{English: "take place", Japanese: "行われる", Mp3Path: "/audio/take_place.mp3"},
		{English: "pass out", Japanese: "配る", Mp3Path: "/audio/pass_out.mp3"},
		{English: "take part in", Japanese: "参加する", Mp3Path: "/audio/take_part_in.mp3"},
		{English: "sign up", Japanese: "申し込む", Mp3Path: "/audio/sign_up.mp3"},
		{English: "responsible for", Japanese: "～の責任がある", Mp3Path: "/audio/responsible_for.mp3"},
		{English: "in charge of", Japanese: "～を担当している", Mp3Path: "/audio/in_charge_of.mp3"},
		{English: "out of stock", Japanese: "売り切れ", Mp3Path: "/audio/out_of_stock.mp3"},
	}

	if err := db.Create(&words3).Error; err != nil {
		return err
	}

	lesson3 := model.Lesson{
		Title:       "TOEIC Part2で出てくる重要な用語43選",
		Description: "このレッスンでは、TOEIC Part2（質問応答）で頻繁に出題される43の重要な英単語とフレーズを学びます。これらの単語は、日常会話やビジネスシーンでよく使われる表現です。各単語には日本語訳と発音確認用の音声が付いており、リスニングとスピーキングのスキル向上に役立ちます。",
	}

	if err := db.Create(&lesson3).Error; err != nil {
		return err
	}

	var lessonWords3 []model.LessonWord
	for _, word := range words3 {
		lw := model.LessonWord{
			LessonID: lesson3.ID,
			WordID:   word.ID,
		}
		lessonWords3 = append(lessonWords3, lw)
	}

	if err := db.Create(&lessonWords3).Error; err != nil {
		return err
	}

	fmt.Println("Lesson 3 seed data inserted successfully")
	return nil
}

func SeedLesson4(db *gorm.DB) error {
	var count int64
	db.Model(&model.Word{}).Where("english IN (?)", []string{
		"due", "delayed", "inclusive", "legible", "fragile", "superb", "thrilled",
		"spacious", "unfortunately", "overlook", "recommend", "involve", "commute",
		"prohibit", "conduct", "disrupt", "undergo", "exclude", "finalize", "foresee",
		"reflect", "postpone", "evaluate", "admission", "summary", "flaw", "extension",
		"draft", "pension", "put off", "put forward", "focus on",
		"drop off", "attribute A to B", "lead to", "attend to", "subscribe to",
		"serve as", "keep up with", "take for granted", "after all", "up to", "on track",
	}).Count(&count)
	if count > 0 {
		fmt.Println("Lesson 4 seed data already exists, skipping seeding")
		return nil
	}

	words4 := []model.Word{
		{English: "due", Japanese: "期日", Mp3Path: "/audio/due.mp3"},
		{English: "delayed", Japanese: "遅延した", Mp3Path: "/audio/delayed.mp3"},
		{English: "inclusive", Japanese: "包括的な", Mp3Path: "/audio/inclusive.mp3"},
		{English: "legible", Japanese: "読みやすい", Mp3Path: "/audio/legible.mp3"},
		{English: "fragile", Japanese: "壊れやすい", Mp3Path: "/audio/fragile.mp3"},
		{English: "superb", Japanese: "素晴らしい", Mp3Path: "/audio/superb.mp3"},
		{English: "thrilled", Japanese: "非常に興奮している", Mp3Path: "/audio/thrilled.mp3"},
		{English: "spacious", Japanese: "広々とした", Mp3Path: "/audio/spacious.mp3"},
		{English: "unfortunately", Japanese: "残念ながら", Mp3Path: "/audio/unfortunately.mp3"},
		{English: "overlook", Japanese: "見落とす", Mp3Path: "/audio/overlook.mp3"},
		{English: "recommend", Japanese: "推奨する", Mp3Path: "/audio/recommend.mp3"},
		{English: "involve", Japanese: "関与する", Mp3Path: "/audio/involve.mp3"},
		{English: "commute", Japanese: "通勤する", Mp3Path: "/audio/commute.mp3"},
		{English: "prohibit", Japanese: "禁止する", Mp3Path: "/audio/prohibit.mp3"},
		{English: "conduct", Japanese: "実施する", Mp3Path: "/audio/conduct.mp3"},
		{English: "disrupt", Japanese: "中断する", Mp3Path: "/audio/disrupt.mp3"},
		{English: "undergo", Japanese: "経験する", Mp3Path: "/audio/undergo.mp3"},
		{English: "exclude", Japanese: "除外する", Mp3Path: "/audio/exclude.mp3"},
		{English: "finalize", Japanese: "最終決定する", Mp3Path: "/audio/finalize.mp3"},
		{English: "foresee", Japanese: "予見する", Mp3Path: "/audio/foresee.mp3"},
		{English: "reflect", Japanese: "反映する", Mp3Path: "/audio/reflect.mp3"},
		{English: "postpone", Japanese: "延期する", Mp3Path: "/audio/postpone.mp3"},
		{English: "evaluate", Japanese: "評価する", Mp3Path: "/audio/evaluate.mp3"},
		{English: "admission", Japanese: "入場", Mp3Path: "/audio/admission.mp3"},
		{English: "summary", Japanese: "要約", Mp3Path: "/audio/summary.mp3"},
		{English: "flaw", Japanese: "欠陥", Mp3Path: "/audio/flaw.mp3"},
		{English: "extension", Japanese: "延長", Mp3Path: "/audio/extension.mp3"},
		{English: "draft", Japanese: "草案", Mp3Path: "/audio/draft.mp3"},
		{English: "pension", Japanese: "年金", Mp3Path: "/audio/pension.mp3"},
		{English: "put off", Japanese: "延期する", Mp3Path: "/audio/put_off.mp3"},
		{English: "put forward", Japanese: "提案する", Mp3Path: "/audio/put_forward.mp3"},
		{English: "focus on", Japanese: "～に焦点を当てる", Mp3Path: "/audio/focus_on.mp3"},
		{English: "drop off", Japanese: "減少する", Mp3Path: "/audio/drop_off.mp3"},
		{English: "attribute A to B", Japanese: "AをBのせいにする", Mp3Path: "/audio/attribute_a_to_b.mp3"},
		{English: "lead to", Japanese: "～につながる", Mp3Path: "/audio/lead_to.mp3"},
		{English: "attend to", Japanese: "対応する", Mp3Path: "/audio/attend_to.mp3"},
		{English: "subscribe to", Japanese: "購読する", Mp3Path: "/audio/subscribe_to.mp3"},
		{English: "serve as", Japanese: "～として機能する", Mp3Path: "/audio/serve_as.mp3"},
		{English: "keep up with", Japanese: "～についていく", Mp3Path: "/audio/keep_up_with.mp3"},
		{English: "take for granted", Japanese: "当然と思う", Mp3Path: "/audio/take_for_granted.mp3"},
		{English: "after all", Japanese: "結局", Mp3Path: "/audio/after_all.mp3"},
		{English: "up to", Japanese: "～まで", Mp3Path: "/audio/up_to.mp3"},
		{English: "on track", Japanese: "順調に進んでいる", Mp3Path: "/audio/on_track.mp3"},
	}

	if err := db.Create(&words4).Error; err != nil {
		return err
	}

	lesson4 := model.Lesson{
		Title:       "TOEIC Part3で出てくる重要な用語43選",
		Description: "このレッスンでは、TOEIC Part3（会話問題）で頻繁に出題される44の重要な英単語とフレーズを学びます。これらの単語は、ビジネスや日常会話の中でよく使われる表現です。各単語には日本語訳と発音確認用の音声が付いており、リスニングとスピーキングのスキル向上に役立ちます。",
	}

	if err := db.Create(&lesson4).Error; err != nil {
		return err
	}

	var lessonWords4 []model.LessonWord
	for _, word := range words4 {
		lw := model.LessonWord{
			LessonID: lesson4.ID,
			WordID:   word.ID,
		}
		lessonWords4 = append(lessonWords4, lw)
	}

	if err := db.Create(&lessonWords4).Error; err != nil {
		return err
	}

	fmt.Println("Lesson 4 seed data inserted successfully")
	return nil
}
