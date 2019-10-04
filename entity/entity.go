package entity

// ErrorResult 处理失败的标准应答
type ErrorResult struct {
	Ret    int64
	Reason string
}

// RandomCardsOKResult 随机获取卡片接口的成功应答
type RandomCardsOKResult struct {
	Ret   int64
	Cards []Card
}

// Desc 卡片/单章/手册的说明 包含标题、分类、标签
type Desc struct {
	Title    string
	Category []string
	Tags     []string
}

// Card 单张卡片 包含说明和内容
type Card struct {
	Desc    Desc
	Content string
}

// Cards 卡片列表 多张卡片
type Cards struct {
	Cards []Card
}

// ShortPage 单章 包含说明和内容
type ShortPage struct {
	Desc    Desc
	Content string
}

// Page 手册中的一页
type Page struct {
	Content string
}

// HandBook 手册
type HandBook struct {
	Desc       Desc
	Preface    Page
	PostScript Page
	Pages      []Page
}
