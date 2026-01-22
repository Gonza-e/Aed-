accion Ejercicio1 (b[j]) es //arreglo con las categorias 
 ambiente 
	EVENTOS = registro 
		codigo: AN(5)
		titulo: AN(30)
		costo: N(4,2)
		cod_categoria: (1...10)
		segmento: (0...2)
	FReg
	arch: archivo de EVENTOS 
	reg: EVENTOS

	menor, cat_menor, mayor, cat_mayor, i, j, k, cant_consignaA: entero 
	costo_usu, monto: real 

	nodo = registro 
		categoria: AN(30)
		segmento: AN(10)
		prox: puntero a nodo 
	FReg
	prim,p,q,ant: puntero a nodo 

	a: arreglo[0...3,1...11,1...2] de real 

	Funcion obtenerSegmento(n: entero): AN 
		segun n hacer
			= 0: obtenerSegmento:= "matinal"
			= 1: obtenerSegmento:= "tarde"
			= 2: obtenerSegmento:= "noche" 
		fsegun
	FFuncion

	Procedimiento cargaOrdenada() es 
		p:= prim; ant:= nil 
		Mientras (p <> nil) y (*q.costo < *p.costo) hacer 
			ant:= p 
			p:= *p.prox 
		FM 

		si ant = nil entonces 
			*q.prox:= prim 
			prim:= q 
		sino 
			*q.prox:= *ant.prox 
			*ant.prox:= q 
		fsi 
	FProcedimiento

	Proceso 
		Abrir(arch); Leer(arch,reg)
		
		Para k:= 1 hasta 2 hacer 
			Para i:= 0 hasta 3 hacer 
				Para j:= 1 hasta 11 hacer 
					a[i]:= 0 
				fpara
			fpara
		fpara

		costo_usu:= 0; menor:= 0; cat_mayor:= 0; cat_menor:= 0; mayor:= 0; menor:= 0; monto:= 0; cant_consignaA:= 0 
		prim:= nil 

		Esc("Ingrese monto: "); Leer(monto)

		Mientras NFDA(arch) hacer 
			i:= reg.segmento; j:= reg.cod_categoria
			k:= 1 

			si (i = 0) y (B[j] = "deportivo") entonces 
				cant_consignaA:= cant_consignaA + 1 
			fsi 

			si monto < reg.costo entonces 
				Nuevo(q), *q.categoria:= b[j]; *q.segmento:= obtenerSegmento(i)

				si prim = nil entonces 
					*q.prox:= prim 
					prim:= q 
				sino 
					cargaOrdenada()
				fsi 
			fsi 

			a[i,j,k]:= a[i,j,k] + 1 
			a[3,j,k]:= a[3,j,k] + 1 
			a[i,11,k]:= a[i,11,k] + 1 
			a[3,11,k]:= a[3,11,k] + 1 

			k:= 2 

			a[i,j,k]:= a[i,j,k] + reg.costo 
			a[3,j,k]:= a[3,j,k] + reg.costo 
			a[i,11,k]:= a[i,11,k] + reg.costo
			a[3,11,k]:= a[3,11,k] + reg.costo
				
			Leer(arch,reg)
		FM 

		Para k:= 1 hasta 2 hacer 
			Para i:= 0 hasta 3 
				Para j:= 1 hasta 11 hacer 
					si (k <> 2) y (i <> 3) y (j <> 11) entonces 
						si menor < a[i,j,k] entonces 
							menor:= a[i,j,k] 
							cat_menor:= j 
						fsi 

						si mayor > a[i,j,k] entonces 
							mayor:= a[i,j,k]
							cat_mayor:= j 
						fsi
						
					fsi 

					si (i = 2) y (k <> 2) entonces 
						porcentaje:= porcentaje + 1
					fsi 
				fpara
			fpara
		fpara 

		Cerrar(arch)
		porcentaje:= (mayor/porcentaje)*100
		Esc("La cantidad de eventos de matinal y deportivo es de: ",cant_consignaA)
		Esc("El total invertido en la categoria: ",b[cat_menor]," es de: ",menor)
		Esc("El porcentaje de: ",b[cat_mayor]," en el segmento noche es de: ",porcentaje)
	FProceso
FAccion