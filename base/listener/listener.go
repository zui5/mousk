package listener

import hook "github.com/robotn/gohook"

// var RegisterKeys []RegisterKey
var registerFuncList []RegisterFunc

type RegisterFunc struct {
	Key          RegisterKey
	CallbackFunc func(hook.Event)
}
type RegisterKey struct {
	Cmd  []string
	When uint8
}

// add one listner func to var registerFncList
func RegisterOne(key RegisterKey, cb func(hook.Event)) {
	registerFuncList = append(registerFuncList, RegisterFunc{key, cb})
}

// ergodic the registerFuncList register event and start
func Start() {
	for _, v := range registerFuncList {
		hook.Register(v.Key.When, v.Key.Cmd, v.CallbackFunc)
	}

	s := hook.Start()
	<-hook.Process(s)
}
