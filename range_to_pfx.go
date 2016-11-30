package main

import (
	"fmt"
	"strconv"
	"strings"
)

type diapazon struct {
	start int
	stop  int
}

func read_from_text(s string) (R []diapazon, err error) {
	fmt.Println("start read_from_text")
	splitted := strings.Split(s, "\n")
	for i, s := range splitted {
		splitted[i] = strings.TrimRight(s, "\r")
	}
	//fmt.Println("input text", s)
	fmt.Println("splitted", splitted)
	for _, v := range splitted {
		//fmt.Println(i, v)
		if v != "" {
			ds := strings.Split(v, ",")
			if a, err := strconv.Atoi(ds[0]); err == nil {
				if b, err := strconv.Atoi(ds[1]); err == nil {
					d := diapazon{a, b}
					R = append(R, d)
				} else {
					fmt.Println("error in read_from_text", err)
					return R, err
				}
			} else {
				return R, err
			}
		}
	}
	return R, err
}
func FullCombine(r []diapazon) []diapazon {
	new_r := Combine(r)
	repeat := true
	if len(new_r) == len(r) {
		return new_r
	}
	for repeat {
		old_len := len(new_r)
		fmt.Println("here")
		new_r = Combine(new_r)
		if old_len == len(new_r) {
			repeat = false
			fmt.Println("exit")
		}
	}
	return new_r
}

//если конец диапазона - 1 = равно начало предыдущего то диапазон объединить
func Combine(r []diapazon) (new_r []diapazon) {
	/*
	   def combine(ranges):
	       #attension tis def modified incoming paramiters
	           i = 0
	           new = []
	           found = True
	           while found == True:
	               if i < len(ranges)-1:
	                   if (int(ranges[i][1]) + 1) == int(ranges[i+1][0]):
	                       r = [ranges[i][0],ranges[i+1][1]]
	                       ranges[i] = r
	                       a = ranges.pop(i+1)
	                   else:
	                       i += 1
	               else:
	                   found = False
	           return ranges
	*/
	i := 0
	found := true
	for found == true {
		if i < len(r)-1 {
			if r[i].stop+1 == r[i+1].start {
				d := diapazon{r[i].start, r[i+1].stop}
				//r[i] = d
				//надо удалить i+1 элемент из массива r
				//r = append(r[:i], r[i+1:]...)
				r = r[:i+copy(r[i:], r[i+1:])]
				r[i] = d
			} else {
				i += 1
			}
		} else {
			found = false
		}
	}
	new_r = r
	return new_r
	//	i := 1
	//	for i < len(r) {
	//		if r[i].start-1 == r[i-1].stop {
	//			d := diapazon{r[i-1].start, r[i].stop}
	//			new_r = append(new_r, d)
	//			//need to delete
	//			r = append(r[:i-1], r[i:]...)
	//		} else {
	//			new_r = append(new_r, r[i-1])
	//		}
	//		i = i + 1
	//	}
	//	if len(new_r) == 0 {
	//		return r
	//	} else {
	//		return new_r
	//	}
}

func appendDigit(str string, digit string, count int) string {
	result := []string{str}
	for i := 0; i < count-len(str); i++ {
		result = append(result, digit)
	}

	return strings.Join(result, "")
}

func range2prefix(start, stop string) []string {
	//Mihails Koreshkovs code
	var prefix string
	i := len(start) - 1
	for ; i >= 0; i-- {
		prefix = start[0:i]
		if strings.HasPrefix(stop, prefix) {
			break
		}
	}

	start_d, _ := strconv.Atoi(start[i : i+1])
	stop_d, _ := strconv.Atoi(stop[i : i+1])

	var prefix_list []string
	for ; start_d <= stop_d; start_d++ {
		prefix_list = append(prefix_list, prefix+strconv.Itoa(start_d))
	}
	//	fmt.Println(prefix)
	//	fmt.Println(prefix_list)

	l := len(prefix_list)
	if l >= 2 {
		if appendDigit(prefix_list[0], "0", len(start)) != start {
			t1 := range2prefix(start, appendDigit(prefix_list[0], "9", len(start)))
			prefix_list = append(t1, prefix_list[1:]...)
		}

		l = len(prefix_list)
		if appendDigit(prefix_list[l-1], "9", len(stop)) != stop {
			t2 := range2prefix(appendDigit(prefix_list[l-1], "0", len(stop)), stop)
			prefix_list = append(prefix_list[:l-1], t2...)
		}

	}
	//code for create prefix smoller ex. 10,11,12,13,14,15,16,17,18,19 => 1
	if len(prefix_list) == 10 {
		tl := (len(prefix_list[0]))
		do_next := true
		for _, p := range prefix_list {
			if tl != len(p) {
				do_next = false
			}
		}
		if do_next {
			if (prefix_list[0][tl-1:] == "0") && (prefix_list[1][tl-1:] == "1") && (prefix_list[2][tl-1:] == "2") && (prefix_list[3][tl-1:] == "3") && (prefix_list[4][tl-1:] == "4") && (prefix_list[5][tl-1:] == "5") && (prefix_list[6][tl-1:] == "6") && (prefix_list[7][tl-1:] == "7") && (prefix_list[8][tl-1:] == "8") && (prefix_list[9][tl-1:] == "9") {
				return []string{prefix_list[0][:tl-1]}
			}
		}
	}
	fmt.Println("pfx=", prefix_list)
	return prefix_list

}

func list2prefix(r []diapazon) (p_list []string) {
	for _, d := range r {
		a := strconv.Itoa(d.start)
		b := strconv.Itoa(d.stop)
		p_list = append(p_list, range2prefix(a, b)...)
	}
	return p_list
}

/*
func main() {
	fmt.Println("start")
	var operator_cap []diapazon
	operator_cap = read_from_text(s)
	//fmt.Println(operator_cap)
	fmt.Println("after combine")
	//combined1 := Combine(operator_cap)
	combined1 := FullCombine(operator_cap)
	//fmt.Println(combined1)
	fmt.Println("combain again")
	//combined2 := Combine(combined1)
	//fmt.Println(combined2)
	//	combined3 := Combine(combined2)
	//	combined4 := Combine(combined3)
	//	fmt.Println("found prefixes")
	results := list2prefix(combined1)
	fmt.Println(results)
	//	for _, v := range results {
	//		fmt.Println(v)
	//	}
	//_ = combined2
	f, err := os.Create("pfx.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	for _, v := range results {
		_, _ = f.WriteString(v + "\r\n")
	}
}
*/
var s string = `
0219100000,0219199999
0219900000,0219939999
0219950000,0219999999
0267550000,0267999999
`
