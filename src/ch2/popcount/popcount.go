package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		//fmt.Printf("pc[%v] %v, pc[%v] %v, byte %v\n", i, pc[i], i/2, pc[i/2], byte(i&1))
	}

	//fmt.Println("pc", pc)
}

func PopCount(x uint64) int {

	//fmt.Println("x>> ", x>>0, x>>8, x>>16, x>>24, x>>32, x>>40, x>>48, x>>56)
	//fmt.Println("byte(x>>) ", byte(x>>0), byte(x>>8), byte(x>>16), byte(x>>24), byte(x>>32), byte(x>>40), byte(x>>48), byte(x>>56))
	//fmt.Println("bit byte(x>>) ", biu.ByteToBinaryString(byte(x>>0)), biu.ByteToBinaryString(byte(x>>8)), biu.ByteToBinaryString(byte(x>>16)), biu.ByteToBinaryString(byte(x>>24)), biu.ByteToBinaryString(byte(x>>32)),
	//	biu.ByteToBinaryString(byte(x>>40)), biu.ByteToBinaryString(byte(x>>48)), biu.ByteToBinaryString(byte(x>>56)))
	//fmt.Println("pc[byte(x>>)]", pc[byte(x>>0)], pc[byte(x>>8)], pc[byte(x>>16)], pc[byte(x>>24)], pc[byte(x>>32)], pc[byte(x>>40)], pc[byte(x>>48)], pc[byte(x>>56)])

	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount23(x uint64) int {
	r := 0
	for i := 0; i < 8; i++ {
		r += int(pc[byte(x>>uint(i*8))])
	}
	return r
}

func PopCount24(x uint64) int {
	r := 0
	var i uint
	for i = 0; i < 64; i++ {
		if (x>>i)&1 == 1 {
			r++
		}
	}
	return r
}

func PopCount25(x uint64) int {
	r := 0
	for x != 0 {
		x = x & (x - 1)
		r++
	}

	return r
}
