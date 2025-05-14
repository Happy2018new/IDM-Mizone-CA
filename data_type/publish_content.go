package data_type

const (
	PublishContentTypeAdvertisement = iota
)

// 发布者
type Publisher struct {
	Name   string
	URL    string
	Detail string
}

// 发布的内容的基本信息
type ContentInfo struct {
	Publisher   Publisher
	PublishTime string

	PublishContentType int
	PublishDetail      []string

	Slogan []string
	HowTo  []string
}

const (
	ContentGeneralFeedbackTrendsTypeTraditional         = iota // 传统的 (广告)
	ContentGeneralFeedbackTrendsTypeSuperStarSuggestion        // 明星推荐
)

// 内容发布后的一些基本数据
type ContentGeneralFeedback struct {
	HashTag []string

	TrendsType       int
	TrendsDetail     string
	SpecificAccounts []Publisher

	KnownViews      bool
	Views           int
	Likes           int
	CommentCounts   int
	Stars           int
	StatisticalTime string
}

// 对发布内容的逆向，
// 以了解其是如何制作的
type ContentInv struct {
	UserContentType  int
	Filters          []string // 滤镜
	Captions         []string // 字幕
	VisualAesthetics []string // 视觉特效
	Impressions      string   // 内容被展示的总次数
}
