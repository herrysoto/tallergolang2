package datos

import (
	bean "taller/entidad"
	utils "taller/utils"
	"time"
	//"fmt"
	"log"
	"github.com/gin-gonic/gin"
)


//Getfirstcombo obtener datos del primer combo
func Getfirstcombo(c *gin.Context) ([]bean.Servicio,error) {

	var descod []bean.Servicio

	_,err := utils.DbmapOracle.Select(&descod, "SELECT CHRCODIGOSERVICIO , VCHDESCRIPCION  FROM SERVICIO ORDER BY CHRCODIGOSERVICIO")
	if err != nil {
		log.Println("error ")
		log.Println(err)
		return nil,err
	}
/*	 for x, i := range descod {
		log.Printf("    %d: %v\n", x, i)
		}*/
		return descod,nil
}

//validación para la duplicacion de código al insertar un nuevo item
func Getcodigovalidacion(c *gin.Context)([]bean.OperacionServicio,error){
	Codopera := c.Params.ByName("Codopera")
	Codoperaservicio := c.Params.ByName("Codoperaservicio")
	var codigovalidacion []bean.OperacionServicio

	_,err := utils.DbmapOracle.Select(&codigovalidacion, `SELECT CHRCODIGOOPERACIONSERVICIO FROM OPERACIONSERVICIOS WHERE VCHCODIGOOPERACION =:Codopera AND CHRCODIGOOPERACIONSERVICIO =:Codoperaservicio`,Codopera,Codoperaservicio)
	 log.Println(len(codigovalidacion))
	   if (len(codigovalidacion) == 0) {codigovalidacion = nil}
	   log.Println(codigovalidacion)

	
	if err != nil {
		log.Println("error")
		log.Println(err)
		return nil,err
	}
	return codigovalidacion,nil
}

//obtener las horas hombre 
func GetHorasHombres(c *gin.Context)([]bean.HorasHombre,error){

	var hh []bean.HorasHombre

	_,err := utils.DbmapOracle.Select(&hh, "SELECT ROUND(NUMHORASHOMBRE,2) NUMHORASHOMBRE FROM (SELECT * FROM HORASHOMBRE WHERE DTEFECHA IS NOT NULL ORDER BY DTEFECHA DESC) WHERE ROWNUM <= 1")
	if err != nil {
		log.Println("error")
		log.Println(err)
		return nil,err
	}
	return hh,nil
}



//actualizar las HorasHombre, esto va primero y luego viene el insert(post)
func PutHorasHombre(c *gin.Context)(bean.UpdateHH,error){
	var puthh bean.UpdateHH
	c.Bind(&puthh)
	if (puthh.NUMPRECIOOFICIAL == 0){
		c.JSON(422, gin.H{"error": "fields are empty"})
		return puthh, nil
	}
	_, err := utils.DbmapOracle.Exec(`UPDATE OPERACIONSERVICIOS SET NUMPRECIOSUGERIDO = (NVL(NUMHORASHOMBRE,0) * :NUMPRECIOOFICIAL),
              NUMPRECIOOFICIAL =:NUMPRECIOOFICIAL WHERE TRIM(VCHCODIGOOPERACION) NOT IN('7B')`,puthh.NUMPRECIOOFICIAL,puthh.NUMPRECIOOFICIAL)
	if err != nil {
		log.Println("error")
		log.Println(err)
	}
	return puthh, nil
}

func PostHorasHombres(c *gin.Context)(bean.HorasHombre,error){
	var inserthh bean.HorasHombre
	c.Bind(&inserthh)
	// Validations
	if inserthh.NUMHORASHOMBRE == 0 {
		c.JSON(422, gin.H{"error": "fields are empty"})
		return inserthh,nil
	}
	_, err := utils.DbmapOracle.Exec(`INSERT INTO HORASHOMBRE(NUMCODIGO, NUMHORASHOMBRE, DTEFECHA)
          VALUES(:NUMCODIGO, :NUMHORASHOMBRE, :DTEFECHA)`,inserthh.NUMCODIGO,inserthh.NUMHORASHOMBRE,time.Now())
	if err != nil {
		log.Println("error")
		log.Println(err)
	}
	return inserthh, nil
}

