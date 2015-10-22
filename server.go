package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"unsafe"
)

var verbose = flag.Bool("v", true, "")

func noun(n int) string {
	if n == 1 {
		return "bottle"
	}
	return "bottles"
}

func nobottles() string {
	kitchenSink()
	nomore := fmt.Sprintf("No more %s of beer on the wall!", noun(0))
	return nomore
}
func bottles(n int, song []string) (r []string) {
	var line string
	if n <= 0 {
		line = nobottles()
		song = append(song, line)
		r = song
	} else {
		line = fmt.Sprintf("%d %s of beer on the wall!", n, noun(n))
		song = append(song, line)
		r = bottles(n-1, song)
	}
	fmt.Printf("%s","")
	return r
}

func Beer(w http.ResponseWriter, r *http.Request) {
	if *verbose {
		log.Println("beer request")
	}
	song := bottles(10, nil)
	fmt.Fprintln(w, "<html><body>\n")
	fmt.Fprintln(w, strings.Join(song, "<br>"))
	fmt.Fprintln(w, "</body></html>\n")
}

func main() {
	flag.Parse()
	http.HandleFunc("/bar", Beer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type FooInterface interface {
	Bar()
}

type FooStruct struct {
	a int
	b string
}

func (f *FooStruct) Bar() {}

type myInt int

func kitchenSink() {
	var (
		Z_bool_false          bool        = false
		Z_bool_true           bool        = true
		Z_int                 int         = -21
		Z_int8                int8        = -121
		Z_int16               int16       = -32321
		Z_int32               int32       = -1987654321
		Z_int64               int64       = -9012345678987654321
		Z_uint                uint        = 21
		Z_uint8               uint8       = 231
		Z_uint16              uint16      = 54321
		Z_uint32              uint32      = 3217654321
		Z_uint64              uint64      = 12345678900987654321
		Z_uintptr             uintptr     = 21
		Z_float32             float32     = 1.54321
		Z_float64             float64     = 1.987654321
		Z_complex64           complex64   = 1.54321 + 2.54321i
		Z_complex128          complex128  = 1.987654321 - 2.987654321i
		Z_array               [5]int8     = [5]int8{-121, 121, 3, 2, 1}
		Z_array_temp          [5]int8     = [5]int8{-121, 121, 3, 2, 1}
		Z_array_empty         [0]int8     = [0]int8{}
		Z_array_of_empties    [2]struct{} = [2]struct{}{struct{}{}, struct{}{}}
		Z_channel             chan int    = make(chan int)
		Z_channel_buffered    chan int    = make(chan int, 10)
		Z_channel_nil         chan int
		Z_func_bar                              = (*FooStruct).Bar
		Z_func_int8_r_int8                      = func(x int8) int8 { return x + 1 }
		Z_func_int8_r_pint8                     = func(x int8) *int8 { y := x + 1; return &y }
		Z_func_copy                             = Z_func_int8_r_int8
		Z_func_nil            func(x int8) int8 = nil
		Z_struct              FooStruct         = FooStruct{a: 21, b: "hi"}
		Z_struct_temp         FooStruct         = FooStruct{a: 21, b: "hi"}
		Z_pointer             *FooStruct        = &Z_struct_temp
		Z_pointer_nil         *FooStruct
		Z_interface           FooInterface = &Z_struct_temp
		Z_interface_typed_nil FooInterface = Z_pointer_nil
		Z_interface_nil       FooInterface
		Z_map                 map[int8]float32 = map[int8]float32{-21: 3.54321}
		Z_map_2               map[int16]int8   = map[int16]int8{1024: 1}
		Z_map_empty           map[int8]float32 = map[int8]float32{}
		Z_map_nil             map[int8]float32
		Z_slice               []byte = []byte{'s', 'l', 'i', 'c', 'e'}
		Z_slice_2             []int8 = Z_array_temp[0:2]
		Z_slice_nil           []byte
		Z_string              string         = "I'm a string"
		Z_uint_temp           uint           = 21
		Z_unsafe_pointer      unsafe.Pointer = unsafe.Pointer(&Z_uint_temp)
		Z_unsafe_pointer_nil  unsafe.Pointer
		Z_byte                byte  = 111
		Z_rune                rune  = '~'
		Z_int_typedef         myInt = 42
	)
	for i := 0; i < 10; i++ {
		maybePrintln(Z_bool_false, Z_bool_true)
		maybePrintln(Z_int, Z_int8, Z_int16, Z_int32, Z_int64)
		maybePrintln(Z_uint, Z_uint8, Z_uint16, Z_uint32, Z_uint64, Z_uintptr)
		maybePrintln(Z_float32, Z_float64, Z_complex64, Z_complex128)
		maybePrintln(Z_array, Z_array_empty, Z_array_of_empties)
		maybePrintln(Z_channel, Z_channel_buffered, Z_channel_nil)
		maybePrintln(Z_func_bar, Z_func_int8_r_int8, Z_func_int8_r_pint8)
		maybePrintln(Z_interface, Z_interface_nil, Z_interface_typed_nil)
		maybePrintln(Z_map, Z_map_2, Z_map_empty, Z_map_nil)
		maybePrintln(Z_pointer, Z_pointer_nil)
		maybePrintln(Z_slice, Z_slice_2, Z_slice_nil)
		maybePrintln(Z_string, Z_struct)
		maybePrintln(Z_unsafe_pointer, Z_unsafe_pointer_nil)
		maybePrintln(Z_func_copy, Z_func_nil, Z_byte, Z_rune, Z_int_typedef)
	}
}

// maybePrintln won't print anything.  We just don't want the compiler to optimize it away.
func maybePrintln(v ...interface{}) {
	var t time.Time
	if time.Now() == t {
		fmt.Println(v...)
	}
}

// This is another version of the file.
