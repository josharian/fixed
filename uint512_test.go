// Code generated by 'gen'. DO NOT EDIT.

package fixed

import (
	"math/big"
	"math/bits"
	"testing"

	"golang.org/x/exp/rand"
)

var (
	big512mask    = bigMask(512)
	bigMaxUint512 = Uint512{}.max().big()
)

func randUint512() Uint512 {
	var x Uint512
	if randBool() {
		x.u0 = randUint64()
	}
	if randBool() {
		x.u1 = randUint64()
	}
	if randBool() {
		x.u2 = randUint64()
	}
	if randBool() {
		x.u3 = randUint64()
	}
	if randBool() {
		x.u4 = randUint64()
	}
	if randBool() {
		x.u5 = randUint64()
	}
	if randBool() {
		x.u6 = randUint64()
	}
	if randBool() {
		x.u7 = randUint64()
	}
	return x
}

func (x Uint512) big() *big.Int {
	var v big.Int
	if bits.UintSize == 32 {
		v.SetBits([]big.Word{
			big.Word(x.u0),
			big.Word(x.u0 >> 32),
			big.Word(x.u1),
			big.Word(x.u1 >> 32),
			big.Word(x.u2),
			big.Word(x.u2 >> 32),
			big.Word(x.u3),
			big.Word(x.u3 >> 32),
			big.Word(x.u4),
			big.Word(x.u4 >> 32),
			big.Word(x.u5),
			big.Word(x.u5 >> 32),
			big.Word(x.u6),
			big.Word(x.u6 >> 32),
			big.Word(x.u7),
			big.Word(x.u7 >> 32),
		})
	} else {
		v.SetBits([]big.Word{
			big.Word(x.u0),
			big.Word(x.u1),
			big.Word(x.u2),
			big.Word(x.u3),
			big.Word(x.u4),
			big.Word(x.u5),
			big.Word(x.u6),
			big.Word(x.u7),
		})
	}
	return &v
}

func TestUint512BitLen(t *testing.T) {
	for i := 0; i < 250_000; i++ {
		x := randUint512()

		got := x.BitLen()
		want := x.big().BitLen()
		if got != want {
			t.Fatalf("expected %d, got %d", want, got)
		}
	}
}

func TestUint512LeadingZeros(t *testing.T) {
	for i := 0; i < 250_000; i++ {
		x := randUint512()

		got := x.LeadingZeros()
		want := 512 - x.big().BitLen()
		if got != want {
			t.Fatalf("expected %d, got %d", want, got)
		}
	}
}

func TestUint512Cmp(t *testing.T) {
	for i := 0; i < 100_000; i++ {
		x := randUint512()
		y := randUint512()

		got := x.Cmp(y)
		want := x.big().Cmp(y.big())
		if got != want {
			t.Fatalf("Cmp(%d, %d): expected %d, got %d",
				x.big(), y.big(), want, got)
		}
	}
}

func TestUint512And(t *testing.T) {
	for i := 0; i < 100_000; i++ {
		x := randUint512()
		y := randUint512()

		z := x.And(y)

		want := new(big.Int).And(x.big(), y.big())
		want.And(want, big512mask)

		if got := z.big(); got.Cmp(want) != 0 {
			t.Fatalf("%d + %d: expected %d, got %d",
				x.big(), y.big(), want, got)
		}
	}
}

func TestUint512Or(t *testing.T) {
	for i := 0; i < 100_000; i++ {
		x := randUint512()
		y := randUint512()

		z := x.Or(y)

		want := new(big.Int).Or(x.big(), y.big())
		want.And(want, big512mask)

		if got := z.big(); got.Cmp(want) != 0 {
			t.Fatalf("%d + %d: expected %d, got %d",
				x.big(), y.big(), want, got)
		}
	}
}

func TestUint512Xor(t *testing.T) {
	for i := 0; i < 100_000; i++ {
		x := randUint512()
		y := randUint512()

		z := x.Xor(y)

		want := new(big.Int).Xor(x.big(), y.big())
		want.And(want, big512mask)

		if got := z.big(); got.Cmp(want) != 0 {
			t.Fatalf("%d + %d: expected %d, got %d",
				x.big(), y.big(), want, got)
		}
	}
}

func TestUint512Lsh(t *testing.T) {
	for i := 0; i < 1_000_000; i++ {
		x := randUint512()
		n := uint(rand.Intn(512 + 1))

		z := x.Lsh(n)

		want := new(big.Int).Lsh(x.big(), n)
		want.And(want, big512mask)

		if got := z.big(); got.Cmp(want) != 0 {
			t.Fatalf("%d << %d: expected %d, got %d",
				x.big(), n, want, got)
		}
	}
}

func TestUint512Rsh(t *testing.T) {
	for i := 0; i < 1_000_000; i++ {
		x := randUint512()
		n := uint(rand.Intn(512 + 1))

		z := x.Rsh(n)

		want := new(big.Int).Rsh(x.big(), n)
		want.And(want, big512mask)

		if got := z.big(); got.Cmp(want) != 0 {
			t.Fatalf("%d >> %d: expected %d, got %d",
				x.big(), n, want, got)
		}
	}
}

