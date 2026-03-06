Accion ej1 es 
	Ambiente 

	Fecha = registro 
		dia: N(2)
		mes: N(2)
		año: N(4)
	Freg 

	Compras = registro 
		f_compra: Fecha 
		dni: N(8)
		cant: N(3)
		imp: N(7)
		jubilado: ("S","N")
		pago: ("QR","efectivo","debito","credito")
	Freg 
	archC: archivo de Compras ordenado por f_compra
	regC: Compras 

	Cliente = registro 
		dni: N(8)
		ayn: AN(30)
		f_add: Fecha 
		cat: ("regular","black")
	Freg 
	arch: archivo de Cliente indexado por dni 
	reg: Cliente 

	nodo = registro 
		ayn: AN(30)
		chances: N(5)
		prox: puntero a nodo 
	Freg 
	prim,p,q,ant: puntero a nodo 

	num,año,desc: entero 
	cat,nom: AN 

	Procedimiento obtInfo()
		reg.dni:= regC.dni; Leer(arch,reg)

		año:= reg.f_add.año 
		cat:= reg.cat 
		nom:= reg.ayn
	FProcedimiento

	Procedimiento cargaOrdenada()
		si prim = nil entonces 
			crearNodo()
			*q.prox:= prim 
			prim:= q 
		sino 
			p:= prim; ant:= nil 
			Mientras p<>nil y (nom > *p.ayn) hacer 
				ant:= p 
				p:= *p.prox 
			FM 

			si nom = *p.ayn entonces 
				actualizar()
			sino 
				crearNodo()
				si prim = p entonces 
					*q.prox:= prim 
					prim:= q 
				sino 
					*ant.prox:= q 
					*q.prox:= p 
				Fsi 
			Fsi 
		Fsi 
	FProcedimiento

	Procedimiento crearNodo()
		Nuevo(q); *q.ayn:= nom
		dif:= año - regC.f_compra.año 

		si dif > 10 entonces 
			*q.chances:= 10 
		sino 
			si dif > 3 y dif < 10 entonces 
				*q.chances:= 5 
			sino 
				*q.chances:= 2 
			Fsi 
		Fsi 
	FProcedimiento
	
	Procedimiento actualizar()
		*q.chances:= *q.chances + (regC.imp div 100)

		si cat = "black" entonces 
			num:= Tirar()
			*q.chances:= *q.chances + A[num]
		Fsi 
	FProcedimiento

	Procedimiento mostrarDescuento(socio:AN)
		desc:= 0 
		si regC.jubilado = "S" entonces 
			desc:= desc+15
		Fsi 
		si regC.pago = "QR" entonces 
			desc:= desc+20
		Fsi  
		Esc("DNI:",regC.dni,"Fecha de compra:",regC.f_compra,"Socio:",socio,"Descuento:",desc)

	FProcedimiento

	Proceso 
		Abrir E/(archC); Leer(archC,regC)
		Abrir E/(arch)

		num:= 0; desc:= 0; año:= 0
		nom:= ""; cat:= ""
		prim:= nil 

		Mientras NFDA(archC) hacer 
			reg.dni:= regC.dni; Leer(arch,reg)
			si EXISTE entonces 
				obtInfo()
				cargaOrdenada()
				mostrarDescuento("SI")
			sino 
				mostrarDescuento("NO")
			Fsi 
			Leer(archC,regC)
		FM 
		Cerrar(arch); Cerrar(archC)
	FProceso 
FAccion 


Accion ej2 es 
	Ambiente 

	nodo = registro 	
		num: N(9)
		ant,prox: puntero a nodo 
	Freg 
	p,q: puntero a nodo 

	Procedimiento eliminar()
		si prim = p entonces 
			prim:= *prim.prox 
			*prim.ant:= nil 
		sino 
			si p = ult entonces
				ult:= *ult.ant 
				*ult.prox:= nil 
			sino 
				si prim = ult entonces 
					prim:= nil
					ult:= nil 
				sino 
					*(*p.ant).prox:= *p.prox 
					*(*p.prox).ant:= *p.ant 
				Fsi 
			Fsi 
		Fsi 
		q:= p; p:= *p.prox
		Disponer(q)
	FProcedimiento

	Funcion pmas(valor,num): booleano 
		si num < 10 entonces 
			si num = valor entonces 
				pmas:= verdadero 
			sino 
				pmas:= falso 
			Fsi 
		sino 
			pmas:= pmas(valor,num div 10)
		Fsi 
	FFuncion 

	num: entero 

	Proceso 
		p:= prim 

		Esc("Ingrese un numero: "); Leer(num)
		Mientras p<>nil hacer 	
			si pmas(num,*p.num) entonces 
				eliminar()
			sino
				p:= *p.prox 
			Fsi 
		FM 

	FProceso 
FAccion



