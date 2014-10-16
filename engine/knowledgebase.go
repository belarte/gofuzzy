package engine

type KnowledgeBase struct {
	functions map[string]Function
	rules     map[string]Expression
}

func NewKnowledgeBase() KnowledgeBase {
	var result KnowledgeBase
	result.functions = make(map[string]Function)
	result.rules = make(map[string]Expression)
	return result
}

func (self KnowledgeBase) GetFunction(name string) (Function, bool) {
	value, check := self.functions[name]
	return value, check
}

func (self KnowledgeBase) GetRule(name string) (Expression, bool) {
	value, check := self.rules[name]
	return value, check
}

func (self KnowledgeBase) AddFunction(name string, f Function) bool {
	_, check := self.functions[name]

	if !check {
		self.functions[name] = f
	}

	return check
}

func (self KnowledgeBase) AddRule(name string, f Expression) bool {
	_, check := self.rules[name]

	if !check {
		self.rules[name] = f
	}

	return check
}
