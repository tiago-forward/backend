package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (wd *workInfoDomainService) FindWorkInfoByUserIdServices(
	userId string,
) (model.WorkInfoDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindWorkInfoByUserId service", zap.String("journey", "findWorkInfo"))

	return wd.workInfoRepository.FindWorkInfoByUserId(userId)
}
