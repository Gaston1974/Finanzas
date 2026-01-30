package dao

import (
	"database/sql"
	"fmt"
	fun "hello/src/pkg/apiDatas"
	"log"
	"strconv"
	//_ "github.com/go-sql-driver/mysql"
)

type Causa struct {
	Nro_causa          string
	Caratula           string
	Juzgado            string
	Fiscalia           string
	Magistrado         string
	Preventor          string
	Preventor_auxiliar string
	Provincia_id       string
	Localidad_id       string
	Domicilio          string
	Nro_sgo            string
	Nro_mto            string
	Tipo_delito        string
	Nombre_fantasia    string
	Fecha              string
	Providencia        string
	Estado             string
	IpAdress           string
	Nombre_archivo     string
	Ruta_archivo       string
	Tipo_documento     string
	Tamano             string
	UsuarioId          string
	Nota_causas        string
}

func (l *Causa) Load(pp *sql.Rows) string {

	err := pp.Scan(&l.Nro_causa, &l.Caratula, &l.Juzgado, &l.Fiscalia, &l.Magistrado,
		&l.Preventor, &l.Preventor_auxiliar, &l.Provincia_id, &l.Localidad_id, &l.Domicilio, &l.Nro_sgo, &l.Nro_mto,
		&l.Tipo_delito, &l.Nombre_fantasia, &l.Fecha, &l.Providencia, &l.Estado, &l.IpAdress, &l.Nombre_archivo, &l.Ruta_archivo,
		&l.Tipo_documento, &l.Tamano, &l.UsuarioId, &l.Nota_causas)

	if err != nil {
		msg := "\n Error en la carga del objeto"
		fmt.Printf("%s \n %s", msg, err)
		return msg
	}

	return ""

}

func (p *Causa) Load2(nro_causa string, caratula string, juzgado string, fiscalia string, magistrado string, estado string,
	Preventor string, Preventor_auxiliar string, provincia_id string, localidad_id string, domicilio string, ip string,
	nro_sgo string, nro_mto string, tipo_delito string, nombre_fantasia string, fecha string, providencia string,
	nombre_archivo string, ruta_archivo string, tipo_documento string, tamano string, nota string) {

	p.Nro_causa = nro_causa
	p.Caratula = caratula
	p.Juzgado = juzgado
	p.Fiscalia = fiscalia
	p.Magistrado = magistrado
	p.Nombre_fantasia = nombre_fantasia
	p.Nro_mto = nro_mto
	p.Nro_sgo = nro_sgo
	p.Preventor = Preventor
	p.Preventor_auxiliar = Preventor_auxiliar
	p.Provincia_id = provincia_id
	p.Localidad_id = localidad_id
	p.Domicilio = domicilio
	p.Tipo_delito = tipo_delito
	p.Fecha = fecha
	p.Providencia = providencia
	p.Estado = estado
	p.Nombre_archivo = nombre_archivo
	p.Ruta_archivo = ruta_archivo
	p.Tipo_documento = tipo_documento
	p.Tamano = tamano
	p.Nota_causas = nota

}

func (p *Causa) Load3(nro_causa string) {

	p.Nro_causa = nro_causa

}

