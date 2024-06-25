package api

import (
	"github.com/ltcong1411/rapid-loyalty/pb"
	"github.com/ltcong1411/rapid-loyalty/utils"
)

// Server serves api requests for our rapid loyalty service.
type Server struct {
	pb.UnimplementedRapidLoyaltyServer
	config utils.Config
}

// NewServer creates a new api server and set up routing.
func NewServer(config utils.Config) (*Server, error) {
	server := &Server{
		config: config,
	}

	return server, nil
}
