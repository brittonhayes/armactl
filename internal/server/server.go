package server

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"
)

type ARMA struct {
	Binary                 string
	CDLC                   string
	Config                 string
	Port                   uint16
	SkipInstall            bool
	ModsLocal              []string
	ModsServer             []string
	World                  string
	LimitFPS               uint8
	HeadlessClients        uint8
	HeadlessClientProfile  string
	HeadlessClientPassword string
	HeadlessClientServer   string
	Params                 string
	Profile                string
	SteamBranch            string
	SteamBranchPassword    string
	SteamUser              string
	SteamPassword          string
	DryRun                 bool
}

type Server interface {
	Update(ctx context.Context, w io.Writer) error
	Start(ctx context.Context, w io.Writer) error
	Stop(ctx context.Context) error
	StartHeadless(ctx context.Context, w io.Writer) error
	StopHeadless(ctx context.Context) error
}

func New(options ...ServerOption) Server {
	s := &ARMA{}
	for _, option := range options {
		option(s)
	}

	return s
}

func (s *ARMA) args() string {
	return fmt.Sprintf("-limitFPS=%d -world=%q -mod=%q -config=%q -port=%d -name=%q -profiles=%q -serverMod=%q",
		s.LimitFPS,
		s.World,
		strings.Join(s.ModsLocal, ";"),
		s.Config,
		s.Port,
		s.Profile,
		s.Profile,
		strings.Join(s.ModsServer, ";"),
	)
}

func (s *ARMA) argsHeadless() string {
	return fmt.Sprintf("-limitFPS=%d -world=%q -mod=%q -config=%q -port=%d -profiles=%q -serverMod=%q -client -connect=%s -password=%q -name=%q",
		s.LimitFPS,
		s.World,
		strings.Join(s.ModsLocal, ";"),
		s.Config,
		s.Port,
		s.Profile,
		strings.Join(s.ModsServer, ";"),
		s.HeadlessClientServer,
		s.HeadlessClientPassword,
		s.HeadlessClientProfile,
	)
}

func (s *ARMA) Start(ctx context.Context, w io.Writer) error {
	args := s.args()
	fmt.Println(s.Binary, args)
	if s.DryRun {
		return nil
	}

	cmd := exec.CommandContext(ctx, s.Binary, args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(w, cmd.Process.Kill())
		return err
	}

	fmt.Fprintln(w, string(out))

	return nil
}

func (s *ARMA) Stop(ctx context.Context) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	return nil
}

func (s *ARMA) StartHeadless(ctx context.Context, w io.Writer) error {
	var wg sync.WaitGroup

	for i := 0; i < int(s.HeadlessClients); i++ {
		wg.Add(1)
		go func(ctx context.Context, w io.Writer, s *ARMA) {
			args := s.argsHeadless()
			fmt.Fprintln(w, s.Binary, args)
			if s.DryRun {
				wg.Done()
			}

			cmd := exec.CommandContext(ctx, s.Binary, args)
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Fprintln(w, err)
				fmt.Fprintln(w, cmd.Process.Kill())
				wg.Done()
			}

			fmt.Fprintln(w, string(out))
			wg.Done()
		}(ctx, w, s)
	}
	wg.Wait()

	return nil
}

func (s *ARMA) StopHeadless(ctx context.Context) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	return nil
}

func (s *ARMA) Update(ctx context.Context, w io.Writer) error {
	cmd := exec.CommandContext(ctx, "steamcmd.exe", "+login", s.SteamUser, s.SteamPassword, "+force_install_dir", "/arma3", "+app_update", "233780", "validate", "+quit")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Fprintln(w, string(out))

	return nil
}
