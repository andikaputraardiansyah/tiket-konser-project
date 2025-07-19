package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Tiket struct {
	ID       string
	Nama     string
	Kategori string
	Harga    int
	Status   string // "pending", "confirmed", "rejected"
	Tanggal  string
}

type Akun struct {
	Username string
	Password string
	Role     string
}

var daftarTiket []Tiket
var daftarAkun []Akun
var loginAkun *Akun = nil
var reader = bufio.NewReader(os.Stdin)
var lastID = 0

// ==================== KONSTANTA WARNA DAN STYLE ====================
const (
	// Warna dasar
	RESET     = "\033[0m"
	BOLD      = "\033[1m"
	DIM       = "\033[2m"
	UNDERLINE = "\033[4m"
	BLINK     = "\033[5m"

	// Warna teks
	BLACK   = "\033[30m"
	RED     = "\033[31m"
	GREEN   = "\033[32m"
	YELLOW  = "\033[33m"
	BLUE    = "\033[34m"
	MAGENTA = "\033[35m"
	CYAN    = "\033[36m"
	WHITE   = "\033[37m"

	// Warna background
	BG_BLACK   = "\033[40m"
	BG_RED     = "\033[41m"
	BG_GREEN   = "\033[42m"
	BG_YELLOW  = "\033[43m"
	BG_BLUE    = "\033[44m"
	BG_MAGENTA = "\033[45m"
	BG_CYAN    = "\033[46m"
	BG_WHITE   = "\033[47m"

	// Kombinasi style
	SUCCESS   = BOLD + GREEN
	ERROR     = BOLD + RED
	WARNING   = BOLD + YELLOW
	INFO      = BOLD + CYAN
	PRIMARY   = BOLD + BLUE
	SECONDARY = BOLD + MAGENTA
)

func main() {
	clearScreen()
	tampilkanSplashScreen()
	muatDataAkun()
	muatDataTiket()
	for {
		menuLogin()
	}
}

// ==================== FUNGSI UTILITAS TAMPILAN ====================

func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

func printBorder(char string, length int, color string) {
	fmt.Print(color)
	for i := 0; i < length; i++ {
		fmt.Print(char)
	}
	fmt.Print(RESET + "\n")
}

func printCentered(text string, width int, color string) {
	padding := (width - len(text)) / 2
	fmt.Print(color)
	fmt.Print(strings.Repeat(" ", padding) + text)
	fmt.Print(RESET + "\n")
}

func printBox(title string, content []string, width int, borderColor string, titleColor string) {
	// Top border
	fmt.Print(borderColor + "╔")
	for i := 0; i < width-2; i++ {
		fmt.Print("═")
	}
	fmt.Print("╗" + RESET + "\n")

	// Title
	if title != "" {
		titlePadding := (width - len(title) - 2) / 2
		fmt.Print(borderColor + "║" + RESET)
		fmt.Print(strings.Repeat(" ", titlePadding))
		fmt.Print(titleColor + title + RESET)
		fmt.Print(strings.Repeat(" ", width-len(title)-titlePadding-2))
		fmt.Print(borderColor + "║" + RESET + "\n")

		// Separator
		fmt.Print(borderColor + "╠")
		for i := 0; i < width-2; i++ {
			fmt.Print("═")
		}
		fmt.Print("╣" + RESET + "\n")
	}

	// Content
	for _, line := range content {
		fmt.Print(borderColor + "║" + RESET)
		fmt.Print(" " + line)
		fmt.Print(strings.Repeat(" ", width-len(line)-3))
		fmt.Print(borderColor + "║" + RESET + "\n")
	}

	// Bottom border
	fmt.Print(borderColor + "╚")
	for i := 0; i < width-2; i++ {
		fmt.Print("═")
	}
	fmt.Print("╝" + RESET + "\n")
}

func tampilkanSplashScreen() {
	clearScreen()

	// Gradient background effect
	fmt.Print(BG_CYAN + strings.Repeat(" ", 80) + RESET + "\n")

	// SpongeBob ASCII Art dengan warna
	fmt.Print(YELLOW + BOLD)
	fmt.Println("                    .-\"\"\"\"-.")
	fmt.Println("                   /        \\")
	fmt.Println("                  /_        _\\")
	fmt.Println("                 // \\      / \\\\")
	fmt.Println("                 |\\__\\    /__/|")
	fmt.Println("                  \\    ||    /")
	fmt.Println("                   \\___||___/")
	fmt.Println("                   /   ><   \\")
	fmt.Println("                  /  ______  \\")
	fmt.Println("                 /  /      \\  \\")
	fmt.Println("                (  (   __   )  )")
	fmt.Println("                 \\  \\______/  /")
	fmt.Println("                  \\__________/")
	fmt.Println("                     |    |")
	fmt.Println("                     |    |")
	fmt.Println("                    /      \\")
	fmt.Println("                   /________\\")
	fmt.Print(RESET)

	// Animated title
	title := []string{
		"🎵 SISTEM TIKET KONSER RT07 RW09 🎵",
		"",
		"✨ Powered by SpongeBob SquarePants ✨",
		"🌊 Bikini Bottom Entertainment System 🌊",
	}

	printBox("", title, 60, CYAN, BOLD+WHITE)

	// Loading animation
	fmt.Print(INFO + "Memuat sistem")
	for i := 0; i < 5; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Print(RESET + "\n\n")

	// Press any key
	fmt.Print(DIM + "Tekan Enter untuk melanjutkan..." + RESET)
	reader.ReadString('\n')
}

func tampilkanTerimaKasihSpongeBob() {
	clearScreen()

	// Animated SpongeBob
	fmt.Print(YELLOW + BOLD)
	fmt.Println("                 🌟✨ TERIMA KASIH! ✨🌟")
	fmt.Println()
	fmt.Println("                    .-\"\"\"\"-.")
	fmt.Println("                   /   😊   \\")
	fmt.Println("                  /_        _\\")
	fmt.Println("                 // \\  👀  / \\\\")
	fmt.Println("                 |\\__\\    /__/|")
	fmt.Println("                  \\    ^^    /")
	fmt.Println("                   \\___||___/")
	fmt.Println("                   /   ><   \\")
	fmt.Println("                  /  ______  \\")
	fmt.Println("                 /  /  😄  \\  \\")
	fmt.Println("                (  (   __   )  )")
	fmt.Println("                 \\  \\______/  /")
	fmt.Println("                  \\__________/")
	fmt.Println("                     |    |")
	fmt.Println("                     |    |")
	fmt.Println("                    /  🦵  \\")
	fmt.Println("                   /________\\")
	fmt.Print(RESET)

	ucapanUnik := []string{
		"🎪 Wah! Konser ini pasti akan lebih seru dari Krusty Krab!",
		"🎵 Gary akan iri dengan tiket keren ini! Meong~",
		"🍔 Tiket ini lebih berharga dari resep rahasia Krabby Patty!",
		"⭐ Patrick pasti akan bilang: 'Ini dia yang namanya ROCK!'",
		"🌊 Seperti gelembung sabun, kebahagiaan Anda akan melayang tinggi!",
		"🏠 Nanas saya tidak sebagus konser yang akan Anda tonton!",
		"🦀 Mr. Krabs akan bangga dengan investasi tiket yang cerdas ini!",
		"🐙 Bahkan Squidward akan tersenyum melihat pertunjukan ini!",
	}

	ucapan := ucapanUnik[time.Now().Second()%len(ucapanUnik)]

	farewell := []string{
		ucapan,
		"",
		"🎊 Sampai jumpa di konser! Jangan lupa bawa semangat! 🎊",
		"",
		"- Dengan cinta dari Bikini Bottom 💙 -",
	}

	printBox("FAREWELL MESSAGE", farewell, 70, CYAN, BOLD+WHITE)
}

func tampilkanSpongeBobKecil() {
	fmt.Print(YELLOW)
	fmt.Println("      .-\"\"\"\"-.    ")
	fmt.Println("     /  😊   \\   ")
	fmt.Println("    /_      _\\   ")
	fmt.Println("   // \\👀 / \\\\  ")
	fmt.Println("   |\\__\\  /__/|  ")
	fmt.Println("    \\   ^^   /   ")
	fmt.Println("     \\__||__/    ")
	fmt.Println("     /  ><  \\    ")
	fmt.Println("    /______\\     ")
	fmt.Println("   (   😄   )    ")
	fmt.Println("    \\______/     ")
	fmt.Println("      |  |       ")
	fmt.Println("     /____\\      ")
	fmt.Print(RESET)
}

// ==================== FUNGSI VALIDASI INPUT ====================

func validasiInputInteger(prompt string, min, max int) int {
	for {
		fmt.Print(PRIMARY + prompt + RESET)
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if err != nil {
			fmt.Println(ERROR + "❌ Error membaca input! Coba lagi." + RESET)
			continue
		}

		if input == "" {
			fmt.Println(ERROR + "❌ Input tidak boleh kosong!" + RESET)
			fmt.Println(WARNING + "🧽 SpongeBob berkata: 'Jangan lupa isi angkanya ya!' 🧽" + RESET)
			continue
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(ERROR + "❌ Input harus berupa angka!" + RESET)
			fmt.Println(WARNING + "🧽 SpongeBob berkata: 'Angka saja ya, bukan huruf!' 🧽" + RESET)
			continue
		}

		if num < min || num > max {
			fmt.Printf(ERROR+"❌ Angka harus antara %d dan %d!"+RESET+"\n", min, max)
			fmt.Println(WARNING + "🧽 SpongeBob berkata: 'Pilih yang ada di menu ya!' 🧽" + RESET)
			continue
		}

		return num
	}
}

