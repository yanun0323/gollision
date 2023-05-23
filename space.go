package gollision

import (
	"sync"

	"github.com/yanun0323/gollection/v2"
)

type idSet = gollection.Set[uint64]

type Space struct {
	lastBodyIDLock sync.Mutex
	lastBodyID     uint64

	collidedMap map[uint64]idSet

	bodyMap gollection.SyncMap[uint64, Body]
}

func NewSpace() Space {
	return Space{
		lastBodyIDLock: sync.Mutex{},
		lastBodyID:     0,
		collidedMap:    map[uint64]idSet{},
		bodyMap:        gollection.NewSyncMap[uint64, Body](),
	}
}

func (s *Space) Update() {
	s.calculateCollided()
}

func (s *Space) NextID() uint64 {
	s.lastBodyIDLock.Lock()
	defer s.lastBodyIDLock.Unlock()
	s.lastBodyID++
	return s.lastBodyID
}

func (s *Space) AddBody(b Body) {
	s.bodyMap.Store(b.ID(), b)
}

func (s *Space) RemoveBody(id uint64, t Type) {
	s.bodyMap.Delete(id)
}

func (s *Space) GetCollided(id uint64) []Body {
	m, ok := s.collidedMap[id]
	if !ok {
		return []Body{}
	}
	var bodies []Body
	for _, id := range m.Iter() {
		if b, ok := s.bodyMap.Load(id); ok {
			bodies = append(bodies, b)
		}
	}
	return bodies
}

func (s *Space) calculateCollided() {
	s.collidedMap = map[uint64]idSet{}
	var queue []Body
	s.bodyMap.Range(func(key uint64, b Body) bool {
		queue = append(queue, b)
		return true
	})
	for i := 0; i < len(queue); i++ {
		a := queue[i]
		aID := a.ID()
		for j := i + 1; j < len(queue); j++ {
			b := queue[j]
			bID := b.ID()
			if aID == bID || a.Type() == b.Type() {
				continue
			}

			if a.PositionedBitmap().And(b.PositionedBitmap()).IsZero() {
				continue
			}

			if s.collidedMap[aID] == nil {
				s.collidedMap[aID] = gollection.NewSet[uint64]()
			}
			s.collidedMap[aID].Insert(bID)

			if s.collidedMap[bID] == nil {
				s.collidedMap[bID] = gollection.NewSet[uint64]()
			}
			s.collidedMap[bID].Insert(aID)
		}
	}
}
