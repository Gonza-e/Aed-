Accion Parcial2024A (prim: puntero a nodo, prim_circ: punteroa circ) es 
 Ambiente 
	Fecha = registro 
		dia: N(2)
		mes: N(2)
		anio: N(4)
	FReg

	nodo = registro 
		id_usuario: AN(20)
		f_opinion: Fecha
		calificacion: (1..5)
		serv_calificado: (1..20)
		prox: puntero a nodo
	FReg
	p,n: puntero a nodo 

	circ = registro 
		cod_descuento: N(4)
		porc_descuento: N(1,2)
		prox: puntero a circ
	FReg
	c: puntero a circ 

	doble = registro 
		id_usuario: AN(20)
		apeynom: AN(30)
		correo: AN(30)
		c_opiniones: N(3)
		prom_opiniones: N(1,2)
		descuento: N(1,2)
		prox,ant: puntero a doble
	FReg
	d,q,s,prim_doble,ult: puntero a doble

	USUARIOS = registro
		id_usuario: AN(20)
		apeynom: AN(30)
		dni: N(8)
		direccion: AN(30)
		correo: aN(30)
		cat: ("SN","SILVER","GOLD")
	FReg
	arch: archivo de USUARIOS indexado por id_usuario
	reg: USUARIOS

	Procedimiento cargar_doble() es 
		si prim_doble = nil entonces
			*q.ant:= nil 
			*q.prox:= nil 
			prim:= q 
			ult:= q  
		sino 
		 	d:= prim_doble 
			Mientras (d <> nil) y (*q.prom_opiniones > *p.prom_opiniones) hacer 
				d:= *d.prox 
			FM 
			si d = prim_doble entonces
				*q.ant:= nil 
				*q.prox:= *d.prox 
				*d.ant:= q 
				prim:= q 
			sino 
				si d = ult entonces 
					*q.prox:= *ult.prox 
					*q.ant:= ult 
					*ult.prox:= q 
					ult:= q 
				sino 
					*(*d.prox).ant:= q 
					*q.prox:= *d.prox 
					*d.prox:= q 
					*q.ant:= d 
				fsi 
			fsi 
		fsi  

	cant_op, num_op: entero 
	f_comp: Fecha
	bandera1, bandera2: boolean

	Proceso 
		Abrir E/(arch)
		p:= prim 
		n:= *p.prox 
		c:= prim_circ
		prim_doble:= nil
		num_op:= *p.calificacion; cant_op:= 1;
		f_comp:= *p.f_opinion
		Mientras (p<>nil) hacer
			Si *p.id_usuario = *n.id_usuario entonces  
				Si *p.f_opinion < *n.f_opinion entonces 
					f_comp:= *n.f_opinion
				Fsi 
				cant_op:= cant_op + 1
				num_op:= num_op + *n.calificacion
				n:= *n.prox 
			Sino 
				reg.id_usuario:= *p.id_usuario
				Leer(arch,reg)
				Si existe entonces
					Nuevo(q) 
					*q.id_usuario:= reg.id_usuario; *q.apeynom:= reg.apeynom; *q.correo:= reg.correo 
					*q.cant_op:= cant_op; *q.prom_opiniones:= (num_op/cant_op); *q.descuento:= *c.porc_descuento
					c:= *c.prox 
				Sino 
					Escribir("Usuario no existe")
				Fsi 
				p:= n
				n:= *n.prox 
				cant_op:= 0; num_op:= 0
				cargar_doble()
			Fsi 
		FM 

		d:= prim_doble
		s:= ult 

		Para i:= 1 a 10 hacer 
			Escribir("Usuario ",i," con mayor promedio: ",*d.id_usuario)
			Escribir("Usuario ",i," con menor promedio: ",*s.id_usuario)
			d:= *d.prox 
			s:= *s.ant
		FPara 
	FProceso 
FAccion
	

			


