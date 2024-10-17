package passwordbench

import (
	"testing"
)

func Benchmark_Baseline(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		_ = Plaintext(32)
	}
}

func Benchmark_Bcrypt_Cost10(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		plaintext := Plaintext(32)
		cost := 10
		_, _ = Bcrypt(plaintext, cost)
	}
}

func Benchmark_Bcrypt_Cost12(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		plaintext := Plaintext(32)
		cost := 12
		_, _ = Bcrypt(plaintext, cost)
	}
}

func Benchmark_PBPDF2_HMAC_SHA256_Iter600000(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		plaintext := Plaintext(32)
		// RFC8018 says the iterations must be at least 1000.
		// https://www.rfc-editor.org/rfc/rfc8018#section-4.2
		// OWASP says the iterations should be 600000.
		// https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#pbkdf2
		iterations := 600000
		_, _ = PBKDF2_HMAC_SHA256(plaintext, iterations)
	}
}

func Benchmark_Argon2id_t3_m12288_p1(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		plaintext := Plaintext(32)
		_, _ = Argon2id(plaintext, 3, 12288, 1)
	}
}

func Benchmark_Argon2id_t4_m9216_p1(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		plaintext := Plaintext(32)
		_, _ = Argon2id(plaintext, 4, 9216, 1)
	}
}

func Benchmark_Argon2id_t5_m7168_p1(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		plaintext := Plaintext(32)
		_, _ = Argon2id(plaintext, 5, 7168, 1)
	}
}

func Benchmark_Scrypt_N32768_r8_p1(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		plaintext := Plaintext(32)
		_, _ = Scrypt(plaintext, 32768, 8, 1)
	}
}

func Benchmark_Scrypt_N32768_r8_p3(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		plaintext := Plaintext(32)
		_, _ = Scrypt(plaintext, 32768, 8, 3)
	}
}

func Benchmark_Scrypt_N16384_r8_p1(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		plaintext := Plaintext(32)
		_, _ = Scrypt(plaintext, 16384, 8, 1)
	}
}

func Benchmark_Scrypt_N16384_r8_p5(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		plaintext := Plaintext(32)
		_, _ = Scrypt(plaintext, 16384, 8, 5)
	}
}
