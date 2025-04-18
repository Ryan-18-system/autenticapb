package handlerException

import (
	"context"
	"errors"
	"fmt"
	"log"
)

func HandlerError(err error) error {
	if errors.Is(err, context.DeadlineExceeded) {
		log.Println("Erro ao executar processo, tempo excedido")
		return fmt.Errorf(":timeout: %w", err)
	}
	return err
}
