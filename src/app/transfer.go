package app

import (
	"fmt"
	"math/rand"

	"github.com/oatsaysai/simple-core-bank/src/custom_error"
	log "github.com/oatsaysai/simple-core-bank/src/logger"
	"github.com/oatsaysai/simple-core-bank/src/model"
	"github.com/shopspring/decimal"
)

func (ctx *Context) TransferIn(params model.TransferInParams) (*model.TransferInResponse, error) {
	logger := ctx.getLogger()
	logger = logger.WithFields(log.Fields{
		"func": "TransferIn",
	})
	logger.Info("Begin")
	logger.Debugf("params: %v", params)
	defer logger.Info("End")

	if err := ValidateInput(params); err != nil {
		logger.Errorf("validateInput error : %s", err)
		return nil, err
	}

	txID, err := ctx.DB.TransferIn(
		params.ToAccountNo,
		decimal.NewFromFloat(params.Amount),
	)
	if err != nil {
		return nil, &custom_error.UserError{
			Code:    custom_error.DBError,
			Message: err.Error(),
		}
	}

	return &model.TransferInResponse{
		TransactionID: *txID,
		ToAccountNo:   params.ToAccountNo,
		Amount:        params.Amount,
	}, nil
}

func (ctx *Context) TransferOut(params model.TransferOutParams) (*model.TransferOutResponse, error) {
	logger := ctx.getLogger()
	logger = logger.WithFields(log.Fields{
		"func": "TransferOut",
	})
	logger.Info("Begin")
	logger.Debugf("params: %v", params)
	defer logger.Info("End")

	if err := ValidateInput(params); err != nil {
		logger.Errorf("validateInput error : %s", err)
		return nil, err
	}

	txID, err := ctx.DB.TransferOut(
		params.FromAccountNo,
		decimal.NewFromFloat(params.Amount),
	)
	if err != nil {
		return nil, &custom_error.UserError{
			Code:    custom_error.DBError,
			Message: err.Error(),
		}
	}

	return &model.TransferOutResponse{
		TransactionID: *txID,
		FromAccountNo: params.FromAccountNo,
		Amount:        params.Amount,
	}, nil
}

func (ctx *Context) Transfer(params model.TransferParams) (*model.TransferResponse, error) {
	logger := ctx.getLogger()
	logger = logger.WithFields(log.Fields{
		"func": "Transfer",
	})
	logger.Info("Begin")
	logger.Debugf("params: %v", params)
	defer logger.Info("End")

	if err := ValidateInput(params); err != nil {
		logger.Errorf("validateInput error : %s", err)
		return nil, err
	}

	txID, err := ctx.DB.Transfer(
		params.FromAccountNo,
		params.ToAccountNo,
		decimal.NewFromFloat(params.Amount),
	)
	if err != nil {
		return nil, &custom_error.UserError{
			Code:    custom_error.DBError,
			Message: err.Error(),
		}
	}

	return &model.TransferResponse{
		TransactionID: *txID,
		FromAccountNo: params.FromAccountNo,
		ToAccountNo:   params.ToAccountNo,
		Amount:        params.Amount,
	}, nil
}

func (ctx *Context) TransferInForLoadTest(params model.TransferForLoadTestParams) (*model.TransferInResponse, error) {
	logger := ctx.getLogger()
	logger = logger.WithFields(log.Fields{
		"func": "TransferInForLoadTest",
	})
	logger.Info("Begin")
	logger.Debugf("params: %v", params)
	defer logger.Info("End")

	if err := ValidateInput(params); err != nil {
		logger.Errorf("validateInput error : %s", err)
		return nil, err
	}

	// Random account_no
	toAccountNo := fmt.Sprintf("%010d", rand.Intn(params.MaxAccountNo)+1)

	txID, err := ctx.DB.TransferIn(
		toAccountNo,
		decimal.NewFromFloat(params.Amount),
	)
	if err != nil {
		return nil, &custom_error.UserError{
			Code:    custom_error.DBError,
			Message: err.Error(),
		}
	}

	return &model.TransferInResponse{
		TransactionID: *txID,
		ToAccountNo:   toAccountNo,
		Amount:        params.Amount,
	}, nil
}

func (ctx *Context) TransferOutForLoadTest(params model.TransferForLoadTestParams) (*model.TransferOutResponse, error) {
	logger := ctx.getLogger()
	logger = logger.WithFields(log.Fields{
		"func": "TransferOutForLoadTest",
	})
	logger.Info("Begin")
	logger.Debugf("params: %v", params)
	defer logger.Info("End")

	if err := ValidateInput(params); err != nil {
		logger.Errorf("validateInput error : %s", err)
		return nil, err
	}

	// Random account_no
	fromAccountNo := fmt.Sprintf("%010d", rand.Intn(params.MaxAccountNo)+1)

	txID, err := ctx.DB.TransferOut(
		fromAccountNo,
		decimal.NewFromFloat(params.Amount),
	)
	if err != nil {
		return nil, &custom_error.UserError{
			Code:    custom_error.DBError,
			Message: err.Error(),
		}
	}

	return &model.TransferOutResponse{
		TransactionID: *txID,
		FromAccountNo: fromAccountNo,
		Amount:        params.Amount,
	}, nil
}

func (ctx *Context) TransferForLoadTest(params model.TransferForLoadTestParams) (*model.TransferResponse, error) {
	logger := ctx.getLogger()
	logger = logger.WithFields(log.Fields{
		"func": "TransferForLoadTest",
	})
	logger.Info("Begin")
	logger.Debugf("params: %v", params)
	defer logger.Info("End")

	if err := ValidateInput(params); err != nil {
		logger.Errorf("validateInput error : %s", err)
		return nil, err
	}

	// Random account_no
	fromAccountNo := fmt.Sprintf("%010d", rand.Intn(params.MaxAccountNo)+1)
	toAccountNo := fmt.Sprintf("%010d", rand.Intn(params.MaxAccountNo)+1)

	txID, err := ctx.DB.Transfer(
		fromAccountNo,
		toAccountNo,
		decimal.NewFromFloat(params.Amount),
	)
	if err != nil {
		return nil, &custom_error.UserError{
			Code:    custom_error.DBError,
			Message: err.Error(),
		}
	}

	return &model.TransferResponse{
		TransactionID: *txID,
		FromAccountNo: fromAccountNo,
		ToAccountNo:   toAccountNo,
		Amount:        params.Amount,
	}, nil
}
