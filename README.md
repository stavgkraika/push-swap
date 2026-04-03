# push-swap

A two-stack sorting checker written in Go. It reads an initial sequence of integers from command-line arguments, applies push-swap instructions from stdin, and prints `OK` if the result is sorted or `KO` otherwise.

---

## Project structure

```
push-swap/
├── cmd/
│   └── main.go          # Entry point
└── internal/
    ├── state.go         # Two-stack state
    ├── stack.go         # Stack data structure
    ├── ops.go           # Push-swap operations
    ├── exec.go          # Instruction dispatcher
    ├── parse.go         # Argument parsing
    └── validate.go      # Sorting validation
```

---

## cmd/main.go

Entry point of the program. Wires the real OS streams into `run` and exits with its return code.

| Function | Description |
|----------|-------------|
| `main()` | Calls `run` with `os.Args`, `os.Stdin`, `os.Stdout`, and `os.Stderr`, then exits with the returned code. |
| `run(args, stdin, stdout, stderr)` | Contains all program logic. Parses arguments, initialises state, reads and executes instructions from stdin line by line, and prints `OK` or `KO`. Returns `0` on success, `1` on any error. Accepts injectable streams so it can be tested without spawning a subprocess. |

---

## internal/state.go

Defines the shared state passed to every push-swap operation, holding both stacks.

| Type / Function | Description |
|-----------------|-------------|
| `State` | Struct with two fields: `A *Stack` (the stack to be sorted) and `B *Stack` (the auxiliary stack). |
| `NewState(values)` | Creates a `State` with the given values loaded into stack A and an empty stack B. |

---

## internal/stack.go

Implements the `Stack` data structure used by both stack A and stack B. The top element is always at index 0 of the underlying slice.

| Type / Function | Description |
|-----------------|-------------|
| `Stack` | LIFO data structure backed by a `[]int` slice. |
| `NewStack(values)` | Creates a Stack pre-loaded with the given values. Copies the slice so the caller's original is never mutated. |
| `PushTop(v)` | Inserts `v` at the top (index 0) of the stack. |
| `PopTop()` | Removes and returns the top element. Returns `(0, false)` if the stack is empty. |
| `PeekTop()` | Returns the top element without removing it. Returns `(0, false)` if the stack is empty. |
| `Size()` | Returns the number of elements in the stack. |
| `IsEmpty()` | Reports whether the stack contains no elements. |
| `Values()` | Returns a copy of all elements in top-to-bottom order. |
| `SwapTopTwo()` | Swaps the two topmost elements. No-op if fewer than 2 elements. |
| `Rotate()` | Moves the top element to the bottom (`ra`/`rb`). No-op if fewer than 2 elements. |
| `ReverseRotate()` | Moves the bottom element to the top (`rra`/`rrb`). No-op if fewer than 2 elements. |

---

## internal/ops.go

Implements all eleven push-swap operations. Each function takes a `*State` and mutates one or both stacks.

| Function | Description |
|----------|-------------|
| `Sa(s)` | Swaps the top two elements of stack A. |
| `Sb(s)` | Swaps the top two elements of stack B. |
| `Ss(s)` | Performs `Sa` and `Sb` simultaneously. |
| `Pa(s)` | Pops the top of stack B and pushes it onto stack A. No-op if B is empty. |
| `Pb(s)` | Pops the top of stack A and pushes it onto stack B. No-op if A is empty. |
| `Ra(s)` | Rotates stack A upward: top element moves to the bottom. |
| `Rb(s)` | Rotates stack B upward: top element moves to the bottom. |
| `Rr(s)` | Performs `Ra` and `Rb` simultaneously. |
| `Rra(s)` | Reverse-rotates stack A: bottom element moves to the top. |
| `Rrb(s)` | Reverse-rotates stack B: bottom element moves to the top. |
| `Rrr(s)` | Performs `Rra` and `Rrb` simultaneously. |

---

## internal/exec.go

Dispatches a single instruction string to the corresponding operation function.

| Function | Description |
|----------|-------------|
| `ExecuteInstruction(state, instruction)` | Matches the instruction string against all eleven valid push-swap instructions and calls the corresponding function. Returns an error for any unrecognised instruction. |

---

## internal/parse.go

Parses and validates raw command-line arguments into a slice of unique integers.

| Symbol | Description |
|--------|-------------|
| `ErrInvalidInteger` | Sentinel error returned when a token cannot be parsed as a base-10 integer. |
| `ErrDuplicateValue` | Sentinel error returned when the same integer appears more than once. |
| `ParseArgs(args)` | Flattens all arguments into individual tokens (supporting both `"3" "1" "2"` and `"3 1 2"` forms), parses each as an integer, rejects duplicates, and returns the ordered slice. Returns an empty non-nil slice when `args` is empty. |

---

## internal/validate.go

Provides a helper to check whether a slice is in sorted order.

| Function | Description |
|----------|-------------|
| `IsSortedAscending(values)` | Returns `true` if the slice is sorted in ascending order, `false` otherwise. |
