package app

import (
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
		return nil, &custom_error.InternalError{
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
		return nil, &custom_error.InternalError{
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
