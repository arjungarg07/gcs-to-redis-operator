package dto

type PostStaticRedisMsetMsg struct {
	Key string
	Val []byte
}

type PostStaticFeatures struct {
	PostID            string  `json:"postId,omitempty"`
	TagID             string  `json:"tagId,omitempty"`
	CreatorID         string  `json:"creatorId,omitempty"`
	Language          string  `json:"language,omitempty"`
	CreatedOn         string  `json:"createdOn,omitempty"`
	LOTopic           string  `json:"LO_topic,omitempty"`
	L1Topic           string  `json:"L1_topic,omitempty"`
	L2Topic           string  `json:"L2_topic,omitempty"`
	CUL1Topic         string  `json:"CU_L1_topic,omitempty"`
	Duration          float64 `json:"duration,omitempty"`
	CreatorIP         string  `json:"creatorIp,omitempty"`
	CreatorCity       string  `json:"creatorCity,omitempty"`
	CreatorState      string  `json:"creatorState,omitempty"`
	CreatorGender     string  `json:"creatorGender,omitempty"`
	CreatorBadge      string  `json:"creatorBadge,omitempty"`
	CreatorType       string  `json:"creatorType,omitempty"`
	PredictedProb     float64 `json:"predictedProb,omitempty"`
	PredictedTopic    string  `json:"predictedTopic,omitempty"`
	ContentType       string  `json:"contentType,omitempty"`
	TagGenre          string  `json:"tagGenre,omitempty"`
	Badge             string  `json:"badge,omitempty"`
	L0Taxonomy        string  `json:"L0_taxonomy,omitempty"`
	L1Taxonomy        string  `json:"L1_taxonomy,omitempty"`
	L2Taxonomy        string  `json:"L2_taxonomy,omitempty"`
	L3Taxonomy        string  `json:"L3_taxonomy,omitempty"`
	L4Taxonomy        string  `json:"L4_taxonomy,omitempty"`
	CommentOff        string  `json:"commentOff,omitempty"`
	PostShareDisabled string  `json:"postShareDisabled,omitempty"`
	Height            string  `json:"height,omitempty"`
	Width             string  `json:"width,omitempty"`
	HybridTopic       string  `json:"hybridTopic,omitempty"`
}
