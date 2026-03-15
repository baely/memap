package models

var (
	work                     = &Node{Label: "Work", Link: "https://atlassian.com", Position: &Position{-37.816089, 144.961657}}
	githubAsABlogService     = &Node{Label: "GitHub as a Blog Service", Link: "https://blog.baileys.dev/posts/github-as-a-blog-service/", Position: &Position{-37.817478342724414, 144.9691839090097}}
	myCodeDidntNeedARefactor = &Node{Label: "My Code Didn't Need a Refactor", Link: "https://blog.baileys.dev/posts/code-with-a-passport/", Position: &Position{-37.81671058380661, 144.97176192067823}}
	n                        = &Node{Label: "日本に到着\n", Link: "https://devhou.se/103/", Position: &Position{-37.81886764818794, 144.96579537017587}}
	n_2                      = &Node{Label: "最初の一日全体\n", Link: "https://devhou.se/122", Position: &Position{-37.819732401727386, 144.96485830729964}}
	n_3                      = &Node{Label: "回復\n", Link: "https://devhou.se/147/", Position: &Position{-37.82000912286001, 144.96627703800942}}
	n_4                      = &Node{Label: "東京のまたもやどんよりした日\n", Link: "https://devhou.se/170/", Position: &Position{-37.82081533350024, 144.96562517771545}}
	n_5                      = &Node{Label: "個人記録\n", Link: "https://dehvou.se/183", Position: &Position{-37.82130623090021, 144.96681435558685}}
	n_6                      = &Node{Label: "旅行\n", Link: "https://devhou.se/203", Position: &Position{-37.821931596452515, 144.9661638046224}}
	n_7                      = &Node{Label: "ユニバーサルスタジオジャパン\n", Link: "https://devhou.se/222", Position: &Position{-37.822270194010706, 144.9672506625963}}
	n_8                      = &Node{Label: "大阪城\n", Link: "https://devhou.se/249", Position: &Position{-37.823041888150776, 144.96676786199754}}
	n_9                      = &Node{Label: "のんびりな日\n", Link: "https://dehvou.se/267", Position: &Position{-37.82353671986982, 144.96787757583087}}
	n_10                     = &Node{Label: "発進\n", Link: "https://devhou.se/293", Position: &Position{-37.82441603255188, 144.967461418206}}
	n_11                     = &Node{Label: "서울, 대한민국\n", Link: "https://devhou.se/329", Position: &Position{-37.824830514099894, 144.96851315017923}}
	n_12                     = &Node{Label: "제주도, 대한민국\n", Link: "https://devhou.se/343/", Position: &Position{-37.825596670411656, 144.96802457537314}}
	n_13                     = &Node{Label: "台北，台灣\n", Link: "https://devhou.se/367/", Position: &Position{-37.826097304520935, 144.96922866008882}}
	n1                       = &Node{Label: "1日目\n", Link: "https://devhou.se/473", Position: &Position{-37.81815244126289, 144.96836305006462}}
	n_14                     = &Node{Label: "松島\n", Link: "https://devhou.se/487/", Position: &Position{-37.818770203782364, 144.96771282508135}}
	hub                      = &Node{Label: "青葉城とHUBでの野球\n", Link: "https://devhou.se/505", Position: &Position{-37.819264309728034, 144.96874512040668}}
	n_15                     = &Node{Label: "ニッカウヰスキー蒸留所\n", Link: "https://devhou.se/520/", Position: &Position{-37.81993869916701, 144.96830817419536}}
	n_16                     = &Node{Label: "東京（「ディーラーズチョイス」以前）\n", Link: "https://dehvou.se/540", Position: &Position{-37.82040907153537, 144.96930146266433}}
	n1_2                     = &Node{Label: "ディーラーズチョイス 1日目\n", Link: "https://devhou.se/560/", Position: &Position{-37.821032595736796, 144.9688726870104}}
	n_17                     = &Node{Label: "渋谷\n", Link: "https://devhou.se/578/", Position: &Position{-37.82149917229464, 144.96994984171735}}
	n_18                     = &Node{Label: "チリング、カピバラ、カレー\n", Link: "https://devhou.se/625/", Position: &Position{-37.82205309609615, 144.9693277045534}}
	n_19                     = &Node{Label: "金沢への旅\n", Link: "https://devhou.se/642/", Position: &Position{-37.82253489526877, 144.97038656106542}}
	n_20                     = &Node{Label: "金沢城\n", Link: "https://devhou.se/662/", Position: &Position{-37.82312094656945, 144.96986541313962}}
	n_21                     = &Node{Label: "金沢での最後の数日間\n", Link: "https://devhou.se/707", Position: &Position{-37.823559968255196, 144.97092615960264}}
	n_22                     = &Node{Label: "福岡\n", Link: "https://devhou.se/750/", Position: &Position{-37.82431226463998, 144.97035879920864}}
	n_23                     = &Node{Label: "岡山、姫路、大阪、東京\n", Link: "https://devhou.se/769/", Position: &Position{-37.82478198630325, 144.97156784184884}}
	devhouse2                = &Node{Label: "Devhouse 2の反省\n", Link: "https://dehvou.se/779", Position: &Position{-37.82560554904477, 144.97094597776533}}
	github                   = &Node{Label: "GitHub", Link: "https://github.com/baely", Description: "My GitHub page", Position: &Position{-37.81489447040561, 144.96700913885638}}
	linkedin                 = &Node{Label: "LinkedIn", Link: "https://linkedin.com/in/baileybutler1", Description: "My LinkedIn page", Position: &Position{-37.81457994700561, 144.9684321225668}}
	flindersSpencer          = &Node{Label: "Flinders & Spencer", Position: &Position{-37.821073, 144.955057}}
	flindersKing             = &Node{Label: "Flinders & King", Position: &Position{-37.820345, 144.957536}}
	flindersWilliam          = &Node{Label: "Flinders & William", Position: &Position{-37.819618, 144.960015}}
	flindersQueen            = &Node{Label: "Flinders & Queen", Position: &Position{-37.81889, 144.962494}}
	flindersElizabeth        = &Node{Label: "Flinders & Elizabeth", Position: &Position{-37.818162, 144.964973}}
	flindersSwanston         = &Node{Label: "Flinders & Swanston", Position: &Position{-37.817435, 144.967453}}
	flindersRussell          = &Node{Label: "Flinders & Russell", Position: &Position{-37.816707, 144.969932}}
	flindersExhibition       = &Node{Label: "Flinders & Exhibition", Position: &Position{-37.81598, 144.972411}}
	flindersSpring           = &Node{Label: "Flinders & Spring", Position: &Position{-37.815252, 144.97489}}
	bourkeSpencer            = &Node{Label: "Bourke & Spencer", Position: &Position{-37.817113, 144.953227}}
	bourkeKing               = &Node{Label: "Bourke & King", Position: &Position{-37.816385, 144.955706}}
	bourkeWilliam            = &Node{Label: "Bourke & William", Position: &Position{-37.815657, 144.958185}}
	bourkeQueen              = &Node{Label: "Bourke & Queen", Position: &Position{-37.81493, 144.960664}}
	bourkeElizabeth          = &Node{Label: "Bourke & Elizabeth", Position: &Position{-37.814202, 144.963143}}
	bourkeSwanston           = &Node{Label: "Bourke & Swanston", Position: &Position{-37.813474, 144.965623}}
	bourkeRussell            = &Node{Label: "Bourke & Russell", Position: &Position{-37.812747, 144.968102}}
	bourkeExhibition         = &Node{Label: "Bourke & Exhibition", Position: &Position{-37.812019, 144.970581}}
	bourkeSpring             = &Node{Label: "Bourke & Spring", Position: &Position{-37.811292, 144.97306}}
	laTrobeSpencer           = &Node{Label: "La Trobe & Spencer", Position: &Position{-37.813152, 144.951397}}
	laTrobeKing              = &Node{Label: "La Trobe & King", Position: &Position{-37.812424, 144.953876}}
	laTrobeWilliam           = &Node{Label: "La Trobe & William", Position: &Position{-37.811697, 144.956355}}
	laTrobeQueen             = &Node{Label: "La Trobe & Queen", Position: &Position{-37.810969, 144.958834}}
	laTrobeElizabeth         = &Node{Label: "La Trobe & Elizabeth", Position: &Position{-37.810242, 144.961313}}
	laTrobeSwanston          = &Node{Label: "La Trobe & Swanston", Position: &Position{-37.809514, 144.963793}}
	laTrobeRussell           = &Node{Label: "La Trobe & Russell", Position: &Position{-37.808786, 144.966272}}
	laTrobeExhibition        = &Node{Label: "La Trobe & Exhibition", Position: &Position{-37.808059, 144.968751}}
	laTrobeSpring            = &Node{Label: "La Trobe & Spring", Position: &Position{-37.807331, 144.97123}}
	collinsQueen             = &Node{Label: "Collins & Queen", Position: &Position{-37.81691, 144.961579}}
	lonsdaleQueen            = &Node{Label: "Lonsdale & Queen", Position: &Position{-37.812949, 144.959749}}
	collinsElizabeth         = &Node{Label: "Collins & Elizabeth", Position: &Position{-37.816182, 144.964058}}
	lonsdaleElizabeth        = &Node{Label: "Lonsdale & Elizabeth", Position: &Position{-37.812222, 144.962228}}
	collinsSwanston          = &Node{Label: "Collins & Swanston", Position: &Position{-37.815455, 144.966538}}
	lonsdaleSwanston         = &Node{Label: "Lonsdale & Swanston", Position: &Position{-37.811494, 144.964708}}
	lonsdaleSpencer          = &Node{Label: "Lonsdale & Spencer", Position: &Position{-37.815132, 144.952312}}
	lonsdaleKing             = &Node{Label: "Lonsdale & King", Position: &Position{-37.814405, 144.954791}}
	lonsdaleWilliam          = &Node{Label: "Lonsdale & William", Position: &Position{-37.813677, 144.95727}}
	lonsdaleSpring           = &Node{Label: "Lonsdale & Spring", Position: &Position{-37.809311, 144.972145}}
	p                        = &Node{Label: "", Position: &Position{-37.82689712162978, 144.9691518173112}}
	p_2                      = &Node{Label: "", Position: &Position{-37.82617694582086, 144.9718554325563}}
	collinsSpencer           = &Node{Label: "Collins & Spencer", Position: &Position{-37.819093, 144.954142}}
	collinsKing              = &Node{Label: "Collins & King", Position: &Position{-37.818365, 144.956621}}
	collinsWilliam           = &Node{Label: "Collins & William", Position: &Position{-37.817638, 144.9591}}
	collinsSpring            = &Node{Label: "Collins & Spring", Position: &Position{-37.813272, 144.973975}}
)

