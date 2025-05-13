package data_type

type Trends string
type UserBehaviour string

// SingleContentGraph will not used, just to provide a whole mind map.
// ignore: _
type SingleContentGraph[T Trends | UserBehaviour, F float32 | int] struct {
	// Hit 1
	AudiencePreferences string // 受众偏好
	Engagement          F      // 参与度
	TUCombie            T      // 趋势 or 用户行为

	// Hit 2
	Hashtags         string // 标签
	Trends           string // 趋势
	SpecificAccounts any    // 特定账号
	Views            int    // 观看次数
	Likes            int    // 点赞
	CommentCount     int    // 评论数
	Shares           int    // 分享
	Stars            int    // 收藏

	// Hit 3.1
	UserContentType UserContentType // 用户正在创建哪些类型的内容？
	Keywords        []string        // 帖子中常用的关键词
	// Hit 3.2 - Video itself
	Filters          []string // 滤镜
	Captions         []string // 字幕
	VisualAesthetics []string // 视觉特效
	// Hit 3.3
	UserFeedbackTone []UserFeedbackToneLevel // 用户回馈语气
	UserInterested   []string                // 用户感兴趣的内容

	// Hit 4
	Reach       int // 看到内容的不同用户数量
	Impressions int // 内容被展示的总次数。

	// Hit 5
	FollowerTimeline         FollowerTimeline // 关注者数量随时间的变化
	BestContentToGetFollower []string         // 哪些内容最能推动关注者增长
	// TimeUserWatchEachVideo []int // ignore: due to we don't know
	// VideoFinishRate []float32    // ignore: due to we don't know
	// CTR []float 					// ignore: due to we don't know

	// Hit 6
	WhichLookBest                       []string // 哪些类型的内容表现最佳
	WhichTimeUserEngageIn               []int64  // 用户在一天中的什么时候参与度最高; Unix Time
	RelationBetweenHashtagAndEngagement []any    // 哈希标签和参与度之间是否存在相关性

	// Hit 7
	ER float32 // 参与率；参与者÷观看者
	// CR float32 // 转化率；行动者÷点击者 // ignore: due to we don't know
}
