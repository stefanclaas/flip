package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var flipTable = map[rune]rune{
	'a': 'ɐ',
	'b': 'q',
	'c': 'ɔ',
	'd': 'p',
	'e': 'ǝ',
	'f': 'ɟ',
	'g': 'ƃ',
	'h': 'ɥ',
	'i': 'ᴉ',
	'j': 'ɾ',
	'k': 'ʞ',
	'l': 'l',
	'm': 'ɯ',
	'n': 'u',
	'o': 'o',
	'p': 'd',
	'q': 'b',
	'r': 'ɹ',
	's': 's',
	't': 'ʇ',
	'u': 'n',
	'v': 'ʌ',
	'w': 'ʍ',
	'x': 'x',
	'y': 'ʎ',
	'z': 'z',
	'A': 'Ɐ',
	'B': 'ᗺ',
	'C': 'Ɔ',
	'D': 'ᗡ',
	'E': 'Ǝ',
	'F': 'Ⅎ',
	'G': '⅁',
	'H': 'H',
	'I': 'I',
	'J': 'ᒋ',
	'K': 'ꓘ',
	'L': '⅂',
	'M': 'ꟽ',
	'N': 'N',
	'O': 'O',
	'P': 'Ԁ',
	'Q': 'ტ',
	'R': 'ᴚ',
	'S': 'S',
	'T': 'Ʇ',
	'U': 'Ո',
	'V': 'Λ',
	'W': 'M',
	'X': 'X',
	'Y': '⅄',
	'Z': 'Z',
	'1': '⇂',
	'2': 'ᘔ',
	'3': 'Ɛ',
	'4': '߈',
	'5': 'ဌ',
	'6': '9',
	'7': 'ㄥ',
	'8': '8',
	'9': '6',
	'0': '0',
	'.': '˙',
	'[': ']',
	'\'': ',',
	',': '\'',
	'(': ')',
	'{': '}',
	'?': '¿',
	'!': '¡',
	'"': ',',
	'<': '>',
	'_': '‾',
	';': '؛',
	'‿': '⁀',
	'⁅': '⁆',
	'∴': '∵',
	'\r': '\n',
	' ': ' ',
}

var reverseFlag = flag.Bool("r", false, "reverse the input string")

func main() {
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		flipped := flip(line, *reverseFlag)
		fmt.Println(flipped)
	}
}

func flip(input string, reverse bool) string {
	runes := []rune(input)
	if !reverse {
		for idx, ch := range runes {
			if val, ok := flipTable[ch]; ok {
				runes[idx] = val
			}
		}
	} else {
		for idx, ch := range runes {
			for key, value := range flipTable {
				if ch == value {
					runes[idx] = key
					break
				}
			}
		}
	}
	return string(runes)
}