var (
	blogStreet = &Path{Label: "Blog Street", Type: PathTypeRoad, Nodes: []*Node{
		flindersSpencer,
		flindersKing,
		flindersWilliam,
		flindersQueen,
		flindersElizabeth,
		flindersSwanston,
		flindersRussell,
		flindersExhibition,
		flindersSpring,
	}}
	bourkeStreet = &Path{Label: "Bourke Street", Type: PathTypeRoad, Nodes: []*Node{
		bourkeSpencer,
		bourkeKing,
		bourkeWilliam,
		bourkeQueen,
		bourkeElizabeth,
		bourkeSwanston,
		bourkeRussell,
		bourkeExhibition,
		bourkeSpring,
	}}
	laTrobeStreet = &Path{Label: "La Trobe Street", Type: PathTypeRoad, Nodes: []*Node{
		laTrobeSpencer,
		laTrobeKing,
		laTrobeWilliam,
		laTrobeQueen,
		laTrobeElizabeth,
		laTrobeSwanston,
		laTrobeRussell,
		laTrobeExhibition,
		laTrobeSpring,
	}}
	queenStreet = &Path{Label: "Queen Street", Type: PathTypeRoad, Nodes: []*Node{
		flindersQueen,
		collinsQueen,
		bourkeQueen,
		lonsdaleQueen,
		laTrobeQueen,
	}}
	elizabethStreet = &Path{Label: "Elizabeth Street", Type: PathTypeRoad, Nodes: []*Node{
		flindersElizabeth,
		collinsElizabeth,
		bourkeElizabeth,
		lonsdaleElizabeth,
		laTrobeElizabeth,
	}}
	swanstonStreet = &Path{Label: "Swanston Street", Type: PathTypeRoad, Nodes: []*Node{
		flindersSwanston,
		collinsSwanston,
		bourkeSwanston,
		lonsdaleSwanston,
		laTrobeSwanston,
	}}
	randomStreetWest = &Path{Label: "Random Street West", Type: PathTypeRoad, Nodes: []*Node{
		lonsdaleSpencer,
		lonsdaleKing,
		lonsdaleWilliam,
		lonsdaleQueen,
		lonsdaleElizabeth,
	}}
	randomStreetEast = &Path{Label: "Random Street East", Type: PathTypeRoad, Nodes: []*Node{
		lonsdaleSpring,
		lonsdaleElizabeth,
	}}
	devhouse1Street = &Path{Label: "devhou.se 1 Street", Type: PathTypeRoad, Nodes: []*Node{
		flindersElizabeth,
		p,
	}}
	devhouse2Street = &Path{Label: "devhou.se 2 Street", Type: PathTypeRoad, Nodes: []*Node{
		flindersSwanston,
		p_2,
	}}
	collinsStreet = &Path{Label: "Collins Street", Type: PathTypeRoad, Nodes: []*Node{
		collinsSpencer,
		collinsKing,
		collinsWilliam,
		collinsQueen,
		collinsElizabeth,
	}}
	professionalLink = &Path{Label: "Professional Link", Type: PathTypeRoad, Nodes: []*Node{
		collinsElizabeth,
		collinsSpring,
	}}
)

var SampleMap = &Map{
	Nodes: []*Node{
		work,
		githubAsABlogService,
		myCodeDidntNeedARefactor,
		n,
		n_2,
		n_3,
		n_4,
		n_5,
		n_6,
		n_7,
		n_8,
		n_9,
		n_10,
		n_11,
		n_12,
		n_13,
		n1,
		n_14,
		hub,
		n_15,
		n_16,
		n1_2,
		n_17,
		n_18,
		n_19,
		n_20,
		n_21,
		n_22,
		n_23,
		devhouse2,
		github,
		linkedin,
	},
	Paths: []*Path{
		blogStreet,
		bourkeStreet,
		laTrobeStreet,
		queenStreet,
		elizabethStreet,
		swanstonStreet,
		randomStreetWest,
		randomStreetEast,
		devhouse1Street,
		devhouse2Street,
		collinsStreet,
		professionalLink,
	},
}
