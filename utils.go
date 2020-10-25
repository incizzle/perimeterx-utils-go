package utils

import (
	"strconv"
)

// Create PC Variable pass in jsonstring payload and uuid:tag:ftag
func createPC(jsonstring string, tag string) string {
	var a = B(jsonstring, tag)
	var e = j(a)
	var r = Yt(e)
	var pc = ""

	for i := 0; i < len(r); i += 2 {
		pc += string(r[i])
	}
	return pc
}

func B(n string, t string) string {
	var e int
	r := V(t)
	o := make([]int, 16)
	i := make([]int, 16)
	if len(r) > 16 {
		W(r, 8 * len(t))
	}
	for e = 0; e < 16; e ++ {
		if (e >= len(r)) {
			o[e] = 909522486
			i[e] = 1549556828
		} else {
			o[e] = 909522486 ^ r[e]
			i[e] = 1549556828 ^ r[e]
		}
	}
	a := W(append(o, V(n)...), 512 + 8 * len(n))
	x := W(append(i, a...), 640)
	return F(x)
}

func V(t string) []int {
	e := make([]int, (len(t) * 8) >> 5 + 1)
    for n := 0; n < 8 * len(t); n += 8 {
		charCode := int([]rune(t)[n / 8])
		e[n >> 5] |= charCode << (n % 32);
	}
	return e
}

func W(t []int, n int) []int {
	var x int32 = 128 << (n % 32)
	for i := len(t); i < ((n >> 5) + 1); i++ {
		t = append(t, 0)
	}
	t[n >> 5] |= int(x)
	var y = 14 + ((n + 64) >> 9 << 4)
	for i := len(t); i < y + 2; i++ {
		t = append(t, 0)
	}
	t[y] = n
	var e int
    var r = 0
    var o = 0
    var i = 0
    var a = 0
    var c = 1732584193
    var u = -271733879
    var f = -1732584194
    var s = 271733878;
	for e = 0; e < len(t); e += 16 {
		r = c
		o = u
		i = f
		a = s
		c = C(c, u, f, s, t[e], 7, -680876936)
		s = C(s, c, u, f, t[e + 1], 12, -389564586)
		f = C(f, s, c, u, t[e + 2], 17, 606105819)
		u = C(u, f, s, c, t[e + 3], 22, -1044525330)
		c = C(c, u, f, s, t[e + 4], 7, -176418897)
		s = C(s, c, u, f, t[e + 5], 12, 1200080426)
		f = C(f, s, c, u, t[e + 6], 17, -1473231341)
		u = C(u, f, s, c, t[e + 7], 22, -45705983)
		c = C(c, u, f, s, t[e + 8], 7, 1770035416)
		s = C(s, c, u, f, t[e + 9], 12, -1958414417)
		f = C(f, s, c, u, t[e + 10], 17, -42063)
		u = C(u, f, s, c, t[e + 11], 22, -1990404162)
		c = C(c, u, f, s, t[e + 12], 7, 1804603682)
		s = C(s, c, u, f, t[e + 13], 12, -40341101)
		f = C(f, s, c, u, t[e + 14], 17, -1502002290)
		u = C(u, f, s, c, t[e + 15], 22, 1236535329)
		c = N(c, u, f, s, t[e + 1], 5, -165796510)
		s = N(s, c, u, f, t[e + 6], 9, -1069501632)
        f = N(f, s, c, u, t[e + 11], 14, 643717713)
        u = N(u, f, s, c, t[e], 20, -373897302)
        c = N(c, u, f, s, t[e + 5], 5, -701558691)
        s = N(s, c, u, f, t[e + 10], 9, 38016083)
        f = N(f, s, c, u, t[e + 15], 14, -660478335)
        u = N(u, f, s, c, t[e + 4], 20, -405537848)
        c = N(c, u, f, s, t[e + 9], 5, 568446438)
        s = N(s, c, u, f, t[e + 14], 9, -1019803690)
        f = N(f, s, c, u, t[e + 3], 14, -187363961)
        u = N(u, f, s, c, t[e + 8], 20, 1163531501)
        c = N(c, u, f, s, t[e + 13], 5, -1444681467)
        s = N(s, c, u, f, t[e + 2], 9, -51403784)
        f = N(f, s, c, u, t[e + 7], 14, 1735328473)
		u = N(u, f, s, c, t[e + 12], 20, -1926607734)
		c = D(c, u, f, s, t[e + 5], 4, -378558)
		s = D(s, c, u, f, t[e + 8], 11, -2022574463)
        f = D(f, s, c, u, t[e + 11], 16, 1839030562)
        u = D(u, f, s, c, t[e + 14], 23, -35309556)
        c = D(c, u, f, s, t[e + 1], 4, -1530992060)
        s = D(s, c, u, f, t[e + 4], 11, 1272893353)
        f = D(f, s, c, u, t[e + 7], 16, -155497632)
        u = D(u, f, s, c, t[e + 10], 23, -1094730640)
        c = D(c, u, f, s, t[e + 13], 4, 681279174)
        s = D(s, c, u, f, t[e], 11, -358537222)
        f = D(f, s, c, u, t[e + 3], 16, -722521979)
        u = D(u, f, s, c, t[e + 6], 23, 76029189)
        c = D(c, u, f, s, t[e + 9], 4, -640364487)
        s = D(s, c, u, f, t[e + 12], 11, -421815835)
        f = D(f, s, c, u, t[e + 15], 16, 530742520)
		u = D(u, f, s, c, t[e + 2], 23, -995338651)
		c = M(c, u, f, s, t[e], 6, -198630844)
		s = M(s, c, u, f, t[e + 7], 10, 1126891415)
        f = M(f, s, c, u, t[e + 14], 15, -1416354905)
        u = M(u, f, s, c, t[e + 5], 21, -57434055)
        c = M(c, u, f, s, t[e + 12], 6, 1700485571)
        s = M(s, c, u, f, t[e + 3], 10, -1894986606)
        f = M(f, s, c, u, t[e + 10], 15, -1051523)
        u = M(u, f, s, c, t[e + 1], 21, -2054922799)
        c = M(c, u, f, s, t[e + 8], 6, 1873313359)
        s = M(s, c, u, f, t[e + 15], 10, -30611744)
        f = M(f, s, c, u, t[e + 6], 15, -1560198380)
        u = M(u, f, s, c, t[e + 13], 21, 1309151649)
        c = M(c, u, f, s, t[e + 4], 6, -145523070)
        s = M(s, c, u, f, t[e + 11], 10, -1120210379)
        f = M(f, s, c, u, t[e + 2], 15, 718787259)
		u = M(u, f, s, c, t[e + 9], 21, -343485551)
		c = R(c, r)
        u = R(u, o)
        f = R(f, i)
        s = R(s, a)
	}
	return []int{c, u, f, s}
}

