package Command

import "testing"

func TestInvoker_ExecuteCommand(t *testing.T) {
	receiver := &Receiver{name: "小张"}
	command := &ConcreteCommand{receiver}
	invoker := &Invoker{}
	invoker.SetCommand(command)
	invoker.ExecuteCommand()
}
