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

func (server *Server) CreateFrontendClient(ctx context.Context, req *pb.CreateFrontendClientRequest) (*pb.CreateFrontendClientResponse, error) {
	md := util.ExtractMetaData(ctx)
	logger.Logger.Warn().Msg(md.ClientIP)
	if !util.IsValidIP(md.ClientIP) {
		logger.Logger.Error().Str("ip", md.ClientIP).Msg("get client register get err")
		return nil, util.InValidateOperation(fmt.Errorf("ip is invalid"))
	}

	arg := repository.CreateFrontendClientParams{
		Ip:     md.ClientIP,
		Region: req.Region,
	}

	entity, err := server.dbDao.CreateFrontendClient(ctx, arg)
	if err != nil {
		errCode := repository.ErrorCode(err)
		if errCode == repository.UniqueViolation {
			err = errors.New("frontend client is already register")
			logger.Logger.Error().Err(err).Msg("create frontend client err")
			return nil, util.InValidateOperation(err)
		} else if errCode == repository.ForeginKeyViolation {
			logger.Logger.Error().Err(err).Msg("create frontend client err")
			return nil, util.InternalError(err)
		}
		logger.Logger.Error().Err(err).Msg("create frontend client err")
		return nil, util.InternalError(err)
	}

	logger.Logger.Info().Msg("create frontend client successed")
	return &pb.CreateFrontendClientResponse{
		Data: cvFrontendClientE2DTO(&entity),
	}, nil
}

func (server *Server) DeleteFrontendClient(ctx context.Context, req *pb.DeleteFrontendClientRequest) (*pb.DeleteFrontendClientResponse, error) {
	violations := validateDeleteFrontendClientRequest(req)
	if violations != nil {
		return nil, util.InvalidArgumentError(violations)
	}

	clientId, _ := uuid.Parse(req.ClientId)

	err := server.dbDao.DeleteFrontendClient(ctx, clientId)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("delete client register get err")
		return nil, util.InternalError(err)
	}

	return &pb.DeleteFrontendClientResponse{
		Result: fmt.Sprintf("success deleted frontend client with id : %s", clientId),
	}, nil
}

func (server *Server) GetFrontendClientByIP(ctx context.Context, req *pb.GetFrontendClientByIPRequest) (*pb.GetFrontendClientByIPResponse, error) {
	md := util.ExtractMetaData(ctx)
	logger.Logger.Warn().Msg(md.ClientIP)
	if !util.IsValidIP(md.ClientIP) {
		err := fmt.Errorf("ip is invalid")
		logger.Logger.Error().Err(err).Str("ip", md.ClientIP).Msg("get client register get err")
		return nil, util.InValidateOperation(err)
	}

	fc, err := server.dbDao.GetFrontendClientByIP(ctx, md.ClientIP)
	if err != nil {
		if err == repository.ErrRecordNotFound {
			return &pb.GetFrontendClientByIPResponse{}, nil
		}
		logger.Logger.Error().Err(err).Msg("get client register get err")
		return nil, util.InternalError(err)
	}
	return &pb.GetFrontendClientByIPResponse{
		Data: cvFrontendClientE2DTO(&fc),
	}, nil
}

func validateDeleteFrontendClientRequest(req *pb.DeleteFrontendClientRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := util.ValidUUID(req.ClientId); err != nil {
		violations = append(violations, util.FieldViolation("client_id", err))
	}
	return violations
}

func cvFrontendClientE2DTO(entity *repository.FrontendClient) *pb.FrontendClient {
	return &pb.FrontendClient{
		ClientId: entity.ClientUid.String(),
		Ip:       entity.Ip,
		Region:   entity.Region,
	}
}
