## Golang stream lib is like Java 8 stream. Only handle slice or array.

## demo

```go

   var data =[]person{{age:11,name:"alice"},{age:19,name:"pig"},{age:5,name:"cat"},{age:21,name:"bob"}}
   
   stream.Of(data).Filter(func(v stream.T) bool {
   		p:=v.(person)
   		fmt.Printf("Filter %v\r\n",v)
   		return  p.age>10
   	}).Peek(func(v stream.T) {
   		fmt.Printf("Peek %v\r\n",v)
   	}).Skip(1).Map(func(v stream.T) stream.T {
   		p:=v.(person)
   		p.name = strings.ToUpper(p.name)
   		return p
   	}).Sorted(func(v1, v2 stream.T) bool {
   		s1:=v1.(person)
   		s2:=v2.(person)
   		return  strings.Compare(s1.name,s2.name)<0
   	}).Limit(2).Distinct(func(v stream.T) stream.T {
   		p:=v.(person)
   		fmt.Printf("Distinct %v\r\n",v)
   		return p.name
   	}).ForEach(func(v stream.T) {
      		fmt.Printf("ForEach %v\r\n",v)
      	})

```

## Todo

### Intermediate

- [x] Filter 
- [x] Map
- [x] peek 

- [x] Sorted 
- [x] Distinct 
- [ ] unordered 

- [x] Skip 

- [ ] parallel 
- [ ] sequential 



### Terminal

- [x] ForEach 
- [ ] ForEachOrdered 
- [ ] toArray 
- [x] Collect 
- [x] Min 
- [x] Max
- [x] Count
- [ ] iterator

### Short-circuiting

- [x] AnyMatch
- [x] AllMatch
- [x] NoneMatch
- [x] FindFirst
- [ ] findAny
- [x] Limit


## ref

http://www.cnblogs.com/CarpenterLee/p/6637118.html

https://www.cnblogs.com/Dorae/p/7779246.html

https://zhuanlan.zhihu.com/p/33313312

