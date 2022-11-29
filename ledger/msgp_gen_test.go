//go:build !skip_msgp_testing
// +build !skip_msgp_testing

package ledger

// Code generated by github.com/algorand/msgp DO NOT EDIT.

import (
	"testing"

	"github.com/algorand/msgp/msgp"

	"github.com/algorand/go-algorand/protocol"
	"github.com/algorand/go-algorand/test/partitiontest"
)

func TestMarshalUnmarshalCatchpointFileHeader(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := CatchpointFileHeader{}
	bts := v.MarshalMsg(nil)
	left, err := v.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func TestRandomizedEncodingCatchpointFileHeader(t *testing.T) {
	protocol.RunEncodingTest(t, &CatchpointFileHeader{})
}

func BenchmarkMarshalMsgCatchpointFileHeader(b *testing.B) {
	v := CatchpointFileHeader{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgCatchpointFileHeader(b *testing.B) {
	v := CatchpointFileHeader{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalCatchpointFileHeader(b *testing.B) {
	v := CatchpointFileHeader{}
	bts := v.MarshalMsg(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestMarshalUnmarshalcatchpointFileBalancesChunkV5(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := catchpointFileBalancesChunkV5{}
	bts := v.MarshalMsg(nil)
	left, err := v.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func TestRandomizedEncodingcatchpointFileBalancesChunkV5(t *testing.T) {
	protocol.RunEncodingTest(t, &catchpointFileBalancesChunkV5{})
}

func BenchmarkMarshalMsgcatchpointFileBalancesChunkV5(b *testing.B) {
	v := catchpointFileBalancesChunkV5{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgcatchpointFileBalancesChunkV5(b *testing.B) {
	v := catchpointFileBalancesChunkV5{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalcatchpointFileBalancesChunkV5(b *testing.B) {
	v := catchpointFileBalancesChunkV5{}
	bts := v.MarshalMsg(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestMarshalUnmarshalcatchpointFileChunkV6(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := catchpointFileChunkV6{}
	bts := v.MarshalMsg(nil)
	left, err := v.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func TestRandomizedEncodingcatchpointFileChunkV6(t *testing.T) {
	protocol.RunEncodingTest(t, &catchpointFileChunkV6{})
}

func BenchmarkMarshalMsgcatchpointFileChunkV6(b *testing.B) {
	v := catchpointFileChunkV6{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgcatchpointFileChunkV6(b *testing.B) {
	v := catchpointFileChunkV6{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalcatchpointFileChunkV6(b *testing.B) {
	v := catchpointFileChunkV6{}
	bts := v.MarshalMsg(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestMarshalUnmarshalcatchpointFirstStageInfo(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := catchpointFirstStageInfo{}
	bts := v.MarshalMsg(nil)
	left, err := v.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func TestRandomizedEncodingcatchpointFirstStageInfo(t *testing.T) {
	protocol.RunEncodingTest(t, &catchpointFirstStageInfo{})
}

func BenchmarkMarshalMsgcatchpointFirstStageInfo(b *testing.B) {
	v := catchpointFirstStageInfo{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgcatchpointFirstStageInfo(b *testing.B) {
	v := catchpointFirstStageInfo{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalcatchpointFirstStageInfo(b *testing.B) {
	v := catchpointFirstStageInfo{}
	bts := v.MarshalMsg(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestMarshalUnmarshalencodedBalanceRecordV5(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := encodedBalanceRecordV5{}
	bts := v.MarshalMsg(nil)
	left, err := v.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func TestRandomizedEncodingencodedBalanceRecordV5(t *testing.T) {
	protocol.RunEncodingTest(t, &encodedBalanceRecordV5{})
}

func BenchmarkMarshalMsgencodedBalanceRecordV5(b *testing.B) {
	v := encodedBalanceRecordV5{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgencodedBalanceRecordV5(b *testing.B) {
	v := encodedBalanceRecordV5{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalencodedBalanceRecordV5(b *testing.B) {
	v := encodedBalanceRecordV5{}
	bts := v.MarshalMsg(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestMarshalUnmarshalencodedBalanceRecordV6(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := encodedBalanceRecordV6{}
	bts := v.MarshalMsg(nil)
	left, err := v.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func TestRandomizedEncodingencodedBalanceRecordV6(t *testing.T) {
	protocol.RunEncodingTest(t, &encodedBalanceRecordV6{})
}

func BenchmarkMarshalMsgencodedBalanceRecordV6(b *testing.B) {
	v := encodedBalanceRecordV6{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgencodedBalanceRecordV6(b *testing.B) {
	v := encodedBalanceRecordV6{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalencodedBalanceRecordV6(b *testing.B) {
	v := encodedBalanceRecordV6{}
	bts := v.MarshalMsg(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestMarshalUnmarshalencodedKVRecordV6(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := encodedKVRecordV6{}
	bts := v.MarshalMsg(nil)
	left, err := v.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func TestRandomizedEncodingencodedKVRecordV6(t *testing.T) {
	protocol.RunEncodingTest(t, &encodedKVRecordV6{})
}

func BenchmarkMarshalMsgencodedKVRecordV6(b *testing.B) {
	v := encodedKVRecordV6{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgencodedKVRecordV6(b *testing.B) {
	v := encodedKVRecordV6{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalencodedKVRecordV6(b *testing.B) {
	v := encodedKVRecordV6{}
	bts := v.MarshalMsg(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}
