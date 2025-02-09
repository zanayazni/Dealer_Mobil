package main

import "fmt"

const NMAX int = 500

type pabrik struct {
	nama     string //nama mobil
	pabrik string //nama pabrik
	warna    string //warna mobil
	tahun    int //tahun 
	kode     int //kode mobil selalu berbeda tiap input untuk mempermudah pencarian data
	jmlhPr   int //jumlah penjualan
}

type arrpabrik [NMAX]pabrik

func main() {
	var arr arrpabrik
	var i int
	var opsi, Pilih,PilihSearch int
	Header()
	fmt.Scan(&opsi)
	for opsi != 7 {
		if opsi == 1 {
			i = inputData(&arr)
			fmt.Println("====== Masukan Selesai =======")
		} else if opsi == 2 {
			header2()
			fmt.Scan(&Pilih)
			for Pilih != 5 {
				if Pilih == 1 {
					printData(arr, i)
				} else if Pilih == 2 {
					sortingInsertion(&arr,i)
					printData(arr, i)
				} else if Pilih == 3 {
					sortingSelection(&arr,i)
					printData(arr,i)
					fmt.Println("====== Masukan Selesai =======")
				} else if Pilih == 4 {
					sortingInsertion(&arr,i)
					for b:=0;b<3;b++{
						fmt.Printf("Nama pabrik: %v\t Nama mobil: %v\t Warna mobil: %v\t Tahun: %d\t Kode mobil: %d\t jumlah terjual: %d\n", arr[b].pabrik, arr[b].nama, arr[b].warna, arr[b].tahun, arr[b].kode,arr[b].jmlhPr)
					}
					fmt.Println("====== Masukan Selesai =======")
				} else {
					fmt.Println("Silahkan Pilih Opsi yang tersedia")
				}
				header2()
				fmt.Scan(&Pilih)
			}
		} else if opsi == 3 {
			searchData(arr, i)
		} else if opsi == 4 {
			editData(&arr, i)
		} else if opsi == 5 {
			i = inputDataAgain(&arr, i)
			fmt.Println("====== Masukan Sudah Terupdate =======")
		} else if opsi == 6 {
			header6()
			fmt.Scan(&PilihSearch)
			fmt.Println("")
			for PilihSearch != 3 {
				if PilihSearch == 1 {
					deleteDataBinary(&arr,&i)
				} else if PilihSearch == 2 {
					deleteDataSeq(&arr, &i)
				} else {
					fmt.Println("Pilihan Search ", PilihSearch," Tidak tersedia")
				}
				header6()
				fmt.Scan(&PilihSearch)
				fmt.Println("")
			}
		} else {
			fmt.Println("Opsi Tidak Ada")
		}
		tampilan()
		fmt.Scan(&opsi)
	}
	fmt.Println("====== Program Selesai =========")
}

func inputData(arr *arrpabrik) int {
	var pabrik, nama, warna string
	var kode, tahun, i int
	var jmlhPr int

	fmt.Println("Masukan Data")
	fmt.Println("==================")
	fmt.Print("Pabrik : ")
	fmt.Scan(&pabrik)
	fmt.Print("Nama Mobil : ")
	fmt.Scan(&nama)
	fmt.Print("Warna : ") 
	fmt.Scan(&warna)
	fmt.Print("Tahun: ")
	fmt.Scan(&tahun)
	fmt.Print("kode : ")
	fmt.Scan(&kode)
	fmt.Print("jumlah produk yang terjual : ")
	fmt.Scan(&jmlhPr)
	fmt.Println("")

	for pabrik != "None" && nama != "None" && warna != "None" && tahun != 0 && kode != 0  && jmlhPr != 0{
		arr[i].pabrik = pabrik
		arr[i].nama = nama
		arr[i].warna = warna
		arr[i].tahun = tahun
		arr[i].kode = kode
		arr[i].jmlhPr = jmlhPr
		i++
		fmt.Print("Pabrik : ")
		fmt.Scan(&pabrik)
		fmt.Print("Nama Mobil : ")
		fmt.Scan(&nama)
		fmt.Print("Warna : ") 
		fmt.Scan(&warna)
		fmt.Print("Tahun: ")
		fmt.Scan(&tahun)
		fmt.Print("kode : ")
		fmt.Scan(&kode)
		fmt.Print("jumlah produk yang terjual : ")
		fmt.Scan(&jmlhPr)
		fmt.Println("")
	}
	return i
}
func tampilan() {
	fmt.Println("")
	fmt.Println("Terdapat opsi dibawah ini")
	fmt.Println("1. Input Data")
	fmt.Println("2. Tampilkan Data")
	fmt.Println("3. Search Data")
	fmt.Println("4. Forum Edit Data")
	fmt.Println("5. Input Data Kembali")
	fmt.Println("6. Delete Data")
	fmt.Println("7. Quit")
	fmt.Println("Masukan Opsi :")
}
func Header() {
	fmt.Println("Selamat Datang di Aplikasi Management Mobil:")
	fmt.Println("")
	fmt.Println("Terdapat opsi dibawah ini")
	fmt.Println("1. Input Data")
	fmt.Println("2. Tampilkan Data")
	fmt.Println("3. Search Data")
	fmt.Println("4. Forum Edit Data")
	fmt.Println("5. Input Data Kembali")
	fmt.Println("6. Delete Data")
	fmt.Println("7. Quit")
	fmt.Println("Masukan Opsi dibawah ini:")
}