func validasiInputString(prompt string, minLen, maxLen int, allowEmpty bool) string {
	for {
		fmt.Print(PRIMARY + prompt + RESET)
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if err != nil {
			fmt.Println(ERROR + "❌ Error membaca input! Coba lagi." + RESET)
			continue
		}

		if !allowEmpty && input == "" {
			fmt.Println(ERROR + "❌ Input tidak boleh kosong!" + RESET)
			fmt.Println(WARNING + "🧽 SpongeBob berkata: 'Jangan lupa diisi ya!' 🧽" + RESET)
			continue
		}

		if allowEmpty && input == "" {
			return input
		}

		if len(input) < minLen {
			fmt.Printf(ERROR+"❌ Input minimal %d karakter!"+RESET+"\n", minLen)
			fmt.Println(WARNING + "🧽 SpongeBob berkata: 'Terlalu pendek, tambah lagi!' 🧽" + RESET)
			continue
		}

		if maxLen > 0 && len(input) > maxLen {
			fmt.Printf(ERROR+"❌ Input maksimal %d karakter!"+RESET+"\n", maxLen)
			fmt.Println(WARNING + "🧽 SpongeBob berkata: 'Terlalu panjang, dipersingkat ya!' 🧽" + RESET)
			continue
		}

		return input
	}
}

func validasiUsername(prompt string) string {
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)

	for {
		input := validasiInputString(prompt, 3, 20, false)

		if !usernameRegex.MatchString(input) {
			fmt.Println(ERROR + "❌ Username hanya boleh huruf, angka, dan underscore (3-20 karakter)!" + RESET)
			fmt.Println(WARNING + "🧽 SpongeBob berkata: 'Username seperti nama di Bikini Bottom!' 🧽" + RESET)
			continue
		}

		return input
	}
}

func validasiPassword(prompt string) string {
	for {
		input := validasiInputString(prompt, 6, 50, false)

		hasLetter := regexp.MustCompile(`[a-zA-Z]`).MatchString(input)
		hasNumber := regexp.MustCompile(`[0-9]`).MatchString(input)

		if !hasLetter || !hasNumber {
			fmt.Println(ERROR + "❌ Password harus mengandung huruf dan angka (minimal 6 karakter)!" + RESET)
			fmt.Println(WARNING + "🧽 SpongeBob berkata: 'Password yang kuat seperti spatula!' 🧽" + RESET)
			continue
		}

		return input
	}
}

func validasiIDTiket(prompt string) string {
	idRegex := regexp.MustCompile(`^TKT\d{3}$`)

	for {
		fmt.Print(PRIMARY + prompt + RESET)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToUpper(input))

		if input == "" {
			fmt.Println(ERROR + "❌ ID tiket tidak boleh kosong!" + RESET)
			fmt.Println(WARNING + "🧽 SpongeBob berkata: 'ID tiket seperti TKT001!' 🧽" + RESET)
			continue
		}

		if !idRegex.MatchString(input) {
			fmt.Println(ERROR + "❌ Format ID tiket salah! Contoh: TKT001, TKT002" + RESET)
			fmt.Println(WARNING + "🧽 SpongeBob berkata: 'Format harus TKT diikuti 3 angka!' 🧽" + RESET)
			continue
		}

		return input
	}
}

func validasiStatus(prompt string) string {
	validStatuses := []string{"pending", "confirmed", "rejected"}

	for {
		input := strings.ToLower(validasiInputString(prompt, 1, 20, false))

		for _, status := range validStatuses {
			if input == status {
				return input
			}
		}

		fmt.Println(ERROR + "❌ Status harus: pending, confirmed, atau rejected!" + RESET)
		fmt.Println(WARNING + "🧽 SpongeBob berkata: 'Pilih salah satu status yang tersedia!' 🧽" + RESET)
	}
}

func validasiKategori(prompt string) string {
	validKategori := map[string]int{
		"reguler":    50000,
		"vip":        100000,
		"early bird": 25000,
	}

	for {
		input := strings.ToLower(validasiInputString(prompt, 1, 20, false))

		if _, exists := validKategori[input]; exists {
			return input
		}

		fmt.Println(ERROR + "❌ Kategori harus: Reguler, VIP, atau Early Bird!" + RESET)
		fmt.Println(WARNING + "🧽 SpongeBob berkata: 'Pilih kategori yang tersedia!' 🧽" + RESET)

		categories := []string{
			"Kategori tersedia:",
			"🎫 Reguler (Rp50.000)",
			"⭐ VIP (Rp100.000)",
			"🐦 Early Bird (Rp25.000)",
		}
		printBox("KATEGORI TIKET", categories, 40, BLUE, WHITE)
	}
}

func validasiKonfirmasi(prompt string) bool {
	for {
		fmt.Print(PRIMARY + prompt + RESET)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "" {
			fmt.Println(ERROR + "❌ Silakan jawab y atau n!" + RESET)
			continue
		}

		if input == "y" || input == "yes" {
			return true
		}

		if input == "n" || input == "no" {
			return false
		}

		fmt.Println(ERROR + "❌ Jawab dengan 'y' untuk ya atau 'n' untuk tidak!" + RESET)
		fmt.Println(WARNING + "🧽 SpongeBob berkata: 'Y atau N saja ya!' 🧽" + RESET)
	}
}

func validasiRole(prompt string) string {
	validRoles := []string{"admin", "user"}

	for {
		input := strings.ToLower(validasiInputString(prompt, 1, 10, false))

		for _, role := range validRoles {
			if input == role {
				return input
			}
		}

		fmt.Println(ERROR + "❌ Role harus 'admin' atau 'user'!" + RESET)
		fmt.Println(WARNING + "🧽 SpongeBob berkata: 'Pilih admin atau user!' 🧽" + RESET)
	}
}

// ==================== MENU UTAMA ====================

func menuLogin() {
	clearScreen()

	// Header dengan gradient effect
	printBorder("═", 60, CYAN)
	printCentered("🎵 SISTEM TIKET KONSER RT07 RW09 🎵", 60, BOLD+WHITE+BG_CYAN)
	printBorder("═", 60, CYAN)

	menuItems := []string{
		"",
		"🔐 1. Login",
		"📝 2. Registrasi",
		"🚪 3. Keluar",
		"",
	}

	printBox("MENU UTAMA", menuItems, 50, BLUE, BOLD+WHITE)

	pilihan := validasiInputInteger("🎯 Pilih menu (1-3): ", 1, 3)

	switch pilihan {
	case 1:
		login()
	case 2:
		registrasi()
	case 3:
		tampilkanTerimaKasihSpongeBob()
		time.Sleep(3 * time.Second)
		os.Exit(0)
	}
}

func registrasi() {
	clearScreen()

	header := []string{
		"🧽 SpongeBob akan membantu membuat akun baru! 🧽",
		"",
		"Silakan isi data berikut:",
	}
	printBox("REGISTRASI AKUN", header, 60, GREEN, BOLD+WHITE)

	username := validasiUsername("👤 Username (3-20 karakter, huruf/angka/underscore): ")

	// Cek apakah username sudah ada
	for _, akun := range daftarAkun {
		if akun.Username == username {
			fmt.Println(ERROR + "❌ Username sudah terdaftar!" + RESET)
			fmt.Println(WARNING + "🧽 SpongeBob berkata: 'Coba username lain ya!' 🧽" + RESET)
			time.Sleep(2 * time.Second)
			return
		}
	}

	password := validasiPassword("🔒 Password (minimal 6 karakter, harus ada huruf dan angka): ")

	// Konfirmasi password
	fmt.Print(PRIMARY + "🔒 Konfirmasi password: " + RESET)
	konfirmasiPassword, _ := reader.ReadString('\n')
	konfirmasiPassword = strings.TrimSpace(konfirmasiPassword)

	if password != konfirmasiPassword {
		fmt.Println(ERROR + "❌ Password tidak cocok!" + RESET)
		fmt.Println(WARNING + "🧽 SpongeBob berkata: 'Password harus sama!' 🧽" + RESET)
		time.Sleep(2 * time.Second)
		return
	}

	role := validasiRole("👑 Role (admin/user): ")

	akunBaru := Akun{username, password, role}
	daftarAkun = append(daftarAkun, akunBaru)
	simpanAkunKeFile()

	// Success message
	clearScreen()
	tampilkanSpongeBobKecil()

	success := []string{
		"🎉 Registrasi berhasil! 🎉",
		"",
		"Selamat bergabung di Bikini Bottom!",
		"Silakan login dengan akun baru Anda.",
	}
	printBox("REGISTRASI BERHASIL", success, 50, GREEN, BOLD+WHITE)

	time.Sleep(3 * time.Second)
}

