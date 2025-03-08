package utils

import "testing"

func TestGenerateSnowflakeID(t *testing.T) {
	worker, _ := NewSnowFlakeWorker(1)
	for i := range 20 {
		id, _ := worker.GenerateId()
		t.Log(i, " ", id)
	}
}
