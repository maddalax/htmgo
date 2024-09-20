package h

import "fmt"

type AttributeMap map[string]any

func (m *AttributeMap) ToMap() map[string]string {
	result := make(map[string]string)
	for k, v := range *m {
		switch v.(type) {
		case AttributeMap:
			m2 := v.(*AttributeMap).ToMap()
			for _, a := range m2 {
				result[k] = a
			}
		case string:
			result[k] = v.(string)
		default:
			result[k] = fmt.Sprintf("%v", v)
		}
	}
	return result
}
