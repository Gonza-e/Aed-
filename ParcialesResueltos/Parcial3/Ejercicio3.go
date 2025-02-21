Accion ejercicio3 (prim: puntero a nodo) es 
 Ambiente 
	CLIENTES = registro 
	 numero_cel: N(13)
	 dni: N(8)
	FinReg 
	
	arch: archivo de Clientes 
	reg: CLIENTES 

	nodo = registro 
	 iden_regalo: AN(255)
	 cant_disponible: N(4)
	 prox: puntero a nodo 
	FinReg 
	p: puntero a nodo 

	simple = registro
	 dni: N(8)
	 iden_regalo: AN(255)
	 prox: puntero a simple
	FinReg
	s,qs,primS: puntero a simple 

	doble = registro
	 iden_regalo: AN(255)
	 cant_disponible: N(4)
	 prox, ant: puntero a doble 
	FinReg
	d,qd,primD: puntero a doble

	doble2 = registro
	 dni: N(8)
	 prox, ant: puntero a doble2
	FinReg
	d2,qd2,primD2: puntero a doble2

	Procedimiento dar_regalo(x:puntero a simple) es 
	 	Mientras (*p.prox <> prim) y (k>0) hacer 
		 *x.iden_regalo:= *p.iden_regalo
		 *p.cant_disponible:= *p.cant_disponible - 1
		 	Si *p.cant_disponible = 0 entonces 
			 p:= *p.prox 
			Fsi 
		 k:= k-1
		FM
		Si *p.prox = prim 
		 *x.iden_regalo:= *p.iden_regalo
		 *p.cant_disponible:= *p.cant_disponible - 1
		 	Si *p.cant_disponible = 0 entonces 
		 	 p:= *p.prox
			Fsi 
	    Sino 
		 k:= k-1
		Fsi 
		Si p = prim entonces 
		 *x.iden_regalo:= ' '
		 bandera1:= true
		Fsi 
	FProcedimiento

	Procedimiento cargar_doble(/*prim*/x:puntero a nodo,/*nuevo*/y:puntero a nodo) es 
	 *y.prox:= *x.prox 
	 	Si *x.prox <> nil entonces 
		 *(*x.prox).ant:= y 
		Fsi 
	 *y.ant:= x 
	 *x.prox:= y 
	FProcedimiento

	bandera1: Booleano
	resg_dni: N(8)
	k: entero 
 Proceso 
	 Abrir E/(arch)
	 Leer(arch,reg)
	 primS:= nil; primD:= nil; primD2:= nil; p:= prim 
	 resg_dni:= reg.dni
	 bandera1:= false
	 	Mientras NFDA(arch) hacer 
		 	Si primS = nil entonces 
			 Nuevo(qs); *qs.dni:= reg.dni 
			 dar_regalo(qs)
			 resg_dni:= reg.dni
			Sino 
			 s:= primS 
			  	Mientras (s <> nil) y (*s.dni <> resg_dni) hacer
				 s:= *s.prox 
				FM 
			 Nuevo(qs); *qs.dni:= reg.dni 
			 dar_regalo(qs) 
			 *qs.prox:= primS 
			 primS:= qs 
			FSi 
		FM 

		Si NO bandera1 entonces 
		 	Mientras (p <> *p.prox) hacer 
		 	 Nuevo(qd); *qd.iden_regalo:= *p.iden_regalo
			 *qd.cant_disponible:= *p.cant_disponible
			 cargar_doble(primD,qd)
			 p:= p.prox 
			FM 
		 Nuevo(qd); *qd.iden_regalo:= *p.iden_regalo
		 *qd.cant_disponible:= *p.cant_disponible
		 cargar_doble(primD,qd)
		Fsi 

	 s:= primS 
	 	Mientras (s <> nil) hacer 
		 	Si *s.iden_regalo = ' ' entonces 
			 Nuevo(qd2); *qd2.dni:= *s.dni 
			 cargar_doble(primD2,qd2)
			Fsi 
		 s:= *s.prox 
		FM 
	 Cerrar(arch)
	FProceso 
FAccion
	

	

	 

			 

 