func header2(){
	fmt.Println("Silahkan pilih Tampilan data yang dinginkan")
	fmt.Println("1. Tampilkan Seluruh Data")
	fmt.Println("2. Tampilkan Berdasarkan Jumlah Produk Mobil")
	fmt.Println("3. Tampilkan Berdasarkan Tahun Keluar")
	fmt.Println("4. Tampilkan 3 daftar mobil dengan penjualan tertinggi")
	fmt.Println("5. Quit")
	fmt.Println("Silakan pilih opsi yang diinginkan : ")
}

func header6(){
	fmt.Println("Silahkan pilih tipe searching yang ingin digunakan")
	fmt.Println("1. Binary Search")
	fmt.Println("2. Sequential Search")
	fmt.Println("3. Quit")
	fmt.Println("Silakan pilih opsi yang diinginkan : ")

}

func inputDataAgain(arr *arrpabrik, startIdx int) int {
	var pabrik, nama, warna string
	var kode, tahun, i,jmlhPr int

	fmt.Println("Masukan Data")
	fmt.Println("==================")
	fmt.Print("Pabrik : ")
	fmt.Scan(&pabrik)
	fmt.Print("Nama Mobil : ")
	fmt.Scan(&nama)
	fmt.Print("Warna : ") 
	fmt.Scan(&warna)
	fmt.Print("Tahun: ")
	fmt.Scan(&tahun)
	fmt.Print("kode : ")
	fmt.Scan(&kode)
	fmt.Print("jumlah produk yang terjual : ")
	fmt.Scan(&jmlhPr)
	fmt.Println("")

	i = startIdx

	for pabrik != "None" && nama != "None" && warna != "None" && tahun != 0 && kode != 0  && jmlhPr != 0{
		arr[i].pabrik = pabrik
		arr[i].nama = nama
		arr[i].warna = warna
		arr[i].tahun = tahun
		arr[i].kode = kode
		arr[i].jmlhPr = jmlhPr
		i++
		fmt.Print("Pabrik : ")
		fmt.Scan(&pabrik)
		fmt.Print("Nama Mobil : ")
		fmt.Scan(&nama)
		fmt.Print("Warna : ") 
		fmt.Scan(&warna)
		fmt.Print("Tahun: ")
		fmt.Scan(&tahun)
		fmt.Print("kode : ")
		fmt.Scan(&kode)
		fmt.Print("jumlah produk yang terjual : ")
		fmt.Scan(&jmlhPr)
		fmt.Println("")
	}
	return i
}

func printData(arr arrpabrik, size int) {
	fmt.Println("================")
	fmt.Println("Data dalam array")
	fmt.Println("================")

	for b := 0; b < size; b++ {
		fmt.Printf("Nama pabrik: %v\t Nama mobil: %v\t Warna mobil: %v\t Tahun: %d\t Kode mobil: %d\t jumlah: %d\n", arr[b].pabrik, arr[b].nama, arr[b].warna, arr[b].tahun, arr[b].kode,arr[b].jmlhPr)
	}
	fmt.Println("")
}

