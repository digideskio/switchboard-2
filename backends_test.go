package switchboard_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf-experimental/switchboard"
	"github.com/pivotal-golang/lager"
)

var _ = Describe("Backends", func() {
	var (
		backends          switchboard.Backends
		backend_ips       []string
		backend_ports     []uint
		healthcheck_ports []uint
		logger            lager.Logger
	)

	var backendChanToSlice = func(c <-chan switchboard.Backend) []switchboard.Backend {
		var result []switchboard.Backend
		for b := range c {
			result = append(result, b)
		}
		return result
	}

	BeforeEach(func() {
		backend_ips = []string{"localhost", "localhost", "localhost"}
		backend_ports = []uint{50000, 50001, 50002}
		healthcheck_ports = []uint{60000, 60001, 60002}
		logger = lager.NewLogger("Backends test")
		backends = switchboard.NewBackends(backend_ips, backend_ports, healthcheck_ports, logger)
	})

	Describe("Concurrent operations", func() {
		It("do not result in a race", func() {
			readySetGo := make(chan interface{})

			doneChans := []chan interface{}{
				make(chan interface{}),
				make(chan interface{}),
				make(chan interface{}),
				make(chan interface{}),
				make(chan interface{}),
			}

			go func() {
				<-readySetGo
				backends.All()
				close(doneChans[0])
			}()

			go func() {
				<-readySetGo
				backends.Active()
				close(doneChans[1])
			}()

			go func() {
				<-readySetGo
				backends.SetHealthy(nil)
				close(doneChans[2])
			}()

			go func() {
				<-readySetGo
				backends.SetUnhealthy(nil)
				close(doneChans[3])
			}()

			go func() {
				<-readySetGo
				backends.Healthy()
				close(doneChans[4])
			}()

			close(readySetGo)

			for _, done := range doneChans {
				<-done
			}
		})
	})

	Describe("All", func() {
		It("allows iterating over all the backends", func() {
			backendsSeen := []string{}
			for backend := range backends.All() {
				backendsSeen = append(backendsSeen, backend.HealthcheckUrl())
			}

			Expect(backendsSeen).To(ContainElement("http://localhost:60000"))
			Expect(backendsSeen).To(ContainElement("http://localhost:60001"))
			Expect(backendsSeen).To(ContainElement("http://localhost:60002"))
		})
	})

	Describe("SetHealthy", func() {
		var unhealthy switchboard.Backend

		BeforeEach(func() {
			unhealthy = backendChanToSlice(backends.Healthy())[0]
			backends.SetUnhealthy(unhealthy)
		})

		It("sets the backend to be healthy", func() {
			Expect(len(backendChanToSlice(backends.Healthy()))).To(Equal(2))
			backends.SetHealthy(unhealthy)
			Expect(len(backendChanToSlice(backends.Healthy()))).To(Equal(3))
		})

		Context("when all backends are unhealthy and there is no active backend", func() {
			BeforeEach(func() {
				healthy := backendChanToSlice(backends.Healthy())
				for _, b := range healthy {
					backends.SetUnhealthy(b)
				}
			})

			It("sets the newly healthy backend as the new active backend", func() {
				Expect(backends.Active()).To(BeNil())
				backend := backends.Any()
				backends.SetHealthy(backend)
				Expect(backends.Active()).To(Equal(backend))
			})
		})
	})

	Describe("SetUnhealthy", func() {
		var healthy switchboard.Backend

		BeforeEach(func() {
			healthy = backendChanToSlice(backends.Healthy())[0]
		})

		It("sets the backend to be unhealthy", func() {
			Expect(len(backendChanToSlice(backends.Healthy()))).To(Equal(3))
			backends.SetUnhealthy(healthy)
			Expect(len(backendChanToSlice(backends.Healthy()))).To(Equal(2))
		})

		Context("when there is at least one healthy backend", func() {
			It("sets another healthy backend as the new active backend", func() {
				numHealthy := len(backendChanToSlice(backends.Healthy()))
				for _ = range backends.Healthy() {
					previousActive := backends.Active()
					backends.SetUnhealthy(previousActive)
					nextActive := backends.Active()
					Expect(nextActive).ToNot(Equal(previousActive))

					numHealthy--
					if numHealthy > 0 { // more healthy backends
						Expect(backends.Active()).ToNot(BeNil())
					} else { // no more healthy backends -> no active backend
						Expect(backends.Active()).To(BeNil())
					}
				}
			})
		})
	})

	Describe("Healthy", func() {
		It("sets the backend to be healthy", func() {
			healthy := backendChanToSlice(backends.Healthy())
			numHealthy := 3
			Expect(len(healthy)).To(Equal(numHealthy))

			for _, b := range healthy {
				backends.SetUnhealthy(b)
				numHealthy--
				healthy = backendChanToSlice(backends.Healthy())
				Expect(len(healthy)).To(Equal(numHealthy))
			}
		})
	})
})
