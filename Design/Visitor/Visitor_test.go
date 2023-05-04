package Visitor

import (
	"fmt"
	"testing"
)

func TestBonusVisitor_VisitManager(t *testing.T) {
	employees := []Employee{
		&Manager{Name: "Alice", Salary: 8000},
		&Developer{Name: "Bob", Salary: 5000},
	}

	bonusVisitor := &BonusVisitor{}

	for _, e := range employees {
		e.Accept(bonusVisitor)
		fmt.Printf("%s's new salary: %.2f\n", e.GetName(), e.GetSalary())
	}
}
