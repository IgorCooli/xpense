package service

import (
	"context"
	"fmt"
	"time"

	"github.com/IgorCooli/xpense/internal/business/model"
	repository "github.com/IgorCooli/xpense/internal/repository/expense"
)

type Service interface {
	AddExpense(ctx context.Context, expense model.Expense) error
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return service{
		repository: repository,
	}
}

func (s service) AddExpense(ctx context.Context, expense model.Expense) error {
	if expense.Installments == 1 {
		return s.repository.InsertOne(ctx, expense)
	}

	installments := buildInstallments(expense)

	return s.repository.InsertMany(ctx, installments)
}

func buildInstallments(expense model.Expense) []model.Expense {
	var installments []model.Expense
	for i := 0; i < int(expense.Installments); i++ {
		expenseItem := buildExpenseInstallment(expense, i+1)

		installments = append(installments, expenseItem)
	}

	return installments
}

func buildExpenseInstallment(expense model.Expense, number int) model.Expense {
	newDescription := buildDescriptionWithInstallments(number, expense)
	newDate := handleDate(expense.PaymentDate, number)

	expenseItem := model.Expense{
		Value:        expense.Value,
		PaymentDate:  newDate,
		Installments: expense.Installments,
		Description:  newDescription,
		Type:         expense.Type,
		Method:       expense.Method,
		Card:         expense.Card,
		CardBrand:    expense.CardBrand,
	}
	return expenseItem
}

func handleDate(time time.Time, number int) time.Time {
	if number == 1 {
		return time
	}

	return time.AddDate(0, number-1, 0)
}

func buildDescriptionWithInstallments(number int, expense model.Expense) string {
	installmentDescription := fmt.Sprintf("%v/%v", number, expense.Installments)
	newDescription := fmt.Sprintf("%s %s", expense.Description, installmentDescription)
	return newDescription
}