func login() {
	clearScreen()

	header := []string{
		"🧽 SpongeBob siap membantu login! 🧽",
		"",
		"Masukkan kredensial Anda:",
	}
	printBox("LOGIN AKUN", header, 50, BLUE, BOLD+WHITE)

	maxPercobaan := 3
	for i := 0; i < maxPercobaan; i++ {
		fmt.Printf(INFO+"--- Percobaan %d dari %d ---"+RESET+"\n", i+1, maxPercobaan)

		username := validasiInputString("👤 Username: ", 1, 50, false)
		password := validasiInputString("🔒 Password: ", 1, 50, false)

		for _, akun := range daftarAkun {
			if akun.Username == username && akun.Password == password {
				loginAkun = &akun

				// Success animation
				clearScreen()
				fmt.Printf(SUCCESS+"✅ Login berhasil sebagai %s!"+RESET+"\n", akun.Role)
				tampilkanSpongeBobKecil()
				fmt.Println(WARNING + "🌟 Siap-siap untuk petualangan tiket yang seru! 🌟" + RESET)
				loading("Memuat menu")
				menuUtama()
				return
			}
		}

		fmt.Printf(ERROR+"❌ Login gagal! Username atau password salah (%d/%d)"+RESET+"\n", i+1, maxPercobaan)
		fmt.Println(WARNING + "🧽 SpongeBob berkata: 'Coba lagi dengan hati-hati!' 🧽" + RESET)

		if i < maxPercobaan-1 {
			if validasiKonfirmasi("🔄 Coba lagi? (y/n): ") {
				continue
			} else {
				fmt.Println(INFO + "Kembali ke menu utama." + RESET)
				time.Sleep(1 * time.Second)
				return
			}
		}
	}

	fmt.Println(ERROR + "🚫 Terlalu banyak percobaan gagal! Kembali ke menu utama." + RESET)
	fmt.Println(WARNING + "🧽 SpongeBob berkata: 'Istirahat dulu ya, nanti coba lagi!' 🧽" + RESET)
	time.Sleep(3 * time.Second)
}

func loading(pesan string) {
	fmt.Print(INFO + pesan + RESET)
	for i := 0; i < 5; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println()
	time.Sleep(500 * time.Millisecond)
}

func menuUtama() {
	for {
		clearScreen()

		// Header dengan role
		roleColor := PRIMARY
		if loginAkun.Role == "admin" {
			roleColor = ERROR
		}

		fmt.Printf(roleColor+"╔══ Menu Utama (%s) ══╗"+RESET+"\n", strings.ToUpper(loginAkun.Role))

		var menuItems []string
		var maxOption int

		if loginAkun.Role == "admin" {
			menuItems = []string{
				"",
				"📋 1. Tampilkan Tiket",
				"🔍 2. Cari Tiket",
				"📊 3. Urutkan Tiket",
				"✅ 4. Kelola Konfirmasi Tiket",
				"🗑️  5. Hapus Tiket",
				"📈 6. Laporan Excel (Terminal)",
				"🚪 7. Logout",
				"",
			}
			maxOption = 7
		} else {
			menuItems = []string{
				"",
				"📋 1. Tampilkan Tiket",
				"🔍 2. Cari Tiket",
				"📊 3. Urutkan Tiket",
				"🛒 4. Beli Tiket",
				"📄 5. Status Tiket Saya",
				"🚪 6. Logout",
				"",
			}
			maxOption = 6
		}

		printBox("MENU UTAMA", menuItems, 50, roleColor, BOLD+WHITE)

		pilihan := validasiInputInteger(fmt.Sprintf("🎯 Pilih menu (1-%d): ", maxOption), 1, maxOption)

		if loginAkun.Role == "admin" {
			switch pilihan {
			case 1:
				tampilkanTiket()
			case 2:
				menuCariTiket()
			case 3:
				menuUrutkanTiket()
			case 4:
				menuKonfirmasiTiket()
			case 5:
				menuHapusTiket()
			case 6:
				tampilkanLaporanExcel()
			case 7:
				tampilkanLogoutSpongeBob()
				loginAkun = nil
				return
			}
		} else {
			switch pilihan {
			case 1:
				tampilkanTiket()
			case 2:
				menuCariTiket()
			case 3:
				menuUrutkanTiket()
			case 4:
				beliTiketUser()
			case 5:
				statusTiketUser()
			case 6:
				tampilkanLogoutSpongeBob()
				loginAkun = nil
				return
			}
		}
	}
}

func tampilkanLogoutSpongeBob() {
	clearScreen()

	fmt.Print(YELLOW)
	fmt.Println("      .-\"\"\"\"-.    ")
	fmt.Println("     /  😢   \\   ")
	fmt.Println("    /_      _\\   ")
	fmt.Println("   // \\👋 / \\\\  ")
	fmt.Println("   |\\__\\  /__/|  ")
	fmt.Println("    \\   ~~   /   ")
	fmt.Println("     \\__||__/    ")
	fmt.Println("     /  ><  \\    ")
	fmt.Println("    /______\\     ")
	fmt.Println("   (   😔   )    ")
	fmt.Println("    \\______/     ")
	fmt.Println("      |  |       ")
	fmt.Println("     /____\\      ")
	fmt.Print(RESET)

	logout := []string{
		"🌊 Sampai jumpa! 🌊",
		"",
		"Jangan lupa kembali ke Bikini Bottom!",
		"Terima kasih telah menggunakan sistem kami.",
	}
	printBox("LOGOUT", logout, 50, CYAN, BOLD+WHITE)

	time.Sleep(2 * time.Second)
}

// ==================== MENU HAPUS TIKET ====================

func menuHapusTiket() {
	clearScreen()

	header := []string{
		"🧽 SpongeBob akan membantu menghapus tiket dengan hati-hati! 🧽",
		"",
		"⚠️  Pastikan Anda yakin sebelum menghapus!",
	}
	printBox("MENU HAPUS TIKET", header, 60, RED, BOLD+WHITE)

	menuItems := []string{
		"",
		"🎯 1. Hapus Tiket Berdasarkan ID",
		"📊 2. Hapus Tiket Berdasarkan Status",
		"💥 3. Hapus Semua Tiket (HATI-HATI!)",
		"🔙 4. Kembali",
		"",
	}
	printBox("PILIHAN HAPUS", menuItems, 50, RED, BOLD+WHITE)

	pilihan := validasiInputInteger("🎯 Pilih menu (1-4): ", 1, 4)

	switch pilihan {
	case 1:
		hapusTiketByID()
	case 2:
		hapusTiketByStatus()
	case 3:
		hapusSemuaTiket()
	case 4:
		return
	}
}

func hapusTiketByID() {
	clearScreen()

	if len(daftarTiket) == 0 {
		noTickets := []string{
			"Tidak ada tiket untuk dihapus.",
			"",
			"🧽 SpongeBob berkata: 'Daftar tiket kosong seperti otak Patrick!' 🧽",
		}
		printBox("TIDAK ADA TIKET", noTickets, 60, WARNING, BOLD+WHITE)
		time.Sleep(3 * time.Second)
		return
	}

	// Tampilkan semua tiket terlebih dahulu
	fmt.Println(INFO + "📋 Daftar tiket yang tersedia:" + RESET)
	tampilkanTiket()

	id := validasiIDTiket("🎯 Masukkan ID tiket yang akan dihapus (contoh: TKT001): ")

	// Cari dan hapus tiket
	for i, tiket := range daftarTiket {
		if tiket.ID == id {
			// Tampilkan detail tiket yang akan dihapus
			clearScreen()
			fmt.Println(WARNING + "⚠️  Detail Tiket yang akan dihapus:" + RESET)
			tampilkanTiketArrayWithStatus([]Tiket{tiket})

			// Konfirmasi penghapusan
			if validasiKonfirmasi("\n⚠️  Apakah Anda yakin ingin menghapus tiket ini? (y/n): ") {
				// Hapus tiket dari slice
				daftarTiket = append(daftarTiket[:i], daftarTiket[i+1:]...)
				simpanTiketKeFile()

				clearScreen()
				tampilkanSpongeBobKecil()

				success := []string{
					fmt.Sprintf("🗑️ Tiket %s berhasil dihapus! 🗑️", id),
					"",
					"🧽 SpongeBob berkata: 'Tiket sudah hilang seperti spatula yang hilang!' 🧽",
				}
				printBox("PENGHAPUSAN BERHASIL", success, 60, SUCCESS, BOLD+WHITE)
			} else {
				cancel := []string{
					"Penghapusan dibatalkan.",
					"",
					"🧽 SpongeBob berkata: 'Phew! Tiket masih aman!' 🧽",
				}
				printBox("DIBATALKAN", cancel, 50, INFO, BOLD+WHITE)
			}
			time.Sleep(3 * time.Second)
			return
		}
	}

	notFound := []string{
		fmt.Sprintf("Tiket dengan ID %s tidak ditemukan.", id),
		"",
		"🧽 SpongeBob berkata: 'ID ini tidak ada di sistem Bikini Bottom!' 🧽",
	}
	printBox("TIKET TIDAK DITEMUKAN", notFound, 60, ERROR, BOLD+WHITE)
	time.Sleep(3 * time.Second)
}

