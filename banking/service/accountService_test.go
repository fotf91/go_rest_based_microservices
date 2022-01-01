package service

import (
	"banking/customErrors"
	domain "banking/domain"
	"banking/dto"
	mockDomain "banking/mocks/domain"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_should_return_a_validation_error_response_when_the_request_is_not_validated(t *testing.T) {
	// Arrange
	request := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      0, // this amount will fail the validation
	}

	service := NewAccountService(nil)

	//  Act
	_, appError := service.NewAccount(request)

	// Assert
	if appError == nil {
		t.Error("failed while testing the new account validation")
	}
}

var mockRepo *mockDomain.MockAccountRepository
var service AccountService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)

	mockRepo = mockDomain.NewMockAccountRepository(ctrl)
	service = NewAccountService(mockRepo)
	return func() {
		service = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_an_error_from_the_server_side_if_the_new_account_cannot_be_created(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      6000,
	}

	account := domain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: dbTSLayout,
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	mockRepo.EXPECT().Save(account).Return(nil, customErrors.NewUnexpectedError("unexpected db error"))

	// Act
	_, appError := service.NewAccount(req)

	// Assert
	if appError == nil {
		t.Error("Test failed while validating error for new account")

	}
}

func Test_should_return_new_account_response_when_a_new_account_is_saved_successfully(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      6000,
	}
	account := domain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: dbTSLayout,
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	accountWithId := account
	accountWithId.AccountId = "201"
	mockRepo.EXPECT().Save(account).Return(&accountWithId, nil)
	// Act
	newAccount, appError := service.NewAccount(req)

	// Assert
	if appError != nil {
		t.Error("Test failed while creating new account")
	}
	if newAccount.AccountId != accountWithId.AccountId {
		t.Error("Failed while mathching new account id")
	}
}
