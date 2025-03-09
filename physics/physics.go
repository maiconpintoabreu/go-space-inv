package physics

import (
	"github.com/maiconspas/go-space-inv/utils"
)

const arenaSize int = 100

type PhysicObject struct {
	Position utils.Vector2d
}

var objects [arenaSize]PhysicObject
var ObjectCount int = 0

func CreatePhysicObject(position utils.Vector2d) {
	objects[ObjectCount] = PhysicObject{Position: position}
	ObjectCount++
}
func RemovePhysicObject(index int) {
	if ObjectCount <= 1 {
		ObjectCount = 0
		return
	}
	objects[index] = objects[ObjectCount-1]
	ObjectCount--

}

func StartPhysicSystem() bool {
	return true
}
func GetObjects() *[arenaSize]PhysicObject {
	return &objects
}

func Update() {
	// Update all objects that are alive
	for range ObjectCount {

	}
}
