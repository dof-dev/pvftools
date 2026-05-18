package world_drop

import "testing"

func TestRenderWorldDropDataKeepsHeaderAndItemOrder(t *testing.T) {
	mgr := &worldDropMgr{
		items: map[int][]*WorldDropItem{
			1: {
				{
					Id:    1001,
					Level: 1,
					Rate:  500,
				},
				{
					Id:    1002,
					Level: 1,
					Rate:  250,
				},
			},
		},
	}

	data := mgr.renderWorldDropData()
	want := []int{1, 0, 1001, 500, 1002, 250, -1}

	for i, v := range want {
		if data[i] != v {
			t.Fatalf("index %d: got %d, want %d", i, data[i], v)
		}
	}
}