func hapusTiketByStatus() {
	clearScreen()

	if len(daftarTiket) == 0 {
		noTickets := []string{
			"Tidak ada tiket untuk dihapus.",
			"",
			"🧽 SpongeBob berkata: 'Daftar tiket kosong!' 🧽",
		}
		printBox("TIDAK ADA TIKET", noTickets, 50, WARNING, BOLD+WHITE)
		time.Sleep(3 * time.Second)
		return
	}

	status := validasiStatus("📊 Masukkan status tiket yang akan dihapus (pending/confirmed/rejected): ")

	// Cari tiket dengan status yang diminta
	var tiketDitemukan []Tiket
	var indexTiket []int

	for i, tiket := range daftarTiket {
		if tiket.Status == status {
			tiketDitemukan = append(tiketDitemukan, tiket)
			indexTiket = append(indexTiket, i)
		}
	}

	if len(tiketDitemukan) == 0 {
		notFound := []string{
			fmt.Sprintf("Tidak ada tiket dengan status '%s'.", status),
			"",
			"🧽 SpongeBob berkata: 'Tidak ada tiket dengan status ini!' 🧽",
		}
		printBox("TIKET TIDAK DITEMUKAN", notFound, 60, WARNING, BOLD+WHITE)
		time.Sleep(3 * time.Second)
		return
	}

	// Tampilkan tiket yang akan dihapus
	clearScreen()
	fmt.Printf(WARNING+"⚠️  %d Tiket dengan status '%s' yang akan dihapus:"+RESET+"\n\n", len(tiketDitemukan), status)
	tampilkanTiketArrayWithStatus(tiketDitemukan)

	// Konfirmasi penghapusan
	if validasiKonfirmasi(fmt.Sprintf("\n⚠️  Apakah Anda yakin ingin menghapus %d tiket dengan status '%s'? (y/n): ", len(tiketDitemukan), status)) {
		// Hapus tiket dari belakang ke depan untuk menghindari perubahan index
		for i := len(indexTiket) - 1; i >= 0; i-- {
			idx := indexTiket[i]
			daftarTiket = append(daftarTiket[:idx], daftarTiket[idx+1:]...)
		}
		simpanTiketKeFile()

		clearScreen()
		tampilkanSpongeBobKecil()

		success := []string{
			fmt.Sprintf("🗑️ %d tiket dengan status '%s' berhasil dihapus! 🗑️", len(tiketDitemukan), status),
			"",
			"🧽 SpongeBob berkata: 'Pembersihan selesai seperti membersihkan Krusty Krab!' 🧽",
		}
		printBox("PENGHAPUSAN BERHASIL", success, 70, SUCCESS, BOLD+WHITE)
	} else {
		cancel := []string{
			"Penghapusan dibatalkan.",
			"",
			"🧽 SpongeBob berkata: 'Tiket-tiket masih aman!' 🧽",
		}
		printBox("DIBATALKAN", cancel, 50, INFO, BOLD+WHITE)
	}
	time.Sleep(3 * time.Second)
}

func hapusSemuaTiket() {
	clearScreen()

	if len(daftarTiket) == 0 {
		noTickets := []string{
			"Tidak ada tiket untuk dihapus.",
			"",
			"🧽 SpongeBob berkata: 'Daftar sudah kosong!' 🧽",
		}
		printBox("TIDAK ADA TIKET", noTickets, 50, WARNING, BOLD+WHITE)
		time.Sleep(3 * time.Second)
		return
	}

	warning := []string{
		"⚠️⚠️⚠️ PERINGATAN KERAS! ⚠️⚠️⚠️",
		"",
		fmt.Sprintf("Anda akan menghapus SEMUA %d tiket!", len(daftarTiket)),
		"Tindakan ini TIDAK DAPAT DIBATALKAN!",
		"",
		"🧽 SpongeBob berkata: 'Ini seperti menghancurkan seluruh Krusty Krab!' 🧽",
	}
	printBox("PERINGATAN BERBAHAYA", warning, 70, ERROR, BOLD+WHITE)

	// Konfirmasi ganda dengan input khusus
	fmt.Print(ERROR + "⚠️  Ketik 'HAPUS SEMUA' untuk melanjutkan (case sensitive): " + RESET)
	konfirmasi, _ := reader.ReadString('\n')
	konfirmasi = strings.TrimSpace(konfirmasi)

	if konfirmasi == "HAPUS SEMUA" {
		if validasiKonfirmasi("\n⚠️  Apakah Anda BENAR-BENAR yakin? (y/n): ") {
			jumlahTiket := len(daftarTiket)
			daftarTiket = []Tiket{} // Kosongkan slice
			lastID = 0              // Reset ID counter
			simpanTiketKeFile()

			clearScreen()

			// SpongeBob sedih
			fmt.Print(YELLOW)
			fmt.Println("      .-\"\"\"\"-.    ")
			fmt.Println("     /  😭   \\   ")
			fmt.Println("    /_      _\\   ")
			fmt.Println("   // \\💧 / \\\\  ")
			fmt.Println("   |\\__\\  /__/|  ")
			fmt.Println("    \\   😢   /   ")
			fmt.Println("     \\__||__/    ")
			fmt.Println("     /  ><  \\    ")
			fmt.Println("    /______\\     ")
			fmt.Println("   (   😭   )    ")
			fmt.Println("    \\______/     ")
			fmt.Println("      |  |       ")
			fmt.Println("     /____\\      ")
			fmt.Print(RESET)

			destroyed := []string{
				fmt.Sprintf("💥 SEMUA %d tiket telah dihapus! 💥", jumlahTiket),
				"",
				"😭 SpongeBob berkata: 'Semuanya hilang seperti formula Krabby Patty!' 😭",
			}
			printBox("PENGHAPUSAN TOTAL", destroyed, 60, ERROR, BOLD+WHITE)
		} else {
			cancel := []string{
				"Penghapusan dibatalkan.",
				"",
				"😌 SpongeBob berkata: 'Fiuh! Semua tiket masih aman!' 😌",
			}
			printBox("DIBATALKAN", cancel, 50, SUCCESS, BOLD+WHITE)
		}
	} else {
		invalid := []string{
			"Konfirmasi tidak tepat. Penghapusan dibatalkan.",
			"",
			"🧽 SpongeBob berkata: 'Lebih baik aman daripada menyesal!' 🧽",
		}
		printBox("KONFIRMASI SALAH", invalid, 60, INFO, BOLD+WHITE)
	}
	time.Sleep(3 * time.Second)
}

// ==================== MENU KONFIRMASI TIKET ====================

func menuKonfirmasiTiket() {
	clearScreen()

	header := []string{
		"🧽 SpongeBob siap membantu mengelola konfirmasi tiket! 🧽",
		"",
		"Pilih aksi yang ingin dilakukan:",
	}
	printBox("KELOLA KONFIRMASI TIKET", header, 60, BLUE, BOLD+WHITE)

	menuItems := []string{
		"",
		"⏳ 1. Lihat Tiket Pending",
		"✅ 2. Konfirmasi Tiket",
		"❌ 3. Tolak Tiket",
		"📊 4. Lihat Semua Status",
		"🔙 5. Kembali",
		"",
	}
	printBox("MENU KONFIRMASI", menuItems, 50, BLUE, BOLD+WHITE)

	pilihan := validasiInputInteger("🎯 Pilih menu (1-5): ", 1, 5)

	switch pilihan {
	case 1:
		tampilkanTiketPending()
	case 2:
		konfirmasiTiket()
	case 3:
		tolakTiket()
	case 4:
		tampilkanSemuaStatus()
	case 5:
		return
	}
}

func tampilkanTiketPending() {
	clearScreen()

	var tiketPending []Tiket
	for _, tiket := range daftarTiket {
		if tiket.Status == "pending" {
			tiketPending = append(tiketPending, tiket)
		}
	}

	if len(tiketPending) == 0 {
		noPending := []string{
			"Tidak ada tiket yang menunggu konfirmasi.",
			"",
			"🧽 SpongeBob berkata: 'Semua tiket sudah diproses!' 🧽",
		}
		printBox("TIDAK ADA TIKET PENDING", noPending, 60, INFO, BOLD+WHITE)
		time.Sleep(3 * time.Second)
		return
	}

	fmt.Println(WARNING + "⏳ Tiket Menunggu Konfirmasi:" + RESET)
	tampilkanTiketArrayWithStatus(tiketPending)

	fmt.Print(DIM + "\nTekan Enter untuk kembali..." + RESET)
	reader.ReadString('\n')
}

func konfirmasiTiket() {
	clearScreen()

	fmt.Println(SUCCESS + "✅ Konfirmasi Tiket" + RESET)
	tampilkanTiketPending()

	if len(daftarTiket) == 0 {
		return
	}

	id := validasiIDTiket("🎯 Masukkan ID tiket yang akan dikonfirmasi: ")

	for i, tiket := range daftarTiket {
		if tiket.ID == id && tiket.Status == "pending" {
			daftarTiket[i].Status = "confirmed"
			simpanTiketKeFile()

			clearScreen()
			tampilkanSpongeBobKecil()

			success := []string{
				fmt.Sprintf("🎉 Tiket %s berhasil dikonfirmasi! 🎉", id),
				"",
				"🌟 SpongeBob berkata: 'I'm ready! I'm ready!' 🌟",
			}
			printBox("KONFIRMASI BERHASIL", success, 60, SUCCESS, BOLD+WHITE)
			time.Sleep(3 * time.Second)
			return
		}
	}

	notFound := []string{
		"Tiket tidak ditemukan atau sudah diproses.",
		"",
		"🧽 SpongeBob berkata: 'Tiket ini tidak bisa dikonfirmasi!' 🧽",
	}
	printBox("TIKET TIDAK DITEMUKAN", notFound, 60, ERROR, BOLD+WHITE)
	time.Sleep(3 * time.Second)
}

