package Simple

import "testing"

func TestNewRestaurant(t *testing.T) {
	NewRestaurant("K").GetFood()
	NewRestaurant("L").GetFood()
}
