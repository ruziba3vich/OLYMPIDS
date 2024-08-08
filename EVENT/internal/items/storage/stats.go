package storage

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	genproto "github.com/ruziba3vich/OLYMPIDS/EVENT/genproto/stats"
	"github.com/ruziba3vich/OLYMPIDS/EVENT/internal/items/redisservice"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	StatisticsStorage struct {
		mongosh *DB
		logger  *log.Logger
		events  chan *genproto.Event
		redisCl *redisservice.RedisService
		genproto.StatsServiceServer
	}
)

func NewStatsStorage(mongosh *DB, events chan *genproto.Event, redisCl *redisservice.RedisService, logger *log.Logger) *StatisticsStorage {
	return &StatisticsStorage{
		mongosh: mongosh,
		logger:  logger,
		redisCl: redisCl,
		events:  events,
	}
}

func (s *StatisticsStorage) CreateTeamStats(ctx context.Context, req *genproto.TeamEvent) (*genproto.TeamEvent, error) {
	_, err := s.mongosh.StatsCollection.InsertOne(ctx, req)
	if err != nil {
		s.logger.Println("Error inserting TeamEvent:", err)
		return nil, err
	}
	value := genproto.Event{
		EventId:   uuid.New().String(),
		TeamEvent: req,
		Finished:  false,
	}
	s.events <- &value
	return req, nil
}

func (s *StatisticsStorage) CreatePlayerOnlyStats(ctx context.Context, req *genproto.PlayerOnly) (*genproto.PlayerOnly, error) {
	_, err := s.mongosh.StatsCollection.InsertOne(ctx, req)
	if err != nil {
		s.logger.Println("Error inserting PlayerOnly:", err)
		return nil, err
	}
	value := genproto.Event{
		EventId:    uuid.New().String(),
		PlayerOnly: req,
		Finished:   false,
	}
	s.events <- &value
	return req, nil
}

func (s *StatisticsStorage) CreateRaceStats(ctx context.Context, req *genproto.Race) (*genproto.Race, error) {
	_, err := s.mongosh.StatsCollection.InsertOne(ctx, req)
	if err != nil {
		s.logger.Println("Error inserting Race:", err)
		return nil, err
	}
	value := genproto.Event{
		EventId:   uuid.New().String(),
		RaceEvent: req,
		Finished:  false,
	}
	s.events <- &value
	return req, nil
}

func (s *StatisticsStorage) UpdateTeamStats(ctx context.Context, req *genproto.UpdateTeamStatsRequest) (*genproto.Team, error) {
	event, err := s.getStatsById(req.EventId)
	if err != nil {
		return nil, err
	}
	
	filter := bson.M{"team_id": req.TeamId}
	if event.TeamEvent.Team1Id == req.TeamId {
		filter["set"] = event.TeamEvent.Team1
		if req.Update != nil {
			event.TeamEvent.Team1.Updates = append(event.TeamEvent.Team1.Updates, req.Update)
		}
		if req.Card != nil {
			event.TeamEvent.Team1.Cards = append(event.TeamEvent.Team1.Cards, req.Card)
		}
		if req.Score != nil {
			event.TeamEvent.Team1.Scores = append(event.TeamEvent.Team1.Scores, req.Score)
		}
	} else {
		filter["set"] = event.TeamEvent.Team2
		if req.Update != nil {
			event.TeamEvent.Team2.Updates = append(event.TeamEvent.Team2.Updates, req.Update)
		}
		if req.Card != nil {
			event.TeamEvent.Team2.Cards = append(event.TeamEvent.Team2.Cards, req.Card)
		}
		if req.Score != nil {
			event.TeamEvent.Team2.Scores = append(event.TeamEvent.Team2.Scores, req.Score)
		}
	}

	update := bson.M{
		"$set": event,
	}
	result := s.mongosh.StatsCollection.FindOneAndUpdate(ctx, filter, update)
	if result.Err() != nil {
		s.logger.Println("Error updating TeamStats:", result.Err())
		return nil, result.Err()
	}

	var updatedTeam genproto.Team
	if err := result.Decode(&updatedTeam); err != nil {
		s.logger.Println("Error decoding updated Team:", err)
		return nil, err
	}

	s.events <- event
	return &updatedTeam, nil
}

func (s *StatisticsStorage) GetStatsByEventId(req *genproto.GetStatsByEventIdRequest, stream genproto.StatsService_GetStatsByEventIdServer) error {

	if err := s.redisCl.GetEventById(stream.Context(), req.EventId); err != nil {
		return err
	}

	event, err := s.getStatsById(req.EventId)
	if err != nil {
		return err
	}
	if err := stream.Send(event); err != nil {
		return err
	}

	for {
		select {
		case event = <-s.events:
			if event.Finished {
				return errors.New("match is not live")
			}
			if event.EventId == req.EventId {
				if err := stream.Send(event); err != nil {
					return err
				}
			} else {
				s.events <- event
			}
		}
	}
}

func (s *StatisticsStorage) getStatsById(statId string) (*genproto.Event, error) {
	filter := bson.M{"event_id": statId}
	cursor := s.mongosh.StatsCollection.FindOne(context.Background(), filter)
	var event genproto.Event
	if err := cursor.Decode(&event); err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return &event, nil
}

// func (s *StatisticsStorage) ListenAndServe() {
// 	for {
// 		select {
// 		case event := <-s.events:
// 			if err := en.stream.Send(event); err != nil {
// 				log.Println("Error sending event:", err)
// 				return
// 			}
// 		case <-t.C:
// 			return
// 		}
// 	}
// }
