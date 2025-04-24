Accion final31 es 
 Ambiente 
	CONTROL = registro 
		mes_visita: N(2)
		nroemp: N(6)
		razsoc: AN(15)
		cat: (1...3)
		habilitado: ("si","no")
		capacidad: N(3)
	freg

	arch: archivo de CONTROL ordenado por mes_visita y nroemp
	reg: CONTROL 
    menor_hab, no_hab, mes_menor, cat_menor, menor_100: entero
	porc: real 
	
	A: arreglo[1...13,1...4] de enteros 

	Funcion mostrarMes(mes: entero): AN 
		segun mes hacer 
			1: mostrarMes:= "Enero"
			2: mostrarMes:= "Febrero"
			3: mostrarMes:= "Marzo"
			4: mostrarMes:= "Abril"
			5: mostrarMes:= "Mayo"
			6: mostrarMes:= "Junio"
			7: mostrarMes:= "Julio"
			8: mostrarMes:= "Agosto"
			9: mostrarMes:= "Septiembre"
			10: mostrarMes:= "Octubre"
			11: mostrarMes:= "Noviembre"
			12: mostrarMes:= "Diciembre"
		fsegun 
	ffuncion

	Proceso 
	 	Abrir E/(arch)
		Leer(arch,reg)
		resg_mes:= reg.mes_visita
		mes_menor:= 0; menor_100:= 0; cat_menor:= 0; menor_hab:= HV; no_hab:= 0

		para i:= 1 a 13 hacer 
			para j:= 1 a 4 hacer 
				A[i,j]:= 0 
			fpara 
		fpara

		Mientras NFDA(arch) hacer 
			si reg.habilitado = "si" entonces
				si reg.capacidad < 100 entonces 
					menor_100:= menor_100 + 1
				fsi 
				
				A[reg.mes_visita,reg.cat]:= A[reg.mes_visita,reg.cat] + 1 
				A[13,4]:= A[13,4] + 1 
			sino 
				no_hab:= no_hab + 1 
			fsi 
		FM 
		cerrar(arch)

		para i:= 1 a 12 hacer 
			para j:= 1 a 3 hacer 
				Esc("Mes: ",mostrarMes(i), "Categoria: ",j, "Cantidad de restaurantes: ",A[i,j]) // Consigna 1
				si menor_hab < A[i,j] entonces 
					mes_menor:= i 
					cat_menor:= j 
				fsi 
			fpara 
		fpara 

		porc:= (no_hab/(no_hab + A[13,4]))*100

		Esc("Cantidad de restaurantes con capacidad menor a 100: ",menor_100) // Consigna 2
		Esc("Restaurantes con menor cantidad de habilitaciones, mes: ",mostrarMes(i), "categoria: ",cat_menor) // Consigna 3
		Esc("Porcentaje de restaurantes no habilitados: ",porc:0:2,"%") // Consigna 4
	FProceso
FAccion


		
				
