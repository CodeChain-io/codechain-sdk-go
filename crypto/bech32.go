// bech32 implementations are based on codechain-primitives-js
package crypto

import (
	"strings"
)

const alphabet = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

func alphabet_map(c rune) int {
	for k, v := range alphabet {
		if c == v {
			return k
		}
	}
	panic(c)
}

func polymodeStep(pre int) int {
	b := pre >> 25
	return (((pre & 0x1ffffff) << 5) ^
		(-((b >> 0) & 1) & 0x3b6a57b2) ^
		(-((b >> 1) & 1) & 0x26508e6d) ^
		(-((b >> 2) & 1) & 0x1ea119fa) ^
		(-((b >> 3) & 1) & 0x3d4233dd) ^
		(-((b >> 4) & 1) & 0x2a1462b3))
}

func prefixChk(prefix string) (chk int) {
	chk = 1
	for _, c := range prefix {
		if c < 33 || c > 126 {
			panic(c)
		}
		chk = polymodeStep(chk) ^ (int(c) >> 5)
	}
	chk = polymodeStep(chk)

	for _, c := range prefix {
		chk = polymodeStep(chk) ^ (int(c) & 0x1f)
	}
	return
}

func bech32Encode(prefix string, words []byte) (result string) {

	prefix = strings.ToLower(prefix)
	var chk = prefixChk(prefix)
	result = prefix

	for _, c := range words {
		chk = polymodeStep(chk) ^ int(c)

		result += string(alphabet[int(c)])
	}

	for i := 0; i < 6; i++ {
		chk = polymodeStep(chk)
	}
	chk ^= 1
	for i := 0; i < 6; i++ {
		v := (chk >> ((5 - uint(i)) * 5)) & 0x1f
		result += string(alphabet[int(v)])
	}
	return
}

func bech32Decode(str string, prefix string) (words []byte) {
	// TODO Handle errors
	var chk = prefixChk(prefix)
	wordChars := str[len(prefix):]
	for i, c := range wordChars {
		v := alphabet_map(c)
		chk = polymodeStep(chk) ^ int(v)

		if i+6 >= len(wordChars) {
			continue
		}
		words = append(words, byte(v))
	}
	return
}

func convert(data []byte, inBits int, outBits int, pad bool) (result []byte) {
	var value = 0
	var bits = 0
	var maxV = (1 << uint(outBits)) - 1

	for _, c := range data {
		value = (value << uint(inBits)) | int(c)
		bits += inBits

		for bits >= outBits {
			bits -= outBits
			result = append(result, byte((value>>uint(bits))&maxV))
		}
	}

	if pad {
		if bits > 0 {
			result = append(result, byte((value<<uint(outBits-bits))&maxV))
		}
	} else {
		if bits >= inBits {
			panic(bits)
		}
		if (value<<uint(outBits-bits))&maxV != 0 {
			panic(value)
		}
	}

	return
}

func fromWords(words []byte) (conv []byte) {
	conv = convert(words, 5, 8, false)
	return
}

func toWords(bytes []byte) (conv []byte) {
	conv = convert(bytes, 8, 5, true)
	return
}
