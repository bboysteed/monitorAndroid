package models


type Message struct {
	*Cpu `json:"cpuInfo"`
	*Ram `json:"ramInfo"`
}