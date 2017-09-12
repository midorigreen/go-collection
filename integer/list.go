package integer

type Collection interface {
	Filter(func(interface{}) bool) Collection
	Each(func(int, int))
	Map() Collection
}

type List []int

func (in List) Filter(fn func(interface{}) bool) List {
	out := make(List, 0, len(in))
	for _, v := range in {
		if fn(v) {
			out = append(out, v)
		}
	}
	return out
}

func (in List) Each(fn func(int, int)) {
	for i, v := range in {
		fn(i, v)
	}
}

func (in List) Map(fn func(interface{}) interface{}) List {
	out := make(List, 0, len(in))
	for _, v := range in {
		out = append(out, fn(v).(int))
	}
	return out
}

func CreateList(cap int) List {
	return make(List, 0, cap)
}