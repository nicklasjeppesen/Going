package db

import (
	relationships "github.com/nicklasjeppesen/going_internal/super/db/relationship"
	. "github.com/nicklasjeppesen/going_internal/super/db/types"
)

type BelongsTo[T IDBConnection[T]] = *relationships.BelongsTo[T]
type HasMany[T IDBConnection[T]] = *relationships.HasManyRelation[T]
type HasOne[T IDBConnection[T]] = *relationships.HasOneRelation[T]
type BelongsToMany[T IDBConnection[T]] = *relationships.BelongsToManyRelation[T]
type BelongsToMorph[T IRepository] = *relationships.BelongsToMorphRelation[T]
type HasManyMorph[T IDBConnection[T]] = *relationships.HasManyMorphRelation[T]
