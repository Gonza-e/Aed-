Accion Ejercicio1 es 
 Ambiente 

	nodo = registro 
		valor: N(3)
		posicion: N(3)
		subconjunto: N(3)
		prox: puntero a nodo 
	FReg 

	prim,p,s,q:puntero a nodo 

	prom, total_num, total_sub, menor, mayor, pos_may, pos, sub, k, numero: entero 

	procedimiento crearNodo() es 
		Nuevo(q) 
		*q.valor:= numero; *q.posicion:= pos; *q.subconjunto:= sub
	FP

 Proceso
	 	prim:= nil
	 	prom:= 0; menor:= HV; mayor:= 0; pos_may:= 0; sub:= 1; numero:= 0; k:= 0; pos:= 1; total_num:= 0; total_sub:= 0;
	 
		repetir 
			Escribir("Ingrese un valor positivo, ingrese un valor negativo si desea terminar")
			Leer(numero)

			si numero >= 0 entonces 
				si prim = nil entonces 
					crearNodo()
					*q.prox:= prim 
					prim:= q 

				sino
					p:= prim 
					Mientras (p<>nil) y (k<3) hacer 
						si *p.valor = numero entonces 
							k:= k + 1
						fsi 
						p:= *p.prox 
					FM

					si k > 2 entonces 
						Escribir("No puede repetirse mas de dos veces este valor")
						k:= 0
					sino 
						crearNodo()
						*q.prox:= prim 
						prim:= q 
						pos:= pos + 1 
					fsi 
				fsi 
			fsi 

			si *prim.valor = 0 entonces 
				*prim.posicion:= 0; *prim.subconjunto:= sub 
				sub:= sub + 1 
				pos:= 0
				total_sub:= total_sub + 1
			fsi 

	 	hasta que (numero < 0)

		k:= 0
		p:= prim; s:= p
		Mientras (s<>nil) hacer 

			Mientras (s<>0) hacer 
				si *s.valor < menor entonces 
					menor:= *s.valor 
				sino 
					si *s.valor > mayor entonces 
						mayor:= *s.valor 
						pos_may:= *s.posicion
						sub:= *s.subconjunto
					fsi 
				fsi 

				total_num:= total_num + *s.valor
				k:= k + 1 
				s:= *s.prox 
			FM 

			prom:= (total_num/k)
			p:= s 
			s:= *s.prox 
			Escribir("En el subconjunto ",*p.subconjunto," el valor minimo es ",menor); menor:= 0 //Consigna 4
			Escribir("El promedio del subconjunto ",*p.subconjunto," es de ", prom:0:2) //Consigna 1
			k:= 0
		FM 

		Escribir("El total de subconjuntos procesados es de: ",total_sub) //Consigna 2
		Escribir("El valor maximo del conjunto es ",mayor," en la posicion ",pos_may," en el subconjunto ",sub)  //Consigna 3

	FProceso 
FAccion

			
			
	 
	