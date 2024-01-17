package mock

type aaa struct {
	Contents string
}

func (r aaa) Get(url string) string {
	return r.Contents
}
