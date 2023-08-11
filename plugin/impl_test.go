package plugin

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegistries(t *testing.T) {
	p := &Plugin{}
	p.settings.Login.RegistriesRaw = `[{"username": "docker_user", "password": "docker_password"}]`

	assert.NoError(t, p.Validate())

	fmt.Println(p.settings.Login.Registries[0].Password)

	if assert.Len(t, p.settings.Login.Registries, 1) {
		assert.EqualValues(t, "docker_user", p.settings.Login.Registries[0].Username)
		assert.EqualValues(t, "docker_password", p.settings.Login.Registries[0].Password)
		assert.EqualValues(t, DefaultRegistry, p.settings.Login.Registries[0].Registry)
	}
}
