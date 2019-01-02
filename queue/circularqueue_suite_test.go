package queue_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCircularqueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Circularqueue Suite")
}