func (p *Causa) LoadDB(operacion string, token string, descripcion string, accion string) (int, string) {

	db := fun.Acceso{}

	_, val0, msg := db.SetCliente()

	if val0 != 1 {
		return 0, msg
	}

	var sqlStatement, sqlStatement2, sqlStatement3, sqlStatement4, sqlStatement5, sqlStatement6 string
	var res, res2, res3 sql.Result
	var res4 *sql.Row
	var err, err2, err3, err4 error
	var count int64

	defer db.Cliente.Close()

	// Test the connection to the database
	if err := db.Cliente.Ping(); err != nil {
		log.Println(err)
		return 0, err.Error()
	} else {
		log.Println("\n Successfully Connected")
	}

	msg = "No se ha podido concretar el alta. "
	m := "Alta creada con exito"
	i := 0

	sqlStatement3 = "SELECT max(id) FROM causas;"
	sqlStatement4 = "DELETE from causas WHERE id = ? ;"
	sqlStatement6 = "DELETE from documentos_causa WHERE id = ? ;"

	if operacion == "alta" {

		sqlStatement = "INSERT INTO causas  " +
			" (numero_causa, caratula, juzgado_id, fiscalia_id, a_cargo_del_magistrado, preventor, preventor_auxiliar, " +
			" provincia_id, localidad_id, domicilio, nro_sgo, nro_mto, tipo_delito, nombre_fantasia, " +
			" fecha_llegada, providencia ) VALUES (?, ?, ?, ?, ?, ?, ?, ( SELECT id FROM provincias WHERE nombre = ?), " +
			" ( SELECT id FROM localidades WHERE nombre = ?), ?, ?, ?, ?, ?, ?, ?);"

		res, err = db.Cliente.Exec(sqlStatement, p.Nro_causa, p.Caratula, p.Juzgado, p.Fiscalia, p.Magistrado,
			p.Preventor, p.Preventor_auxiliar, p.Provincia_id, p.Localidad_id, p.Domicilio, p.Nro_sgo, p.Nro_mto, p.Tipo_delito,
			p.Nombre_fantasia, p.Fecha, p.Providencia)

		if p.Nombre_archivo != "" && p.Ruta_archivo != "" && err == nil {
			res4 = db.Cliente.QueryRow(sqlStatement3)
			err4 = res4.Scan(&i)
			if err4 != nil {
				fmt.Printf("\n %s error: %s", msg, err4.Error())
				return 0, msg
			}

			sqlStatement2 = "INSERT INTO documentos_causa  " +
				" (causa_id, nombre_archivo, ruta_archivo, tipo_documento, tamano, subido_por) " +
				" VALUES (?, ?, ?, ?, ?, ?);"

			res2, err2 = db.Cliente.Exec(sqlStatement2, i, p.Nombre_archivo, p.Ruta_archivo, p.Tipo_documento,
				p.Tamano, token)

		}

		if p.Nota_causas != "" && err2 == nil {

			sqlStatement5 = "INSERT INTO notas_causa  " +
				" (causa_id, contenido, creado_por) " +
				" VALUES (?, ?, ?);"

			res3, err3 = db.Cliente.Exec(sqlStatement5, i, p.Nota_causas, token)

		}

	} else {

		// type json struct {
		// 	Caratula string
		// 	Estado   string
		// }

		// var jsonFields json

		// jsonFields.Caratula = p.Caratula
		// jsonFields.Estado = p.Estado

		// _, dat := fun.WriteJson(jsonFields)

		sqlStatement = "INSERT INTO historial_causas  " +
			" (numero_causa, accion, usuario_id, descripcion) " +
			" VALUES (?, ?, ?, ?);"

		res, err = db.Cliente.Exec(sqlStatement, p.Nro_causa, accion, token, descripcion)

	}

	if err != nil {
		fmt.Printf("\n %s error: %s", msg, err.Error())
		return 0, msg
	} else if err2 != nil {
		fmt.Printf("\n %s error: %s", msg, err2.Error())
		db.Cliente.Exec(sqlStatement4, i)
		return 0, msg
	} else if err3 != nil {
		fmt.Printf("\n %s error: %s", msg, err3.Error())
		db.Cliente.Exec(sqlStatement6, i)
		return 0, msg
	} else {

		if res != nil {
			count, _ = res.RowsAffected()
			fmt.Println("\n registros afectados en historial de causas: ", count)
		}
		if res2 != nil {
			count, _ = res2.RowsAffected()
			fmt.Println("\n registros afectados en documentos causas: ", count)
		}
		if res3 != nil {
			count, _ = res3.RowsAffected()
			fmt.Println("\n registros afectados en notas causas: ", count)
		}
		fmt.Printf("\n %s ", m)
	}

	return 1, m
}

func (p *Causa) LoadDBUpdt(nroCausa string, token string, descripcion string) (int, string) { // caso modificacion de precios  -  info map[int]float64

	db := fun.Acceso{}

	_, val0, msg := db.SetCliente()

	if val0 != 1 {
		return 0, msg
	}

	defer db.Cliente.Close()

	// Test the connection to the database
	if err := db.Cliente.Ping(); err != nil {
		log.Println(err)
		return 0, err.Error()
	} else {
		log.Println("\n Successfully Connected")
	}

	msg = "No se ha podido concretar la actualizacion. "
	m := "Actualizacion realizada"

	fis, _ := strconv.Atoi(p.Fiscalia)
	juz, _ := strconv.Atoi(p.Juzgado)

	sqlStatement1 := "UPDATE causas " +
		" SET caratula = '" + p.Caratula + "'," +
		" fiscalia_id =  ? ," +
		" juzgado_id = ? ," +
		" a_cargo_del_magistrado = '" + p.Magistrado + "'," +
		" nombre_fantasia = '" + p.Nombre_fantasia + "'," +
		" nro_mto = '" + p.Nro_mto + "'," +
		" nro_sgo = '" + p.Nro_sgo + "'," +
		" Preventor = '" + p.Preventor + "'," +
		" Preventor_auxiliar = '" + p.Preventor_auxiliar + "'," +
		" provincia_id = ( SELECT id FROM provincias WHERE nombre = '" + p.Provincia_id + "'), " +
		" localidad_id = ( SELECT id FROM localidades WHERE nombre = '" + p.Localidad_id + "'), " +
		" domicilio = '" + p.Domicilio + "'," +
		" tipo_delito = '" + p.Tipo_delito + "'," +
		" fecha_llegada = '" + p.Fecha + "'," +
		" estado = '" + p.Estado + "'," +
		" providencia = '" + p.Providencia + "'" +
		" WHERE numero_causa = ?;"

	res, err := db.Cliente.Exec(sqlStatement1, fis, juz, nroCausa)

	if err != nil {
		fmt.Printf("\n %s error: %s", msg, err.Error())
		return 0, msg

	} else {
		count, _ := res.RowsAffected()
		if count == 0 {
			fmt.Printf("\n %s error: %s cpount: %d", msg, err, count)
			return 0, msg
		}
		fmt.Println("\n registros afectados: ", count)
		fmt.Printf("\n %s ", m)
	}

	p.LoadDB("", token, descripcion, "actualizacion")

	return 1, m

}

