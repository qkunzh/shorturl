package util

const N int = 10
type IntBase62 struct {
	buf [N]byte
}
var (
	charToInt = map[byte]int {
		'0':0,
		'1':1,
		'2':2,
		'3':3,
		'4':4,
		'5':5,
		'6':6,
		'7':7,
		'8':8,
		'9':9,
		'a':10,
		'b':11,
		'c':12,
		'd':13,
		'e':14,
		'f':15,
		'g':16,
		'h':17,
		'i':18,
		'j':19,
		'k':20,
		'l':21,
		'm':22,
		'n':23,
		'o':24,
		'p':25,
		'q':26,
		'r':27,
		's':28,
		't':29,
		'u':30,
		'v':31,
		'w':32,
		'x':33,
		'y':34,
		'z':35,
		'A':36,
		'B':37,
		'C':38,
		'D':39,
		'E':40,
		'F':41,
		'G':42,
		'H':43,
		'I':44,
		'J':45,
		'K':46,
		'L':47,
		'M':48,
		'N':49,
		'O':50,
		'P':51,
		'Q':52,
		'R':53,
		'S':54,
		'T':55,
		'U':56,
		'V':57,
		'W':58,
		'X':59,
		'Y':60,
		'Z':61,
	}
	intToChar = map[int]byte {
		0:'0',
		1:'1',
		2:'2',
		3:'3',
		4:'4',
		5:'5',
		6:'6',
		7:'7',
		8:'8',
		9:'9',
		10:'a',
		11:'b',
		12:'c',
		13:'d',
		14:'e',
		15:'f',
		16:'g',
		17:'h',
		18:'i',
		19:'j',
		20:'k',
		21:'l',
		22:'m',
		23:'n',
		24:'o',
		25:'p',
		26:'q',
		27:'r',
		28:'s',
		29:'t',
		30:'u',
		31:'v',
		32:'w',
		33:'x',
		34:'y',
		35:'z',
		36:'A',
		37:'B',
		38:'C',
		39:'D',
		40:'E',
		41:'F',
		42:'G',
		43:'H',
		44:'I',
		45:'J',
		46:'K',
		47:'L',
		48:'M',
		49:'N',
		50:'O',
		51:'P',
		52:'Q',
		53:'R',
		54:'S',
		55:'T',
		56:'U',
		57:'V',
		58:'W',
		59:'X',
		60:'Y',
		61:'Z',
	}
)

func NewIntBase62(str string) *IntBase62 {
	ret := &IntBase62{}
	for i := 0; i < N; i++ {
		ret.buf[i] = '0'
	}
	bytes := []byte(str)
	n := len(bytes)
	for i := 0; i < n; i++ {
		ret.buf[N-1-i] = bytes[n-1-i]
	}
	return ret
}

func (this *IntBase62) Acc() {
	this.addOne(true)
}

func (this *IntBase62) Dec() {
	this.addOne(false)
}

func (this *IntBase62) addOne(positive bool) {
	carry := -1
	if positive {
		carry = 1
	}
	for i := N-1; i >= 0; i-- {
		num := charToInt[this.buf[i]]
		tmp := (carry + num) % 62
		carry = (carry + num) / 62
		this.buf[i] = intToChar[tmp]
	}
}

func (this *IntBase62) GreaterThan(other *IntBase62) bool {
	for i := 0; i < N; i++ {
		num1 := charToInt[this.buf[i]]
		num2 := charToInt[other.buf[i]]
		if num1 > num2 {
			return true
		}
	}
	return false
}
func (this *IntBase62) String() string {
	i := 0
	for ; i < N; i++ {
		if this.buf[i] != '0' {
			break
		}
	}
	return string(this.buf[i:])
}