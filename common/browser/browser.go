package browser

import (
	"github.com/nowhszh/hack-browser-data/common/browingdata"
)

type Browser interface {
	// Name is browser's name
	Name() string
	// BrowsingData returns all browsing data in the browser.
	BrowsingData() (*browingdata.Data, error)
}