func tolakTiket() {
	clearScreen()

	fmt.Println(ERROR + "❌ Tolak Tiket" + RESET)
	tampilkanTiketPending()

	if len(daftarTiket) == 0 {
		return
	}

	id := validasiIDTiket("🎯 Masukkan ID tiket yang akan ditolak: ")

	for i, tiket := range daftarTiket {
		if tiket.ID == id && tiket.Status == "pending" {
			daftarTiket[i].Status = "rejected"
			simpanTiketKeFile()

			clearScreen()
			rejected := []string{
				fmt.Sprintf("❌ Tiket %s telah ditolak!", id),
				"",
				"😔 SpongeBob berkata: 'Aww, tartar sauce!' 😔",
			}
			printBox("TIKET DITOLAK", rejected, 50, ERROR, BOLD+WHITE)
			time.Sleep(3 * time.Second)
			return
		}
	}

	notFound := []string{
		"Tiket tidak ditemukan atau sudah diproses.",
		"",
		"🧽 SpongeBob berkata: 'Tiket ini tidak bisa ditolak!' 🧽",
	}
	printBox("TIKET TIDAK DITEMUKAN", notFound, 60, ERROR, BOLD+WHITE)
	time.Sleep(3 * time.Second)
}

func tampilkanSemuaStatus() {
	clearScreen()

	pending := 0
	confirmed := 0
	rejected := 0

	for _, tiket := range daftarTiket {
		switch tiket.Status {
		case "pending":
			pending++
		case "confirmed":
			confirmed++
		case "rejected":
			rejected++
		}
	}

	stats := []string{
		fmt.Sprintf("📋 Total Tiket: %d", len(daftarTiket)),
		fmt.Sprintf("⏳ Pending: %d", pending),
		fmt.Sprintf("✅ Confirmed: %d", confirmed),
		fmt.Sprintf("❌ Rejected: %d", rejected),
		"",
	}
	printBox("STATISTIK TIKET", stats, 40, INFO, BOLD+WHITE)

	if len(daftarTiket) > 0 {
		fmt.Println(INFO + "📋 Detail Semua Tiket:" + RESET)
		tampilkanTiketArrayWithStatus(daftarTiket)
	} else {
		noTickets := []string{
			"🧽 SpongeBob berkata: 'Belum ada tiket yang terdaftar!' 🧽",
		}
		printBox("BELUM ADA TIKET", noTickets, 60, WARNING, BOLD+WHITE)
	}

	fmt.Print(DIM + "\nTekan Enter untuk kembali..." + RESET)
	reader.ReadString('\n')
}

func statusTiketUser() {
	clearScreen()

	var tiketUser []Tiket
	for _, tiket := range daftarTiket {
		if tiket.Nama == loginAkun.Username {
			tiketUser = append(tiketUser, tiket)
		}
	}

	if len(tiketUser) == 0 {
		noTickets := []string{
			"Anda belum memiliki tiket.",
			"",
			"🎵 SpongeBob berkata: 'Ayo beli tiket untuk konser seru!' 🎵",
		}
		printBox("BELUM ADA TIKET", noTickets, 60, WARNING, BOLD+WHITE)
		tampilkanSpongeBobKecil()
		time.Sleep(3 * time.Second)
		return
	}

	fmt.Println(INFO + "📄 Status Tiket Saya:" + RESET)
	tampilkanTiketArrayWithStatus(tiketUser)

	fmt.Print(DIM + "\nTekan Enter untuk kembali..." + RESET)
	reader.ReadString('\n')
}

// ==================== PEMBELIAN TIKET ====================

func beliTiketUser() {
	clearScreen()

	header := []string{
		"🧽 SpongeBob siap membantu pembelian tiket Anda! 🧽",
		"",
		"Pilih kategori tiket yang diinginkan:",
	}
	printBox("PEMBELIAN TIKET", header, 60, GREEN, BOLD+WHITE)

	categories := []string{
		"",
		"🎫 1. Reguler - Rp50.000",
		"⭐ 2. VIP - Rp100.000",
		"🐦 3. Early Bird - Rp25.000",
		"",
	}
	printBox("KATEGORI TIKET", categories, 40, GREEN, BOLD+WHITE)

	kategori := validasiKategori("🎯 Masukkan tipe tiket (Reguler/VIP/Early Bird): ")

	var harga int
	switch kategori {
	case "reguler":
		harga = 50000
	case "vip":
		harga = 100000
	case "early bird":
		harga = 25000
	}

	// Konfirmasi pembelian
	clearScreen()
	confirmation := []string{
		"--- Konfirmasi Pembelian ---",
		"",
		fmt.Sprintf("👤 Nama: %s", loginAkun.Username),
		fmt.Sprintf("🎫 Kategori: %s", strings.Title(kategori)),
		fmt.Sprintf("💰 Harga: Rp%s", formatRupiah(harga)),
		"",
	}
	printBox("KONFIRMASI PEMBELIAN", confirmation, 50, YELLOW, BOLD+BLACK)

	if !validasiKonfirmasi("🛒 Lanjutkan pembelian? (y/n): ") {
		cancel := []string{
			"Pembelian dibatalkan.",
			"",
			"🧽 SpongeBob berkata: 'Tidak apa-apa, nanti beli lagi ya!' 🧽",
		}
		printBox("PEMBELIAN DIBATALKAN", cancel, 60, INFO, BOLD+WHITE)
		time.Sleep(3 * time.Second)
		return
	}

	tiket := Tiket{
		ID:       buatIDTiket(),
		Nama:     loginAkun.Username,
		Kategori: strings.Title(kategori),
		Harga:    harga,
		Status:   "pending",
		Tanggal:  time.Now().Format("2006-01-02"),
	}
	daftarTiket = append(daftarTiket, tiket)
	simpanTiketKeFile()

	loading("SpongeBob sedang menyiapkan struk")
	cetakStruk(tiket)
}

// ==================== PENCARIAN TIKET ====================

func menuCariTiket() {
	clearScreen()

	header := []string{
		"🔍 SpongeBob akan membantu mencari tiket! 🔍",
		"",
		"Pilih metode pencarian:",
	}
	printBox("MENU PENCARIAN", header, 50, CYAN, BOLD+WHITE)

	searchOptions := []string{
		"",
		"👤 1. Cari berdasarkan Nama",
		"🎫 2. Cari berdasarkan Kategori",
		"🆔 3. Cari berdasarkan ID",
		"📊 4. Cari berdasarkan Status",
		"🔙 5. Kembali",
		"",
	}
	printBox("OPSI PENCARIAN", searchOptions, 50, CYAN, BOLD+WHITE)

	pilihan := validasiInputInteger("🎯 Pilih metode pencarian (1-5): ", 1, 5)

	switch pilihan {
	case 1:
		cariTiketByNama()
	case 2:
		cariTiketByKategori()
	case 3:
		cariTiketByID()
	case 4:
		cariTiketByStatus()
	case 5:
		return
	}
}

func cariTiketByNama() {
	clearScreen()

	nama := validasiInputString("👤 Masukkan nama (minimal 1 karakter): ", 1, 50, false)
	nama = strings.ToLower(nama)

	var hasil []Tiket
	for _, tiket := range daftarTiket {
		if strings.Contains(strings.ToLower(tiket.Nama), nama) {
			hasil = append(hasil, tiket)
		}
	}

	clearScreen()
	if len(hasil) == 0 {
		notFound := []string{
			"Tiket tidak ditemukan.",
			"",
			"🧽 SpongeBob berkata: 'Tidak ada yang cocok, coba kata kunci lain!' 🧽",
		}
		printBox("HASIL PENCARIAN", notFound, 60, WARNING, BOLD+WHITE)
	} else {
		fmt.Printf(SUCCESS+"🎉 SpongeBob menemukan %d tiket! 🎉"+RESET+"\n\n", len(hasil))
		tampilkanTiketArrayWithStatus(hasil)
	}

	fmt.Print(DIM + "\nTekan Enter untuk kembali..." + RESET)
	reader.ReadString('\n')
}

