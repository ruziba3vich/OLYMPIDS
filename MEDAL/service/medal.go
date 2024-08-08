package service

import (
	"context"
	"log"
	pb "medal-service/genproto/medals"
	"medal-service/models"
	"medal-service/repositroy"
	"medal-service/storage/postgres"
	"strconv"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MedalService struct {
	repo repositroy.MedalRepository
	pb.UnimplementedMedalServiceServer
	countryMedals postgres.CountryMedals
}

func NewMedalService(repo repositroy.MedalRepository) *MedalService {
	return &MedalService{repo: repo}
}

func (m *MedalService) CreateMedal(ctx context.Context, req *pb.CreateMedalRequest) (*pb.Medal, error) {
	medal := &models.CreateMedal{
		Description: req.GetDescription(),
		AthleteID:   uuid.MustParse(req.GetAthleteId()),
		Type:        req.GetType(),
		Country:     req.GetCountry(),
	}

	createdMedal, err := m.repo.CreateMedal(ctx, medal)
	if err != nil {
		return nil, err
	}

	return &pb.Medal{
		Id:          createdMedal.ID.String(),
		Description: createdMedal.Description,
		AthleteId:   createdMedal.AthleteID.String(),
		Type:        createdMedal.Type,
		Country:     createdMedal.Country,
		CreatedAt:   timestamppb.New(createdMedal.CreatedAt),
		UpdatedAt:   timestamppb.New(createdMedal.UpdatedAt),
	}, nil
}

func (m *MedalService) DeleteMedal(ctx context.Context, req *pb.DeleteMedalRequest) (*pb.DeleteMedalResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}

	err = m.repo.DeleteMedalByID(ctx, id.String())
	if err != nil {
		return nil, err
	}

	return &pb.DeleteMedalResponse{Success: true}, nil
}

func (m *MedalService) GetMedal(ctx context.Context, req *pb.GetMedalRequest) (*pb.GetMedalResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}

	medal, err := m.repo.GetMedalByID(ctx, id.String())
	if err != nil {
		return nil, err
	}

	return &pb.GetMedalResponse{
		Medal: &pb.Medal{
			Id:          medal.ID.String(), // Convert uuid.UUID to string
			Description: medal.Description,
			AthleteId:   medal.AthleteID.String(),
			Type:        medal.Type,
			Country:     medal.Country,
			CreatedAt:   timestamppb.New(medal.CreatedAt),
			UpdatedAt:   timestamppb.New(medal.UpdatedAt),
		},
	}, nil
}

func (m *MedalService) GetMedals(ctx context.Context, req *pb.GetMedalsRequest) (*pb.GetMedalsResponse, error) {
	medals, err := m.repo.ListMedals(ctx, req.GetPage(), req.GetLimit(), req.GetSortOrder(), req.GetTypeFilter())
	if err != nil {
		return nil, err
	}

	pbMedals := make([]*pb.Medal, len(medals))
	for i, medal := range medals {
		pbMedals[i] = &pb.Medal{
			Id:          medal.ID.String(),
			Description: medal.Description,
			AthleteId:   medal.AthleteID.String(),
			Type:        medal.Type,
			Country:     medal.Country,
			CreatedAt:   timestamppb.New(medal.CreatedAt),
			UpdatedAt:   timestamppb.New(medal.UpdatedAt),
		}
	}

	return &pb.GetMedalsResponse{Medals: pbMedals}, nil
}

func (m *MedalService) GetMedalsByAthlete(ctx context.Context, req *pb.GetMedalsByAthleteRequest) (*pb.GetMedalsResponse, error) {
	athleteID, err := uuid.Parse(req.GetAthleteId())
	if err != nil {
		return nil, err
	}

	medals, err := m.repo.GetMedalsByAthlete(ctx, athleteID.String())
	if err != nil {
		return nil, err
	}

	pbMedals := make([]*pb.Medal, len(medals))
	for i, medal := range medals {
		pbMedals[i] = &pb.Medal{
			Id:          medal.ID.String(),
			Description: medal.Description,
			AthleteId:   medal.AthleteID.String(),
			Type:        medal.Type,
			Country:     medal.Country,
			CreatedAt:   timestamppb.New(medal.CreatedAt),
			UpdatedAt:   timestamppb.New(medal.UpdatedAt),
		}
	}

	return &pb.GetMedalsResponse{Medals: pbMedals}, nil
}

