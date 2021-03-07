package menu

type MenuVo struct {
	Name       string `json:"name"`
	Component  string `json:"component"`
	Path       string `json:"path"`
	Hidden     int8 `json:"hidden"`
	AlwaysShow bool `json:"alwaysShow"`
	Redirect   string `json:"redirect"`
	Meta       MenuMetaVo `json:"meta"`
	Children   []MenuVo `json:"children"`
}

type MenuMetaVo struct {
	Title string `json:"title"`
	Icon string `json:"icon"`
	NoCache bool `json:"noCache"`
}
