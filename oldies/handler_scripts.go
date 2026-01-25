package oldies

/*
import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	funciones "hello/src/pkg/apiDatas"
	scr "hello/src/pkg/scripts"
	"io"
	"net/http"
)

func HandlerCryptos(w http.ResponseWriter, r *http.Request) {

	var err error

	type solicitante struct {
		Script string
	}

	sol := solicitante{}

	b, err := io.ReadAll(r.Body)
	fmt.Println(string(b))
	if err != nil {
		msg := "Falla interna al leer el body del mensaje"
		fmt.Printf("%s", msg)
		funciones.ResponseWithJSON(w, 400, msg)
		return
	}

	err = json.Unmarshal(b, &sol)
	if err != nil {
		msg := "\nFalla durante parseo de parametros del Request: "
		fmt.Printf("%s", msg)
		funciones.ResponseWithJSON(w, 400, msg)
		return
	}

	switch sol.Script {

	case "GALICIA.pdf":

		scr.Info1()

	default:

		fmt.Println("No se encuentra el script ingresado")

	}

	//fmt.Printf("\nresultado: %s", msg)
	//funciones.ResponseWithJSON(w, 200, msg)
	//funciones.Log(msg, ".log.txt")

	//os.Remove("./Bancos/" + bank)
	//os.Remove("./Convertidos/" + name)
	//os.Remove("./Bancos/" + "file.pdf")
	//os.Remove(xlsPath)

}

// *********************************************************************


*/
