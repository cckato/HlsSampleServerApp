package utils

import (
	"testing"
)

func TestGetRequestIndex(t *testing.T) {
	n1 := GetRequestIndex("index=31584")
	assertEqualInt(t, 31584, n1)

	n2 := GetRequestIndex("")
	assertEqualInt(t, 0, n2)

	n3 := GetRequestIndex("hoge=huga&index=410")
	assertEqualInt(t, 410, n3)

	n4 := GetRequestIndex("index=6&hige=hage")
	assertEqualInt(t, 6, n4)

	n5 := GetRequestIndex("index=u29e3")
	assertEqualInt(t, 0, n5)
}

func assertEqualInt(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("Expected [%d], but was [%d]", expected, actual)
	} else {
		t.Log("パスしたよ")
	}
}


func TestIsExtinf(t *testing.T) {
	if IsExtinf("#EXTM3U") {
		t.Errorf("Expected false, target:%s", "#EXTM3U")
	}
	if IsExtinf("#EXT-X-VERSION:3") {
		t.Errorf("Expected false, target:%s", "#EXT-X-VERSION:3")
	}
	if IsExtinf("#EXTINF:14.000000,") == false {
		t.Errorf("Expected true, target:%s", "#EXTINF:14.000000,")
	}
}

func TestIsTsFile(t *testing.T) {
	if IsTsFile("#EXTM3U") {
		t.Errorf("Expected false, target:%s", "#EXTM3U")
	}
	if IsTsFile("#EXTINF:14.000000,") {
		t.Errorf("Expected false, target%s", "#EXTINF:14.000000,")
	}
	if IsTsFile("hogets") {
		t.Errorf("Expected false, target%s", "hogets")
	}
	if IsTsFile("http://hoge.com:8080/kin/moza_000.ts") == false {
		t.Errorf("Expected true, target%s", "http://hoge.com:8080/kin/moza_000.ts")
	}
}