func Getnumcodigo(c *gin.Context)([]bean.HorasHombre,error){
	var numcod []bean.HorasHombre

	_,err := utils.DbmapOracle.Select(&numcod, "SELECT NVL(MAX(NUMCODIGO),0) + 1  NUMCODIGO FROM HORASHOMBRE")
	if err != nil {
		log.Println("error")
		log.Println(err)
		return nil,err
	}
	return numcod,nil
}

//prueba para que me traiga el input por default
func Getfirstcombo2(c *gin.Context) ([]bean.Servicio,error) {
		var codiserv []bean.Servicio

	_, err := utils.DbmapOracle.Select(&codiserv, "SELECT VCHDESCRIPCION  FROM SERVICIO where CHRCODIGOSERVICIO ='01' ORDER BY CHRCODIGOSERVICIO")
	if err != nil {
		log.Println("error ")
		log.Println(err)
		return nil,err
	}
/*	 for x, i := range descod {
		log.Printf("    %d: %v\n", x, i)
		}*/
		return codiserv,nil
}


func Getsecondcombo(c *gin.Context) ([]bean.Servicio,error) {
		CodServicio := c.Params.ByName("CodServicio")
		var codserv []bean.Servicio
		_,err := utils.DbmapOracle.Select(&codserv,`SELECT VCHCODIGOOPERACION,VCHDESCRIPCION,CHRCODIGOSERVICIO  FROM OPERACION WHERE CHRCODIGOSERVICIO=:CodServicio ORDER BY VCHCODIGOOPERACION`,CodServicio)
		if err != nil{
			log.Println("error ")
			log.Println(err)
			
			
		}
		return codserv,nil
}

func Getgrid(c *gin.Context) ([]bean.OperacionServicio,error){
		Operserv := c.Params.ByName("Operserv")
		var opser []bean.OperacionServicio
		_,err := utils.DbmapOracle.Select(&opser,`SELECT nvl(NUMNROPIEZA,0) NUMNROPIEZA,VCHNROTRABAJO,VCHCODIGOOPERACION || '-' || CHRCODIGOOPERACIONSERVICIO  CODIGO,VCHCODIGOOPERACION,CHRCODIGOOPERACIONSERVICIO, VCHDESCRIPCION, ROUND(NUMHORASHOMBRE,2) NUMHORASHOMBRE,
                   ROUND(NUMPRECIOSUGERIDO,2) NUMPRECIOSUGERIDO, ROUND(NUMDESCUENTO,2) NUMDESCUENTO, NVL(ROUND(NUMPRECIOESTIMADO + (ROUND(NUMPRECIOESTIMADO,2) * 0.18),2),0)  NUMPRECIOESTIMADO ,ROUND(NUMPRECIOOFICIAL,2) NUMPRECIOOFICIAL,
                   nvl(NUMMANOOBRA,0) NUMMANOOBRA,nvl(NUMREPUESTO,0) NUMREPUESTO,nvl(NUMTOTAL,0) NUMTOTAL
            FROM OPERACIONSERVICIOS
            WHERE VCHCODIGOOPERACION =: Operserv
            ORDER BY CHRCODIGOOPERACIONSERVICIO`,Operserv)
		if err !=nil{
			log.Println("error")
			log.Println(err)
		} 
		return opser,nil
}

