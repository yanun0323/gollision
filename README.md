# Gollision

A simple collision package.

# Requirement
_GO 1.20 or Higher_

# Usage

```go
// Init space
space := NewSpace()

// Create body
CollisionType := gollision.Type(1)
bodyA := NewBody(space, CollisionType)

// Update body image
imageData := [][]uint8{...}
bodyA.UpdateBitmap(imageHeight, imageWidth, imageData)

// Update body position
bodyA.SetPosition(10, 10) /* Change position to (10, 10) */
bodyA.AddPosition(10, 10) /* Change position to (originalX + 10, originalY + 10) */

// Calculate body collisions in the space
space.Update()

// Get the other bodies colliding with this body
collided := bodyA.GetCollided() 
collided := space.GetCollided(bodyA.ID())

```
