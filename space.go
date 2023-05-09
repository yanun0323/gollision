package gollision

import (
	"sync"

	"github.com/yanun0323/gollection/v2"
)

type idSet = gollection.Set[uint64]

type space struct {
	lastBodyIDLock sync.Mutex
	lastBodyID     uint64

	collidedMap map[uint64]idSet

	bodyMap gollection.SyncMap[uint64, Body]
}

func NewSpace() space {
	return space{
		lastBodyIDLock: sync.Mutex{},
		lastBodyID:     0,
		collidedMap:    map[uint64]idSet{},
		bodyMap:        gollection.NewSyncMap[uint64, Body](),
	}
}

func (s *space) Update() {
	s.calculateCollided()
}

func (s *space) NextID() uint64 {
	s.lastBodyIDLock.Lock()
	defer s.lastBodyIDLock.Unlock()
	s.lastBodyID++
	return s.lastBodyID
}

func (s *space) AddBody(b Body) {
	s.bodyMap.Store(b.ID(), b)
}

func (s *space) RemoveBody(id uint64, t Type) {
	s.bodyMap.Delete(id)
}

func (s *space) GetCollided(id uint64) []Body {
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

func (s *space) calculateCollided() {
	s.collidedMap = map[uint64]idSet{}
	var queue []Body
	s.bodyMap.Range(func(key uint64, b Body) bool {
		queue = append(queue, b)
		return true
	})
	println("queue:", len(queue))
	for i := 0; i < len(queue); i++ {
		a := queue[i]
		aID := a.ID()
		v := a.UpdatePosition(Vector{})
		println("Body ID:", aID, "X:", v.X, "Y:", v.Y, "Bitmap:", a.PositionedBitmap().Bitmap()[2])
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
