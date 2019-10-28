package hotpepper

// Shop : Hotpeperの店舗情報の型
type Shop struct {
	NameKana  string `json:"name_kana"`
	OtherMemo string `json:"other_memo"`
	Photo     struct {
		Mobile struct {
			L string `json:"l"`
			S string `json:"s"`
		} `json:"mobile"`
		Pc struct {
			L string `json:"l"`
			M string `json:"m"`
			S string `json:"s"`
		} `json:"pc"`
	} `json:"photo"`
	LargeArea struct {
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"large_area"`
	PartyCapacity    string `json:"party_capacity"`
	LargeServiceArea struct {
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"large_service_area"`
	MobileAccess string `json:"mobile_access"`
	ID           string `json:"id"`
	Address      string `json:"address"`
	Lng          string `json:"lng"`
	Course       string `json:"course"`
	Show         string `json:"show"`
	Parking      string `json:"parking"`
	NonSmoking   string `json:"non_smoking"`
	Horigotatsu  string `json:"horigotatsu"`
	Name         string `json:"name"`
	Genre        struct {
		Name  string `json:"name"`
		Catch string `json:"catch"`
		Code  string `json:"code"`
	} `json:"genre"`
	Open     string `json:"open"`
	Card     string `json:"card"`
	Tatami   string `json:"tatami"`
	Charter  string `json:"charter"`
	Wifi     string `json:"wifi"`
	SubGenre struct {
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"sub_genre"`
	ShopDetailMemo string `json:"shop_detail_memo"`
	Band           string `json:"band"`
	MiddleArea     struct {
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"middle_area"`
	Lat       string `json:"lat"`
	Karaoke   string `json:"karaoke"`
	LogoImage string `json:"logo_image"`
	Midnight  string `json:"midnight"`
	Budget    struct {
		Average string `json:"average"`
		Name    string `json:"name"`
		Code    string `json:"code"`
	} `json:"budget"`
	Urls struct {
		Pc string `json:"pc"`
	} `json:"urls"`
	English     string `json:"english"`
	Lunch       string `json:"lunch"`
	ServiceArea struct {
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"service_area"`
	Close       string `json:"close"`
	BudgetMemo  string `json:"budget_memo"`
	Tv          string `json:"tv"`
	PrivateRoom string `json:"private_room"`
	CouponUrls  struct {
		Sp string `json:"sp"`
		Pc string `json:"pc"`
	} `json:"coupon_urls"`
	BarrierFree string `json:"barrier_free"`
	SmallArea   struct {
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"small_area"`
	Wedding     string `json:"wedding"`
	Access      string `json:"access"`
	Child       string `json:"child"`
	Capacity    string `json:"capacity"`
	Pet         string `json:"pet"`
	KtaiCoupon  string `json:"ktai_coupon"`
	FreeFood    string `json:"free_food"`
	StationName string `json:"station_name"`
	Catch       string `json:"catch"`
	FreeDrink   string `json:"free_drink"`
}
