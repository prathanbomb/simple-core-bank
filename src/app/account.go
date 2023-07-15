package app

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/oatsaysai/simple-core-bank/src/custom_error"
	log "github.com/oatsaysai/simple-core-bank/src/logger"
	"github.com/oatsaysai/simple-core-bank/src/model"
	"github.com/shopspring/decimal"
)

const ACCOUNT_NO_PREFIX = "007"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (ctx *Context) CreateAccount(params model.CreateAccountParams) (*model.CreateAccountResponse, error) {
	logger := ctx.getLogger()
	logger = logger.WithFields(log.Fields{
		"func": "CreateAccount",
	})
	logger.Info("Begin")
	logger.Debugf("params: %v", params)
	defer logger.Info("End")

	if err := ValidateInput(params); err != nil {
		logger.Errorf("validateInput error : %s", err)
		return nil, err
	}

	// Random account_no
	accountNo := fmt.Sprintf("%s%07d", ACCOUNT_NO_PREFIX, rand.Intn(1000000))

	for {
		if result, err := ctx.DB.AccountExists(accountNo); err != nil {
			logger.Errorf("Failed to check account exists: %s", err)
			return nil, &custom_error.InternalError{
				Code:    custom_error.DBError,
				Message: err.Error(),
			}
		} else if !result {
			break
		}

		// if account exists, generate new accountNo and check again
		accountNo = fmt.Sprintf("%s%07d", ACCOUNT_NO_PREFIX, rand.Intn(1000000))
	}

	if err := ctx.DB.InsertAccount(
		accountNo,
		params.AccountName,
		decimal.NewFromInt(0),
	); err != nil {
		logger.Errorf("Failed to insert account: %s", err)
		return nil, &custom_error.InternalError{
			Code:    custom_error.DBError,
			Message: err.Error(),
		}
	}

	return &model.CreateAccountResponse{
		AccountNo:   accountNo,
		AccountName: params.AccountName,
	}, nil
}

func (ctx *Context) GetAccount(params *model.GetAccountParams) (*model.GetAccountResponse, error) {
	logger := ctx.getLogger()
	logger = logger.WithFields(log.Fields{
		"func": "GetAccount",
	})
	logger.Info("Begin")
	logger.Debugf("params: %v", params)
	defer logger.Info("End")

	if err := ValidateInput(params); err != nil {
		logger.Errorf("validateInput error : %s", err)
		return nil, err
	}

	accountNo := params.AccountNo
	accNo, accName, balance, err := ctx.DB.GetAccount(accountNo)
	if err != nil {
		logger.Errorf("Failed to get account: %s", err)
		return nil, &custom_error.InternalError{
			Code:    custom_error.DBError,
			Message: err.Error(),
		}
	}

	balanceFloat, _ := balance.Float64()

	return &model.GetAccountResponse{
		AccountNo:   *accNo,
		AccountName: *accName,
		Balance:     balanceFloat,
	}, nil
}
