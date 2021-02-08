package avp_test

import (
	"testing"

	"github.com/nilsocket/avp"
)

var t1Formats = avp.Formats{
	{Resolution: 1080, AudioBitrate: 196},
	{Resolution: 4320},
	{Resolution: 2160},
	{Resolution: 720},
	{AudioBitrate: 48},
	{AudioBitrate: 316},
	{AudioBitrate: 128},
}

var t1Result = map[avp.Quality]avp.Formats{
	1: {
		{AudioBitrate: 316},
		{Resolution: 4320},
	},
	2: {
		{Resolution: 1080, AudioBitrate: 196},
	},
	3: {
		{AudioBitrate: 128},
		{Resolution: 720},
	},
	4: {
		{AudioBitrate: 128},
		{Resolution: 720},
	},
}

func TestAll(t *testing.T) {
	a := avp.New(t1Formats)
	check(t, a, t1Result)
}

func TestResolutionToInt(t *testing.T) {
	resolutionInput := []string{"1920x1080", "720p"}
	expectedOutput := []int{1080, 720}

	for i, ri := range resolutionInput {
		res, _ := avp.ResolutionToInt(ri)
		if expectedOutput[i] != res {
			t.Error("Expected:", expectedOutput[i], "Got:", res, "For:", ri)
		}
	}

}

func check(t *testing.T, got *avp.AVP, res map[avp.Quality]avp.Formats) {
	equal(t, got.Best(), res[1])
	equal(t, got.High(), res[2])
	equal(t, got.Medium(), res[3])
	equal(t, got.Low(), res[4])
}

func equal(t *testing.T, a, b avp.Formats) {
	if len(a) != len(b) {
		t.Error("Expected len(a) == len(b), but len(a)=", a, "len(b)=", b)
	}

	for i, aa := range a {
		if !fmtEqual(aa, b[i]) {
			t.Error("Expected:", b[i], "Got:", aa)
		}
	}
}

func fmtEqual(a, b *avp.Format) bool {
	// video
	if a.Resolution != b.Resolution || a.VideoCodec != b.VideoCodec || a.VideoBitrate != b.VideoBitrate || a.VideoHDR != b.VideoHDR || a.VideoHFR != b.VideoHDR {
		return false
	}

	if a.AudioBitrate != b.AudioBitrate || a.AudioChannels != b.AudioChannels || a.AudioCodec != b.AudioCodec || a.AudioVBR != b.AudioVBR {
		return false
	}
	return true
}
