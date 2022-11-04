package test

import (
	"encoding/json"
	"os"
	"testing"
)

type FileDataSample struct {
	Title string
}

func TestFileCreate(t *testing.T) {
	f, err := os.Create("../data/fileCreateTest")
	if err != nil {
		t.Error(err.Error())
	}
	defer f.Close()
	t.Log("파일이 성공적으로 생성됬습니다")

	s, err := json.Marshal(FileDataSample{Title: "생성이 잘된다!"})
	if err != nil {
		t.Error("마샬 실패")
	}
	f.WriteString(string(s))
}

func TestFileRead(t *testing.T) {
	_, err := os.Stat("../data/fileCreateTest")
	if err != nil {
		t.Log("파일 정보 불러오기 실패, 새로생성합니다")
		TestFileCreate(t)
	}

	fr, err := os.ReadFile("../data/fileCreateTest")
	if err != nil {
		t.Error("읽기 실패헀어요")
	}
	t.Log(string(fr))
	t.Log("파일 내용을 go의 스트럭쳐에 맞게 변환합니다")

	temp := new(map[string]string)
	json.Unmarshal(fr, &temp)
	t.Log(*temp)
}

func TestFileDelete(t *testing.T) {
	_, err := os.Stat("../data/fileCreateTest")
	if err != nil {
		t.Log("파일이 없습니다, 생성합니다")
		TestFileCreate(t)
	}

	t.Log("파일을 삭제 합니다")
	os.Remove("../data/fileCreateTest")

}
