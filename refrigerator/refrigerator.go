package refrigerator

import (
	"math/rand"
	"time"
)

// All meat is stored in the refrigerator 
type refrigerator struct {
	meat []Meat
}

// There are different kinds of meat
type Meat struct {
	Name        string // The name of the meat 
	OperateTime time.Duration // Time Required to Process the meat 
	Num         int // The number of the meat 
}

type iRefrigerator interface {
	Get() Meat
	Remove(s int) []Meat
	Add(meat string , operateTime time.Duration , num int)
}

func NewRefrigerator() iRefrigerator {
	return &refrigerator{}
}

// Get the meat from the refrigerator
func (m *refrigerator) Get() Meat {

	// Get the meat by random
	randNum := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(m.meat))

	// Calculate remaining quantity 
	m.meat[randNum].Num = m.meat[randNum].Num - 1
	rawMeat := m.meat[randNum]

	// if the remaining quantity = 0 , remove the kind of meat from refrigerator
	if m.meat[randNum].Num == 0 {
		m.meat = m.Remove(randNum)
	}
	return rawMeat
}

func (m *refrigerator) Remove(s int) []Meat {
	return append(m.meat[:s], m.meat[s+1:]...)
}

func (m *refrigerator) Add(meat string , operateTime time.Duration , num int){
	newMeat := Meat{
		Name:        meat,
		OperateTime: operateTime,
		Num:         num,
	}
	m.meat = append(m.meat , newMeat)
}