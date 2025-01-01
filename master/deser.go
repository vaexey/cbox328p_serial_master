package master

import (
	"encoding/json"
	"os"
)

func (m *Master) Serialize() ([]byte, error) {
	data, err := json.MarshalIndent(*m, "", "    ")

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *Master) Save(path string) error {
	var err error

	bytes, err := m.Serialize()

	if err != nil {
		return err
	}

	err = os.WriteFile(path, bytes, 0644)

	if err != nil {
		return err
	}

	return nil
}

func (m *Master) Deserialize(str []byte) error {
	err := json.Unmarshal(str, m)

	if err != nil {
		return err
	}

	return nil
}

func (m *Master) Load(path string) error {
	var err error

	bytes, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	err = m.Deserialize(bytes)

	if err != nil {
		return err
	}

	return nil
}