func BuscarOperServiciosParam(c *gin.Context)([]bean.OperacionServicio,error){
	param := c.Params.ByName("param")
		var opserparam []bean.OperacionServicio
		_,err := utils.DbmapOracle.Select(&opserparam,`SELECT nvl(NUMNROPIEZA,0) NUMNROPIEZA,VCHCODIGOOPERACION || '-' || CHRCODIGOOPERACIONSERVICIO CODIGO,VCHNROTRABAJO,VCHCODIGOOPERACION,CHRCODIGOOPERACIONSERVICIO, VCHDESCRIPCION,ROUND(NUMHORASHOMBRE,2) NUMHORASHOMBRE ,
              ROUND(NUMPRECIOSUGERIDO,2) NUMPRECIOSUGERIDO , ROUND(NUMDESCUENTO,2) NUMDESCUENTO,NVL(ROUND(NUMPRECIOESTIMADO + (ROUND(NUMPRECIOESTIMADO,2) * 0.18),2),0)  NUMPRECIOESTIMADO,
              ROUND(NUMPRECIOOFICIAL,2) NUMPRECIOOFICIAL,
            nvl(NUMMANOOBRA,0) NUMMANOOBRA,nvl(NUMREPUESTO,0) NUMREPUESTO,nvl(NUMTOTAL,0) NUMTOTAL
        FROM OPERACIONSERVICIOS
        WHERE TRIM(UPPER(VCHDESCRIPCION)) LIKE '%'|| TRIM(UPPER(:param)) ||'%'`,param)
		if err !=nil{
			log.Println("error")
			log.Println(err)
		} 
		return opserparam,nil
}

func LISTAR_OPERSERVCONTENIDOS(c *gin.Context)([]bean.OperacionServiciosContenidos,[]bean.OperacionServicio){
	param := c.Params.ByName("param")
	param1 := c.Params.ByName("param1")
		var opserparam1 []bean.OperacionServiciosContenidos
		_,err := utils.DbmapOracle.Select(&opserparam1,`SELECT nvl(OPCS.NUMNROPIEZA,0) NUMNROPIEZA, OPC.VCHCODIGOOPERACION || '-' || OPC.CHRCODIGOOPERACIONSERVICIO  CODIGO,OPCS.VCHNROTRABAJO,TRIM(OPC.VCHCODIGOOPERACION) VCHCODIGOOPERACION,TRIM(OPC.CHRCODIGOOPERACIONSERVICIO) CHRCODIGOOPERACIONSERVICIO,OPCS.VCHDESCRIPCION,
                OPCS.NUMPRECIOSUGERIDO,OPCS.NUMHORASHOMBRE,OPCS.NUMDESCUENTO,nvl(OPCS.NUMMANOOBRA,0) NUMMANOOBRA,nvl(OPCS.NUMREPUESTO,0) NUMREPUESTO,nvl(OPCS.NUMTOTAL,0) NUMTOTAL        
        FROM    OPERACIONSERVICIOSCONTENIDOS OPC,OPERACIONSERVICIOS OPCS
        WHERE   OPC.VCHCODIGOOPERACION = OPCS.VCHCODIGOOPERACION
        AND     OPC.CHRCODIGOOPERACIONSERVICIO = OPCS.CHRCODIGOOPERACIONSERVICIO
        AND     TRIM(OPC.VCHOPERACIONMAESTRA) = TRIM(:param)||''||TRIM(:param1)`,param,param1)
		if err !=nil{
			log.Println("error")
			log.Println(err)
		} 
		return opserparam1,nil
}

