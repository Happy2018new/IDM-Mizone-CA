package data_type

type UserContentType int

const (
	UserContentTypeImage = iota
	UserContentTypeVideo
	UserContentTypeReels // 快拍
	UserContentTypeStory
)

const (
	RepeatThingsTypeThemes = iota // 主题
	RepeatThingsTypeStyle         // 风格
	RepeatThingsTypeFormat        // 格式
)

type RepeatThingsType struct {
	EnumID  int
	Payload any
}

type UserFeedbackToneLevel int

const (
	UserFeedbackToneVeryHappy = iota
	UserFeedbackToneHappy
	UserFeedbackToneMiddle
	UserFeedbackToneUnhappy
	UserFeedbackToneVeryUnhappy
)
