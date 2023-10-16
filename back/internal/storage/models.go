package storage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	consts "nixietech/internal"
)

type ClockWithoutId struct {
	Name      string           `bson:"name"`
	Avatar    string           `bson:"avatar"`
	Photos    []string         `bson:"photos"`
	Price     int              `bson:"price"`
	Existence bool             `bson:"existence"`
	Type      consts.ClockType `bson:"type"`
}
type Clock[T primitive.ObjectID] struct {
	Id        T                `bson:"_id"`
	Name      string           `bson:"name"`
	Avatar    string           `bson:"avatar"`
	Photos    []string         `bson:"photos"`
	Price     int              `bson:"price"`
	Existence bool             `bson:"existence"`
	Type      consts.ClockType `bson:"type"`
}

type OrderWithoutId[T primitive.ObjectID] struct {
	Wishes     string        `bson:"wishes"`
	Contact    string        `bson:"contact"`
	ClockId    T             `bson:"clock_id"`
	Base       BasedSettings `bson:"base"`
	SettingsId T             `bson:"settings_id"`
}
type Order[T primitive.ObjectID] struct {
	Id         T             `bson:"_id"`
	Wishes     string        `bson:"wishes"`
	Contact    string        `bson:"contact"`
	ClockId    T             `bson:"clock_id"`
	Base       BasedSettings `bson:"base"`
	SettingsId T
}
type BasedSettings struct {
	LegsType      []string `bson:"legs_type"`
	EngravingType []string `bson:"engraving_type"`
	PackageType   []string `bson:"package_type"`
}

type TypeOneSettingsWithoutId struct {
	LampsType           []string `bson:"lamp_types"`
	DecorativeRingsType []string `bson:"decorative_rings_type"`
}
type TypeOneSettings[T primitive.ObjectID] struct {
	Id                  T        `bson:"_id"`
	LampsType           []string `bson:"lamp_types"`
	DecorativeRingsType []string `bson:"decorative_rings_type"`
}

type TypeTwoSettingsWithoutId struct {
	Test []string `bson:"test"`
}
type TypeTwoSettings[T primitive.ObjectID] struct {
	Id   T        `bson:"_id"`
	Test []string `bson:"test"`
}