func searchData(arr arrpabrik, size int) {
	var cari, pabrik, nama, warna string
	var kode, tahun int

	fmt.Println("================")
	fmt.Println("Opsi yang tersedia:")
	fmt.Println("pabrik")
	fmt.Println("nama")
	fmt.Println("warna")
	fmt.Println("tahun")
	fmt.Println("kode")
	fmt.Println("================")
	fmt.Println("Masukan opsi yang ingin dicari:")
	fmt.Scan(&cari)

	if cari == "pabrik" {
		fmt.Println("Masukan Nama pabriknya")
		fmt.Scan(&pabrik)
		fmt.Println("======  Hasil  =======")
		for c := 0; c < size; c++ {
			if arr[c].pabrik == pabrik {
				fmt.Println("Nama Mobil tersebut:", arr[c].nama, "Dengan warna:", arr[c].warna, "Tahun:", arr[c].tahun, "Kode:", arr[c].kode)
			}
		}
	} else if cari == "nama" {
		fmt.Println("Masukan Nama Mobilnya")
		fmt.Scan(&nama)
		fmt.Println("======Hasil=======")
		for c := 0; c < size; c++ {
			if arr[c].nama == nama {
				fmt.Println("pabrik Mobil tersebut:", arr[c].pabrik, "Dengan warna:", arr[c].warna, "Tahun:", arr[c].tahun, "Kode:", arr[c].kode)
			}
		}
	} else if cari == "warna" {
		fmt.Println("Masukan Warna Mobilnya")
		fmt.Scan(&warna)
		fmt.Println("======Hasil=======")
		for c := 0; c < size; c++ {
			if arr[c].warna == warna {
				fmt.Println("pabrik Mobil tersebut:", arr[c].pabrik, "Dengan Nama:", arr[c].nama, "Tahun:", arr[c].tahun, "Kode:", arr[c].kode)
			}
		}
	} else if cari == "tahun" {
		fmt.Println("Masukan Tahun Keluaran Mobilnya")
		fmt.Scan(&tahun)
		fmt.Println("======Hasil=======")
		for c := 0; c < size; c++ {
			if arr[c].tahun == tahun {
				fmt.Println("pabrik Mobil tersebut:", arr[c].pabrik, "Dengan Nama:", arr[c].nama, "Warna Mobil Tersebut:", arr[c].warna, "Kode:", arr[c].kode)
			}
		}
	} else if cari == "kode" {
		fmt.Println("Masukan Kode Mobilnya")
		fmt.Scan(&kode)
		fmt.Println("======Hasil=======")
		for c := 0; c < size; c++ {
			if arr[c].kode == kode {
				fmt.Println("pabrik Mobil tersebut:", arr[c].pabrik, "Dengan Nama:", arr[c].nama, "Tahun:", arr[c].tahun, "Warna Mobil Tersebut:", arr[c].warna)
			}
		}
	}
	fmt.Println("")
}

func editData(arr *arrpabrik, size int) {
	var kode, tahun, j, jmlhPr int
	var cari, pabrik, nama, warna string

	fmt.Println("=================")
	fmt.Println("Forum Pengeditan")
	fmt.Println("=================")

	fmt.Println("Masukan Kode Mobil:")
	fmt.Scan(&kode)
	fmt.Println("")

	for j = 0; j < size; j++ {
		if arr[j].kode == kode {
			fmt.Println("Masukan nama variabel yang ingin diubah:")
			fmt.Scan(&cari)
			if cari == "pabrik" {
				fmt.Println("Masukan Perubahannya:")
				fmt.Scan(&pabrik)
				arr[j].pabrik = pabrik
			} else if cari == "nama" {
				fmt.Println("Masukan Perubahannya:")
				fmt.Scan(&nama)
				arr[j].nama = nama
			} else if cari == "warna" {
				fmt.Println("Masukan Perubahannya:")
				fmt.Scan(&warna)
				arr[j].warna = warna
			} else if cari == "tahun" {
				fmt.Println("Masukan Perubahannya:")
				fmt.Scan(&tahun)
				arr[j].tahun = tahun
			} else if cari == "terjual" {
				fmt.Println("Masukan Perubahannya:")
				fmt.Scan(&jmlhPr)
				arr[j].jmlhPr = jmlhPr
			} else {
				fmt.Println("Pilihan kategori tidak tersedia")
			}
		}
	}
	fmt.Println("")
	fmt.Println("=========== HASIL SUDAH EDIT =============")
	fmt.Println("")
	// data array kembali dipanggil
	for j := 0; j < size; j++ {
		fmt.Printf("Nama pabrik: %v\t Nama mobil: %v\t Warna mobil: %v\t Tahun: %d\t Kode mobil: %d\t jumlah terjual: %d\n", arr[j].pabrik, arr[j].nama, arr[j].warna, arr[j].tahun, arr[j].kode, arr[j].jmlhPr)
	}
	fmt.Println("")
}

