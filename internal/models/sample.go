package models

// Melbourne CBD Hoddle Grid street network.
// Pre-computed from corner coordinates:
//   SW (Flinders & Spencer):  -37.821073, 144.955057
//   SE (Flinders & Spring):   -37.815252, 144.974890
//   NW (La Trobe & Spencer):  -37.813152, 144.951397
//   NE (La Trobe & Spring):   -37.807331, 144.971230
// Victoria St endpoint:        -37.805864, 144.955975

// --- Grid intersection nodes ---

var (
	// Flinders Street (row 0, southernmost)
	flindersSpencer    = Node{Label: "Flinders & Spencer", Position: Position{-37.821073, 144.955057}}
	flindersKing       = Node{Label: "Flinders & King", Position: Position{-37.820345, 144.957536}}
	flindersWilliam    = Node{Label: "Flinders & William", Position: Position{-37.819618, 144.960015}}
	flindersQueen      = Node{Label: "Flinders & Queen", Position: Position{-37.818890, 144.962494}}
	flindersElizabeth  = Node{Label: "Flinders & Elizabeth", Position: Position{-37.818162, 144.964973}}
	flindersSwanston   = Node{Label: "Flinders & Swanston", Position: Position{-37.817435, 144.967453}}
	flindersRussell    = Node{Label: "Flinders & Russell", Position: Position{-37.816707, 144.969932}}
	flindersExhibition = Node{Label: "Flinders & Exhibition", Position: Position{-37.815980, 144.972411}}
	flindersSpring     = Node{Label: "Flinders & Spring", Position: Position{-37.815252, 144.974890}}

	// Collins Street (row 1)
	collinsSpencer    = Node{Label: "Collins & Spencer", Position: Position{-37.819093, 144.954142}}
	collinsKing       = Node{Label: "Collins & King", Position: Position{-37.818365, 144.956621}}
	collinsWilliam    = Node{Label: "Collins & William", Position: Position{-37.817638, 144.959100}}
	collinsQueen      = Node{Label: "Collins & Queen", Position: Position{-37.816910, 144.961579}}
	collinsElizabeth  = Node{Label: "Collins & Elizabeth", Position: Position{-37.816182, 144.964058}}
	collinsSwanston   = Node{Label: "Collins & Swanston", Position: Position{-37.815455, 144.966538}}
	collinsRussell    = Node{Label: "Collins & Russell", Position: Position{-37.814727, 144.969017}}
	collinsExhibition = Node{Label: "Collins & Exhibition", Position: Position{-37.813999, 144.971496}}
	collinsSpring     = Node{Label: "Collins & Spring", Position: Position{-37.813272, 144.973975}}

	// Bourke Street (row 2)
	bourkeSpencer    = Node{Label: "Bourke & Spencer", Position: Position{-37.817113, 144.953227}}
	bourkeKing       = Node{Label: "Bourke & King", Position: Position{-37.816385, 144.955706}}
	bourkeWilliam    = Node{Label: "Bourke & William", Position: Position{-37.815657, 144.958185}}
	bourkeQueen      = Node{Label: "Bourke & Queen", Position: Position{-37.814930, 144.960664}}
	bourkeElizabeth  = Node{Label: "Bourke & Elizabeth", Position: Position{-37.814202, 144.963143}}
	bourkeSwanston   = Node{Label: "Bourke & Swanston", Position: Position{-37.813474, 144.965623}}
	bourkeRussell    = Node{Label: "Bourke & Russell", Position: Position{-37.812747, 144.968102}}
	bourkeExhibition = Node{Label: "Bourke & Exhibition", Position: Position{-37.812019, 144.970581}}
	bourkeSpring     = Node{Label: "Bourke & Spring", Position: Position{-37.811292, 144.973060}}

	// Lonsdale Street (row 3)
	lonsdaleSpencer    = Node{Label: "Lonsdale & Spencer", Position: Position{-37.815132, 144.952312}}
	lonsdaleKing       = Node{Label: "Lonsdale & King", Position: Position{-37.814405, 144.954791}}
	lonsdaleWilliam    = Node{Label: "Lonsdale & William", Position: Position{-37.813677, 144.957270}}
	lonsdaleQueen      = Node{Label: "Lonsdale & Queen", Position: Position{-37.812949, 144.959749}}
	lonsdaleElizabeth  = Node{Label: "Lonsdale & Elizabeth", Position: Position{-37.812222, 144.962228}}
	lonsdaleSwanston   = Node{Label: "Lonsdale & Swanston", Position: Position{-37.811494, 144.964708}}
	lonsdaleRussell    = Node{Label: "Lonsdale & Russell", Position: Position{-37.810766, 144.967187}}
	lonsdaleExhibition = Node{Label: "Lonsdale & Exhibition", Position: Position{-37.810039, 144.969666}}
	lonsdaleSpring     = Node{Label: "Lonsdale & Spring", Position: Position{-37.809311, 144.972145}}

	// La Trobe Street (row 4, northernmost grid row)
	laTrobeSpencer    = Node{Label: "La Trobe & Spencer", Position: Position{-37.813152, 144.951397}}
	laTrobeKing       = Node{Label: "La Trobe & King", Position: Position{-37.812424, 144.953876}}
	laTrobeWilliam    = Node{Label: "La Trobe & William", Position: Position{-37.811697, 144.956355}}
	laTrobeQueen      = Node{Label: "La Trobe & Queen", Position: Position{-37.810969, 144.958834}}
	laTrobeElizabeth  = Node{Label: "La Trobe & Elizabeth", Position: Position{-37.810242, 144.961313}}
	laTrobeSwanston   = Node{Label: "La Trobe & Swanston", Position: Position{-37.809514, 144.963793}}
	laTrobeRussell    = Node{Label: "La Trobe & Russell", Position: Position{-37.808786, 144.966272}}
	laTrobeExhibition = Node{Label: "La Trobe & Exhibition", Position: Position{-37.808059, 144.968751}}
	laTrobeSpring     = Node{Label: "La Trobe & Spring", Position: Position{-37.807331, 144.971230}}

	// --- Franklin Street nodes (one block north of La Trobe) ---

	franklinWilliam   = Node{Label: "Franklin & William", Position: Position{-37.809717, 144.955440}}
	franklinQueen     = Node{Label: "Franklin & Queen", Position: Position{-37.808989, 144.957919}}
	franklinElizabeth = Node{Label: "Franklin & Elizabeth", Position: Position{-37.808261, 144.960398}}
	franklinSwanston  = Node{Label: "Franklin & Swanston", Position: Position{-37.807534, 144.962878}}
	franklinRussell   = Node{Label: "Franklin & Russell", Position: Position{-37.806806, 144.965357}}
	franklinVictoria  = Node{Label: "Franklin & Victoria", Position: Position{-37.806776, 144.965459}}

	// --- Victoria Street nodes (diagonal, above the grid) ---

	victoriaPeel       = Node{Label: "Victoria & Peel", Position: Position{-37.805864, 144.955975}}
	victoriaQueen      = Node{Label: "Victoria & Queen", Position: Position{-37.806090, 144.958322}}
	victoriaElizabeth  = Node{Label: "Victoria & Elizabeth", Position: Position{-37.806198, 144.959445}}
	victoriaSwanston   = Node{Label: "Victoria & Swanston", Position: Position{-37.806481, 144.962391}}
	victoriaRussell    = Node{Label: "Victoria & Russell", Position: Position{-37.806764, 144.965338}}
	victoriaExhibition = Node{Label: "Victoria & Exhibition", Position: Position{-37.807048, 144.968284}}

	// --- Other nodes ---
	work = Node{Label: "Work", Position: Position{-37.816089, 144.961657}}
)

