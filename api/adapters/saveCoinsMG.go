package adapters

//import (
//	"golang.org/x/mod/sumdb/storage"
//)
//
//type serviceCoins struct {
//	storage.Storage
//}

//func (s serviceCoins) Create(ctx context.Context, coin entities.CoinsData) (coins entities.CoinsData, err error) {
//
//}

//var db *sql.DB

//var conf = config.GetConfig()
//var pass = conf.PosgrePass

//func init() {
//
//	conn, err := pgx.Connect(context.Background(), "jdbc:postgresql://localhost:5432/crypto")
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
//		os.Exit(1)
//	}
//
//	var greeting string
//	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
//		os.Exit(1)
//	}
//
//	fmt.Println(greeting)
//
//	//DBConnection := PosgreDB{
//	//	Host:     "localhost",
//	//	port:     5432,
//	//	user:     "msylniahin",
//	//	password: config.GetConfig().PosgrePass,
//	//	dbsName:  "crypto",
//	//}
//	// connection string
//
//	fmt.Println("Connected!")
//}
