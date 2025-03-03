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

	prim,p,s,q: puntero a nodo
	
	procedimiento actualizarNodo(x:puntero a nodo; nombre:AN; golpe,puntos:entero ) es 
		x:= prim
		Mientras (x<>nil) y (*x.heroe <> nombre) hacer 
			x:= *x.prox 
		FM 
		
		si *x.heroe = nombre entonces 
			*x.g_totales_recibidos:= *x.g_totales_recibidos + golpes 
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
				Nuevo(q); *q.heroe:= reg.heroe1
				*q.g_totales_recibidos:= *q.g_totales_recibidos + (reg.g_totales_heroe2)
				*q.puntos:= 0 
				*q.prox:= prim; prim:= q 
				Nuevo(q); *q.heroe:= reg.heroe2
				*q.g_totales_recibidos:= *q.g_totales_recibidos + (reg.g_totales_heroe1)
				*q.puntos:= 0
				*q.prox:= prim; prim:= q 
				s:= prim; p:= *prim.prox 
			sino 
				actualizarNodo(p,reg.heroe1,reg.g_totales_heroe2)
				actualizarNodo(s,reg.heroe2,reg.g_totales_heroe1)

				si ((reg.g_parciales_heroe1/2)+(reg.g_totales_heroe1)) < ((reg.g_parciales_heroe1/2)+(reg.g_totales_heroe1)) entonces 
					*s.puntos:= *s.puntos + 5 
					*p.puntos:= *p.puntos + 1 
				sino 
					si ((reg.g_parciales_heroe1/2)+(reg.g_totales_heroe1)) > ((reg.g_parciales_heroe1/2)+(reg.g_totales_heroe1)) entonces 
						*p.puntos:= *p.puntos + 5
						*s.puntos:= *s.puntos + 1 
					sino 
						*p.puntos:= *p.puntos + 2 
						*s.puntos:= *s.puntos + 2
					fsi 
				fsi 
			fsi 

			Leer(arch,reg)
		FM 

		p:= prim 
		Mientras (p<>nil) hacer 
			
			si *p.puntos < menor entonces //Consigna 1
				menor:= *p.puntos 
				super:= *p.heroe 
			fsi 

			si *p.g_totales_recibidos < 10 entonces 
				Escribir("El superheroe ",*p.heroe," recibio menos de 10 golpes totales") //Consigna 3
			fsi 

			p:= *p.prox 
		FM 
	FProceso 
FAccion
		 


	
		