func cariTiketByKategori() {
	clearScreen()

	kategori := validasiInputString("🎫 Masukkan kategori: ", 1, 20, false)
	kategori = strings.ToLower(kategori)

	var hasil []Tiket
	for _, tiket := range daftarTiket {
		if strings.Contains(strings.ToLower(tiket.Kategori), kategori) {
			hasil = append(hasil, tiket)
		}
	}

	clearScreen()
	if len(hasil) == 0 {
		notFound := []string{
			"Tiket tidak ditemukan.",
			"",
			"🧽 SpongeBob berkata: 'Kategori ini belum ada, coba yang lain!' 🧽",
		}
		printBox("HASIL PENCARIAN", notFound, 60, WARNING, BOLD+WHITE)
	} else {
		fmt.Printf(SUCCESS+"🎉 SpongeBob menemukan %d tiket kategori %s! 🎉"+RESET+"\n\n", len(hasil), kategori)
		tampilkanTiketArrayWithStatus(hasil)
	}

	fmt.Print(DIM + "\nTekan Enter untuk kembali..." + RESET)
	reader.ReadString('\n')
}

func cariTiketByID() {
	clearScreen()

	id := validasiIDTiket("🆔 Masukkan ID (contoh: TKT001): ")

	for _, tiket := range daftarTiket {
		if tiket.ID == id {
			clearScreen()
			fmt.Println(SUCCESS + "🎯 SpongeBob menemukan tiket yang tepat! 🎯" + RESET)
			tampilkanTiketArrayWithStatus([]Tiket{tiket})
			fmt.Print(DIM + "\nTekan Enter untuk kembali..." + RESET)
			reader.ReadString('\n')
			return
		}
	}

	notFound := []string{
		"Tiket tidak ditemukan.",
		"",
		"🧽 SpongeBob berkata: 'ID ini tidak ada di sistem!' 🧽",
	}
	printBox("HASIL PENCARIAN", notFound, 50, ERROR, BOLD+WHITE)
	time.Sleep(3 * time.Second)
}

func cariTiketByStatus() {
	clearScreen()

	status := validasiStatus("📊 Masukkan status (pending/confirmed/rejected): ")

	var hasil []Tiket
	for _, tiket := range daftarTiket {
		if tiket.Status == status {
			hasil = append(hasil, tiket)
		}
	}

	clearScreen()
	if len(hasil) == 0 {
		notFound := []string{
			"Tiket tidak ditemukan.",
			"",
			"🧽 SpongeBob berkata: 'Tidak ada tiket dengan status ini!' 🧽",
		}
		printBox("HASIL PENCARIAN", notFound, 60, WARNING, BOLD+WHITE)
	} else {
		fmt.Printf(SUCCESS+"📊 SpongeBob menemukan %d tiket dengan status %s! 📊"+RESET+"\n\n", len(hasil), status)
		tampilkanTiketArrayWithStatus(hasil)
	}

	fmt.Print(DIM + "\nTekan Enter untuk kembali..." + RESET)
	reader.ReadString('\n')
}

// ==================== PENGURUTAN TIKET ====================

func menuUrutkanTiket() {
	clearScreen()

	header := []string{
		"🧽 SpongeBob akan membantu mengurutkan tiket! 🧽",
		"",
		"Pilih metode pengurutan:",
	}
	printBox("MENU PENGURUTAN TIKET", header, 60, MAGENTA, BOLD+WHITE)

	sortOptions := []string{
		"",
		"💰 1. Selection Sort (berdasarkan Harga)",
		"👤 2. Insertion Sort (berdasarkan Nama)",
		"🆔 3. Built-in Sort (berdasarkan ID)",
		"🔙 4. Kembali",
		"",
	}
	printBox("METODE PENGURUTAN", sortOptions, 50, MAGENTA, BOLD+WHITE)

	pilihan := validasiInputInteger("🎯 Pilih metode pengurutan (1-4): ", 1, 4)

	switch pilihan {
	case 1:
		selectionSortByHarga()
	case 2:
		insertionSortByNama()
	case 3:
		builtinSortByID()
	case 4:
		return
	}
}

// Selection Sort - Mengurutkan berdasarkan harga (ascending)
func selectionSortByHarga() {
	clearScreen()

	if len(daftarTiket) == 0 {
		noTickets := []string{
			"Tidak ada tiket untuk diurutkan.",
			"",
			"🧽 SpongeBob berkata: 'Daftar tiket kosong!' 🧽",
		}
		printBox("TIDAK ADA TIKET", noTickets, 50, WARNING, BOLD+WHITE)
		time.Sleep(3 * time.Second)
		return
	}

	header := []string{
		"🧽 SpongeBob akan mengurutkan tiket seperti mengurutkan spatula! 🧽",
		"",
		"Metode: Selection Sort berdasarkan Harga (Ascending)",
	}
	printBox("SELECTION SORT", header, 70, YELLOW, BOLD+BLACK)

	// Buat salinan untuk sorting
	tiketCopy := make([]Tiket, len(daftarTiket))
	copy(tiketCopy, daftarTiket)

	fmt.Println(INFO + "📋 Sebelum diurutkan:" + RESET)
	tampilkanTiketArrayWithStatus(tiketCopy)

	// Selection Sort Algorithm
	n := len(tiketCopy)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if tiketCopy[j].Harga < tiketCopy[minIdx].Harga {
				minIdx = j
			}
		}
		// Swap
		tiketCopy[i], tiketCopy[minIdx] = tiketCopy[minIdx], tiketCopy[i]

		// Tampilkan proses sorting
		fmt.Printf(WARNING+"🔄 Langkah %d: SpongeBob menukar posisi %d dengan %d"+RESET+"\n", i+1, i, minIdx)
		time.Sleep(1 * time.Second)
	}

	clearScreen()
	fmt.Println(SUCCESS + "✅ Hasil Selection Sort (Harga Ascending):" + RESET)
	fmt.Println(WARNING + "🎉 SpongeBob berkata: 'Terurut rapi seperti Krabby Patty!' 🎉" + RESET)
	tampilkanTiketArrayWithStatus(tiketCopy)

	if validasiKonfirmasi("\n💾 Apakah ingin menyimpan hasil pengurutan? (y/n): ") {
		daftarTiket = tiketCopy
		simpanTiketKeFile()

		clearScreen()
		tampilkanSpongeBobKecil()

		saved := []string{
			"💾 Pengurutan disimpan! SpongeBob bangga! 💾",
		}
		printBox("TERSIMPAN", saved, 50, SUCCESS, BOLD+WHITE)
		time.Sleep(2 * time.Second)
	}
}

// Insertion Sort - Mengurutkan berdasarkan nama (ascending)
func insertionSortByNama() {
	clearScreen()

	if len(daftarTiket) == 0 {
		noTickets := []string{
			"Tidak ada tiket untuk diurutkan.",
			"",
			"🧽 SpongeBob berkata: 'Daftar tiket kosong!' 🧽",
		}
		printBox("TIDAK ADA TIKET", noTickets, 50, WARNING, BOLD+WHITE)
		time.Sleep(3 * time.Second)
		return
	}

	header := []string{
		"🧽 SpongeBob akan menyisipkan nama seperti menyisipkan keju di burger! 🧽",
		"",
		"Metode: Insertion Sort berdasarkan Nama (Ascending)",
	}
	printBox("INSERTION SORT", header, 70, BLUE, BOLD+WHITE)

	// Buat salinan untuk sorting
	tiketCopy := make([]Tiket, len(daftarTiket))
	copy(tiketCopy, daftarTiket)

	fmt.Println(INFO + "📋 Sebelum diurutkan:" + RESET)
	tampilkanTiketArrayWithStatus(tiketCopy)

	// Insertion Sort Algorithm
	n := len(tiketCopy)
	for i := 1; i < n; i++ {
		key := tiketCopy[i]
		j := i - 1

		// Pindahkan elemen yang lebih besar dari key ke posisi setelahnya
		for j >= 0 && strings.ToLower(tiketCopy[j].Nama) > strings.ToLower(key.Nama) {
			tiketCopy[j+1] = tiketCopy[j]
			j--
		}
		tiketCopy[j+1] = key

		// Tampilkan proses sorting
		fmt.Printf(WARNING+"📝 Langkah %d: SpongeBob memasukkan '%s' ke posisi yang tepat"+RESET+"\n", i, key.Nama)
		time.Sleep(1 * time.Second)
	}

	clearScreen()
	fmt.Println(SUCCESS + "✅ Hasil Insertion Sort (Nama Ascending):" + RESET)
	fmt.Println(WARNING + "🎉 SpongeBob berkata: 'Nama-nama tersusun alfabetis!' 🎉" + RESET)
	tampilkanTiketArrayWithStatus(tiketCopy)

	if validasiKonfirmasi("\n💾 Apakah ingin menyimpan hasil pengurutan? (y/n): ") {
		daftarTiket = tiketCopy
		simpanTiketKeFile()

		clearScreen()
		tampilkanSpongeBobKecil()

		saved := []string{
			"💾 Pengurutan disimpan! Gary juga senang! 💾",
		}
		printBox("TERSIMPAN", saved, 50, SUCCESS, BOLD+WHITE)
		time.Sleep(2 * time.Second)
	}
}

