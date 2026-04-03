package internal

import "fmt"

// ExecuteInstruction dispatches a single push-swap instruction string to the
// corresponding operation function, applying it to the given state.
//
// Recognised instructions: sa, sb, ss, pa, pb, ra, rb, rr, rra, rrb, rrr.
// Returns an error for any unrecognised instruction string.
func ExecuteInstruction(state *State, instruction string) error {
	switch instruction {
	case "sa":
		Sa(state)
	case "sb":
		Sb(state)
	case "ss":
		Ss(state)
	case "pa":
		Pa(state)
	case "pb":
		Pb(state)
	case "ra":
		Ra(state)
	case "rb":
		Rb(state)
	case "rr":
		Rr(state)
	case "rra":
		Rra(state)
	case "rrb":
		Rrb(state)
	case "rrr":
		Rrr(state)
	default:
		return fmt.Errorf("invalid instruction: %s", instruction)
	}
	return nil
}
