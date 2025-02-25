// Create by Yale 2019/2/25 16:08
package example

import (
	"fmt"
	"github.com/yale8848/stream"
	"strings"
	"testing"
)

var data =[]person{{age:11,name:"alice"},{age:19,name:"pig"},{age:5,name:"cat"},{age:21,name:"bob"},{age:13,name:"pig"},{age:6,name:"lili"}}

func st() stream.Stream  {
	st,_:=stream.Of(data)
	return st.Filter(func(v stream.T) bool {
		p:=v.(person)
		fmt.Printf("Filter %v\r\n",v)
		return  p.age>10
	}).Map(func(v stream.T) stream.T {
		p:=v.(person)
		p.name = strings.ToUpper(p.name)
		fmt.Printf("Mapper %v\r\n",p)
		return p
	}).Peek(func(v stream.T) {
		fmt.Printf("Peek %v\r\n",v)
	}).Sorted(func(v1, v2 stream.T) bool {
		va:=v1.(person)
		vb:=v2.(person)
		return strings.Compare(va.name,vb.name)<0
	})
}

func TestForEach(t *testing.T) {
	st:=st()
	st.ForEach(func(v stream.T)bool {
		fmt.Printf("ForEach %v\r\n",v)
		return true
	})
}
func TestCount(t *testing.T)  {
	fmt.Println(st().Count())
}
func TestFindFirst(t *testing.T)  {
	fmt.Println(st().FindFirst())
}
func TestCollect(t *testing.T)  {
	fmt.Println(st().Collect())
}
func TestSkip(t *testing.T)  {
	fmt.Println(st().Skip(2).Collect())
}
func TestLimit(t *testing.T)  {
	fmt.Println(st().Skip(0).Limit(3).Collect())
}
func TestDistinct(t *testing.T)  {
	fmt.Println(st().Skip(1).Limit(4).Distinct(func(v stream.T) stream.T {
		v1:=v.(person)
		return v1.name
	}).Collect())
}
func TestMax(t *testing.T)  {
	fmt.Println(st().Skip(0).Limit(3).Max(func(max stream.T, v stream.T) bool {
		m:=max.(person)
		v1:=v.(person)
		return m.age<v1.age
	}))
}
func TestMin(t *testing.T)  {
	fmt.Println(st().Skip(0).Limit(3).Min(func(min stream.T, v stream.T) bool {
		m:=min.(person)
		v1:=v.(person)
		return m.age>v1.age
	}))
}
func TestAnyMatch(t *testing.T)  {
	fmt.Println(st().Skip(0).Limit(3).AnyMatch(func(v stream.T) bool {
		v1:=v.(person)
		return strings.Contains(v1.name,"I")
	}))
}
func TestAllMatch(t *testing.T)  {
	fmt.Println(st().Skip(0).Limit(3).AllMatch(func(v stream.T) bool {
		v1:=v.(person)
		return strings.Contains(v1.name,"I")
	}))
}
func TestNoneMatch(t *testing.T)  {
	fmt.Println(st().Skip(0).Limit(3).NoneMatch(func(v stream.T) bool {
		v1:=v.(person)
		return strings.Contains(v1.name,"I")
	}))
}
func TestGroup(t *testing.T)  {
	st().Group(7, func(ts []stream.T) {
		fmt.Println(ts)
	})
}
func TestGroupRoutine(t *testing.T)  {
	st().GroupRoutine(2, func(ts []stream.T) {
		fmt.Println(ts)
	}, func(err interface{}) {
		fmt.Println(err)
	})
}
func TestSum(t *testing.T)  {
	fmt.Println(st().Sum(func(v stream.T) int64 {
		v1:=v.(person)
		return int64(v1.age)
	}))
}


func TestReduce(t *testing.T) {
	s, _ := stream.Of(data)
	reduce := s.Filter(func(v stream.T) bool {
		return v.(person).age > 5
	}).Reduce(0, func(x, y stream.T) stream.T {
		return x.(int) + y.(person).age
	})
	fmt.Printf("TestReduce result is %d", reduce)
}

func TestReduce2(t *testing.T) {
	s, _ := stream.Of(data)
	reduce := s.Filter(func(v stream.T) bool {
		return v.(person).age > 5
	}).Map(func(v stream.T) stream.T {
		return v.(person).age
	}).Reduce(0, func(x, y stream.T) stream.T {
		return x.(int) + y.(int)
	})
	fmt.Printf("TestReduce2 result is %d", reduce)
}