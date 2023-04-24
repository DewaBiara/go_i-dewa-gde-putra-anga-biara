package main

type kendaraan struct {
	// totalroda       int    <- tidak menggunakan camelCase pada nama variable
	totalRoda int
	// kecepatanperjam int 	  <- tidak menggunakan camelCase pada nama variable
	kecepatanPerJam int
}

type mobil struct {
	kendaraan
}

func (m *mobil) berjalan() {
	// m.tambahkecepatan(10)  <- mengganti sesuai dengan nama method yang ada di dalam struct
	m.tambahKecepatan(10)
}

// func (m *mobil) tambahkecepatan(kecepatanbaru int) { 	 <- tidak menggunakan camelCase pada nama function dan argumennya
func (m *mobil) tambahKecepatan(kecepatanBaru int) {
	// m.kecepatanperjam = m.kecepatanperjam + kecepatanbaru <- mengganti sesuai dengan nama variable yang ada di dalam struct
	m.kecepatanPerJam = m.kecepatanPerJam + kecepatanBaru
}

func main() {
	// mobilcepat := mobil{}	<- tidak menggunakan camelCase untuk nama variable
	mobilCepat := mobil{}
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()

	// mobillamban := mobil{}	<- tidak menggunakan camelCase untuk nama variable
	mobilLamban := mobil{}
	mobilLamban.berjalan()
}
