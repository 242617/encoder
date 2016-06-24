package main

import (
	"testing"
)

func TestBlackEncoding(t *testing.T) {
	img := load("test/black.png")
	res, _ := encode(img, 160)
	for _, v := range res {
		if v != 0 {
			t.Fatal("Incorrect black image encoding!")
		}
	}
}

func TestWhiteEncoding(t *testing.T) {
	img := load("test/white.png")
	res, _ := encode(img, 160)
	for _, v := range res {
		if v != 255 {
			t.Fatal("Incorrect white image encoding!")
		}
	}
}

func TestInterlacedEncoding(t *testing.T) {
	img := load("test/interlaced.png")
	res, _ := encode(img, 160)
	if res[0] != 85 {
		t.Fatal("Incorrect interlaced image encoding!")
	}
}

func TestWhiteAndBlackEncoding(t *testing.T) {
	img := load("test/white_and_black.png")
	res, _ := encode(img, 160)
	black, white := 0, 0
	for _, v := range res {
		if v > 127 {
			white++
		} else {
			black++
		}
	}
	if black != white {
		t.Fatal("Incorrect white and black image encoding!")
	}
}

func TestThresholdEncoding(t *testing.T) {
	img := load("test/gray.png")

	black, white := 0, 0
	res, _ := encode(img, 255)
	for _, v := range res {
		if v > 127 {
			white++
		} else {
			black++
		}
	}
	if black != 384 || white != 0 {
		t.Fatal("Incorrect threshold image encoding!")
	}

	black, white = 0, 0
	res, _ = encode(img, 1)
	for _, v := range res {
		if v > 127 {
			white++
		} else {
			black++
		}
	}
	if black != 0 || white != 384 {
		t.Fatal("Incorrect threshold image encoding!")
	}
}

func BenchmarkEncoding(b *testing.B) {
	img := load("test/white_and_black.png")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encode(img, 160)
	}
}
