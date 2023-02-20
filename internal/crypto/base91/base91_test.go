/*
 * A Go implementation of Joachim Henke's code from http://base91.sourceforge.net.
 *
 * Original by Joachim Henke, this implementation by Michael Traver.
 * License from Joachim Henke's source:
 *
 * Copyright (c) 2000-2006 Joachim Henke
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 *   - Redistributions of source code must retain the above copyright notice, this
 *     list of conditions and the following disclaimer.
 *   - Redistributions in binary form must reproduce the above copyright notice,
 *     this list of conditions and the following disclaimer in the documentation
 *     and/or other materials provided with the distribution.
 *   - Neither the name of Joachim Henke nor the names of his contributors may be
 *     used to endorse or promote products derived from this software without
 *     specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
 * ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
 * ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package base91

import (
	"bytes"
	"fmt"
	"testing"
)

type pair struct {
	decoded, encoded string
}

var pairs = []pair{
	{
		"May your trails be crooked, winding, lonesome, dangerous, leading to the most amazing view. May your mountains rise into and above the clouds.",
		"8D9KR`0eLUd/ZQFl62>vb,1RL%%&~8bju\"sQ;mmaU=UfU)1T70<^rm?i;Ct)/p;R(&^m5PKimf2+H[QSd/[E<oTPgZh>DZ%y;#,aIl]U>vP:3pPIqSwPmLwre3:W.{6U)/wP;mYBxgP[UCsS)/[EOiqMgZR*Sk<Rd/=8jL=ibg7+b[C",
	},

	// Random bytes
	{
		"\x35\x5e\x56\xe0\xc6\x29\x38\xf4\x81\x00\xab\x81\x7e\xd7\x08\x95\x62\x20\xa7\xda\x64\xa2\xce\xb3\xc5",
		"~_1H=x_t{|$AjJX(nMFdjL~:?1b3HgM",
	},

	// RFC 3548 examples (adapted from base64 to base91)
	{"\x14\xfb\x9c\x03\xd9\x7e", "Q<c[2!,C"},
	{"\x14\xfb\x9c\x03\xd9", "Q<c[2!B"},
	{"\x14\xfb\x9c\x03", "Q<c[A"},

	// RFC 4648 examples (adapted from base64 to base91)
	{"", ""},
	{"f", "LB"},
	{"fo", "drD"},
	{"foo", "dr.J"},
	{"foob", "dr/2Y"},
	{"fooba", "dr/2s)A"},
	{"foobar", "dr/2s)uC"},
}

func TestEncode(t *testing.T) {
	for i, p := range pairs {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			dst := make([]byte, StdEncoding.EncodedLen(len(p.decoded)))

			n := StdEncoding.Encode(dst, []byte(p.decoded))
			got := dst[:n]
			if !bytes.Equal(got, []byte(p.encoded)) {
				t.Errorf("Expected %v, got %v", []byte(p.encoded), got)
			}
		})
	}
}

func TestEncodeToString(t *testing.T) {
	for i, p := range pairs {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			got := StdEncoding.EncodeToString([]byte(p.decoded))
			if got != p.encoded {
				t.Errorf("Expected %v, got %v", p.encoded, got)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	for i, p := range pairs {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			dst := make([]byte, StdEncoding.DecodedLen(len(p.encoded)))

			n, err := StdEncoding.Decode(dst, []byte(p.encoded))
			if err != nil {
				t.Errorf("Got decoding error: %v", err)
			} else {
				got := dst[:n]
				if !bytes.Equal(got, []byte(p.decoded)) {
					t.Errorf("Expected %v, got %v", []byte(p.decoded), got)
				}
			}
		})
	}
}

func TestDecodeInvalidData(t *testing.T) {
	cases := []string{
		"~_1H=x_t{ |$AjJX(nMFdjL~:?1b3HgM", // Spaces are not in the standard encoding alphabet.
		"-", "\\", "'",                     // These characters are not in the standard encoding alphabet.
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			dst := make([]byte, StdEncoding.DecodedLen(len(tc)))

			_, err := StdEncoding.Decode(dst, []byte(tc))
			if err == nil {
				t.Errorf("Expected decoding error, got nil")
			}
		})
	}
}

func TestDecodeString(t *testing.T) {
	for i, p := range pairs {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			got, err := StdEncoding.DecodeString(p.encoded)
			if err != nil {
				t.Errorf("Got decoding error: %v", err)
			} else {
				if !bytes.Equal(got, []byte(p.decoded)) {
					t.Errorf("Expected %v, got %v", []byte(p.decoded), got)
				}
			}
		})
	}
}
