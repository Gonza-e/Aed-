Accion Ejercicio2 es 
 Ambiente 

 	SUPERHEROES = registro 
		heroe1: AN(15)
		heroe2: AN(15)
		g_parciales_heroe1: N(2)
		g_parciales_heroe2: N(2)
		g_totales_heroe1: N(2)
		g_totales_heroe2: N(2)
	FReg

	arch: archivo de SUPERHEROES 
	reg: SUPERHEROES 

	nodo = registro 
		heroe: AN(15)
		g_totales_recibidos: N(3)
		puntos: N(3)
		prox: puntero a nodo 
	FReg 

	prim,p,s,q,t: puntero a nodo
	
	procedimiento actualizarNodo(x:puntero a nodo, golpes: entero) es 
		*x.g_totales_recibidos:= golpes 
		calcularPuntos()
	FP 

	procedimiento calcularPuntos(x1,x2:puntero a nodo; parciales1,parciales2,totales1,totales2: entero) es 
		si ((parciales1/2)+(totales1)) < ((parciales2/2)+(totales2)) entonces 
			*x2.puntos:= *x2.puntos + 5 
			*x1.puntos:= *x1.puntos + 1 
		sino 
			si ((parciales1/2)+(totales1)) > ((parciales2/2)+(totales2)) entonces 
				*x1.puntos:= *x1.puntos + 5 
				*x2.puntos:= *x2.puntos + 1 
			sino 
				*x1.puntos:= *x1.puntos + 2
				*x1.puntos:= *x1.puntos + 2 
			fsi 
		fsi 
	FP

	menor: entero 
	super: AN(15)

 Proceso
		Abrir E/(arch)
		Leer(arch,reg)
	 	prim:= nil 
		menor:= HV

		Mientras NFDA(arch) hacer 
			si prim = nil entonces 
				Nuevo(q); *q.heroe:= reg.heroe2; *q.g_totales_recibidos:= *q.g_totales_recibidos + reg.g_totales_heroe1; p:= q 
				*q.prox:= prim; prim:= q 
				Nuevo(q); *q.heroe:= reg.heroe1; *q.g_totales_recibidos:= *q.g_totales_recibidos + reg.g_totales_heroe2; s:= q
				*q.prox:= prim; prim:= q 
			sino 
				p:= prim; s:= prim; ant1:= nil; ant2:= nil  
				Mientras (p<>nil) y (*p.heroe <> reg.heroe1) hacer
					ant1:= s
					p:= *p.prox 
				FM 
				Mientras (s<>nil) y (*p.heroe <> reg.heroe2) hacer 
					ant2:= s
					s:= *s.prox 
				FM 
				
				si (*s.heroe = reg.heroe2) y (*p.heroe = reg.heroe1) entonces
					actualizarNodo(p,s,reg.g_parciales_heroe1,reg.g_parciales_heroe2,reg.g_totales_heroe1,reg.g_totales_heroe2)
				fsi 
				
				si (s = nil) entonces 
					Nuevo(q); *q.heroe:= reg.heroe2; *q.g_totales_recibidos:= *q.g_totales_recibidos + reg.g_totales_heroe1; s:= q
					ant2:= q; *q.prox:= s 
				sino 
					si (p = nil) entonces  
						Nuevo(q); *q.heroe:= reg.heroe1; *q.g_totales_recibidos:= *q.g_totales_recibidos + reg.g_totales_heroe2; p:= q
						ant1:= q; *q.prox:= p
					fsi 
					actualizarNodo(p,s,reg.g_parciales_heroe1,reg.g_parciales_heroe2,reg.g_totales_heroe1,reg.g_totales_heroe2)
				fsi 
			fsi 	 		
		FM 

		Mientras (p <> nil) hacer 
			si *p.g_totales_recibidos < 10 entonces 
				Esc("El superheroe ",*p.heroe," recibio menos de 10 golpes") //Consigna 3  
			fsi 
			
			si *p.puntos < menor entonces 
				menor:= *p.puntos 
				super:= *p.heroe 
			fsi 
		FM 

		Esc("El superheroe con menos puntos es: ",super) //Consigna 1

		Cerrar(arch)
	FProceso 
FAccion
		 


	
		
