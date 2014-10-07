package engine

var (
	knowledgeBase KnowledgeBase
)

func Init() {
	knowledgeBase = NewKnowledgeBase()
}
