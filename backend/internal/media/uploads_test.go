package media

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

const (
	a = "0123456789abcdef0123456789abcdef"
	b = "fedcba9876543210fedcba9876543210"
)

func TestExtractNames(t *testing.T) {
	cases := []struct {
		name  string
		texts []string
		want  []string
	}{
		{"sampul", []string{"/uploads/" + a + ".png"}, []string{a + ".png"}},
		{"markdown", []string{"Teks ![x](/uploads/" + b + ".webp) lagi"}, []string{b + ".webp"}},
		{"unik dan berurutan", []string{"![](/uploads/" + a + ".jpg) ![](/uploads/" + b + ".gif)", "/uploads/" + a + ".jpg"}, []string{a + ".jpg", b + ".gif"}},
		{"situs lain diabaikan", []string{"https://contoh.com/uploads/" + a + ".png"}, nil},
		{"nama tidak valid diabaikan", []string{"/uploads/../../etc/passwd", "/uploads/abc.png", "/uploads/" + a + ".svg", "/uploads/" + a[:31] + ".png"}, nil},
		{"huruf besar diabaikan", []string{"/uploads/" + "ABCDEF0123456789ABCDEF0123456789" + ".png"}, nil},
		{"kosong", []string{"", "teks biasa"}, nil},
	}
	for _, c := range cases {
		if got := ExtractNames(c.texts...); !reflect.DeepEqual(got, c.want) {
			t.Errorf("%s: dapat %v, mau %v", c.name, got, c.want)
		}
	}
}

func TestRemove(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, a+".png")
	if err := os.WriteFile(file, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	other := filepath.Join(dir, b+".png")
	os.WriteFile(other, []byte("x"), 0o644)

	if err := Remove(dir, a+".png"); err != nil {
		t.Fatalf("hapus: %v", err)
	}
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		t.Error("file seharusnya terhapus")
	}
	if _, err := os.Stat(other); err != nil {
		t.Error("file lain tidak boleh ikut terhapus")
	}
	if err := Remove(dir, a+".png"); err != nil {
		t.Errorf("file yang sudah tidak ada harus dianggap berhasil: %v", err)
	}
	for _, bad := range []string{"../x.png", "..", "", "a/b.png", b + ".png/../" + a + ".png", a + ".svg"} {
		if err := Remove(dir, bad); err == nil {
			t.Errorf("nama %q seharusnya ditolak", bad)
		}
	}
	if _, err := os.Stat(other); err != nil {
		t.Error("file lain terhapus oleh nama berbahaya")
	}
}
