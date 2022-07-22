package output

type Output interface {
	Write(url string)
}
