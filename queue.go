package goext

type IQueue[T any] interface {
	Push(data T)
	Close()
	OnError(handler func(err error))
}

type QueueImpl[T any] struct {
	channel      chan T
	errorHandler func(err error)
}

func (queue *QueueImpl[T]) Push(data T) {
	queue.channel <- data
}

func (queue *QueueImpl[T]) Close() {
	close(queue.channel)
}

func (queue *QueueImpl[T]) OnError(handler func(err error)) {
	queue.errorHandler = handler
}

// Queue processes data sequentially by the given `handler` function and prevents concurrency
// conflicts, it returns a queue instance that we can push data into.
//
// `bufferSize` is the maximum capacity of the underlying channel, once reached, the push
// operation will block until there is new space available. Bu default, this option is not set and
// use a non-buffered channel instead.
func Queue[T any](handler func(data T), bufferSize int) IQueue[T] {
	queue := &QueueImpl[T]{channel: make(chan T, bufferSize)}

	go func() {
		_, err := Try(func() int {
			for data := range queue.channel {
				_, err := Try(func() int {
					handler(data)
					return 0
				})

				if err != nil && queue.errorHandler != nil {
					queue.errorHandler(err)
				}
			}

			return 0
		})

		if err != nil && queue.errorHandler != nil {
			queue.errorHandler(err)
		}
	}()

	return queue
}
