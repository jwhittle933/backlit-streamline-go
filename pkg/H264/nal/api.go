// package nal, for nal (Network Abstraction Layer)
// See https://yumichan.net/video-processing/video-compression/introduction-to-h264-nal-unit/
// And https://yumichan.net/video-processing/video-compression/explanation-of-descriptors-in-itu-t-publication-on-h-264-coding-standardrecommendation-with-example/
package nal

// First byte of nal unit: nal unit header
// Rest are the payload

const (
	NALU_TYPE_UNSPECIFIED byte = iota
	NALU_TYPE_SLICE_NON_IDR
	NALU_TYPE_DATA_PARTITION_A
	NALU_TYPE_DATA_PARTITION_B
	NALU_TYPE_DATA_PARTITION_C
	NALU_TYPE_SLICE_IDR
	NALU_TYPE_SUPPLEMENTAL_ENHANCEMENT_INFO
	NALU_TYPE_SEQUENCE_PARAMETER_SET
	NALU_TYPE_PICTURE_PARAMETER_SET
	NALU_TYPE_ACCESS_UNIT_DELIMITER
	NALU_TYPE_END_OF_SEQ
	NALU_TYPE_END_OF_STREAM
	NALU_TYPE_FILLER
	NALU_TYPE_SEQ_PARAM_SET_EXT
	NALU_TYPE_PREFIX
	NALU_TYPE_SUBSET_SEQ_PARAM_SET
	_ // 16-18 reserved
	_
	_
	NALU_TYPE_SLICE_AUX_CODED_PICTURE_WITHOUT_PARTITION
	NALU_TYPE_SLICE_EXT
	NALU_TYPE_SLICE_EXT_DEPTH_VIEW
	_ // 22-23 reserved
	_
	// 24-31 unspecified
)

type UnitHeader byte

type Unit struct {
	// forbidden_zero_bit (1)
	// nal_ref_idc (3)
	// nal_unit_type (5)
	header  UnitHeader
	payload []byte
}

// VCL units are those which have a nal_unit_type
// equal to 1 - 5, inclusive. All remaining are
// considered non-VCL
// This follows Annex A
func (uh *UnitHeader) isVCL() bool {
	return false
}

// If a non-reference field/frame/picture unit, nal_ref_idc is 0
func (uh *UnitHeader) isReference() bool {
	return false
}
