package app

import (
	"github.com/shopspring/decimal"
	"math/rand"
	"time"

	"github.com/oatsaysai/simple-core-bank/src/custom_error"
	log "github.com/oatsaysai/simple-core-bank/src/logger"
	"github.com/oatsaysai/simple-core-bank/src/model"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (ctx *Context) PreGenerateAccountNumbers(params model.PreGenerateAccountNoParams) (*model.PreGenerateAccountNoResponse, error) {
	logger := ctx.getLogger()
	logger = logger.WithFields(log.Fields{
		"func": "PreGenerateAccountNumbers",
	})
	logger.Info("Begin")
	logger.Debugf("params: %v", params)
	defer logger.Info("End")

	// Pre-generate account numbers
	err := ctx.DB.PreGenerateAccountNo(params.BatchSize)
	if err != nil {
		logger.Errorf("Failed to pre-generate account numbers: %s", err)
		return nil, &custom_error.InternalError{
			Code:    custom_error.DBError,
			Message: err.Error(),
		}
	}

	return &model.PreGenerateAccountNoResponse{
		Success:   true,
		BatchSize: params.BatchSize,
	}, nil
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

	// Get the next available account number, mark it as used and insert it into accounts
	accountNo, err := ctx.DB.GetAccountNoAndInsertAccount(params.AccountName, decimal.NewFromInt(0))
	if err != nil {
		logger.Errorf("Failed to get and mark next available account number as used and insert it: %s", err)
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
