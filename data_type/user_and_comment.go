package data_type

// 单个用户的信息，只携带其名和 UID
type User struct {
	Name string
	UID  string
}

// 单个用户评论的信息
type Comment struct {
	User                  User
	CommitTime            string
	ContentText           string
	PeopleStarThisComment int

	HandledSrc            string
	HandledTone           string
	UserFeedbackToneLevel UserFeedbackToneLevel
	ReplyCount            int

	UserBackgroud string // 通过观察其笔记简要得出的
	AIComment     string // **最后** 把数据丢给 Deep seek 后其给出的分析 (较慢，让我们坐与放松，然后睡眠)
}
