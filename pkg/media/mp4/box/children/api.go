package children

import "github.com/jwhittle933/streamline/pkg/media/mp4/box"

type BoxFactory func(*box.Info) box.Boxed

type Registry map[string]BoxFactory
