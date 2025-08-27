accion ejercicio1 es 
 Ambiente 
	sec: secuencia de caracteres
	v: caracter 

	cod, pref: AN 
	edad_cli, prom_cli, total_cli: entero
	
	nodo = registro 
		codprod: AN(40)
		cant: N(5)
		prefe: arreglo[1...3] de enteros //1=M 2=D 3=N
		prox: puntero a nodo 
	freg
	prim,q,p: puntero a nodo 

	procedimiento crearNodo() es
		Nuevo(q)
		*q.codprod:= cod; *q.cant:= 1 
		calcularPref(q)
	FProcedimiento

	procedimiento calcularPref(p: puntero a nodo) es 
		segun pref hacer 
			="M": i:= 1
			="D": i:= 2
			="N": i:= 3 
		fsegun
		*p.prefe[i]:= *p.prefe[i] + 1  
	FProcedimiento

	Proceso 
		ARR(sec,v); AVZ(sec,v)
		cod:= " "; pref:= " "
		edad_cli:= 0; prom_cli:= 0; total_cli:= 0 

		prim:= nil 

		Mientras NFDS(sec) hacer 
			Para i:= 1 hasta 7 hacer 
				AVZ(sec,v)
			FPara 

			Para i:= 1 hasta 2 hacer 
				edad_cli:= (edad_cli * 10) + conv_entero(v)
				AVZ(sec,v)
			FPara 

			Mientras v <> "-" hacer 
				Concatenar(cod,v); AVZ(sec,v)
			FM 
			AVZ(sec,v)
			
			pref:= v
			
			si pref = "N" entonces 
				prom_cli:= prom_cli + edad_cli
				total_cli:= total_cli + 1 
			fsi 

			si prim = nil entonces 
				crearNodo()
				*q.prox:= nil
				prim:= q 
			sino 
				p:= prim 
				Mientras (p<>nil) y (cod <> *p.codprod) hacer 
					p:= *p.prox 
				FM 

				si cod = *p.codprod entonces 
					calcularPref(p); *p.cant:= *p.cant + 1 
				sino 
					crearNodo(q); 
					*q.prox:= prim 
					prim:= q
				fsi 
			fsi 
		FM 
		Cerrar(sec)

		p:= prim 
		Mientras p <> nil hacer 
			Esc("Cod producto: ",*p.codprod," Cantidad de compras: ",*p.cant)
			si *p.prefe[3] > *p.prefe[2] entonces 
				Esc("Este producto tuvo mas preferencia nula que deseable")
			fsi 
			p:= *p.prox 
		FM 

		prom:= (prom/total_cli); Esc("Promedio de edad de clientes de preferencia nula: ",prom)
	FProceso 
FAccion


