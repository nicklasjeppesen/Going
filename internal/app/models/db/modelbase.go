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

func NewHasManyMorph[T IDBConnection[T]](current T, relationFrom IRepository) *relationships.HasManyMorphRelation[T] {
	return relationships.NewHasManyMorph(current, relationFrom)
}

func NewBelongsToMorph[T IRepository](relations []T, delegateAble string, relationTo IRepository) *relationships.BelongsToMorphRelation[T] {
	return relationships.NewBelongsToMorph(delegateAble, relations, relationTo)
}

func NewBelongsToMany[T IDBConnection[T]](relation T, relationTo IRepository) *relationships.BelongsToManyRelation[T] {
	return relationships.NewBelongsToMany(relation, relationTo)
}

func NewHasOne[T IDB[T]](holder T, relationFrom IRepository) *relationships.HasOneRelation[T] {
	return relationships.NewHasOne(holder, relationFrom)
}

func NewHasMany[T IDB[T]](current T, relationFrom IRepository) *relationships.HasManyRelation[T] {
	return relationships.NewHasMany(current, relationFrom)
}

func NewBelongsTo[T IDBConnection[T]](current T, relationFrom IRepository) *relationships.BelongsTo[T] {
	return relationships.NewBelongsTo(current, relationFrom)
}
