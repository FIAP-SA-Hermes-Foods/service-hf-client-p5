package useCase

import (
	"errors"
	"service-hf-client-p5/internal/core/domain/entity/dto"
	"service-hf-client-p5/internal/core/domain/useCase"
)

var _ useCase.ClientUseCase = (*clientUseCase)(nil)

type clientUseCase struct {
}

func NewClientUseCase() clientUseCase {
	return clientUseCase{}
}

func (c clientUseCase) SaveClient(reqClient dto.RequestClient) error {
	client := reqClient.Client()

	if err := client.CPF.Validate(); err != nil {
		return err
	}

	return nil
}

func (c clientUseCase) GetClientByCPF(cpf string) error {
	if len(cpf) == 0 {
		return errors.New("the cpf is not valid for consult")
	}
	return nil
}
