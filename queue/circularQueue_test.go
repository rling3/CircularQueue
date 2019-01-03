package queue_test

import (
	"circularqueue/queue"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CircularQueue", func() {

	var (
		//data = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n"}
		q queue.Queue
	)

	Describe("Testing basic push and pop", func() {

		BeforeEach(func() {
			q = queue.New(nil)
		})

		//Context("Test popping from empty array", func() {
		//	It("Should return correct error", func() {
		//		val, err := q.Pop()
		//		Expect(val).To(BeNil())
		//		Expect(err).To(Equal(queue.E_EMPTY))
		//	})
		//})
		//Context("Push and pop single character", func() {
		//	It("Should pop same element back out", func() {
		//		q.Push(data[0])
		//		val, err := q.Pop()
		//		Expect(err).To(BeNil())
		//		Expect(val).To(Equal(data[0]))
		//	})
		//})
		//Context("Push and pop multiple characters", func() {
		//	It("Should pop 3 elements in queue order", func() {
		//		for i := 0; i < 3; i++ {
		//			q.Push(data[i])
		//		}
		//		Expect(q.Length()).To(Equal(8))
		//		Expect(q.Pop()).To(Equal(data[0]))
		//		Expect(q.Pop()).To(Equal(data[1]))
		//		Expect(q.Pop()).To(Equal(data[2]))
		//	})
		//	It("Should grow and shrink queue", func() {
		//		for i := 0; i < 4; i++ {
		//			q.Push(data[i])
		//		}
		//		Expect(q.Size()).To(Equal(4))
		//		Expect(q.Length()).To(Equal(8))
		//		Expect(q.Pop()).To(Equal(data[0]))
		//		Expect(q.Pop()).To(Equal(data[1]))
		//		Expect(q.Pop()).To(Equal(data[2]))
		//		Expect(q.Pop()).To(Equal(data[3]))
		//		Expect(q.Length()).To(Equal(4))
		//		Expect(q.Size()).To(Equal(0))
		//	})
		//})
		//Context("Push and pop a lot of characters", func() {
		//	It("Should work", func() {
		//		for _, val := range data {
		//			q.Push(val)
		//			verifySize(q)
		//		}
		//		Expect(q.Size()).To(Equal(len(data)))
		//		Expect(q.Length()).To(Equal(32))
		//		for i := 0; i < len(data); i++ {
		//			Expect(q.Pop()).To(Equal(data[i]))
		//			verifySize(q)
		//		}
		//	})
		//})
		Context("Push and pop a lot of characters", func() {
			It("Should work", func() {
				var data []int
				for i := 0; i < 10000; i++ {
					data = append(data, i)
				}
				var popped []int
				for _, i := range data {
					if i%100 == 0 {
						fmt.Print(i)
					}
					q.Push(i)
					if i%15 >= 10 {
						v, err := q.Pop()
						Expect(err).ToNot(HaveOccurred())
						popped = append(popped, v.(int))
						//fmt.Printf("%v\n", v)
					}
				}
				//fmt.Printf("\nLen: %v\n", q.Length())
				for q.Size() > 0 {
					v, err := q.Pop()
					Expect(err).ToNot(HaveOccurred())
					popped = append(popped, v.(int))
					//fmt.Printf("%v\n", v)
				}
				//fmt.Printf("\nLen: %v\n", q.Length())
				Expect(popped).To(Equal(data))
			})
		})
	})

})

func verifySize(q queue.Queue) {
	switch {
	case q.Size() > 8:
		Expect(q.Length()).To(Equal(32))
	case q.Size() > 4:
		Expect(q.Length()).To(Equal(16))
	case q.Size() > 2:
		Expect(q.Length()).To(Equal(8))
	default:
		Expect(q.Length()).To(Equal(queue.DEFAULT_SIZE))
	}
}
