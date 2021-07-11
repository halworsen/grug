package grug

func init() {
	actionStore := make(map[string]interface{})

	AllActions = append(AllActions, []Action{
		{
			// Store arg 1 in a field named arg 0
			// Passes the stored value back through as return value
			Name: "Store",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				field := atostr(args[0])
				actionStore[field] = args[1]
				return args[1], nil
			},
		},
		{
			// Fetch and return the value stored in the given field
			// Arg 1 may be supplied as a default value in case the field doesn't already exist
			Name: "Load",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				field := atostr(args[0])
				var defaultValue interface{}
				defaultValue = nil
				if len(args) > 1 {
					defaultValue = args[1]
				}

				val, exists := actionStore[field]
				if !exists {
					val = defaultValue
				}
				return val, nil
			},
		},
	}...)
}
