package main

import (
	"os"
)

func segmentGhcEnv(p *powerline) {
	var ghcEnv string
		ghcEnv, _ = os.LookupEnv("GHC_ENVIRONMENT")
	if ghcEnv == "" {
		return
	} else {
		p.appendSegment("ghc-env", segment{
            // We use the constant string "ghc-env" because the content
            // of the env var is usually a long path
			content:    "ghc-env",
			foreground: p.theme.GhcEnvFg,
			background: p.theme.GhcEnvBg,
		})
	}
}
