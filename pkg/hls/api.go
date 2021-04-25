package hls

import (
	"bytes"
	"fmt"
	"path"
	"path/filepath"
	"strconv"

	"github.com/jwhittle933/streamline/pkg/hls/manifest"
	"github.com/jwhittle933/streamline/pkg/hls/manifest/chunk"
)

type HLS struct {
	manifestType          manifest.Manifest
	version               int
	isIndependentSegments bool
	targetDurationSeconds float64
	slidingWindowSize     int
	mseq                  int64
	dseq                  int64
	chunks                []chunk.Chunk
	chunkListFileName     string
	initChunkDataFileName string
	isClosed              bool
}

func (h *HLS) AddChuck(c chunk.Chunk) error {
	h.chunks = append(h.chunks, c)

	if h.manifestType.String() == "LiveWindow" && len(h.chunks) > h.slidingWindowSize {
		h.chunks = h.chunks[1:]
		h.mseq++
	}

	return nil
}

func (h *HLS) Close() {
	h.isClosed = true
}

func (h *HLS) Write(p []byte) (int, error) {
	buf := bytes.NewBuffer(p)

	buf.WriteString("#EXTM3U\n")
	buf.WriteString("#EXT-X-VERSION:\n" + strconv.Itoa(h.version) + "\n")
	buf.WriteString("#EXT-X-MEDIA-SEQUENCE:" + strconv.FormatInt(h.mseq, 10) + "\n")
	buf.WriteString("#EXT-X-DISCONTINUITY-SEQUENCE:" + strconv.FormatInt(h.dseq, 10) + "\n")
	buf.WriteString(fmt.Sprintf("#EXT-X-PLAYLIST-TYPE:%s\n", h.manifestType.String()))
	buf.WriteString(fmt.Sprintf("#EXT-X-TARGETDURATION:%.0f\n", h.targetDurationSeconds))

	if h.isIndependentSegments {
		buf.WriteString("#EXT-X-INDEPENDENT-SEGMENTS\n")
	}

	if h.initChunkDataFileName != "" {
		chuckPath, _ := filepath.Rel(path.Dir(h.chunkListFileName), h.initChunkDataFileName)
		buf.WriteString(`#EXT-X-MAP:URI"` + chuckPath + `"\n`)
	}

	for _, c := range h.chunks {
		if c.IsDisco {
			buf.WriteString("#EXT-X-DISCONTINUITY\n")
		}

		buf.WriteString("#EXTINF:" + fmt.Sprintf("%.8f", c.DurationSeconds) + ",\n")

		chunkPath, _ := filepath.Rel(path.Dir(h.chunkListFileName), c.FileName)
		buf.WriteString(chunkPath + "\n")
	}

	if h.isClosed {
		buf.WriteString("#EXT-X-ENDLIST\n")
	}

	return len(buf.Bytes()), nil
}