func TestUint512Add(t *testing.T) {
	for i := 0; i < 100_000; i++ {
		x := randUint512()
		y := randUint512()

		z, c := x.AddCheck(y)

		want := new(big.Int).Add(x.big(), y.big())
		if carry := want.Cmp(bigMaxUint512) > 0; carry != (c == 1) {
			t.Fatalf("%d + %d: expected %t, got %t",
				x.big(), y.big(), carry, c == 1)
		}
		want.And(want, big512mask)

		if c == 0 && x.Add(y) != z {
			t.Fatalf("%d: %d * %d: %d != %d",
				i, x.big(), y.big(), x.Add(y), z)
		}
		if got := z.big(); got.Cmp(want) != 0 {
			t.Fatalf("%d + %d: expected %d, got %d",
				x.big(), y.big(), want, got)
		}
	}
}

func TestUint512Add64(t *testing.T) {
	for i := 0; i < 250_000; i++ {
		x := randUint512()
		y := randUint64()

		z, c := x.addCheck64(y)

		ybig := new(big.Int).SetUint64(y)
		want := new(big.Int).Add(x.big(), ybig)
		if carry := want.Cmp(bigMaxUint512) > 0; carry != (c == 1) {
			t.Fatalf("%d + %d: expected %t, got %t",
				x.big(), ybig, carry, c == 1)
		}
		want.And(want, big512mask)

		if c == 0 && x.add64(y) != z {
			t.Fatalf("%d: %d * %d: %d != %d",
				i, x.big(), ybig, x.add64(y), z)
		}
		if got := z.big(); got.Cmp(want) != 0 {
			t.Fatalf("%d + %d: expected %d, got %d",
				x.big(), ybig, want, got)
		}
	}
}

func TestUint512Sub(t *testing.T) {
	for i := 0; i < 100_000; i++ {
		x := randUint512()
		y := randUint512()

		z, b := x.SubCheck(y)

		want := new(big.Int).Sub(x.big(), y.big())
		if borrow := want.Sign() < 0; borrow != (b == 1) {
			t.Fatalf("%d - %d: expected %t, got %t",
				x.big(), y.big(), borrow, b == 1)
		}
		want.And(want, big512mask)

		if b == 0 && x.Sub(y) != z {
			t.Fatalf("%d: %d * %d: %d != %d",
				i, x.big(), y.big(), x.Sub(y), z)
		}
		if got := z.big(); got.Cmp(want) != 0 {
			t.Fatalf("%d - %d: expected %d, got %d",
				x.big(), y.big(), want, got)
		}
	}
}

func TestUint512Sub64(t *testing.T) {
	for i := 0; i < 250_000; i++ {
		x := randUint512()
		y := randUint64()

		z, b := x.subCheck64(y)

		ybig := new(big.Int).SetUint64(y)
		want := new(big.Int).Sub(x.big(), ybig)
		if borrow := want.Sign() < 0; borrow != (b == 1) {
			t.Fatalf("%d - %d: expected %t, got %t",
				x.big(), ybig, borrow, b == 1)
		}
		want.And(want, big512mask)

		if b == 0 && x.sub64(y) != z {
			t.Fatalf("%d: %d * %d: %d != %d",
				i, x.big(), ybig, x.sub64(y), z)
		}
		if got := z.big(); got.Cmp(want) != 0 {
			t.Fatalf("%d - %d: expected %d, got %d",
				x.big(), ybig, want, got)
		}
	}
}

func TestUint512Mul(t *testing.T) {
	for i := 0; i < 100_000; i++ {
		x := randUint512()
		y := randUint512()

		z, ok := x.MulCheck(y)

		want := new(big.Int).Mul(x.big(), y.big())
		if (want.Cmp(bigMaxUint512) <= 0) != ok {
			t.Fatalf("%d: %d * %d: expected %t",
				i, x.big(), y.big(), !ok)
		}
		want.And(want, big512mask)

		if ok && x.Mul(y) != z {
			t.Fatalf("%d: %d * %d: %d != %d",
				i, x.big(), y.big(), x.Mul(y), z)
		}
		z = x.Mul(y)
		if got := z.big(); got.Cmp(want) != 0 {
			t.Fatalf("%d: %d * %d: expected %d, got %d",
				i, x.big(), y.big(), want, got)
		}
	}
}

func TestUint512Mul64(t *testing.T) {
	for i := 0; i < 100_000; i++ {
		x := randUint512()
		y := randUint64()

		z, ok := x.mulCheck64(y)

		ybig := new(big.Int).SetUint64(y)
		want := new(big.Int).Mul(x.big(), ybig)
		if (want.Cmp(bigMaxUint512) <= 0) != ok {
			t.Fatalf("%d: %d * %d: expected %t",
				i, x.big(), ybig, !ok)
		}
		want.And(want, big512mask)

		if ok && x.mul64(y) != z {
			t.Fatalf("%d: %d * %d: %d != %d",
				i, x.big(), ybig, x.mul64(y), z)
		}
		z = x.mul64(y)
		if got := z.big(); got.Cmp(want) != 0 {
			t.Fatalf("%d: %d * %d: expected %d, got %d",
				i, x.big(), ybig, want, got)
		}
	}
}

