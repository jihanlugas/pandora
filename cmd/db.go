package cmd

import (
	"github.com/jihanlugas/pandora/constant"
	"github.com/jihanlugas/pandora/cryption"
	"github.com/jihanlugas/pandora/db"
	"github.com/jihanlugas/pandora/model"
	"github.com/jihanlugas/pandora/utils"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Run server",
	Long: `With this command you can
	up : create database table
	down :  drop database table
	seed :  insert data table
	`,
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Up table",
	Long:  "Up table",
	Run: func(cmd *cobra.Command, args []string) {
		up()
	},
}

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Down table",
	Long:  "remove public schema, create public schema, restore the default grants",
	Run: func(cmd *cobra.Command, args []string) {
		down()
	},
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed data table",
	Long:  "Seed data table",
	Run: func(cmd *cobra.Command, args []string) {
		seed()
	},
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Down, up, seed table",
	Long:  "Down, up, seed table",
	Run: func(cmd *cobra.Command, args []string) {
		down()
		up()
		seed()
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(upCmd)
	dbCmd.AddCommand(downCmd)
	dbCmd.AddCommand(resetCmd)
	dbCmd.AddCommand(seedCmd)
}

func up() {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	// table
	err = conn.Migrator().AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	err = conn.Migrator().AutoMigrate(&model.Ktp{})
	if err != nil {
		panic(err)
	}

	// view
	vUser := conn.Model(&model.User{}).
		Select("users.*, u1.fullname as create_name, u2.fullname as update_name").
		Joins("left join users u1 on u1.id = users.create_by").
		Joins("left join users u2 on u2.id = users.update_by")

	err = conn.Migrator().CreateView(model.VIEW_USER, gorm.ViewOption{
		Replace: true,
		Query:   vUser,
	})
	if err != nil {
		panic(err)
	}

	vCompany := conn.Model(&model.Ktp{}).
		Select("ktps.*, u1.fullname as create_name, u2.fullname as update_name").
		Joins("left join users u1 on u1.id = ktps.create_by").
		Joins("left join users u2 on u2.id = ktps.update_by")

	err = conn.Migrator().CreateView(model.VIEW_KTP, gorm.ViewOption{
		Replace: true,
		Query:   vCompany,
	})
	if err != nil {
		panic(err)
	}

	tx := conn.Begin()

	userID := utils.GetUniqueID()
	password, err := cryption.EncryptAES64("123456")

	if err != nil {
		panic(err)
	}
	users := []model.User{
		{ID: userID, Role: constant.RoleAdmin, Email: "jihanlugas2@gmail.com", Username: "jihanlugas", NoHp: "6287770333043", Fullname: "Jihan Lugas", Passwd: password, PassVersion: 1, IsActive: true, PhotoID: "", LastLoginDt: nil, CreateBy: userID, UpdateBy: userID},
	}
	tx.Create(&users)

	err = tx.Commit().Error
	if err != nil {
		panic(err)
	}
}

// remove public schema
func down() {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	err = conn.Exec("DROP SCHEMA public CASCADE").Error
	if err != nil {
		panic(err)
	}

	err = conn.Exec("CREATE SCHEMA public").Error
	if err != nil {
		panic(err)
	}

	err = conn.Exec("GRANT ALL ON SCHEMA public TO postgres").Error
	if err != nil {
		panic(err)
	}

	err = conn.Exec("GRANT ALL ON SCHEMA public TO public").Error
	if err != nil {
		panic(err)
	}
}

