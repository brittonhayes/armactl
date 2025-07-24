package steam

// Server represents a steam game server
type Server struct {
	Name           string
	Game           string
	PlayerCount    uint8
	MaxPlayerCount uint8
}