// Built-in Sort - Mengurutkan berdasarkan ID
func builtinSortByID() {
	clearScreen()

	if len(daftarTiket) == 0 {
		noTickets := []string{
			"Tidak ada tiket untuk diurutkan.",
			"",
			"🧽 SpongeBob berkata: 'Daftar tiket kosong!' 🧽",
		}
		printBox("TIDAK ADA TIKET", noTickets, 50, WARNING, BOLD+WHITE)
		time.Sleep(3 * time.Second)
		return
	}

	header := []string{
		"🧽 SpongeBob menggunakan magic sorting seperti Sandy! 🧽",
		"",
		"Metode: Built-in Sort berdasarkan ID (Ascending)",
	}
	printBox("BUILT-IN SORT", header, 60, GREEN, BOLD+WHITE)

	fmt.Println(INFO + "📋 Sebelum diurutkan:" + RESET)
	tampilkanTiket()

	loading("SpongeBob sedang melakukan magic sorting")

	sort.Slice(daftarTiket, func(i, j int) bool {
		return daftarTiket[i].ID < daftarTiket[j].ID
	})

	clearScreen()
	fmt.Println(SUCCESS + "✅ Hasil Built-in Sort (ID Ascending):" + RESET)
	fmt.Println(WARNING + "⚡ SpongeBob berkata: 'Cepat seperti kilat!' ⚡" + RESET)
	tampilkanTiket()

	simpanTiketKeFile()
	tampilkanSpongeBobKecil()

	saved := []string{
		"💾 Pengurutan disimpan otomatis! 💾",
	}
	printBox("TERSIMPAN", saved, 40, SUCCESS, BOLD+WHITE)
	time.Sleep(2 * time.Second)
}

// ==================== LAPORAN ====================

func tampilkanLaporanExcel() {
	clearScreen()

	if len(daftarTiket) == 0 {
		noTickets := []string{
			"Tidak ada data tiket untuk ditampilkan.",
			"",
			"🧽 SpongeBob berkata: 'Belum ada tiket untuk dilaporkan!' 🧽",
		}
		printBox("TIDAK ADA DATA", noTickets, 60, WARNING, BOLD+WHITE)
		time.Sleep(3 * time.Second)
		return
	}

	loading("SpongeBob sedang menyiapkan laporan")
	clearScreen()

	// SpongeBob ASCII untuk laporan
	fmt.Print(YELLOW)
	fmt.Println("                 📊 LAPORAN DARI BIKINI BOTTOM 📊")
	fmt.Println("                    .-\"\"\"\"-.")
	fmt.Println("                   /  🤓   \\")
	fmt.Println("                  /_      _\\")
	fmt.Println("                 // \\👓 / \\\\")
	fmt.Println("                 |\\__\\  /__/|")
	fmt.Println("                  \\   📋   /")
	fmt.Println("                   \\__||__/")
	fmt.Println("                   /  ><  \\")
	fmt.Println("                  /______\\")
	fmt.Println("                 (  📈📊  )")
	fmt.Println("                  \\______/")
	fmt.Println("                     |  |")
	fmt.Println("                    /____\\")
	fmt.Print(RESET)

	// Header Laporan
	reportHeader := []string{
		"📊 LAPORAN TIKET KONSER RT07 RW09",
		"",
		fmt.Sprintf("📅 Tanggal: %s", time.Now().Format("2006-01-02 15:04:05")),
		"🧽 Disajikan oleh SpongeBob SquarePants 🧽",
	}
	printBox("LAPORAN TIKET", reportHeader, 80, CYAN, BOLD+WHITE)

	// Statistik Ringkas
	pending, confirmed, rejected := 0, 0, 0
	totalHarga, totalPendapatan := 0, 0
	kategoriCount := make(map[string]int)

	for _, tiket := range daftarTiket {
		totalHarga += tiket.Harga
		kategoriCount[tiket.Kategori]++

		switch tiket.Status {
		case "pending":
			pending++
		case "confirmed":
			confirmed++
			totalPendapatan += tiket.Harga
		case "rejected":
			rejected++
		}
	}

	stats := []string{
		"📈 RINGKASAN STATISTIK (Lebih akurat dari ramalan Patrick!)",
		"",
		fmt.Sprintf("├─ Total Tiket Terjual    : %d tiket", len(daftarTiket)),
		fmt.Sprintf("├─ ⏳ Menunggu Konfirmasi : %d tiket", pending),
		fmt.Sprintf("├─ ✅ Dikonfirmasi        : %d tiket", confirmed),
		fmt.Sprintf("├─ ❌ Ditolak             : %d tiket", rejected),
		fmt.Sprintf("├─ 💰 Total Nilai Tiket   : Rp %s", formatRupiah(totalHarga)),
		fmt.Sprintf("└─ 💵 Pendapatan Bersih   : Rp %s (Lebih banyak dari gaji di Krusty Krab!)", formatRupiah(totalPendapatan)),
	}
	printBox("STATISTIK", stats, 80, GREEN, BOLD+WHITE)

	// Breakdown per Kategori
	categoryStats := []string{
		"🎫 BREAKDOWN PER KATEGORI (Gary juga ikut menghitung!)",
		"",
	}
	for kategori, jumlah := range kategoriCount {
		persentase := float64(jumlah) / float64(len(daftarTiket)) * 100
		categoryStats = append(categoryStats, fmt.Sprintf("├─ %-15s : %d tiket (%.1f%%)", kategori, jumlah, persentase))
	}
	printBox("KATEGORI", categoryStats, 60, BLUE, BOLD+WHITE)

	// Tabel Data Detail
	fmt.Println(BOLD + BLUE + "\n📋 DATA DETAIL TIKET KONSER" + RESET)
	fmt.Println(DIM + "🎵 Setiap tiket adalah karya seni! 🎵" + RESET)
	tampilkanTiketArrayWithStatus(daftarTiket)

	// Opsi simpan ke file
	if validasiKonfirmasi("\n💾 Simpan laporan ke file teks? SpongeBob siap membantu! (y/n): ") {
		simpanLaporanKeFile()
	}

	fmt.Print(DIM + "\nTekan Enter untuk kembali..." + RESET)
	reader.ReadString('\n')
}

func formatRupiah(angka int) string {
	str := strconv.Itoa(angka)
	n := len(str)
	if n <= 3 {
		return str
	}

	result := ""
	for i, digit := range str {
		if i > 0 && (n-i)%3 == 0 {
			result += "."
		}
		result += string(digit)
	}
	return result
}

func simpanLaporanKeFile() {
	filename := fmt.Sprintf("laporan_tiket_spongebob_%s.txt", time.Now().Format("2006-01-02_15-04-05"))
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf(ERROR+"Error membuat file: %v"+RESET+"\n", err)
		fmt.Println(WARNING + "🧽 SpongeBob berkata: 'Oops! Tidak bisa buat file!' 🧽" + RESET)
		return
	}
	defer file.Close()

	// Tulis header dengan SpongeBob
	fmt.Fprintf(file, "🧽 LAPORAN TIKET KONSER RT07 RW09 - BY SPONGEBOB 🧽\n")
	fmt.Fprintf(file, "Tanggal: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Fprintf(file, "Dibuat dengan cinta dari Bikini Bottom! 💙\n\n")

	// Tulis statistik
	pending, confirmed, rejected := 0, 0, 0
	totalHarga, totalPendapatan := 0, 0

	for _, tiket := range daftarTiket {
		totalHarga += tiket.Harga
		switch tiket.Status {
		case "pending":
			pending++
		case "confirmed":
			confirmed++
			totalPendapatan += tiket.Harga
		case "rejected":
			rejected++
		}
	}

	fmt.Fprintf(file, "STATISTIK KONSER:\n")
	fmt.Fprintf(file, "Total Tiket: %d (Lebih banyak dari gelembung yang pernah saya tiup!)\n", len(daftarTiket))
	fmt.Fprintf(file, "Pending: %d\n", pending)
	fmt.Fprintf(file, "Confirmed: %d\n", confirmed)
	fmt.Fprintf(file, "Rejected: %d\n", rejected)
	fmt.Fprintf(file, "Total Nilai: Rp %s\n", formatRupiah(totalHarga))
	fmt.Fprintf(file, "Pendapatan: Rp %s (Mr. Krabs akan senang!)\n\n", formatRupiah(totalPendapatan))

	// Tulis data detail
	fmt.Fprintf(file, "DATA DETAIL TIKET:\n")
	fmt.Fprintf(file, "No\tID\tNama\tKategori\tHarga\tStatus\tTanggal\n")
	for i, tiket := range daftarTiket {
		fmt.Fprintf(file, "%d\t%s\t%s\t%s\tRp %s\t%s\t%s\n",
			i+1, tiket.ID, tiket.Nama, tiket.Kategori,
			formatRupiah(tiket.Harga), tiket.Status, tiket.Tanggal)
	}

	fmt.Fprintf(file, "\n🌟 Terima kasih telah menggunakan sistem tiket SpongeBob! 🌟\n")
	fmt.Fprintf(file, "Sampai jumpa di konser! - SpongeBob SquarePants 🧽\n")

	clearScreen()
	tampilkanSpongeBobKecil()

	success := []string{
		fmt.Sprintf("🎉 Laporan berhasil disimpan: %s 🎉", filename),
		"",
		"🧽 SpongeBob berkata: 'File tersimpan dengan sempurna!' 🧽",
	}
	printBox("LAPORAN TERSIMPAN", success, 70, SUCCESS, BOLD+WHITE)
	time.Sleep(3 * time.Second)
}

// ==================== FUNGSI TAMPILAN ====================

