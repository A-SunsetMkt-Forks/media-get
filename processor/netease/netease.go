package netease

import (
	"strings"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/meta"
)

type Core struct {
	Opts *args.Options
}

const Album = "Netease"

func (c *Core) FetchMetaAndResourceInfo() (mediaMeta *meta.MediaMeta, audios []*meta.Audio, err error) {
	if strings.Contains(c.Opts.Url, "/program") {
		return c.fetchFromProgram()
	}
	if strings.Contains(c.Opts.Url, "/song") {
		return c.fetchFromSong()
	}
	if strings.Contains(c.Opts.Url, "/mv") || strings.Contains(c.Opts.Url, "/video") {
		return c.fetchFromMV()
	}
	return nil, nil, err
}
