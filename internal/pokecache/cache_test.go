package pokecache

import (
    "testing"
    "time"
)

func TestAddGet(t *testing.T) {
    c := NewCache(5 * time.Second)

    c.Add("key1", []byte("value1"))
    c.Add("key2", []byte("value2"))

    got, ok := c.Get("key1")
    if !ok {
        t.Fatalf("expected key1 to exist")
    }
    if string(got) != "value1" {
        t.Fatalf("key1: expected %q, got %q", "value1", string(got))
    }

    got, ok = c.Get("key2")
    if !ok {
        t.Fatalf("expected key2 to exist")
    }
    if string(got) != "value2" {
        t.Fatalf("key2: expected %q, got %q", "value2", string(got))
    }
}

func TestGetMissing(t *testing.T) {
    c := NewCache(5 * time.Second)

    got, ok := c.Get("nope")
    if ok {
        t.Fatalf("expected missing key to return ok=false")
    }
    if got != nil {
        t.Fatalf("expected nil for missing key, got %v", got)
    }
}

func TestAddReplaces(t *testing.T) {
    c := NewCache(5 * time.Second)

    c.Add("key", []byte("first"))
    c.Add("key", []byte("second"))

    got, ok := c.Get("key")
    if !ok {
        t.Fatalf("expected key to exist")
    }
    if string(got) != "second" {
        t.Fatalf("expected replaced value %q, got %q", "second", string(got))
    }
}

func TestKeysDontOverwriteEachOther(t *testing.T) {
    c := NewCache(5 * time.Second)

    c.Add("a", []byte("A"))
    c.Add("b", []byte("B"))
    c.Add("c", []byte("C"))

    checks := map[string]string{"a": "A", "b": "B", "c": "C"}
    for key, want := range checks {
        got, ok := c.Get(key)
        if !ok {
            t.Fatalf("expected %q to exist", key)
        }
        if string(got) != want {
            t.Fatalf("%q: expected %q, got %q", key, want, string(got))
        }
    }
}

func TestFreshEntryStays(t *testing.T) {
    c := NewCache(100 * time.Millisecond)

    c.Add("fresh", []byte("value"))

    // Wait less than the interval — the entry should still be there.
    time.Sleep(50 * time.Millisecond)

    got, ok := c.Get("fresh")
    if !ok {
        t.Fatalf("expected fresh entry to still exist")
    }
    if string(got) != "value" {
        t.Fatalf("expected %q, got %q", "value", string(got))
    }
}

func TestOldEntryIsReaped(t *testing.T) {
    c := NewCache(50 * time.Millisecond)

    c.Add("old", []byte("value"))

    // Wait several intervals so the reaper fires at least once
    // and the entry is definitely older than the interval.
    time.Sleep(200 * time.Millisecond)

    _, ok := c.Get("old")
    if ok {
        t.Fatalf("expected old entry to be reaped")
    }
}