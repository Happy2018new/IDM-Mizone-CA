package hand_data_collected

import "IDM-Mizone-CA/data_type"

// ContantB & RAW (Delete unknown rune)
//
// 雨琦dd篮球少年，大学生篮球联赛开赛啦！
// 🏀「脉动+电解质大学生篮球联赛」今天正式开赛！等你来战！
// 看谁能一路升级，竞逐顶级街球圈 #脉搏不止篮球不休
// 🌀要有球场好状态，别忘了带上雨琦同款脉动+电解质 455mg高能电解质，一口满电，痛快打一场
//
// #宋雨琦脉动品牌代言人 #脉动回来  #篮球  #大学生篮球

// Step 0.
//
// - Publisher table -
var CommentBPublisherTable = data_type.ContentInfo{
	Publisher: data_type.Publisher{
		Name:   "Mizone Chinese",
		Detail: "Official Account",
		URL:    "https://www.xiaohongshu.com/user/profile/5e4d669f0000000001000a36",
	},
	PublishTime:        "2025-04-21",
	PublishContentType: data_type.PublishContentTypeAdvertisement,
	PublishDetail: []string{
		"Published before the content A",
		// TODO: Finish WIP
		"Invite @宋雨琦_G-I-DLE in it (Super star; User Data? <WIP>; https://www.xiaohongshu.com/user/profile/5cc10d72000000001200f6b8)",
	},
	Slogan: []string{
		"脉动+电解质",
		"一口满电455",
		"脉动+电解质篮球赛, 等你来战!",
		"大学生篮球联赛",
		"轻量顶级街球圈, 痛快打一场",
		"看谁能一路升级, 竞逐顶级街球圈",
	},
	HowTo: []string{"线下举办面向大学生的篮球活动"},
}

// Step 1.
//
// - Video table (Hit 2) -
var ContentBVideoTable = data_type.ContentGeneralFeedback{
	HashTag: []string{
		"#脉搏不止篮球不休",
		"#宋雨琦脉动品牌代言人",
		"#脉动回来",
		"#篮球",
		"#大学生篮球",
	},
	TrendsType:       data_type.ContentGeneralFeedbackTrendsTypeSuperStarSuggestion,
	TrendsDetail:     "通过演示代言人打篮球，然后通过喝脉动“找回状态”、“一口满电455” 来推销脉动“电解质”产品，同时开展“脉动+电解质篮球赛”吸引大学生群体购买",
	SpecificAccounts: []data_type.Publisher{CommentBPublisherTable.Publisher},
	KnownViews:       false,
	Views:            0,
	Likes:            6725,
	CommentCounts:    214,
	Stars:            1868,
	StatisticalTime:  "2025-05-14 08:46",
}

// - Hit 3 -
var ContentBHit3 = data_type.ContentInv{
	UserContentType: data_type.UserContentTypeReels,
	Filters:         []string{"00:27+: 以脉动瓶子窗户的滤镜收束并给出“大学生篮球赛”的具体信息"},
	Captions: []string{
		"口号和号召部分使用“方块体字幕” (e.g. 00:16 脉动+电解质)",
		"口号和号召的字幕还在周围加了“闪电”符号，符合脉动本次活动要推销的产品，也就是“电解质”",
		"普通字幕呈现蓝色色调，给人以脉动的基调印象",
		"视频结尾的大学生活动号召文本使用标准的脉动字体，并且均为深蓝色的字体",
	},
	VisualAesthetics: []string{
		"00:02: 小型汗滴动图突出代言人打篮球后“失去能量”需要脉动补充",
		"00:05~00:08: 代言人接过脉动后闪电特效一闪而过(体现本次产品“电解质”主题)，并且代言人换装为 Mizone 455 蓝色衣服",
		"00:08~00:10: 代言人喝电解质脉动时，每喝一口就有一个小闪电图标出现",
		"00:10~00:12: 代言人喝完脉动后，以代言人的眼睛为起点释放了颜色相同但形状不同的闪电特效，表示代言人“满电”",
	},
	Impressions: "脉动贯穿全文",
}

