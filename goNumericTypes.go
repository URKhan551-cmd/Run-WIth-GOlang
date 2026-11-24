// *how to decied as an engineer about the data representation, correctness, memory, range, and performance.
// Go gives us a full control what kind of number we are dealing with.
// the choice to choose affect the compilation, computation, axuarte, data overflows corretcly behave under load.

// NUMERIC TYPES IN GO
// Integer --> Whole Numbers
// int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr

// FLOATIG POINT.
// fractional / approximate 
// float32, float64

// COMPLEX NUMBERS:
// real / imaginary components
// complex64,  complex128

// differeny repreenattion have different capabilities & limitation

// BITS : small unit of binary information 0 or 1
// if a  type has 8bits  2^8 = 256 distinct bit pattern
// if it has 16bits = 2^16 = 65,536 patterns

var count int = 78

func main(){  // infer as int for all of these
  msgCount := 100
  retryCount := 3
  userCount := 7887
}

messages := []string{"umair", "principal engineer", "khan"}
for i := 0; i < len(messages); i++ {
  fmt.Println(messages[i])
}
// int is a natural choice for LOOP and INDEXES.

// IN JavaSccript 
// const messageCount = 700;
// Number type represent both whole and fractional 

// in GO 
messageCount := 899
// compiler infer as int not float64

// NumericTypes have Sizes: 

