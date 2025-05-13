package data_type

type UserContentType int

const (
	UserContentTypeImage = iota
	UserContentTypeVideo
	UserContentTypeReels // 快拍
	UserContentTypeStory
)

type RepeatThingsType struct {
	EnumID  int
	Payload any
}

const (
	RepeatThingsTypeThemes = iota // 主题
	RepeatThingsTypeStyle         // 风格
	RepeatThingsTypeFormat        // 格式
)

type UserFeedbackToneLevel int

const (
	UserFeedbackToneVeryHappy = iota
	UserFeedbackToneHappy
	UserFeedbackToneMiddle
	UserFeedbackToneUnhappy
	UserFeedbackToneVeryUnhappy
)
