// Package sgpd (Sample Group Description)
package sgpd

const (
	SGPD string = "sgpd"
)

type Box struct {
	GroupingType                  [4]byte
	DefaultLength                 uint32
	DefaultSampleDescriptionIndex uint32
	EntryCount                    uint32
	RollDistances                 []int16
	RollDistanceLens              []RollDistanceLengths
	AlternativeStartupEntries     []AlternativeStartupEntry
	AlternativeStartupEntriesLens []AlternativeStartupEntryWithLen
	VisualRandomAccessEntries     []VisualRandomAccessEntry
	VisualRandomAccessEntriesLens []VisualRandomAccessEntryWithLen
	TemporalLevelEntries          []TemporalLevelEntry
	TemporalLevelEntriesLens      []TemporalLevelEntryWithLen
	Unsupported                   []byte
}

type RollDistanceLengths struct {
	DescriptionLength uint32
	RollDistance      int16
}

type AlternativeStartupEntry struct {
	RollCount         uint16
	FirstOutputSample uint16
	SampleOffset      []uint32
	Opts              []AlternativeStartupEntryOpts
}

type AlternativeStartupEntryWithLen struct {
	AlternativeStartupEntry
	DescriptionLen uint32
}

type AlternativeStartupEntryOpts struct {
	OutputSampleLen uint16
	TotalSampleLen  uint16
}

type VisualRandomAccessEntry struct {
	LeadingSamplesKnown bool
	LeadingSamplesLen   uint8
}

type VisualRandomAccessEntryWithLen struct {
	VisualRandomAccessEntry
	DescriptionLength uint32
}

type TemporalLevelEntry struct {
	LevelIndependentlyDecodable bool
	_reserved                   uint8
}

type TemporalLevelEntryWithLen struct {
	TemporalLevelEntry
	DescriptionLength uint32
}

func (Box) Type() string {
	return SGPD
}
