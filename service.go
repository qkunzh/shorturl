package shorturl

type URLCompressor interface {
	Generator(url string) (string, error)
	Revert(url string) (string,error)
}
