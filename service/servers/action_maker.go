package servers

import "flint/service/contracts"

type ActionMaker struct {
	remote contracts.RemoteAction
}

func NewActionMaker(remote contracts.RemoteAction) *ActionMaker {
	return &ActionMaker{remote: remote}
}
