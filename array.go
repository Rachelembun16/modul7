package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Increment, used to avoid doubled customer id
var Increment int = 0

type Customer struct {
	Id    int
	Name  string
	Hours int
	Price int
}

var ListCustomers []Customer

type CustomerRepository interface {
	Add(name string, hours int) Customer
	Delete(id int)
	GetAll() ([]Customer, int)
	GetAverageHours() int
	GetMinHour() []Customer
	GetMinAverageUsage() []Customer
}

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepositoryImpl{}
}

func (repo *CustomerRepositoryImpl) Add(name string, hours int) Customer {
	Increment += 1
	customer := Customer{
		Id:    Increment,
		Name:  name,
		Hours: hours,
		Price: hours * 60 * 1000,
	}

	ListCustomers = append(ListCustomers, customer)

	return customer
}

func (repo *CustomerRepositoryImpl) Delete(id int) {
	for index, customer := range ListCustomers {
		if customer.Id == id {
			ListCustomers = append(ListCustomers[:index], ListCustomers[index+1:]...)
		}
	}

}

func (repo *CustomerRepositoryImpl) GetAll() ([]Customer, int) {
	// Sort by customer id before return the result
	sort.Slice(ListCustomers, func(i, j int) bool {
		return ListCustomers[i].Id < ListCustomers[j].Id
	})

	return ListCustomers, len(ListCustomers)
}

func (repo *CustomerRepositoryImpl) GetAverageHours() int {
	counter := 0
	for _, customer := range ListCustomers {
		counter += customer.Hours
	}
	return counter / len(ListCustomers)
}

func (repo *CustomerRepositoryImpl) GetMinHour() []Customer {
	var customers []Customer

	// Sort by Minimum Hours usage
	sort.Slice(ListCustomers, func(i, j int) bool {
		return ListCustomers[i].Hours < ListCustomers[j].Hours
	})
	if len(ListCustomers) >= 3 {
		for i := 0; i < 3; i++ {
			customers = append(customers, ListCustomers[i])
		}
	} else {
		for i := 0; i < len(ListCustomers); i++ {
			customers = append(customers, ListCustomers[i])
		}
	}

	// Append to new customers list

	return customers
}

func (repo *CustomerRepositoryImpl) GetMinAverageUsage() []Customer {
	var customers []Customer

	AverageUsageHours := repo.GetAverageHours()

	// Append customers who under minimum average usage to new customers list
	for _, customer := range ListCustomers {
		if customer.Hours < AverageUsageHours {
			customers = append(customers, customer)
		}

	}

	return customers
}

type CustomerView interface {
	MainMenu()
	AddNewCustomerMenu()
	DeleteCustomerMenu()
	AvarageUsageMenu()
	MinimumUsageMenu()
	MinAverageUsageMenu()
	ShowAllCustomer()
}

// CustomerViewImpl, adalah implementasi (mirip) inheritance dari CustomerRepository
// Contoh penggunaanya dapat dilihat di methode AddNewCustomerMenu()
type CustomerViewImpl struct {
	Repo CustomerRepository
}

func NewCustomerView(customerRepository CustomerRepository) CustomerView {
	return &CustomerViewImpl{Repo: customerRepository}
}

func (view *CustomerViewImpl) MainMenu() {
	fmt.Println("\n===MENU PENGOLAH DATA WARNET===")
	fmt.Println("1. Memasukkan Data")
	fmt.Println("2. Menghapus Data")
	fmt.Println("3. Menampilkan Keseluruhan Data")
	fmt.Println("4. Menampilkan Rata-Rata jumlah jam penggunaan")
	fmt.Println("5. Menampilkan 3 buah data dengan jam penggunaan paling sedikit")
	fmt.Println("6. Menampilkan data costumer dengan jumlah penyewaan komputer dibawah rata rata")
	fmt.Println("0. Keluar")
	fmt.Print("Masukkan Pilihan Anda \t\t: ")
}

