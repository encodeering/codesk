package config

func Combine (configurations ...*Config) (config *Config) {
    config = DefaultConfig ()

    for _, current := range configurations {
        if                        current.Box.User.Name != "" {
           config.Box.User.Name = current.Box.User.Name
        }

        if                                current.Box.Distribution.Name != "" {
           config.Box.Distribution.Name = current.Box.Distribution.Name
        }

        env        :=  &config.Command.Environment
        envcurrent := &current.Command.Environment

               resolution := envcurrent.Resolution
        switch resolution {
            case "first":  env.Var = WslVarCombine (env.Var, envcurrent.Var)
            case "parent": env.Var =                env.Var
            case "self":   env.Var =                         envcurrent.Var
            case "last":   env.Var = WslVarCombine (         envcurrent.Var, env.Var)
        }

        config.Command.Environment.Resolution = resolution
    }

    return
}

func WslVarCombine (priority, economy []WslVar) (result []WslVar) {
	result = make ([]WslVar, 0)

	putif := func (to map[string]WslVar, from map[string]WslVar) {
		for k, _ := range from {
			if _, ok := to[k]; ! ok {
				to[k] = from[k]
			}
		}
	}

	       all := make (map[string]WslVar, len (priority))
	putif (all, WslVarAsMap (priority))
	putif (all, WslVarAsMap (economy))

	for _, v := range all {
		result = append (result, v)
	}

	return
}

func WslVarAsMap (values []WslVar) (result map[string]WslVar) {
	result = make (map[string]WslVar, len (values))

	for _, item := range values {
		result[item.Key] = item
	}

	return
}
