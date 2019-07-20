//* package counters provides alert counter support

package counters

//* alertCounter is an unexported type that contains
//* an integer counter for alerts
type alertCounter int

//* function New creates and returns the value of the
//* unexported identifier => type alert
func New(value int) alertCounter {
	return alertCounter(value)
}
