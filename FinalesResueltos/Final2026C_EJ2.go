Accion ej2 (desc: arreglo[0...24] de AN) es 
	Ambiente 
	
	sec: secuencia de caracteres 
	s: caracter 
	pat,mot: AN 
	tip,ide,monto,mProm,mTotal,cM,opcion: entero 

	A: arreglo[0...24] de DATO 
	B: arreglo[0...24] de entero 

	DATO = registro 
		monto: N(10)
		cant: N(4)
	Freg 

	PAGOS = registro 
		id: N(6)
		fecha: Fecha 
		enTermino: booleano 
	Freg 
	arch: archivo de PAGOS indexado por id 
	reg: PAGOS 

	nodo = registro 
		id: N(6)
		patente: AN(7)
		tipo: N(2)
		monto: N(10)
		motivo: AN(200)
		prox: puntero a nodo 
	Freg 
	prim,p,q,ant: puntero a nodo 

	Procedimiento info()
		Para i:= 1 hasta 6 hacer 
			ide:= ide*10 + conv(s)
			Avz(sec,s)
		FPara
		
		Para i:= 1 hasta 7 hacer 
			pat:= concatenar(pat,s)
			Avz(sec,s)
		FPara 

		Para i:= 1 hasta 2 hacer 
			tip:= tip*10 + conv(s)
			Avz(sec,s)
		FPara 

		Para i:= 1 hasta 8 hacer 
			monto:= monto*10 + conv(s)
			Avz(sec,s)
		FPara 

		Mientras (s <> "-") y s <> "*" hacer
			mot:= concatenar(mot,s)
			Avz(sec,s) 
		FM 
		Avz(sec,s)

		mTotal:= mTotal+monto 
		cM:= cM+1
		A[tip].cant:= A[tip].cant+1; A[tip].monto:= A[tip].monto+monto

	FP 

	Procedimiento cargar()
		Nuevo(q); *q.id:= ide; *q.patente:= pat; *q.tipo:= tip; *q.monto:= monto; *q.motivo:= mot 

		si prim = nil entonces 
			*q.prox:= prim 
			prim:= q 
		sino 
			p:= prim; ant:= nil 
			Mientras (p<>nil) y ((*q.tipo < *p.tipo) o ((*q.tipo = *p.tipo) y (*q.id < *p.id))) hacer 
				ant:= p 
				p:= *p.prox 
			FM 

			si prim = p entonces 
				*q.prox:= prim 
				prim:= q 
			sino 
				*ant.prox:= q 
				*q.prox:= p 
			Fsi 
		Fsi 

		ide:= 0; pat:= ""; tip:= 0; monto:= 0; mot:= ""
	FP 

	Procedimiento conA()
		p:= prim; mProm:= mTotal DIV cM  

		Mientras p<>nil hacer 
			si *p.monto > mProm entonces 
				Esc("ID:",*p.id,"PATENTE:",*p.patente,"TIPO:",desc(*p.tipo),"MONTO:",*p.monto)
			Fsi 
			p:= *p.prox 
		FM 
	FP 

	Procedimiento conB()
		Ambiente
			idUSU: entero 
		Proceso 
			Esc("Ingrese la ID de la multa:"); Leer(idUSU)

			p:= prim 
			Mientras p<>nil y (*p.id<>idUSU) hacer 
				p:= *p.prox
			FM 

			si p = nil entonces 
				Esc("No se encontro la multa")
			sino 
				reg.id:= idUSU; Leer(arch,reg)
				si EXISTE entonces 
					si reg.enTermino entonces 
						*p.monto:= *p.monto - (*p.monto * 0.15)
					sino 
						*p.monto:= *p.monto + (*p.monto * 0.1)
					Fsi 
				sino 
					Esc("No se registro ningun pago")
				Fsi 
			Fsi 
		FProceso 
	FP 

	Procedimiento conC()
		Ambiente 
			total: entero
			patente: AN  
		Proceso 
			total:= 0 
			Esc("Ingrese una patente:"); Leer(patente)
			Para i:= 0 hasta 24 hacer 
				B[i]:= 0
			FPara 

			p:= prim 
			Mientras p<>nil hacer 
				si *p.patente = patente entonces 
					B[*p.tipo]:= B[*p.tipo]+*p.monto 
				Fsi 
				total:= total+*p.monto 
			FM 

			Para i:= 0 hasta 24 hacer 
				si B[i] <> 0 entonces 
					Esc("Tipo:",desc(i),"Monto por tipo:",B[i])
				Fsi 
			FPara 

			Esc("Total final a abonar:",total)
		FProceso 
	FP 

	Procedimiento conD()
		Ambiente 
			tMen,men,cMen: entero 
		Proceso 
			tMen:= 0; men:= HV; cMen:= 0 
			
			Para i:= 0 hasta 24 hacer 
				Esc("Tipo:",desc(i),"Monto:",A[i].monto)
			
				si A[i].cant < cMen entonces 
					cMen:= A[i].cant 
					tMen:= i 
					men:= A[i].monto 
				Fsi 
			FPara 

			Esc("Tipo menos frecuente:",desc(tMen),"Promedio de monto a pagar:",men DIV cMen)
		FProceso
	FP 

	Procedimiento init()
		pat:= ""; ide:= 0; monto:= 0; tip:= 0; mot:= ""
		mProm:= 0; mTotal:= 0; cM:= 0, opcion:= 0
	FP 

	Proceso 
		Arr(sec); Avz(sec,s)

		Para i:= 0 hasta 24 hacer 
			A[i].cant:= 0; A[i].monto:= 0
			B[i]:= 0 
		FPara 

		init()

		Mientras s <> "*" hacer 
			info()
			cargar()
		FM 
		Cerrar(sec)

		
		Repetir 
			Esc("Ingrese opcion | 0: Salir | 1: Multa con monto mayor al promedio | 2: Buscar multa | 3: Buscar patente | 4: Multa menos frecuente")
			Leer(opcion)

			segun opcion hacer 
				1: conA() 
				2: conB() 
				3: conC()
				4: conD()
			Fsegun 

		Hasta (opcion = 0)
	FProceso 
FAccion 



					



