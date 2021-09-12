package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "good"
	str1 := "GOOD"
	str2 := "good good study, day day up"
	str3 := "gggggggggg"
	str4 := "   aaabbbaaa   "
	str5 := "a,b,c,d,e,"

	fmt.Println("\n前缀和后缀")
	fmt.Printf("src=%s func=%s target=%s result=%t\n", str, "HasPrefix", "go", strings.HasPrefix(str, "go"))
	fmt.Printf("src=%s func=%s target=%s result=%t\n", str, "HasSuffix", "od", strings.HasSuffix(str, "od"))
	
	fmt.Println("\n字符串包含关系")
	fmt.Printf("src=%s func=%s target=%s result=%t\n", str, "Contains", "oo", strings.Contains(str, "oo"))

	fmt.Println("\n判断子字符串或字符在父字符串中出现的位置（索引）")
	fmt.Printf("src=%s func=%s target=%s result=%d\n", str, "Index", "oo", strings.Index(str, "oo"))
	fmt.Printf("src=%s func=%s target=%s result=%d\n", str, "Index", "g", strings.Index(str, "g"))
	fmt.Printf("src=%s func=%s target=%s result=%d\n", str, "Index", "d", strings.Index(str, "d"))
	fmt.Printf("src=%s func=%s target=%s result=%d\n", str, "Index", "s", strings.Index(str, "s"))

	fmt.Println("\n字符串替换")
	fmt.Printf("src=%s func=%s target=%s result=%s\n", str, "ReplaceAll", "oo->ee", strings.ReplaceAll(str,"oo","ee"))
	fmt.Printf("src=%s func=%s target=%s result=%s\n", str, "Replace", "o->e 1", strings.Replace(str,"o","e",1))
	fmt.Printf("src=%s func=%s target=%s result=%s\n", str, "Replace", "o->e 2", strings.Replace(str,"o","e",2))
	
	fmt.Println("\n统计字符串出现次数")
	fmt.Printf("src=%s func=%s target=%s result=%d\n", str, "Count", "g", strings.Count(str, "g"))
	fmt.Printf("src=%s func=%s target=%s result=%d\n", str, "Count", "g", strings.Count(str, "o"))
	fmt.Printf("src=%s func=%s target=%s result=%d\n", str2, "Count", "o", strings.Count(str2, "o"))
	fmt.Printf("src=%s func=%s target=%s result=%d\n", str2, "Count", "oo", strings.Count(str2, "oo"))
	fmt.Printf("src=%s func=%s target=%s result=%d\n", str3, "Count", "g", strings.Count(str3, "g"))
	fmt.Printf("src=%s func=%s target=%s result=%d\n", str3, "Count", "gg", strings.Count(str3, "gg"))
	fmt.Printf("src=%s func=%s target=%s result=%d\n", str3, "Count", "ggg", strings.Count(str3, "ggg"))

	fmt.Println("\n重复字符串")
	fmt.Printf("src=%s func=%s target=%s result=%s\n", str, "Repeat", "3", strings.Repeat(str, 3))
	
	fmt.Println("\n修改字符串大小写")
	fmt.Printf("src=%s func=%s target=%s result=%s\n", str1, "ToLower", "", strings.ToLower(str1))
	fmt.Printf("src=%s func=%s target=%s result=%s\n", str, "ToUpper", "", strings.ToUpper(str))
	
	fmt.Println("\n修剪字符串")
	fmt.Printf("src=\"%s\" func=%s target=%s result=\"%s\"\n", str4, "TrimSpace", " ", strings.TrimSpace(str4))
	fmt.Printf("src=\"%s\" func=%s target=%s result=\"%s\"\n", str4, "Trim", " ", strings.Trim(str4, " "))
	
	fmt.Println("\n分割字符串")
	var rst1 = strings.Fields(str2)
	// for _, val := range rst1 {
	// 	fmt.Printf("src=\"%s\" val=%s\n",str2, val)
	// }
	for i := 0; i < len(rst1); i++ {
		fmt.Printf("src=\"%s\" index=%d val=%s\n",str2, i, rst1[i])
	}
	var rst2 = strings.Split(strings.Trim(str5,","), ",")
	for i := 0; i < len(rst2); i++ {
		fmt.Printf("src=\"%s\" index=%d val=%s\n",str5, i, rst2[i])
	}

	fmt.Println("\n拼接 slice 到字符串")
	fmt.Printf("src=\"%s\" func=%s target=%s result=\"%s\"\n", str5, "Join", "|", strings.Join(rst2,"|"))
	
}
