package stats

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/genproto/stats"
)

type (
	StatsClient struct {
		client stats.StatsServiceClient
		logger *log.Logger
	}
)

func NewStatClient(client stats.StatsServiceClient, logger *log.Logger) *StatsClient {
	return &StatsClient{
		client: client,
		logger: logger,
	}
}

func (s *StatsClient) CreateStatsHandler(c *gin.Context) {
	var req stats.CreateStatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		s.logger.Println("ERROR WHILE BINDING THE OBJECT :", err.Error())
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	switch req.EventType {
		case "team"
	}
}

// func (s *StatsClient) CreateTeamStats(ctx context.Context, in *stats.TeamEvent, opts ...grpc.CallOption) (*stats.TeamEvent, error)
// func (s *StatsClient) CreatePlayerOnlyStats(ctx context.Context, in *stats.PlayerOnly, opts ...grpc.CallOption) (*stats.PlayerOnly, error)
// func (s *StatsClient) CreateRaceStats(ctx context.Context, in *stats.Race, opts ...grpc.CallOption) (*stats.Race, error)
// func (s *StatsClient) UpdateTeamStats(ctx context.Context, in *stats.UpdateTeamStatsRequest, opts ...grpc.CallOption) (*stats.Team, error)
// func (s *StatsClient) GetStatsByEventId(ctx context.Context, in *stats.GetStatsByEventIdRequest, opts ...grpc.CallOption) (stats.StatsService_GetStatsByEventIdClient, error)
