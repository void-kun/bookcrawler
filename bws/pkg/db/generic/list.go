package generic

type ListQueryCondition struct {
	Key   string
	Op    string
	Value string
}

type ListQuery []ListQueryCondition

type ListOptions struct {
	Skip  int
	Limit int
	Sort  []ListSort
}
