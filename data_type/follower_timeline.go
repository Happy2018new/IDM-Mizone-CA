package data_type

// 关注者随时间的变化
type FollowerTimeline []FollowerTimePoint

// 单个时间点的关注者数据
type FollowerTimePoint struct {
	FollowerCount int
	UnixTime      int
}
