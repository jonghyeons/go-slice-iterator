package go_slice_iterator

import (
	"fmt"
	"testing"
)

/*
TODO
메모리상에서 떨어져 있는 슬라이스를
붙어있는 메모리 변수로(배열로) 바꿔서
L1 캐시 히트율을 올려보자

아주 긴 리스트에 적용해야 속도 향상이 체감될껄?

iterator := userCollection.createIterator()
for iterator.hasNext() {
	user := iterator.getNext()
	fmt.Printf("User is %+v\n", user)
}

이런 식으로 쓰게 만들자
*/

// ------------------------------------
type Person struct {
	Age    int
	Name   string
	Job    string
	Email  string
	Region string
	//...
}

func TestIterate(t *testing.T) {
	var pointerPeople []*Person
	for i := 0; i < 100; i++ {
		pointerPeople = append(pointerPeople, &Person{})
	}
	fmt.Println(pointerPeople)

	var valuePeople []Person
	for i := range pointerPeople {
		valuePeople = append(valuePeople, *pointerPeople[i])
	}
	fmt.Println(valuePeople)

	fmt.Println("***************************")

	itr := GetIterator(pointerPeople)
	for itr.HasNext() {
		v := itr.GetNext()
		fmt.Println(v)
	}
}

func BenchmarkDefaultIterator(b *testing.B) {
	var people []*Person
	for i := 0; i < b.N; i++ {
		people = append(people, &Person{})
	}

	for i := range people {
		people[i].Age++
	}
}

func TestCustomIterator(t *testing.T) {
	var people []*Person
	for i := 0; i < 100; i++ {
		people = append(people, &Person{})
	}

	iterator := GetIterator(people)
	for iterator.HasNext() {
		person := iterator.GetNext()
		fmt.Println(*person)
	}
}

func BenchmarkCustomIterator(b *testing.B) {
	var people []*Person
	for i := 0; i < b.N; i++ {
		people = append(people, &Person{})
	}

	iterator := GetIterator(people)
	for iterator.HasNext() {
		person := *iterator.GetNext()
		person.Age++
	}
}
