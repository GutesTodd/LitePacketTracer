// Packages repositories
package repositories

import "litepackettracer/internal/domain/topology"

type LabRepository interface {
	Get(id topology.LabID) (*topology.Lab, error)
	Save(lab *topology.Lab) error
}