func seed() {
	////now := time.Now()
	//password, err := cryption.EncryptAES64("123456")
	//if err != nil {
	//	panic(err)
	//}
	//
	//conn, closeConn := db.GetConnection()
	//defer closeConn()
	//
	//tx := conn.Begin()
	//
	//userID := utils.GetUniqueID()
	//demoUserID := utils.GetUniqueID()
	//
	//demoCompanyID := utils.GetUniqueID()
	//
	//luffyPlayerID := utils.GetUniqueID()
	//zoroPlayerID := utils.GetUniqueID()
	//sakazukiPlayerID := utils.GetUniqueID()
	//ishoPlayerID := utils.GetUniqueID()
	//robinPlayerID := utils.GetUniqueID()
	//namiPlayerID := utils.GetUniqueID()
	//ussopPlayerID := utils.GetUniqueID()
	//sanjiPlayerID := utils.GetUniqueID()
	//chopperPlayerID := utils.GetUniqueID()
	//frankyPlayerID := utils.GetUniqueID()
	//brookPlayerID := utils.GetUniqueID()
	//jinbePlayerID := utils.GetUniqueID()
	//enelPlayerID := utils.GetUniqueID()
	//buggyPlayerID := utils.GetUniqueID()
	//arlongPlayerID := utils.GetUniqueID()
	//kuroPlayerID := utils.GetUniqueID()
	//kriegPlayerID := utils.GetUniqueID()
	//smokerPlayerID := utils.GetUniqueID()
	//tashigiPlayerID := utils.GetUniqueID()
	//hinaPlayerID := utils.GetUniqueID()
	//wapolPlayerID := utils.GetUniqueID()
	//crocodilePlayerID := utils.GetUniqueID()
	//lucciPlayerID := utils.GetUniqueID()
	//kakuPlayerID := utils.GetUniqueID()
	//
	//btcGorWahyu := utils.GetUniqueID()
	//btcGorPrs := utils.GetUniqueID()
	//
	//users := []model.User{
	//	{ID: userID, Role: constant.RoleAdmin, Email: "jihanlugas2@gmail.com", Username: "jihanlugas", NoHp: "6287770333043", Fullname: "Jihan Lugas", Passwd: password, PassVersion: 1, IsActive: true, PhotoID: "", LastLoginDt: nil, CreateBy: userID, UpdateBy: userID},
	//	{ID: demoUserID, Role: constant.RoleUser, Email: "demo@gmail.com", Username: "demo", NoHp: "6287770331234", Fullname: "Admin Demo", Passwd: password, PassVersion: 1, IsActive: true, PhotoID: "", LastLoginDt: nil, CreateBy: userID, UpdateBy: userID},
	//}
	//tx.Create(&users)
	//
	//companies := []model.Company{
	//	{ID: demoCompanyID, Name: "Demo", Description: "Demo Company", Balance: 50000, CreateBy: userID, UpdateBy: userID},
	//}
	//tx.Create(&companies)
	//
	//usercompanies := []model.Usercompany{
	//	{UserID: demoUserID, CompanyID: demoCompanyID, IsDefaultCompany: true, IsCreator: true, CreateBy: userID, UpdateBy: userID},
	//}
	//tx.Create(&usercompanies)
	//
	//var players []model.Player
	//playersBtc := []model.Player{
	//	{ID: luffyPlayerID, CompanyID: demoCompanyID, Name: "Monkey D. Luffy", Email: "luffy@gmail.com", NoHp: utils.FormatPhoneTo62("08123456789"), Address: "Fusha Mura", Gender: constant.GENDER_MALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: zoroPlayerID, CompanyID: demoCompanyID, Name: "Roronoa Zoro", Email: "zoro@gmail.com", NoHp: utils.FormatPhoneTo62("08123456777"), Address: "Jl. Kehidupan", Gender: constant.GENDER_MALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: sakazukiPlayerID, CompanyID: demoCompanyID, Name: "Sakazuki", Email: "sakazuki@gmail.com", NoHp: utils.FormatPhoneTo62("08123456779"), Address: "Jl. Perkara", Gender: constant.GENDER_MALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: ishoPlayerID, CompanyID: demoCompanyID, Name: "Isho", Email: "isho@gmail.com", NoHp: utils.FormatPhoneTo62("081234654789"), Address: "Jl. Yang Salah", Gender: constant.GENDER_MALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: robinPlayerID, CompanyID: demoCompanyID, Name: "Nico Robin", Email: "robin@gmail.com", NoHp: utils.FormatPhoneTo62("081234654789"), Address: "Jl. Yang Tersesat", Gender: constant.GENDER_FEMALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: namiPlayerID, CompanyID: demoCompanyID, Name: "Nami", Email: "nami@gmail.com", NoHp: utils.FormatPhoneTo62("081234654789"), Address: "Jl. Yang Salah", Gender: constant.GENDER_FEMALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: ussopPlayerID, CompanyID: demoCompanyID, Name: "Ussop", Email: "ussop@gmail.com", NoHp: utils.FormatPhoneTo62("081234654123"), Address: "Jl. Yang Bohong", Gender: constant.GENDER_MALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: sanjiPlayerID, CompanyID: demoCompanyID, Name: "Vinsmoke Sanji", Email: "sanji@gmail.com", NoHp: utils.FormatPhoneTo62("081234654111"), Address: "Jl. Ke Wanita", Gender: constant.GENDER_MALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: chopperPlayerID, CompanyID: demoCompanyID, Name: "Tony Tony Chopper", Email: "chopper@gmail.com", NoHp: utils.FormatPhoneTo62("081234654112"), Address: "Jl. Medis", Gender: constant.GENDER_MALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: frankyPlayerID, CompanyID: demoCompanyID, Name: "Franky", Email: "franky@gmail.com", NoHp: utils.FormatPhoneTo62("081234654112"), Address: "Jl. Lelaki", Gender: constant.GENDER_MALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: brookPlayerID, CompanyID: demoCompanyID, Name: "Brook", Email: "brook@gmail.com", NoHp: utils.FormatPhoneTo62("081234654114"), Address: "Jl. Menepati Janji", Gender: constant.GENDER_MALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: jinbePlayerID, CompanyID: demoCompanyID, Name: "Jinbe", Email: "jinbe@gmail.com", NoHp: utils.FormatPhoneTo62("081234654115"), Address: "Jl. Yang Dihormati", Gender: constant.GENDER_MALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: enelPlayerID, CompanyID: demoCompanyID, Name: "Enel", Email: "enel@gmail.com", NoHp: utils.FormatPhoneTo62("081234654115"), Address: "Jl. Yang Dihormati", Gender: constant.GENDER_MALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: buggyPlayerID, CompanyID: demoCompanyID, Name: "Buggy", Email: "buggy@gmail.com", NoHp: utils.FormatPhoneTo62("081234654115"), Address: "Jl. Yang Dihormati", Gender: constant.GENDER_MALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: arlongPlayerID, CompanyID: demoCompanyID, Name: "Arlong", Email: "arlong@gmail.com", NoHp: utils.FormatPhoneTo62("081234654115"), Address: "Jl. Yang Dihormati", Gender: constant.GENDER_MALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: kuroPlayerID, CompanyID: demoCompanyID, Name: "Kuro", Email: "kuro@gmail.com", NoHp: utils.FormatPhoneTo62("081234654115"), Address: "Jl. Yang Dihormati", Gender: constant.GENDER_MALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: kriegPlayerID, CompanyID: demoCompanyID, Name: "Krieg", Email: "krieg@gmail.com", NoHp: utils.FormatPhoneTo62("081234654115"), Address: "Jl. Yang Dihormati", Gender: constant.GENDER_MALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: smokerPlayerID, CompanyID: demoCompanyID, Name: "Smoker", Email: "smoker@gmail.com", NoHp: utils.FormatPhoneTo62("081234654115"), Address: "Jl. Yang Dihormati", Gender: constant.GENDER_MALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: tashigiPlayerID, CompanyID: demoCompanyID, Name: "Tashigi", Email: "tashigi@gmail.com", NoHp: utils.FormatPhoneTo62("081234654115"), Address: "Jl. Yang Dihormati", Gender: constant.GENDER_FEMALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: hinaPlayerID, CompanyID: demoCompanyID, Name: "Hina", Email: "hina@gmail.com", NoHp: utils.FormatPhoneTo62("081234654115"), Address: "Jl. Yang Dihormati", Gender: constant.GENDER_FEMALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: wapolPlayerID, CompanyID: demoCompanyID, Name: "Wapol", Email: "wapol@gmail.com", NoHp: utils.FormatPhoneTo62("081234654115"), Address: "Jl. Yang Dihormati", Gender: constant.GENDER_MALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: crocodilePlayerID, CompanyID: demoCompanyID, Name: "Crocodile", Email: "crocodile@gmail.com", NoHp: utils.FormatPhoneTo62("081234654115"), Address: "Jl. Yang Dihormati", Gender: constant.GENDER_MALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: lucciPlayerID, CompanyID: demoCompanyID, Name: "Lucci", Email: "lucci@gmail.com", NoHp: utils.FormatPhoneTo62("081234654115"), Address: "Jl. Yang Dihormati", Gender: constant.GENDER_MALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//	{ID: kakuPlayerID, CompanyID: demoCompanyID, Name: "Kaku", Email: "kaku@gmail.com", NoHp: utils.FormatPhoneTo62("081234654115"), Address: "Jl. Yang Dihormati", Gender: constant.GENDER_MALE, IsActive: true, PhotoID: "", CreateBy: userID, UpdateBy: userID},
	//}
	//
	//players = append(players, playersBtc...)
	//
	//tx.Create(&players)
	//
	//gors := []model.Gor{
	//	{ID: btcGorWahyu, CompanyID: demoCompanyID, Name: "Gor Wahyu", Description: "Gor Wahyu Gobah", Address: "Jl. Sumatra", NormalGamePrice: 8000, RubberGamePrice: 11000, BallPrice: 3000, CreateBy: userID, UpdateBy: userID},
	//	{ID: btcGorPrs, CompanyID: demoCompanyID, Name: "Gor PRS", Description: "Gor Panam Raya Square", Address: "Jl. HR. Subrantas", NormalGamePrice: 7000, RubberGamePrice: 10000, BallPrice: 3000, CreateBy: userID, UpdateBy: userID},
	//}
	//tx.Create(&gors)
	//
	//gameDt := time.Date(2023, time.September, 6, 20, 0, 0, 0, time.Local)
	//
	//games := []model.Game{}
	//for i := 0; i < 20; i++ {
	//	gameDt = gameDt.Add(time.Hour * 24 * 7)
	//	gameDtNext := gameDt.Add(time.Hour * 24)
	//	newGame := []model.Game{
	//		{ID: utils.GetUniqueID(), CompanyID: demoCompanyID, GorID: btcGorPrs, Name: fmt.Sprintf("Game %d", i*2+1), Description: fmt.Sprintf("Game %d Generated", i*2+1), NormalGamePrice: 7000, RubberGamePrice: 10000, BallPrice: 3000, GameDt: gameDt, IsFinish: true, CreateBy: userID, UpdateBy: userID},
	//		{ID: utils.GetUniqueID(), CompanyID: demoCompanyID, GorID: btcGorWahyu, Name: fmt.Sprintf("Game %d", i*2+2), Description: fmt.Sprintf("Game %d Generated", i*2+2), NormalGamePrice: 8000, RubberGamePrice: 11000, BallPrice: 3000, GameDt: gameDtNext, IsFinish: true, CreateBy: userID, UpdateBy: userID},
	//	}
	//	games = append(games, newGame...)
	//}
	//tx.Create(&games)
	//
	//gameplayers := []model.Gameplayer{}
	//for _, game := range games {
	//	newGameplayer := []model.Gameplayer{
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: luffyPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: zoroPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: sakazukiPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: ishoPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: robinPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: namiPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: ussopPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: sanjiPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: chopperPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: frankyPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: brookPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: jinbePlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: enelPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: buggyPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: arlongPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: kuroPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: kriegPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: smokerPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: tashigiPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: hinaPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: wapolPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: crocodilePlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: lucciPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//		{CompanyID: demoCompanyID, GameID: game.ID, PlayerID: kakuPlayerID, NormalGame: rand.Int63n(5), RubberGame: rand.Int63n(3), Ball: rand.Int63n(10), IsPay: true, SetWin: rand.Int63n(8), Point: rand.Int63n(8) - 4, CreateBy: userID, UpdateBy: userID},
	//	}
	//
	//	gameplayers = append(gameplayers, newGameplayer...)
	//}
	//tx.Create(&gameplayers)
	//
	//err = tx.Commit().Error
	//if err != nil {
	//	panic(err)
	//}
}