func C(t int, n int, e int, r int, o int, i int, a int) int {
	return k(n & e | ^n & r, t, n, o, i, a)
}

func N(t int, n int, e int, r int, o int, i int, a int) int {
    return k(n & r | e & ^r, t, n, o, i, a)
}

func D(t int, n int, e int, r int, o int, i int, a int) int {
    return k(n ^ e ^ r, t, n, o, i, a)
}

func M(t int, n int, e int, r int, o int, i int, a int) int {
    return k(e ^ (n | ^r), t, n, o, i, a)
}

func k(t int, n int, e int, r int, o int, i int) int {
	return R(O(R(R(n, t), R(r, i)), o), e)
}

func R(t int, n int) int {
	var e = (65535 & t) + (65535 & n);
	return int(int32((int(int32(t >> 16)) + int(int32(n >> 16)) + int(int32(e >> 16))) << 16)) | 65535 & e
}

func O(t int, n int) int {
	return int(int32(t << n)) | (int(uint32(t) >> (32 - n)))
}

func F(t []int) string {
	var n int
	var e = ""
	for n = 0; n < 32 * len(t); n+=8 {
		e += string((int(int32(int(uint32(t[n >> 5])) >> (n % 32)))) & 255)
	}
	return e
}

func j(t string) string {
    var n = "0123456789abcdef"
    var e = ""
    var r = 0
	var o int
    for o = 0; o < 16; o++ {
        r = int([]rune(t)[o])
        e += string(n[int(uint32(r) >> 4) & 15 ]) + string(n[15 & r])
	}
	return e
}

func Yt(t string) string {
	var n string = ""
	var e string = ""
	var r int
    for r = 0; r < len(t); r++ {
		var o = int([]rune(t)[r])
		if o >= 48 && o <= 57 {
			n += string(t[r])
		} else {
			e += strconv.Itoa(o % 10)
		}
	}
    return n + e
}

// Simple Function to Obfuscate string using a factor
func obfuscateString(text string, amount int) string {
	var e string = ""
	for r := 0; r < len(text); r++ {
		e += string(amount ^ int([]rune(text)[r]))
	}
	return e
}