// Step 2.
//
// - Comment Table -
var ContentBCommentTable = []data_type.Comment{
	{
		User:                  data_type.User{Name: "泥的炒猪血爱豪总", UID: "泥的炒猪血爱豪总"},
		CommitTime:            "04-22",
		ContentText:           "雨琦联名的必须买一箱",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "十分支持代言人及其周边 (Also reply, 5 people want, total 6 people reply)",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 81,
		ReplyCount:            11,
		UserBackgroud:         "没有在小红书笔记和收藏中发现关于代言人或脉动的笔记，但是是蛋仔派对玩家",
	}, // 1
	{
		User:                  data_type.User{Name: "宋STAR", UID: "2770401679"},
		CommitTime:            "04-21",
		ContentText:           "**一张代言人穿着与视频同款的 Mizone 455 蓝色衣服在夜晚的海边的图片**",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "希望借图片表现代言人很美? (评论区回复代言人“好美啊”)",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 46,
		ReplyCount:            2,
		UserBackgroud:         "爱宋雨琦，檀健次，并且追剧",
	}, // 2
	{
		User:                  data_type.User{Name: "Yu", UID: "1516975180"},
		CommitTime:            "04-30",
		ContentText:           "**一个关于和代言人结婚的表情包**",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "评论区在通过表情包讨论谁应该获得跟代言人的结婚证，十分搞笑活泼",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 9,
		ReplyCount:            8,
		UserBackgroud:         "小红书自述主担宋雨琦",
	}, // 3
	{
		User:                  data_type.User{Name: "笑", UID: "5615575674"},
		CommitTime:            "04-22",
		ContentText:           "谁家雨琦最好看？",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "评论区在讨论代言人应该属于谁，个个都在表达对代言人的喜爱",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 5,
		ReplyCount:            4,
		UserBackgroud:         "其小红书收藏了四篇与代言人有关的笔记，但未发现与脉动有关的",
	}, // 4
	{
		User:                  data_type.User{Name: "叼着小鱼干", UID: "11629096632"},
		CommitTime:            "04-21",
		ContentText:           "代言人宋雨琦太美啦+一个代言人穿着蓝色半透明裙子的图片",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "表达对代言人的喜爱",
		UserFeedbackToneLevel: data_type.UserFeedbackToneVeryHappy, // 0
		PeopleStarThisComment: 15,
		ReplyCount:            0,
		UserBackgroud:         "其发布的小红书笔记中有大量关于代言人的，可能是一名代言人重度粉丝",
	}, // 5
	{
		User:                  data_type.User{Name: "第一狗塑-", UID: "796329072"},
		CommitTime:            "04-21",
		ContentText:           "脉脉可以让代言人教我打篮球吗+一个代言人穿着蓝色半透明裙子的图片",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "用户反映出“让代言人教我打篮球”，并且以亲切的语气称呼脉动为“脉脉”，体现用户对代言人和脉动活动的喜爱",
		UserFeedbackToneLevel: data_type.UserFeedbackToneVeryHappy, // 0
		PeopleStarThisComment: 8,
		ReplyCount:            1,
		UserBackgroud:         "其小红书笔记中基本都是关于代言人的，并且基本上皆是各种声援代言人的文本笔记",
	}, // 6
	{
		User:                  data_type.User{Name: "龙女", UID: "2740798864"},
		CommitTime:            "05-01",
		ContentText:           "@祈羽铃颂（曲琦） (回复: 我已经喝上了)",
		HandledSrc:            "可能因为朋友是代言人粉丝所以被推流了，但其本身算是路人?",
		HandledTone:           "应该是推荐给朋友看视频，然后他朋友也回复并称自己已经喝上了脉动",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "蛋仔派对玩家",
	}, // 7
	{
		User:                  data_type.User{Name: "饰礁", UID: "9655126339"},
		CommitTime:            "05-01",
		ContentText:           "@念渔Aa @＾娜水音🥥 妈呀帅死我了",
		HandledSrc:            "可能是了解一些关于代言人的内容，然后被平台推流了",
		HandledTone:           "推荐给朋友看视频，并且表示代言人打篮球很帅",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		PeopleStarThisComment: 2,
		ReplyCount:            1,
		UserBackgroud:         "似乎小红书头像是代言人，但其没有发布小红书笔记，且收藏的笔记中没有发现太多关于代言人或脉动的内容；可能追星",
	}, // 8
	{
		User:                  data_type.User{Name: "淋漠本漠💕", UID: "116853287A"},
		CommitTime:            "05-01",
		ContentText:           "瓶子上面印着一个雨琦，我天天喝。+一个抽象的用于表达对代言人喜爱的表情包",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "表达对代言人的喜爱，并且很支持代言人代言的脉动产品",
		UserFeedbackToneLevel: data_type.UserFeedbackToneVeryHappy, // 0
		PeopleStarThisComment: 3,
		ReplyCount:            0,
		UserBackgroud:         "代言人粉丝，并且很喜欢买她的周边，是蛋仔派对游戏的主播",
	}, // 9
	{
		User:                  data_type.User{Name: "淋漠本漠💕", UID: "11516607492"},
		CommitTime:            "04-25",
		ContentText:           "我打篮球喝了好好喝",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "第一层意思是表达脉动的产品很好用，第二层意思是其作为代言人粉丝的身份，支持代言人的产品",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "头像和自述都表明其为代言人粉丝；是女生；学生；15岁",
	}, // 10
	{
		User:                  data_type.User{Name: "小珠粉丝头子", UID: "1388809417"},
		CommitTime:            "04-25",
		ContentText:           "我去，以后除了脉动，其他的饮料有我一律不喝",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "通过表达以后买饮料只买脉动，体现了自己很喜欢脉动产品",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "小红书笔记和收藏中各找到一篇关于代言人的内容，可能是轻度粉丝；女生；初中生?；王者，蛋仔，沙维玛，和平精英玩家",
	}, // 11
	{
		User:                  data_type.User{Name: "。", UID: "153252053"},
		CommitTime:            "04-25",
		ContentText:           "@love琦 Yuqi",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "可能是艾特朋友来看视频",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "女生，且小红书中收藏了多篇关于代言人的内容，尚且未发现跟脉动有关的内容",
	}, // 12
	{
		User:                  data_type.User{Name: "跨年悸少", UID: "42931444100"},
		CommitTime:            "04-24",
		ContentText:           "我今天去买一箱",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "表达自己很支持代言人推荐的产品",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "小红书头像似乎是代言人，并且其收藏了一篇关于代言人的内容，没有发现与脉动有关的内容",
	}, // 13
	{
		User:                  data_type.User{Name: "安利手^鹿雨欣🌙🎀", UID: "577340812"},
		CommitTime:            "04-24",
		ContentText:           "@🌷雨琦🌷-跟班",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "可能是艾特朋友来看视频 (朋友回复了一个赞的表情)",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 1,
		ReplyCount:            1,
		UserBackgroud:         "女生，主推白鹿，副推代言人",
	}, // 14
	{
		User:                  data_type.User{Name: "哦好的哦", UID: "94718122101"},
		CommitTime:            "04-24",
		ContentText:           "感觉电解质瓶装设计一股男风，其实女孩也运动啊，这瓶子太工业风",
		HandledSrc:            "可能是路人?",
		HandledTone:           "希望脉动的设计可以贴近女性化偏好，并且表达对目前男风工业风的不满",
		UserFeedbackToneLevel: data_type.UserFeedbackToneUnhappy, // 3
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "女生，大学生，学习党?",
	}, // 15
	{
		User:                  data_type.User{Name: "番茄不炒蛋(禁言中)", UID: "2894978461"},
		CommitTime:            "04-24",
		ContentText:           "我要买100箱",
		HandledSrc:            "可能了解代言人的粉丝?",
		HandledTone:           "希望通过大量购入脉动来支持代言人或者脉动的产品",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "女生；蛋仔派对玩家；小红书笔记和收藏中没有发现与代言人或脉动有关的内容",
	}, // 16
	{
		User:                  data_type.User{Name: "沈梦雨🍪", UID: "8934972278"},
		CommitTime:            "04-23",
		ContentText:           "告诉我这个品牌在哪买",
		HandledSrc:            "可能是代言人的粉丝?",
		HandledTone:           "希望能从评论区知晓脉动的这个新产品如何购买，可能是因为代言人慕名而来",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "女生，似乎追星",
	}, // 17
	{
		User:                  data_type.User{Name: "蓝莓曲奇🍪", UID: "11578437661"},
		CommitTime:            "04-23",
		ContentText:           "好美",
		HandledSrc:            "代言人粉丝?",
		HandledTone:           "表达了对代言人美貌的高度评价",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "小红书的收藏的三个笔记中有一个是与代言人有关的，可能是代言人的粉丝",
	}, // 18
	{
		User:                  data_type.User{Name: "美神。。", UID: "loveyuq1"},
		CommitTime:            "04-23",
		ContentText:           "美神。。",
		HandledSrc:            "代言人粉丝?",
		HandledTone:           "表达了对代言人美貌的高度评价",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "小红书的头像是代言人，双鱼座",
	}, // 19
	{
		User:                  data_type.User{Name: "别下雨啦🥛", UID: "9644342921"},
		CommitTime:            "04-23",
		ContentText:           "**一个眼睛放光嘴巴张开表示很赞叹的表情包**",
		HandledSrc:            "代言人粉丝?",
		HandledTone:           "表达了对代言人或者脉动产品的赞叹?",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "其在小红书发布的三个内容中有一个是与代言人有关的",
	}, // 20
	{
		User:                  data_type.User{Name: "折翠云", UID: "2878634069"},
		CommitTime:            "04-23",
		ContentText:           "**一个代言人比心的表情包**",
		HandledSrc:            "可能是喜欢代言人的轻度粉丝(或者路人?)",
		HandledTone:           "表达了对代言人的喜欢",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "没有在他小红书发现与代言人或者脉动有关的内容",
	}, // 21
	{
		User:                  data_type.User{Name: "珉珉珉珉珉.", UID: "6223502275"},
		CommitTime:            "04-23",
		ContentText:           "我靠",
		HandledSrc:            "可能是喜欢代言人的粉丝(或者路人?)",
		HandledTone:           "表现自己对脉动或代言人的惊叹?",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "女生；小红书收藏没有对外公开，其也没有发布任何笔记，无法知晓更多消息",
	}, // 22
	{
		User:                  data_type.User{Name: "失一棠爱", UID: "7437142768"},
		CommitTime:            "04-22",
		ContentText:           "雨琦宝宝我爱你",
		HandledSrc:            "可能是喜欢代言人的轻度粉丝(或者路人?)",
		HandledTone:           "表达自己对代言人的喜爱",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "女生，学生，自称幼稚园小宝宝；小红书收藏和发布的笔记中没有发现与脉动或者跟代言人有关的内容",
	}, // 23
	{
		User:                  data_type.User{Name: "嘎嘎乐吐司🍞", UID: "3687875344"},
		CommitTime:            "04-22",
		ContentText:           "那我得多买几箱了",
		HandledSrc:            "可能是喜欢代言人的轻度粉丝(或者路人?)",
		HandledTone:           "通过购买产品来表达对代言人或者脉动的支持",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 1,
		ReplyCount:            0,
		UserBackgroud:         "蛋仔派对玩家，喜马拉雅留学生，13岁；无法从其了解额外的信息",
	}, // 24
	{
		User:                  data_type.User{Name: "帅气小狗愛吃曲琦", UID: "26039034608"},
		CommitTime:            "04-22",
		ContentText:           "第一眼：什麼廣告 第二眼：哇雨琦買一个",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "本来是不打算看广告的，但看到是自家偶像便津津乐道",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "自述自己主担代言人；该用户似乎来自中国香港",
	}, // 25
	{
		User:                  data_type.User{Name: "Dear_雨琦🍪", UID: "3889137370"},
		CommitTime:            "04-22",
		ContentText:           "买一瓶脉动送一个宋雨琦吗？",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "通过调侃希望能通过买脉动买到自家偶像表达对代言人的喜爱",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "小红书收藏了三篇与代言人有关的小红书笔记，可能是代言人的粉丝",
	}, // 26
	{
		User:                  data_type.User{Name: "白念熙.Qi", UID: "95614067773"},
		CommitTime:            "04-22",
		ContentText:           "@蓝海属于你🦋 @^芝士^ @意迟 @郑在想名字．",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "可能是艾特朋友来观看视频",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "小红书自述声称其主担是代言人",
	}, // 27
	{
		User:                  data_type.User{Name: "伊露", UID: "4983608813"},
		CommitTime:            "04-22",
		ContentText:           "@🐟 +拿着马头的无语表情",
		HandledSrc:            "可能是朋友了解代言人，然后被平台推流了",
		HandledTone:           "可能是艾特朋友来观看视频，但表情包可能表现出一些不满?",
		UserFeedbackToneLevel: data_type.UserFeedbackToneUnhappy, // 3
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "女生，但小红书笔记和收藏没有找到跟代言人或脉动相关的内容",
	}, // 28
	{
		User:                  data_type.User{Name: "墨色七月", UID: "8967877366"},
		CommitTime:            "04-22",
		ContentText:           "@瑾榆.琦^",
		HandledSrc:            "未知? 可能是因为其对这些内容可能有兴趣，被平台推流了",
		HandledTone:           "可能是艾特朋友来观看视频",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "(路人?) 女生，蛋仔派对玩家，其小红书收藏没有公开，但对外发布的小红书笔记没有找到跟代言人或脉动相关的内容",
	}, // 29
	{
		User:                  data_type.User{Name: "莞风琦灵", UID: "6803447258"},
		CommitTime:            "04-22",
		ContentText:           "**一个关于和代言人结婚的表情包**",
		HandledSrc:            "可能是了解代言人的粉丝? 然后被平台推流了",
		HandledTone:           "通过表情包很好的表达了自己希望代言人是自己家女朋友这件事，表达了对代言人的喜爱",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            1,
		UserBackgroud:         "其小红书笔记和收藏没有找到跟代言人或跟脉动相关的内容，可能是路人",
	}, // 30
	{
		User:                  data_type.User{Name: "雨琦的粉", UID: "42964624820"},
		CommitTime:            "04-22",
		ContentText:           "厉害",
		HandledSrc:            "可能是代言人的粉丝",
		HandledTone:           "“厉害”二字简单凝练的表达了该用户对代言人及脉动产品的支持",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "在其小红书收藏有两篇与代言人有关",
	}, // 31
	{
		User:                  data_type.User{Name: "法兰西不吃西兰花", UID: "11423290418"},
		CommitTime:            "04-22",
		ContentText:           "**一张代言人穿着与视频同款的 Mizone 455 蓝色衣服在夜晚的海边的照片**",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "通过发送代言人的美照表达对代言人的喜爱",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "其小红书自述其担代言人，似乎还是二次元?",
	}, // 32
	{
		User:                  data_type.User{Name: "盐汽水🥤", UID: "1018422883"},
		CommitTime:            "04-22",
		ContentText:           "@奶昔恋雪～🍭僵🐍🐍",
		HandledSrc:            "可能是朋友了解代言人，然后被平台推流了?",
		HandledTone:           "可能是艾特朋友观看视频",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "女生，但其小红书笔记和收藏基本与男偶像有关，没有找到跟代言人或者跟脉动相关的内容", // May be?
	}, // 33
	{
		User:                  data_type.User{Name: "xtc小帕.告别总在冬季🐰", UID: "878761532"},
		CommitTime:            "04-22",
		ContentText:           "最纯爱那年是宋雨琦代言的东西，我都会疯狂买（好吧，其实现在也是",
		HandledSrc:            "可能是代言人的轻度粉丝",
		HandledTone:           "通过表达自己曾经和现在都疯狂购买代言人的周边来表达自己强推代言人及其代言产品",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "女生，13岁，喜爱小马宝莉和一些女偶像 (e.g. 白鹿)；没有从其小红书收藏和小红书笔记中找到跟脉动或代言人有关的内容",
	}, // 34
	{
		User:                  data_type.User{Name: "小冬日绵绵^winyeon🐰🐶", UID: "9463974836"},
		CommitTime:            "04-22",
		ContentText:           "买一瓶送代言人吗？",
		HandledSrc:            "可能是对这类视频感兴趣，然后被平台算法推流",
		HandledTone:           "调侃能不能通过买脉动得到代言人，而实质可能是表达对代言人的强烈喜爱?", // 说实话我没看出来，所以让我使用 Middle Level
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle,     // 2
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "女生，15岁；没有从其小红书收藏及笔记中找到跟代言人或者脉动相关的内容",
	}, // 35
	{
		User:                  data_type.User{Name: "琦贝贝", UID: "6896622783"},
		CommitTime:            "04-22",
		ContentText:           "**一个由网友魔改的代言人比心的有点抽象的表情**",
		HandledSrc:            "代言人粉丝?",
		HandledTone:           "通过发送比心表情表达自己对代言人及其代言产品的支持",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 1,
		ReplyCount:            0,
		UserBackgroud:         "小红书头像似乎是代言人? 但没有从其小红书笔记中找到跟代言人或脉动有关的内容",
	}, // 36
	{
		User:                  data_type.User{Name: "琦贝贝", UID: "6896622783"},
		CommitTime:            "04-22",
		ContentText:           "@水煮美人鱼.🦄 @LNM金so喽 @小陈很好.🐰 @Jen鲶帝. @滚行不行 @.... @小调皮& @喵喵咪. @KBANB、 @STP🍟",
		HandledSrc:            "代言人粉丝?",
		HandledTone:           "可能是艾特朋友前来观看",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1; 这里参考了 36 所以使用 Happy Level 而不是 Middle Level
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "小红书头像似乎是代言人? 但没有从其小红书笔记中找到跟代言人或脉动有关的内容",
	}, // 37
	{
		User:                  data_type.User{Name: "小姩.（杂货铺）", UID: "6222476915"},
		CommitTime:            "04-22",
		ContentText:           "雨琦代言的必须多买点",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "表达对代言人及其代言产品脉动的肯定",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            3,
		UserBackgroud:         "女生、潮流博主、其小红书包含四篇跟代言人有关的小红书笔记",
	}, // 38
	{
		User:                  data_type.User{Name: "炸龟饼（欧丽达西）", UID: "4949088914"},
		CommitTime:            "04-22",
		ContentText:           "我买一瓶代言人可以送给我吗",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "表达了自己想要得到偶像的强烈情感?",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "女生；其小红书收藏出现了多个与代言人有关的内容",
	}, // 39
	{
		User:                  data_type.User{Name: "呃", UID: "108999828"},
		CommitTime:            "04-22",
		ContentText:           "**一个赞的表情**",
		HandledSrc:            "可能是路人, 也可能是对此感兴趣的路人, 也可能是追星者, 也可能是轻度粉丝",
		HandledTone:           "表达了对代言人或者对脉动产品的肯定?",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "其小红书收藏没有公开，并且其发布的小红书笔记没有与代言人或是脉动相关的笔记",
	}, // 40
	{
		User:                  data_type.User{Name: "🆔宋曲琦🍪💧", UID: "94107184090"},
		CommitTime:            "04-22",
		ContentText:           "**一张代言人穿着与视频同款的 Mizone 455 蓝色衣服在夜晚的海边的照片**",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "通过在评论区发送代言人的美照表达对代言人的喜爱",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 1,
		ReplyCount:            0,
		UserBackgroud:         "根据其发布的小红书笔记和其主页自述，其应该是一名代言人的粉丝",
	}, // 41
	{
		User:                  data_type.User{Name: "🆔宋曲琦🍪💧", UID: "94107184090"},
		CommitTime:            "04-22",
		ContentText:           "@刷刷的宁檬舒",
		HandledSrc:            "可能是邀请好友前往查看",
		HandledTone:           "通过在评论区发送代言人的美照表达对代言人的喜爱",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1; 参考自 41
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "根据其发布的小红书笔记和其主页自述，其应该是一名代言人的粉丝",
	}, // 42
	{
		User:                  data_type.User{Name: "一筐lemon（见过马丽版", UID: "5293674377"},
		CommitTime:            "04-22",
		ContentText:           "**一个赞的表情**",
		HandledSrc:            "可能是路人或者是了解代言人的路人，也可能是轻度粉丝",
		HandledTone:           "直切了当的表达自己对代言人或者脉动产品的肯定",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "处女座；其小红书收藏隐藏，且其发布的小红书笔记中没有找到跟代言人有关的内容",
	}, // 43
	{
		User:                  data_type.User{Name: "ko", UID: "106524201"},
		CommitTime:            "04-22",
		ContentText:           "太美了",
		HandledSrc:            "可能是路人",
		HandledTone:           "简洁明了的表达了自己认为代言人很美这件事",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "男生；其小红书没有发布任何笔记，且其收藏没有公开，主页也没有自述", // 终于看见男生了，这令人感动
	}, // 44
	{
		User:                  data_type.User{Name: "ko", UID: "106524201"},
		CommitTime:            "04-22",
		ContentText:           "好漂亮的代言人啊",
		HandledSrc:            "可能是路人",
		HandledTone:           "再度表达了自己对代言人的喜爱? (可能粉了吗?)",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "男生；其小红书没有发布任何笔记，且其收藏没有公开，主页也没有自述", // 终于看见男生了，这令人感动
	}, // 45
	{
		User:                  data_type.User{Name: "泥的炒猪血爱豪总", UID: "42195482097"},
		CommitTime:            "04-22",
		ContentText:           "给个链接吧",
		HandledSrc:            "可能是路人，也可能是轻度粉丝",
		HandledTone:           "可能是希望了解大学生篮球赛的报名链接，也可能只是单纯希望了解相关脉动产品的购买链接", // 真没人讨论大学生篮球赛啊，这令人抽象
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy,             // 用户有意向提出点击链接，说明其具有购买的欲望，应该还是 happy 的；1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "蛋仔派对玩家；没有发现其收藏或发布了跟代言人或者跟脉动有关的欸日",
	}, // 46
	{
		User:                  data_type.User{Name: "琦异果-", UID: "1077531248"},
		CommitTime:            "04-22",
		ContentText:           "宋雨琦做得好",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "表达对代言人代言脉动产品的肯定，也间接表达了对脉动产品的肯定",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 1,
		ReplyCount:            0,
		UserBackgroud:         "其小红书自述“top @宋雨琦_G-I-DLE”",
	}, // 47
	{
		User:                  data_type.User{Name: "柯北1998", UID: "3539468343"},
		CommitTime:            "04-22",
		ContentText:           "**一个赞的表情**",
		HandledSrc:            "?代言人轻度粉丝/路人/对追星有兴趣的人",
		HandledTone:           "通过一个赞表达出自己对代言人或者脉动产品的肯定",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "射手座；其小红书收藏隐藏，且对外发布的小红书笔记或其自述也没有跟脉动或者跟代言人有关的内容",
	}, // 48
	{
		User:                  data_type.User{Name: "宁雨jiejie", UID: "11594562257"},
		CommitTime:            "04-22",
		ContentText:           "雨琦美美哒",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "表达对代言人样貌的夸赞",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "其发布了很多个跟代言人有关的小红书笔记，并且其自述自己是代言人的担，应该是重度粉丝? (甚至头像似乎都是代言人)",
	}, // 49
	{
		User:                  data_type.User{Name: "会一直支持Yuqi的葡萄干", UID: "495688506"},
		CommitTime:            "04-21",
		ContentText:           "漂亮",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "表达自己对代言人的喜爱",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		PeopleStarThisComment: 0,
		ReplyCount:            0,
		UserBackgroud:         "其小红书笔记发布了很多跟代言人相关的小红书笔记，并且自述称自己“只爱代言人”",
	}, // 50
}

// Step 3.
//
// - Filter key words -
var ContentBKeyWords = []string{
	"必须买一箱", "跟宋雨琦结婚", "宋雨琦最美", "谁家雨琦最好看", "雨琦教我打篮球",
	"我已经喝上了", "雨琦帅死我了", "拍照印雨琦我天天喝", "打篮球喝了好好喝", "饮料只喝脉动",
	"赞", "瓶装设计太男风", "女孩也运动呀", "瓶子太工业风", "买 100 箱",
	"这个品牌在哪买", "美神", "哇", "比心", "多买几箱", "买脉动送雨琦",
	"厉害", "雨琦代言的多买点", "太美了", "好漂亮的代言人啊", "给个链接吧",
	"宋雨琦做得好", "雨琦美美哒", "漂亮",
}

// Setp 4.
//
// - Compute data and get simple result -
// TODO: WIP
