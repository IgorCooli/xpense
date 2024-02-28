package expense

import (
	"context"
	"fmt"
	"time"

	"github.com/IgorCooli/xpense/internal/business/model"
	repository "github.com/IgorCooli/xpense/internal/repository/expense"
	"github.com/google/uuid"
)

type Service interface {
	AddExpense(ctx context.Context, expense model.Expense) error
	Search(ctx context.Context, userId string) []model.Expense
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
	buildExpenseId(&expense)

	if expense.Installments == 1 {
		return s.repository.InsertOne(ctx, expense)
	}

	installments := buildInstallments(expense)

	return s.repository.InsertMany(ctx, installments)
}

func buildInstallments(expense model.Expense) []model.Expense {
	var installments []model.Expense
	//TODO dividir o valor total nas parcelas
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
		ID:           expense.ID,
		Value:        expense.Value,
		PaymentDate:  newDate,
		Installments: expense.Installments,
		Description:  newDescription,
		Type:         expense.Type,
		Method:       expense.Method,
		Card:         expense.Card,
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

func buildExpenseId(expense *model.Expense) {
	UUID, err := uuid.NewUUID()

	if err != nil {
		panic("Could not generate uuid")
	}

	expenseId := UUID.String()

	expense.ID = expenseId
}

func (s service) Search(ctx context.Context, userId string) []model.Expense {
	return s.repository.Search(ctx, userId)
}
