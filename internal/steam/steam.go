package steam

import (
	"time"

	"github.com/rumblefrog/go-a2s"
)

type Client interface {
	Inspect() (*Server, error)
	Players() ([]string, error)
	Close() error
}

type steam struct {
	*a2s.Client
}

// New creates a new steam server readonly client.
func New(address string) Client {
	c, err := a2s.NewClient(address, a2s.TimeoutOption(5*time.Second))
	if err != nil {
		panic(err)
	}

	return &steam{
		Client: c,
	}
}

// Players returns a list of players on the server.
func (s *steam) Players() ([]string, error) {
	p, err := s.Client.QueryPlayer()
	if err != nil {
		return nil, err
	}

	var players []string
	for _, player := range p.Players {
		players = append(players, player.Name)
	}

	return players, nil
}

// Inspect returns information about the server.
func (s *steam) Inspect() (*Server, error) {
	i, err := s.Client.QueryInfo()
	if err != nil {
		return nil, err
	}

	return &Server{
		Name:           i.Name,
		Game:           i.Game,
		PlayerCount:    i.Players,
		MaxPlayerCount: i.MaxPlayers,
	}, nil
}
