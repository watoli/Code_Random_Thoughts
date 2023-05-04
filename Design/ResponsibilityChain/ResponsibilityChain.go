package ResponsibilityChain

import "fmt"

type Department interface {
	execute(*Patient)
	setNext(Department)
}

type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(next Department) {
	r.next = next
}

type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}

type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(next Department) {
	m.next = next
}

type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}
