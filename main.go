package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	for _, word := range flag.Args() {
		fmt.Print(word + "\x20=\x20")

		for i, c := range word {
			if i%2 == 0 {
				fmt.Print("\x1b[0;94m" + alphabet[c] + "\x1b[0m ")
			} else {
				fmt.Print(alphabet[c] + "\x20")
			}
		}

		fmt.Println()
	}
}

// source: https://en.wikipedia.org/wiki/NATO_phonetic_alphabet
var alphabet = map[rune]string{
	'0': "Zero",     // (ZEE-RO)
	'1': "One",      // (WUN)
	'2': "Two",      // (TOO)
	'3': "Three",    // (TREE)
	'4': "Four",     // (FOW-ER)
	'5': "Five",     // (FIFE)
	'6': "Six",      // (SIX)
	'7': "Seven",    // (SEV-EN)
	'8': "Eight",    // (AIT)
	'9': "Nine",     // (NIN-ER)
	'A': "Alfa",     // (AL-FAH)
	'B': "Bravo",    // (BRAH-VOH)
	'C': "Charlie",  // (CHAR-LEE)
	'D': "Delta",    // (DELL-TAH)
	'E': "Echo",     // (ECK-OH)
	'F': "Foxtrot",  // (FOKS-TROT)
	'G': "Golf",     // (GOLF)
	'H': "Hotel",    // (HOH-TEL)
	'I': "India",    // (IN-DEE-AH)
	'J': "Juliett",  // (JEW-LEE-ETT)
	'K': "Kilo",     // (KEY-LOH)
	'L': "Lima",     // (LEE-MAH)
	'M': "Mike",     // (MIKE)
	'N': "November", // (NO-VEM-BER)
	'O': "Oscar",    // (OSS-CAH)
	'P': "Papa",     // (PAH-PAH)
	'Q': "Quebec",   // (KEH-BECK)
	'R': "Romeo",    // (ROW-ME-OH)
	'S': "Sierra",   // (SEE-AIR-RAH)
	'T': "Tango",    // (TANG-GO)
	'U': "Uniform",  // (YOU-NEE-FORM)
	'V': "Victor",   // (VIK-TAH)
	'W': "Whiskey",  // (WISS-KEY)
	'X': "Xray",     // (ECKS-RAY)
	'Y': "Yankee",   // (YANG-KEY)
	'Z': "Zulu",     // (ZOO-LOO)
	'a': "Alfa",     // (AL-FAH)
	'b': "Bravo",    // (BRAH-VOH)
	'c': "Charlie",  // (CHAR-LEE)
	'd': "Delta",    // (DELL-TAH)
	'e': "Echo",     // (ECK-OH)
	'f': "Foxtrot",  // (FOKS-TROT)
	'g': "Golf",     // (GOLF)
	'h': "Hotel",    // (HOH-TEL)
	'i': "India",    // (IN-DEE-AH)
	'j': "Juliett",  // (JEW-LEE-ETT)
	'k': "Kilo",     // (KEY-LOH)
	'l': "Lima",     // (LEE-MAH)
	'm': "Mike",     // (MIKE)
	'n': "November", // (NO-VEM-BER)
	'o': "Oscar",    // (OSS-CAH)
	'p': "Papa",     // (PAH-PAH)
	'q': "Quebec",   // (KEH-BECK)
	'r': "Romeo",    // (ROW-ME-OH)
	's': "Sierra",   // (SEE-AIR-RAH)
	't': "Tango",    // (TANG-GO)
	'u': "Uniform",  // (YOU-NEE-FORM)
	'v': "Victor",   // (VIK-TAH)
	'w': "Whiskey",  // (WISS-KEY)
	'x': "Xray",     // (ECKS-RAY)
	'y': "Yankee",   // (YANG-KEY)
	'z': "Zulu",     // (ZOO-LOO)
}
