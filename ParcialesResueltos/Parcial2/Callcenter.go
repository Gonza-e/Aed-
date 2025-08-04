Accion Ejercicio2 es 
 Ambiente
	Fecha = registro 
		dia: N(2)
		mes: N(2)
		anio: N(4)
	FinRegistro
	RECLAMOS = registro 
		cod_recl: N(10)
		f_reclamo: Fecha 
		mailcliente: AN(20)
		urgencia: ("A","M","B")
		detalle: AN(100)
		region: N(2)
	FinRegistro 
	arch: archivo de RECLAMOS
	reg: RECLAMOS
	i, j, mayor, mes_may, c_meses: entero 
	Funcion obtener_urgencia(j: entero): AN
		Segun j hacer 
			1: Esc("Alta")
			2: Esc("Media")
			3: Esc("Baja")
		FinSegun
	FinFuncion
	Procedimiento urgencia() es
		segun reg.urgencia hacer 
			"A": j:=1
			"M": j:=2
			"B": j:=3
		FinSegun 
	FinProcedimiento
	reclamos: arreglo[1...13,1...4] de enteros
	Proceso 
		Abrir E/(arch)
		Leer(arch,reg)
		i:=0; j:=0; mayor:=0; mes_may:=0, c_meses:=0 
		Para i:= 1 a 13 hacer 
			Para j:= 1 a 4 hacer 
				reclamos[i,j]:=0 
			FinPara
		FinPara 
		Mientras NFDA(arch) hacer 
			i:= reg.f_reclamo.mes
			urgencia()
			reclamos[i,j]:= reclamos[i,j] + 1 
			reclamos[13,j]:= reclamos[13,j] + 1 
			reclamos[i,4]:= reclamos[i,4] + 1 
			Leer(arch,reg)
		FinMientras
		Para i:= 1 a 12 hacer
			Para j:= 1 a 3 hacer 
				Esc("Mes:",i,"Urgencia",obtener_urgencia(j),"reclamos:",reclamos[i,j])
			 	Si reclamos[i,j] > mayor entonces 
					mayor:= reclamos[i,j] 
					mes_may:= i 
				FinSi
				Si reclamos[i,2] < 10 entonces 
				 	c_meses:= c_meses + 1 
				FinSi
			FinPara
		FinPara
		Esc("Reclamos en enero:",reclamos[1,1])
		Esc("Reclamos de urgencia Alta",reclamos[13,1],"Media",reclamos[13,2],"Baja",reclamos[13,3])
		Esc("Mes con mas reclamos",mes_may)
		Esc("Cantidad de meses con menos de 10 reclamos de urgencia media:",c_meses)
		Cerrar(arch)
	FinProceso 
FinAccion

	 
