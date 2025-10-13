Accion corredores (Corredor: arreglo[1...2,1...20] de AN) es
 Ambiente
	Historico = registro 
		año: N(4)
		id_carrera: N(2)
		nro_equipo: (1...20)
		nro_piloto: (1...2)
		p_obtenidos: N(2)
		pos_final: (1...20)
	FReg

	arch: archivo de Historico
	reg: Historico

	A: arreglo[1...3,1...20] de enteros
	
	i,j,masPuntos,pilotoMayJ,pilotoMayI: enteros 

	Proceso 
		Abrir E/(arch); Leer(arch,reg)
		masPuntos:= 0; pilotoMayJ:= 0; pilotoMayI:= 0

		Para i:= 1 hasta 3 hacer 
			Para j:= 1 hasta 20 hacer 
				A[i,j]:= 0 
			Fpara
		Fpara

		Mientras NFDA(arch) hacer 
			i:= reg.nro_piloto
			j:= reg.nro_equipo

			A[i,j]:= A[i,j] + reg.p_obtenidos 

			si reg.pos_final = 1 entonces
				i:= 3 
				A[i,j]:= A[i,j] + 1 
			Fsi 

			Leer(arch,reg)
		FM 
		Cerrar(arch)

		Para i:= 1 hasta 2 hacer 
			Para j:= 1 hasta 20 hacer
				si A[i,j] > masPuntos entonces 
					masPuntos:= A[i,j]
					pilotoMayJ:= j
					pilotoMayI:= i 
				fsi 

				Esc("La cantidad de triunfos del equipo: ",Equipo(j)," es de: ", A[3,j])
				Esc("Diferencia de puntos entre pilotos del equipo ",Equipo(j),": ",A[1,j] - A[2,j])
			Fpara
		Fpara 

		Esc("El corredor con mas puntos es: ",Corredor(pilotoMayI,pilotoMayJ)," con: ",masPuntos," puntos")
	FProceso 
FAccion


			