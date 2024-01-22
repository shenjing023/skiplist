package skiplist

type options struct {
	SkipListP float64
	MaxLevel  int
}

func NewOptions() *options {
	return &options{
		SkipListP: SKIPLIST_P,
		MaxLevel:  MAX_LEVEL,
	}
}

type Option func(*options)
