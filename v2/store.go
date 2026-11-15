package di

import (
	"errors"
	"slices"
	"sync"
)

// scopeData хранит регистрации и кэш scoped-экземпляров одного уровня:
// корневого контейнера или конкретного скоупа.
type scopeData struct {
	mu          sync.Mutex
	regs        sync.Map // any(typeKey[T]) -> []any (*registration[T])
	order       []any    // живые регистрации в порядке регистрации
	scoped      map[any]any
	scopedOrder []any
}

func newScopeData() *scopeData {
	return &scopeData{scoped: make(map[any]any)}
}

// store сохраняет регистрацию по ключу. При many=true запись добавляется к
// уже существующим, иначе список заменяется.
func (d *scopeData) store(key any, entry any, many bool) {
	d.mu.Lock()
	defer d.mu.Unlock()

	var existing []any
	if value, ok := d.regs.Load(key); ok {
		existing, _ = value.([]any)
	}

	if many {
		next := make([]any, 0, len(existing)+1)
		next = append(next, existing...)
		next = append(next, entry)

		d.regs.Store(key, next)
		d.order = append(d.order, entry)

		return
	}

	d.dropFromOrder(existing)
	d.regs.Store(key, []any{entry})
	d.order = append(d.order, entry)
}

// load возвращает первую (основную) регистрацию по ключу.
func (d *scopeData) load(key any) (any, bool) {
	value, ok := d.regs.Load(key)
	if !ok {
		return nil, false
	}

	list, ok := value.([]any)
	if !ok || len(list) == 0 {
		return nil, false
	}

	return list[0], true
}

func (d *scopeData) remove(key any) {
	d.mu.Lock()
	defer d.mu.Unlock()

	value, ok := d.regs.LoadAndDelete(key)
	if !ok {
		return
	}

	list, _ := value.([]any)
	d.dropFromOrder(list)
}

// getOrCreateScoped возвращает кэш scoped-экземпляра по ключу, создавая его
// при первом обращении.
func (d *scopeData) getOrCreateScoped(key any, build func() any) any {
	d.mu.Lock()
	defer d.mu.Unlock()

	if value, ok := d.scoped[key]; ok {
		return value
	}

	value := build()
	d.scoped[key] = value
	d.scopedOrder = append(d.scopedOrder, value)

	return value
}

func (d *scopeData) snapshotOrder() []any {
	d.mu.Lock()
	defer d.mu.Unlock()

	return append([]any(nil), d.order...)
}

func (d *scopeData) snapshotScopedOrder() []any {
	d.mu.Lock()
	defer d.mu.Unlock()

	return append([]any(nil), d.scopedOrder...)
}

func (d *scopeData) dropFromOrder(entries []any) {
	if len(entries) == 0 {
		return
	}

	removed := make(map[any]struct{}, len(entries))
	for _, entry := range entries {
		removed[entry] = struct{}{}
	}

	kept := d.order[:0]
	for _, entry := range d.order {
		if _, ok := removed[entry]; !ok {
			kept = append(kept, entry)
		}
	}

	d.order = kept
}

// disposeAll закрывает экземпляры в обратном порядке регистрации (LIFO).
func disposeAll(order []any) error {
	var errs []error

	for _, o := range slices.Backward(order) {
		item, ok := o.(disposable)
		if !ok {
			continue
		}

		if err := item.dispose(); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}
