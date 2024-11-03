package tests

import (
	"errors"
	"fmt"
	"log"
	"testing"

	"go.uber.org/zap"
)

func TestLog(t *testing.T) {
	errorSysFatal := errors.New("system fatal error")
	fmt.Println(errorSysFatal)
	fmt.Println(errorSysFatal.Error())
	log.Print(errorSysFatal)
	log.Print(errorSysFatal.Error())

	logger, _ := zap.NewProduction()

	logger.Error("system fatal error", zap.Error(errorSysFatal))
	logger.Error(errorSysFatal.Error(), zap.Error(errorSysFatal))

}