// func PutOperaServi(c *gin.Context)(bean.UpdateOS,error){
// 	putOpeSer := c.Params.ByName("putOpeSer")
// 	putOpeSer1 := c.Params.ByName("putOpeSer1")
// 	sysdate := time.Now().Format("2017-01-02 15:04:05-0700")
// 	var insop bean.UpdateOS
// 	c.Bind(&insop)
// 	if (insop.VCHCODIGOOPERACION == "") && (insop.CHRCODIGOOPERACIONSERVICIO =="") {
// 		c.JSON(422, gin.H{"error": "fields are empty"})
// 		return insop,nil
// 	}
// 	_,err := utils.DbmapOracle.Exec(`UPDATE OPERACIONSERVICIOS SET VCHDESCRIPCION=:VCHDESCRIPCION, NUMPRECIOSUGERIDO=:NUMPRECIOSUGERIDO,
// 									NUMHORASHOMBRE=:NUMHORASHOMBRE,NUMDESCUENTO=:NUMDESCUENTO,
//             						DTEMODIFICACION =:DTEMODIFICACION, NUMPRECIOESTIMADO =:NUMPRECIOESTIMADO
//         							WHERE UPPER(TRIM(CHRCODIGOOPERACIONSERVICIO)) =: UPPER(TRIM(putOpeSer)) AND  UPPER(TRIM(VCHCODIGOOPERACION)) =: UPPER(TRIM(putOpeSer1))`,
// 		insop.VCHDESCRIPCION,insop.NUMPRECIOSUGERIDO,insop.NUMHORASHOMBRE,insop.NUMDESCUENTO,sysdate,insop.NUMPRECIOESTIMADO,putOpeSer,putOpeSer1)
// 	if err !=nil{
// 			log.Println("error")
// 			log.Println(err)
// 		} 
// 		return insop,nil
// }

func PutOperaServi(c *gin.Context) (bean.UpdateOS, error) {
	// putOpeSer := c.Params.ByName("putOpeSer")
	// putOpeSer1 := c.Params.ByName("putOpeSer1")
	var insop bean.UpdateOS
	c.Bind(&insop)
	// sysdate := time.Now().Format("2017-01-02 15:04:05-0700")
	if (insop.VCHCODIGOOPERACION == "") && (insop.CHRCODIGOOPERACIONSERVICIO == "") {
		c.JSON(422, gin.H{"error": "fields are empty"})
		return insop, nil
	}
	_, err := utils.DbmapOracle.Exec(`UPDATE OPERACIONSERVICIOS SET VCHNROTRABAJO=:VCHNROTRABAJO,VCHDESCRIPCION=:VCHDESCRIPCION, NUMPRECIOSUGERIDO=:NUMPRECIOSUGERIDO,
									NUMHORASHOMBRE=:NUMHORASHOMBRE,NUMDESCUENTO=:NUMDESCUENTO,
            						DTEMODIFICACION =:DTEMODIFICACION, NUMTOTAL =:NUMTOTAL
        							WHERE UPPER(TRIM(CHRCODIGOOPERACIONSERVICIO)) = UPPER(TRIM(:putOpeSer)) AND  UPPER(TRIM(VCHCODIGOOPERACION)) = UPPER(TRIM(:putOpeSer1))`,insop.VCHNROTRABAJO,

		insop.VCHDESCRIPCION, insop.NUMPRECIOSUGERIDO, insop.NUMHORASHOMBRE, insop.NUMDESCUENTO, time.Now(), insop.NUMTOTAL, insop.CHRCODIGOOPERACIONSERVICIO, insop.VCHCODIGOOPERACION)
	if err != nil {
		log.Println("error")
		log.Println(err)
	}
	return insop, nil
}

func Getnumcodigoop(c *gin.Context)([]bean.Numcodigo,error){
	var numcodop []bean.Numcodigo

	_,err := utils.DbmapOracle.Select(&numcodop,"SELECT NVL(MAX(NUMCODIGO),0) + 1  NUMCODIGO FROM operacionservicios")
	if err != nil {
		log.Println("error")
		log.Println(err)
		return nil,err
	}
	return numcodop,nil
}

