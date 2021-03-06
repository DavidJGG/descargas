package main

import (
	"context"
	"fmt"
	"log"
	"time"
	adminDatos "./conexion"
	"google.golang.org/grpc"

	"io/ioutil"
	"net/http"
	"encoding/json"
	"strconv"
	str "strings"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/datos", datos)
	log.Printf("listening on port 9000")
	log.Fatal(http.ListenAndServe(":9000", mux))	
}

type Resultado struct{
	msg string
	msg2 string
}

func mandarPorgRPC(nombre string, depto string, edad string, fc string, est string) (*Resultado, error) {
	time.Sleep(2000 * time.Millisecond)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("cnt-python-svc:50051", grpc.WithInsecure())
	//conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("No se puedo conectar: %s", err)
		return nil,err;
	}
	defer conn.Close()

	c := adminDatos.NewServicioDatosClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err2 := c.ObtenerDatos(ctx, &adminDatos.Datos{
		Nombre:nombre,
		Departamento:depto,
		Edad:edad,
		FormaContagio:fc,
		Estado:est,
	})

	if err2 != nil {
		fmt.Println("Error: %s", err2)
		return nil,err2;
	}
	fmt.Println("Response from server: ", response)
	var rr = Resultado{}
	rr.msg=strconv.FormatBool(response.Enviado)
	rr.msg2=response.Msg
	return &rr,nil;
}

type datosRecibidos struct{
	Nombre string `json:"name"`
	Departamento string `json:"depto"`
	Edad int `json:"age"`
	FormaContagio string `json:"form"`
	Estado string `json:"state"`
}

func datos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//var results []string
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading request body", http.StatusInternalServerError)
		}
		//results = append(results, string(body))
		dat := datosRecibidos{}
		err = json.Unmarshal(body, &dat)
		if err!=nil{
			fmt.Println("Error al convertir los datos: %s", err)
			w.Write([]byte(err.Error()))
		}
		fmt.Println(dat.Nombre)
		//fmt.Println(results)
		
		rr, errGRPC:=mandarPorgRPC(
			normalizar(dat.Nombre),
			normalizar(dat.Departamento), 
			strconv.FormatInt(int64(dat.Edad), 10), 
			normalizar(dat.FormaContagio), 
			normalizar(dat.Estado))

		if errGRPC!=nil{
			w.Write([]byte(errGRPC.Error()))
		}else{
			if rr.msg=="true"{
				w.Write([]byte("Enviado -> "+rr.msg2))
			}else{
				w.Write([]byte("No se pudo enviar"))
			}			
		}
		

		//fmt.Println("-------------------------------------------")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}


func normalizar(txt string) string{
	txt=str.Replace(txt, "á", "a", -1);
	txt=str.Replace(txt, "é", "e", -1);
	txt=str.Replace(txt, "í", "i", -1);
	txt=str.Replace(txt, "ó", "o", -1);
	txt=str.Replace(txt, "ú", "u", -1);
	
	txt=str.Replace(txt, "Á", "A", -1);
	txt=str.Replace(txt, "É", "E", -1);
	txt=str.Replace(txt, "Í", "I", -1);
	txt=str.Replace(txt, "Ó", "O", -1);
	txt=str.Replace(txt, "Ú", "U", -1);

	txt=str.Replace(txt, "ä", "a", -1);
	txt=str.Replace(txt, "ë", "e", -1);
	txt=str.Replace(txt, "ï", "i", -1);
	txt=str.Replace(txt, "ö", "o", -1);
	txt=str.Replace(txt, "ü", "u", -1);
	
	txt=str.Replace(txt, "Ä", "A", -1);
	txt=str.Replace(txt, "Ë", "E", -1);
	txt=str.Replace(txt, "Ï", "I", -1);
	txt=str.Replace(txt, "Ö", "O", -1);
	txt=str.Replace(txt, "Ü", "U", -1);

	txt=str.Replace(txt, "ñ", "ni", -1);
	txt=str.Replace(txt, "Ñ", "NI", -1);
	return txt;
}
