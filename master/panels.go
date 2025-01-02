package master

import (
	"cbox328p/serial_master/panel"
	"fmt"
)

func (m *Master) SendDirtyPanels() error {
	var errCount int = 0

	for pi := range m.Panels {
		p := &m.Panels[pi]

		if p.Dirty {

			err := m.SendDirect(p.Id, p.RawColor)

			if err != nil {
				errCount++

				//continue
			}

			p.Dirty = false
		}
	}

	if errCount != 0 {
		return fmt.Errorf("could not update %d panels", errCount)
	}

	return nil
}

func (m *Master) GenerateIndex() {
	if m.panelIdMap != nil && m.panelGroupIdMap != nil {
		return
	}

	m.RegenerateIndex()
}

func (m *Master) RegenerateIndex() {
	m.panelIdMap = make(map[string]*panel.Panel)
	m.panelGroupIdMap = make(map[string][]*panel.Panel)

	for pi := range m.Panels {
		p := &m.Panels[pi]

		m.panelIdMap[p.Id] = p

		for _, gid := range append(p.GroupIds, p.Id) {
			if m.panelGroupIdMap[gid] == nil {
				m.panelGroupIdMap[gid] = make([]*panel.Panel, 0)
			}

			m.panelGroupIdMap[gid] = append(m.panelGroupIdMap[gid], p)
		}
	}
}

func (m *Master) FindPanel(id string) *panel.Panel {
	m.GenerateIndex()

	return m.panelIdMap[id]
}

func (m *Master) FindPanels(id string) []*panel.Panel {
	m.GenerateIndex()

	return m.panelGroupIdMap[id]
}
