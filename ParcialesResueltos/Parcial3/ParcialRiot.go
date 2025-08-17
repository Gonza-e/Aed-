Accion RIOT (prim: puntero a nodo) es 
 Ambiente 
 	CLIENTE = registro 
		nom_cliente: AN(45)
		localizacion: ("LAN","EUW","LAS")
		nivel: (1...500)
	FReg
	arch: archivo de CLIENTE
	reg: CLIENTE

	nodo = registro 
		servidor: N(10)
		localizacion: ("LAN","EUW","LAS")
		estado: ("en mantenimiento", "libre")
		cant_usu: (1...10)
		latencia: (1...1500)
		prox: puntero a nodo 
	FReg
	p: puntero a nodo 

	datos = registro 
		nom_cliente: AN(45)
		servidor: N(10)
		localizacion: ("LAN", "EUW", "LAS")
		nivel: (1...500)
		prox: puntero a datos
	FReg
	prim_datos, s, q: puntero a datos 

	libres = registro 
		servidor: N(10)
		localizacion: ("LAN","EUW","LAS")
		latencia: (1...1500)
		prox, ant: puntero a libres 
	FReg
	prim_libres, ult, p1, q1: puntero a libres 

	Procedimiento cargar_datos() es 
	 	*q.nom_cliente:= reg.nom_cliente; *q.servidor:= *p.servidor; *q.localizacion:= reg.localizacion; *q.nivel:= reg.nivel
	FProcedimiento 

 	Procedimiento guardar() es 
		Nuevo(q); cargar_datos()
		*q.prox:= prim_datos 
		prim_datos:= q 
	FProcedimiento
  
	Procedimiento cargar_doble() es 
	 	Nuevo(q1); *q.servidor:= *p.servidor; *q.localizacion:= *p.localizacion; *q.latencia:= *p.latencia
	 	Si prim_libres = nil entonces 
			*q1.prox:= nil 
			*q1.ant:= nil 
			prim_libres:= q1 
			ult:= q1 
		Sino 
			p1:= prim_libres
			Mientras (s<>nil) Y (*q1.latencia > *p1.latencia) hacer 
				p1:= *p1.prox 
			FM 
			Si prim_libres = p1 entonces 
				*q1.prox:= p1 
				*q1.ant:= nil 
				*p1.ant:= q1 
				prim_libres:= q1 
			Sino 
				Si p1 = nil entonces 
					*q1.prox:= nil 
					*q1.ant:= ult 
					*ult.prox:= q1 
					ult:= q1
				Sino 
					*q1.prox:= *p1.prox 
					*q1.ant:= p1 
					*(*p1.prox).ant:= q1 
					*p1.prox:= q1
				FSi 
			FSi 
		FSi 
	FProcedimiento

	Procedimiento cargar_servidor() es 
		p:= prim
		Mientras (p<>nil) Y ((*p.cant_usu = 10) O (*p.estado <> "libre")) hacer 
			p:= *p.prox 
 	 	FM 
 		Si (*p.estado = libre) Y (*p.cant_usu < 10) entonces
 		 	*p.cant_usu:= *p.cant_usu + 1 
 		FSi 
	FProcedimiento

	Proceso
		Abrir E/(arch)
		Leer(arch,reg)
		prim_datos:= nil 
		prim_libres:= nil 
		Mientras NFDA(arch) hacer 
	 	 	Si prim_datos = nil entonces
				cargar_servidor()
				guardar() 
			Sino 
				s:= prim 
				Mientras (s<>nil) Y (*s.nom_cliente <> reg.nom_cliente) hacer 
					s:= *s.prox 
				FM 
				Si s = nil entonces 
					cargar_servidor()
					guardar()
				Sino 
				 	Esc("Ya esta asignado")
				FSi 
			FSi
		 	Leer(arch,reg)
		FM 
		p:= prim 
		Mientras (p<>nil) hacer
	 	 	Si (*p.cant_usu = 10) Y (*p.estado <> "libre") entonces 
			 	cargar_doble()
			FSi 
		 	p:= *p.prox
		FM 
	FProceso 
FinAccion


		 	 

			 

	 	 
				 
