Accion parcialB (prim: puntero a nodo, prim_circ: puntero a nodo_circ) es 
 Ambiente 
	Fecha = registro 
		dia: N(2)
		mes: N(2)
		anio: N(4)
	FReg

 	nodo = registro 
		id_usuario: AN(20)
		f_opinion: Fecha
		calificacion: (1...5)
		serv_calificado: (1...20)
		prox: puntero a nodo
	FReg 
	p, s: puntero a nodo 

	nodo_doble = registro
		id_usuario: AN(20)
		apeynom: AN(30)
		correo: AN(20)
		c_opiniones: N(2)
		prom_opiniones: N(2,2)
		descuento: N(1,2)
		prox, ant: puntero a nodo_doble
	FReg
	prim1, p1, q1, ult: puntero a nodo_doble

	nodo_circ = registro 
		cod_descuento: AN(10)
		porc_descuento: N(1,2)
		prox: puntero a nodo_circ
	FReg
	p2: puntero a nodo_circ

	usuarios = registro 
		id_usuario: AN(20)
		apeynom: AN(30)
		dni: N(8)
		direccion: AN(20)
		correo: AN(20)
		cat: ("SN","silver","gold")
	FReg
	arch: archivo indexado por id_usuario
	reg: usuarios

	comparar: Fecha
	desc: N(1,2)
	cant_op, num_op, k: entero 
	bandera1, bandera2, asignar: logico

	Procedimiento obtener_desc() es 
	 	Mientras (*p2.prox <> nil) y (k>0) entonces 
			desc:= *p2.porc_descuento 
			k:= k-1
		FM
	FP
	Procedimiento agregar_nodo() es 
		Si prim1 = nil entones 
			*q.ant:= nil 
			*q.prox:= nil 
			prim:= q 
			ult:= q 
		Sino 
		 p1:= prim1 
		 	Mientras (p1<>nil) y (*q1.c_opiniones > *p1.c_opiniones) hacer 
			 	*p1:= *p1.prox 
			FM 
			Si prim1 = p1 entonces 
				*q1.ant:= nil 
				*q1.prox:= p1
				p1:= q1 	
				prim1:= q1 
			Sino 
				Si p1 = ult entonces 
					*q1.ant:= ult 
					*q1.prox:= *ult.prox 
					ult:= q1 
				Sino 
					*(*p1.prox).ant:= q1 
					*q1.prox:= *p1.prox 
					*p1.prox:= q1
					*q1.ant:= p1  
				FSi
			FSi
		FSi
	FP

	Procedimiento crear_nodo() es 
		Nuevo(q1)
		*q1.id_usuario:= reg.id_usuario; *q1.apeynom:= reg.apeynom; *q1.correo:= reg.correo
		*q1.c_opiniones:= cant_op; *q1.prom_opiniones:= num_op/cant_op; *q1.descuento:= desc
	FP

	Proceso 
		Abrir E/(arch)
		k:= 1
		p:= prim
		s:= *p.prox 
		p2:= prim_circ
		prim1:= nil
		bandera1:= falso
		Mientras (p<>nil) hacer 

	 	 	Si *p.id_usuario = *s.id_usuario entonces 
				cant_op:= cant_op + 1 
				num_op:= num_op + *p.calificacion
				Si *p.f_opinion < *s.f_opinion entonces 
				 	comparar:= *p.fecha 
				Sino 
				 	comparar:= *s.fecha 
				FSi 
			 	s:= *s.prox 
			Sino 
			 	bandera1:= verdadero
			FSi 
			
			Si bandera1 entonces 
				reg.id_usuario:= *p.id_usuario
				Leer(arch,reg)
			 	Si existe entonces 
				 	asignar:= bonificacion(reg.cat, comparar)
				 	Si asignar entonces
						crear_nodo() 
						agregar_nodo()
					FSi 
				Sino 
				 	Esc("Usuario no existe")
				FSi 
				p:= s 
				s:= *s.prox 
				bandera1:= falso 
				k:= 1 
				cant_op:= 0; num_op:= 0; comparar:= 0 
			FSi 
		FM
	 
		s:= ult
		p:= prim1 
		k:= 10
		Mientras (p<>nil) y (k>0) hacer 
			Esc("Usuario con mayor porcentaje: ", *p.id_usuario, *p.apeynom, *p.correo)
			Esc("Usuario con menor porcentaje: ", *s.id_usuario, *s.apeynom, *s.correo)
			p:= *p.prox 
			s:= *s.ant 
			k:= k-1 
		FM 
	FProceso 
FAccion

					 
Accion Ejercicio2 (prim: puntero a nodo, ult: puntero a nodo) es 
 Ambiente
	Fecha = registro 
		dia: N(2)
		mes: N(2)
		anio: N(4)
	FReg

 	nodo = registro
		id_compra: N(10)
		f_compra: Fecha 
		metodo_pago: AN(3)
		nro_tarjeta: arreglo[1...20] de enteros 
		monto: N(7)
	FReg
	p,s: puntero a nodo 

	Procedimiento eliminar() es 
	 	Si prim = nil entonces 
		 	Esc("La lista está vacia")
		Sino 
		 	Si prim = p entonces 
				prim:= *p.prox 
				*prim.ant:= nil 
			Sino 
				Si p = ult entonces 
					ult:= *ult.ant 
					*ult.prox:= *p.prox 
				Sino 
				 	Si prim = ult entonces 
						prim:= nil 
						ult:= nil 
					Sino 
						*(*p.prox).ant:= *p.ant 
						*(*p.ant).prox:= *p.prox 
					FinSi 
				FinSi 
			FinSi 
		FinSi 
		s:= p 
		p:= *p.prox 
		Disponer(s)
	FP 

	Funcion verificar(A:arreglo[1...20] de enteros,suma:entero,i:entero): booleano 
	 	Si i = 0 entonces
			Si (suma MOD 5) = 0 entonces 
			 	verificar:= verdadero
			Sino 
			 	verificar:= falso 
			Fsi 
		Sino 
			Si (i MOD 2 <> 0) entonces 
			 	A[i]:= A[i] * 2
				Si A[i] > 9 entonces 
				 	A[i]:= (A[i] MOD 10) + (A[i] DIV 10) 
				Fsi 
			 	verificar:= verificar(suma + A[i], i-1)
			Sino 
			 	verificar:= verificar(suma, i-1)
			Fsi 
		Fsi 
	FFuncion

	Proceso 
		p:= prim 
	 	Mientras (p <> nil) hacer 
		 	Si (*p.metodo_pago = "TC") entonces
				Si (verificar(*p.nro_tarjeta,0,20)) entonces
				 	Esc("Cumple")
				Sino 
				 	eliminar()
				Fsi 
			Fsi 
		 	p:= *p.prox 
		FM 
	FProceso 
FAccion
			  
			


				 
		 



	 	 
	 	 



	 
	
