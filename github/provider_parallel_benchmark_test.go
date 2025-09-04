package github

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func BenchmarkProvider_ParallelRequests(b *testing.B) {
	b.Run("Serial", func(b *testing.B) {
		benchmarkRequests(b, false)
	})
	
	b.Run("Parallel", func(b *testing.B) {
		benchmarkRequests(b, true)
	})
}

func benchmarkRequests(b *testing.B, parallel bool) {
	// Create a mock server that simulates API latency
	var requestCount int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&requestCount, 1)
		// Simulate API latency
		time.Sleep(50 * time.Millisecond)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"id": 1, "name": "test-repo"}`)
	}))
	defer server.Close()

	// Create transport
	client := http.DefaultClient
	if parallel {
		client.Transport = NewRateLimitTransport(http.DefaultTransport, WithParallelRequests(true))
	} else {
		client.Transport = NewRateLimitTransport(http.DefaultTransport, WithParallelRequests(false))
	}

	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		// Reset counter
		atomic.StoreInt32(&requestCount, 0)
		
		// Simulate multiple requests like Terraform would make
		numRequests := 10
		
		if parallel {
			var wg sync.WaitGroup
			for j := 0; j < numRequests; j++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					req, _ := http.NewRequest("GET", server.URL+"/repos/test/test", nil)
					client.Do(req)
				}()
			}
			wg.Wait()
		} else {
			for j := 0; j < numRequests; j++ {
				req, _ := http.NewRequest("GET", server.URL+"/repos/test/test", nil)
				client.Do(req)
			}
		}
		
		// Verify all requests were made
		if int(atomic.LoadInt32(&requestCount)) != numRequests {
			b.Errorf("Expected %d requests, got %d", numRequests, requestCount)
		}
	}
}

func TestProvider_ParallelRequestsSpeedup(t *testing.T) {
	// Create a mock server with controlled latency
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate 100ms API latency
		time.Sleep(100 * time.Millisecond)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"id": 1, "name": "test-repo", "description": "test"}`)
	}))
	defer server.Close()

	numRequests := 10

	// Test serial requests
	serialClient := &http.Client{
		Transport: NewRateLimitTransport(http.DefaultTransport, WithParallelRequests(false)),
	}
	
	serialStart := time.Now()
	for i := 0; i < numRequests; i++ {
		req, _ := http.NewRequest("GET", server.URL+"/repos/test/test", nil)
		resp, err := serialClient.Do(req)
		if err != nil {
			t.Fatalf("Serial request failed: %v", err)
		}
		resp.Body.Close()
	}
	serialDuration := time.Since(serialStart)

	// Test parallel requests
	parallelClient := &http.Client{
		Transport: NewRateLimitTransport(http.DefaultTransport, WithParallelRequests(true)),
	}
	
	parallelStart := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			req, _ := http.NewRequest("GET", server.URL+"/repos/test/test", nil)
			resp, err := parallelClient.Do(req)
			if err != nil {
				t.Errorf("Parallel request failed: %v", err)
			}
			if resp != nil {
				resp.Body.Close()
			}
		}()
	}
	wg.Wait()
	parallelDuration := time.Since(parallelStart)

	// Calculate and verify speedup
	speedup := float64(serialDuration) / float64(parallelDuration)
	
	t.Logf("Serial: %v for %d requests", serialDuration, numRequests)
	t.Logf("Parallel: %v for %d requests", parallelDuration, numRequests)
	t.Logf("Speedup: %.2fx", speedup)

	// With 100ms latency and 10 requests:
	// Serial should take ~1000ms (10 * 100ms)
	// Parallel should take ~100ms (all concurrent)
	// Expect at least 5x speedup
	
	if speedup < 5.0 {
		t.Errorf("Expected at least 5x speedup with parallel requests, got %.2fx", speedup)
	}
	
	// Verify serial took approximately the expected time (with some tolerance)
	expectedSerial := time.Duration(numRequests) * 100 * time.Millisecond
	if serialDuration < expectedSerial*8/10 || serialDuration > expectedSerial*12/10 {
		t.Errorf("Serial duration %v outside expected range around %v", serialDuration, expectedSerial)
	}
	
	// Verify parallel took approximately 100ms (with some tolerance)
	expectedParallel := 100 * time.Millisecond
	if parallelDuration < expectedParallel*5/10 || parallelDuration > expectedParallel*20/10 {
		t.Errorf("Parallel duration %v outside expected range around %v", parallelDuration, expectedParallel)
	}
}