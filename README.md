```sh
$ go test -bench . -benchmem -benchtime 5s
goos: darwin
goarch: amd64
pkg: passwordhash
cpu: Intel(R) Core(TM) i5-7267U CPU @ 3.10GHz
Benchmark_Baseline-4                        	 3796885	      1505 ns/op	     160 B/op	       3 allocs/op
Benchmark_Bcrypt_Cost10-4                   	      92	  63040934 ns/op	    5542 B/op	      15 allocs/op
Benchmark_Bcrypt_Cost12-4                   	      22	 275535410 ns/op	    5584 B/op	      15 allocs/op
Benchmark_PBPDF2_HMAC_SHA256_Iter600000-4   	      16	 323931011 ns/op	    1013 B/op	      15 allocs/op
Benchmark_Argon2id_t3_m12288_p1-4           	     210	  27708793 ns/op	12588082 B/op	      45 allocs/op
Benchmark_Argon2id_t4_m9216_p1-4            	     211	  28208118 ns/op	 9442549 B/op	      53 allocs/op
Benchmark_Argon2id_t5_m7168_p1-4            	     225	  29167244 ns/op	 7345632 B/op	      61 allocs/op
Benchmark_Scrypt_N32768_r8_p1-4             	      52	  97451026 ns/op	33559276 B/op	      27 allocs/op
Benchmark_Scrypt_N32768_r8_p3-4             	      20	 284210739 ns/op	33561313 B/op	      27 allocs/op
Benchmark_Scrypt_N16384_r8_p1-4             	     132	  45973887 ns/op	16782060 B/op	      27 allocs/op
Benchmark_Scrypt_N16384_r8_p5-4             	      26	 224935420 ns/op	16786412 B/op	      27 allocs/op
PASS
ok  	passwordhash	83.577s
```
