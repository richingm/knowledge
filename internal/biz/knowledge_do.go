package biz

type KnowledgePo struct {
	Hello string
}

func (K KnowledgePo) ConvertToDo() KnowledgeDo {
	return KnowledgeDo{}
}
