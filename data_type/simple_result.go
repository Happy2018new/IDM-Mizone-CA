package data_type

import (
	"cmp"
	"fmt"
	"slices"
)

// 这是简单的分析结果，可以通过“人工”智能得出
type SimpleResult struct {
	// Hit 1
	AudiencePreferences []string // 受众偏好
	Treads              string   // 趋势
	UserAction          []string // 用户行为

	// Hit 3.3
	UserFeedbackTone []UserFeedbackToneLevel // 用户回馈语气 (排名); 此字段自动计算
	UserInterested   []string                // 用户感兴趣的内容

	// Hit 6
	// WhichTimeUserEngageIn               []int64  // 用户在一天中的什么时候参与度最高; Unix Time // 这玩意做不了，我可不想一个个 F12 看时间戳，这很抽象
	RelationBetweenHashtagAndEngagement []string // 哈希标签和参与度之间是否存在相关性

	// Hit 7
	ER float32 // 参与率；参与者÷观看者

	// Custom
	SimpleCount            int // 样本数量; 此字段自动计算
	VeryHappyPeopleCount   int // 感到非常快乐的人类; 此字段自动计算
	HappyPeopleCount       int // 感到快乐的人类; 此字段自动计算
	MiddlePeopleCount      int // 感到 OK 的人类; 此字段自动计算
	UnhappyPeopleCount     int // 感到不快乐的人类; 此字段自动计算
	VeryUnhappyPeopleCount int // 感到非常不愉快的人类; 此字段自动计算
}

// // 更复杂的工作还是让 AI 完成吧，这些数据让 AI 做；
// // 当然，这个部分的生成还是需要把 SimpleResult 丢给 AI
// type FurtherResult struct {
// 	// Hit 4
// 	Reach       int    // 看到内容的不同用户数量
// 	ReachDetail string // AI 关于 Reach 的阐述

// 	// Custom
// 	Suggestions string // AI 给出更多分析建议
// } // ignore

// 把 SimpleResult 和 FurtherResult 合起来，通过“人工”智能变成我们最终的分析报告
type FinalResult string

type tempStruct struct {
	ID      int
	Payload float32
}

// InitSimpleCountFromSlice 简要统计一下样本量和那些快乐悲伤人类的人数及数据。
// 其他字段还是得人工填写 :(
func InitSimpleCountFromSlice(origin SimpleResult, s []Comment) SimpleResult {
	result := origin
	result.SimpleCount = len(s)

	for _, value := range s {
		switch value.UserFeedbackToneLevel {
		case UserFeedbackToneVeryHappy:
			result.VeryHappyPeopleCount++
		case UserFeedbackToneHappy:
			result.HappyPeopleCount++
		case UserFeedbackToneMiddle:
			result.MiddlePeopleCount++
		case UserFeedbackToneUnhappy:
			result.UnhappyPeopleCount++
		case UserFeedbackToneVeryUnhappy:
			result.VeryUnhappyPeopleCount++
		default:
			panic(fmt.Sprintf("Unknown user deed back tone level %d", value.UserFeedbackToneLevel))
		}
	}

	temp := []tempStruct{
		{ID: UserFeedbackToneVeryHappy, Payload: float32(result.VeryHappyPeopleCount) / float32(result.SimpleCount)},
		{ID: UserFeedbackToneHappy, Payload: float32(result.HappyPeopleCount) / float32(result.SimpleCount)},
		{ID: UserFeedbackToneMiddle, Payload: float32(result.MiddlePeopleCount) / float32(result.SimpleCount)},
		{ID: UserFeedbackToneUnhappy, Payload: float32(result.UnhappyPeopleCount) / float32(result.SimpleCount)},
		{ID: UserFeedbackToneVeryUnhappy, Payload: float32(result.VeryUnhappyPeopleCount) / float32(result.SimpleCount)},
	}
	slices.SortStableFunc(temp, func(a tempStruct, b tempStruct) int {
		return cmp.Compare(b.Payload, a.Payload)
	})

	for _, value := range temp {
		result.UserFeedbackTone = append(result.UserFeedbackTone, UserFeedbackToneLevel(value.ID))
	}

	return result
}
