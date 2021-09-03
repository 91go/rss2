package porn

import (
	"fmt"
	"testing"

	"github.com/91go/rss2/core"
)

func TestPorn(t *testing.T) {
	url := "https://jiuse911.com/author/Hhonswifelonely"

	doc := core.FetchHTML(url).Text()

	fmt.Println(doc)
}
