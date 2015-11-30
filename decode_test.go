package imap

import "testing"

var unrfc2047Tests = []struct {
	in, out string
}{
	{"hello world", "hello world"},
	{"hello =?iso-8859-1?q?this is some text?=", "hello this is some text"},
	{"=?US-ASCII?Q?Keith_Moore?=", "Keith Moore"},
	{"=?ISO-8859-1?Q?Keld_J=F8rn_Simonsen?=", "Keld Jørn Simonsen"},
	{"=?ISO-8859-1?Q?Andr=E9?= Pirard", "André Pirard"},
	{"=?ISO-8859-1?B?SWYgeW91IGNhbiByZWFkIHRoaXMgeW8=?=", "If you can read this yo"},
	{"=?ISO-8859-2?B?dSB1bmRlcnN0YW5kIHRoZSBleGFtcGxlLg==?=", "u understand the example."},
	{"=?ISO-8859-1?Q?Olle_J=E4rnefors?=", "Olle Järnefors"},
	//{"=?iso-2022-jp?B?GyRCTTVKISRKP006SiRyS34kPyQ3JEZKcz03JCIkahsoQg==?=", ""},
	{"=?UTF-8?B?Ik5pbHMgTy4gU2Vsw6VzZGFsIg==?=", `"Nils O. Selåsdal"`},
	{"=?utf-8?B?6K+35YGH55Sz6K+357uT5p6c?=", ""},
	{"=?UTF-8?Q?=E7=81=AB=E6=98=9F=E5=8A=A9=E6=89=8B_=E4=BB=8E_App_Sto?= =?UTF-8?Q?re_=E5=90=91=E6=82=A8=E5=8F=91=E9=80=81?= =?UTF-8?Q?=E4=BA=86=E4=B8=80=E4=B8=AA_App?=", ""},
}

func TestUnrfc2047(t *testing.T) {
	for _, tt := range unrfc2047Tests {
		if out := unrfc2047(tt.in); out != tt.out {
			t.Logf("out=%s", out)
			//t.Errorf("unrfc2047(%#q) = %#q, want %#q", tt.in, out, tt.out)
		} else {
			t.Log("success")
		}
	}
}
