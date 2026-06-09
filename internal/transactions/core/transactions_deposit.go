package core

import (
	"errors"
	"mini-project-a/internal/shared/constant"
	"mini-project-a/internal/transactions/core/entity"
)

func (s *TransactionsService) Deposit(
	accountNumber string,
	transaction *entity.Transactions,
) (*entity.Transactions, error) {
	account, err := s.accountReaderPort.GetAccountDetailByAccountNumber(accountNumber)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, errors.New("account not found")
	}

	if account.Status == constant.ACCOUNTS_STATUS_CLOSED {
		return nil, errors.New("Account is closed")
	}

	balanceBefore := account.Balance
	balanceAfter := account.Balance + transaction.Amount

	err = s.accountReaderPort.UpdateBalance(account.ID, balanceAfter)
	if err != nil {
		return nil, err
	}

	transaction.AccountID = account.ID
	transaction.TransactionType = constant.TRANSACTIONS_TYPE_DEPOSIT
	transaction.BalanceBefore = balanceBefore
	transaction.BalanceAfter = balanceAfter
	return s.CreateTransaction(transaction)
}
