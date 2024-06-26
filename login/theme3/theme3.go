package theme3

import (
	"github.com/go-hq/components/login"
	"github.com/go-hq/go-admin/template"
)

type Theme3 struct {
	*template.BaseComponent
}

func (*Theme3) GetAssetList() []string {
	return AssetsList
}

func (*Theme3) GetAsset(name string) ([]byte, error) {
	return Asset(name)
}

func (*Theme3) GetHTML() string {
	return List["login"]
}

func init() {
	login.Register("theme3", new(Theme3))
}
