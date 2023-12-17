package gapi

import (
	"context"
	"errors"
	"fmt"

	repository "github.com/RoyceAzure/go-stockinfo-distributor/repository/db/sqlc"
	logger "github.com/RoyceAzure/go-stockinfo-distributor/repository/logger_distributor"
	"github.com/RoyceAzure/go-stockinfo-distributor/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util"
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) CreateClientRegister(ctx context.Context, req *pb.CreateClientRegisterRequest) (*pb.CreateClientRegisterResponse, error) {
	_, err := server.authorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}
	violations := validateCreateClientRegisterRequest(req)
	if violations != nil {
		return nil, util.InvalidArgumentError(violations)
	}
	var fc repository.FrontendClient
	if req.ClientId == "" {
		md := util.ExtractMetaData(ctx)
		if !util.IsValidIP(md.ClientIP) {
			err := fmt.Errorf("ip is invalid")
			logger.Logger.Error().Err(err).Msg("get client register get err")
			return nil, util.InValidateOperation(err)
		}
		fc, err = server.dbDao.GetFrontendClientByIP(ctx, md.ClientIP)
	} else {
		clientId, _ := uuid.Parse(req.ClientId)
		fc, err = server.dbDao.GetFrontendClientByID(ctx, clientId)
	}
	if err != nil {
		msg := "can't find client with ip/id"
		logger.Logger.Error().Err(err).Msg(msg)
		return nil, util.InternalError(errors.New(msg))
	}

	crParm := repository.CreateClientRegisterParams{
		ClientUid: fc.ClientUid,
		StockCode: req.StockCode,
	}

	entity, err := server.dbDao.CreateClientRegister(ctx, crParm)
	if err != nil {
		errCode := repository.ErrorCode(err)
		if errCode == repository.UniqueViolation {
			err = errors.New("client register is already register")
			logger.Logger.Error().Err(err).Msg("create client register err")
			return nil, util.InValidateOperation(err)
		} else if errCode == repository.ForeginKeyViolation {
			logger.Logger.Error().Err(err).Msg("create client register err")
			return nil, util.InternalError(err)
		}
		logger.Logger.Error().Err(err).Msg("create client register err")
		return nil, util.InternalError(err)
	}
	return &pb.CreateClientRegisterResponse{
		ClientId:  entity.ClientUid.String(),
		StockCode: entity.StockCode,
	}, nil
}

func (server *Server) DeleteClientRegister(ctx context.Context, req *pb.DeleteClientRegisterRequest) (*pb.DeleteClientRegisterResponse, error) {
	_, err := server.authorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}
	violations := validateDeleteClientRegisterRequest(req)
	if violations != nil {
		return nil, util.InvalidArgumentError(violations)
	}

	clientId, _ := uuid.Parse(req.ClientId)
	arg := repository.DeleteClientRegisterParams{
		ClientUid: clientId,
		StockCode: req.StockCode,
	}

	err = server.dbDao.DeleteClientRegister(ctx, arg)
	if err != nil {
		err = fmt.Errorf("delete client register get err : %w", err)
		return nil, util.InternalError(err)
	}

	return &pb.DeleteClientRegisterResponse{
		Result: fmt.Sprintf("success deleted client with id : %s", clientId),
	}, nil
}

func (server *Server) GetClientRegisterByClientUID(ctx context.Context, req *pb.GetClientRegisterByClientUIDRequest) (*pb.GetClientRegisterResponse, error) {
	_, err := server.authorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}
	violations := validateGetClientRegisterByClientUIDRequest(req)
	if violations != nil {
		return nil, util.InvalidArgumentError(violations)
	}

	clientId, _ := uuid.Parse(req.ClientId)

	entity, err := server.dbDao.GetClientRegisterByClientUID(ctx, clientId)
	if err != nil {
		if err == repository.ErrRecordNotFound {
			return nil, nil
		}
		return nil, util.InternalError(err)
	}

	res := pb.GetClientRegisterResponse{}

	for _, item := range entity {
		res.Data = append(res.Data, cvClientRegisterE2DTO(&item))
	}

	return &res, nil
}

func validateDeleteClientRegisterRequest(req *pb.DeleteClientRegisterRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := util.ValidUUID(req.ClientId); err != nil {
		violations = append(violations, util.FieldViolation("client_id", err))
	}
	if err := util.ValidateEmptyString(req.StockCode); err != nil {
		violations = append(violations, util.FieldViolation("stock_code", err))
	}
	return violations
}

func validateCreateClientRegisterRequest(req *pb.CreateClientRegisterRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.ClientId != "" {
		if err := util.ValidUUID(req.ClientId); err != nil {
			violations = append(violations, util.FieldViolation("client_id", err))
		}
	}
	if err := util.ValidateEmptyString(req.StockCode); err != nil {
		violations = append(violations, util.FieldViolation("stock_code", err))
	}
	return violations
}

func validateGetClientRegisterByClientUIDRequest(req *pb.GetClientRegisterByClientUIDRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := util.ValidUUID(req.ClientId); err != nil {
		violations = append(violations, util.FieldViolation("client_id", err))
	}
	return violations
}

func cvClientRegisterE2DTO(entity *repository.ClientRegister) *pb.ClientRegister {
	return &pb.ClientRegister{
		ClientId:  entity.ClientUid.String(),
		StockCode: entity.StockCode,
	}
}
