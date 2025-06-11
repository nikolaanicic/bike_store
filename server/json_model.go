package server

type JsonModel interface {
	AsJson() ([]byte, error)
	FromJson(data []byte) error
}
