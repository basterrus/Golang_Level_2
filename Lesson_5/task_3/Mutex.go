package task_3

import "sync"

type SetUpStruct struct {
	sync.Mutex
	m map[int]struct{}
}

func SetUp() *SetUpStruct {
	return &SetUpStruct{
		m: map[int]struct{}{},
	}

}

func (s *SetUpStruct) Add(i int) {
	s.Lock()
	s.m[i] = struct{}{}
	s.Unlock()
}

func (s *SetUpStruct) Has(i int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.m[i]
	return ok
}

//3. Протестируйте производительность операций чтения и записи на множестве
//действительных чисел, безопасность которого обеспечивается sync.Mutex и
//sync.RWMutex для разных вариантов использования: 10% запись, 90% чтение; 50%
//запись, 50% чтение; 90% запись, 10% чтение
