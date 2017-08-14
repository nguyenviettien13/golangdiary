package main

import (
	"fmt"
	"strings"
)

func Index(strs []string, t string) int {
	for i,v := range strs{
		if v==t {
			return i
		}
	}
	return -1
}

func Include(strs []string, t string) bool {
	if Index(strs,t) >0 {
		return true
	}else {
		return false
	}
}

//return true if one of element in slice satisfy the predicate 'f'
func Any(strs []string,f func(string) bool ) bool {
	for _, v := range strs{
		if f(v){
			return true
		}
	}
	return false
}

//return true if all of element in slice sastify the predicate 'f'
func All(strs []string, f func(string) bool) bool  {
	for _, v := range strs {
		if !f(v) {
			return false
		}
	}
	return  true
}

//return a new slice contains all string in the slice that satisfy predicate f
func Filter(strs []string,f func(string) bool) []string {
	fs := make([]string,0)
	for _, v := range strs{
		if f(v) {
			fs = append(fs, v)
		}
	}
	return fs
}
//return a new slice contain results of string after applying the function f to each string  in
//original string
func Map(strs[]string, f func(string) string) []string  {
	ms := make([]string,len(strs))
	for i,v := range strs{
		ms[i]= f(v)
	}
	return ms
}

func main()  {
	var strs =[]string{"hoanf","hoan","bang", "thuy","tien"}
	fmt.Println(Index(strs,"thuy"))
	fmt.Println(Include(strs,"Huong"))
	fmt.Println(Any(strs, func(s string) bool {
		return strings.HasPrefix(s, "t")
	}))
	fmt.Println(All(strs, func(s string) bool {
		return strings.HasPrefix(s, "h")
	}))
	fmt.Println(Filter(strs, func(v string) bool {
        return strings.Contains(v, "a")
    }))
	fmt.Println(Map(strs, strings.ToUpper))
}