package bangumi

type SubjectType int

const (
	BOOK SubjectType = iota + 1
	ANIME
	MUSIC
	GAME
	REAL
)

// Subject 条目信息
type Subject struct {
	ID            int           `json:"id"`                 // ID
	Type          int           `json:"type"`               // 类型（如 1=书籍, 2=动画等）
	Name          string        `json:"name"`               // 原名
	NameCN        string        `json:"name_cn"`            // 中文名
	Summary       string        `json:"summary"`            // 简介
	Series        bool          `json:"series"`             // 是否系列作品
	Nsfw          bool          `json:"nsfw"`               // 是否为 NSFW (R18)
	Locked        bool          `json:"locked"`             // 是否锁定（需权限查看）
	Date          string        `json:"date,omitempty"`     // 发售/播出日期（YYYY-MM-DD 或 YYYY-MM）
	Platform      string        `json:"platform,omitempty"` // 平台（如 Nintendo Switch、PC 等）
	Images        Images        `json:"images"`             // 封面图
	Infobox       []interface{} `json:"infobox,omitempty"`  // Infobox 数据（通常是键值对数组）
	Volumes       int           `json:"volumes"`            // 卷数（书籍）
	Eps           int           `json:"eps"`                // 话数（动画）
	TotalEpisodes int           `json:"total_episodes"`     // 总集数
	Rating        Rating        `json:"rating"`             // 评分信息
	Collection    Collection    `json:"collection"`         // 用户收藏状态统计
	MetaTags      []string      `json:"meta_tags"`          // 公共标签
	Tags          []Tag         `json:"tags"`               // 用户标签列表
}

// Images 封面图地址（通常有多种尺寸）
type Images struct {
	Common string `json:"common"` // 通用图
	Grid   string `json:"grid"`   // 网格展示图
	Large  string `json:"large"`  // 大图
	Small  string `json:"small"`  // 小图
	Thumb  string `json:"thumb"`  // 缩略图
}

// Rating 评分详情
type Rating struct {
	Count struct {
		One   int `json:"1"`
		Two   int `json:"2"`
		Three int `json:"3"`
		Four  int `json:"4"`
		Five  int `json:"5"`
		Six   int `json:"6"`
		Seven int `json:"7"`
		Eight int `json:"8"`
		Nine  int `json:"9"`
		Ten   int `json:"10"`
	} `json:"count"`
	Average float64 `json:"average"` // 平均分
}

// Collection 收藏统计
type Collection struct {
	Doing    int `json:"doing"`    // 在看/在读
	Complete int `json:"complete"` // 完成
	Wish     int `json:"wish"`     // 想看/想读
	OnHold   int `json:"on_hold"`  // 搁置
	Dropped  int `json:"dropped"`  // 放弃
}

// Tag 用户贡献的标签
type Tag struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
