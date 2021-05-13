package zoo

import (
	"sync"

	"github.com/teratron/bingo/pkg"
	"github.com/teratron/bingo/pkg/utils"
)

// Neuroner
type Neuroner interface {
}

// Synapser
type Synapser interface {
}

// Name of the neural network architecture.
const Name = "configurator"

// Declare conformity with NeuralNetwork interface.
var _ pkg.NeuralNetwork = (*NN)(nil)

type NN struct {
	pkg.NeuralNetwork

	isInit bool
	config utils.Filer
	mutex  sync.Mutex
}

type neuron struct {
	Synapser

	value    pkg.FloatType // Neuron value
	axon     []*axon       // All incoming axons
	specific Neuroner      // Specific option of neuron: miss (error) or other
}

type axon struct {
	weight  pkg.FloatType
	synapse map[string]Synapser
}

// New return Perceptron neural network.
func New() *NN {
	return &NN{
		/*Name: Name,
		Activation: params.ModeSIGMOID,
		Loss:       params.ModeMSE,
		Limit:      .01,
		Rate:       pkg.FloatType(params.DefaultRate),*/
	}
}
