package PokeAPI

type NamedAPIResourceList struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationArea struct {
	Id                     int                   `json:"id"`
	Name                   string                `json:"name"`
	Game_index             int                   `json:"game_index"`
	Encounter_method_rates []EncounterMethodRate `json:"encounter_method_rates"`
	Location               NamedAPIResource      `json:"location"`
	Names                  []Name                `json:"names"`
	Pokemon_encounters     []PokemonEncounter    `json:"pokemon_encounters"`
}

type Name struct {
	Name      string           `json:"name"`
	Languagei NamedAPIResource `json:"language"`
}

type EncounterMethodRate struct {
	Encounter_method NamedAPIResource          `json:"encounter_method"`
	Version_details  []EncounterVersionDetails `json:"version_details"`
}

type EncounterVersionDetails struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type PokemonEncounter struct {
	Pokemon         NamedAPIResource         `json:"pokemon"`
	Version_details []VersionEncounterDetail `json:"version_details"`
}

type VersionEncounterDetail struct {
	Version           NamedAPIResource `json:"version"`
	Max_chance        int              `json:"max_chance"`
	Encounter_details []Encounter      `json:"encounter_details"`
}

type Encounter struct {
	Min_level        int                `json:"min_level"`
	Max_level        int                `json:"max_level"`
	Condition_values []NamedAPIResource `json:"condition_values"`
	Chance           int                `json:"chance"`
	Method           NamedAPIResource   `json:"method"`
}

type Pokemon struct {
	Id                       int                  `json:"id"`
	Name                     string               `json:"name"`
	Base_experience          int                  `json:"base_experience"`
	Height                   int                  `json:"height"`
	Is_default               bool                 `json:"is_default"`
	Order                    int                  `json:"order"`
	Weight                   int                  `json:"weight"`
	Abilities                []PokemonAbility     `json:"abilities"`
	Forms                    []NamedAPIResource   `json:"forms"`
	Game_indices             []VersionGameIndex   `json:"game_indices"`
	Held_items               []PokemonHeldItem    `json:"held_items"`
	Location_area_encounters string               `json:"location_area_encounters"`
	Moves                    []PokemonMove        `json:"moves"`
	Past_types               []PokemonTypePast    `json:"past_types"`
	Past_abilities           []PokemonAbilityPast `json:"past_abilities"`
	Sprites                  PokemonSprites       `json:"sprites"`
	Cries                    PokemonCries         `json:"cries"`
	Species                  NamedAPIResource     `json:"species"`
	Stats                    []PokemonStat        `json:"stats"`
	Types                    []PokemonType        `json:"types"`
}

type PokemonAbility struct {
	Is_hidden bool             `json:"is_hidden"`
	Slot      int              `json:"slot"`
	Ability   NamedAPIResource `json:"ability"`
}

type PokemonType struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

type PokemonForm struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

type PokemonTypePast struct {
	Generation NamedAPIResource `json:"generation"`
	Types      []PokemonType    `json:"types"`
}

type PokemonAbilityPast struct {
	Generation NamedAPIResource `json:"generation"`
	Ability    NamedAPIResource `json:"ability"`
}

type PokemonHeldItem struct {
	Item            NamedAPIResource         `json:"item"`
	Version_details []PokemonHeldItemVersion `json:"version_details"`
}

type PokemonHeldItemVersion struct {
	Version NamedAPIResource `json:"version"`
	Rarity  int              `json:"rarity"`
}

type PokemonMove struct {
	Move                  NamedAPIResource     `json:"move"`
	Version_group_details []PokemonMoveVersion `json:"version_group_details"`
}

type PokemonMoveVersion struct {
	Move_learn_method NamedAPIResource `json:"move_learn_method"`
	Version_group     NamedAPIResource `json:"version_group"`
	Level_learned_at  int              `json:"level_learned_at"`
	Order             int              `json:"order"`
}

type PokemonSprites struct {
	Front_default      string `json:"front_default"`
	Front_shiny        string `json:"front_shiny"`
	Front_female       string `json:"front_female"`
	Front_shiny_female string `json:"front_shiny_female"`
	Back_default       string `json:"back_default"`
	Back_shiny         string `json:"back_shiny"`
	Back_female        string `json:"back_female"`
	Back_shiny_female  string `json:"back_shiny_female"`
}

type PokemonCries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type VersionGameIndex struct {
	Game_index int              `json:"game_index"`
	Version    NamedAPIResource `json:"version"`
}

type PokemonStat struct {
	Stat      NamedAPIResource `json:"stat"`
	Effort    int              `json:"effort"`
	Base_stat int              `json:"base_stat"`
}
