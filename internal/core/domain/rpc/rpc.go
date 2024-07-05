package rpc

import "service-hf-client-p5/internal/core/domain/entity/dto"

type ClientRPC interface {
	GetClientByCPF(cpf string) (*dto.OutputClient, error)
	SaveClient(cpf dto.RequestClient) (*dto.OutputClient, error)
}

type ClientWorkerRPC interface {
	GetClientByCPF(cpf string) (*dto.OutputClient, error)
	SaveClient(cpf dto.RequestClient) (*dto.OutputClient, error)
}
