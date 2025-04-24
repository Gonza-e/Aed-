Accion final29 es 
 Ambiente 
	nodo = registro 
		subconjunto: N(3)
		valor: N(3)
		posicion: N(3)
		prox: puntero a nodo 
	Freg 

	prim,p,s,q,ant: puntero a nodo 

	procedimiento cargaEncolada() es 
		p:= prim; ant:= nil
		Mientras p <> nil hacer 
			ant:= p 
			p:= *p.prox 
		FM
		*q.prox:= *ant.prox 
		*ant.prox:= q 
	FP 

	i, j, k, prom, total_subconjuntos, valor_min; valor_max, sub_max, pos_max, valor_temp, lotes_vacios: entero

	Proceso 
		prim:= nil 
		prom:= 0; total_subconjuntos:= 0; valor_min:= HV; valor_max:= 0; sub_max:= 0; sub_max:= 0; pos_max:= 0; lotes_vacios:= 0
		i:= 1; j:= 1, k:= 0
		repetir 
			Esc("Ingrese un valor mayor a 0")
			Leer(valor_temp)

			si valor_temp > 0 entonces 
				si prim = nil entonces 
					Nuevo(q); *q.subconjunto:= i, *q.posicion:= j, *q.valor:= valor_temp
					cargaEncolada()
				sino 
					p:= prim; ant:= nil
					Mientras p <> nil hacer 
						si valor_temp = *p.valor entonces 
							k:= k + 1 
						fsi 
						ant:= p 
						p:= *p.prox 
					FM 

					si k >= 2 entonces 
						Esc("Error el valor se repite mas de dos veces")
						k:= 0
					sino 
						si valor_temp <> 0 entonces 
							j:= j + 1
							Nuevo(q); *q.subconjunto:= i; *q.posicion:= j; *q.valor:= valor_temp
						sino 
							j:= 0
							Nuevo(q); *q.subconjunto:= i; *q.posicion:= j; *q.valor:= valor_temp
							i:= i + 1 
						fsi 
						cargaEncolada()
					fsi 
				fsi 
			fsi 
		hasta que (valor_temp < 0)

		p:= prim; s:= *prim.prox  
		i:= 0
		si prim = nil entonces
			Esc("El lote esta vacio")
		sino 
			Mientras *p.prox <> nil hacer
				Mientras *s.valor <> 0 hacer 
					i:= i + 1 
					prom:= prom + *s.valor
					
					si *s.valor < valor_min entonces 
						valor_min:= *s.valor 
					fsi 

					si *s.valor > valor_max entonces
						valor_max:= *s.valor 
						sub_max:= *s.subconjunto
						pos_max:= *s.posicion
					fsi 
					s:= *s.valor
				FM 

				si i = 0 entonces 
					lotes_vacios:= lotes_vacios + 1 
				sino 
					Esc("El valor minimo del subconjunto ",*s.subconjunto," es ",valor_min)
					Esc("El promedio de valores del subconjunto",*s.subconjunto," es ",(prom/i))
					i:= 0; prom:= 0; valor_min:= HV 
					p:= s; s:= *s.prox 
				fsi 
				total_subconjuntos:= total_subconjuntos + 1 
			FM 
		fsi 

		Esc("El valor maximo fue: ",valor_max," en el subconjunto ",sub_max," en la posicion ",pos_max)
		Esc("El porcentaje de subconjuntos vacios es de: "(lotes_vacios/total_subconjuntos)*100,"%")
	FProceso 
FAccion







