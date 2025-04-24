Accion Ejercicio1 es 
 Ambiente 
	sec, salida1, salida2: secuencia de caracteres
	v: caracter 

	nodo = registro 
		dato: AN(1)
		prox: puntero a nodo 
	FReg 

	p, prim, q, ant: puntero a nodo 
	zona1, infraccion, masdedos: booleano 
	cant_z1, cant_infraccion, cant_total, mayor: entero
	
	procedimiento cargaEncolada() es 
		si prim = nil entonces 
			*q.prox:= prim  
			prim:= q 
		sino 
			p:= prim; ant:= nil
			Mientras (p<>nil) hacer 
				ant:= p 	
				p:= *p.prox
			FM 
			
			*ant.prox:= q 
			*q.prox:= p 
		fsi 
	FP 
	
	A: arreglo[1...4] de enteros 

	procedimiento cargarTipo() es 
		segun *p.dato hacer 
			= "A": A[1]:= A[1] + 1 
			= "B": A[2]:= A[2] + 1 
			= "C": A[3]:= A[3] + 1 
			= "D": A[4]:= A[4] + 1 
		Fsegun
	FP 

 Proceso 
	 	Arr(sec)
		Avz(sec,v)
		Crear(salida1); Crear(salida2)
		zona1:= falso; infraccion:= falso; masdedos:= falso 
		cant_z1:=0; cant_infraccion:=0; cant_total:=0; mayor:=0 
		prim:= nil

		Mientras NFDS(sec) hacer 
			Mientras v <> "*" hacer 
				Mientras v <> "+" hacer  
					Nuevo(q); *q.dato:= v 
					cargaEncolada()
				FM 
				Avz(sec,v)
				cant_total:= cant_total + 1 

				p:= prim 
				Mientras (prim<>nil) hacer 
					Mientras (*p.dato <> "-") hacer 
						p:= *p.prox 
					FM 
					p:= *p.prox 

					cargarTipo(); p:= *p.prox 

					si conv_entero(*p.dato) > 2 entonces 
						masdedos:= verdadero
					fsi 
					p:= *p.prox 
					
					si *p.dato = "S" entonces 
						infraccion:= verdadero 
						cant_infraccion:= cant_infraccion + 1
					fsi 
					p:= *p.prox; p:= *p.prox

					si *p.dato = "1" entonces 
						zona1:= verdadero
						cant_z1:= cant_z1 + 1 
					fsi 

					p:= prim 
					Mientras (p<>nil) hacer 
						si (infraccion) y (masdedos) entonces 
							grabar(salida2,*p.dato)
						fsi 
						si (zona1) entonces 
							grabar(salida1,*p.dato)
						fsi 

						prim:= *p.prox 
						disponer(p)
						p:= prim 
					FM 

					infraccion:= falso; masdedos:= falso; zona1:= falso 
				FM 
			FM 
		FM
		
		Para i:= 1 a 4 hacer 
			si mayor < A[i] entonces 
				mayor:= A[i]
			fsi 
		FPara

	 	Esc("Vehiculos z1: ",cant_z1," Porcentaje: ",(cant_z1/cant_total)*100,"%")
		Esc("Vehiculos con infraccion: ",cant_infraccion," Porcentaje: ",(cant_infraccion/cant_total)*100,"%")
		Esc("")
		Cerrar(salida1); Cerrar(salida2); Cerrar(sec)
	FProceso
FAccion



					
					


		
