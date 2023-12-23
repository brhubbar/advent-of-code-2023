package main

import (
	"brh/aoc2023/internal/helpers"
	"fmt"
	"strings"
)

func Day20Part1(data string) int {
	n_modules := strings.Count(data, "\n")
	network := Network{
		moduleMap: make(map[string]uint8, n_modules),
		modules:   make([]Module, 0, n_modules),
		queue:     make([]Pulse, 0, len(data)),
		NLo:       0,
		NHi:       0,
	}

	var mod Module
	// Initialize all of the modules.
	for _, hookup := range strings.Split(data, "\n") {
		if strings.TrimSpace(hookup) == "" {
			continue
		}

		type_, hookup := hookup[0], hookup[1:]
		name, destination_str, _ := strings.Cut(hookup, " -> ")
		destinations := strings.Split(destination_str, ", ")

		switch type_ {
		case '%':
			mod = &FlipFlop{
				name:         name,
				destinations: destinations,
				state:        false,
			}
		case '&':
			mod = &Conjunction{
				name:         name,
				destinations: destinations,
				states:       make(map[string]bool, n_modules),
			}
		case 'b':
			// Broadcaster.
			mod = &Broadcaster{
				name:         "broadcaster",
				destinations: destinations,
			}
		default:
			panic("Unexpected type")
		}
		network.AddModule(mod)
	}

	// Register all module hookups.
	for name := range network.moduleMap {
		currentModule := network.GetModuleByName(name)
		for _, destName := range currentModule.GetDestinations() {
			// fmt.Printf("Attaching %v as a source to %v\n", name, destName)
			network.GetModuleByName(destName).AttachSource(name)
		}
	}

	for i := 0; i < 1000; i++ {
		// fmt.Println()
		network.queue = append(
			network.queue,
			Pulse{
				SourceModule:      "button",
				DestinationModule: "broadcaster",
				IsHi:              false,
			},
		)
		isQueue := true
		for isQueue {
			// fmt.Printf("\n%v\n", network.queue)
			isQueue = network.ProcessNextPulse(0)
		}
	}

	return int(network.NHi) * int(network.NLo)
}

func Day20Part2(data string) int {
	n_modules := strings.Count(data, "\n")
	network := Network{
		moduleMap: make(map[string]uint8, n_modules),
		modules:   make([]Module, 0, n_modules),
		queue:     make([]Pulse, 0, len(data)),
		NLo:       0,
		NHi:       0,
	}

	var mod Module
	var rxSourceConj *Conjunction
	// Initialize all of the modules.
	for _, hookup := range strings.Split(data, "\n") {
		if strings.TrimSpace(hookup) == "" {
			continue
		}

		type_, hookup := hookup[0], hookup[1:]
		name, destination_str, _ := strings.Cut(hookup, " -> ")
		destinations := strings.Split(destination_str, ", ")

		switch type_ {
		case '%':
			mod = &FlipFlop{
				name:         name,
				destinations: destinations,
				state:        false,
			}
		case '&':
			mod = &Conjunction{
				name:         name,
				destinations: destinations,
				states:       make(map[string]bool, n_modules),
			}
			if destinations[0] == "rx" {
				// Assume that one and only one conjunction provides an input to rx.
				rxSourceConj = mod.(*Conjunction)
			}
		case 'b':
			// Broadcaster.
			mod = &Broadcaster{
				name:         "broadcaster",
				destinations: destinations,
			}
		default:
			panic("Unexpected type")
		}
		network.AddModule(mod)
	}

	// Register all module hookups.
	for name := range network.moduleMap {
		currentModule := network.GetModuleByName(name)
		for _, destName := range currentModule.GetDestinations() {
			// fmt.Printf("Attaching %v as a source to %v\n", name, destName)
			network.GetModuleByName(destName).AttachSource(name)
		}
	}

	// Get ready to track the cycle length for each input to rxSourceConj
	network.NPresses = make(map[string]int, len(rxSourceConj.states))
	for source := range rxSourceConj.states {
		network.NPresses[source] = 0
	}
	network.RxSourceName = rxSourceConj.GetName()

	nBtnPresses := 0

run:
	for {
		nBtnPresses++
		// fmt.Println()
		network.queue = append(
			network.queue,
			Pulse{
				SourceModule:      "button",
				DestinationModule: "broadcaster",
				IsHi:              false,
			},
		)
		isQueue := true
		for isQueue {
			// fmt.Printf("\n%v\n", network.queue)
			isQueue = network.ProcessNextPulse(nBtnPresses)
		}
		// fmt.Printf("NLoRx: %v  ", network.NLoRx)
		for _, N := range network.NPresses {
			if N == 0 {
				// Not all of the cycle lengths have been found.
				continue run
			}
		}
		// All cycle lengths have been found.
		break
	}
	fmt.Printf("%v\n", network.NPresses)

	nPresses := make([]int, 0, len(network.NPresses))
	for _, value := range network.NPresses {
		nPresses = append(nPresses, value)
	}
	return helpers.LCM(nPresses...)
}

type Pulse struct {
	IsHi              bool
	SourceModule      string
	DestinationModule string
}

