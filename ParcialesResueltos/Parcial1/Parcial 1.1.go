Parcial 1.1 

La primera secuencia contiene a los fanáticos que hacen la “fila virtual”, esta secuencia contiene la hora en la que se registró en la fila (hhmm), el número en la fila (6 caracteres), el medio de pago con el que piensa pagar (“T”- tarjeta, “C”-efectivo, “E”-transferencia) y cuántas entradas piensa comprar (como máximo puede comprar 5 entradas).
En estos sistemas de fila, al colocarse en ella se obtiene un identificador o número de fila, pero recién el usuario inicia sesión al llegar a su lugar en la fila y entrar a la sección de Compras.
En una segunda secuencia se encuentran las compras reales realizadas, esta secuencia contiene el número en la fila (6 caracteres), el nombre de usuario (cantidad indefinida, termina en “+”), y luego contiene los datos de para quién es cada entrada en la compra, incluyendo su DNI y nombre (separados con ‘.’). Una compra finaliza con un “?”.

Secuencia Fila_Virtual:
//horafilanumerofilamediopagocantentradas#
Ejemplo:
0301123456T3#0345234567E4#

Secuencia Compras:
//numerofilanombre_usuario+dninombre.dninombre.dninombre?
Ejemplo:
123456unsuario+33254787Juan.27895614Melisa.36257489Pedro?

Si por algún motivo algún fanático deja su lugar en la fila, en la secuencia de Compras en nombre de usuario aparece un carácter “#”, el signo “+” y luego la marca “?”. Esto implica una correspondencia 1 a 1 entre las dos secuencias.
Escriba un algoritmo en pseudocódigo que resuelva las siguientes consignas:
a) Generar una secuencia que contenga el número de fila y DNI para cada entrada comprada por esa persona, solo de aquellas compras que entraron en la cola antes de las 10:00 am. con el siguiente formato.
//EJ: numerofila+dni+dni#numerofila+dni+dni#
b) Se desea conocer la cantidad de usuarios que estaban en la fila, pero no compraron entradas (dejaron la fila).

Accion parcial1.1 es 
Ambiente 
 	f_virtual: secuencia de caracteres 
 	v1: caracter
 	compras: secuencia de caracteres 
 	v2: caracter 

 	funcion conv_entero(caracter):entero
		segun v hacer 
	 	 ="1": conv_entero := 1
	 	 ="2": conv_entero := 2
	 	 ="3": conv_entero := 3 
	 	 ="4": conv_entero := 4
	 	 ="5": conv_entero := 5 
	 	 ="6": conv_entero := 6
	  	 ="7": conv_entero := 7
	 	 ="8": conv_entero := 8
	 	 ="9": conv_entero := 9
	 	 ="0": conv_entero := 0
		FinSegun 
	FinFuncion 
 cant_abandono, hora: entero 
 digitos: (1,2,3,4,5,6,7,8,9,0)
 bandera, bandera1: booleano 

 Proceso 
		Arr(f_virtual)
		Arr(compras)
		Crear(salida)
		Avz(f_virtual,v1)
		Avz(compras,v2)
		cant_abandono:= 0 
		hora:= 0 
		bandera:= falso 
		bandera1:= falso

	 	Mientras NFDS(f_virtual) Y NFDS(compras) hacer
	 	 	conv_entero(v1) 
	 	 	hora:= hora + (conv_entero(v1)*10)
		 	Avz(f_virtual,v1)
		 	conv_entero(v1)
		 	hora:= hora + conv_entero(v1)

		 	Si hora = 10 entonces 
			 	Mientras v2 <> "?" hacer 
			 	 	Mientras v2 <> "+" hacer 

				 		Si v2 = "#" entonces 
						 	cant_abandono:= cant_abandono + 1
						 	bandera:= verdadero
						Sino 
						 	Avz(compras,v2)
						FinSi

				 	 	Si v2 EN digitos entonces 
						 	Grabar(salida,v2)
						 	Avz(compras,v2)
						Sino 
						 	Avz(compras,v2)
						FinSi
					 	Grabar(salida,"+")
					FinMientras
				 Avz(compras,v2)

				 	Mientras v2 <> "?" hacer 
					 	Mientras v2 <> "." hacer 

				 		 	Si v2 EN digitos entonces
					 		 	Grabar(salida,v2)
							 	Avz(compras,v2)
							 	bandera1:= verdadero
							Sino 
							 	Avz(compras,v2)
							FinSi
						FinMientras
					FinMientras

					Si bandera1 entonces 
					 	Grabar(salida,"+")
					FinSi
				 	Avz(compras,v2)
				FinMientras
			FinSi

			Si bandera entonces
			 Mientras v1 <> "#" hacer 	
			 	 	Avz(f_virtual,v1)
				FinMientras
			FinSi
		 	hora:= 0 
		FinMientras
		
	 	Cerrar(f_virtual)
	 	Cerrar(compras)
	 	Cerrar(salida)
	 	Escribir("La cantidad e personas que abandonaron es de",cant_abandono)
	FinProceso 
FinAccion
			 
					
