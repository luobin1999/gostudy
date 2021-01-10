package mock

type Retriever struct {
	Contents string
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}

func (r *Retriever) Post(url string, con map[string]string) string {
	r.Contents = con["contents"]
	return "ok"
}
