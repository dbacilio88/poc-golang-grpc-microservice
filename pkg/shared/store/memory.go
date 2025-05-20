package store

/**
 * memory
 * <p>
 * This file contains core data structures and logic used throughout the application.
 *
 * <p><strong>Copyright © 2025 – All rights reserved.</strong></p>
 *
 * <p>This source code is distributed under a collaborative license.</p>
 *
 * <p>
 * Contributions, suggestions, and improvements are welcome!
 * You are free to fork, modify, and submit pull requests under the terms of the repository's license.
 * Please ensure proper attribution to the original author(s) and preserve this notice in derivative works.
 * </p>
 *
 * @author Christian Bacilio De La Cruz
 * @email dbacilio88@outlook.es
 * @since 5/7/2025
 */

type Memory struct {
	Key string
	Val string
}

type MemoryMap struct {
	items   map[string]string
	addChan chan Memory
	getChan chan string
}

func (m *MemoryMap) run() {

	for {
		select {
		case item := <-m.addChan:
			m.items[item.Key] = item.Val
		case key := <-m.getChan:
			val, ok := m.items[key]
			if !ok {
				val = ""
			}
			m.getChan <- val
		}
	}

}

func NewMemory() *MemoryMap {
	store := &MemoryMap{
		items:   make(map[string]string),
		addChan: make(chan Memory),
		getChan: make(chan string),
	}
	go store.run()
	return store
}

func (m *MemoryMap) Put(key string, val string) {
	m.addChan <- Memory{
		Key: key,
		Val: val,
	}
}

func (m *MemoryMap) Get(key string) string {
	m.getChan <- key
	return <-m.getChan
}
