package incoming

type KillProcessParams struct {
	Pid  string `json:"pId" form:"pId" example:"12345"`
	Name string `json:"name" form:"name" example:"notepad.exe"`
}

type ExcecuteScriptParams struct {
	Path string `json:"path" form:"path" example:".\\script.ps1"`
}

type GetFileParams struct {
	Path string `json:"path" form:"path" example:".\\script.ps1"`
}
