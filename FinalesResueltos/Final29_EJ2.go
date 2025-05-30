Accion final29 es 
 Ambiente 
	DATOS = registro 
		idviaje: N(6)
		provincia: (1...23)
		mes: (1...12)
		cant_dias: N(2)
		cant_personas: N(2)
		alojamiento: ("H","D","NI")
		deuda: N(6,2)
		credito: N(6,2)
	freg 
	arch: archivo de DATOS 
	reg: DATOS 

	DATOS_MATRIZ = registro
		cant_viajes: N(3)
		cant_dias: N(2)
		cant_personas: N(2)
		deuda: N(6,2) 
	freg 

	nom_provincias: arreglo[1...23] de AN 

	matriz: arreglo[1...23,1...12] de DATOS_MATRIZ

	cant_x_duracion: arreglo[1...3] de enteros 

	provincia_mayor, provincia_menor, cant_may, cant_men, cant_may2, mes, provincia: entero

	i,j: entero
 Proceso 
		Abrir E/(arch)
		Leer(arch,reg)
		
		provincia_mayor:= 0; provincia_menor:= 0; cant_may:= LV, cant_men:= HV

		para i:= 1 a 23 hacer 
			para j:= 1 a 12 hacer
				matriz[i,j]:= 0 
			fpara 
		fpara 

		Mientras NFDA(arch) hacer 
			i:= reg.provincia; j:= reg.mes 

			matriz[i,j].cant_viajes:= matriz[i,j].cant_viajes + 1 
			matriz[i,j].cant_dias:= matriz[i,j].cant_dias + reg.cant_dias
			matriz[i,j].cant_personas:= matriz[i,j].cant_personas + reg.cant_personas
			matriz[i,j].deuda:= matriz[i,j].deuda + reg.deuda 

			si reg.cant_dias <= 3 entonces 
				cant_x_duracion[1]:= cant_x_duracion + 1 
			sino 
				si (reg.cant_dias > 3) y (reg.cant_dias < 10) entonces 
					cant_x_duracion[2]:= cant_x_duracion[2] + 1 
				sino 
					cant_x_duracion[3]:= cant_x_duracion[3] + 1 
				fsi 
			fsi 

			Leer(arch,reg) 
		FM 

		Cerrar(Arch)

		para i:= 1 a 23 hacer 
			para j:= 1 a 12 hacer 
				si matriz[i,j].cant_personas < cant_men entonces 
					provincia_menor:= i 
					cant_men:= matriz[i,j].cant_personas
				fsi 
				si matriz[i,j].cant_personas > cant_mayor entonces 
					provincia_mayor:= i 
					cant_may:= matriz[i,j].cant_personas
				fsi 

				Escribir("En el mes ",obtenerMes(j)," en la provincia ",obtenerProvincia(i), " se adeuda: ",matriz[i,j].deuda) // consigna C

				si matriz[i,j].cant_viajes > cant_may2 entonces 
					mes:= j 
					provincia:= i 
					cant_may:= matriz[i,j].cant_viajes
				sino 
					si matriz[i,j].cant_viajes = 0 entonces 
						Escribir("En el mes ",obtenerMes(j)," la provincia ",obtenerProvincia(i)," no se recibio viajes") // consigna E
					fsi 
				fsi 

				total_viajes:= total_viajes + matriz[i,j].cant_viajes
			fpara 
		fpara 

		Escribir("Viajes de menos de 3 dias: ",cant_x_duracion[1],"viajes de entre 3 y 10 dias: ",cant_x_duracion[2]," viajes de mas de 10 dias: ",cant_x_duracion[3]) // consigna A
		Escribir("Provincia con mas personas: ",obtenerProvincia(provincia_mayor)," cantidad: ",cant_may) // consigna B
		Escribir("Provincia con menos personas: ",obtenerProvincia(provincia_menor)," cantidad: ",cant_men) // consigna B
		Escribir("Provincia con mas viajes: ",obtenerProvincia(provincia)," mes con mas viajes: ",obtenerMes(mes)) // consigna D 
		Escribir("El total de viajes es de: ",total_viajes) // consigna F a medias 
	FProceso
FAccion




