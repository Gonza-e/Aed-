Accion ejercicio1 (prim1: puntero a ruleta) es 
 Ambiente 
 	Fecha = registro 
	 dia: N(2)
	 mes: N(2)
	 anio: N(4)
	FReg

	COMPRAS = registro 
	 f_compra: Fecha 
	 dni_cliente: N(8)
	 c_articulos: N(3)
	 importe_total: N(8)
	FReg 
	arch: archivo de COMPRAS ordenado por f_compra
	reg: COMPRAS

	FIDELIZACION = registro
	 dni_cliente: N(8)
	 apeynom: AN(20)
	 f_adhesion: Fecha 
	 categoria: AN(10)
	FReg
	arch_ind: archivo de FIDELIZACION indexado por dni_cliente
	reg_ind: FIDELIZACION

 	chances, c_vueltas: entero 

	nodo = registro 
	 apeynom: AN(20)
	 chances_total: N(4)
	 prox, ant: puntero a nodo 
	FReg
 	q,aux,s,ult,prim: puntero a nodo 

	ruleta = registro 
	 chance: N(3)
	 prox: puntero a ruleta
	FReg
	p: puntero a ruleta

	Procedimiento cargar_cliente() es 
	 Nuevo(q)
	 *q.apeynom:= reg.apeynom
	 *q.chances_total:= 5 
	FProcedimiento

	Procedimiento chances_aleatorias() es 
	 	Si reg_ind.categoria = "Black" entonces
		 c_vueltas:= Tirar()
		 p:= prim1 
		 	Para i:= 1 a c_vueltas hacer 
		 	 p:= *p.prox 
			FM 
		 prim1:= p 
		 *q.chances_total:= *q.chances_total + *p.chance
		FSi 
	FProcedimiento
	 
 Proceso 
	 Abrir E/(arch)
	 Abrir E/(arch_ind)
	 Leer(arch,reg)
	 prim:= nil 
	 chances:= 0; c_vueltas:= 0 
	 Mientras NFDA(arch) hacer 
	 	 reg_ind.dni_cliente:= reg.dni_cliente
		 Leer(arch_ind,reg_ind)
		 	Si EXISTE entonces 
			 	Si prim = nil entonces 
				 cargar_cliente()
				 chances_aleatorias()
				 *q.ant:= nil 
				 *q.prox:= nil 
				 prim:= q 
				 ult:= q 
				Sino
				 s:= prim 
				 Mientras (s <> nil) hacer 
				 		Si *s.apeynom = reg_ind.apeynom entonces 
						 *s.chances_total:= *s.chances_total + (importe_total DIV 100)
						Sino 
						 aux:= prim 
						 cargar_cliente()
						 chances_aleatorias()
						 Mientras (aux <> nil) Y (*q.apeynom < *aux.apeynom) hacer 
							 aux:= *aux.prox 
							FM 
							Si aux = prim entonces
							 *q.prox:= *aux.prox 
							 *(*aux.prox).ant:= q 
							 *aux.prox:= q
							 *q.ant:= aux 
							Sino 
								Si aux = nil entonces 
								 *q.prox:= nil 
								 *ult.prox:= q 
								 *q.ant:= ult 
								 ult:= q 
								Sino 
								 *(*aux.prox).ant:= q 
								 *q.prox:= *aux.prox 
								 *aux.prox:= q 
								 *q.ant:= aux 
								FSi 
							FSi
						FSi     
					 s:= *s.prox 					 
					FM 
				FSi
			FSi
		 Leer(arch,reg)
		FM
	 Cerrar(arch)
	 Cerrar(arch_ind)
	FProceso 
FAccion
					

Accion ejercicio2 (prim: puntero a nodo) es
 Ambiente 
 	nodo = registro
	 dni_paciente: N(8)
	 apeynom: AN(30)
	 edad: N(2)
	 nro_cama: N(5) 
	 nro_habitacion: N(5)
	 temperaturas: arreglo[1...4] de enteros 
	 prox: puntero a nodo 
	FReg 
	p,e: puntero a nodo
	
	Funcion verificar(i, cant_med: enteros, A: arreglo[1...4] de enteros)
		Si i = 0 entonces 
		 	Si cant_med >= 1 entonces 
			 verificar:= falso 
			Sino 
			 verificar:= verdadero 
			FSi 
		Sino 
			Si A[i] > 36,5 entonces 
			 verificar:= verificar(i - 1, cant_med + 1, A[i])
			Sino 
			 verificar:= verificar(i - 1, cant_med, A[i])
			FSi 
		FSi 
	FFuncion
 
 Proceso 
	 p:= prim 
	 Mientras (p <> nil) hacer 
	 	 	Si verificar(4,0,*p.temperaturas) entonces 
			 Esc("Paciente sin temperatura alta", *p.apeynom)
			 e:= p 
			 p:= *p.prox 
			Sino 
			 p:= *p.prox 
			FSi 
		 disponer(e)
		FM 
	FProceso 
FAccion



	

				 
				 
