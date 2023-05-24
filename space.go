package gollision

import (
	"sync"

	"github.com/yanun0323/gollection/v2"
)

type Space interface {
	// Calculate collision of all bodies in this space
	Update()

	// Return true if the body which matches the ID is colliding something
	IsCollided(id uint64) bool

	// Get the other bodies colliding with the body which matches the ID
	GetCollided(id uint64) []Body

	// Get the next usable ID
	nextID() uint64

	// Add a body into this space
	addBody(b Body)

	// Remove a body from this space
	removeBody(id uint64)
}

type idSet = gollection.Set[uint64]

type space struct {
	lastBodyIDLock sync.Mutex
	lastBodyID     uint64

	collidedMap gollection.SyncMap[uint64, idSet]

	bodyMap gollection.SyncMap[uint64, Body]
}

func NewSpace() Space {
	return &space{
		lastBodyIDLock: sync.Mutex{},
		lastBodyID:     0,
		collidedMap:    gollection.NewSyncMap[uint64, idSet](),
		bodyMap:        gollection.NewSyncMap[uint64, Body](),
	}
}

func (s *space) Update() {
	s.calculateCollided()
}

func (s *space) IsCollided(id uint64) bool {
	m, ok := s.collidedMap.Load(id)
	if !ok {
		return false
	}
	return m.Len() != 0
}

func (s *space) GetCollided(id uint64) []Body {
	m, ok := s.collidedMap.Load(id)
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

func (s *space) nextID() uint64 {
	s.lastBodyIDLock.Lock()
	defer s.lastBodyIDLock.Unlock()
	s.lastBodyID++
	return s.lastBodyID
}

func (s *space) addBody(b Body) {
	s.bodyMap.Store(b.ID(), b)
}

func (s *space) removeBody(id uint64) {
	s.bodyMap.Delete(id)
}

func (s *space) calculateCollided() {
	s.collidedMap = gollection.NewSyncMap[uint64, idSet]()
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

			aBm := a.positionedBitmap()
			if aBm.isEmpty() {
				continue
			}

			bBm := b.positionedBitmap()
			if bBm.isEmpty() {
				continue
			}

			jBm := aBm.and(bBm)
			if jBm.isEmpty() {
				continue
			}

			if v, ok := s.collidedMap.Load(aID); !ok || v == nil {
				s.collidedMap.Store(aID, gollection.NewSet[uint64]())
			}
			v, _ := s.collidedMap.Load(aID)
			v.Insert(bID)

			if v, ok := s.collidedMap.Load(bID); !ok || v == nil {
				s.collidedMap.Store(bID, gollection.NewSet[uint64]())
			}
			v, _ = s.collidedMap.Load(bID)
			v.Insert(aID)
		}
	}
}
