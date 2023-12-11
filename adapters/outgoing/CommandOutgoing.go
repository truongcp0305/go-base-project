package outgoing

type ListProcessOutgoing struct {
	Data string `json:"data" example:"notepad.exe      1021"`
}

type KillProcessOutgoing struct {
	Data string `json:"data" example:"process killed!"`
}

type ExecuteScriptOutgoing struct {
	Data string `json:"data" example:"excute done!"`
}

type OpenFileSucessOutgoing struct {
	Data string `json:"data" example:"opening file"`
}

type OpenFileFailedOutgoing struct {
	Data string `json:"data" example:"cannot find file with path"`
}
