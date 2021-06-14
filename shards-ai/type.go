package shardsai

type Template struct {
	Path     string
	Indices  []string
	Priority int
}

type IndexTemplate struct {
	IndexPatterns []string `json:"index_patterns"`
	Template      struct {
		Settings struct {
			Index struct {
				Codec   string `json:"codec"`
				Routing struct {
					Allocation struct {
						Require struct {
							Type string `json:"type"`
						} `json:"require"`
					} `json:"allocation"`
				} `json:"routing"`
				NumberOfShards   string `json:"number_of_shards"`
				NumberOfReplicas string `json:"number_of_replicas"`
			} `json:"index"`
		} `json:"settings"`
	} `json:"template"`
	Priority int `json:"priority"`
}

type CatIndex struct {
	Health       string `json:"health"`
	Status       string `json:"status"`
	Index        string `json:"index"`
	UUID         string `json:"uuid"`
	Pri          string `json:"pri"`
	Rep          string `json:"rep"`
	DocsCount    string `json:"docs.count"`
	DocsDeleted  string `json:"docs.deleted"`
	StoreSize    string `json:"store.size"`
	PriStoreSize string `json:"pri.store.size"`
}