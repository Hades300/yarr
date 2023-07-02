//go:build !windows && !macos
// +build !windows,!macos

package platform

import (
	"github.com/hades300/yarr/src/server"
)

func Start(s *server.Server) {
	s.Start()
}
