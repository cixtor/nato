package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()

	// source: https://en.wikipedia.org/wiki/NATO_phonetic_alphabet
	alphabet := map[string]string{
		"A": "Alfa",     // (AL-FAH)
		"B": "Bravo",    // (BRAH-VOH)
		"C": "Charlie",  // (CHAR-LEE)
		"D": "Delta",    // (DELL-TAH)
		"E": "Echo",     // (ECK-OH)
		"F": "Foxtrot",  // (FOKS-TROT)
		"G": "Golf",     // (GOLF)
		"H": "Hotel",    // (HOH-TEL)
		"I": "India",    // (IN-DEE-AH)
		"J": "Juliett",  // (JEW-LEE-ETT)
		"K": "Kilo",     // (KEY-LOH)
		"L": "Lima",     // (LEE-MAH)
		"M": "Mike",     // (MIKE)
		"N": "November", // (NO-VEM-BER)
		"O": "Oscar",    // (OSS-CAH)
		"P": "Papa",     // (PAH-PAH)
		"Q": "Quebec",   // (KEH-BECK)
		"R": "Romeo",    // (ROW-ME-OH)
		"S": "Sierra",   // (SEE-AIR-RAH)
		"T": "Tango",    // (TANG-GO)
		"U": "Uniform",  // (YOU-NEE-FORM)
		"V": "Victor",   // (VIK-TAH)
		"W": "Whiskey",  // (WISS-KEY)
		"X": "Xray",     // (ECKS-RAY)
		"Y": "Yankee",   // (YANG-KEY)
		"Z": "Zulu",     // (ZOO-LOO)
		"1": "One",      // (WUN)
		"2": "Two",      // (TOO)
		"3": "Three",    // (TREE)
		"4": "Four",     // (FOW-ER)
		"5": "Five",     // (FIFE)
		"6": "Six",      // (SIX)
		"7": "Seven",    // (SEV-EN)
		"8": "Eight",    // (AIT)
		"9": "Nine",     // (NIN-ER)
		"0": "Zero",     // (ZEE-RO)
	}

	var phonics []string
	everything := strings.Join(flag.Args(), "")

	for i, c := range everything {
		letter := strings.ToUpper(string(c))
		if i%2 == 0 {
			phonics = append(phonics, "\x1b[0;94m"+alphabet[letter]+"\x1b[0m")
		} else {
			phonics = append(phonics, alphabet[letter])
		}
	}

	fmt.Println(strings.Join(phonics, "\x20"))
}
