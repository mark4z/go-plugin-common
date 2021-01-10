package filter

type FilterFunc func(Context)

type Context interface {
	API(a string)
	GetAPI() string
}
