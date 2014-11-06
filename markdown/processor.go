package markdown

import (
	"io"
	"io/ioutil"

	"github.com/russross/blackfriday"
)

type Markdown struct{}

func New() *Markdown {
	return &Markdown{}
}

func (m *Markdown) Process(r io.Reader) ([]byte, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return []byte{}, err
	}
	return blackfriday.MarkdownCommon(data), nil
}
