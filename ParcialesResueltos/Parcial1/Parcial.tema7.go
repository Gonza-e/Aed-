1)El centro bioquímico chaco cuenta con información de análisis realizados durante el día. Esta información se encuentra almacenada en una secuencia de caracteres la cual inicia con el código de protocolo (5 caracteres, el primero siempre una letra), luego el apellido y nombre del paciente (finaliza en una ","), cantidad de estudios realizados (2 caracteres) y los códigos de cada uno (4 caracteres). El Código de cada estudio inicia con una vocal, que indica el tipo. Si el tipo "A", cuesta $300, si es de tipo "E" cuesta $420 y si es de tipo "I" cuesta $670. La secuencia finaliza con "*"

Ejemplo:
//A2462Reina Isabel,03A123E345E333P2342Rey Leon,01E888*

Se solicita: 
A) Generar otra secuencia de salida que almacene los estudios solicitados tipo "A", siempre y cuando la cantidad de estudios de esa persona supere los 2.
B) Informar total recaudado en el día.
C) Promedio de estudios por paciente. 

Accion tema7 es 
 Ambiente 
 sec: secuencia de caracteres 
 v: caracter
 sal: secuencia de caracteres 
 prom, tot_rec, cant_est, cant_estudios, cant_cli: entero 
 funcion conv_entero(v:caracter):entero 
	 	segun v hacer 
	 	 ="1": conv_entero:= 1
		 ="2": conv_entero:= 2
		 ="3": conv_entero:= 3
		 ="4": conv_entero:= 4
		 ="5": conv_entero:= 5
		 ="6": conv_entero:= 6
		 ="7": conv_entero:= 7
		 ="8": conv_entero:= 8
		 ="9": conv_entero:= 9
		 ="0": conv_entero:= 0
		FinSegun
	FinFuncion
 Proceso 
	 Arr(sec)
	 Avz(sec,v)
	 Crear(sal)
	 prom:= 0 
	 tot_rec:= 0
	 cant_est:= 0 
	 cant_cli:= 0
	 cant_estudios:= 0
 	 Mientras v <> "*" hacer 
		 Mientras v <> "," hacer 
		 	 Avz(sec,v)
			FinMientras
		 Avz(sec,v)
		 cant_cli:= cant_cli + 1
		 cant_estudios:= cant_estudios + (conv_entero(v) * 10)
		 Avz(sec,v) 
		 cant_estudios:= cant_estudios + conv_entero(v)
		 cant_est:= cant_est + cant_estudios
		 	Si cant_estudios > 2 entonces
				Para i:= 1 a (cant_estudios*4) hacer
			 		Segun v hacer 
				 	 "A": tot_rec:= tot_rec + 300
				 		Para i:= 1 a 4 hacer 
					 	 Grabar(sal,v)
					 	 Avz(sec,v)
						FinPara
					 "E": tot_rec:= tot_rec + 420 
					 	Para i:= 1 a 4 hacer
					     Avz(sec,v)
						FinPara
					 "I": tot_rec:= tot_rec + 670
						Para i:= 1 a 4 hacer 
					 	 Avz(sec,v)
						FinPara
					FinSegun
				FinPara
			FinSi
		 cant_estudios:= 0
		FinMientras
	 Cerrar(sec)
	 Cerrar(sal)
	 prom:= prom + (cant_est/cant_cli)
	 Escribir("El total recaudado en el dia es de",tot_rec,"$")
	 Escribir("El promedio de estudios por paciente es de",prom)
	FinProceso
FinAccion 
	

			    

		  
	 