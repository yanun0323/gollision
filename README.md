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
body := NewBody(space, CollisionType)

// Update body image
imageData := [][]uint8{...}
body.UpdateBitmap(imageHeight, imageWidth, imageData)

// Update body position
body.SetPosition(10, 10) /* Change position to (10, 10) */
body.AddPosition(10, 10) /* Change position to (origin x + 10, origin y + 10) */

// Calculate body collisions in the space
space.Update()

// Get the other bodies colliding with this body
collided := body.GetCollided() 
collided := space.GetCollided(body.ID())

```
