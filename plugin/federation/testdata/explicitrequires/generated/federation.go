// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/99designs/gqlgen/plugin/federation/fedruntime"
)

var (
	ErrUnknownType  = errors.New("unknown type")
	ErrTypeNotFound = errors.New("type not found")
)

func (ec *executionContext) __resolve__service(ctx context.Context) (fedruntime.Service, error) {
	if ec.DisableIntrospection {
		return fedruntime.Service{}, errors.New("federated introspection disabled")
	}

	var sdl []string

	for _, src := range sources {
		if src.BuiltIn {
			continue
		}
		sdl = append(sdl, src.Input)
	}

	return fedruntime.Service{
		SDL: strings.Join(sdl, "\n"),
	}, nil
}

func (ec *executionContext) __resolve_entities(ctx context.Context, representations []map[string]interface{}) []fedruntime.Entity {
	list := make([]fedruntime.Entity, len(representations))

	repsMap := map[string]struct {
		i []int
		r []map[string]interface{}
	}{}

	// We group entities by typename so that we can parallelize their resolution.
	// This is particularly helpful when there are entity groups in multi mode.
	buildRepresentationGroups := func(reps []map[string]interface{}) {
		for i, rep := range reps {
			typeName, ok := rep["__typename"].(string)
			if !ok {
				// If there is no __typename, we just skip the representation;
				// we just won't be resolving these unknown types.
				ec.Error(ctx, errors.New("__typename must be an existing string"))
				continue
			}

			_r := repsMap[typeName]
			_r.i = append(_r.i, i)
			_r.r = append(_r.r, rep)
			repsMap[typeName] = _r
		}
	}

	isMulti := func(typeName string) bool {
		switch typeName {
		case "MultiHello":
			return true
		case "MultiHelloMultipleRequires":
			return true
		case "MultiHelloRequires":
			return true
		case "MultiHelloWithError":
			return true
		case "MultiPlanetRequiresNested":
			return true
		default:
			return false
		}
	}

	resolveEntity := func(ctx context.Context, typeName string, rep map[string]interface{}, idx []int, i int) (err error) {
		// we need to do our own panic handling, because we may be called in a
		// goroutine, where the usual panic handling can't catch us
		defer func() {
			if r := recover(); r != nil {
				err = ec.Recover(ctx, r)
			}
		}()

		switch typeName {
		case "Hello":
			resolverName, err := entityResolverNameForHello(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Hello": %w`, err)
			}
			switch resolverName {

			case "findHelloByName":
				id0, err := ec.unmarshalNString2string(ctx, rep["name"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findHelloByName(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindHelloByName(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Hello": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "HelloMultiSingleKeys":
			resolverName, err := entityResolverNameForHelloMultiSingleKeys(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "HelloMultiSingleKeys": %w`, err)
			}
			switch resolverName {

			case "findHelloMultiSingleKeysByKey1AndKey2":
				id0, err := ec.unmarshalNString2string(ctx, rep["key1"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findHelloMultiSingleKeysByKey1AndKey2(): %w`, err)
				}
				id1, err := ec.unmarshalNString2string(ctx, rep["key2"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 1 for findHelloMultiSingleKeysByKey1AndKey2(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindHelloMultiSingleKeysByKey1AndKey2(ctx, id0, id1)
				if err != nil {
					return fmt.Errorf(`resolving Entity "HelloMultiSingleKeys": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "HelloWithErrors":
			resolverName, err := entityResolverNameForHelloWithErrors(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "HelloWithErrors": %w`, err)
			}
			switch resolverName {

			case "findHelloWithErrorsByName":
				id0, err := ec.unmarshalNString2string(ctx, rep["name"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findHelloWithErrorsByName(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindHelloWithErrorsByName(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "HelloWithErrors": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "PlanetMultipleRequires":
			resolverName, err := entityResolverNameForPlanetMultipleRequires(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "PlanetMultipleRequires": %w`, err)
			}
			switch resolverName {

			case "findPlanetMultipleRequiresByName":
				id0, err := ec.unmarshalNString2string(ctx, rep["name"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findPlanetMultipleRequiresByName(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindPlanetMultipleRequiresByName(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "PlanetMultipleRequires": %w`, err)
				}

				err = ec.PopulatePlanetMultipleRequiresRequires(ctx, entity, rep)
				if err != nil {
					return fmt.Errorf(`populating requires for Entity "PlanetMultipleRequires": %w`, err)
				}
				list[idx[i]] = entity
				return nil
			}
		case "PlanetRequires":
			resolverName, err := entityResolverNameForPlanetRequires(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "PlanetRequires": %w`, err)
			}
			switch resolverName {

			case "findPlanetRequiresByName":
				id0, err := ec.unmarshalNString2string(ctx, rep["name"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findPlanetRequiresByName(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindPlanetRequiresByName(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "PlanetRequires": %w`, err)
				}

				err = ec.PopulatePlanetRequiresRequires(ctx, entity, rep)
				if err != nil {
					return fmt.Errorf(`populating requires for Entity "PlanetRequires": %w`, err)
				}
				list[idx[i]] = entity
				return nil
			}
		case "PlanetRequiresNested":
			resolverName, err := entityResolverNameForPlanetRequiresNested(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "PlanetRequiresNested": %w`, err)
			}
			switch resolverName {

			case "findPlanetRequiresNestedByName":
				id0, err := ec.unmarshalNString2string(ctx, rep["name"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findPlanetRequiresNestedByName(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindPlanetRequiresNestedByName(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "PlanetRequiresNested": %w`, err)
				}

				err = ec.PopulatePlanetRequiresNestedRequires(ctx, entity, rep)
				if err != nil {
					return fmt.Errorf(`populating requires for Entity "PlanetRequiresNested": %w`, err)
				}
				list[idx[i]] = entity
				return nil
			}
		case "World":
			resolverName, err := entityResolverNameForWorld(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "World": %w`, err)
			}
			switch resolverName {

			case "findWorldByHelloNameAndFoo":
				id0, err := ec.unmarshalNString2string(ctx, rep["hello"].(map[string]interface{})["name"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findWorldByHelloNameAndFoo(): %w`, err)
				}
				id1, err := ec.unmarshalNString2string(ctx, rep["foo"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 1 for findWorldByHelloNameAndFoo(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindWorldByHelloNameAndFoo(ctx, id0, id1)
				if err != nil {
					return fmt.Errorf(`resolving Entity "World": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "WorldName":
			resolverName, err := entityResolverNameForWorldName(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "WorldName": %w`, err)
			}
			switch resolverName {

			case "findWorldNameByName":
				id0, err := ec.unmarshalNString2string(ctx, rep["name"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findWorldNameByName(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindWorldNameByName(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "WorldName": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "WorldWithMultipleKeys":
			resolverName, err := entityResolverNameForWorldWithMultipleKeys(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "WorldWithMultipleKeys": %w`, err)
			}
			switch resolverName {

			case "findWorldWithMultipleKeysByHelloNameAndFoo":
				id0, err := ec.unmarshalNString2string(ctx, rep["hello"].(map[string]interface{})["name"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findWorldWithMultipleKeysByHelloNameAndFoo(): %w`, err)
				}
				id1, err := ec.unmarshalNString2string(ctx, rep["foo"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 1 for findWorldWithMultipleKeysByHelloNameAndFoo(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindWorldWithMultipleKeysByHelloNameAndFoo(ctx, id0, id1)
				if err != nil {
					return fmt.Errorf(`resolving Entity "WorldWithMultipleKeys": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			case "findWorldWithMultipleKeysByBar":
				id0, err := ec.unmarshalNInt2int(ctx, rep["bar"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findWorldWithMultipleKeysByBar(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindWorldWithMultipleKeysByBar(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "WorldWithMultipleKeys": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}

		}
		return fmt.Errorf("%w: %s", ErrUnknownType, typeName)
	}

	resolveManyEntities := func(ctx context.Context, typeName string, reps []map[string]interface{}, idx []int) (err error) {
		// we need to do our own panic handling, because we may be called in a
		// goroutine, where the usual panic handling can't catch us
		defer func() {
			if r := recover(); r != nil {
				err = ec.Recover(ctx, r)
			}
		}()

		switch typeName {

		case "MultiHello":
			resolverName, err := entityResolverNameForMultiHello(ctx, reps[0])
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "MultiHello": %w`, err)
			}
			switch resolverName {

			case "findManyMultiHelloByNames":
				_reps := make([]*MultiHelloByNamesInput, len(reps))

				for i, rep := range reps {
					id0, err := ec.unmarshalNString2string(ctx, rep["name"])
					if err != nil {
						return errors.New(fmt.Sprintf("Field %s undefined in schema.", "name"))
					}

					_reps[i] = &MultiHelloByNamesInput{
						Name: id0,
					}
				}

				entities, err := ec.resolvers.Entity().FindManyMultiHelloByNames(ctx, _reps)
				if err != nil {
					return err
				}

				for i, entity := range entities {
					list[idx[i]] = entity
				}
				return nil

			default:
				return fmt.Errorf("unknown resolver: %s", resolverName)
			}

		case "MultiHelloMultipleRequires":
			resolverName, err := entityResolverNameForMultiHelloMultipleRequires(ctx, reps[0])
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "MultiHelloMultipleRequires": %w`, err)
			}
			switch resolverName {

			case "findManyMultiHelloMultipleRequiresByNames":
				_reps := make([]*MultiHelloMultipleRequiresByNamesInput, len(reps))

				for i, rep := range reps {
					id0, err := ec.unmarshalNString2string(ctx, rep["name"])
					if err != nil {
						return errors.New(fmt.Sprintf("Field %s undefined in schema.", "name"))
					}

					_reps[i] = &MultiHelloMultipleRequiresByNamesInput{
						Name: id0,
					}
				}

				entities, err := ec.resolvers.Entity().FindManyMultiHelloMultipleRequiresByNames(ctx, _reps)
				if err != nil {
					return err
				}

				for i, entity := range entities {
					entity.Key1, err = ec.unmarshalNString2string(ctx, reps[i]["key1"])
					if err != nil {
						return err
					}
					entity.Key2, err = ec.unmarshalNString2string(ctx, reps[i]["key2"])
					if err != nil {
						return err
					}
					list[idx[i]] = entity
				}
				return nil

			default:
				return fmt.Errorf("unknown resolver: %s", resolverName)
			}

		case "MultiHelloRequires":
			resolverName, err := entityResolverNameForMultiHelloRequires(ctx, reps[0])
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "MultiHelloRequires": %w`, err)
			}
			switch resolverName {

			case "findManyMultiHelloRequiresByNames":
				_reps := make([]*MultiHelloRequiresByNamesInput, len(reps))

				for i, rep := range reps {
					id0, err := ec.unmarshalNString2string(ctx, rep["name"])
					if err != nil {
						return errors.New(fmt.Sprintf("Field %s undefined in schema.", "name"))
					}

					_reps[i] = &MultiHelloRequiresByNamesInput{
						Name: id0,
					}
				}

				entities, err := ec.resolvers.Entity().FindManyMultiHelloRequiresByNames(ctx, _reps)
				if err != nil {
					return err
				}

				for i, entity := range entities {
					entity.Key1, err = ec.unmarshalNString2string(ctx, reps[i]["key1"])
					if err != nil {
						return err
					}
					list[idx[i]] = entity
				}
				return nil

			default:
				return fmt.Errorf("unknown resolver: %s", resolverName)
			}

		case "MultiHelloWithError":
			resolverName, err := entityResolverNameForMultiHelloWithError(ctx, reps[0])
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "MultiHelloWithError": %w`, err)
			}
			switch resolverName {

			case "findManyMultiHelloWithErrorByNames":
				_reps := make([]*MultiHelloWithErrorByNamesInput, len(reps))

				for i, rep := range reps {
					id0, err := ec.unmarshalNString2string(ctx, rep["name"])
					if err != nil {
						return errors.New(fmt.Sprintf("Field %s undefined in schema.", "name"))
					}

					_reps[i] = &MultiHelloWithErrorByNamesInput{
						Name: id0,
					}
				}

				entities, err := ec.resolvers.Entity().FindManyMultiHelloWithErrorByNames(ctx, _reps)
				if err != nil {
					return err
				}

				for i, entity := range entities {
					list[idx[i]] = entity
				}
				return nil

			default:
				return fmt.Errorf("unknown resolver: %s", resolverName)
			}

		case "MultiPlanetRequiresNested":
			resolverName, err := entityResolverNameForMultiPlanetRequiresNested(ctx, reps[0])
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "MultiPlanetRequiresNested": %w`, err)
			}
			switch resolverName {

			case "findManyMultiPlanetRequiresNestedByNames":
				_reps := make([]*MultiPlanetRequiresNestedByNamesInput, len(reps))

				for i, rep := range reps {
					id0, err := ec.unmarshalNString2string(ctx, rep["name"])
					if err != nil {
						return errors.New(fmt.Sprintf("Field %s undefined in schema.", "name"))
					}

					_reps[i] = &MultiPlanetRequiresNestedByNamesInput{
						Name: id0,
					}
				}

				entities, err := ec.resolvers.Entity().FindManyMultiPlanetRequiresNestedByNames(ctx, _reps)
				if err != nil {
					return err
				}

				for i, entity := range entities {
					entity.World.Foo, err = ec.unmarshalNString2string(ctx, reps[i]["world"].(map[string]interface{})["foo"])
					if err != nil {
						return err
					}
					list[idx[i]] = entity
				}
				return nil

			default:
				return fmt.Errorf("unknown resolver: %s", resolverName)
			}

		default:
			return errors.New("unknown type: " + typeName)
		}
	}

	resolveEntityGroup := func(typeName string, reps []map[string]interface{}, idx []int) {
		if isMulti(typeName) {
			err := resolveManyEntities(ctx, typeName, reps, idx)
			if err != nil {
				ec.Error(ctx, err)
			}
		} else {
			// if there are multiple entities to resolve, parallelize (similar to
			// graphql.FieldSet.Dispatch)
			var e sync.WaitGroup
			e.Add(len(reps))
			for i, rep := range reps {
				i, rep := i, rep
				go func(i int, rep map[string]interface{}) {
					err := resolveEntity(ctx, typeName, rep, idx, i)
					if err != nil {
						ec.Error(ctx, err)
					}
					e.Done()
				}(i, rep)
			}
			e.Wait()
		}
	}
	buildRepresentationGroups(representations)

	switch len(repsMap) {
	case 0:
		return list
	case 1:
		for typeName, reps := range repsMap {
			resolveEntityGroup(typeName, reps.r, reps.i)
		}
		return list
	default:
		var g sync.WaitGroup
		g.Add(len(repsMap))
		for typeName, reps := range repsMap {
			go func(typeName string, reps []map[string]interface{}, idx []int) {
				resolveEntityGroup(typeName, reps, idx)
				g.Done()
			}(typeName, reps.r, reps.i)
		}
		g.Wait()
		return list
	}
}

func entityResolverNameForHello(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["name"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findHelloByName", nil
	}
	return "", fmt.Errorf("%w for Hello", ErrTypeNotFound)
}

func entityResolverNameForHelloMultiSingleKeys(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["key1"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		m = rep
		val, ok = m["key2"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findHelloMultiSingleKeysByKey1AndKey2", nil
	}
	return "", fmt.Errorf("%w for HelloMultiSingleKeys", ErrTypeNotFound)
}

func entityResolverNameForHelloWithErrors(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["name"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findHelloWithErrorsByName", nil
	}
	return "", fmt.Errorf("%w for HelloWithErrors", ErrTypeNotFound)
}

func entityResolverNameForMultiHello(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["name"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findManyMultiHelloByNames", nil
	}
	return "", fmt.Errorf("%w for MultiHello", ErrTypeNotFound)
}

func entityResolverNameForMultiHelloMultipleRequires(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["name"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findManyMultiHelloMultipleRequiresByNames", nil
	}
	return "", fmt.Errorf("%w for MultiHelloMultipleRequires", ErrTypeNotFound)
}

func entityResolverNameForMultiHelloRequires(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["name"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findManyMultiHelloRequiresByNames", nil
	}
	return "", fmt.Errorf("%w for MultiHelloRequires", ErrTypeNotFound)
}

func entityResolverNameForMultiHelloWithError(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["name"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findManyMultiHelloWithErrorByNames", nil
	}
	return "", fmt.Errorf("%w for MultiHelloWithError", ErrTypeNotFound)
}

func entityResolverNameForMultiPlanetRequiresNested(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["name"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findManyMultiPlanetRequiresNestedByNames", nil
	}
	return "", fmt.Errorf("%w for MultiPlanetRequiresNested", ErrTypeNotFound)
}

func entityResolverNameForPlanetMultipleRequires(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["name"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findPlanetMultipleRequiresByName", nil
	}
	return "", fmt.Errorf("%w for PlanetMultipleRequires", ErrTypeNotFound)
}

func entityResolverNameForPlanetRequires(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["name"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findPlanetRequiresByName", nil
	}
	return "", fmt.Errorf("%w for PlanetRequires", ErrTypeNotFound)
}

func entityResolverNameForPlanetRequiresNested(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["name"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findPlanetRequiresNestedByName", nil
	}
	return "", fmt.Errorf("%w for PlanetRequiresNested", ErrTypeNotFound)
}

func entityResolverNameForWorld(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["hello"]
		if !ok {
			break
		}
		if m, ok = val.(map[string]interface{}); !ok {
			break
		}
		val, ok = m["name"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		m = rep
		val, ok = m["foo"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findWorldByHelloNameAndFoo", nil
	}
	return "", fmt.Errorf("%w for World", ErrTypeNotFound)
}

func entityResolverNameForWorldName(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["name"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findWorldNameByName", nil
	}
	return "", fmt.Errorf("%w for WorldName", ErrTypeNotFound)
}

func entityResolverNameForWorldWithMultipleKeys(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["hello"]
		if !ok {
			break
		}
		if m, ok = val.(map[string]interface{}); !ok {
			break
		}
		val, ok = m["name"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		m = rep
		val, ok = m["foo"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findWorldWithMultipleKeysByHelloNameAndFoo", nil
	}
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["bar"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findWorldWithMultipleKeysByBar", nil
	}
	return "", fmt.Errorf("%w for WorldWithMultipleKeys", ErrTypeNotFound)
}
