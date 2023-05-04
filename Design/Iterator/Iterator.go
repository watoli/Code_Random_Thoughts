package Iterator

type Collection interface {
	createIterator() Iterator
}

type Iterator interface {
	hasNext() bool
	getNext() *User
}
type User struct {
	name string
	age  int
}

type UserIterator struct {
	index int
	users []*User
}

func (i *UserIterator) hasNext() bool {
	if i.index < len(i.users) {
		return true
	}
	return false
}

func (i *UserIterator) getNext() *User {
	if i.hasNext() {
		user := i.users[i.index]
		i.index++
		return user
	}
	return nil
}

type UserCollection struct {
	users []*User
}

func (c UserCollection) createIterator() Iterator {
	return &UserIterator{
		index: 0,
		users: c.users,
	}
}
