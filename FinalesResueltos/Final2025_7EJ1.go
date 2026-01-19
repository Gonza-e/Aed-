Accion ejercicio1 es 
 Ambiente 
	Fecha = registro 
		dia: N(2)
		mes: N(2)
		año: N(4)
	Freg

	VENTAS = registro 
		codrepartidor: AN(10)
		codproducto: N(6)
		cantidad: N(4)
		precio: N(7,2)
		fchVta: Fecha 
	Freg
	arch_v: archivo de VENTAS ordenado por codrepartidor y codproducto
	reg_v: VENTAS  

	COMISIONES_PRODUCTO = registro 
		codrepartidor: AN(10)
		codproducto: N(6)
		comision: N(7,2)
	Freg
	arch_c: archivo de COMISIONES_PRODUCTO indexado por codrepartidor y codproducto
	reg_c: COMISIONES_PRODUCTO

	nodo = registro 
		codrepartidor: AN(10)
		comision_total: N(7,2)
		comision: N(7,2)
		prox: puntero a nodo 
	Freg
	prim,ant,q,p: puntero a nodo 

	nodo2 = registro 
		codrepartidor: AN(10)
		comision: N(7,2)
		prox: puntero a nodo 
	Freg
	prim2,p2,q2: puntero a nodo2 

	procedimiento cargaOrdenada() 
		si prim = nil entonces 
			prim:= q 
			*q.prox:= nil
		sino 
			ant:= nil; p:= prim 
			Mientras p <> nil y (*q.comision < *p.comision) hacer 
				ant:= p 
				p:= *p.prox 
			FM 
			
			si p = prim entonces 
				*q.prox:= p 
				prim:= q 
			sino 
				*q.prox:= *ant.prox 
				*ant.prox:= q 
			fsi 
		fsi 
	FProcedimiento

	procedimiento corte_repartidor() 
		corte_producto()
		Nuevo(q2); *q2.codrepartidor:= res_repartidor; *q2.comision:= comRepartidor
		si comRepartidor > 10000 entonces 
			*q2.comision:= *q2.comision + (*q2.comision * 0.10)
		fsi 
		cargaOrdenada()
		comRepartidor:= 0 
		res_repartidor:= reg_v.codrepartidor
	FProcedimiento

	procedimiento corte_producto()
		reg_c.codrepartidor:= res_repartidor 
		reg_c.codproducto:= res_producto
		Leer(arch_c,reg_c)
		Nuevo(q2); *q2.codrepartidor:= res_repartidor
		si EXISTE entonces 
			*q2.comision:= reg_c.comision
			comRepartidor:= comRepartidor + reg_c.comision 
		sino 
			*q2.comision:= generica 
			comRepartidor:= comRepartidor + generica 
		fsi 
		*q2.prox:= prim2 
		prim2:= q2
		res_producto:= reg_v.codproducto
	FProcedimiento

	generica, comRepartidor, res_producto: entero 
	res_repartidor: AN 	

	Proceso 

		Abrir E/ (arch_v); Leer(reg_v)
		Abrir E/ (arch_c)
		prim:= nil; prim2:= nil 
		res_producto:= reg_v.codproducto
		res_repartidor:= reg_v.codrepartidor

		Esc("Ingrese una comision generica para todos los repartidores: ")
		Leer(generica)

		Mientras NFDA(arch_v) hacer 
			si reg_v.codrepartidor <> res_repartidor entonces 
				corte_repartidor()
			sino 
				si reg_v.codproducto <> res_producto entonces 
					corte_producto()
				fsi 
			fsi
			Leer(arch_v,reg_v) 
		FM 
		corte_repartidor()
		Cerrar(arch_v); Cerrar(arch_c)

		p2:= prim2 
		Mientras p2 <> nil hacer 
			Esc("Codigo repartidor: ",*p2.codrepartidor)
			Esc("Comision:",*p2.comision)
			p2:= *p2.prox 
		FM 

		Esc("Ranking de comisiones a cobrar")
		p:= prim 
		Mientras p <> nil hacer 
			Esc("Codigo repartidor:",*p.codrepartidor)
			Esc("Comision: ",*p.comision)
			p:= *p.prox 
		FM 
	FProceso 
FAccion


