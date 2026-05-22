package activityModels

type ActivityInfo struct {
	Id          int32   `bson:"id" json:"id"`
	Index       int32   `bson:"index" json:"index"`
	TagId       int16   `bson:"tag_id" json:"tag_id"`
	Type        int8    `bson:"type" json:"type"`
	Banner      string  `bson:"banner" json:"banner"`
	Icon        string  `bson:"icon" json:"icon"`
	Title       string  `bson:"title" json:"title"`
	Desc        string  `bson:"desc" json:"desc"`
	PreviewTime int64   `bson:"preview_time" json:"preview_time"`
	BeginTime   int64   `bson:"begin_time" json:"begin_time"`
	EndTime     int64   `bson:"end_time" json:"end_time"`
	CloseTime   int64   `bson:"close_time" json:"close_time"`
	Servers     []int16 `bson:"servers" json:"servers"`
	ShowRewards bool    `bson:"show_rewards" json:"show_rewards"`
	ShowRanks   bool    `bson:"show_ranks" json:"show_ranks"`
	Arg1        int32   `bson:"arg1" json:"arg1"`
	Arg2        int32   `bson:"arg2" json:"arg2"`
	LangLst     []int8  `bson:"servers" json:"lang_lst"`
}

type ActivityList struct {
	ActivityList []ActivityInfo `json:"activity_list"`
}

type AddActivityRes struct {
	Error int16 `json: "error"`
}

type ModActivityRes struct {
	Error int16 `json: "error"`
}

type DelActivityRes struct {
	Error int16 `json: "error"`
}
