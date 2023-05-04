package Command

import "fmt"

type Command interface {
	Execute()
}

type ConcreteCommand struct {
	receiver *Receiver
}

func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}

type Receiver struct {
	name string
}

func (r *Receiver) Action() {
	fmt.Println("接收者", r.name, "正在执行命令")
}

type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(c Command) {
	i.command = c
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}
