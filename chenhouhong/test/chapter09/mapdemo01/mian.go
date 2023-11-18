package main

import "fmt"

func main() {
	var m map[string]string
	fmt.Printf("m: %v\n", m)
	//分配内存，map会覆盖
	//会自动排序，会自动扩容
	m = make(map[string]string, 2)
	m["n1"] = "hello"
	m["n3"] = "world"
	m["n6"] = "!!!"
	m["n1"] = "Hello"
	fmt.Printf("m: %v\n", m)

	//遍历是乱序的
	for k, v := range m {
		fmt.Printf("key : %v---value : %v", k, v)
	}

	//第二种定义方式
	mapper := make(map[string]string, 10)
	mapper["NO1"] = "李白"
	mapper["NO2"] = "杜甫"
	fmt.Printf("mapper: %v\n", mapper)

	//第三种
	//最后也要加逗号，不可key重复
	mm := map[string]string{
		"no1": "hello",
		"no2": "world",
		"no3": "pick",
	}
	fmt.Printf("mm: %v\n", mm)

	//复杂案例
	mapping := make(map[string]map[string]string, 10)
	mapping["stu01"] = make(map[string]string, 3)
	mapping["stu02"] = make(map[string]string, 3)

	mapping["stu01"]["name"] = "tom"
	mapping["stu01"]["age"] = "19"
	mapping["stu01"]["address"] = "guangzhou"

	mapping["stu02"]["name"] = "mike"
	mapping["stu02"]["age"] = "99"
	mapping["stu02"]["address"] = "shenzhen"

	//根据key删除
	delete(mapping["stu01"], "name")

	//查找
	val, ok := mapping["stu01"]
	if ok {
		fmt.Printf("找到了val: %v\n", val)
	} else {
		fmt.Println("没抄到")
	}
	fmt.Println("xiugaile")
	//for - range 遍历
	for k1, v1 := range mapping {
		fmt.Printf("k1 : %v\n", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\tk2 : %v , v2 : %v", k2, v2)
		}
		fmt.Println()
	}
	fmt.Printf("mapping: %v\n", mapping)
	//长度
	fmt.Printf("len(mapping): %v\n", len(mapping))

	noteFrequency := map[string]float32{
		"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
		"G0": 24.50, "A0": 27.50, "B0": 30.87, "A4": 440} //紧跟在后面就不需要加，
	fmt.Printf("noteFrequency: %v\n", noteFrequency)

	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}
}
