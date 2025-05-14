package hand_data_collected

import "IDM-Mizone-CA/data_type"

// ContantA & RAW (Delete unknown rune)
//
// 返场状态告急？不怕！
// 脉动时刻守护 @宋雨琦_G-I-DLE 的好状态
// 维C拉满，大口畅饮
// 满分活力回来💯尽情享受舞台💥
// 💬快来评论区晒出【生活中被脉动回来的时刻】
// ✌️抽5人各送代言人限定周边*1份！
// #宋雨琦脉动品牌代言人   #这一刻不容错过脉动回来   #让宋雨琦活力返场的状态救星

// Step 0.
//
// - Publisher table -
var CommentAPublisherTable = data_type.ContentInfo{
	Publisher: data_type.Publisher{
		Name:   "Mizone Chinese",
		Detail: "Official Account",
		URL:    "https://www.xiaohongshu.com/user/profile/5e4d669f0000000001000a36",
	},
	PublishTime:        "2025-03-28",
	PublishContentType: data_type.PublishContentTypeAdvertisement,
	PublishDetail: []string{
		"Latest one",
		// TODO: Finish WIP
		"Invite @宋雨琦_G-I-DLE in it (Super star; User Data? <WIP>; https://www.xiaohongshu.com/user/profile/5cc10d72000000001200f6b8)",
	},
	Slogan: []string{
		"维C拉满，大口畅饮", "这一刻不容错过，脉动回来", "快来评论区晒出【生活中被脉动回来的时刻】",
	},
	HowTo: []string{"抽5人各送代言人限定周边*1份！"},
}

// Step 1.
//
// - Video table (Hit 2) -
var ContentAVideoTable = data_type.ContentGeneralFeedback{
	HashTag: []string{
		"#宋雨琦脉动品牌代言人", "#这一刻不容错过脉动回来", "#让宋雨琦活力返场的状态救星",
	},
	TrendsType:       data_type.ContentGeneralFeedbackTrendsTypeTraditional,
	TrendsDetail:     "同样的口号和宣传模式，即“脉动回来”",
	SpecificAccounts: []data_type.Publisher{CommentAPublisherTable.Publisher},
	KnownViews:       false,
	Views:            0,
	Likes:            23000,
	CommentCounts:    1188,
	Stars:            5265,
	StatisticalTime:  "2025-05-13 19:09",
}

// - Hit 3 -
var ContentAHit3 = data_type.ContentInv{
	Filters:         []string{"开头有一个以脉动瓶子的窗户开场的滤镜特效?"},
	UserContentType: data_type.UserContentTypeReels,
	Captions: []string{
		"脉动风格的字幕 (官方制作)", "大、清晰、容易认出这是脉动字幕",
	},
	VisualAesthetics: []string{
		"拍摄是代言人在舞台上进行的，并且舞台以蓝色基调，给人以脉动的直接联想",
		"00:06: 脉动救我 (于是一个脉动瓶子从天而降，然后用户得到了脉动，有水花特效)",
		"00:09: 大口畅饮 (大型水花特效及声音)",
		"00:11~00:12: 脉动回来 (人跳起来，伴随着水花出现，以及屏幕两边的烟花特效)",
	},
	Impressions: "脉动贯穿全文",
}

