// MIT License

// Copyright (c) 2020 Stephen Afam-Osemene

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package util

import (
	"context"
	"time"
)

// Every sends the time to the returned channel at the specified intervals
func Every(ctx context.Context, start time.Time, interval time.Duration) <-chan time.Time {
	// Create the channel which we will return
	stream := make(chan time.Time, 3)

	// Calculating the first start time in the future
	// Need to check if the time is zero (e.g. if time.Time{} was used)
	if !start.IsZero() {
		diff := time.Until(start)
		if diff < 0 {
			total := diff - interval
			times := total / interval * -1

			start = start.Add(times * interval)
		}
	}

	// Run this in a goroutine, or our function will block until the first event
	go func() {

		// Run the first event after it gets to the start time
		timer := time.NewTimer(time.Until(start))
		defer timer.Stop() // Make sure to stop the timer when we're done

		t := <-timer.C
		stream <- t

		// Open a new ticker
		ticker := time.NewTicker(interval)
		defer ticker.Stop() // Make sure to stop the ticker when we're done

		// Listen on both the ticker and the context done channel to know when to stop
		for {
			select {
			case t2 := <-ticker.C:
				stream <- t2
			case <-ctx.Done():
				close(stream)
				return
			}
		}
	}()

	return stream
}

// WaitUntil will block until the given time.
// Can be cancelled by cancelling the context
func WaitUntil(ctx context.Context, t time.Time) {
	diff := t.Sub(time.Now())
	if diff <= 0 {
		return
	}

	WaitFor(ctx, diff)
}

// WaitFor will block for the specified duration or the context is cancelled
func WaitFor(ctx context.Context, diff time.Duration) {
	timer := time.NewTimer(diff)
	defer timer.Stop()

	select {
	case <-timer.C:
		return
	case <-ctx.Done():
		return
	}
}

// BeginningOfMonth returns the beginning of the month given a date in that month
func BeginningOfMonth(date time.Time) time.Time {
	year, month, _ := date.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, date.Location())
}

// EndOfMonth returns the end of the month given a date in that month
func EndOfMonth(date time.Time) time.Time {
	year, month, _ := date.Date()
	return time.Date(year, month+1, 0, 23, 59, 59, 999999999, date.Location())
}

// BeginningOfDay returns the beginning of the day given a time in that day
func BeginningOfDay(date time.Time) time.Time {
	year, month, day := date.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, date.Location())
}

// EndOfDay returns the end of the day given a time in that day
func EndOfDay(date time.Time) time.Time {
	year, month, day := date.Date()
	return time.Date(year, month, day, 23, 59, 59, 999999999, date.Location())
}

// BeginningOfHour returns the beginning of the hour given an time in that hour
func BeginningOfHour(date time.Time) time.Time {
	year, month, day := date.Date()
	hour, _, _ := date.Clock()
	return time.Date(year, month, day, hour, 0, 0, 0, date.Location())
}

// EndOfHour returns the end of the hour given a time in that hour
func EndOfHour(date time.Time) time.Time {
	year, month, day := date.Date()
	hour, _, _ := date.Clock()
	return time.Date(year, month, day, hour, 59, 59, 999999999, date.Location())
}

// BeginningOfMinute returns the beginning of the minute given a time in that minute
func BeginningOfMinute(date time.Time) time.Time {
	year, month, day := date.Date()
	hour, minute, _ := date.Clock()
	return time.Date(year, month, day, hour, minute, 0, 0, date.Location())
}

// EndOfMinute returns the end of the minute given a time in that minute
func EndOfMinute(date time.Time) time.Time {
	year, month, day := date.Date()
	hour, minute, _ := date.Clock()
	return time.Date(year, month, day, hour, minute, 59, 999999999, date.Location())
}

// BeginningOfSecond returns the beginning of the second given a time in that second
func BeginningOfSecond(date time.Time) time.Time {
	year, month, day := date.Date()
	hour, minute, second := date.Clock()
	return time.Date(year, month, day, hour, minute, second, 0, date.Location())
}

// EndOfSecond returns the end of the second given a time in that second
func EndOfSecond(date time.Time) time.Time {
	year, month, day := date.Date()
	hour, minute, second := date.Clock()
	return time.Date(year, month, day, hour, minute, second, 999999999, date.Location())
}
