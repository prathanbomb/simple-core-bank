package app

import (
	log "github.com/oatsaysai/simple-core-bank/src/logger"
	"github.com/oatsaysai/simple-core-bank/src/model"
)

func (ctx *Context) GetTransactionByAccountNo(params *model.GetTransactionParams) ([]model.Transaction, error) {
	logger := ctx.getLogger()
	logger = logger.WithFields(log.Fields{
		"func": "GetTransactionByAccountNo",
	})
	logger.Info("Begin")
	logger.Debugf("params: %v", params)
	defer logger.Info("End")

	if err := ValidateInput(params); err != nil {
		logger.Errorf("validateInput error : %s", err)
		return nil, err
	}

	return ctx.DB.GetTransactionByAccountNo(params.AccountNo)
}