// Step 2.
//
// - Comment Table -
var ContentACommentTable = []data_type.Comment{
	{
		User:                  data_type.User{Name: "宋吃吃（雨琦天天开心）", UID: "674298046"},
		CommitTime:            "03-28",
		ContentText:           "尝尝青草味的 代言人好美",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "想要代言人限定周边产品 (Also reply, 5 people want, total 7 reply)",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 体现：支持周边
		PeopleStarThisComment: 366,
		ReplyCount:            8,
		UserBackgroud:         "代言人重度粉丝, 几乎每个小红书笔记都是她, 最后一个动态发布于 05/09",
	}, // 1
	{
		User:                  data_type.User{Name: "我担雨琦天下第一", UID: "26274516098"},
		CommitTime:            "03-28",
		ContentText:           "官方太用心了。雨琦造型好美。",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "夸赞代言人很美 :)",
		UserFeedbackToneLevel: data_type.UserFeedbackToneVeryHappy, // 0
		PeopleStarThisComment: 139,
		ReplyCount:            6,
		UserBackgroud:         "共三个小红书笔记，都与偶像追星相关，不确定是不是代言人，但主页评论“主推宋雨琦(代言人)”",
	}, // 2
	{
		User:                  data_type.User{Name: "Chloe.辰", UID: "6258382517"},
		CommitTime:            "04-05",
		ContentText:           "买脉动送不送代言人",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "调侃：希望买脉动可以得到代言人",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 0
		PeopleStarThisComment: 143,
		ReplyCount:            11,
		UserBackgroud:         "看起来是普通用户，但个人评论有写喜欢代言人",
	}, // 3
	{
		User:                  data_type.User{Name: "我是抽象派", UID: "1550740413"},
		CommitTime:            "03-28",
		ContentText:           "我买一瓶脉动，能把雨琦买过来吗？",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "调侃：希望买脉动可以得到代言人",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 0
		PeopleStarThisComment: 78,
		ReplyCount:            1,
		UserBackgroud:         "普通用户?",
	}, // 4
	{
		User:                  data_type.User{Name: "woogie", UID: "1104556877"},
		CommitTime:            "03-28",
		ContentText:           "支持脉动品牌代言人宋雨琦",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "该评论的回复提到买脉动“送徽章没有小卡”，同时讨论了口味，并且讨论他们买了，同时觉得口味还可以", // 参与度 +1?
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy,                   // 0
		PeopleStarThisComment: 65,
		ReplyCount:            4,
		UserBackgroud:         "他的主页小红书笔记有较多关于代言人的，应该属于比较了解代言人的粉丝，最后一个相关的帖子发布在 04/07",
	}, // 5
	{
		User:                  data_type.User{Name: "小小", UID: "1118430808"},
		CommitTime:            "04-12",
		ContentText:           "我推。",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "评论区在讨论应该用“推”还是“担”来称呼三次元追星行为",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		PeopleStarThisComment: 0,
		ReplyCount:            6,
		UserBackgroud:         "看起来小红书主页是空的，个人简介也没有?",
	}, // 6
	{
		User:                  data_type.User{Name: "阿辞i锐锐", UID: "42949808261"},
		CommitTime:            "04-15",
		ContentText:           "第一眼广告，不看。第二眼雨琦，直接看完。",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "另外一个人回复了他，表示赞同",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		PeopleStarThisComment: 5,
		ReplyCount:            1,
		UserBackgroud:         "从主页来看是个学生，然后里面有很多关于代言人的小红书笔记，最后一个发布于 02-20",
	}, // 7
	{
		User:                  data_type.User{Name: "魚思年✨", UID: "6284088875"},
		CommitTime:            "04-13",
		ContentText:           "**一个关于和代言人结婚的表情包**",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "评论区在讨论谁应该跟代言人结婚这件事",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		PeopleStarThisComment: 1,
		ReplyCount:            2,
		UserBackgroud:         "主页没有发现与代言人相关的小红书笔记，个人简介也没写，但头像有点像?",
	}, // 8
	{
		User:                  data_type.User{Name: "一只芷苓吖（薯）", UID: "1167995316"},
		CommitTime:            "04-10",
		ContentText:           "买脉动送不送代言人呀",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "评论区在同问",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		PeopleStarThisComment: 3,
		ReplyCount:            1,
		UserBackgroud:         "主页没有发现与代言人相关的小红书笔记，个人简介也没写",
	}, // 9
	{
		User:                  data_type.User{Name: "小红薯6128DF4C", UID: "4201662903"},
		CommitTime:            "04-21",
		ContentText:           "我怎么感觉明星长的全都一样？",
		HandledSrc:            "路人（小学生）",
		HandledTone:           "评论区回复“因为你不够了解她们(狗头)”",
		UserFeedbackToneLevel: data_type.UserFeedbackToneUnhappy, // 3
		PeopleStarThisComment: 0,
		ReplyCount:            1,
		UserBackgroud:         "路人、小学生、不参与追星这件事",
	}, // 10
	{
		User:                  data_type.User{Name: "冷酷奇犽", UID: "798268617"},
		CommitTime:            "04-20",
		ContentText:           "又渴又累，状态缩水",
		HandledSrc:            "路人?",
		HandledTone:           "评论看上去是在跟脉动唱反调，给人一种抽象感；评论区的回复(回复“6”)显示人们觉得这个评论很抽象",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		PeopleStarThisComment: 1,
		ReplyCount:            1,
		UserBackgroud:         "看起来是路人，并且是学生",
	}, // 11
	{
		User:                  data_type.User{Name: "寻椿莫迟", UID: "Breakfast250"},
		CommitTime:            "04-20",
		ContentText:           "逼我买脉动？宋雨琦你不要小看你的魅力了",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "评论区的回复(回复“6”)显示人们觉得这个评论很抽象",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		ReplyCount:            0,
		PeopleStarThisComment: 2,
		UserBackgroud:         "二次元并且似乎也追星",
	}, // 12
	{
		User:                  data_type.User{Name: "林虞兮", UID: "95456646562"},
		CommitTime:            "04-19",
		ContentText:           "谁懂 我现在就在喝",
		HandledSrc:            "路人? 未知?",
		HandledTone:           "用户的评论让人觉得他正在喝脉动的产品就刷到了视频，很巧合又很激动", // 参与+1
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy,    // 1
		ReplyCount:            0,
		PeopleStarThisComment: 1,
		UserBackgroud:         "主页无法很好体现他是否是代言人的粉丝",
	}, // 13
	{
		User:                  data_type.User{Name: "又不见春", UID: "5421993395"},
		CommitTime:            "04-19",
		ContentText:           "哇~",
		HandledSrc:            "路人? 未知?",
		HandledTone:           "感叹的语气看上去是比较肯定脉动的产品(或者代言人)",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		ReplyCount:            0,
		PeopleStarThisComment: 0,
		UserBackgroud:         "主页无法很好体现他是否是代言人的粉丝",
	}, // 14
	{
		User:                  data_type.User{Name: "是元宝呀！", UID: "225128517"},
		CommitTime:            "04-19",
		ContentText:           "宋雨旑",
		HandledSrc:            "路人? 未知?",
		HandledTone:           "评论似乎只体现出这个用户认识或者了解一些代言人，但似乎算不上代言人的粉丝",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		ReplyCount:            0,
		PeopleStarThisComment: 1,
		UserBackgroud:         "主页看上去是这个用户最近养了鹦鹉，但跟脉动和代言人没有什么关联",
	}, // 15
	{
		User:                  data_type.User{Name: "晚姩爱奇迹🐞", UID: "26280436451"},
		CommitTime:            "04-19",
		ContentText:           "雨琦♥",
		HandledSrc:            "可能是因为代言人才看的视频",
		HandledTone:           "评论体现出这个人比较喜欢代言人?",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		ReplyCount:            0,
		PeopleStarThisComment: 0,
		UserBackgroud:         "主页看上去单推的是一个动漫，并且似乎有一定的追星，但没有发现跟代言人有关的内容",
	}, // 16
	{
		User:                  data_type.User{Name: "凤梨", UID: "26791228499"},
		CommitTime:            "04-19",
		ContentText:           "买脉动能送雨琦吗",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "表达很喜爱代言人，并且其最后一个跟代言人有关的小红书动态发布在 05-05",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		ReplyCount:            0,
		PeopleStarThisComment: 2,
		UserBackgroud:         "15岁学生，追星，平常在小红书上发布绘画、漫画",
	}, // 17
	{
		User:                  data_type.User{Name: "凤梨", UID: "26791228499"},
		CommitTime:            "04-19",
		ContentText:           "@邪恶蜿延 @人机",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "应该是艾特好朋友来看代言人的视频",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		ReplyCount:            0,
		PeopleStarThisComment: 2,
		UserBackgroud:         "15岁学生，追星，平常在小红书上发布绘画、漫画",
	}, // 18
	{
		User:                  data_type.User{Name: "妮茶延面", UID: "1029666874"},
		CommitTime:            "04-18",
		ContentText:           "**一个泪流成河的表情包**",
		HandledSrc:            "可能是被小红书推流看见的视频",
		HandledTone:           "泪目",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		ReplyCount:            0,
		PeopleStarThisComment: 0,
		UserBackgroud:         "是女生，可能感兴趣的内容包含追星",
	}, // 19
	{
		User:                  data_type.User{Name: "浅^月", UID: "7401146856"},
		CommitTime:            "04-17",
		ContentText:           "我就在喝那个桃子味的",
		HandledSrc:            "代言人的粉丝",
		HandledTone:           "表达自己已经用喝上了脉动，并且还在看视频",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		ReplyCount:            0,
		PeopleStarThisComment: 1,
		UserBackgroud:         "追星且主推代言人，海湾第五人格，女团是t t V韩娱团",
	}, // 20
	{
		User:                  data_type.User{Name: "蛋仔嵛苳", UID: "520FSY1314"},
		CommitTime:            "04-17",
		ContentText:           "@呆呆.🌷",
		HandledSrc:            "未知",
		HandledTone:           "可能是艾特朋友来看这个帖子",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		ReplyCount:            0,
		PeopleStarThisComment: 0,
		UserBackgroud:         "小红书没有太多笔记，个人主页也没有介绍，只知道是女生并且看头像是二次元",
	}, // 21
	{
		User:                  data_type.User{Name: "Asa存湖", UID: "11675815696"},
		CommitTime:            "04-17",
		ContentText:           "@殇忆 @李^Heler^森筱🥑 @辞忧 我还等着yy给咱们买呢",
		HandledSrc:            "可能是了解代言人然后被平台推流了",
		HandledTone:           "看起来像是等 yy 这个人给他买脉动?",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		ReplyCount:            0,
		PeopleStarThisComment: 1,
		UserBackgroud:         "小红书主页多为热门歌曲相关的笔记，并且还包含一些追星的",
	}, // 22
	{
		User:                  data_type.User{Name: "🙄💅", UID: "🙄💅"},
		CommitTime:            "04-16",
		ContentText:           "雨琦",
		HandledSrc:            "可能是了解代言人然后被平台推流了",
		HandledTone:           "表达出自己了解代言人然后恰好看到视频了?",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		ReplyCount:            0,
		PeopleStarThisComment: 0,
		UserBackgroud:         "小红书主页没有笔记，也没有个人介绍，但可以了解是女生，并且看头像可能是抽象网络带师",
	}, // 23
	{
		User:                  data_type.User{Name: "LihKiMuu.", UID: "302050736"},
		CommitTime:            "04-16",
		ContentText:           "一群官方在干嘛",
		HandledSrc:            "路人?",
		HandledTone:           "对评论区的评论表示不解",
		UserFeedbackToneLevel: data_type.UserFeedbackToneUnhappy, // 3s
		ReplyCount:            0,
		PeopleStarThisComment: 0,
		UserBackgroud:         "小学生",
	}, // 24
	{
		User:                  data_type.User{Name: "林木沐", UID: "9530314699"},
		CommitTime:            "04-15",
		ContentText:           "**一个泪流成河的表情包**",                // 同 19
		HandledSrc:            "可能是被小红书推流看见的视频",                // 同 19
		HandledTone:           "泪目",                            // 同 19
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 同 19
		ReplyCount:            0,
		PeopleStarThisComment: 0,
		UserBackgroud:         "是女生，根据小红书主页看起来她追星",
	}, // 25
	{
		User:                  data_type.User{Name: "小凶手？", UID: "95384976356"},
		CommitTime:            "04-15",
		ContentText:           "我要买六箱啊，买脉动。之后呢，里面有宋雨琦。",
		HandledSrc:            "代言人粉丝?",
		HandledTone:           "表达出要买很多脉动以显示自己很爱代言人",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		ReplyCount:            0,
		PeopleStarThisComment: 0,
		UserBackgroud:         "根据小红书主页推断其是蛋仔派对的玩家，但没有发现跟代言人和脉动相关联的小红书笔记",
	}, // 26
	{
		User:                  data_type.User{Name: "小凶手？", UID: "95384976356"},
		CommitTime:            "04-15",
		ContentText:           "直接叫他宋雨琦来我家。",
		HandledSrc:            "代言人粉丝?",
		HandledTone:           "甚至表达出希望大主播(代言人)去他家的意思",             // 说实话我觉得这个很抽象，因为这一点也没有表达对代言人的尊重
		UserFeedbackToneLevel: data_type.UserFeedbackToneVeryHappy, // 0
		ReplyCount:            0,
		PeopleStarThisComment: 1,
		UserBackgroud:         "根据小红书主页推断其是蛋仔派对的玩家，但没有发现跟代言人和脉动相关联的小红书笔记",
	}, // 27
	{
		User:                  data_type.User{Name: "太好了，是雨琦我们有救了", UID: "1637604980"},
		CommitTime:            "04-14",
		ContentText:           "太好了，是雨琦我们有救了",
		HandledSrc:            "代言人粉丝?",
		HandledTone:           "表达对代言人的喜爱?",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		ReplyCount:            0,
		PeopleStarThisComment: 0,
		UserBackgroud:         "蛋仔派对的玩家，且追星，但似乎没有跟脉动或代言人相关的小红书笔记",
	}, // 28
	{
		User:                  data_type.User{Name: "🌷花间梦🌷", UID: "95388452466"},
		CommitTime:            "04-14",
		ContentText:           "我家琦宝真好看哦",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "夸赞代言人很美，表达对其的喜爱",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		ReplyCount:            0,
		PeopleStarThisComment: 1,
		UserBackgroud:         "女生，爱好可爱清新的东西，但似乎没有跟脉动或代言人相关的小红书笔记",
	}, // 29
	{
		User:                  data_type.User{Name: "小绵吖", UID: "234941091"},
		CommitTime:            "04-14",
		ContentText:           "我们学校就有",
		HandledSrc:            "路人?",
		HandledTone:           "看脉动视频时恰好想起自己学校有卖这款脉动，于是发表了评论?",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		ReplyCount:            0,
		PeopleStarThisComment: 0,
		UserBackgroud:         "玩王者荣耀，但没有看到跟代言人或脉动相关的小红书笔记",
	}, // 30
	{
		User:                  data_type.User{Name: "woonini", UID: "cyh090920121234"},
		CommitTime:            "04-14",
		ContentText:           "@披萨🍕 @ID.总裁荞麦面 @终焉",
		HandledSrc:            "路人?",
		HandledTone:           "可能是艾特朋友来看视频",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		ReplyCount:            0,
		PeopleStarThisComment: 0,
		UserBackgroud:         "处女座，根据小红书主页似乎追星，但未发现与代言人或脉动相关的小红书笔记",
	}, // 31
	{
		User:                  data_type.User{Name: "曲琦味的鹿茸🍪", UID: "265465353"},
		CommitTime:            "04-14",
		ContentText:           "宋雨琦！！！+一个关于代言人的大头照",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "可能是艾特朋友来看视频",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		ReplyCount:            0,
		PeopleStarThisComment: 1,
		UserBackgroud:         "小红书没有笔记和个人介绍，但头像是代言人的图，推测是代言人的粉丝",
	}, // 32
	{
		User:                  data_type.User{Name: "+の、芷欣.🔍🍄💕", UID: "Kuinai5201314n"},
		CommitTime:            "04-14",
		ContentText:           "笑死我了哈我平板能偷窥刚拿一个脉动刷下一个视频，就是这个脉动哈",
		HandledSrc:            "路人 (也可能是来自联合推流)",
		HandledTone:           "表达自己刚买脉动就刷到脉动视频这件事",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		ReplyCount:            0,
		PeopleStarThisComment: 1,
		UserBackgroud:         "不太清楚其背景信息，但其小红书个人简介其为金牛座",
	}, // 33
	{
		User:                  data_type.User{Name: "白鹿是我老婆🌷", UID: "26963428025"},
		CommitTime:            "04-14",
		ContentText:           "第1次看广告都这么津津有味",
		HandledSrc:            "可能是代言人粉丝?",
		HandledTone:           "可能是表达自己本来是广告不想看的，但是由于是代言人的广告，所以看得很起劲?",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		ReplyCount:            0,
		PeopleStarThisComment: 1,
		UserBackgroud:         "爱玩原神🌷蛋仔派对🌷有快手账号💦可以深交👌粉白鹿，敖瑞鹏😊",
	}, // 34
	{
		User:                  data_type.User{Name: "落坞", UID: "9452379639"},
		CommitTime:            "04-14",
		ContentText:           "太爱脉动了",
		HandledSrc:            "脉动粉丝", // 终于看到了一个脉动粉丝？这令人抽象
		HandledTone:           "可能是表达自己本来是广告不想看的，但是由于是代言人的广告，所以看得很起劲?",
		UserFeedbackToneLevel: data_type.UserFeedbackToneVeryHappy, // 1
		ReplyCount:            0,
		PeopleStarThisComment: 0,
		UserBackgroud:         "男生 (脉动的目标用户之一), 大学生(安徽大学龙河校区), 二次元(根据头像)",
	}, // 35
	{
		User:                  data_type.User{Name: "章妤小丸子", UID: "95632856976"},
		CommitTime:            "04-14",
		ContentText:           "@小乖不乖",
		HandledSrc:            "平台推流 (路人)",
		HandledTone:           "推荐给朋友看，然后他朋友回复了，并表示很喜欢代言人",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		ReplyCount:            1,
		PeopleStarThisComment: 0,
		UserBackgroud:         "喜欢看漫画和番剧",
	}, // 36
	{
		User:                  data_type.User{Name: "轩轩", UID: "5391819883"},
		CommitTime:            "04-12",
		ContentText:           "我同学看到这条广告直接哭了。上课看开心锤锤的时候。",
		HandledSrc:            "可能是平台推流 (路人?)",
		HandledTone:           "表达即便是非粉丝，其身边也有一些代言人的粉丝",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		ReplyCount:            0,
		PeopleStarThisComment: 0,
		UserBackgroud:         "普通人，但看起来喜欢鹦鹉",
	}, // 37
	{
		User:                  data_type.User{Name: "杂鱼.", UID: "95654861441"},
		CommitTime:            "04-13",
		ContentText:           "感覺可以買100箱喔",
		HandledSrc:            "平台推流 (路人)",
		HandledTone:           "表达自己很喜欢喝脉动，一次性可以购入 100 箱?",     // 喜欢脉动 +1
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		ReplyCount:            0,
		PeopleStarThisComment: 1,
		UserBackgroud:         "未知, 但看起来来自中国台湾",
	}, // 38
	{
		User:                  data_type.User{Name: "雨陪我哭泣、", UID: "5200901qjn"},
		CommitTime:            "04-13",
		ContentText:           "我宣布麦冬将成为我的新宠。",
		HandledSrc:            "平台推流 (路人)",
		HandledTone:           "表达自己很喜欢喝脉动?",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		ReplyCount:            0,
		PeopleStarThisComment: 0,
		UserBackgroud:         "女生, 小学生~高中生",
	}, // 39
	{
		User:                  data_type.User{Name: "曲奇饼干《^~^》", UID: "2130144372"},
		CommitTime:            "04-13",
		ContentText:           "我觉得可以叫我妈买两箱",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "由于是代言人，想要买脉动支持一下",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		ReplyCount:            0,
		PeopleStarThisComment: 1,
		UserBackgroud:         "根据小红书头像，似乎这个用户应该追星，然后她的收藏中包含一些与代言人有关的内容，但未发现与脉动有关的内容",
	}, // 40
	{
		User:                  data_type.User{Name: "赵鞠婉-小云朵", UID: "2686260087"},
		CommitTime:            "04-13",
		ContentText:           "如果是雨琦打的广告哈，那我就必须喝",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "表示大力支持代言人",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		ReplyCount:            0,
		PeopleStarThisComment: 1,
		UserBackgroud:         "蛋仔派对玩家，同时似乎一定程度上追星",
	}, // 41
	{
		User:                  data_type.User{Name: "知珉爱吃柚子", UID: "8848265310"},
		CommitTime:            "04-13",
		ContentText:           "@Yu^qiii🍪 @Yulu^siqi✨🍪 @yu^qiii的小曲奇",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "可能是艾特好朋友来看视频",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		ReplyCount:            0,
		PeopleStarThisComment: 0,
		UserBackgroud:         "根据小红书主页，似乎热爱网红，可能追星，但没有发现跟代言人或脉动具有强烈关联的内容",
	}, // 42
	{
		User:                  data_type.User{Name: "💬暴富小霖.", UID: "LIN0621"},
		CommitTime:            "04-13",
		ContentText:           "@宋雨琦_G-I-DLE 你真好看",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "十分直白的表示自己对代言人的喜爱",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		ReplyCount:            0,
		PeopleStarThisComment: 0,
		UserBackgroud:         "第五人格玩家，小马宝莉爱好者，最近在分享一些学习日常和工具，自称“手工”博主",
	}, // 43
	{
		User:                  data_type.User{Name: "曲琦小饼干.", UID: "曲琦小饼干"},
		CommitTime:            "04-13",
		ContentText:           "脉动，你等我，我去抢1000箱",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "通过希望大量购入脉动来表示对代言人的喜爱，并且回复她的用户也持有一致的观点，并且也希望这么做(买很多脉动)",
		UserFeedbackToneLevel: data_type.UserFeedbackToneVeryHappy, // 0
		ReplyCount:            1,
		PeopleStarThisComment: 1,
		UserBackgroud:         "根据小红书主页，其包含四个跟代言人有关的内容，并且最新发布的笔记是关于代言人的",
	}, // 44
	{
		User:                  data_type.User{Name: "滧酬🍡（重回巅峰）.", UID: "198789184"},
		CommitTime:            "04-13",
		ContentText:           "雨琦？",
		HandledSrc:            "未知，可能是平台推流来的路人",
		HandledTone:           "似乎是认出了代言人是哪位，但不太清楚她进一步要表达什么意思",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		ReplyCount:            0,
		PeopleStarThisComment: 0,
		UserBackgroud:         "看弹丸论破，但没有发现与代言人或脉动有关的小红书笔记",
	}, // 45
	{
		User:                  data_type.User{Name: "苏瑶", UID: "6915684671"},
		CommitTime:            "04-13",
		ContentText:           "哇偶",
		HandledSrc:            "可能是代言人的粉丝，然后被平台推流了",
		HandledTone:           "表达了对代言人(脉动)的赞叹?",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		ReplyCount:            0,
		PeopleStarThisComment: 0,
		UserBackgroud:         "其收藏有一个关于代言人的小红书笔记，不太清楚其是否真的是代言人的粉丝",
	}, // 46
	{
		User:                  data_type.User{Name: "软小棉", UID: "9460430939"},
		CommitTime:            "04-13",
		ContentText:           "我们在学校看电影的时候广告也出来这个了，那时候我看到宋雨琦激动了好久",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "真挚的表达了自己对代言人的喜爱",
		UserFeedbackToneLevel: data_type.UserFeedbackToneVeryHappy, // 0
		ReplyCount:            0,
		PeopleStarThisComment: 0,
		UserBackgroud:         "蛋仔派对玩家，似乎是小学生，然后喜欢看小马宝莉",
	}, // 47
	{
		User:                  data_type.User{Name: "一只聪明的小青蛙✨", UID: "1563231143"},
		CommitTime:            "04-13",
		ContentText:           "@Prawn🍪（已被作业千刀万剐） @一只聪明的小蝌蚪🍟",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "推荐好友来查看视频",
		UserFeedbackToneLevel: data_type.UserFeedbackToneHappy, // 1
		ReplyCount:            0,
		PeopleStarThisComment: 1,
		UserBackgroud:         "女生，追星，并且包含代言人",
	}, // 48
	{
		User:                  data_type.User{Name: "困困鱼.冰清清", UID: "7136023146"},
		CommitTime:            "04-13",
		ContentText:           "tffgtghygdddxdc", // 这令人感到抽象 :(
		HandledSrc:            "未知，可能是路人也可能是粉丝推流",
		HandledTone:           "未知，因为使用了非人类可读语言",
		UserFeedbackToneLevel: data_type.UserFeedbackToneMiddle, // 2
		ReplyCount:            0,
		PeopleStarThisComment: 0,
		UserBackgroud:         "小学生? 未知更多信息",
	}, // 49
	{
		User:                  data_type.User{Name: "流浪", UID: "6023024472"},
		CommitTime:            "04-13",
		ContentText:           "我去！雨琦！曲奇们！快来！",
		HandledSrc:            "代言人粉丝",
		HandledTone:           "表达对代言人的热爱，并希望天下所有人都了解代言人并且观看这个视频",
		UserFeedbackToneLevel: data_type.UserFeedbackToneVeryHappy, // 0
		ReplyCount:            0,
		PeopleStarThisComment: 1,
		UserBackgroud:         "原神玩家，似乎喜欢流浪者，根据其小红书自述知晓其喜欢代言人，但没有得到更多的跟脉动有关的内容",
	}, // 50
}