func Postgrid(c *gin.Context)(bean.InsertOS,error){
	var insertOS bean.InsertOS
	c.Bind(&insertOS)
	// Validations
	if insertOS.VCHCODIGOOPERACION == "" {
		c.JSON(422, gin.H{"error": "fields are empty"})
		return insertOS,nil
	}
	if insertOS.VCHCODIGOOPERACION =="1Z"{
	_, err := utils.DbmapOracle.Exec(`INSERT INTO OPERACIONSERVICIOS (CHRCODIGOOPERACIONSERVICIO,VCHCODIGOOPERACION,VCHDESCRIPCION,
                                       NUMPRECIOSUGERIDO,NUMPRECIOOFICIAL,NUMHORASHOMBRE,NUMDESCUENTO,
                                       CHRESTADO,DTECREACION,NUMTIPOSERVICIO,DTEMODIFICACION,NUMPRECIOESTIMADO,NUMNROPIEZA,VCHNROTRABAJO,NUMCODIGO,NUMMANOOBRA,NUMREPUESTO,NUMTOTAL)
        values (:CHRCODIGOOPERACIONSERVICIO,:VCHCODIGOOPERACION,:VCHDESCRIPCION,:NUMPRECIOSUGERIDO,:NUMPRECIOOFICIAL,:NUMHORASHOMBRE,:NUMDESCUENTO,:CHRESTADO,
               :DTECREACION,:NUMTIPOSERVICIO,:DTEMODIFICACION,:NUMPRECIOESTIMADO,:NUMNROPIEZA,:VCHNROTRABAJO,:NUMCODIGO,:NUMMANOOBRA,:NUMREPUESTO,:NUMTOTAL)`, 
			   insertOS.CHRCODIGOOPERACIONSERVICIO, insertOS.VCHCODIGOOPERACION, insertOS.VCHDESCRIPCION,
			    insertOS.NUMPRECIOSUGERIDO,insertOS.NUMPRECIOOFICIAL, insertOS.NUMHORASHOMBRE, insertOS.NUMDESCUENTO,0,time.Now(),1,time.Now()/*:DTEMODIFICACION*/,
				4/*:NUMPRECIOESTIMADO*/,0/*:NUMNROPIEZA*/,insertOS.VCHNROTRABAJO/*:VCHNROTRABAJO*/,insertOS.NUMCODIGO/*:NUMCODIGO*/,0/*:NUMMANOOBRA*/,0/*:NUMREPUESTO*/,insertOS.NUMTOTAL/*:NUMTOTAL*/)
				
		if err != nil {
		log.Println("error")
		log.Println(err)
	}}
	if insertOS.VCHCODIGOOPERACION =="1W"{
	_, err := utils.DbmapOracle.Exec(`INSERT INTO OPERACIONSERVICIOS (CHRCODIGOOPERACIONSERVICIO,VCHCODIGOOPERACION,VCHDESCRIPCION,
                                       NUMPRECIOSUGERIDO,NUMPRECIOOFICIAL,NUMHORASHOMBRE,NUMDESCUENTO,
                                       CHRESTADO,DTECREACION,NUMTIPOSERVICIO,DTEMODIFICACION,NUMPRECIOESTIMADO,NUMNROPIEZA,VCHNROTRABAJO,NUMCODIGO,NUMMANOOBRA,NUMREPUESTO,NUMTOTAL)
        values (:CHRCODIGOOPERACIONSERVICIO,:VCHCODIGOOPERACION,:VCHDESCRIPCION,:NUMPRECIOSUGERIDO,:NUMPRECIOOFICIAL,:NUMHORASHOMBRE,:NUMDESCUENTO,:CHRESTADO,
               :DTECREACION,:NUMTIPOSERVICIO,:DTEMODIFICACION,:NUMPRECIOESTIMADO,:NUMNROPIEZA,:VCHNROTRABAJO,:NUMCODIGO,:NUMMANOOBRA,:NUMREPUESTO,:NUMTOTAL)`, insertOS.CHRCODIGOOPERACIONSERVICIO, insertOS.VCHCODIGOOPERACION, insertOS.VCHDESCRIPCION,
			    insertOS.NUMPRECIOSUGERIDO,insertOS.NUMPRECIOOFICIAL, insertOS.NUMHORASHOMBRE, insertOS.NUMDESCUENTO,0,time.Now(),3,
				time.Now()/*:DTEMODIFICACION*/,4/*:NUMPRECIOESTIMADO*/,0/*:NUMNROPIEZA*/,insertOS.VCHNROTRABAJO/*:VCHNROTRABAJO*/,insertOS.NUMCODIGO/*:NUMCODIGO*/,0/*:NUMMANOOBRA*/,0/*:NUMREPUESTO*/,insertOS.NUMTOTAL/*:NUMTOTAL*/)

		if err != nil {
		log.Println("error")
		log.Println(err)
	}}
	if insertOS.VCHCODIGOOPERACION =="1X"{
	_, err := utils.DbmapOracle.Exec(`INSERT INTO OPERACIONSERVICIOS (CHRCODIGOOPERACIONSERVICIO,VCHCODIGOOPERACION,VCHDESCRIPCION,NUMHORASHOMBRE,
                                       NUMPRECIOSUGERIDO,NUMPRECIOOFICIAL,NUMDESCUENTO,
                                       CHRESTADO,DTECREACION,NUMTIPOSERVICIO,DTEMODIFICACION,NUMPRECIOESTIMADO,NUMNROPIEZA,VCHNROTRABAJO,NUMCODIGO,NUMMANOOBRA,NUMREPUESTO,NUMTOTAL)
        values (:CHRCODIGOOPERACIONSERVICIO,:VCHCODIGOOPERACION,:VCHDESCRIPCION,:NUMHORASHOMBRE,:NUMPRECIOSUGERIDO,:NUMPRECIOOFICIAL,:NUMDESCUENTO,:CHRESTADO,
               :DTECREACION,:NUMTIPOSERVICIO,:DTEMODIFICACION,:NUMPRECIOESTIMADO,:NUMNROPIEZA,:VCHNROTRABAJO,:NUMCODIGO,:NUMMANOOBRA,:NUMREPUESTO,:NUMTOTAL)`, insertOS.CHRCODIGOOPERACIONSERVICIO, insertOS.VCHCODIGOOPERACION, insertOS.VCHDESCRIPCION, insertOS.NUMHORASHOMBRE,
			    insertOS.NUMPRECIOSUGERIDO,insertOS.NUMPRECIOOFICIAL, insertOS.NUMDESCUENTO,0,time.Now(),4,time.Now()/*:DTEMODIFICACION*/,
				4/*:NUMPRECIOESTIMADO*/,0/*:NUMNROPIEZA*/,insertOS.VCHNROTRABAJO/*:VCHNROTRABAJO*/,insertOS.NUMCODIGO/*:NUMCODIGO*/,0/*:NUMMANOOBRA*/,0/*:NUMREPUESTO*/,insertOS.NUMTOTAL/*:NUMTOTAL*/)

		if err != nil {
		log.Println("error")
		log.Println(err)
	}}
	if insertOS.VCHCODIGOOPERACION =="1V"{
	_, err := utils.DbmapOracle.Exec(`INSERT INTO OPERACIONSERVICIOS (CHRCODIGOOPERACIONSERVICIO,VCHCODIGOOPERACION,VCHDESCRIPCION,
                                       NUMPRECIOSUGERIDO,NUMPRECIOOFICIAL,NUMHORASHOMBRE,NUMDESCUENTO,
                                       CHRESTADO,DTECREACION,NUMTIPOSERVICIO,DTEMODIFICACION,NUMPRECIOESTIMADO,NUMNROPIEZA,VCHNROTRABAJO,NUMCODIGO,NUMMANOOBRA,NUMREPUESTO,NUMTOTAL)
        values (:CHRCODIGOOPERACIONSERVICIO,:VCHCODIGOOPERACION,:VCHDESCRIPCION,:NUMPRECIOSUGERIDO,:NUMPRECIOOFICIAL,:NUMHORASHOMBRE,:NUMDESCUENTO,:CHRESTADO,
               :DTECREACION,:NUMTIPOSERVICIO,:DTEMODIFICACION,:NUMPRECIOESTIMADO,:NUMNROPIEZA,:VCHNROTRABAJO,:NUMCODIGO,:NUMMANOOBRA,:NUMREPUESTO,:NUMTOTAL)`, insertOS.CHRCODIGOOPERACIONSERVICIO, insertOS.VCHCODIGOOPERACION, insertOS.VCHDESCRIPCION,
			    insertOS.NUMPRECIOSUGERIDO,insertOS.NUMPRECIOOFICIAL, insertOS.NUMHORASHOMBRE, insertOS.NUMDESCUENTO,0,time.Now(),10,
				time.Now()/*:DTEMODIFICACION*/,	4/*:NUMPRECIOESTIMADO*/,0/*:NUMNROPIEZA*/,insertOS.VCHNROTRABAJO/*:VCHNROTRABAJO*/,insertOS.NUMCODIGO/*:NUMCODIGO*/,0/*:NUMMANOOBRA*/,0/*:NUMREPUESTO*/,insertOS.NUMTOTAL/*:NUMTOTAL*/)

		if err != nil {
		log.Println("error")
		log.Println(err)
	}}

	if (insertOS.VCHCODIGOOPERACION !="1V") && (insertOS.VCHCODIGOOPERACION !="1Z") && (insertOS.VCHCODIGOOPERACION !="1W") && (insertOS.VCHCODIGOOPERACION !="1X") && (insertOS.VCHCODIGOOPERACION !=""){
		_, err := utils.DbmapOracle.Exec(`INSERT INTO OPERACIONSERVICIOS (CHRCODIGOOPERACIONSERVICIO,VCHCODIGOOPERACION,VCHDESCRIPCION,
                                       NUMPRECIOSUGERIDO,NUMPRECIOOFICIAL,NUMHORASHOMBRE,NUMDESCUENTO,
                                       CHRESTADO,DTECREACION,NUMTIPOSERVICIO,DTEMODIFICACION,NUMPRECIOESTIMADO,NUMNROPIEZA,VCHNROTRABAJO,NUMCODIGO,NUMMANOOBRA,NUMREPUESTO,NUMTOTAL)
        values (:CHRCODIGOOPERACIONSERVICIO,:VCHCODIGOOPERACION,:VCHDESCRIPCION,:NUMPRECIOSUGERIDO,:NUMPRECIOOFICIAL,:NUMHORASHOMBRE,:NUMDESCUENTO,:CHRESTADO,
               :DTECREACION,:NUMTIPOSERVICIO,:DTEMODIFICACION,:NUMPRECIOESTIMADO,:NUMNROPIEZA,:VCHNROTRABAJO,:NUMCODIGO,:NUMMANOOBRA,:NUMREPUESTO,:NUMTOTAL)`, insertOS.CHRCODIGOOPERACIONSERVICIO, insertOS.VCHCODIGOOPERACION, insertOS.VCHDESCRIPCION,
			    insertOS.NUMPRECIOSUGERIDO,/*insertOS.NUMPRECIOOFICIAL*/0, insertOS.NUMHORASHOMBRE, insertOS.NUMDESCUENTO,0,time.Now(),0,time.Now()/*:DTEMODIFICACION*/,
				4/*:NUMPRECIOESTIMADO*/,0/*:NUMNROPIEZA*/,insertOS.VCHNROTRABAJO/*:VCHNROTRABAJO*/,insertOS.NUMCODIGO/*:NUMCODIGO*/,0/*:NUMMANOOBRA*/,0/*:NUMREPUESTO*/,insertOS.NUMTOTAL/*:NUMTOTAL*/)

		if err != nil {
		log.Println("error")
		log.Println(err)
	}}

	return insertOS,nil
	// curl -X POST -H "Content-Type: application/json" -H "Cache-Control: no-cache" -d '{"country_id": "PE","COUNTRY_NAME": "PERU","region_id": 1}' "http://localhost:8082/api/v1/country/"}
}