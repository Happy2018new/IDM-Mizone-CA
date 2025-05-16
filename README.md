# IDM Mizone Content Analysis





# Catalogue
- [IDM Mizone Content Analysis](#idm-mizone-content-analysis)
- [Catalogue](#catalogue)
	- [Summary](#summary)
	- [Contents to be analyzed](#contents-to-be-analyzed)
	- [How to analyzed](#how-to-analyzed)
		- [Prepare](#prepare)
		- [For each video](#for-each-video)
		- [After all video](#after-all-video)
	- [LICENSE](#license)





## Summary
This is the repository that records the analysis process and results of some videos related to Mizone published on Xiaohongshu.





## Contents to be analyzed
The following contents are to be analyzed.

- [Contant A](https://www.xiaohongshu.com/explore/67e52aed000000000e00565a)





## How to analyzed
### Prepare
By referring to the courseware of Week 10 and the content of the workshop, we formulated some basic constants and data structures to represent the data to be collected later more efficiently

Helpful constants is listed below.

```go
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
```

```go
const (
	PublishContentTypeAdvertisement = iota
)

const (
	ContentGeneralFeedbackTrendsTypeTraditional = iota // 传统的 (广告)
)
```



### For each video
0. Filter basic data
1. Filter video data
2. See comments
3. Filter key words
4. Compute data and get simple result



### After all video
1. States all date
2. Combine to completely result






## LICENSE
Use **MIT** License for this project.