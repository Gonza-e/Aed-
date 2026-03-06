Accion ej1 es 
	Ambiente 

	MOVIM = registro 
		mes: N(2)
		cli: N(13)
		cod: ("D","E","T")
		moneda: ("P","U","E")
		imp: N(13,2)
	Freg 
	arch: archivo de MOVIM ordenado por mes y cli 
	reg: MOVIM 

	nodo = registro 
		cli: N(13)
		tot: N(13,2)
		cant: N(5)
		IAF: N(13,2)
		prox: puntero a nodo 
	Freg 
	prim,primOrd,p,q,x,ant: puntero a nodo 

	resCli,resMes,cM,cT,cC,cEx,tr: entero 
	ext,mMes,mCli,mTotal: real  
	A: arreglo[1...3] de real 

	Procedimiento crearNodo()
		Nuevo(q); *q.cli:= resCli; *q.tot:= mCli; *q.cant:= cC; *q.IAF:= 0
	FP 

	Procedimiento cargaActualiza()
		si prim = nil entonces 
			crearNodo()
			*q.prox:= prim 
			prim:= q 
		sino 
			p:= prim 
			Mientras p<>nil y (*p.cli<>resCli) hacer 
				p:= *p.prox 
			FM 

			si p = nil entonces 
				crearNodo()
				*q.prox:= prim 
				prim:= q 
			sino 
				*p.tot:= *p.tot+mCli
				*p.cant:= *p.cant+cC 
			Fsi 
		Fsi 
	FP 

	Procedimiento cargaOrdenada()
		Nuevo(q); *q.cli:= *p.cli; *q.tot:= *p.tot; *q.cant:= *p.cant; *q.IAF:= (*p.tot) * (*p.cant)
		si primOrd = nil entonces 
			*q.prox:= primOrd
			primOrd:= q 
		sino 
			x:= primOrd; ant:= nil  
			Mientras x<>nil y ((*q.IAF < *x.IAF) o ((*q.IAF = *x.IAF) y (*q.cant < *x.cant))) hacer 
				ant:= x
				x:= *x.prox 
			FM 

			si primOrd = x entonces 
				*q.prox:= primOrd 
				primOrd:= q 
			sino 
				*ant.prox:= q 
				*q.prox:= x 
			Fsi 
		Fsi 
	FP 

	Procedimiento tratarReg()
		si reg.moneda = "U" entonces 
			monto:= convertir_en_pesos(reg.imp,"U"); cEx:= cEx+1
		sino 
			si reg.moneda = "E" entonces 
				monto:= convertir_en_pesos(reg.imp,"E"); cEx:= cEx+1
			sino 
				monto:= reg.imp 
			Fsi 
		Fsi 

		mCli:= mCli+monto 
		cC:= cC+1 

		segun reg.cod hacer 
			"D": A[1]:= A[1]+1
			"E": A[2]:= A[2]+1; ext:= ext+monto 
			"T": A[3]:= A[3]+1 
		Fsegun 

		monto:= 0 
	FP 

	Procedimiento corteCli()
		cargaActualiza()
		cM:= cM+cC; mMes:= mMes+mCli 
		cC:= 0; mCli:= 0 
		resCli:= reg.cli 
	FP 

	Procedimiento corteMes()
		Esc("Mes:",resMes,"Monto:",mMes,"Operaciones:",cM)
		Esc("Porcentaje de extracciones:",(ext/mMes)*100,"%")
		Para i:= 1 hasta 3 hacer 
			si A[i] > may entonces 
				may:= A[i]
				tr:= i
			Fsi 
		FPara 
		Esc("La transaccion mas solicitada en el mes es:",tran(i))
		Esc("Porcentaje de operaciones en moneda extranjera:",(cEx/cM)*100,"%")
		mTotal:= mTotal+mMes; cT:= cT+cM 
		arreglo()
		mMes:= 0; cM:= 0; ext:= 0; cEx:= 0
		resMes:= reg.mes 
	FP 

	Funcion tran(x:entero): AN 
		segun x hacer 
			1: tran:= "Deposito"
			2: tran:= "Extraccion"
			3: tran:= "Transferencia"
		Fsegun 
	FF 

	Procedimiento arreglo()
		Para i:= 1 hasta 3 hacer 
			A[i]:= 0 
		FPara 
	FP 

	Procedimiento init()
		arreglo()
		cM:= 0; cC:= 0; cT:= 0; cEx:= 0; tr:= 0; resCli:= reg.cli; resMes:= reg.mes 
		mMes:= 0; mTotal:= 0; exp:= 0; mCli:= 0
		prim:= nil; primOrd:= nil 
	FP 

	Proceso 
		Abrir E/(arch); Leer(arch,reg)
		init()

		Mientras NFDA(arch) hacer 
			si resMes <> reg.mes entonces 
				corteMes()
			sino 
				si resCli <> reg.cli entonces 
					corteCli()
				Fsi 
			Fsi 

			tratarReg()
			Leer(arch,reg)
		FM 
		corteMes(); Cerrar(arch)

		p:= prim 
		Mientras p<>nil hacer 
			cargaOrdenada()
			p:= *p.prox 
		FM 

		p:= primOrd 
		Para i:= 1 hasta 5 hacer 
			Esc("Posicion:",i,"Cliente:",*p.cli,"Total operado:",*p.tot,"Operaciones:",*p.cant,"IAF:",*p.IAF)
			p:= *p.prox 
		FPara 

		Esc("Monto anual:",mTotal,"Operaciones anuales:",cT)
	FProceso 
FAccion 






		

		