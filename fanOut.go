package main

type fanOut func(pipe, []pipe)

// split implements a round-roubin scheduled fan-out
func split(ch pipe, n int, fanout fanOut) []pipe {
	cs := make([]pipe, 0)

	for i := 0; i < n; i++ {
		cs = append(cs, make(pipe))
	}

	go fanout(ch, cs)

	return cs
}

func roundRobin(ch pipe, cs []pipe) {
	defer func(cs []pipe) { // at the end, close all incoming channels
		for _, c := range cs {
			close(c)
		}
	}(cs)

	for {
		for _, c := range cs {
			select {
			case val, ok := <-ch:
				if !ok {
					return
				}

				c <- val
			}
		}
	}
}
