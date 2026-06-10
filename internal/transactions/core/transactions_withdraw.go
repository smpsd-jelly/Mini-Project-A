package core

import (
	"errors"
	"mini-project-a/internal/shared/constant"
	"mini-project-a/internal/transactions/core/entity"
)

func (s *TransactionsService) Withdraw(accountNumber string, transaction *entity.Transactions) (*entity.Transactions, error) {
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

	if account.Balance <= 0 {
		return nil, errors.New("Insufficient balance")
	}

	if transaction.Amount <= 0 {
		return nil, errors.New("Amount must be greater than 0")
	}

	if transaction.Amount > account.Balance {
		return nil, errors.New("Amount exceeds available balance")
	}

	balanceBefore := account.Balance
	balanceAfter := balanceBefore - transaction.Amount

	err = s.accountReaderPort.UpdateBalance(account.ID, balanceAfter)
	if err != nil {
		return nil, err
	}

	transaction.AccountID = account.ID
	transaction.TransactionType = constant.TRANSACTIONS_TYPE_WITHDRAWAL
	transaction.BalanceBefore = balanceBefore
	transaction.BalanceAfter = balanceAfter

	return s.CreateTransaction(transaction)
}
