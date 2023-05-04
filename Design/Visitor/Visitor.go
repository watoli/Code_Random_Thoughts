package Visitor

// Employee 定义员工结构体
type Employee interface {
	Accept(v Visitor)
	GetName() string
	GetSalary() float64
}

type Manager struct {
	Name   string
	Salary float64
}

func (m *Manager) Accept(v Visitor) {
	v.VisitManager(m)
}

func (m *Manager) GetName() string {
	return m.Name
}

func (m *Manager) GetSalary() float64 {
	return m.Salary
}

type Developer struct {
	Name   string
	Salary float64
}

func (d *Developer) Accept(v Visitor) {
	v.VisitDeveloper(d)
}

func (d *Developer) GetName() string {
	return d.Name
}

func (d *Developer) GetSalary() float64 {
	return d.Salary
}

// Visitor 定义访问者接口
type Visitor interface {
	VisitManager(e *Manager)
	VisitDeveloper(e *Developer)
}

// BonusVisitor 实现访问者接口
type BonusVisitor struct{}

func (b *BonusVisitor) VisitManager(m *Manager) {
	m.Salary *= 1.2 // 给经理加 20% 的奖金
}

func (b *BonusVisitor) VisitDeveloper(d *Developer) {
	d.Salary *= 1.1 // 给开发人员加 10% 的奖金
}
