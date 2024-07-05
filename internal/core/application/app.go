package application

import (
	"context"
	"errors"
	l "service-hf-client-p5/external/logger"
	ps "service-hf-client-p5/external/strings"
	"service-hf-client-p5/internal/core/domain/entity/dto"
	"service-hf-client-p5/internal/core/domain/rpc"
)

type Application interface {
	GetClientByCPF(msgID string, cpf string) (*dto.OutputClient, error)
	SaveClient(msgID string, client dto.RequestClient) (*dto.OutputClient, error)
}

type application struct {
	ctx             context.Context
	clientRPC       rpc.ClientRPC
	clientWorkerRPC rpc.ClientWorkerRPC
}

func (app *application) setMessageIDCtx(msgID string) {
	
}

func NewApplication(ctx context.Context, clientRPC rpc.ClientRPC, clientWorkerRPC rpc.ClientWorkerRPC) Application {
	return application{
		ctx:             ctx,
		clientRPC:       clientRPC,
		clientWorkerRPC: clientWorkerRPC,
	}
}

func (app application) GetClientByCPF(msgID string, cpf string) (*dto.OutputClient, error) {
	app.setMessageIDCtx(msgID)

	l.Infof(msgID, "GetClientByCPFApp: ", " | ", cpf)

	go app.clientRPC.GetClientByCPF(cpf) // pub

	clientRpc, err := app.clientWorkerRPC.GetClientByCPF(cpf)

	if err != nil {
		l.Errorf(msgID, "GetClientByCPFApp error: ", " | ", err)
		return nil, err
	}

	if clientRpc == nil {
		l.Infof(msgID, "GetClientByCPFApp output: ", " | ", nil)
		return nil, nil
	}

	client := &dto.OutputClient{
		UUID:      		clientRpc.UUID,
		Name:      		clientRpc.Name,
		CPF:       		clientRpc.CPF,
		Email:     		clientRpc.Email,
		PhoneNumber:	clientRpc.PhoneNumber,
		Address:   		clientRpc.Address,
		CreatedAt: 		clientRpc.CreatedAt,
	}

	l.Infof(msgID, "GetClientByCPFApp output: ", " | ", client)
	return client, nil
}

func (app application) SaveClient(msgID string, client dto.RequestClient) (*dto.OutputClient, error) {
	app.setMessageIDCtx(msgID)

	l.Infof(msgID, "SaveClientApp: ", " | ", ps.MarshalString(client))

	go app.clientRPC.SaveClient(client) // pub

	pRepo, err := app.clientWorkerRPC.SaveClient(client)

	if err != nil {
		l.Errorf(msgID, "SaveClientApp error: ", " | ", err)
		return nil, err
	}

	if pRepo == nil {
		l.Infof(msgID, "SaveClientApp output: ", " | ", nil)
		return nil, errors.New("is not possible to save client because it's null")
	}

	out := &dto.OutputClient{
		UUID:          		pRepo.UUID,
		Name:      	   		pRepo.Name,
		CPF:           		pRepo.CPF,
		Email:   	   		pRepo.Email,
		PhoneNumber:        pRepo.PhoneNumber,
		Address:       		pRepo.Address,
		CreatedAt:     		pRepo.CreatedAt,
	}

	l.Infof(msgID, "SaveClientApp output: ", " | ", ps.MarshalString(out))
	return out, nil
}

