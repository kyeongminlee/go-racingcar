package domain

type Car struct {
	name     string
	position int
}

func (c Car) Run(movementStrategy MovementStrategy) Car {
	if movementStrategy.Movement() {
		c.position += 1
	}

	return c
}

func (c *Car) GetName() string {
	return c.name
}

func (c *Car) GetCurrentPosition() int {
	return c.position
}
