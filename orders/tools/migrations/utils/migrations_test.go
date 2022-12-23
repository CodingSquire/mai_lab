package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/innerave/mai_lab/orders/internal/db"
	"github.com/innerave/mai_lab/orders/internal/dotenv"
)

func TestParseNoArgs(t *testing.T) {
	os.Args = []string{"migrations"}
	_, _, _, err := ParseArgs()
	if err.Error() != "Not enough arguments" {
		t.Errorf("Expected error, got %s", err)
	}
}

func TestParseUpArg(t *testing.T) {
	os.Args = []string{"migrations", "up"}
	_, _, _, err := ParseArgs()
	if err == nil {
		t.Errorf("Expected error, got %s", err)
	}
}

func TestParseUpArgPath(t *testing.T) {
	tmp := os.TempDir()
	os.Args = []string{"migrations", "up", tmp}
	cmd, path, _, err := ParseArgs()
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if cmd != "up" {
		t.Errorf("Expected cmd to be up, got %s", cmd)
	}
	if path != tmp {
		t.Errorf("Expected path to be %s, got %s", tmp, path)
	}
}

func TestParsePersisrFlagArg(t *testing.T) {
	tmp := os.TempDir()
	os.Args = []string{"migrations", "up", tmp, "--persist"}
	cmd, path, flags, err := ParseArgs()
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if cmd != "up" {
		t.Errorf("Expected cmd to be up, got %s", cmd)
	}
	if path != tmp {
		t.Errorf("Expected path to be %s, got %s", tmp, path)
	}
	if flags.Persist != true {
		t.Errorf("Expected persist to be true, got %t", flags.Persist)
	}
}

func TestRun(t *testing.T) {
	dotenv.Config()
	db, err := db.Init()
	if err != nil {
		t.Errorf("Expected db to be initialized, got %s", err)
	}
	defer db.Close()

	tmp, err := os.MkdirTemp("", "migrations")
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	defer os.RemoveAll(tmp)

	t.Run("RunUpTest", func(t *testing.T) {
		err = os.WriteFile(filepath.Join(tmp, "1_up.sql"), []byte("CREATE TABLE test (id int)"), 0644)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		err = Run_up(tmp, db)
		if err != nil {
			t.Errorf("Expected run up to be successful, got %s", err)
		}

		_, err = db.Exec("INSERT INTO test (id) VALUES (1)")
		if err != nil {
			t.Errorf("Expected value to be inserted, got %s", err)
		}

		_, err = db.Exec("DROP TABLE test")
		if err != nil {
			t.Errorf("Expected table to be dropped, got %s", err)
		}
	})

	t.Run("RunDownTest", func(t *testing.T) {
		err = os.WriteFile(filepath.Join(tmp, "1_down.sql"), []byte("DROP TABLE test"), 0644)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		_, err = db.Exec("CREATE TABLE test (id int)")
		if err != nil {
			t.Errorf("Expected table to be created, got %s", err)
		}

		err = Run_down(tmp, db)
		if err != nil {
			t.Errorf("Expected run down to be successful, got %s", err)
		}

		_, err = db.Exec("INSERT INTO test (id) VALUES (1)")
		if err == nil {
			t.Errorf("Expected value to not be inserted, got %s", err)
		}
	})
}
