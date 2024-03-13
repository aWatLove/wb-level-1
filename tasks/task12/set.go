package main

import "fmt"

// Set структура, которая содержит мапу, где ключ - string, а значение - пустая структура, т.к. значение нам не нужно
type Set struct {
	values map[string]struct{}
}

// конструктор для структуры Set
func NewSet() *Set {
	return &Set{values: make(map[string]struct{})}
}

// метод добавления элемента в Set
func (s Set) Add(elem string) bool {
	s.values[elem] = struct{}{}
	return true
}

// метод, который проверяет, содержится ли элементе в Set'е
func (s Set) Contains(elem string) bool {
	_, ok := s.values[elem]
	if ok {
		return true
	}
	return false
}

// метод удаляющий элемент из Set'а
func (s Set) Remove(elem string) {
	delete(s.values, elem)
}

// метод преобразовывающий Set в слайс строк
func (s Set) ToArray() []string {
	res := make([]string, 0)
	for k, _ := range s.values {
		res = append(res, k)
	}
	return res
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"} // слайс со строками
	set := NewSet()                                     // создаем новый сет
	for _, e := range arr {                             // цикл, в котором добавляем элементы из слайса в сет
		set.Add(e)
	}
	// проверяем есть ли "dog" в сете
	fmt.Println(set.Contains("dog")) // true
	// проверяем есть ли "dog" в сете
	fmt.Println(set.Contains("bear")) // false
	fmt.Println(set.ToArray())        // [cat dog tree]
}
