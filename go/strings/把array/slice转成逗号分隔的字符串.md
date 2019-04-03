```
array_or_slice := []int{1,2,3}
strings.Replace(strings.Trim(fmt.Sprint(array_or_slice), "[]"), " ", ",", -1)
```
结果："1,2,3"