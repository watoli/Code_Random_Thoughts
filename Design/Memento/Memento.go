package Memento

type Memento struct {
	state string
}

func (m *Memento) getSavedState() string {
	return m.state
}

type Originator struct {
	state string
}

func (e *Originator) createMemento() *Memento {
	return &Memento{state: e.state}
}

func (e *Originator) restoreMemento(m *Memento) {
	e.state = m.getSavedState()
}

func (e *Originator) setState(state string) {
	e.state = state
}

func (e *Originator) getState() string {
	return e.state
}

type Caretaker struct {
	mementoArray []*Memento
}

func (c *Caretaker) addMemento(m *Memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *Caretaker) getMemento(index int) *Memento {
	return c.mementoArray[index]
}
