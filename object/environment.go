package object

// Environment Store of variables and there associated objects
type Environment struct {
	store map[string]Object
}

// NewEnvironment Creates an empty environment and retruns a reference to it
func NewEnvironment() *Environment {
	return &Environment{
		store: make(map[string]Object),
	}
}

// Get Tries to get an object from the environment, returns an Error if unsuccessful
func (env *Environment) Get(name string) (Object, *Error) {
	obj, ok := env.store[name]
	if !ok {
		return nil, NewErrorf("Identifier not found: %s", name)
	}
	return obj, nil
}

// Set Adds a object to the environment with a given name as a key
func (env *Environment) Set(name string, obj Object) Object {
	env.store[name] = obj
	return obj
}
