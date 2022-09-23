package repositories

import (
	"dnovel/models/novel"
	"sync"
)

type Query func(source novel.BookSource) bool

type BookSourceRepository interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)
	Select(query Query) (movie novel.BookSource, found bool)
	SelectMany(query Query, limit int) (results []novel.BookSource)
}

func NewBookSourceRepository(source map[int64]novel.BookSource) BookSourceRepository {
	return &bookSourceRepository{
		source: source,
	}
}

type bookSourceRepository struct {
	source map[int64]novel.BookSource
	mu     sync.RWMutex
}

const (
	// ReadOnlyMode will RLock(read) the data .
	ReadOnlyMode = iota
	// ReadWriteMode will Lock(read/write) the data.
	ReadWriteMode
)

func (r *bookSourceRepository) Exec(query Query, action Query, actionLimit int, mode int) (ok bool) {
	loops := 0

	if mode == ReadOnlyMode {
		r.mu.RLock()
		defer r.mu.RUnlock()
	} else {
		r.mu.Lock()
		defer r.mu.Unlock()
	}

	for _, movie := range r.source {
		ok = query(movie)
		if ok {
			if action(movie) {
				loops++
				if actionLimit >= loops {
					break // break
				}
			}
		}
	}

	return
}

func (r *bookSourceRepository) Select(query Query) (movie novel.BookSource, found bool) {
	found = r.Exec(query, func(m novel.BookSource) bool {
		movie = m
		return true
	}, 1, ReadOnlyMode)
	//设置一个空的datamodels.Movie，如果根本找不到的话。
	if !found {
		movie = novel.BookSource{}
	}
	return
}

func (r *bookSourceRepository) SelectMany(query Query, limit int) (results []novel.BookSource) {
	r.Exec(query, func(m novel.BookSource) bool {
		results = append(results, m)
		return true
	}, limit, ReadOnlyMode)

	return
}
