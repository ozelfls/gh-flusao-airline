package app

import (
	"context"
	"strings"
	"testing"
)

func TestPresentationBootstrapAndIntegrity(t *testing.T) {
	t.Setenv("DEMO_MODE", "true")
	a, err := New(t.TempDir() + "/demo.db")
	if err != nil {
		t.Fatal(err)
	}
	defer a.DB.Close()
	ctx := context.Background()
	if err = a.Bootstrap(ctx); err != nil {
		t.Fatal(err)
	}
	for table, want := range map[string]int{"voo": 40, "passageiro": 51, "bilhete": 255, "escala": 46, "aeronave": 14, "aeroporto": 12} {
		if got := count(ctx, a.DB, "SELECT count(*) FROM "+table); got != want {
			t.Fatalf("%s: got %d want %d", table, got, want)
		}
	}
	if err = a.Bootstrap(ctx); err != nil {
		t.Fatal(err)
	}
	if count(ctx, a.DB, "SELECT count(*) FROM bilhete") != 255 {
		t.Fatal("bootstrap duplicated tickets")
	}
	checks := []string{
		`SELECT count(*) FROM pragma_foreign_key_check`,
		`SELECT count(*) FROM bilhete b JOIN voo v ON v.id=b.voo_id JOIN aeronave a ON a.id=v.aeronave_id WHERE b.assento>a.capacidade OR b.origem_id<>v.origem_id OR b.destino_id<>v.destino_id`,
		`SELECT count(*) FROM voo v JOIN aeronave a ON a.id=v.aeronave_id WHERE v.empresa_id<>a.empresa_id`,
		`SELECT count(*) FROM voo a JOIN voo b ON a.id<b.id AND a.aeronave_id=b.aeronave_id AND a.partida<b.chegada AND a.chegada>b.partida WHERE a.situacao IN('Agendado','Em Voo') AND b.situacao IN('Agendado','Em Voo')`,
		`SELECT count(*) FROM bilhete WHERE checkin_em<emitido_em OR embarque_em<checkin_em OR cancelado_em<emitido_em`,
	}
	for _, q := range checks {
		if n := count(ctx, a.DB, q); n != 0 {
			t.Fatalf("integrity count %d: %s", n, q)
		}
	}
	if count(ctx, a.DB, `SELECT count(*) FROM bilhete WHERE voo_id=1032 AND situacao<>'Cancelado'`) != 8 {
		t.Fatal("sold out scenario missing")
	}
	if count(ctx, a.DB, `SELECT count(*) FROM escala WHERE voo_id=1004`) != 2 {
		t.Fatal("multi stop scenario missing")
	}
	if err = a.Sync(ctx); err != nil {
		t.Fatal(err)
	}
	for _, q := range advancedCatalog {
		if q.Kind != "read" {
			continue
		}
		result := runAdvancedTest(t, a, q.ID, nil)
		if len(result.Raw) == 0 {
			t.Fatalf("demo query %d has no example rows", q.ID)
		}
	}
	for _, table := range []string{"pais", "estado", "cidade", "empresa", "aeroporto", "aeronave", "voo", "escala", "passageiro", "bilhete", "auditoria", "manutencao"} {
		rs, err := rows(ctx, a.DB, "SELECT * FROM "+table)
		if err != nil {
			t.Fatal(err)
		}
		for _, row := range rs {
			for _, value := range row {
				if strings.ContainsAny(value, "\u2014\u2013\u2022\u00b7") {
					t.Fatalf("forbidden punctuation in %s: %s", table, value)
				}
			}
		}
	}
	for _, role := range []int64{1, 2, 3} {
		cookie := session(t, a, role)
		for _, path := range []string{"/", "/?entity=bilhete", "/detail/voo/1004", "/print/1000"} {
			w := request(a, "GET", path, cookie, nil)
			if w.Code != 200 {
				t.Fatalf("role %d path %s: %d", role, path, w.Code)
			}
			if strings.ContainsAny(w.Body.String(), "\u2014\u2013\u2022\u00b7") {
				t.Fatalf("forbidden punctuation in %s", path)
			}
		}
	}
}

func TestPresentationProductionAndRollback(t *testing.T) {
	ctx := context.Background()
	a := testApp(t)
	t.Setenv("DEMO_MODE", "false")
	if err := a.Bootstrap(ctx); err != nil {
		t.Fatal(err)
	}
	if count(ctx, a.DB, "SELECT count(*) FROM voo") != 6 {
		t.Fatal("production mode changed data")
	}
	t.Setenv("DEMO_MODE", "true")
	if _, err := a.DB.Exec(`INSERT INTO pais(id,nome,continente,sigla) VALUES(1000,'Conflito','Teste','ZZ')`); err != nil {
		t.Fatal(err)
	}
	if err := a.SeedPresentation(ctx); err == nil {
		t.Fatal("expected collision to roll back")
	}
	if count(ctx, a.DB, "SELECT count(*) FROM voo") != 6 {
		t.Fatal("partial seed committed")
	}
}
