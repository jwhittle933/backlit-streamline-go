// package nal, for NAL (Network Abstraction Layer)
// See https://yumichan.net/video-processing/video-compression/introduction-to-h264-nal-unit/
// And https://yumichan.net/video-processing/video-compression/explanation-of-descriptors-in-itu-t-publication-on-h-264-coding-standardrecommendation-with-example/
package nal

// First byte of NAL unit: NAL unit header
// Rest are the payload

type Unit struct {
	header byte
}
