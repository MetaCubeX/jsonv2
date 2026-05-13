package maps

import "github.com/go-json-experiment/json/internal/go120/iter"

func All[Map ~map[K]V, K comparable, V any](m Map) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

func Keys[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range m {
			if !yield(k) {
				return
			}
		}
	}
}

func Values[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range m {
			if !yield(v) {
				return
			}
		}
	}
}

func Collect[K comparable, V any](seq iter.Seq2[K, V]) map[K]V {
	m := make(map[K]V)
	seq(func(k K, v V) bool {
		m[k] = v
		return true
	})
	return m
}

func Clone[Map ~map[K]V, K comparable, V any](m Map) Map {
	if m == nil {
		return nil
	}
	clone := make(Map, len(m))
	Copy(clone, m)
	return clone
}

func Copy[Dst ~map[K]V, Src ~map[K]V, K comparable, V any](dst Dst, src Src) {
	for k, v := range src {
		dst[k] = v
	}
}

func DeleteFunc[Map ~map[K]V, K comparable, V any](m Map, del func(K, V) bool) {
	for k, v := range m {
		if del(k, v) {
			delete(m, k)
		}
	}
}

func Equal[Map1 ~map[K]V, Map2 ~map[K]V, K, V comparable](m1 Map1, m2 Map2) bool {
	return EqualFunc(m1, m2, func(v1, v2 V) bool { return v1 == v2 })
}

func EqualFunc[Map1 ~map[K]V1, Map2 ~map[K]V2, K comparable, V1, V2 any](m1 Map1, m2 Map2, eq func(V1, V2) bool) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || !eq(v1, v2) {
			return false
		}
	}
	return true
}