func (p *Causa) LoadBaja(tabla string, token string, descripcion string, accion string, estado string) (int, string) { // actualiza informacion del cliente

	db := fun.Acceso{}

	_, val0, msg := db.SetCliente()

	if val0 != 1 {
		return 0, msg
	}

	defer db.Cliente.Close()

	p.LoadDB("", token, descripcion, accion)

	sqlStatement := "UPDATE causas " +
		" SET is_active = 0,  " +
		" estado = ? " +
		" WHERE numero_causa = ? ;"

	msg = "Falla al eliminar la causa -"

	var res1 sql.Result
	var err3 error

	res1, err3 = db.Cliente.Exec(sqlStatement, estado, p.Nro_causa)

	if err3 != nil {

		fmt.Printf("\n %s error: %s", msg, err3.Error())
		return 0, msg

	} else {

		count, _ := res1.RowsAffected()
		if count == 0 {
			fmt.Printf("\n %s cpount: %d", msg, count)
			return 2, msg
		}
		fmt.Println("registros afectados al eliminar la causa: ", count)
		return 1, "Baja realizada con exito"
	}

}

func (p *Causa) Exist() int {

	db := fun.Acceso{}

	_, val0, msg := db.SetCliente()

	if val0 != 1 {
		return 0
	}

	var sqlStatement string
	var res1 *sql.Row
	var err error

	sqlStatement = "SELECT id from causas  " +
		" where numero_causa = ? ;"

	defer db.Cliente.Close()

	var i int

	res1 = db.Cliente.QueryRow(sqlStatement, p.Nro_causa)

	err = res1.Scan(&i)

	if err != nil {
		fmt.Printf("\n %s error: %s", msg, err.Error())
		return 0
	}

	return 1

}

func (u *Causa) Atr() []string {

	var atr = []string{"Nro_causa", "Caratula", "Juzgado", "Fiscalia", "Magistrado", "Preventor", "Preventor_auxiliar", "Provincia_id",
		"Localidad_id", "Domicilio", "Nro_sgo", "Tipo_delito", "Nombre_fantasia", "Fecha", "Providencia", "Estado", "IpAdress", "Nombre_archivo",
		"Ruta_archivo", "Tipo_documento", "Tamano", "UsuarioId", "Nota_causas"}

	return atr

}

func (u *Causa) Get(field string) string {

	f := u.Fecha[0:10]

	switch field {

	case "Nro_causa":
		return u.Nro_causa
	case "Caratula":
		return u.Caratula
	case "Juzgado":
		return u.Juzgado
	case "Fiscalia":
		return u.Fiscalia
	case "Magistrado":
		return u.Magistrado
	case "Preventor":
		return u.Preventor
	case "Preventor_auxiliar":
		return u.Preventor_auxiliar
	case "Provincia_id":
		return u.Provincia_id
	case "Localidad_id":
		return u.Localidad_id
	case "Domicilio":
		return u.Domicilio
	case "Nro_sgo":
		return u.Nro_sgo
	case "Tipo_delito":
		return u.Tipo_delito
	case "Nombre_fantasia":
		return u.Nombre_fantasia
	case "Fecha":
		return f
	case "Providencia":
		return u.Providencia
	case "Estado":
		return u.Estado
	case "IpAdress":
		return u.IpAdress
	case "Nombre_archivo":
		return u.Nombre_archivo
	case "Ruta_archivo":
		return u.Ruta_archivo
	case "Tipo_documento":
		return u.Tipo_documento
	case "Tamano":
		return u.Tamano
	case "UsuarioId":
		return u.UsuarioId
	case "Nota_causas":
		return u.Nota_causas
	default:
		return ""

	}

}
