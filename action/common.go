package action

type Action interface {
	Key() string
	Unique() string
	Columns() string
	Type() string
	Valid() bool
}

type Actions []Action

func (slice Actions) Len() int {
	return len(slice)
}

func (slice Actions) Less(i, j int) bool {
	return slice[i].Key() < slice[j].Key()
}

func (slice Actions) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
