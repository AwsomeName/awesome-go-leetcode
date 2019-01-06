package problem0972

func isRationalEqual(S string, T string) bool {
	i := 0
	IS, IT := 0, 0
	for i < len(S) && S[i] != '.' {
		// if S[i]!=T[i] {
		//     return false
		// }
		IS = IS*10 + int(S[i]-'0')
		IT = IT*10 + int(T[i]-'0')
		i++
	}

	SEP := false
	if IS-IT == 1 || IS-IT == -1 {
		SEP = true
	} else {
		if IS != IT {
			return false
		}
	}
	fmt.Println("sep:", SEP)

	i++
	j := i
	FS := make([]int, 0) //xiao shu dian hou 10wei ,yong lai bi jiao
	FT := make([]int, 0)

	IFS, IFT := 0, 0
	for i < len(S) && S[i] != '(' {
		FS = append(FS, int(S[i]-'0'))
		IFS = IFS*10 + int(S[i]-'0')

		i++
	}
	for j < len(T) && T[j] != '(' {
		FT = append(FT, int(T[j]-'0'))
		IFT = IFT*10 + int(T[j]-'0')
		j++
	}

	SSEP := false
	if (IFS-IFT == 1 || IFS-IFT == -1) && (i < len(S) || j < len(T)) && !SEP && i == j {
		SSEP = true
	}
	fmt.Println("ssep:", SSEP)

	i++
	j++
	RS := make([]int, 0)
	RT := make([]int, 0)

	for i < len(S) && S[i] != ')' {
		RS = append(RS, int(S[i]-'0'))
		i++
	}
	for j < len(T) && T[j] != ')' {
		RT = append(RT, int(T[j]-'0'))
		j++
	}

	ii := 0
	for x := len(FS); x <= 8; x++ {
		tmp := 0
		if len(RS) > 0 {
			tmp = RS[ii]
			ii++
			if ii >= len(RS) {
				ii = 0
			}
		}
		FS = append(FS, tmp)
	}
	ii = 0
	for x := len(FT); x <= 8; x++ {
		tmp := 0
		if len(RT) > 0 {
			tmp = RT[ii]
			ii++
			if ii >= len(RT) {
				ii = 0
			}
		}
		FT = append(FT, tmp)
	}

	if !SEP {
		if !SSEP {
			for m, n := range FT {
				if n != FS[m] {
					return false
				}
			}
		} else {
			if IFS < IFT {
				for m, n := range RT {
					if n != 0 || RS[m] != 9 {
						return false
					}
				}
			} else {
				for m, n := range RT {
					if n != 9 || RS[m] != 0 {
						return false
					}
				}
			}
		}
	} else {
		fmt.Println("FS:", FS, IS)
		fmt.Println("FT:", FT, IT)
		if IS < IT {
			for m, n := range FT {
				if n != 0 || FS[m] != 9 {
					return false
				}
			}
		} else {
			for m, n := range FT {
				if n != 9 || FS[m] != 0 {
					return false
				}
			}
		}
	}

	return true
}