// Step 3.
//
// - Filter key words -
var ContentAKeyWords = []string{
	"代言人好美", "雨琦造型好美", "送不送代言人", "买雨琦", "支持雨琦",
	"雨琦", "我现在就在喝", "哇~", "雨琦♥", "我就在喝",
	"我要买六箱", "是雨琦我们有救了", "琦宝真好看", "我们学校就有",
	"看广告都这么津津有味", "太爱脉动了", "泪目", "買100箱", "哇偶",
	"麦冬将成为我的新宠", "叫我妈买两箱", "好看", "抢1000箱", "看到宋雨琦激动了好久",
}

// Setp 4.
//
// - Compute data and get simple result -
var ContentASimpleResult = data_type.InitSimpleCountFromSlice(
	data_type.SimpleResult{
		AudiencePreferences: []string{
			"从小学生到大学生都出现在样本中，说明年龄适应性广泛",
			"用户中存在一些游戏玩家，但多为蛋仔派对玩家",
			"用户群体几乎都是女性，只有少量是男性",
		},
		Treads: "用户多趋向于代言人(宋雨琦)的粉丝而观看脉动的广告",
		UserAction: []string{
			"希望大量购买脉动产品来表示对代言人(宋雨琦)的支持",
			"大部分用户夸赞代言人(宋雨琦)很美，并且是因为她而来观看视频",
			"部分用户对代言人(宋雨琦)代理脉动进行宣传感到惊叹",
		},
		UserInterested: []string{
			"用户大多表达对代言人(宋雨琦)的喜爱，少量用户直接表达自己对脉动的喜爱，并且样本中几乎很多人都是因为代言人而观看视频",
		},
		RelationBetweenHashtagAndEngagement: []string{
			"“#宋雨琦脉动品牌代言人” 跟样本中用户的实际参与基本吻合，样本中大部分参与的用户都是因为该明星而观看视频",
			"“#这一刻不容错过脉动回来” 跟评论区中用户的交谈内容基本没有吻合，但该标签体现了脉动品牌及要宣传的产品，从影响度来看，是成功的",
			"“#让宋雨琦活力返场的状态救星” 也是跟样本中用户的讨论契合，并且部分用户的评论直接出现了“我们有救了”的字样",
		},
		ER: float32(ContentAVideoTable.CommentCounts) / float32(ContentAVideoTable.Likes),
	},
	ContentACommentTable,
)
