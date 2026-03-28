package sequence

import "testing"

func TestDeterministic(t *testing.T) {
	left := New(42)
	right := New(42)
	for i := 0; i < 100; i++ {
		if got, want := left.Next(), right.Next(); got != want {
			t.Fatalf("record %d = %+v, want %+v", i, got, want)
		}
	}
}

func TestKnownSequence(t *testing.T) {
	generator := New(42)
	want := []uint64{
		13679457532755275413,
		2949826092126892291,
		5139283748462763858,
	}
	for i, expected := range want {
		record := generator.Next()
		if record.Index != uint64(i) || record.Value != expected {
			t.Fatalf("record %d = %+v, want value %d", i, record, expected)
		}
		if len(record.Checksum) != 16 {
			t.Fatalf("checksum %q has length %d, want 16", record.Checksum, len(record.Checksum))
		}
	}
}

func TestDifferentSeedsDiverge(t *testing.T) {
	if New(1).Next() == New(2).Next() {
		t.Fatal("different seeds produced the same first record")
	}
}