func (view *CustomerViewImpl) AddNewCustomerMenu() {
	var name string
	var nama string
	var waktu int

	fmt.Println("===Tambah data Pelanggan WARNET===")
	fmt.Print("Masukkan Nama Pelanggan \t\t: ")
	scan := bufio.NewReader(os.Stdin)
	name, _ = scan.ReadString('\n')
	nama = strings.TrimSpace(name)

	fmt.Print("Waktu Penyewaan Komputer (jam) \t\t: ")
	fmt.Scanln(&waktu)
	// Untuk pemanggilan methode CustomRepository, dapat dilakukan dengan cara seperti ini
	// Mungkin mirip dengan inheritance
	view.Repo.Add(nama, waktu)
	fmt.Println("Berhasil menambahkan data pelanggan!")
	view.ShowAllCustomer()
}

func (view *CustomerViewImpl) DeleteCustomerMenu() {
	view.ShowAllCustomer()
	var customerID int
	fmt.Println("===Hapus data Pelanggan WARNET===")

	fmt.Print("Masukkan ID Pelanggan \t\t: ")
	fmt.Scanln(&customerID)
	view.Repo.Delete(customerID)
}

func (view *CustomerViewImpl) AvarageUsageMenu() {
	avarageUsage := view.Repo.GetAverageHours()
	fmt.Printf("Rata-rata Jam Penggunaan: %d", avarageUsage)
}

func (view *CustomerViewImpl) MinimumUsageMenu() {
	customers := view.Repo.GetMinHour()
	fmt.Println("=== Data Pelanggan Penggunaan Paling sedikit ===")

	fmt.Println("================================================================")
	fmt.Printf("Nomor\t Id\t Nama\t\t Penggunaan\t\t Tagihan\n")
	for index, customer := range customers {
		fmt.Printf(
			"%d\t %d\t %s\t\t %d\t\t\t %d\n",
			index+1, customer.Id, customer.Name, customer.Hours, customer.Price,
		)
	}
	fmt.Println("================================================================")

}

func (view *CustomerViewImpl) MinAverageUsageMenu() {
	customers := view.Repo.GetMinAverageUsage()
	fmt.Println("=== Data Pelanggan Penggunaan dibawah Rerata ===")

	fmt.Println("================================================================")
	fmt.Printf("Nomor\t Id\t Nama\t\t Penggunaan\t\t Tagihan\n")
	for index, customer := range customers {
		fmt.Printf(
			"%d\t %d\t %s\t\t %d\t\t\t %d\n",
			index+1, customer.Id, customer.Name, customer.Hours, customer.Price,
		)
	}
	fmt.Println("================================================================")

}

// ShowAllCustomer, digunakan hanya untuk menampilkan data customer
func (view *CustomerViewImpl) ShowAllCustomer() {
	customers, total := view.Repo.GetAll()

	fmt.Println("================================================================")
	fmt.Printf("Nomor\t Id\t Nama\t\t Penggunaan\t\t Tagihan\n")
	for index, customer := range customers {
		fmt.Printf(
			"%d\t %d\t %s\t\t %d\t\t\t %d\n",
			index+1, customer.Id, customer.Name, customer.Hours, customer.Price,
		)
	}
	fmt.Printf("\nTotal Pelanggan: %d\n", total)
	fmt.Println("================================================================")

}

func main() {
	var pilih int
	customerRepository := NewCustomerRepository()
	customerView := NewCustomerView(customerRepository)

	for {
		customerView.MainMenu()
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			customerView.AddNewCustomerMenu()
		case 2:
			customerView.DeleteCustomerMenu()
		case 3:
			customerView.ShowAllCustomer()
		case 4:
			customerView.AvarageUsageMenu()
		case 5:
			customerView.MinimumUsageMenu()
		case 6:
			customerView.MinAverageUsageMenu()
		case 0:
			fmt.Println("Program berakhir...")
			os.Exit(0)
		}
	}

}