func seqSeacrh(arr arrpabrik, size, kode int) int {
	var i int
	for i < size && kode != arr[i].kode {
		i++
	}
	if i == size {
		return -1
	} else {
		return i
	}
}

func sortingInsertion(arr *arrpabrik,size int){
	//melakukan sorting berdasarkan jumlah mobil 
	var pass, i int
	var merah pabrik
	for pass = 1 ; pass < size ; pass++{
		i = pass
		merah = arr[i]
		for i > 0 && merah.jmlhPr > arr[i-1].jmlhPr {
			arr[i] = arr[i-1]
			i--
		}
		arr[i] = merah
	}
}

func sortingInsertion_kode(arr *arrpabrik,size int){
	//melakukan sorting berdasarkan kode mobil descending
	var pass, i int
	var merah pabrik
	for pass = 1 ; pass < size ; pass++{
		i = pass
		merah = arr[i]
		for i > 0 && merah.kode > arr[i-1].kode {
			arr[i] = arr[i-1]
			i--
		}
		arr[i] = merah
	}
	printData(*arr,size)
}

func sortingSelection(arr *arrpabrik, size int){
	//melakukan soring berdasarkan tahun keluar ascending
	var i,j,maxIndx int
	var temp pabrik
	for i = 0; i< size-1;i++{
		maxIndx = i
		for j=i+1;j<size;j++{
			if arr[j].tahun < arr[maxIndx].tahun{
				maxIndx = j
			}
		}
		temp = arr[maxIndx]
		arr[maxIndx] = arr[i]
		arr[i] = temp
	}

}

func binarySearch(arr arrpabrik, size, kode int) int {
	var left,right,mid,found int
	left = 0
	right = size-1
	found = -1

	for left <= right && found == -1{
		mid = (left + right) / 2
		if kode < arr[mid].kode {
			left = mid + 1
		} else if  kode > arr[mid].kode {
			right = mid -1 
		} else {
			found = mid
		}
	}
	return found
}

func deleteDataBinary(arr *arrpabrik, size *int) {
	var kode, j, idx int
	sortingInsertion_kode(arr,*size)
	fmt.Println("=================")
	fmt.Println("Hapus Data")
	fmt.Println("=================")

	fmt.Println("Masukan Kode Mobil:")
	fmt.Scan(&kode)
	fmt.Println("")
	idx = binarySearch(*arr,*size,kode)

	if idx == -1 {
		fmt.Println("Data Tidak Ditemukan")
	} else {
		for j = idx; j < *size-1; j++ {
			arr[j] = arr[j+1]
		}
		*size--
		fmt.Println("Data berhasil dihapus")
	}
	printData(*arr,*size)
}

func deleteDataSeq(arr *arrpabrik, size *int) {
	var kode, j, idx int

	printData(*arr,*size)

	fmt.Println("=================")
	fmt.Println("Hapus Data")
	fmt.Println("=================")

	fmt.Println("Masukan Kode Mobil:")
	fmt.Scan(&kode)
	fmt.Println("")

	idx = seqSeacrh(*arr, *size, kode)

	if idx == -1 {
		fmt.Println("Data Tidak Ditemukan")
		fmt.Println("")
	} else {
		for j = idx; j < *size-1; j++ {
			arr[j] = arr[j+1]
		}
		*size--
		fmt.Println("Data berhasil dihapus")
		fmt.Println("")
	}
	printData(*arr,*size)
}