func tampilkanTiketArrayWithStatus(tikets []Tiket) {
	if len(tikets) == 0 {
		noTickets := []string{
			"Tidak ada tiket.",
		}
		printBox("TIDAK ADA TIKET", noTickets, 40, WARNING, BOLD+WHITE)
		return
	}

	// Header tabel
	fmt.Print(BLUE + "╔")
	fmt.Print(strings.Repeat("═", 95))
	fmt.Print("╗" + RESET + "\n")

	fmt.Print(BLUE + "║" + RESET)
	fmt.Print(BOLD + WHITE + "                                   🎫 DAFTAR TIKET KONSER 🎫                                   " + RESET)
	fmt.Print(BLUE + "║" + RESET + "\n")

	fmt.Print(BLUE + "║" + RESET)
	fmt.Print(DIM + "                              🧽 Dikelola oleh SpongeBob SquarePants 🧽                       " + RESET)
	fmt.Print(BLUE + "║" + RESET + "\n")

	fmt.Print(BLUE + "╠")
	fmt.Print(strings.Repeat("═", 95))
	fmt.Print("╣" + RESET + "\n")

	// Header kolom
	fmt.Printf(BLUE+"║"+RESET+" %-8s "+BLUE+"║"+RESET+" %-15s "+BLUE+"║"+RESET+" %-12s "+BLUE+"║"+RESET+" %-12s "+BLUE+"║"+RESET+" %-11s "+BLUE+"║"+RESET+" %-12s "+BLUE+"║"+RESET+"\n",
		"ID", "Nama", "Kategori", "Harga", "Status", "Tanggal")

	fmt.Print(BLUE + "╠")
	fmt.Print(strings.Repeat("═", 95))
	fmt.Print("╣" + RESET + "\n")

	// Data tiket
	for _, tiket := range tikets {
		statusIcon := ""
		statusColor := ""
		switch tiket.Status {
		case "pending":
			statusIcon = "⏳"
			statusColor = WARNING
		case "confirmed":
			statusIcon = "✅"
			statusColor = SUCCESS
		case "rejected":
			statusIcon = "❌"
			statusColor = ERROR
		}

		fmt.Printf(BLUE+"║"+RESET+" %-8s "+BLUE+"║"+RESET+" %-15s "+BLUE+"║"+RESET+" %-12s "+BLUE+"║"+RESET+" %-12s "+BLUE+"║"+RESET+" %s%s%-10s"+RESET+" "+BLUE+"║"+RESET+" %-12s "+BLUE+"║"+RESET+"\n",
			tiket.ID, tiket.Nama, tiket.Kategori, formatRupiah(tiket.Harga),
			statusColor, statusIcon, tiket.Status, tiket.Tanggal)
	}

	fmt.Print(BLUE + "╚")
	fmt.Print(strings.Repeat("═", 95))
	fmt.Print("╝" + RESET + "\n")
}

func tampilkanTiket() {
	clearScreen()

	if len(daftarTiket) == 0 {
		noTickets := []string{
			"Belum ada tiket yang terdaftar.",
			"",
			"🧽 SpongeBob berkata: 'Ayo mulai jual tiket!' 🧽",
		}
		printBox("BELUM ADA TIKET", noTickets, 50, WARNING, BOLD+WHITE)
		time.Sleep(3 * time.Second)
		return
	}

	tampilkanTiketArrayWithStatus(daftarTiket)

	fmt.Print(DIM + "\nTekan Enter untuk kembali..." + RESET)
	reader.ReadString('\n')
}

func buatIDTiket() string {
	lastID++
	return fmt.Sprintf("TKT%03d", lastID)
}

func cetakStruk(tiket Tiket) {
	clearScreen()

	// Struk dengan design yang lebih menarik
	fmt.Print(YELLOW + "╔")
	fmt.Print(strings.Repeat("═", 53))
	fmt.Print("╗" + RESET + "\n")

	fmt.Print(YELLOW + "║" + RESET)
	fmt.Print(BOLD + WHITE + "              🧽 STRUK PEMBELIAN TIKET 🧽           " + RESET)
	fmt.Print(YELLOW + "║" + RESET + "\n")

	fmt.Print(YELLOW + "║" + RESET)
	fmt.Print(DIM + "              Dari SpongeBob dengan ❤️           " + RESET)
	fmt.Print(YELLOW + "║" + RESET + "\n")

	fmt.Print(YELLOW + "╠")
	fmt.Print(strings.Repeat("═", 53))
	fmt.Print("╣" + RESET + "\n")

	// Detail tiket
	fmt.Printf(YELLOW+"║"+RESET+" ID Tiket    : %-30s "+YELLOW+"║"+RESET+"\n", tiket.ID)
	fmt.Printf(YELLOW+"║"+RESET+" Nama        : %-30s "+YELLOW+"║"+RESET+"\n", tiket.Nama)
	fmt.Printf(YELLOW+"║"+RESET+" Kategori    : %-30s "+YELLOW+"║"+RESET+"\n", tiket.Kategori)
	fmt.Printf(YELLOW+"║"+RESET+" Harga       : Rp%-27s "+YELLOW+"║"+RESET+"\n", formatRupiah(tiket.Harga))
	fmt.Printf(YELLOW+"║"+RESET+" Status      : %-30s "+YELLOW+"║"+RESET+"\n", strings.Title(tiket.Status))
	fmt.Printf(YELLOW+"║"+RESET+" Tanggal     : %-30s "+YELLOW+"║"+RESET+"\n", tiket.Tanggal)
	fmt.Printf(YELLOW+"║"+RESET+" Waktu Beli  : %-30s "+YELLOW+"║"+RESET+"\n", time.Now().Format("2006-01-02 15:04:05"))

	fmt.Print(YELLOW + "╠")
	fmt.Print(strings.Repeat("═", 53))
	fmt.Print("╣" + RESET + "\n")

	fmt.Print(YELLOW + "║" + RESET)
	fmt.Print(WARNING + "  🎵 Tiket Anda sedang menunggu konfirmasi admin!  " + RESET)
	fmt.Print(YELLOW + "║" + RESET + "\n")

	fmt.Print(YELLOW + "║" + RESET)
	fmt.Print(SUCCESS + "     🌟 Terima kasih telah mempercayai kami! 🌟    " + RESET)
	fmt.Print(YELLOW + "║" + RESET + "\n")

	fmt.Print(YELLOW + "╚")
	fmt.Print(strings.Repeat("═", 53))
	fmt.Print("╝" + RESET + "\n")

	tampilkanSpongeBobKecil()

	success := []string{
		"🎊 SpongeBob berkata: 'Yeay! Tiket berhasil dibeli!' 🎊",
	}
	printBox("PEMBELIAN BERHASIL", success, 60, SUCCESS, BOLD+WHITE)

	time.Sleep(3 * time.Second)
}

// ==================== FUNGSI FILE I/O ====================

func simpanTiketKeFile() {
	file, err := os.Create("tiket.txt")
	if err != nil {
		fmt.Printf(ERROR+"Error menyimpan tiket: %v"+RESET+"\n", err)
		return
	}
	defer file.Close()

	for _, tiket := range daftarTiket {
		fmt.Fprintf(file, "%s|%s|%s|%d|%s|%s\n", tiket.ID, tiket.Nama, tiket.Kategori, tiket.Harga, tiket.Status, tiket.Tanggal)
	}
}

func muatDataTiket() {
	file, err := os.Open("tiket.txt")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) >= 4 {
			harga, err := strconv.Atoi(parts[3])
			if err != nil {
				continue // Skip invalid data
			}

			status := "pending"
			tanggal := time.Now().Format("2006-01-02")

			if len(parts) >= 5 {
				status = parts[4]
			}
			if len(parts) >= 6 {
				tanggal = parts[5]
			}

			tiket := Tiket{
				ID:       parts[0],
				Nama:     parts[1],
				Kategori: parts[2],
				Harga:    harga,
				Status:   status,
				Tanggal:  tanggal,
			}
			daftarTiket = append(daftarTiket, tiket)

			// Update lastID
			if strings.HasPrefix(parts[0], "TKT") && len(parts[0]) == 6 {
				idNum, err := strconv.Atoi(parts[0][3:])
				if err == nil && idNum > lastID {
					lastID = idNum
				}
			}
		}
	}
}

func simpanAkunKeFile() {
	file, err := os.Create("akun.txt")
	if err != nil {
		fmt.Printf(ERROR+"Error menyimpan akun: %v"+RESET+"\n", err)
		return
	}
	defer file.Close()

	for _, akun := range daftarAkun {
		fmt.Fprintf(file, "%s|%s|%s\n", akun.Username, akun.Password, akun.Role)
	}
}

func muatDataAkun() {
	file, err := os.Open("akun.txt")
	if err != nil {
		// Jika file tidak ada, buat akun default
		daftarAkun = append(daftarAkun, Akun{"admin", "admin123", "admin"})
		daftarAkun = append(daftarAkun, Akun{"user", "user123", "user"})
		simpanAkunKeFile()
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) == 3 {
			akun := Akun{
				Username: parts[0],
				Password: parts[1],
				Role:     parts[2],
			}
			daftarAkun = append(daftarAkun, akun)
		}
	}
}
