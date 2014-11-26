package engine

type KnowledgeBase struct {
	functions map[string]Function
	rules     map[string]Rule
}

func NewKnowledgeBase() KnowledgeBase {
	var result KnowledgeBase
	result.functions = make(map[string]Function)
	result.rules = make(map[string]Rule)
	return result
}

func (self KnowledgeBase) Function(name string) (Function, bool) {
	value, check := self.functions[name]
	return value, check
}

func (self KnowledgeBase) Functions() map[string]Function {
	return self.functions
}

func (self KnowledgeBase) Rule(name string) (Rule, bool) {
	value, check := self.rules[name]
	return value, check
}

func (self KnowledgeBase) Rules() map[string]Rule {
	return self.rules
}

func (self KnowledgeBase) AddFunction(name string, f Function) bool {
	_, check := self.functions[name]

	if !check {
		self.functions[name] = f
	}

	return check
}

func (self KnowledgeBase) AddRule(name string, r Rule) bool {
	_, check := self.rules[name]

	if !check {
		self.rules[name] = r
	}

	return check
}