// --- Street paths ---

var (
	// East-west streets
	flindersStreet = Path{Label: "Flinders Street", Type: PathTypeRoad, Nodes: []Node{
		flindersSpencer, flindersKing, flindersWilliam, flindersQueen,
		flindersElizabeth, flindersSwanston, flindersRussell, flindersExhibition, flindersSpring,
	}}
	collinsStreet = Path{Label: "Collins Street", Type: PathTypeRoad, Nodes: []Node{
		collinsSpencer, collinsKing, collinsWilliam, collinsQueen,
		collinsElizabeth, collinsSwanston, collinsRussell, collinsExhibition, collinsSpring,
	}}
	bourkeStreet = Path{Label: "Bourke Street", Type: PathTypeRoad, Nodes: []Node{
		bourkeSpencer, bourkeKing, bourkeWilliam, bourkeQueen,
		bourkeElizabeth, bourkeSwanston, bourkeRussell, bourkeExhibition, bourkeSpring,
	}}
	lonsdaleStreet = Path{Label: "Lonsdale Street", Type: PathTypeRoad, Nodes: []Node{
		lonsdaleSpencer, lonsdaleKing, lonsdaleWilliam, lonsdaleQueen,
		lonsdaleElizabeth, lonsdaleSwanston, lonsdaleRussell, lonsdaleExhibition, lonsdaleSpring,
	}}
	laTrobeStreet = Path{Label: "La Trobe Street", Type: PathTypeRoad, Nodes: []Node{
		laTrobeSpencer, laTrobeKing, laTrobeWilliam, laTrobeQueen,
		laTrobeElizabeth, laTrobeSwanston, laTrobeRussell, laTrobeExhibition, laTrobeSpring,
	}}
	franklinStreet = Path{Label: "Franklin Street", Type: PathTypeRoad, Nodes: []Node{
		franklinWilliam, franklinQueen, franklinElizabeth, franklinSwanston, franklinRussell, franklinVictoria,
	}}
	victoriaStreet = Path{Label: "Victoria Street", Type: PathTypeRoad, Nodes: []Node{
		victoriaPeel, victoriaQueen, victoriaElizabeth, victoriaSwanston, victoriaRussell, victoriaExhibition, laTrobeSpring,
	}}

	// North-south streets
	spencerStreet = Path{Label: "Spencer Street", Type: PathTypeRoad, Nodes: []Node{
		flindersSpencer, collinsSpencer, bourkeSpencer, lonsdaleSpencer, laTrobeSpencer,
	}}
	kingStreet = Path{Label: "King Street", Type: PathTypeRoad, Nodes: []Node{
		flindersKing, collinsKing, bourkeKing, lonsdaleKing, laTrobeKing,
	}}
	williamStreet = Path{Label: "William Street", Type: PathTypeRoad, Nodes: []Node{
		flindersWilliam, collinsWilliam, bourkeWilliam, lonsdaleWilliam, laTrobeWilliam,
		franklinWilliam, victoriaPeel,
	}}
	queenStreet = Path{Label: "Queen Street", Type: PathTypeRoad, Nodes: []Node{
		flindersQueen, collinsQueen, bourkeQueen, lonsdaleQueen, laTrobeQueen,
		franklinQueen, victoriaQueen,
	}}
	elizabethStreet = Path{Label: "Elizabeth Street", Type: PathTypeRoad, Nodes: []Node{
		flindersElizabeth, collinsElizabeth, bourkeElizabeth, lonsdaleElizabeth, laTrobeElizabeth,
		victoriaElizabeth,
	}}
	swanstonStreet = Path{Label: "Swanston Street", Type: PathTypeRoad, Nodes: []Node{
		flindersSwanston, collinsSwanston, bourkeSwanston, lonsdaleSwanston, laTrobeSwanston,
		victoriaSwanston,
	}}
	russellStreet = Path{Label: "Russell Street", Type: PathTypeRoad, Nodes: []Node{
		flindersRussell, collinsRussell, bourkeRussell, lonsdaleRussell, laTrobeRussell,
		victoriaRussell,
	}}
	exhibitionStreet = Path{Label: "Exhibition Street", Type: PathTypeRoad, Nodes: []Node{
		flindersExhibition, collinsExhibition, bourkeExhibition, lonsdaleExhibition, laTrobeExhibition,
		victoriaExhibition,
	}}
	springStreet = Path{Label: "Spring Street", Type: PathTypeRoad, Nodes: []Node{
		flindersSpring, collinsSpring, bourkeSpring, lonsdaleSpring, laTrobeSpring,
	}}
)

var SampleMap = Map{
	Nodes: []Node{
		work,
	},
	Paths: []Path{
		// East-west
		flindersStreet, collinsStreet, bourkeStreet, lonsdaleStreet, laTrobeStreet,
		franklinStreet, victoriaStreet,
		// North-south
		spencerStreet, kingStreet, williamStreet, queenStreet,
		elizabethStreet, swanstonStreet, russellStreet, exhibitionStreet, springStreet,
	},
}
