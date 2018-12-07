package main

import (
	"fmt"
	"strings"
)

func main() {
	prefix := strings.HasPrefix("让我一次爱个够", "让")
	suffix := strings.HasSuffix("让我一次爱个够", "够够")
	contains := strings.Contains("让我一次爱个够", "爱个够")
	index := strings.Index("让我一次爱个够", "爱个够")
	count := strings.Count("123213123123123", "12")
	repeat := strings.Repeat("微辣．猪", 10)
	lower := strings.ToLower("HELLO WOELD")
	upper := strings.ToUpper(lower)
	space := strings.TrimSpace("  dadadsdad ads  ad   ")
	trim := strings.Trim(space, "ad")
	fmt.Print(prefix,suffix,contains,index,count,"\n",repeat,"\n",lower,upper,
		"\n",space,"\n",trim,"\n")
	fields := strings.Fields("I am very shy, and I don't know how to resolve?")

	for index, val := range fields {
		fmt.Printf("  index:%d,value:%s  ",index,val)
	}
	fmt.Println()
	split := strings.Split("if,you,do,you,will,go,die", ",")
	for index, val := range split {
		fmt.Printf("  index:%d,value:%s  ",index,val)
	}


	fmt.Print(fields,"\n","\n")
	reader := strings.NewReader("This is a Reader")
	fmt.Print(*reader)
}