type Network struct {
	// Number of low pulses processed.
	NLo uint64
	// Number of high pulses processed.
	NHi uint64
	// Number of presses until a high is sent to the input conjunction for rx.
	RxSourceName string
	NPresses     map[string]int

	// Map of module names to their index in Modules (idx + 1)
	moduleMap map[string]uint8
	// Slice of the actual modules. Indexable using moduleMap.
	modules []Module
	// FIFO queue of pulses to process.
	queue []Pulse
}

type Module interface {
	GetName() string
	GetDestinations() []string
	Process(pulse bool, source string) []Pulse
	AttachSource(source string)
}

type FlipFlop struct {
	name         string
	state        bool
	destinations []string
}

type Conjunction struct {
	name         string
	states       map[string]bool
	destinations []string
}

type Broadcaster struct {
	name         string
	destinations []string
}

type NoOp struct{}

// Pops the next pulse out of the FIFO. Returns True if there are more pulses in the
// queue.
func (this *Network) ProcessNextPulse(nBtnPresses int) bool {
	// Pop the queue.
	pulse := this.queue[0]
	if len(this.queue) > 1 {
		this.queue = this.queue[1:]
	} else {
		// The queue is currently empty. The above does some weird memory overflow ish,
		// so just refresh the whole shebang.
		this.queue = make([]Pulse, 0, 5)
	}

	// Print for debug.
	// fmt.Printf("%v -%v-> %v\n", pulse.SourceModule, pulse.IsHi, pulse.DestinationModule)

	// Grab the destination module and send it the signal.
	result := this.GetModuleByName(pulse.DestinationModule).Process(pulse.IsHi, pulse.SourceModule)

	// fmt.Printf("Result: %v\n", result)
	// Add any resulting sends to the message queue.
	this.queue = append(this.queue, result...)

	// Count the highs/lows.
	if pulse.IsHi {
		this.NHi += 1
		if pulse.DestinationModule == this.RxSourceName && this.NPresses[pulse.SourceModule] == 0 {
			// Update the cycle length if this is the first time discovering it.
			this.NPresses[pulse.SourceModule] = nBtnPresses
		}
	} else {
		this.NLo += 1
	}

	return len(this.queue) > 0
}

func (this *Network) AddModule(mod Module) {
	// Add to the list.
	this.modules = append(this.modules, mod)
	// Save off where in the list this module will live (+1 to the actual address on
	// purpose).
	this.moduleMap[mod.GetName()] = uint8(len(this.modules))
}

// Get the module based on its name.
func (this *Network) GetModuleByName(name string) Module {
	// Module IDX are stored at a +1 from their actual position because map[invalidKey]
	// returns 0, which would overlap with the first value in the array.
	modIdx := this.moduleMap[name] - 1
	return this.GetModuleByIndex(modIdx)
}

// Get the module based on its location in the module array. Useful for iterating
// through all modules.
func (this *Network) GetModuleByIndex(idx uint8) Module {
	if idx >= uint8(len(this.modules)) {
		return &NoOp{}
	}
	return this.modules[idx]
}

// func (this *Network) GetStateHash() uint64 {

// }

// Flips when receiving a low pulse. Sends the new state to all destination modules.
// Does nothing on receiving a high pulse. `source` is ignored.
func (this *FlipFlop) Process(pulse bool, source string) (pulses []Pulse) {
	if pulse {
		return
	}
	this.state = !this.state
	for _, destination := range this.destinations {
		pulses = append(pulses, Pulse{this.state, this.name, destination})
	}
	return
}

// Records the value of the `pulse“ for a given `source`. If the last state of all
// known sources is high, then a low pulse is sent to all destination modules. If not, a
// high pulse is sent.
func (this *Conjunction) Process(pulse bool, source string) (pulses []Pulse) {
	this.states[source] = pulse
	// fmt.Printf("%v states: %v\n", this.GetName(), this.states)
	// Assume that all are high.
	send_pulse := false
	for _, state := range this.states {
		if !state {
			// One is not high; that's all it takes.
			send_pulse = true
			break
		}
	}

	for _, destination := range this.destinations {
		pulses = append(pulses, Pulse{send_pulse, this.name, destination})
	}
	return
}

// Sends low to all destinations when receiving a low signal.
func (this *Broadcaster) Process(pulse bool, source string) (pulses []Pulse) {
	if pulse {
		return
	}
	for _, destination := range this.destinations {
		pulses = append(pulses, Pulse{false, this.name, destination})
	}
	return
}

func (this *NoOp) Process(pulse bool, source string) (pulses []Pulse) { return }

// Do nothing. Source doesn't matter for flipflops.
func (this *FlipFlop) AttachSource(source string)    {}
func (this *Broadcaster) AttachSource(source string) {}
func (this *NoOp) AttachSource(source string)        {}

// Add the module to the list of sources and set its initial state to low.
func (this *Conjunction) AttachSource(source string) {
	this.states[source] = false
}

// Returns the module name.
func (this *FlipFlop) GetName() string    { return this.name }
func (this *Conjunction) GetName() string { return this.name }
func (this *Broadcaster) GetName() string { return this.name }
func (this *NoOp) GetName() string        { return "output" }

// Returns the module destinations.
func (this *FlipFlop) GetDestinations() []string    { return this.destinations }
func (this *Conjunction) GetDestinations() []string { return this.destinations }
func (this *Broadcaster) GetDestinations() []string { return this.destinations }
func (this *NoOp) GetDestinations() []string        { return make([]string, 0) }
