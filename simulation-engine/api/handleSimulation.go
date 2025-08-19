package api

import (
	"encoding/json"
	"log"
	"net/http"
	"simengine/simulator"
	"simengine/types"
)

func handleSimulation(w http.ResponseWriter, r *http.Request) {

	data := []byte(`{
		"nodes": [
			{
				"id": "entry",
				"type": "Entry",
				"resources": {
					"cpu": 0,
					"memoryMB": 0,
					"replicas": 1
				}
			},
			{
				"id": "lb-1",
				"type": "LoadBalancer",
				"resources": {
					"cpu": 0.2,
					"memoryMB": 256,
					"replicas": 1
				}
			},
			{
				"id": "api-1",
				"type": "API",
				"resources": {
					"cpu": 1.0,
					"memoryMB": 1024,
					"replicas": 2
				}
			},
			{
				"id": "queue-1",
				"type": "Queue",
				"resources": {
					"cpu": 0.5,
					"memoryMB": 512,
					"replicas": 1
				}
			},
			{
				"id": "worker-1",
				"type": "Worker",
				"resources": {
					"cpu": 1.5,
					"memoryMB": 2048,
					"replicas": 3
				}
			},
			{
				"id": "db-1",
				"type": "Database",
				"resources": {
					"cpu": 2.0,
					"memoryMB": 4096,
					"replicas": 1
				}
			},
			{
				"id": "exit",
				"type": "Exit",
				"resources": {
					"cpu": 0,
					"memoryMB": 0,
					"replicas": 1
				}
			}
		],
		"connections": [
			{ "from": "entry", "to": "lb-1" },
			{ "from": "lb-1", "to": "api-1" },
			{ "from": "api-1", "to": "queue-1" },
			{ "from": "queue-1", "to": "worker-1" },
			{ "from": "worker-1", "to": "db-1" },
			{ "from": "db-1", "to": "exit" }
		]
	}`)

	var design types.Design
	if err := json.Unmarshal(data, &design); err != nil {
		panic(err)
	}

	log.Print(simulator.BuildGraph(design))
	log.Print(simulator.FindIsolatedNodes(design))

	RespondWithJSON(w, 200, map[string]string{"msg":"done"})

}