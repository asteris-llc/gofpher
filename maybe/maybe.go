// Copyright © 2016 Asteris, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package maybe

import (
	"fmt"

	"github.com/asteris-llc/gofpher/monad"
)

// Just is a plain value
type just struct {
	Val interface{}
}

// Nothing is the nil value for the maybe
type nothing struct{}

// Maybe provides an implementation of monad.Monad
type Maybe struct {
	internal interface{}
}

// AndThen provides the monadic implementation of AndThen
func (m Maybe) AndThen(f func(interface{}) monad.Monad) monad.Monad {
	if asJust, ok := m.internal.(just); ok {
		return f(asJust.Val)
	}
	return m
}

// LogAndThen provides the monadic implementation of LogAndThen
func (m Maybe) LogAndThen(f func(interface{}) monad.Monad, logger func(interface{})) monad.Monad {
	logger(m)
	if asJust, ok := m.internal.(just); ok {
		return f(asJust.Val)
	}
	return m
}

// Return provides the monadic implementation of Return
func (m Maybe) Return(i interface{}) monad.Monad {
	return Maybe{internal: just{Val: i}}
}

// Nothing creates a Nothing value
func Nothing() monad.Monad {
	return Maybe{internal: nothing{}}
}

// Just creates a just value
func Just(i interface{}) monad.Monad {
	return Maybe{internal: just{Val: i}}
}

// IsJust returns true if the value is a just value
func (m Maybe) IsJust() bool {
	_, ok := m.internal.(just)
	return ok
}

// FromJust gets the just value and panics if it's nothing
func (m Maybe) FromJust() interface{} {
	return m.internal.(just).Val
}

// FromMaybe gets the just value or returns the default
func (m Maybe) FromMaybe(defaultVal interface{}) interface{} {
	if m.IsJust() {
		return m.FromJust()
	}
	return defaultVal
}

func (m Maybe) String() string {
	if m.IsJust() {
		return fmt.Sprintf("Just %v", m.FromJust())
	}
	return "Nothing"
}