func TestUint512Exp(t *testing.T) {
	x := U512(1)
	ten := U512(10)
	for i := 1; ; i++ {
		want, ok := x.mulCheck64(10)
		if !ok {
			break
		}
		got := ten.Exp(U512(uint64(i)), U512(0))
		if got != want {
			t.Fatalf("#%d: expected %q, got %q", i, want, got)
		}
		x = want
	}
}

func TestUint512QuoRem(t *testing.T) {
	for i := 0; i < 100_000; i++ {
		x := randUint512()
		y := randUint512()
		if y.IsZero() {
			y = U512(1)
		}

		q, r := x.QuoRem(y)

		wantq := new(big.Int)
		wantr := new(big.Int)
		wantq.QuoRem(x.big(), y.big(), wantr)
		wantq.And(wantq, big512mask)

		if got := q.big(); got.Cmp(wantq) != 0 {
			t.Fatalf("%d / %d expected quotient of %d, got %d",
				x.big(), y.big(), wantq, got)
		}
		if got := r.big(); got.Cmp(wantr) != 0 {
			t.Fatalf("%d / %d expected remainder of %d, got %d",
				x.big(), y.big(), wantr, got)
		}
	}
}

func TestUint512QuoRemHalf(t *testing.T) {
	for i := 0; i < 100_000; i++ {
		x := randUint512()
		y := randUint256()
		if y.IsZero() {
			y = U256(1)
		}

		q, r := x.quoRem256(y)

		wantq := new(big.Int)
		wantr := new(big.Int)
		wantq.QuoRem(x.big(), y.big(), wantr)
		wantq.And(wantq, big512mask)

		if got := q.big(); got.Cmp(wantq) != 0 {
			t.Fatalf("%d / %d expected quotient of %d, got %d",
				x.big(), y.big(), wantq, got)
		}
		if got := r.big(); got.Cmp(wantr) != 0 {
			t.Fatalf("%d / %d expected remainder of %d, got %d",
				x.big(), y.big(), wantr, got)
		}
	}
}

func TestUint512QuoRem64(t *testing.T) {
	for i := 0; i < 100_000; i++ {
		x := randUint512()
		y := randUint64()
		if y == 0 {
			y = 1
		}

		q, r := x.quoRem64(y)

		ybig := new(big.Int).SetUint64(y)
		wantq := new(big.Int)
		wantr := new(big.Int)
		wantq.QuoRem(x.big(), ybig, wantr)
		wantq.And(wantq, big512mask)

		if got := q.big(); got.Cmp(wantq) != 0 {
			t.Fatalf("%d / %d expected quotient of %d, got %d",
				x.big(), ybig, wantq, got)
		}
		if got := new(big.Int).SetUint64(r); got.Cmp(wantr) != 0 {
			t.Fatalf("%d / %d expected remainder of %d, got %d",
				x.big(), ybig, wantr, got)
		}
	}
}

func TestUint512String(t *testing.T) {
	test := func(x Uint512) {
		want := x.big().String()
		got := x.String()
		if want != got {
			t.Fatalf("expected %q, got %q", want, got)
		}
	}
	test(Uint512{})       // min
	test(Uint512{}.max()) // max
	for i := 0; i < 10_000; i++ {
		test(randUint512())
	}
}

func TestParseUint512(t *testing.T) {
	for i := 0; i < 10_000; i++ {
		want := randUint512()
		b := want.big()
		for base := 2; base <= 36; base++ {
			s := b.Text(base)
			got, err := ParseUint512(s, base)
			if err != nil {
				t.Fatalf("%q in base %d: unexpected error: %v", s, base, err)
			}
			if got != want {
				t.Fatalf("%q in base %d: expected %#v, got %#v",
					s, base, want, got)
			}
		}
	}
}

func BenchmarkUint512Add(b *testing.B) {
	s := make([]Uint512, 1000)
	for i := range s {
		s[i] = randUint512()
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		x := s[i%len(s)]
		y := s[(i+1)%len(s)]
		sink.Uint512 = x.Add(y)
	}
}

func BenchmarkUint512Sub(b *testing.B) {
	s := make([]Uint512, 1000)
	for i := range s {
		s[i] = randUint512()
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		x := s[i%len(s)]
		y := s[(i+1)%len(s)]
		sink.Uint512 = x.Sub(y)
	}
}

func BenchmarkUint512Mul(b *testing.B) {
	s := make([]Uint512, 1000)
	for i := range s {
		s[i] = randUint512()
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		x := s[i%len(s)]
		y := s[(i+1)%len(s)]
		sink.Uint512 = x.Mul(y)
	}
}

func BenchmarkUint512QuoRem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink.Uint512, sink.Uint512 = U512(uint64(i + 2)).QuoRem(U512(uint64(i + 1)))
	}
}

func BenchmarkUint512QuoRem64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink.Uint512, sink.uint64 = U512(uint64(i + 2)).quoRem64(uint64(i + 1))
	}
}
