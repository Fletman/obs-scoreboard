package locks

import "sync"

var Broadcast_Mutex sync.Mutex

var Score_Mutex sync.RWMutex

var Bracket_Mutex sync.RWMutex
