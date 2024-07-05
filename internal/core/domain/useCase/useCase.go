package useCase

import "service-hf-client-p5/internal/core/domain/entity/dto"

type ClientUseCase interface {
	SaveClient(reqProduct dto.RequestClient) error
	GetClientByCPF(cpf string) error
}
