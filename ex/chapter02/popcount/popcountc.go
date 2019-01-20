package popcount

// static inline int popcount(unsigned long long x)
// {
// 	return __builtin_popcountll(x);
// }
//
// static inline int __attribute__((optimize("O0"))) popcount10000000(unsigned long long x)
// {
//      int i;
//
//      for (i = 0; i < 10000000; i++)
// 	    __builtin_popcountll(x);
//
// 	return __builtin_popcountll(x);
// }
import "C"

func PopCountC(x uint64) int {
	return int(C.popcount(C.ulonglong(x)))
}

func PopCountC10000000(x uint64) int {
	return int(C.popcount10000000(C.ulonglong(x)))
}
