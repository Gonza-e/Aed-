1) El regimen de Jubilaciones para Docentes Universitarios establece que, segun las leyes vigentes podran acceder al beneficio jubilatorio quienes cumplan determinados requisitos :

Edad; Hombres: Entre 65 y 70 año y Mujeres: Entre 60 y 70 años 
Aportes; Minimo 30 años de aportes 

La Facultad Regional Resistencia cuenta con dos secuencias, una de caracteres y otra de enteros, con las siguientes caracteristicas:

Secuencia de caracteres DOCENTES: contiene los siguientes datos de cada docente :
REGIONAL: 3 caracteres, SEXO: 1 caracter ("M" o "F"), LEGAJO: 5 caracteres, NOMBRE COMPLETO: cantidad indefinida de caracteres, finaliza con el simbolo "#"

Los docentes de una misma regional terminan en un "+"
La regional siempre es FR seguido de un caracter, que es la inicial de la provincia 

Ejemplo:

//FRMM22711PedroGonzalez#FRMF19245MonicaMartinez+FRC...#...#...FDS

Secuencia de enteros EDAD contiene, con correspondencia 1 a 1, para cada docente los siguientes datos: 

edad: 2 caracteres, años de aporte: 2 caracteres 

Ejemplo:

//6720,5030...FDS 

a) Generar una secuencia de caracteres JUBILAR, solo para docentes de una regional especifica ingresada por el usuario, la cual debe tener: legajo, nombre completo de los docentes que cumplam las condiciones para jubilarse 

b) Informar por regional la cantidad de docentes que puedan jubilarse y la cantidad que no pueden

Accion Parcial es 
 Ambiente 
 doc: secuencia de caracteres 
 v1: caracter 
 edad: secuencia de enteros 
 v2: entero
 salida: secuencia de caracteres
 reg: AN
 cant_si, cant_no, ap, ed: entero
 bandera: booleano
 Proceso 
	 Arr(doc)
	 Arr(edad)
	 Crear(salida)
	 Avz(doc,v1)
	 Avz(edad,v2)
	 cant_si:= 0 
	 cant_no:= 0 
	 ap:= 0 
	 ed:= 0 
	 bandera:= falso
	 Esc("Ingrese una regional")
	 Leer(reg)
	 Mientras NFDS(doc) Y NFDS(edad) hacer 
	 	 Avz(doc,v1)
		 Avz(doc,v1)
		 	Si v1 = reg entonces
			 ed:= ed + (v2 DIV 100)
			 ap:= ap + (v2 MOD 100)
			 Avz(doc,v1)
			 Segun v1 hacer 
			 	 "M": 	
				 	Si (ed > 65) Y (ed < 70) entonces
				     	Si ap >= 30 entonces
						 Mientras v1 <> "#" hacer 
						 	 Mientras v1 <> "+" hacer 
							 	 Grabar(salida,v1)
								 Avz(doc,v1)
								FinMientras
							 Avz(doc,v1)
							FinMientras
						 cant_si:= cant_si + 1 
						Sino 
						 cant_no:= cant_no + 1
						 bandera:= verdadero 
						 Mientras v1 <> "#" hacer 
			 				 Avz(doc,v1)
							FinMientras
						FinSi
					Sino 
					 cant_no:= cant_no + 1
					 bandera:= verdadero 
					 Mientras v1 <> "#" hacer 
			 			 Avz(doc,v1)
						FinMientras
					FinSi
				 "F": 	
			 	 	Si (ed > 60) Y (ed < 70) entonces
				     	Si ap >= 30 entonces
						 Mientras v1 <> "#" hacer 
						 	 Mientras v1 <> "+" hacer 
							 	 Grabar(salida,v1)
								 Avz(doc,v1)
								FinMientras
							 Avz(doc,v1)
							FinMientras
						 cant_si:= cant_si + 1 
						Sino 
						 cant_no:= cant_no + 1
						 bandera:= verdadero 
						 Mientras v1 <> "#" hacer 
			 				 Avz(doc,v1)
							FinMientras
						FinSi
					Sino 
					 bandera:= verdadero 
					 Mientras v1 <> "#" hacer 
			 			 Avz(doc,v1)
						FinMientras
					FinSi
				FinSegun
			Sino 
			 bandera:= verdadero 
			 Mientras v1 <> "#" hacer 
			 	 Avz(doc,v1)
				FinMientras
			FinSi
		 ed:= 0 
		 ap:= 0 
		 	Si bandera entonces 
			 Avz(edad,v2)
			FinSi
		FinMientras
	 Cerrar(edad)
	 Cerrar(doc)
	 Cerrar(salida)
	 Esc("La cantidad de docentes que pueden jubilarse es de",cant_si)
	 Esc("La cantidad de docentes que no pueden jubilarse es de",cant_no)
	FinProceso 
FinAccion
		 


			 