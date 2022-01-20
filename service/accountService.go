package service

import (
	"bankingV2/domain"
	"bankingV2/dto"
	"bankingV2/errs"
	"log"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}
type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		log.Println("SOMETHING WENT WRONG")
		return nil, err

	}
	if req.IsTransactionTypeWithdrawal() {
		account, err := s.repo.ById(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}
	log.Println(req)

	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	transaction, appError := s.repo.SaveTransaction(t)
	log.Println(transaction)
	if appError != nil {
		return nil, appError
	}
	response := transaction.ToTransactionResponseDto()
	return &response, nil
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		log.Println("SOMETHING WENT WRONG")
		return nil, err

	}
	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	log.Println(a)
	NewAccount, err := s.repo.Save(a)
	if err != nil {
		log.Println("SOMETHING WENT WRONG")
		return nil, err
	}
	response := NewAccount.ToNewAccountResponseDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