func (m *MedalService) GetMedalsByCountry(ctx context.Context, req *pb.GetMedalsByCountryRequest) (*pb.GetMedalsResponse, error) {
	medals, err := m.repo.GetMedalsByCountry(ctx, req.GetCountry())
	if err != nil {
		return nil, err
	}

	pbMedals := make([]*pb.Medal, len(medals))
	for i, medal := range medals {
		pbMedals[i] = &pb.Medal{
			Id:          medal.ID.String(),
			Description: medal.Description,
			AthleteId:   medal.AthleteID.String(),
			Type:        medal.Type,
			Country:     medal.Country,
			CreatedAt:   timestamppb.New(medal.CreatedAt),
			UpdatedAt:   timestamppb.New(medal.UpdatedAt),
		}
	}

	return &pb.GetMedalsResponse{Medals: pbMedals}, nil
}

func (m *MedalService) GetMedalsByTimeRange(ctx context.Context, req *pb.GetMedalsByTimeRangeRequest) (*pb.GetMedalsResponse, error) {
	startDate := req.StartTime.AsTime()
	endDate := req.EndTime.AsTime()
	medals, err := m.repo.GetMedalsByTimeRange(ctx, startDate, endDate, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}
	log.Println(medals)
	var medalResponses []*pb.Medal
	for _, medal := range medals {
		medalResponses = append(medalResponses, &pb.Medal{
			Id:          medal.ID.String(),
			Description: medal.Description,
			AthleteId:   medal.AthleteID.String(),
			Type:        medal.Type,
			Country:     medal.Country,
			CreatedAt:   timestamppb.New(medal.CreatedAt),
			UpdatedAt:   timestamppb.New(medal.UpdatedAt),
		})
	}

	return &pb.GetMedalsResponse{Medals: medalResponses}, nil
}

func (m *MedalService) UpdateMedal(ctx context.Context, req *pb.UpdateMedalRequest) (*pb.UpdateMedalResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}

	athletId, err := uuid.Parse(req.GetAthleteId())
	if err != nil {
		return nil, err
	}
	medal := &models.UpdateMedal{
		ID:          id,
		Description: req.GetDescription(),
		AthleteID:   athletId,
		Type:        req.GetType(),
		Country:     req.GetCountry(),
	}

	updatedMedal, err := m.repo.UpdateMedal(ctx, medal)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateMedalResponse{
		Medal: &pb.Medal{
			Id:          updatedMedal.ID.String(),
			Description: updatedMedal.Description,
			AthleteId:   updatedMedal.AthleteID.String(),
			Type:        updatedMedal.Type,
			Country:     updatedMedal.Country,
			CreatedAt:   timestamppb.New(updatedMedal.CreatedAt),
			UpdatedAt:   timestamppb.New(updatedMedal.UpdatedAt),
		},
	}, nil
}

func (m *MedalService) RankingByCountry(ctx context.Context, req *pb.GetRankingByCountryRequest) (*pb.GetRankingResponse, error) {
	limit, err := strconv.Atoi(req.Limit)
	if err != nil {
		return nil, err
	}

	ranking, err := m.countryMedals.GetTopCountries(ctx, limit)
	if err != nil {
		return nil, err
	}
	// log.Println(req)
	var res pb.GetRankingResponse
	for _, country := range ranking {
		var rank pb.RankingResponse
		rank.BronzeCount = int32(country.BronzeCount)
		rank.GoldCount = int32(country.GoldCount)
		rank.SilverCount = int32(country.SilverCount)
		rank.Name = country.Name
		res.Rankings = append(res.Rankings, &rank)
	}
	return &res, nil
}
