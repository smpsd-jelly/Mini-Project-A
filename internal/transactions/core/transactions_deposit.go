package core

import (
	"database/sql"
	"errors"
	"mini-project-a/internal/shared/constant"
	"mini-project-a/internal/transactions/core/entity"
)

func (s *TransactionsService) Deposit(
	accountNumber string,
	transaction *entity.Transactions,
) (*entity.Transactions, error) {
	account, err := s.accountReaderPort.GetAccountDetailByAccountNumberRepo(accountNumber)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAccountNotFound
		}

		return nil, err
	}

	if account.Status == constant.ACCOUNTS_STATUS_CLOSED {
		return nil, ErrAccountIsClosed
	}

	balanceBefore := account.Balance
	balanceAfter := account.Balance + transaction.Amount

	transaction.AccountID = account.ID
	transaction.TransactionType = constant.TRANSACTIONS_TYPE_DEPOSIT
	transaction.BalanceBefore = balanceBefore
	transaction.BalanceAfter = balanceAfter

	tx, err := s.transactionsRepo.BeginTx()
	if err != nil {
		return nil, err
	}

	if err := s.accountReaderPort.UpdateBalanceTx(tx, account.ID, balanceAfter); err != nil {
		return nil, err
	}

	result, err := s.transactionsRepo.CreateTransactionTx(tx, transaction)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return result, nil
}
