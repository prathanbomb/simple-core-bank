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
		"func": "GetAdminAPIKey",
	})
	logger.Info("Begin")
	logger.Debugf("params: %v", params)
	defer logger.Info("End")

	if err := ValidateInput(params); err != nil {
		logger.Errorf("validateInput error : %s", err)
		return nil, err
	}

	// Random account_no
	accountNo := fmt.Sprintf("%s%d", ACCOUNT_NO_PREFIX, rand.Intn(1000000))

	// TODO: Check duplicate account_no
	err := ctx.DB.InsertAccount(
		accountNo,
		params.AccountName,
		decimal.NewFromInt(0),
	)
	if err != nil {
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
