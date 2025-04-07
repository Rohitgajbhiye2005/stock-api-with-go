package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rohitgajbhiye2005/stock-api/models"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading the .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Println("succesfully connected to postgres")
	return db
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("Unable to decode the request %v", err)
	}
	insertId := insertStock(stock)
	res := response{
		ID:      insertId,
		Message: "stock created succesfully",
	}
	json.NewEncoder(w).Encode(res)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("unable to convert the string into the int %v", err)
	}

	stock, err := getStock(int64(id))

	if err != nil {
		log.Fatalf("unable to get stock %v", err)
	}

	json.NewEncoder(w).Encode(stock)

}

func GetAllStock(w http.ResponseWriter, r *http.Request) {
	stock, err := getAllStock()

	if err != nil {
		log.Fatalf("unable to get all stock %v", err)
	}

	json.NewEncoder(w).Encode(stock)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("unable to convert into the int %v", err)
	}
	var stock models.Stock

	err = json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("unable to decod ethe request %v", err)
	}
	updatedRows := updateStock(int64(id), stock)
	msg := fmt.Sprintf("stock updated succesfully and Total rows affected %v", updatedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)

}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatalf("unable to convert into the int %v", err)
	}

	deletedRows := deleteStock(int64(id))
	msg := fmt.Sprintf("stock is deleted succesfully total affected is %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)

}

func insertStock(stock models.Stock)int64{
	db:=createConnection()
	defer db.Close()
	sqlStatement:=`INSERT INTO stocks(name,price,company) VALUES($1,$2,$3) RETURNING stockid`

	var id int64

	err:=db.QueryRow(sqlStatement,stock.Name,stock.Price,stock.Company).Scan(&id)

	if err!=nil{
		log.Fatalf("unable to execute the query %v",err)
	}
	fmt.Printf("Inserted a single row %v",id)
	return id
}

func getStock(id int64)(models.Stock,error){
	db:=createConnection()
	defer db.Close()

	var stock models.Stock

	sqlStatement:=`SELECT * FROM stocks WHERE stockid=$1`

	row:=db.QueryRow(sqlStatement,id)

	err:=row.Scan(&stock.StockId,&stock.Name,&stock.Price,&stock.Company)

	switch err{
	case sql.ErrNoRows:
		fmt.Println("now rows are returned!")
		return stock,nil
	case nil:
		return stock,nil
	default:
		log.Fatalf("unable to scan the row %v",err)
	}
	return stock,err
}

func getAllStock()([]models.Stock,error){
	db:=createConnection()
	defer db.Close()

	var stocks[] models.Stock

	sqlStatement:=`SELECT * FROM stocks`

	rows,err:=db.Query(sqlStatement)
	if err!=nil{
		log.Fatalf("unable to execute the query %v",err)
	}

	defer db.Close()

	for rows.Next(){
		var stock models.Stock
		err:=rows.Scan(&stock.StockId,&stock.Name,&stock.Price,&stock.Company)
		if err!=nil{
			log.Fatalf("unable to scan the row %v",err)
		}
		stocks=append(stocks,stock)
	}
	return stocks,err

}

func updateStock(id int64,stock models.Stock) int64{
	db := createConnection()

	
	defer db.Close()

	
	sqlStatement := `UPDATE stocks SET name=$2, price=$3, company=$4 WHERE stockid=$1`

	
	res, err := db.Exec(sqlStatement, id, stock.Name, stock.Price, stock.Company)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}


func deleteStock(id int64)int64{
	db := createConnection()

	
	defer db.Close()

	
	sqlStatement := `DELETE FROM stocks WHERE stockid=$1`

	
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
