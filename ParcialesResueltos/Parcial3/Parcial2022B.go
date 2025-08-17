Accion ejercicio1 (prim: puntero a nodo) es 
 Ambiente 
 	nodo = registro
		pais: AN(20)
		edades: arreglo [1...26] de enteros 
		grupo: AN(1)
		puntos: N(2)
		t_amarillas: N(2)
		t_rojas: N(2)
		prox:= puntero a nodo 
	FReg
	p, p_est, p_mov, p_menor: puntero a nodo 

	grupo_men: AN
	pais_aux, grupo_aux: AN
	edades_aux: arreglo[1...26] de enteros 
	prom_menor, prom_mayor, suma: entero

	Procedimiento ordenar_lista() es 
		p_est:= prim 
		p_mov:= *prim.prox
		Para i:= 1 a 31 hacer 
			grupo_men:= *p_est.grupo
			pais_aux:= *p_est.pais
			edades_aux:= *p_est.edades
			p_menor:= p_est
		 	Para j:= 2 a 32 hacer 
		 	 	Si grupo_men < *p_mov.grupo 
					grupo_men:= *p_mov.grupo
					pais_aux:= *p_mov.pais
					edades_aux:= *p_mov.edades
					p_menor:= p_mov
				FSi
			 	p_mov:= *p_mov.prox
			FPara 
			*p_menor.pais:= *p_est.pais 
			*p_menor.edades:= *p_est.edades 
			*p_menor.grupo:= *p_est.grupo 
			*p_est.pais:= pais_aux
			*p_est.edades:= edades_aux
			*p_est.grupo:= grupo_men
			p_est:= *p_est.prox
		FPara
	FProcedimiento

	Funcion promedio(indice: entero, A: arreglo[1...26] de enteros): entero
	 	Si indice = 26 entonces 
		 	promedio:= A[indice] DIV 26 
		Sino 
		 	promedio:= A[indice] DIV 26 + promedio(indice + 1, A[i])
		FSi
	FFuncion
	
 Proceso 
		prom_mayor:= 0
		prom_menor:= HV
		suma:= 0 
		ordenar_lista()
		p:= prim 
		Mientras (p <> nil) hacer 
	 	 	Si promedio(suma,1,*p.edades) < prom_menor entonces
			 	prom_menor:= promedio()
			Sino 
			 	Si promedio(suma,1,*p_edades) > prom_mayor entonces 
				 	prom_mayor:= promedio()
				FSi
			FSi
		FM
	FProceso
			 


