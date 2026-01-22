Accion ej1 es 
	Ambiente 

	Fecha = registro 
		dia: N(2)
		mes: N(2)
		año: N(4)
	Freg

	VENTAS = registro 
		codRep: (1...20)
		codProd: AN(20)
		cant: N(4)
		precio: N(6,2)
		fchVta: Fecha 
	Freg 
	arch_v: archivo de VENTAS ordenado por codRep y codProd 
	reg_v VENTAS 

	REPARTIDORES = registro 
		codRep: (1...20)
		nombre: AN(20)
		porc: N(1,2)
		domicilio: AN(30)
	Freg 
	arch_r: archivo de REPARTIDORES indexado por codRep 
	reg_r: REPARTIDORES

	COMISIONES = registro 
		codRep: (1...20)
		codProd: AN(20)
		com: N(1,2)
	Freg 
	arch_c: archivo de COMISIONES indexado por codRep y codProd 
	reg: COMISIONES

	PRODUCTOS = registro 
		codProd: AN(20)
		desc: AN(50)
		stock: N(4)
		com: N(1,2)
		ultfch: Fecha 
	Freg 
	arch_p: archivo de PRODUCTOS indexado por codProd 
	reg_p: PRODUCTOS 

	nodo = registro 
		nombre: AN(20)
		comtotal: N(6,2)
		prox: puntero a nodo 
	Freg 
	prim,p,q,ant: puntero a nodo 

	Procedimiento cargaOrdenada()
		si prim = nil entonces 
			*q.prox:= prim 
			prim:= q 
		sino 
			ant:= nil; p:= prim 
			Mientras p <> nil y *q.comtotal < *p.comtotal hacer 
				ant:= p 
				p:= *p.prox 
			FM 

			si p = prim entonces 
				*q.prox:= p 
				prim:= q 
			sino 
				*ant.prox:= q 
				*q.prox:= p 
			Fsi 
		Fsi 
	FProcedimiento

	Procedimiento tratarReg()
		monto:= monto + (reg_v.cant * reg_v.precio)
	FProcedimiento 

	Procedimiento aplicarCom()
		reg_c.codRep:= res_rep; Leer(arch_c,reg_c)
		si EXISTE entonces 
			monto:= monto * reg_c.com 
		sino 
			reg_p.codProd:= res_prod 
			si EXISTE entonces 
				monto:= monto * reg_c.com 
			Fsi 
		Fsi 
	FProcedimiento

	Procedimiento totalCobrar()
		reg_r.codRep:= res_rep; Leer(arch_r,reg_r)
		si EXISTE entonces 
			nombre:= reg_r.nombre 
			si montoRep > 10000 entonces 
				montoRep:= montoRep + (montoRep * reg_r.porc)
			Fsi 
		Fsi 
	FProcedimiento 

	Procedimiento corteProd()
		aplicarCom()
		Esc("Codigo de producto:",res_prod,"Comision:",monto)
		montoRep:= montoRep + monto 
		monto:= 0 
		res_prod:= reg_v.prod 
	FProcedimiento

	Procedimiento corteRep()
		corteProd()
		totalCobrar()
		Esc("Codigo de repartidor:",res_rep,"Comision:",montoRep)
		Nuevo(q); *q.nombre:= nombre; *q.comtotal:= montoRep 
		cargaOrdenada()
		montoRep:= 0 
		res_rep:= reg_v.codRep 
	FProcedimiento

	Procedimiento init()
		montoRep:= 0; monto:= 0
		res_prod:= reg_v.codProd; res_rep:= reg_v.codRep  
		nombre:= " "
		prim:= nil 
	FProcedimiento

	Proceso 
		Abrir E/(arch_v); Leer(arch_v,reg_v)
		Abrir E/(arch_r); Abrir E/(arch_p); Abrir E/(arch_c)

		init()

		Mientras NFDA(arch_v) hacer 
			si res_rep <> reg_v.codRep entonces 	
				corteRep()
			sino 
				si res_prod <> reg_v.codProd entonces 
					corteProd()
				Fsi 
			Fsi 

			tratarReg()
			Leer(arch_v,reg_v)
		FM 
		corteRep()
		Cerrar(arch_v); Cerrar(arch_r); Cerrar(arch_p); Cerrar(arch_c)

		p:= prim; k:= 1 
		Mientras p <> nil hacer 
			Esc("Top",k)
			Esc("Repartidor:",*p.nombre,"Comision:",*p.comtotal)
			k:= k+1 
			p:= *p.prox 
		FM 
	FProceso 
FAccion