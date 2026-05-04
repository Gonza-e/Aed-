/*1)El centro bioquímico chaco cuenta con información de análisis realizados durante el día. Esta información se encuentra
almacenada en una secuencia de caracteres la cual inicia con el código de protocolo (5 caracteres, el primero siempre una letra),
luego el apellido y nombre del paciente (finaliza en una ","), cantidad de estudios realizados (2 caracteres) y los códigos de 
cada uno (4 caracteres). El Código de cada estudio inicia con una vocal, que indica el tipo. 
Si el tipo "A", cuesta $300, si es de tipo "E" cuesta $420 y si es de tipo "I" cuesta $670. La secuencia finaliza con "*"

Ejemplo:
//A2462Reina Isabel,03A123E345E333P2342Rey Leon,01E888*

Se solicita: 
A) Generar otra secuencia de salida que almacene los estudios solicitados tipo "A", siempre y cuando la cantidad de estudios de esa persona supere los 2.
B) Informar total recaudado en el día.
C) Promedio de estudios por paciente. */

Accion tema7 es 
	Ambiente 
		sec,sal: secuencia de caracteres 
		v: caracter
		prom, tot_rec, acumEstudios, cant_est, cant_estudios, cant_cli: entero 
		Funcion conv(v:caracter):entero 
			segun v hacer 
				"1": conv:= 1
				"2": conv:= 2
				"3": conv:= 3
				"4": conv:= 4
				"5": conv:= 5
				"6": conv:= 6
				"7": conv:= 7
				"8": conv:= 8
				"9": conv:= 9
				"0": conv:= 0
			FinSegun
		FF

		Funcion costo(x:AN):entero 
			segun x hacer 
				"A": costo:= 300
				"E": costo:= 420
				"I": costo:= 670
			Fsegun 
		FFuncion 

	Proceso 
		Arr(sec); Avz(sec,v); Crear(sal)
		prom:= 0; tot_rec:= 0; cant_est:= 0; cant_cli:= 0; cant_estudios:= 0; acumEstudios:= 0

		Mientras v <> "*" hacer 
			cant_cli:= cant_cli+1 
			Mientras v <> "," hacer
				Avz(sec,v)
			FM 
			Avz(sec,v)
			
			cant_estudios:= 0
			Para i:= 1 hasta 2 hacer 
				cant_estudios:= cant_estudios*10 + conv(v)
				Avz(sec,v)
			FPara 

			acumEstudios:= acumEstudios+cant_estudios 

			Para i:= 1 hasta cant_estudios hacer 
				Para i:= 1 hasta 4 hacer 
					si i=1 y (v="A") entonces 
						bandera:= verdadero 
					Fsi 
					si i=1 entonces 
						tot_rec:= tot_rec+costo(v)
					Fsi 
					si bandera entonces 
						Grabar(sal,v)
					Fsi 
					Avz(sec,v)
				FPara 
				bandera:= falso 
			FPara 
		FM 
		Cerrar(sec); Cerrar(sal)

		Esc("Promedio de estudios por paciente: ",acumEstudios DIV cant_cli)
		Esc("Total recaudado en el dia: $",tot_rec)					 
	FinProceso
FinAccion 
	

			    

		  
	 