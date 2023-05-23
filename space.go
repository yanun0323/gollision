package gollision

import (
	"sync"

	"github.com/yanun0323/gollection/v2"
)

type Space interface {
	Update()
	GetCollided(id uint64) []Body

	nextID() uint64
	addBody(b Body)
	removeBody(id uint64, t Type)
}

type idSet = gollection.Set[uint64]

type space struct {
	lastBodyIDLock sync.Mutex
	lastBodyID     uint64

	collidedMap map[uint64]idSet

	bodyMap gollection.SyncMap[uint64, Body]
}

func NewSpace() Space {
	return &space{
		lastBodyIDLock: sync.Mutex{},
		lastBodyID:     0,
		collidedMap:    map[uint64]idSet{},
		bodyMap:        gollection.NewSyncMap[uint64, Body](),
	}
}

func (s *space) Update() {
	s.calculateCollided()
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

func (s *space) nextID() uint64 {
	s.lastBodyIDLock.Lock()
	defer s.lastBodyIDLock.Unlock()
	s.lastBodyID++
	return s.lastBodyID
}

func (s *space) addBody(b Body) {
	s.bodyMap.Store(b.ID(), b)
}

func (s *space) removeBody(id uint64, t Type) {
	s.bodyMap.Delete(id)
}

func (s *space) calculateCollided() {
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
