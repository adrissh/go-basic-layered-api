package models

type Employee struct {
	ID       string  `json:"id"`       // ID unik karyawan (contoh: "EMP001")
	Name     string  `json:"name"`     // Nama karyawan (contoh: "John Doe")
	Age      int     `json:"age"`      // Usia karyawan (contoh: 30)
	Position string  `json:"position"` // Posisi jabatan karyawan (contoh: "Software Engineer")
	Salary   float64 `json:"salary"`   // Gaji karyawan (contoh: 75000.50)
	HireDate string  `json:"hireDate"` // Tanggal mulai bekerja (contoh: "2022-04-15")
}
